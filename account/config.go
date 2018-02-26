package account

import (
	"errors"
	"reflect"

	"github.com/spf13/viper"
	//"fmt"
)

func mapAccountsByField(field string) map[string]Account {
	var config Accounts

	viper.Unmarshal(&config)

	confMap := map[string]Account{}
	for _, v := range config.Accounts {
		r := reflect.ValueOf(v)

		if r.FieldByName(field).IsValid() == true {
			f := reflect.Indirect(r).FieldByName(field)
			value := f.Interface().(string)
			confMap[value] = v
		}
	}
	return confMap
}


// func FindOneBy(field string) (Account, error) {

// }

//GetBaseAccounts Return all base accounts
// func GetBaseAccounts() []BaseAccount {

// 	var baseAccount []BaseAccount

// 	err := viper.UnmarshalKey("base_accounts", &baseAccount)

// 	if err != nil {
// 		panic("Unable to unmarshal hosts")
// 	}

// 	return baseAccount
// }

// //GetCrossAccounts Return all base accounts
// func GetCrossAccounts() []CrossAccount {

// 	err := viper.UnmarshalKey("cross_accounts", &account)

// 	if err != nil {
// 		panic("Unable to unmarshal hosts")
// 	}

// 	return account
// }
func WriteAccountToConfig(account Account) bool {

	var config Accounts

	viper.Unmarshal(&config)
	newAccounts := append(config.Accounts, account)

	//config.Accounts = newAccounts

	viper.Set("accounts", newAccounts)
	err := viper.WriteConfig()

	if err != nil {
		return false
	}
	return true
}

// FindAllbyType return accounts by type
func FindAllbyType(accountType string) ([]Account, error) {

	confMap := mapAccountsByField("ProfileName")

	if allowedTypes(accountType) == false {
		return nil, errors.New("invalid account type")
	}

	var accounts []Account

	for _, a := range confMap {
		if isBaseAccount(a) == true && accountType == "base" {
			accounts = append(accounts, a)
		} else if isBaseAccount(a) == false && accountType == "cross" {
			accounts = append(accounts, a)
		}
	}

	return accounts, nil
}

func allowedTypes(accountType string) bool {
	if accountType == "base" || accountType == "cross" {
		return true
	}

	return false
}

func isBaseAccount(account Account) bool {
	if account.SourceProfile == "" {
		return true
	}
	return false
}