package database

import (
	"database/sql"
	"fmt"
	"log"

	//Postgresql driver

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host = "database"
	port = 5432
)

func InitDB(user, password, dbname string) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db = database
	log.Println("Connected to the database!")
}

// func InsertContact(input model.NewContact) (*model.Contact, error) {
// 	stmt, err := db.Prepare(`INSERT INTO contacts VALUES(DEFAULT, $1, $2, $3, $4, $5, $6) RETURNING id;`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	currentTime := time.Now()

// 	var id int
// 	err = stmt.QueryRow(input.FirstName, input.LastName, input.Email, input.Phone, currentTime, currentTime).Scan(&id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	contact := model.Contact{
// 		ID:        id,
// 		FirstName: input.FirstName,
// 		LastName:  input.LastName,
// 		Email:     input.Email,
// 		Phone:     input.Phone,
// 		CreatedAt: currentTime,
// 		UpdatedAt: currentTime,
// 	}

// 	return &contact, nil
// }

// func GetAllContacts() ([]*model.Contact, error) {

// 	var contacts []*model.Contact

// 	rows, err := db.Query(`SELECT * FROM contacts`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var contact model.Contact

// 		err = rows.Scan(
// 			&contact.ID,
// 			&contact.FirstName,
// 			&contact.LastName,
// 			&contact.Email,
// 			&contact.Phone,
// 			&contact.CreatedAt,
// 			&contact.UpdatedAt,
// 		)
// 		if err != nil {
// 			log.Println("error 3")
// 			return nil, err
// 		}

// 		contacts = append(contacts, &contact)
// 	}

// 	return contacts, nil
// }

// func DeleteContactByID(id int) (*int, error) {
// 	// check if contact with given id exists
// 	if !CheckIfContactExists(id) {
// 		return nil, fmt.Errorf("contact with given id does not exists")
// 	}
// 	stmt, err := db.Prepare(`DELETE FROM contacts WHERE id = $1;`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result, err := stmt.Exec(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rowsAffected1, err := result.RowsAffected()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rowsAffected2 := int(rowsAffected1)

// 	return &rowsAffected2, nil
// }

// func CheckIfContactExists(id int) bool {
// 	err := db.QueryRow(`SELECT id FROM contacts WHERE id = $1`, id).Scan(&id)
// 	switch {
// 	case err == sql.ErrNoRows:
// 		return false
// 	case err != nil:
// 		log.Fatal(err)
// 		return false
// 	default:
// 		return true
// 	}
// }

// func UpdateContact(input model.UpdateContact) (*model.Contact, error) {
// 	// check if contact with given id exists
// 	if !CheckIfContactExists(input.ID) {
// 		return nil, fmt.Errorf("contact with given id does not exists")
// 	}

// 	stmt, err := db.Prepare(`UPDATE contacts SET firstName = $1, lastName = $2, email = $3, phone = $4, updatedAt = $5 WHERE id = $6 RETURNING createdAt;`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var createdAt time.Time
// 	currentTime := time.Now()

// 	err = stmt.QueryRow(input.FirstName, input.LastName, input.Email, input.Phone, currentTime, input.ID).Scan(&createdAt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	contact := model.Contact{
// 		ID:        input.ID,
// 		FirstName: input.FirstName,
// 		LastName:  input.LastName,
// 		Email:     input.Email,
// 		Phone:     input.Phone,
// 		CreatedAt: createdAt,
// 		UpdatedAt: currentTime,
// 	}

// 	return &contact, nil
// }
