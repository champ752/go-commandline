package usecase

import (
	"io/ioutil"
	"testing"
)

type PlusResult struct {
	firstNumber int
	secondNumber int
	expectedResult int
}

var plusResults = []PlusResult{
	{1,2,3},
	{500,700,1200},
	{500,-700,-200},
	{17,-12,5},
	{-50,36,-14},
	{-50,55,5},
	{-50,-36,-86},
}

func TestUsecase_Plus(t *testing.T) {
	for _, test := range plusResults {
		uc := CreateUsecase()
		result := uc.Plus(test.firstNumber,test.secondNumber)
		if result != test.expectedResult {
			t.Errorf("Error first number: %d second number: %d, expect: %d actual: %d",test.firstNumber,test.secondNumber,test.expectedResult,result)
		}
	}
}
func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "Hello world" {
		t.Fatal("String contents do not match expected")
	}
}
