package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/mateors/msql"
)

// client add query
func AddClient(name, mobile, email, address string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "client")
	data.Set("dbtype", "sqlite3")
	data.Set("name", name)
	data.Set("mobile", mobile)
	data.Set("email", email)
	data.Set("address", address)

	pid, err := msql.InsertIntoAnyTable(data, db)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	fmt.Println("Successfully inserted", pid)
	return pid, nil
}

// client get
func GetClient() []map[string]interface{} {
	qs := "SELECT * FROM client;"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		panic(err)
	}
	return rows
}
