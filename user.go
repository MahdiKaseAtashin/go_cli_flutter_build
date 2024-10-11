package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type UserItem struct {
	Index       int           `json:"Index"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	Age         int           `json:"age"`
	PhoneNumber string        `json:"phoneNumber"`
	Field       WorkFieldItem `json:"field"`
	WonBefore   bool          `json:"wonBefore"`
}

func (u UserItem) String() string {
	return fmt.Sprintf("%v)\t%v\t%v\t%v\t%v\t%v\t%v", u.Index, u.FirstName, u.LastName, u.Age, u.PhoneNumber, u.Field, u.WonBefore)
}

func saveToFile(users []UserItem, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	j, err := json.Marshal(users)
	if err != nil {
		return err
	}
	_, err = f.Write(j)
	if err != nil {
		return err
	}
	return nil
}

func loadFromFile(filename string) ([]UserItem, error) {
	var users []UserItem
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	d := json.NewDecoder(f)
	err = d.Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getWinner(users []UserItem) UserItem {
	return users[rand.Int31n(int32(len(users)))]
}
