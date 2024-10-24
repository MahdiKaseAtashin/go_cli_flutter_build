package main

import (
	"fmt"
	"lottery/cmd"
	"lottery/models"
)

func main() {
	users, _ := models.LoadFromFile("users.txt")
	fmt.Println(users)
	cmd.Execute()
}
