package main

import (
	"fmt"

	"k8s.io/client-go/rest"
)

func main() {
	fmt.Println("bonjour")
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
}
