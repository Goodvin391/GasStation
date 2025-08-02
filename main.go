package main

import (
	"bufio"
	"fmt"
	getData "gasstation/getdata"
	"gasstation/models"
	"os"
	"strconv"
)

func newGasStation(fuels []models.Fuel) *models.GasStation {
	return &models.GasStation{
		FuelTypes: fuels,
	}

}
func newDriver(d models.Driver) *models.Driver {
	return &d

}

func newCar(c models.Car) *models.Car {
	return &c
}

func fillUpCar(car *models.Car, g *models.GasStation) {
	text := map[string]string{
		"goodBay":        "До свидания",
		"notEnoughMoney": "У вас не хватает денег до полного",
		"notEnoughFuel":  "На заправке не хватает топлива, приезжайте в другой раз.",
		"fillUpCar":      "Машина заправлена",
		"confirmation":   "Хотите заправить машину?\n 1.Да \n 2. Нет\n",
		"confirmation2":  "Хотите заправить машину на все деньги что у вас есть? \n 1.Да \n 2. Нет\n ",
		"dontneedfuel":   "Вам топливо не нужно",
	}
	var status int
	currentFuel := g.FuelTypes[car.FuelType]
	howMuchFuelNeed := car.MaxGasTank - car.CurrentGas
	howMuchMoneyNeed := float64(howMuchFuelNeed) * currentFuel.Price

	balanceWallet := car.Driver.BalanceWallet
	var checkFuelGasStation bool
	var checkEnoughMoney bool

	if howMuchFuelNeed == 0 { // проверка требуется ли топливо машине
		fmt.Println()
		fmt.Println(text["dontneedfuel"])
		os.Exit(0)
	}

	if currentFuel.AmountFuel-float64(howMuchFuelNeed) >= 0 { // проверка наличия топлива на заправке
		checkFuelGasStation = true
		if balanceWallet-howMuchMoneyNeed >= 0 { // проверка хватает ли денег на топливо
			checkEnoughMoney = true
		} else {
			fmt.Println(text["notEnoughMoney"])
			status = 2

		}

	} else {
		fmt.Print(text["notEnoughFuel"])
		fmt.Println(text["goodBay"])
		os.Exit(0)
	}

	if checkFuelGasStation && checkEnoughMoney {
		status = 1
	}

	switch status {
	case 1: // подтверждаем заправку до полного
		answer := offerFillUp(howMuchFuelNeed, howMuchMoneyNeed, balanceWallet)

		if answer == 1 {
			currentFuel.AmountFuel -= float64(howMuchFuelNeed)
			balanceWallet -= howMuchMoneyNeed
			car.CurrentGas = car.MaxGasTank
			fmt.Println(text["fillUpCar"])
		}
	case 2: // предлагаем на оставшиеся деньги
		howMuchCanBye := balanceWallet / currentFuel.Price
		answer := offer(howMuchCanBye)
		if answer == 1 {
			howMuchCanBye := balanceWallet / currentFuel.Price
			car.CurrentGas += int(howMuchCanBye)
			currentFuel.AmountFuel -= howMuchCanBye
			fmt.Println(text["fillUpCar"])
		}

	}
	fmt.Println(text["goodBay"])

}
func offer(howMuchCanBye float64) int {

	text := map[string]string{
		"confirmation":     "Хотите заправить машину на оставщиеся деньги?\n 1.Да \n 2.Нет\n",
		"howMuchCanFillUp": "Кол-во топлива которе можете заправить: %2.f \n",
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(text["confirmation"])
	fmt.Printf(text["howMuchCanFillUp"], howMuchCanBye)
	for {

		scanner.Scan()
		input := scanner.Text()
		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Введите число выбора")
			continue
		}
		switch value {
		case 1:
			return value
		case 2:
			return value
		}

	}

}
func offerFillUp(howMuchFuelNeed int, howMuchMoneyNeed, balanceWallet float64) int {
	text := map[string]string{
		"confirmation":    "Хотите заправить машину?\n 1.Да \n 2.Нет\n",
		"amountMoneyNeed": "Сумму до полного: %2.f \n Кол-во заправляемого топлива: %d \n",
		"willRemain":      "У вас останется денег: %2.f \n",
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(text["confirmation"])
	fmt.Printf(text["amountMoneyNeed"], howMuchMoneyNeed, howMuchFuelNeed)

	fmt.Printf(text["willRemain"], balanceWallet-howMuchMoneyNeed)
	for {

		scanner.Scan()
		input := scanner.Text()
		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Введите число выбора")
			continue
		}
		switch value {
		case 1:
			return value
		case 2:
			return value
		}

	}
}

func main() {

	fuels := []models.Fuel{
		{"diesel", 55.5, 100},
		{"petrol", 60.6, 100},
	}

	gasStation := newGasStation(fuels)
	inputDataCar := getData.GetData(gasStation)
	car := newCar(inputDataCar)

	fillUpCar(car, gasStation)

}
