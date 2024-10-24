package models

import (
	"fmt"
	"testing"
	"time"
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
				{ID: 1, FirstName: "John", LastName: "Doe", BirthDate: time.Date(1993, time.January, 1, 0, 0, 0, 0, time.UTC), PhoneNumber: "1234567890", Field: Mobile, HasWon: false, WinningDate: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
				{ID: 2, FirstName: "Jane", LastName: "Doe", BirthDate: time.Date(1998, time.February, 2, 0, 0, 0, 0, time.UTC), PhoneNumber: "0987654321", Field: Web, HasWon: true, WinningDate: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)},
				{ID: 3, FirstName: "Bob", LastName: "Smith", BirthDate: time.Date(1983, time.March, 3, 0, 0, 0, 0, time.UTC), PhoneNumber: "5555555555", Field: Backend, HasWon: true, WinningDate: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)},
				{ID: 4, FirstName: "Alice", LastName: "Johnson", BirthDate: time.Date(1988, time.April, 4, 0, 0, 0, 0, time.UTC), PhoneNumber: "4444444444", Field: DevOps, HasWon: false},
				{ID: 5, FirstName: "Mike", LastName: "Brown", BirthDate: time.Date(1973, time.May, 5, 0, 0, 0, 0, time.UTC), PhoneNumber: "3333333333", Field: Manager, HasWon: true, WinningDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)},
				{ID: 6, FirstName: "Emily", LastName: "Davis", BirthDate: time.Date(1978, time.June, 6, 0, 0, 0, 0, time.UTC), PhoneNumber: "2222222222", Field: Owner, HasWon: false},
				{ID: 7, FirstName: "David", LastName: "Miller", BirthDate: time.Date(1963, time.July, 7, 0, 0, 0, 0, time.UTC), PhoneNumber: "1111111111", Field: Mobile, HasWon: true, WinningDate: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
			filename: "users.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := SaveToFile(tc.items, tc.filename)
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
			users, err := LoadFromFile(tc.filename)
			if err != nil {
				t.Fatalf("error loading file: %v", err)
			}
			for _, user := range users {
				fmt.Println(user)
			}
		})
	}
}
