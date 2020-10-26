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


type CommandLineHandler struct {
	ut util.Utilize
	uc usecase.Usecaser
}

func CreateCommandLineHandler(utility util.Utilize,uc usecase.Usecaser) Handler  {
	return CommandLineHandler{
		ut: utility,
		uc: uc,
	}
}


func (h CommandLineHandler) inputHandler(firstNumber *int, secondNumber *int) error {

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

func (h CommandLineHandler) parseArgumentToMenu(argumentPlace int) (int, error) {
	if h.ut.CheckArgument(argumentPlace) {
		argument := os.Args[argumentPlace]
		if argument != "" {
			selectedMenu, err := strconv.Atoi(argument)
			if err != nil {
				return 0, errors.New("invalid input argument")
			}
			return selectedMenu, nil
		}
	}

	return 0, nil
}

func (h CommandLineHandler) Minus() (int ,error)  {
	var firstNumber, secondNumber int
	if err := h.inputHandler(&firstNumber, &secondNumber); err != nil {
		return 0 ,err
	}
	return h.uc.Minus(firstNumber, secondNumber),nil
}

func (h CommandLineHandler) Plus() (int ,error)  {
	var firstNumber, secondNumber int
	if err := h.
		inputHandler(&firstNumber, &secondNumber); err != nil {
		return 0 ,err
	}
	return h.uc.Plus(firstNumber, secondNumber),nil
}


func (h CommandLineHandler) Serve()  {
	selectedMenu, err := h.parseArgumentToMenu(1)
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
			result,err :=h.Plus()
			if err != nil {
				fmt.Println(err.Error())
				selectedMenu = h.ut.Reset()
				continue
			}
			fmt.Println("Result is:", result)

		case constant.MINUS:
			result,err :=h.Minus()
			if err != nil {
				fmt.Println(err.Error())
				selectedMenu = h.ut.Reset()
				continue
			}
			fmt.Println("Result is:", result)
		}
		selectedMenu = 0
	}
}

