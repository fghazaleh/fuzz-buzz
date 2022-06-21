package service

type fuzz struct{}

func NewFuzz() FuzzBuzzer {
	return &fuzz{}
}

func (f fuzz) IsValid(n int) bool {
	return n%3 == 0
}

func (f fuzz) GetOutput() string {
	return "Fuzz"
}
