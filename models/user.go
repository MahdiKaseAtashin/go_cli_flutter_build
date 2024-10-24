package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type UserItem struct {
	ID          int           `json:"id"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	BirthDate   time.Time     `json:"birthDate"`
	PhoneNumber string        `json:"phoneNumber"`
	Field       WorkFieldItem `json:"field"`
	HasWon      bool          `json:"hasWon"`
	WinningDate time.Time     `json:"winningDate"`
}

func (u UserItem) String() string {
	return fmt.Sprintf("%v)\t%v\t%v\t%v\t%v\t%v\t%v\t%v", u.ID, u.FirstName, u.LastName, u.BirthDate.Format("2006-01-02"), u.PhoneNumber, u.Field, u.HasWon, u.WinningDate.Format("2006-01-02"))
}

func SaveToFile(users []UserItem, filename string) error {
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

func LoadFromFile(filename string) ([]UserItem, error) {
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
