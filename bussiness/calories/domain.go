package calories

type Domain struct {
	ID           int
	Weight       int
	Height       int
	Gender       string
	Age          int
	ActivityType int
	Calorie      float64
	UserID       int
}

type Service interface {
	CountCalorie(calorie *Domain) (*Domain, error)
	CreateCalorie(calorie *Domain, id int) (*Domain, error)
	GetCalorieByUserID(id int) (*Domain, error)
}

type Repository interface {
	Insert(calorie *Domain, idUser int) (*Domain, error)
	Update(calorie *Domain, id int) (*Domain, error)
	GetCalorieByUserID(id int) (*Domain, error)
}
