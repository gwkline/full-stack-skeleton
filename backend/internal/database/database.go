package database

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gwkline/full-stack-infra/backend/internal/graph/model"

	//Postgresql driver

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase(conn *sql.DB) *Database {
	return &Database{Conn: conn}
}

const (
	host          = "database"
	port          = 5432
	maxRetries    = 10
	retryInterval = 5 * time.Second
)

func InitDB(user, password, dbname string) *Database {
	fmt.Println("Initializing database")
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	if os.Getenv("ENV") == "development" {
		fmt.Println(psqlInfo)
	}

	db, err := waitForDatabase(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.RunMigrations()
	log.Println("Connected to the database")

	return db
}

func waitForDatabase(psqlInfo string) (*Database, error) {
	fmt.Println("Waiting for database connection")

	for i := 0; i < maxRetries; i++ {
		dbConn, err := sql.Open("postgres", psqlInfo)
		db := NewDatabase(dbConn)
		if err == nil {
			err = db.Conn.Ping()
			if err == nil {
				return db, nil
			}
		}

		log.Printf("Failed to connect to database. Retry %d/%d. Waiting for %v before retrying...", i+1, maxRetries, retryInterval)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("failed to connect to database after %d tries", maxRetries)
}

// IsDirEmpty checks if a directory is empty
func IsDirEmpty(dir string) (bool, error) {
	// Open the directory
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// Read directory contents
	_, err = f.Readdir(1)

	// Check for EOF, which means the directory is empty
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func (d *Database) RunMigrations() {
	isEmpty, err := IsDirEmpty("./internal/database/migrations")
	if err != nil {
		os.Mkdir("./internal/database/migrations", os.ModePerm)
		_, err := IsDirEmpty("./internal/database/migrations")
		if err != nil {
			log.Fatalf("Failed to check directory (maybe missing): %v", err)

		} else {
			fmt.Println("No migrations directory found, but was successfully created")
		}
		return
	}

	if isEmpty {
		fmt.Println("Directory is empty - skipping migrations")
		return
	}
	driver, _ := postgres.WithInstance(d.Conn, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/database/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal("error creating migrations instance: ", err)
	}

	// Only call m.Up() if m isn't nil
	if m != nil {
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("error running migrations: ", err)
		}
	}
}

func (d *Database) InsertUser(input model.NewUser) (*model.User, error) {
	stmt, err := d.Conn.Prepare(`
	INSERT INTO users (email, passwordHash, otpSecret, phone, createdAt, updatedAt) 
	VALUES($1, $2, $3, $4, TIMESTAMP 'epoch' + $5 * INTERVAL '1 second', TIMESTAMP 'epoch' + $6 * INTERVAL '1 second') 
	RETURNING id;
	`)
	if err != nil {
		return nil, err
	}

	currentTime := int(time.Now().UnixMilli())

	var id int
	err = stmt.QueryRow(input.Email, input.Password, input.Otp, input.Phone, currentTime, currentTime).Scan(&id)
	if err != nil {
		return nil, err
	}
	user := model.User{
		ID:        strconv.Itoa(id),
		Email:     input.Email,
		Phone:     input.Phone,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	return &user, nil
}

func (d *Database) GetAllUsers() ([]*model.User, error) {

	var users []*model.User

	rows, err := d.Conn.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("error 3")
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (d *Database) DeleteUserByID(id string) (*int, error) {
	// check if user with given id exists
	_, err := d.FindUser(id)
	if err != nil {
		return nil, fmt.Errorf("user with given id does not exists")
	}
	stmt, err := d.Conn.Prepare(`DELETE FROM users WHERE id = $1;`)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	rowsAffected1, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	rowsAffected2 := int(rowsAffected1)

	return &rowsAffected2, nil
}

func (d *Database) FindUser(identifier string, colName ...string) (model.User, error) {

	var col string
	if len(colName) > 0 {
		col = colName[0]
	} else {
		col = "id"
	}

	queryString := fmt.Sprintf(`SELECT id, email, passwordHash, otpSecret, phone, 
	EXTRACT(EPOCH FROM createdAt)::bigint, 
	EXTRACT(EPOCH FROM updatedAt)::bigint 
FROM users WHERE %s = $1`, col)

	var user model.User
	err := d.Conn.QueryRow(queryString, identifier).Scan(&user.ID, &user.Email, &user.Password, &user.OtpSecret, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return model.User{}, fmt.Errorf("no user found")
	case err != nil:
		log.Fatal(err)
		return model.User{}, fmt.Errorf("no user found")
	default:
		return user, nil
	}
}

func (d *Database) UpdateUser(input model.User) (*model.User, error) {
	// check if user with given id exists
	_, err := d.FindUser(input.ID)
	if err != nil {
		return nil, fmt.Errorf("user with given id does not exists")
	}

	stmt, err := d.Conn.Prepare(`UPDATE users SET email = $1, phone = $2, otpSecret = $3, updatedAt = NOW() WHERE id = $4 RETURNING createdAt;`)
	if err != nil {
		return nil, err
	}

	var createdAt time.Time
	currentTime := int(time.Now().Unix())

	err = stmt.QueryRow(input.Email, input.Phone, input.OtpSecret, input.ID).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	intCreatedAt := int(createdAt.UnixMilli())

	user := model.User{
		ID:        input.ID,
		Email:     input.Email,
		Phone:     input.Phone,
		CreatedAt: intCreatedAt,
		UpdatedAt: currentTime,
	}

	return &user, nil
}
