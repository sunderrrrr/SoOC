package api

import (
	"encoding/json"
	"net/http"
)

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
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
