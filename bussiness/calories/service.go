package calories

import (
	_users "daily-tracker-calories/bussiness/users"
)

type serviceCalorie struct {
	repository  Repository
	userService _users.Service
}

func NewService(repositoryCalorie Repository, userServ _users.Service) Service {
	return &serviceCalorie{
		repository:  repositoryCalorie,
		userService: userServ,
	}
}

func (s *serviceCalorie) CountCalorie(calorie *Domain) (*Domain, error) {
	activityValue := calorie.ActivityType
	weight := calorie.Weight
	weightFloat := float64(weight)
	height := calorie.Height
	heightFloat := float64(height)
	age := calorie.Age
	ageFloat := float64(age)
	gender := calorie.Gender
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
		calories = (10 * weightFloat) + (6.25 * heightFloat) - (5*ageFloat)*valueActivity
	} else {
		calories = ((10 * weightFloat) + (6.25 * heightFloat) - (5 * ageFloat) - 161) * valueActivity
	}
	calorie.Calorie = calories
	return &Domain{}, nil
}

func (s serviceCalorie) CreateCalorie(calorie *Domain, idUser int) (*Domain, error) {
	calorie.UserID = idUser
	_, err := s.CountCalorie(calorie)
	if err != nil {
		return &Domain{}, err
	}
	validId, err := s.repository.GetCalorieByUserID(idUser)
	if validId.ID == 0 {
		result, err := s.repository.Insert(calorie, idUser)
		if err != nil {
			return &Domain{}, err
		}
		return result, nil
	}
	result, err := s.repository.Update(calorie, idUser)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (s *serviceCalorie) GetCalorieByUserID(id int) (*Domain, error) {
	result, err := s.repository.GetCalorieByUserID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
