package actions

import (
	"fmt"
	"gasstation/clearterminal"
	"gasstation/models"
	"strconv"
	"time"
)

func FillUpCar(gasStation *models.GasStation, user *models.Driver) {
	choiseFuelForCar(*gasStation, user)

	chosenFuel := gasStation.FuelTypes[user.Car.FuelType]
	howMuchFuelNeed := user.Car.MaxGasTank - user.Car.CurrentGas
	howMuchMoneyNeed := howMuchFuelNeed * chosenFuel.Price

	if user.Car.CurrentGas == user.Car.MaxGasTank {
		fmt.Println("У вас полный бак")
		time.Sleep(1 * time.Second)
		clearterminal.ClearTerminal()
		return
	}

	if int(chosenFuel.AmountFuel-howMuchFuelNeed) < 0 {
		fmt.Println("На автозапровочной станции не хватит топлива, чтобы заполнить до полного")
		time.Sleep(1 * time.Second)
		clearterminal.ClearTerminal()
		return
	}

	if int(user.Wallet.BalanceWallet-howMuchMoneyNeed) < 0 {
		fmt.Println("У вас недостаточно денег")
		time.Sleep(1 * time.Second)
		clearterminal.ClearTerminal()

		fillUpToFullAllMoney(gasStation, user)

		return
	}

	fillUpToFull(gasStation, user)

}

func choiseFuelForCar(gasStation models.GasStation, user *models.Driver) {
	var input string

	for {
		fmt.Println("Выберете тип топлива которым хотите заправить автомобиль")
		for index, value := range gasStation.FuelTypes {
			fmt.Printf("%d. - %s. Цена - %d\n", index, value.Name, value.Price)

		}
		fmt.Scanln(&input)
		val, err := strconv.Atoi(input)

		numFuel := len(gasStation.FuelTypes) - 1

		if err != nil || val < 0 || val > numFuel {
			clearterminal.ClearTerminal()
			fmt.Println("Введите цифру топлива")
			time.Sleep(1 * time.Second)
			continue
		}
		user.Car.FuelType = val
		fmt.Printf("Выбрано топливо: %s (индекс: %d)\n", gasStation.FuelTypes[val].Name, val)

		clearterminal.ClearTerminal()
		return
	}

}

func fillUpToFull(gasStation *models.GasStation, user *models.Driver) {
	var input string
	chosenFuel := gasStation.FuelTypes[user.Car.FuelType]
	howMuchFuelNeed := user.Car.MaxGasTank - user.Car.CurrentGas
	howMuchMoneyNeed := howMuchFuelNeed * chosenFuel.Price

	for {
		fmt.Print("Хотите заправиться до полного? \n1.Да\n2.Нет\n ")

		fmt.Scanln(&input)

		switch input {
		case "1":
			gasStation.FuelTypes[user.Car.FuelType].AmountFuel -= howMuchFuelNeed
			user.Wallet.BalanceWallet -= howMuchMoneyNeed

			user.Car.CurrentGas = user.Car.MaxGasTank
			fmt.Println("Ваш бак заполнен до полного")
			time.Sleep(1 * time.Second)
			clearterminal.ClearTerminal()
			return

		case "2":
			clearterminal.ClearTerminal()
			return
		default:
			clearterminal.ClearTerminal()
			continue
		}

	}
}

func fillUpToFullAllMoney(gasStation *models.GasStation, user *models.Driver) {
	var input string
	chosenFuel := gasStation.FuelTypes[user.Car.FuelType]
	howMuchFuelCanBuy := user.Wallet.BalanceWallet / chosenFuel.Price

	for {
		fmt.Printf("Хотите заправиться на оставшиеся деньги - %d?\n1.Да\n2.Нет\n", user.Wallet.BalanceWallet)
		fmt.Printf("Кол-во топлива на которе хватит денег - %d.\n", howMuchFuelCanBuy)

		fmt.Scanln(&input)

		switch input {
		case "1":
			gasStation.FuelTypes[user.Car.FuelType].AmountFuel -= howMuchFuelCanBuy
			user.Wallet.BalanceWallet = 0

			user.Car.CurrentGas += howMuchFuelCanBuy
			fmt.Printf("Ваш бак заполнен до %d\n", user.Car.CurrentGas)
			time.Sleep(1 * time.Second)
			clearterminal.ClearTerminal()
			return

		case "2":
			clearterminal.ClearTerminal()
			return
		default:
			clearterminal.ClearTerminal()
			continue
		}

	}
}
