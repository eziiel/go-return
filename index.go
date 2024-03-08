package main

import (
	"./teste" // Alteração do caminho absoluto para relativo
	"fmt"
)

func main() {
	fmt.Printf("Hello, world")
	teste.Test()
}