package dao

type Activity struct {
	Name string `json:"name"`
	Id string `json:"id"`
	UserId string `json:"app_user_id"`
	DateCompleted string `json:"datetime_completed"`
}
