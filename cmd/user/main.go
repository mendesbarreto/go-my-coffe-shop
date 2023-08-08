package main

import (
	"fmt"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/user/config"
)

func main() {
	config := config.GetConfig()

	fmt.Printf("App name: %v\n", config)
}
