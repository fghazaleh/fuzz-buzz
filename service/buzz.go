package service

type buzz struct{}

func NewBuzz() FuzzBuzzer {
	return &buzz{}
}

func (f buzz) IsValid(n int) bool {
	return n%5 == 0
}

func (f buzz) GetOutput() string {
	return "Buzz"
}
