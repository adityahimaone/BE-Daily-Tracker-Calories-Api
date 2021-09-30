package calories

type Domain struct {
	ID           int
	Weight       int
	Height       int
	Gender       string
	Age          int
	ActivityType int
	Calorie      float64
}

type Service interface {
	CountCalorie(calorie *Domain) (*Domain, error)
	CreateCalorie(calorie *Domain) (*Domain, error)
	UpdateCalorie(calorie *Domain) (*Domain, error)
}

type Repository interface {
	Insert(calorie *Domain) (*Domain, error)
	Update(calorie *Domain) (*Domain, error)
}
