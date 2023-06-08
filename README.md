# Golang-with-postgresql
This code is a Go program that interacts with a PostgreSQL database using the database/sql package and the github.com/lib/pq driver. It performs the following tasks:

Imports the necessary packages: database/sql for database connectivity, encoding/json for JSON serialization, and _ "github.com/lib/pq" to import the PostgreSQL driver without directly referencing it.

Defines constants for the database connection parameters: hostname, db, username, pwd, and portID.

Declares a global variable conn of type *sql.DB to store the database connection.

Defines a struct type StaffRecord that represents the structure of a record in the STAFF_RECORD table. The struct fields are annotated with json tags to specify the JSON key names for serialization.

Implements a function createConnection() that establishes a connection to the PostgreSQL database using the provided connection parameters. It constructs the connection string and opens a new connection using sql.Open().

Implements a function createDBTable() that drops the STAFF_RECORD table if it already exists and creates a new table with the specified columns using conn.Query().

Implements a function insertDataIntoDB() that inserts records into the STAFF_RECORD table. It calls createDBTable() to ensure the table exists and then executes an INSERT query for each record using conn.Exec().

Implements a function queryDB() that executes a query on the database and retrieves the result as a slice of StaffRecord structs. It iterates over the result set using rows.Next() and scans the values into a new StaffRecord struct using rows.Scan().

Implements a function getResultInJSON() that queries the STAFF_RECORD table using queryDB() and converts the result into a JSON string using json.Marshal().

In the main() function:

Calls createConnection() to establish the database connection and assigns it to the global variable conn.
Defers the closure of the database connection using defer conn.Close() to ensure it's closed after the main function finishes.
Calls insertDataIntoDB() to insert records into the table.
Calls getResultInJSON() to retrieve the contents of the table as a JSON string.
The code handles errors throughout the process and prints relevant error messages if any errors occur.
