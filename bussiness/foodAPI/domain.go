package foodAPI

type Domain struct {
	Name    string
	Photo   string
	Calorie float64
}

type Repository interface {
	GetFoodByName(name string) (*Domain, error)
}
