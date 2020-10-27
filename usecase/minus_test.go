package usecase

import "testing"

type MinusTestSuit struct {
	firstNumber int
	secondNumber int
	expectedResult int
}

var minusTestSuit = []MinusTestSuit {
	{1,10,-9},
	{-500,10,-510},
	{14,-10,24},
	{-89,-10,-79},
}

func TestUsecase_Minus(t *testing.T) {
	for _,test := range minusTestSuit {
		uc := CreateUsecase()
		result := uc.Minus(test.firstNumber,test.secondNumber)
		if result != test.expectedResult {
			t.Errorf("Error first number: %d second number: %d, expect: %d actual: %d",test.firstNumber,test.secondNumber,test.expectedResult,result)
		}
	}
}
