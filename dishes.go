package sooc

type Dishes struct {
	DishyId   int    `json:"dish_id"`
	DishName  string `json:"dish_name"`
	CostPrice string `json:"cost_price"`
	ImgUrl    string `json:"url_photo"`
}
