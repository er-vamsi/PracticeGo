package napodate

import ("context"
		"time")

type Service interface{
	Status(ctx context.Context) (string, error)  // Returns status of MicroService whether up and running or not.
	Get(ctx context.Context) (string, error)     // Get the current date
	Validate(ctx context.Context, date string) (bool, error) //Validates the format of date dd/mm/yyyy and return true/false
}

type dateService struct{}

//NewService makes a new Service
func NewService() Service{
	return dateService{}
}

//Status only tell us that our service is ok!
func (dateService) Status(ctx context.Context) (string,error){
	return "ok",nil
}

//Get will return only today's date

func (dateService) Get(ctx context.Context) (string,error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}

//Validate will check if the date is today's date

func (dateService) Validate(ctx context.Context, date string) (bool,error){
	_, err := time.Parse("02/01/2006",date)
	if err != nil {
		return false, err
	}
	return true,nil
}

