package auth

import (
	"fmt"
	"gasstation/actions"
	"gasstation/clearterminal"
	"gasstation/models"
	"strconv"
	"time"
)

func InputData(driver *models.Driver) {


	actions.InputAmmountMoney(driver)
	AddCar(driver)

}

func AddCar(user *models.Driver) {
	var input string
	fmt.Println("Введите модель вашей машины: ")
	fmt.Scan(&input)
	newCar := models.Car{
		Model:      input,
		MaxGasTank: 60,
	}
	user.Car = &newCar

	clearterminal.ClearTerminal()

	for {
		fmt.Println("Введите уровень топлива машины: ")
		fmt.Scan(&input)
		val, err := strconv.Atoi(input)
		if err != nil || val < 0 {
			fmt.Println("Не корректное число")
			time.Sleep(1 * time.Second)
			clearterminal.ClearTerminal()
			continue

		}
		if uint(val) > user.Car.MaxGasTank {
			fmt.Println("Уровень топлива не может быть выше максимального объема бака. (Максимальный объем бака 60л)")
			time.Sleep(1 * time.Second)
			clearterminal.ClearTerminal()
			continue
		}

		user.Car.CurrentGas = uint(val)

		return

	}

}
