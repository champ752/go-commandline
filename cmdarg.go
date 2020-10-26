package main

import (
	"errors"
	"fmt"
	"goCmd/actions"
	"goCmd/constant"
	"os"
	"strconv"
)


func InputHandler(firstNumber *int,secondNumber *int) error {
	fmt.Print("Please enter 1st number: ")
	_, err := fmt.Scanf("%d\n",firstNumber)
	if err != nil {
		return errors.New("invalid 1st number")
	}
	fmt.Print("Please enter 2nd number: ")
	_, err = fmt.Scanf("%d\n",secondNumber)
	if err != nil {
		return errors.New("invalid 2nd number")
	}
	return nil
}
func DisplayMenu() string {
	outputString := fmt.Sprintf("%s\n%s\n%s","Please select menu" ,"1 plus" ,"2 minus" )
	return outputString
}
func CheckArgument(argumentPlace int) bool {
	if len(os.Args) > argumentPlace {
		return true
	}
	return false
}
func ParseArgumentToMenu(argumentPlace int) (int,error) {
	var argument string
	var selectedMenu int
	if CheckArgument(argumentPlace) {
		argument = os.Args[argumentPlace]
		if argument != "" {
			var err error
			selectedMenu, err = strconv.Atoi(argument)
			if err != nil {
				return 0, errors.New("invalid input argument")
			}
		}
		return selectedMenu,nil
	}
	return 0,nil
}
func Reset() int {
	var discard string
	_, _ = fmt.Scanln(&discard)
	return 0
}
func main()  {
	var firstNumber,secondNumber,selectedMenu int
	
	selectedMenu,err := ParseArgumentToMenu(1)
	if err != nil {
		fmt.Println(err)
	}

	for {
		if selectedMenu == 0 {
			fmt.Println(DisplayMenu())
			fmt.Print("Please enter command: ")
			fmt.Scanf("%d\n",&selectedMenu)
		}


		switch constant.MenuType(selectedMenu) {
		case constant.PLUS: if err :=InputHandler(&firstNumber,&secondNumber);
		err!=nil{fmt.Println(err.Error());selectedMenu=Reset();continue};
		actions.Plus(firstNumber,secondNumber)
		case constant.MINUS:if err :=InputHandler(&firstNumber,&secondNumber);
		err!=nil{fmt.Println(err.Error());selectedMenu=Reset();continue};
		actions.Minus(firstNumber,secondNumber)
		}
		selectedMenu = 0
	}
}
