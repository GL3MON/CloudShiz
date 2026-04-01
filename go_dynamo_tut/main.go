package main

import (
	"dynamo-test/db"
)

func main() {

	client, err := db.NewClient()
	if err != nil {
		panic(err)
	}

	// err = db.CreateTable(client)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Putting item...")
	// db.PutItem(client)

	// fmt.Println("Getting item...")
	// db.GetItem(client)

	// fmt.Println("Updating item...")
	// db.UpdateItem(client)

	// fmt.Println("done")

	// fmt.Println("Deleting item...")
	// db.DeleteItem(client)

	db.DeleteTable(client)
}
