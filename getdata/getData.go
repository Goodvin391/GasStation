package getData

import (
	"bufio"
	"errors"
	"fmt"
	"gasstation/models"
	"os"
	"strconv"
)

func GetData(gasStation *models.GasStation) models.Car {

	scanner := bufio.NewScanner(os.Stdin)

	text := map[string]string{
		"greeting":      "Добро пожаловать. Вы на заправке. Введите свое имя: ",
		"modelCar":      "Введите модель автомобиля: ",
		"typeFuel":      "Введите тип топлива автомобиля под нужной цифрой: ",
		"walletBalance": "Введите баланс кошелька в рублях",
	}
	fmt.Println(text["greeting"])

	scanner.Scan()
	name := scanner.Text()

	fmt.Println(text["modelCar"])
	scanner.Scan()
	modelCar := scanner.Text()

	fmt.Println(text["typeFuel"])
	typeFuel, err := choiceFuelCar(gasStation)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	maxGasTank, currentGas := fuelIndicators()

	fmt.Println(text["walletBalance"])
	walletBalance := inputAmmountMoney()

	return models.Car{

		Model:      modelCar,
		FuelType:   typeFuel,
		MaxGasTank: maxGasTank,
		CurrentGas: currentGas,
		Driver: models.Driver{
			Name:          name,
			BalanceWallet: walletBalance,
		},
	}

}

func fuelIndicators() (maxGasTank int, currentGas int) {
	text := map[string]string{
		"maxGasTank":       "Введите объем вашего бензобака",
		"currentFuelLevel": "Введите текущий уровень топлива",
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(text["maxGasTank"])
	for {

		scanner.Scan()
		input := scanner.Text()
		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Введите  число")
			continue
		}

		if value < 0 {
			fmt.Println("Введите положительное число")
		} else {
			maxGasTank = value
			break
		}

	}

	fmt.Println(text["currentFuelLevel"])
	for {

		scanner.Scan()
		input := scanner.Text()
		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Введите  число")
			continue
		}

		if value < 0 {
			fmt.Println("Введите положительное число")
		} else if value > maxGasTank {
			fmt.Println("Текущий уровень топлива не может быть выше объема бензобака")
		} else {
			currentGas = value
			break
		}

	}
	return

}

func inputAmmountMoney() float64 {
	var walletBalance float64
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Введите  число")
			continue
		}
		if value < 0 {
			fmt.Println("Введите положительное число")
		} else {
			walletBalance = value
			break
		}
	}
	return walletBalance
}

func choiceFuelCar(gasStation *models.GasStation) (int, error) {
	text := map[string]string{
		"waiting":      "Если нужного топлива нет, нажмите цифру 0",
		"incorrectNum": "Введите только доступные номера топлива",
		"noFuel":       "Нам очень жаль, удачи на дорогах",
	}
	scanner := bufio.NewScanner(os.Stdin)
	var choisenFuel int
	var err error
	fmt.Println(text["waiting"])

	for i := 1; i < len(gasStation.FuelTypes)+1; i++ {
		f := gasStation.FuelTypes[i-1]
		fmt.Printf("%d. %s. Price: %.2f \n", i, f.Name, f.Price)
	}

	for {

		scanner.Scan()
		input := scanner.Text()
		numb, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(text["incorrectNum"])
			continue
		}
		if numb < 0 || numb > len(gasStation.FuelTypes) {
			fmt.Println(text["incorrectNum"])
			continue
		} else if numb == 0 {
			err = errors.New("У клиента нет подходящего топлива")
			fmt.Println(text["noFuel"])
			return 0, err
		} else {
			choisenFuel = numb - 1
			break
		}

	}
	return choisenFuel, err

}
