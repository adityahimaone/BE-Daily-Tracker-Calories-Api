package calories

type serviceCalorie struct {
	repository Repository
}

func NewService(repositoryCalorie Repository) Service {
	return &serviceCalorie{
		repository: repositoryCalorie,
	}
}

func (s serviceCalorie) CountCalorie(user *Domain) (*Domain, error) {
	activityValue := user.ActivityType
	weight := user.Weight
	weightFloat := float64(weight)
	height := user.Height
	heightFloat := float64(height)
	age := user.Age
	ageFloat := float64(age)
	gender := user.Gender
	valueActivity := 0.0
	calories := 0.0
	switch activityValue {
	case 1:
		valueActivity = 1.2
	case 2:
		valueActivity = 1.375
	case 3:
		valueActivity = 1.55
	case 4:
		valueActivity = 1.725
	case 5:
		valueActivity = 1.9
	default:
		valueActivity = 1.0
	}
	if gender == "male" {
		calories = (10 * weightFloat) + (6.25 * heightFloat) - (5 * ageFloat) * valueActivity
	}else {
		calories = ((10 * weightFloat) + (6.25 * heightFloat) - (5 * ageFloat) - 161) * valueActivity
	}
	user.Calorie = calories
	return nil, nil
}

func (s serviceCalorie) Create(user *Domain) (*Domain, error) {
	panic("implement me")
}
