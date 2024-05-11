package api

import (
	"encoding/json"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)

	default:
		w.Write([]byte("Use GET protocol for this endpoint"))
	}

}
