package models

type AccountHolder struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	JobTitle string `json:"jobTitle"`
	Updated  int    `json:"updated"`
	Created  int    `json:"created"`
}

//addAccount
//removeAccount
//listAccounts
//getAccountHolderDetails
