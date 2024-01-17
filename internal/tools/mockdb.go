package tools

import (
  "time"
)

type mockDB struct {

}

var mockLoginDetails = map[string]LoginDetails{
  "ian": {
    AuthToken: "123ABC",
    Username: "ian",
  },
  "jason": {
    AuthToken: "ABC123",
    Username: "jason",
  },
  "alex": {
    AuthToken: "XYZ123",
    Username: "alex",
  },
}

var mockCoinDetails = map[string]CoinDetails{
  "ian": {
    Coins: 100,
    Username: "ian",
  },
  "jason": {
    Coins: 200,
    Username: "jason",
  },
  "alex": {
    Coins: 300,
    Username: "alex",
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

