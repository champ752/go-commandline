package usecase

type Usecaser interface {
	Plus(int,int) int
	Minus(int,int) int
}

type Usecase struct {

}

func CreateUsecase() Usecaser  {
	return Usecase{}
}
