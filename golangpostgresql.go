package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname = "127.0.0.1"
	db       = "apache_age"
	username = "postgres"
	pwd      = "Admin"
	portID   = 5432
)

var conn *sql.DB

type StaffRecord struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Phone  string `json:"phone"`
}

func createConnection() (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostname, portID, username, pwd, db)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	fmt.Println("Opened database successfully")
	return db, nil
}

func createDBTable() error {
	cur, err := conn.Query("DROP TABLE IF EXISTS STAFF_RECORD")
	if err != nil {
		return err
	}
	defer cur.Close()

	sql := `CREATE TABLE STAFF_RECORD(
		USER_ID SERIAL PRIMARY KEY,
		NAME CHAR(20) NOT NULL,
		AGE INT,
		PHONE CHAR(20)
	)`

	cur, err = conn.Query(sql)
	if err != nil {
		return err
	}
	defer cur.Close()

	fmt.Println("Table created successfully........")
	return nil
}

func insertDataIntoDB() error {
	createDBTable()

	insertQuery := `INSERT INTO STAFF_RECORD (NAME, AGE, PHONE) VALUES ($1, $2, $3)`
	records := []StaffRecord{
		{Name: "Jenny", Age: 34, Phone: "091128282"},
		{Name: "Tom", Age: 29, Phone: "1-800-123-1234"},
		{Name: "John", Age: 28, Phone: ""},
	}

	for _, record := range records {
		_, err := conn.Exec(insertQuery, record.Name, record.Age, record.Phone)
		if err != nil {
			return err
		}
		fmt.Printf("Record for %s........ was created successfully\n", record.Name)
	}

	return nil
}

func queryDB(query string, args ...interface{}) ([]StaffRecord, error) {
	rows, err := conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []StaffRecord
	for rows.Next() {
		var record StaffRecord
		err := rows.Scan(&record.UserID, &record.Name, &record.Age, &record.Phone)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	return result, nil
}

func getResultInJSON() (string, error) {
	queryResult, err := queryDB("SELECT * FROM public.staff_record")
	if err != nil {
		return "", err
	}

	jsonOutput, err := json.Marshal(queryResult)
	if err != nil {
		return "", err
	}

	return string(jsonOutput), nil
}

func main() {
	var err error
	conn, err = createConnection()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer conn.Close()

	err = insertDataIntoDB()
	if err != nil {
		fmt.Println("Failed to insert records into staff record table:", err)
		return
	}

	jsonOutput, err := getResultInJSON()
	if
