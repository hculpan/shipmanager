package data

type Ship struct {
	Id            string
	Name          string
	Tons          int
	CargoCapacity int
	HighPassage   int
	MiddlePassage int
	BasicPassage  int
	LowPassage    int
	StewardRating int

	TotalCost      float32
	RemainingCost  float32
	MonthlyPayment int
}
