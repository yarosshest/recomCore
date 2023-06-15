package main

import (
	"fmt"
	db "recomCore/bdHandler"
	p "recomCore/parsing"
)

func main() {
	fmt.Print("db initialisation\n")
	db.DbInit()
	fmt.Print("parsing start\n")
	p.JsonParsing()
}
