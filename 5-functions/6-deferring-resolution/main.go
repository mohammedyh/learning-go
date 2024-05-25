package main

import "fmt"

func queryDatabase(query string) string {
	var result string
	connectDatabase()

	defer disconnectDatabase()

	if query == "SELECT * FROM coolTable;" {
		result = "NAME\t\t\t|\tDOB\nVincent Van Gogh\t|\tMarch 30, 1853"
	}
	fmt.Println(result)
	return result
}

func connectDatabase() {
	fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
	fmt.Println("Disconnecting from the database.")
}

func main() {
	queryDatabase("SELECT * FROM coolTable;")
}
