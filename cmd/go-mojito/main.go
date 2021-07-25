package main

import (
	"fmt"

	"github.com/anookinov/go-mojito/pkg/server"
)

func main() {
	fmt.Printf("Go Mojito!!")
	r := server.SetupRouter()
	r.Run() //listen and serve on 0.0.0.0:8080
}