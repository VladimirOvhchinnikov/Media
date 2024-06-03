package models

type CashInfo struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}

type Response struct {
	Exchanges [][]int `json:"exchanges"`
}
