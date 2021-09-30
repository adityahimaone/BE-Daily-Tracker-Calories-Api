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
	CountCalorie(user *Domain) (*Domain, error)
	Create(user *Domain) (*Domain, error)
}

type Repository interface {
	Insert(user *Domain) (*Domain, error)
}
