package actions

import (
	"fmt"
	"gasstation/clearterminal"
	"gasstation/models"
	"strconv"
	"time"
)

func InputAmmountMoney(currentDriver *models.Driver) {
	for {
		var input string
		fmt.Println("Введите сумму пополнения:")
		fmt.Scanln(&input)
		val, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			fmt.Println("Не корректное число")
			clearterminal.ClearTerminal()
			continue

		}
		uintInput := uint(val)
		currentDriver.Wallet.BalanceWallet += uintInput
		fmt.Println("Ваш кошелек пополнен")
		time.Sleep(1 * time.Second)
		clearterminal.ClearTerminal()
		return
	}

}

func ShowBalance(currentDriver *models.Driver) {
	res := currentDriver.Wallet.BalanceWallet
	fmt.Printf("Ваш баланс: %v\n", res)

}
