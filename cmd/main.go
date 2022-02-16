package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rohandas-max/ghCrwaler/pkg/controller"
)

func main() {
	fmt.Println("Enter a github username")
	var username string
	fmt.Scan(&username)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()
	if err := controller.Controller(ctx, username); err != nil {
		log.Fatal(err)
	}
}
