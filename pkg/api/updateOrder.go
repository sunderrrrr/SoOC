package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("Use PUT protocol for this endpoint")
		w.Write([]byte("Use PUT protocol for this endpoint"))
	}

}
