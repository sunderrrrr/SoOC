package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Order struct {
	ID       int    `json:"id"`       //Id
	Dish     string `json:"dish"`     //Название
	Quantity int    `json:"quantity"` //Количество
	IsReady  bool   `json:"is_ready"`
	IsServed bool   `json:"is_served"`
}

var orders []Order

func getOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)

	default:
		w.Write([]byte("Use GET protocol for this endpoint"))
	}

}

func createOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		var newOrder Order
		_ = json.NewDecoder(r.Body).Decode(&newOrder)
		newOrder.ID = len(orders) + 1
		orders = append(orders, newOrder)
		json.NewEncoder(w).Encode(newOrder)

	default:
		w.Write([]byte("Use POST protocol for this endpoint"))
	}

}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		fmt.Println("Пришел запрос на update")
		w.Header().Set("Content-Type", "application/json")
		var updatedOrder Order
		_ = json.NewDecoder(r.Body).Decode(&updatedOrder)

		for i, order := range orders {
			if order.ID == updatedOrder.ID {
				orders[i].IsReady = updatedOrder.IsReady
				json.NewEncoder(w).Encode(orders[i])
				return
			}
		}
		json.NewEncoder(w).Encode(&Order{})

	default:
		w.Write([]byte("Use PUT protocol for this endpoint"))
	}

}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		w.Header().Set("Content-Type", "application/json")
		var deletedOrder Order
		_ = json.NewDecoder(r.Body).Decode(&deletedOrder)

		for i, order := range orders {
			if order.ID == deletedOrder.ID {
				orders = append(orders[:i], orders[i+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(orders)

	default:
		w.Write([]byte("Use DELETE protocol for this endpoint"))
	}

}
func index_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "../html/index.html")
		return
	}

	if r.URL.Path == "/style.css" {
		http.ServeFile(w, r, "../html/style.css")
		return
	}
}

// Чтобы раздавать файл нужно использовать путь /эндпоин/нужный файл
func create(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/create" {
		http.ServeFile(w, r, "../html/create.html")
		return
	}

	if r.URL.Path == "/create/style.css" {
		http.ServeFile(w, r, "../html/style.css")
		return
	}
}

func runServer() {
	//Api
	http.HandleFunc("/api/order/list", getOrders)
	http.HandleFunc("/api/order/create", createOrder)
	http.HandleFunc("/api/order/update", updateOrder)
	http.HandleFunc("/api/order/delete", deleteOrder)
	//FrontEnd
	http.HandleFunc("/", index_page)
	http.HandleFunc("/create", create)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("../html/"))))
	http.Handle("/html/assets/", http.StripPrefix("/html/assets", http.FileServer(http.Dir("../html/assets"))))

	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	orders = append(orders, Order{ID: 1, Dish: "Пицца Пеперони", Quantity: 1, IsReady: true, IsServed: false}, Order{2, "Пельмени жаренные", 2, false, false})
	runServer()
}
