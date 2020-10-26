package util

import (
	"fmt"
	"os"
)

type Utilize interface {
	DisplayMenu() string
	CheckArgument(argumentPlace int) bool
	Reset() int
}
type Utility struct {

}

func CreateUtility() Utilize {
	return Utility{}
}
func  (ut Utility)  DisplayMenu() string {
	outputString := fmt.Sprintf("%s\n%s\n%s","Please select menu" ,"1 plus" ,"2 minus" )
	return outputString
}
func  (ut Utility) CheckArgument(argumentPlace int) bool {
	if len(os.Args) > argumentPlace {
		return true
	}
	return false
}
func (ut Utility) Reset() int {
	var discard string
	_, _ = fmt.Scanln(&discard)
	return 0
}