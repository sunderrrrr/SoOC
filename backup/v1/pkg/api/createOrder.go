package api

import (
	"encoding/json"
	"net/http"
)

// Создание нового заказа
// Пример запроса curl -X POST -H "Content-Type: application/json" -d "{\"Dish\": \"Burger\", \"Quantity\": 1, \"IsReady\": false, \"IsServed\": false}" http://localhost:8080/api/order/create

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		var newOrder Order
		_ = json.NewDecoder(r.Body).Decode(&newOrder) // Распаковываем массив и добавляем в newOrder
		newOrder.ID = len(orders) + 1
		orders = append(orders, newOrder) //помещаем новый элемент в конец массива orders
		json.NewEncoder(w).Encode(newOrder)

	default:
		w.Write([]byte("Use POST protocol for this endpoint"))
	}

}
