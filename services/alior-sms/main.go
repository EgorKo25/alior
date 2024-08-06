package main

import (
	"alior-sms/src/config"
	"alior-sms/src/database"
	"alior-sms/src/types"
	"context"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	connString := config.BuildConnString(&cfg.Databases.SMSDatabase)

	ctx := context.Background()

	migrationDir := "."

	DB, err := database.NewDB(ctx, connString, migrationDir)
	if err != nil {
		log.Fatal(err)
	}

	ser := &types.Service{ID: 1, Name: "serv1", Description: "desc1", Price: 2}
	id, err := DB.InsertService(ctx, ser)

	if err != nil {
		log.Fatal(err)
	}

	ser2, err := DB.GetServiceByID(ctx, id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(ser2.Name)
	servs, err := DB.GetPaginatedServices(ctx, 2, 3)

	if err != nil {
		log.Fatal(err)
	}

	for _, service := range servs {
		log.Println(service.Name, service.ID)
	}
}
