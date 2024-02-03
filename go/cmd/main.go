package main

import (
	"fmt"
	"go/internal/osInterface"
	"go/internal/shredder"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cmd/main.go <file_path>")
		return
	}

	filePath := os.Args[1]
	fmt.Printf("Shredding %s\n", filePath)
	realOs := &osInterface.RealOs{}
	shredder := shredder.NewShredder(realOs)
	err := shredder.Shred(filePath)

	if err != nil {
		fmt.Printf("error performing operation %s", err.Error())
	} else {
		fmt.Println("operation performed successfully")
	}
}
