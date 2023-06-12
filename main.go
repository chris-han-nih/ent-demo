package main

import (
	"context"
	"github.com/chris-han-nih/ent-demo/ent"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=ent password=1234qwer sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)
	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//user, err := CreateUser(context.Background(), client)
	//if err != nil {
	//	log.Fatalf("failed creating user: %v", err)
	//	return
	//}
	//log.Println("user was created: ", user)
	//
	//users, err := QueryUser(context.Background(), client)
	//if err != nil {
	//	log.Fatalf("failed querying user: %v", err)
	//}
	//for _, u := range users {
	//	log.Println("user was queried: ", u)
	//}
}
