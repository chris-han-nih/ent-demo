package main

import (
	"context"
	"github.com/chris-han-nih/ent-demo/ent"
	"github.com/chris-han-nih/ent-demo/ent/migrate"
	"github.com/chris-han-nih/ent-demo/model"
	"log"
	"net/url"

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
	if err = client.Schema.Create(context.Background(), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	user := &model.User{
		Age:    30,
		Rank:   1.0,
		Active: true,
		Name:   "yougjae",
		Url: &url.URL{
			Scheme: "http",
			Host:   "localhost",
		},
		State: "on",
	}

	id, err := user.Create(context.Background(), client)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
		return
	}
	log.Println("user was created: ", id)

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
