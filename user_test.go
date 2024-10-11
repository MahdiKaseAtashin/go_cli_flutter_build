package main

import (
	"fmt"
	"testing"
)

func TestSaveToFile(t *testing.T) {
	tests := []struct {
		name     string
		items    []UserItem
		filename string
	}{
		{
			name: "test save user to file",
			items: []UserItem{
				{Index: 1, FirstName: "John", LastName: "Doe", Age: 30, PhoneNumber: "1234567890", Field: Mobile, WonBefore: false},
				{Index: 2, FirstName: "Jane", LastName: "Doe", Age: 25, PhoneNumber: "0987654321", Field: Web, WonBefore: true},
				{Index: 3, FirstName: "Bob", LastName: "Smith", Age: 40, PhoneNumber: "5555555555", Field: Backend, WonBefore: true},
				{Index: 4, FirstName: "Alice", LastName: "Johnson", Age: 35, PhoneNumber: "4444444444", Field: DevOps, WonBefore: false},
				{Index: 5, FirstName: "Mike", LastName: "Brown", Age: 50, PhoneNumber: "3333333333", Field: Manager, WonBefore: true},
				{Index: 6, FirstName: "Emily", LastName: "Davis", Age: 45, PhoneNumber: "2222222222", Field: Owner, WonBefore: false},
				{Index: 7, FirstName: "David", LastName: "Miller", Age: 60, PhoneNumber: "1111111111", Field: Mobile, WonBefore: true},
			},
			filename: "users.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := saveToFile(tc.items, tc.filename)
			if err != nil {
				t.Fatalf("error saving file: %v", err)
			}
		})
	}
}

func TestLoadFromFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{name: "test load user from file", filename: "users.json"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			users, err := loadFromFile(tc.filename)
			if err != nil {
				t.Fatalf("error loading file: %v", err)
			}
			for _, user := range users {
				fmt.Println(user)
			}
		})
	}
}
