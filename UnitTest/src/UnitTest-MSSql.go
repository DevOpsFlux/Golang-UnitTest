package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

var server = "127.0.0.1" //"<your_server.database.windows.net>"
var port = 1433
var user = "dev"         //"<your_username>"
var password = "@NTdev"  //"<your_password>"
var database = "devdb" //"<your_database>"

func main() {
	test()
}

func test() {
	fmt.Println("UnitTest MS-SQL ")
	//fmt.Println(val)

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	fmt.Println("11111")

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	db.SetMaxOpenConns(10)
	if err != nil {
		fmt.Println("Error creating connection pool: ", err.Error())
	}

	fmt.Println("22222")

	if err != nil || db.Ping() != nil {
		fmt.Println("Ping Error : ", err.Error())
	}
	/*
		ctx := context.Background()
		fmt.Println("33333")
		err = db.PingContext(ctx)

		fmt.Println("44444")
		if err != nil {
			fmt.Println(err.Error())
		}
	*/
	fmt.Printf("Connected!\n")

	// Read employees
	count, err := ReadEmployees()
	if err != nil {
		fmt.Printf("Error reading Employees: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

	/*
		// Create employee
		createID, err := CreateEmployee("Jake", "United States")
		if err != nil {
			log.Fatal("Error creating Employee: ", err.Error())
		}
		fmt.Printf("Inserted ID: %d successfully.\n", createID)

		// Read employees
		count, err := ReadEmployees()
		if err != nil {
			log.Fatal("Error reading Employees: ", err.Error())
		}
		fmt.Printf("Read %d row(s) successfully.\n", count)

		// Update from database
		updatedRows, err := UpdateEmployee("Jake", "Poland")
		if err != nil {
			log.Fatal("Error updating Employee: ", err.Error())
		}
		fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

		// Delete from database
		deletedRows, err := DeleteEmployee("Jake")
		if err != nil {
			log.Fatal("Error deleting Employee: ", err.Error())
		}
		fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)
	*/
	/*
		// sql.DB 객체
		// db, err := sql.Open("mssql", "server=(local);user id=sa;password=pwd;database=pubs")
		db, err := sql.Open("mssql", "server=172.25.5.205;user id=N08704;password=@NTN08704;database=Ticket2000")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// 하나의 Row를 갖는 SQL 쿼리 : QueryRow()
		var lname, fname string
		err = db.QueryRow("SELECT * FROM TestSchema.Employees").Scan(&lname, &fname)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fname, lname)

		// 복수 Row를 갖는 SQL 쿼리 : Query()
		rows, err := db.Query("SELECT au_lname, au_fname FROM authors WHERE au_lname=?", "Ringer")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close() //반드시 닫는다 (지연하여 닫기)

		for rows.Next() {
			err := rows.Scan(&lname, &fname)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(fname, lname)
		}
	*/

}

// CreateEmployee inserts an employee record
func CreateEmployee(name string, location string) (int64, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		err = errors.New("CreateEmployee: db is null")
		return -1, err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `
      INSERT INTO TestSchema.Employees (Name, Location) VALUES (@Name, @Location);
      select isNull(SCOPE_IDENTITY(), -1);
    `

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", name),
		sql.Named("Location", location))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

// ReadEmployees reads all employee records
func ReadEmployees() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var name, location string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
		count++
	}

	return count, nil
}

// UpdateEmployee updates an employee's information
func UpdateEmployee(name string, location string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE TestSchema.Employees SET Location = @Location WHERE Name = @Name")

	// Execute non-query with named parameters
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Location", location),
		sql.Named("Name", name))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

// DeleteEmployee deletes an employee from the database
func DeleteEmployee(name string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name = @Name;")

	// Execute non-query with named parameters
	result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

/*
CREATE SCHEMA TestSchema;
GO

CREATE TABLE TestSchema.Employees (
  Id       INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name     NVARCHAR(50),
  Location NVARCHAR(50)
);
GO

INSERT INTO TestSchema.Employees (Name, Location) VALUES
  (N'Jared',  N'Australia'),
  (N'Nikita', N'India'),
  (N'Tom',    N'Germany');
GO

SELECT * FROM TestSchema.Employees;
GO
*/
