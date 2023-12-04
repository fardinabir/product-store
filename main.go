package main

import (
	_ "github.com/fardinabir/product-store/docs"
	"github.com/fardinabir/product-store/server"
)

func main() {

	server.StartServer("8085")
}
