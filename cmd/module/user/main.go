package main

import (
	"fmt"
	"net/http"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
)

func main() {
	config := config.GetConfig()
	fmt.Printf("App name: %v\n", config)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
