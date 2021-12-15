package main

import (
	"GoProject_1/repositories/filesystem"
	"fmt"
)

func main() {
	repo := filesystem.UserFileRepository{}

	repo.GetByEmail("")

	fmt.Println(repo)
}
