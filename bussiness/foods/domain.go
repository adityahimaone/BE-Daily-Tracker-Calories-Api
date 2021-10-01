package foods

import "time"

type Domain struct {
	ID        int
	Name      string
	Calorie   float64
	Photo     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	GetFoodByName(name string) (*Domain, error)
	GetFoodAPI(food *Domain) (*Domain, error)
	SaveFood(food *Domain) (*Domain, error)
}


type Repository interface {
	GetFoodByName(name string) (*Domain, error)
	Insert(food *Domain) (*Domain, error)
}
