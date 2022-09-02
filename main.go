package main

import (
	"SuperStar/routes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	app := routes.New()

	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))
	log.Fatal(app.Listen(":3000"))
}
