package database

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gwkline/full-stack-infra/backend/graph/model"

	//Postgresql driver

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host          = "database"
	port          = 5432
	maxRetries    = 10
	retryInterval = 5 * time.Second
)

func InitDB(user, password, dbname string) {
	fmt.Println("Initializing database")
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	if os.Getenv("ENV") == "development" {
		fmt.Println(psqlInfo)
	}

	database, err := waitForDatabase(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db = database
	RunMigrations()
	log.Println("Connected to the database")
}

func waitForDatabase(psqlInfo string) (*sql.DB, error) {
	var err error

	fmt.Println("Waiting for database connection")

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err == nil {
			err = db.Ping()
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

func RunMigrations() {
	isEmpty, err := IsDirEmpty("./database/migrations")
	if err != nil {
		log.Fatalf("Failed to check directory (maybe missing): %v", err)
		return
	}

	if isEmpty {
		fmt.Println("Directory is empty - skipping migrations")
		return
	}
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/migrations",
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

func InsertUser(input model.NewUser) (*model.User, error) {
	stmt, err := db.Prepare(`
	INSERT INTO users (email, passwordHash, otpSecret, phone, createdAt, updatedAt) 
	VALUES($1, $2, $3, $4, TIMESTAMP 'epoch' + $5 * INTERVAL '1 second', TIMESTAMP 'epoch' + $6 * INTERVAL '1 second') 
	RETURNING id;
	`)
	if err != nil {
		return nil, err
	}

	currentTime := int(time.Now().Unix())

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

func GetAllUsers() ([]*model.User, error) {

	var users []*model.User

	rows, err := db.Query(`SELECT * FROM users`)
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

func DeleteUserByID(id string) (*int, error) {
	// check if user with given id exists
	_, err := FindUser(id)
	if err != nil {
		return nil, fmt.Errorf("user with given id does not exists")
	}
	stmt, err := db.Prepare(`DELETE FROM users WHERE id = $1;`)
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

func FindUser(identifier string, colName ...string) (model.User, error) {

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
	fmt.Println(queryString)

	var user model.User
	err := db.QueryRow(queryString, identifier).Scan(&user.ID, &user.Email, &user.Password, &user.OtpSecret, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return model.User{}, fmt.Errorf("no user found")
	case err != nil:
		log.Fatal(err)
		return model.User{}, fmt.Errorf("no user found")
	default:
		fmt.Println(user)
		return user, nil
	}
}

func UpdateUser(input model.User) (*model.User, error) {
	// check if user with given id exists
	_, err := FindUser(input.ID)
	if err != nil {
		return nil, fmt.Errorf("user with given id does not exists")
	}

	stmt, err := db.Prepare(`UPDATE users SET email = $1, phone = $2, updatedAt = $3 WHERE id = $4 RETURNING createdAt;`)
	if err != nil {
		return nil, err
	}

	var createdAt int
	currentTime := int(time.Now().Unix())

	err = stmt.QueryRow(input.Email, input.Phone, currentTime, input.ID).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:        input.ID,
		Email:     input.Email,
		Phone:     input.Phone,
		CreatedAt: createdAt,
		UpdatedAt: currentTime,
	}

	return &user, nil
}
