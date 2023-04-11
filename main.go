package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Notes struct {
	Id        int
	Noteslist string
}

func main() {

	db, err := sql.Open("mysql", "root:notes@(127.0.0.1:13306)/notes")

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM noted")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var note Notes
		err := res.Scan(&note.Id, &note.Noteslist)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%b\n", note)
	}

	sql := "insert into noted (ID, Noteslist) values (5, 'This is my fifth note from GOapp');"
	res2, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res2.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)

}
