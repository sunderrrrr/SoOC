package api

import (
	"SoCC/frontend"
	"fmt"
	"net/http"
)

type Order struct { //Структура заказа
	ID       int    `json:"id"`       //Id
	Dish     string `json:"dish"`     //Название
	Quantity int    `json:"quantity"` //Количество
	IsReady  bool   `json:"is_ready"`
	IsServed bool   `json:"is_served"`
}

var orders []Order

type api struct { //Структура Api
	address string
	router  *http.ServeMux
}

func New(address string, router *http.ServeMux) *api { //Инициализаця Api
	return &api{address: address, router: router}
}

func (api *api) FillEndpoints() {
	//FrontEnd endpoints
	api.router.HandleFunc("/", frontend.Index)
	api.router.HandleFunc("/create", frontend.Create)
	api.router.HandleFunc("/guide", frontend.Guide)
	api.router.Handle("../html/", http.StripPrefix("../html/", http.FileServer(http.Dir("../html/"))))
	api.router.Handle("../html/assets/", http.StripPrefix("../html/assets", http.FileServer(http.Dir("../html/assets"))))

	//API endpoints
	api.router.HandleFunc("/api/order/list", GetOrders)
	api.router.HandleFunc("/api/order/create", CreateOrder)
	api.router.HandleFunc("/api/order/update", UpdateOrder)
	api.router.HandleFunc("/api/order/delete", DeleteOrder)
	fmt.Println("API запущено")

	//Запролнения массивов с заказами хоть чем-то
	orders = append(orders, Order{ID: 1, Dish: "Пицца Пеперони", Quantity: 1, IsReady: true, IsServed: false}, Order{2, "Пельмени жаренные", 2, false, false})

}

// Метод запуска сервера для Api
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.address, api.router)
}
