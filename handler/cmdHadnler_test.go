package handler

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// pull error to Err1stNumber (constant instead)
type InputHandlerResult struct {
	input string
	firstNumber int
	secondNumber int
	err error
}

var inputHandlerResult = []InputHandlerResult{
	{"a\n1234\n",0,0,errors.New("invalid 1st number")},
	{"700\n1200\n",700,1200,nil},
	{"700\n1200a\n",0,0,errors.New("invalid 2nd number")},
	{"700ab\n1200a\n",0,0,errors.New("invalid 1st number")},
}

func TestCommandLineHandler_InputHandler(t *testing.T) {


	for _, test := range inputHandlerResult {

		in, err := ioutil.TempFile("", "")
		if err != nil {
			t.Fatal(err)
		}
		defer in.Close()

		_, err = io.WriteString(in, test.input)
		if err != nil {
			t.Fatal(err)
		}

		_, err = in.Seek(0, os.SEEK_SET)
		if err != nil {
			t.Fatal(err)
		}

		num1, num2, err := InputHandler(in)
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Error("error fail")
			}
		}
		if num1 != test.firstNumber || num2 != test.secondNumber {

			t.Errorf("received input error expect num1,num2 %d,%d actual %d,%d",num1,num2,test.firstNumber,test.secondNumber)
		}
	}


}
