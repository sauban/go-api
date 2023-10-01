package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"sauban": {
		AuthToken: "1234",
		Username:  "sauban",
	},
	"amina": {
		AuthToken: "123FE4",
		Username:  "amina",
	},
	"seilat": {
		AuthToken: "123PPE",
		Username:  "seilat",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"sauban": {
		Coins:    100,
		Username: "sauban",
	},
	"amina": {
		Coins:    200,
		Username: "amina",
	},
	"seilat": {
		Coins:    300,
		Username: "seilat",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
