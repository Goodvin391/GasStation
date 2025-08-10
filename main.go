package main

import (
	"fmt"
	"gasstation/actions"
	"gasstation/auth"
	"gasstation/clearterminal"
	"gasstation/models"
	"os"
	"time"
)

var Fuels []models.Fuel = []models.Fuel{
	{"diesel", 55, 100},
	{"petrol", 60, 100},
}

func gasStation(drivers map[string]models.Driver, currentDriver *models.Driver, gasStation *models.GasStation) {
	clearterminal.ClearTerminal()
	text := map[string]string{
		"greeting":        "Добро пожаловать. Вы на заправке. ",
		"actionselection": "Выберете дальнейшие действия %s:\n 1. Пополнить баланс \n 2. Показать баланс\n 3. Заправить автомобиль\n 4. Выйти из игры\n 5. Остановить приложение\n",
		"modelCar":        "Введите модель автомобиля: ",
		"typeFuel":        "Введите тип топлива автомобиля под нужной цифрой: ",
		"walletBalance":   "Введите баланс кошелька в рублях",
	}
	fmt.Println(text["greeting"])

	for {
		var input string
		fmt.Printf("Ваш автомобиль: %s. Текущий уровень топлива: %d\n", currentDriver.Car.Model, currentDriver.Car.CurrentGas)
		fmt.Printf(text["actionselection"], currentDriver.Name)
		fmt.Scanln(&input)

		clearterminal.ClearTerminal()

		switch input {
		case "1":
			input = ""
			actions.InputAmmountMoney(currentDriver)
		case "2":
			input = ""
			actions.ShowBalance(currentDriver)
			time.Sleep(2 * time.Second)
			clearterminal.ClearTerminal()

		case "3":
			input = ""
			actions.FillUpCar(gasStation, currentDriver)
		case "4":
			return
		case "5":
			os.Exit(0)

		default:
			input = ""
			clearterminal.ClearTerminal()
			fmt.Println("Введите число.")
		}

	}

}

func app() {
	drivers := make(map[string]models.Driver)
	newGasStation := models.GasStation{Fuels}
	var currentDriver models.Driver

	for {
		drivers[currentDriver.Name] = currentDriver
		var input string
		fmt.Printf("Выбирете действие: \n 1. Начать игру \n 2. Завершить приложение\n")
		fmt.Scanln(&input)

		switch input {
		case "1":
			clearterminal.ClearTerminal()
			auth.Registarion(&drivers, &currentDriver)

			if currentDriver.Name != "" {
				gasStation(drivers, &currentDriver, &newGasStation)
			}

		case "2":
			os.Exit(0)
		default:
			clearterminal.ClearTerminal()
			continue
		}

	}

}

func main() {
	app()
}
