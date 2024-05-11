package api

import (
	"encoding/json"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
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
