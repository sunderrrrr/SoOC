package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Обновление состояния заказа
// Пример запроса (для терминала) curl -X PUT -H "Content-Type: application/json" -d "{\"ID\": 1, \"Dish\": \"Pizza\", \"Quantity\": 2, \"IsReady\": true, \"IsServed\": true}" http://localhost:8080/api/order/update

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		//fmt.Println("Пришел запрос на update")
		w.Header().Set("Content-Type", "application/json")
		var updatedOrder Order
		_ = json.NewDecoder(r.Body).Decode(&updatedOrder) //Декодирование тела запроса

		for i, order := range orders {
			if updatedOrder.ID == order.ID {
				orders[i].IsReady = updatedOrder.IsReady
				json.NewEncoder(w).Encode(orders[i])
				return
			}
		}
		json.NewEncoder(w).Encode(&Order{}) //Если цикл завершается и не находится нужный заказ, то возвращается пустой заказ

	default:
		fmt.Println("Use PUT protocol for this endpoint")
		w.Write([]byte("Use PUT protocol for this endpoint"))
	}

}
