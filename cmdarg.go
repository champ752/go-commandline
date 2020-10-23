package main

import (
	"fmt"
	"goCmd/actions"
	"goCmd/constant"
	"os"
	"strconv"
)

func InputHandler(firstNumber *int,secondNumber *int) error {
	fmt.Print("Please enter 1st number: ");
	_, err := fmt.Scanf("%d\n",firstNumber)
	if err != nil {
		fmt.Println("Invalid 1st number")
		return err
	}

	fmt.Print("Please enter 2nd number: ");
	_, err = fmt.Scanf("%d\n",secondNumber)
	if err != nil {
		fmt.Println("Invalid 2st number")
		return err
	}
	return nil
}
func DisplayMenu()  {
	fmt.Println("Please select selectedMenu")
	fmt.Println("1 plus")
	fmt.Println("2 minus")
}
func CheckArg(arg *string, selectedMenu *int) {
	if len(os.Args) > 1 {
		*arg = os.Args[1]
	}
	if *arg != "" {
		var err error
		*selectedMenu, err =  strconv.Atoi(*arg)
		if err != nil {
			fmt.Println("Invalid input arg")
			*selectedMenu = 0
		}
	}
}
func Reset(selectedMenu *int) {
	*selectedMenu = 0
	var discard string
	fmt.Scanln(&discard)
}
func main()  {
	var firstNumber,secondNumber,selectedMenu int
	var arg string
	CheckArg(&arg, &selectedMenu)
	for {
		if selectedMenu == 0 {
			DisplayMenu()
			fmt.Print("Please enter command: ")
			fmt.Scanf("%d\n",&selectedMenu)
		}

		switch constant.MenuType(selectedMenu) {
		case constant.PLUS: if err :=InputHandler(&firstNumber,&secondNumber); err!=nil{Reset(&selectedMenu);continue}; actions.Plus(firstNumber,secondNumber)
		case constant.MINUS:if err :=InputHandler(&firstNumber,&secondNumber); err!=nil{Reset(&selectedMenu);continue};  actions.Minus(firstNumber,secondNumber)
		}
		selectedMenu = 0
	}
}
