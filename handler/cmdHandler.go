package handler

import (
	"errors"
	"fmt"
	"goCmd/constant"
	"goCmd/usecase"
	"goCmd/util"
	"os"
	"strconv"
)


func inputHandler(firstNumber *int, secondNumber *int) error {
	fmt.Print("Please enter 1st number: ")
	_, err := fmt.Scanf("%d\n", firstNumber)
	if err != nil {
		return errors.New("invalid 1st number")
	}
	fmt.Print("Please enter 2nd number: ")
	_, err = fmt.Scanf("%d\n", secondNumber)
	if err != nil {
		return errors.New("invalid 2nd number")
	}
	return nil
}

func parseArgumentToMenu(argumentPlace int, utility util.Utility) (int, error) {
	var argument string
	var selectedMenu int
	if utility.CheckArgument(argumentPlace) {
		argument = os.Args[argumentPlace]
		if argument != "" {
			var err error
			selectedMenu, err = strconv.Atoi(argument)
			if err != nil {
				return 0, errors.New("invalid input argument")
			}
		}
		return selectedMenu, nil
	}
	return 0, nil
}

type CommandLineHandler struct {
	ut util.Utility
}

func CreateCommandLineHandler(utility util.Utility) HandlerInterface  {
	return CommandLineHandler{
		ut: utility,
	}
}


func (h CommandLineHandler) Start()  {
	var firstNumber, secondNumber, selectedMenu int
	selectedMenu, err := parseArgumentToMenu(1, h.ut)
	if err != nil {
		fmt.Println(err)
	}

	for {
		if selectedMenu == 0 {
			fmt.Println(h.ut.DisplayMenu())
			fmt.Print("Please enter command: ")
			fmt.Scanf("%d\n", &selectedMenu)
		}

		switch constant.MenuType(selectedMenu) {
		case constant.PLUS:
			if err := inputHandler(&firstNumber, &secondNumber);
				err != nil {
				fmt.Println(err.Error())
				selectedMenu = h.ut.Reset()
				continue
			}
			usecase.Plus(firstNumber, secondNumber)
		case constant.MINUS:
			if err := inputHandler(&firstNumber, &secondNumber); err != nil {
				fmt.Println(err.Error())
				selectedMenu = h.ut.Reset()
				continue
			}
			usecase.Minus(firstNumber, secondNumber)
		}
		selectedMenu = 0
	}
}