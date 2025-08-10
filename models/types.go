package models

type GasStation struct {
	FuelTypes []Fuel
}
type Fuel struct {
	Name       string
	Price      uint
	AmountFuel uint
}

type Car struct {
	Model      string
	FuelType   int
	MaxGasTank uint
	CurrentGas uint
}
type Driver struct {
	Name   string
	Car    *Car
	Wallet Wallet
}

type Wallet struct {
	BalanceWallet uint
}
