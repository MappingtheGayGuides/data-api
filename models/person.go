package models

import (
	"database/sql"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./mgg.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
//names must be capitalized to be exported.
type Person struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	StreetAddress string `json:"street_address"`
	City string `json:"city"`
	State string `json:"state"`
	Amenities string `json:"amenity_features"`
	LocationType string `json:"location_type"`
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
	Status string `json:"status"`
	Year  int `json:"Year"`
}

func GetPersons(count int) ([]Person, error) {

	rows, err := DB.Query("SELECT ID, title, description, streetaddress, city, state, amenityfeatures, type, lon, lat, status, Year from mgg LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	people := make([]Person, 0)
//each name must be capitalized to be exported.
	for rows.Next() {
		singlePerson := Person{}
		err = rows.Scan(&singlePerson.ID, &singlePerson.Title, &singlePerson.Description, &singlePerson.StreetAddress, &singlePerson.City, &singlePerson.State, &singlePerson.Amenities, &singlePerson.LocationType, &singlePerson.Lon, &singlePerson.Lat, &singlePerson.Status, &singlePerson.Year)

		if err != nil {
			return nil, err
		}

		people = append(people, singlePerson)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return people, err
}

//this function gets a location by its ID
func GetPersonById(ID string) (Person, error) {
	//always prepeare your statement so that it can fail before any sql is actually run. More secure this way.
	stmt, err := DB.Prepare("SELECT ID, title, description, streetaddress, city, state, amenityfeatures, type, lon, lat, status, Year from mgg WHERE ID = ?")

	//return an empty person if there is an error.
	if err != nil {
		return Person{}, err
	}
	//create an instance of the person struct -- this gets returned at end of function.
	person := Person{}

	//pass the prepared statement to QueryRow. Passes in the id as a parameter and then scans in the row that is returned to the person instance that we created earlier.
	sqlErr := stmt.QueryRow(ID).Scan(&person.ID, &person.Title, &person.Description, &person.StreetAddress, &person.City, &person.State, &person.Amenities, &person.LocationType, &person.Lon, &person.Lat, &person.Status, &person.Year)

	//this is error checking. One assumption that we make is if the query returns 0 rows that isn't necessarily an "error". sqlErr == sql.ErrNoRow catches that and returns a blank person.
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Person{}, nil
		}
		return Person{}, sqlErr
	}
	return person, nil
}
//
// func AddPerson(newPerson Person) (bool, error) {
//
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return false, err
// 	}
//
// 	stmt, err := tx.Prepare("INSERT INTO people (first_name, last_name, email, ip_address) VALUES (?, ?, ?, ?)")
//
// 	if err != nil {
// 		return false, err
// 	}
//
// 	defer stmt.Close()
//
// 	_, err = stmt.Exec(newPerson.FirstName, newPerson.LastName, newPerson.Email, newPerson.IpAddress)
//
// 	if err != nil {
// 		return false, err
// 	}
//
// 	tx.Commit()
//
// 	return true, nil
// }
//
// func UpdatePerson(ourPerson Person, id int) (bool, error) {
//
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return false, err
// 	}
//
// 	stmt, err := tx.Prepare("UPDATE people SET first_name = ?, last_name = ?, email = ?, ip_address = ? WHERE Id = ?")
//
// 	if err != nil {
// 		return false, err
// 	}
//
// 	defer stmt.Close()
//
// 	_, err = stmt.Exec(ourPerson.FirstName, ourPerson.LastName, ourPerson.Email, ourPerson.IpAddress, id)
//
// 	if err != nil {
// 		return false, err
// 	}
//
// 	tx.Commit()
//
// 	return true, nil
// }
//
// func DeletePerson(personId int) (bool, error) {
//
// 	tx, err := DB.Begin()
//
// 	if err != nil {
// 		return false, err
// 	}
//
// 	stmt, err := DB.Prepare("DELETE from people where id = ?")
//
// 	if err != nil {
// 		return false, err
// 	}
//
// 	defer stmt.Close()
//
// 	_, err = stmt.Exec(personId)
//
// 	if err != nil {
// 		return false, err
// 	}
//
// 	tx.Commit()
//
// 	return true, nil
// }
