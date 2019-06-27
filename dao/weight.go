package dao

type Weight struct {
	Id string `json:"id"`
	UserId string `json:"app_user_id"`
	Weight float32 `json:"weight"`
	WeightDate string `json:"weight_date"`
}

