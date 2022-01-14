package main

import "fmt"

func addClient(name, mobile, email, address string) int64 {

	//INSERT INTO client (name,mobile,email,address) VALUES("RAIHAN", "01672710028", "raihan@gmail.com", "Kustia");
	sql := fmt.Sprintf(`INSERT INTO client (name,mobile,email,address) VALUES("%s", "%s", "%s", "%s");`, name, mobile, email, address)
	fmt.Println(sql)

	return 0
}
