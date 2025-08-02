package models

type Fuel struct {
	Name       string
	Price      float64
	AmountFuel float64
}

type GasStation struct {
	FuelTypes []Fuel
}

type Car struct {
	Model      string
	FuelType   int
	MaxGasTank int
	CurrentGas int
	Driver     Driver
}

type Driver struct {
	Name          string
	BalanceWallet float64
}
