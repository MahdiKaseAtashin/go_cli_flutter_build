package main

import (
	"fmt"
)

func main() {
	users, _ := loadFromFile("users.txt")
	fmt.Println(users)
}
