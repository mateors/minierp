package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/mateors/msql"
)

func addClient(name, mobile, email, address string) (int64, error) {

	//INSERT INTO client (name,mobile,email,address) VALUES("RAIHAN", "01672710028", "raihan@gmail.com", "Kustia");
	//sql := fmt.Sprintf(`INSERT INTO client (name,mobile,email,address) VALUES("%s", "%s", "%s", "%s");`, name, mobile, email, address)
	//fmt.Println(sql)

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
