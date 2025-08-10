package auth

import (
	"fmt"

	"gasstation/clearterminal"
	"gasstation/models"
	"time"
)

func Registarion(users *map[string]models.Driver, currentDriver *models.Driver) {

	var input string

	name, ok := ChekReg(*users)

	if ok {

		fmt.Println("Вы уже зарегистрированы!")

		time.Sleep(1 * time.Second)
		clearterminal.ClearTerminal()
		return
	}

	if !ok {
		fmt.Printf("Вы не зарегистрированны - %s. Хотите зарегистрироваться?\n 1.Да\n 2.Нет\n ", name)
		for {
			fmt.Scanln(&input)
			clearterminal.ClearTerminal()
			switch input {
			case "1":

				currentDriver.Name = name
				InputData(currentDriver)
				fmt.Println("Вы зарегистрированы!")

				time.Sleep(1 * time.Second)
				clearterminal.ClearTerminal()
				return

			case "2":
				fmt.Println("Вы не зарегистрировались!")
				time.Sleep(1 * time.Second)
				clearterminal.ClearTerminal()
				return
			default:
				clearterminal.ClearTerminal()
				fmt.Printf("Хотите зарегистрироваться?\n 1.Да\n 2.Нет\n ")
				fmt.Println("Введите число.")
			}

		}

	}

}

func ChekReg(users map[string]models.Driver) (name string, res bool) {
	for {

		fmt.Println("Введите ваше имя: ")
		fmt.Scanln(&name)
		clearterminal.ClearTerminal()
		if name == "" {
			continue
		}
		break
	}
	_, ok := users[name]
	if !ok {
		res = false
		return
	}
	res = true
	return
}
