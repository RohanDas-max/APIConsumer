package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rohandas-max/ghCrwaler/pkg/handler"
)

func main() {
	fmt.Println("Enter a github username")
	var username string
	fmt.Scan(&username)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	handler.Handler(ctx, username, 1*time.Second)
}
