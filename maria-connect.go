package main

import (
	"database/sql"
	"fmt"
//	"github.com/go-sql-driver/mysql"
         _ "github.com/go-sql-driver/mysql"

	"golang.org/x/crypto/ssh"
)

type ViaSSHDialer struct {
	client *ssh.Client
}


func main() {

	dbUser := "root"         // DB username
	dbPass := "XXXXX"         // DB Password
	dbHost := "localhost:32969" // DB Hostname/IP
	dbName := "test"       // Database name



		// And now we can use our new driver with the regular mysql connection string tunneled through the SSH connection
		if db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)); err == nil {

			fmt.Printf("Successfully connected to the db\n")

			if rows, err := db.Query("SELECT name, age FROM team"); err == nil {
				for rows.Next() {
					var name string
                                        var age int64
					rows.Scan(&name,&age)
					fmt.Printf("Name: %s Age: %d\n", name, age)
				}
				rows.Close()
			} else {
				fmt.Printf("Failure: %s", err.Error())
			}

			db.Close()

		} else {

			fmt.Printf("Failed to connect to the db: %s\n", err.Error())
		}

}
