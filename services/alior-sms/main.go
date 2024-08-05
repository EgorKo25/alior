package main

import (
	"alior-sms/src/database"
	"alior-sms/src/types"
	"context"
	"log"
)

func main() {

	ctx := context.Background()
	connString := ""
	migrationDir := "."
	DB, err := database.NewDB(ctx, connString, migrationDir)
	if err != nil {
		log.Fatal(err)
	}
	ser := &types.Service{ID: 1}
	id, err := DB.InsertService(ctx, ser)
	if err != nil {
		log.Fatal(err)
	}
	ser2, err := DB.GetServiceByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ser2.Name)
}
