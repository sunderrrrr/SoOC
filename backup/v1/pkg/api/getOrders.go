package api

import (
	"encoding/json"
	"net/http"
)

// Получение списка всех заказов ввиде массива JSON
// Пример запроса curl -X GET http://localhost:8080/api/order/list
func GetOrders(w http.ResponseWriter, r *http.Request) {
	//login, passw, ok := r.BasicAuth()
	//if ok && LoginU(login, passw) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)

	default:
		w.Write([]byte("Use GET protocol for this endpoint"))
	}
	//}

}
