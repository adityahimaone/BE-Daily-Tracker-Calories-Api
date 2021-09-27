package calories

type Domain struct {
	ID int
	Weight int
	Calorie int
}

type Service interface{
	Create ()
}

type Repository interface{

}
