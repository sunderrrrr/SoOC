package api

import (
	"SoCC/frontend"
	"SoCC/pkg/auth"
	"SoCC/pkg/postgres"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	//DB Connect
	fmt.Println("Подгрузка конфигов")
	if err := initConfig(); err != nil {
		logrus.Fatalf("error: failed config read attempt: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to load env config: %s", err.Error())
	}
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error: Failed db init: %s", err.Error())
	}
	postgres.Initial(db)

	//FrontEnd endpoints
	go api.router.HandleFunc("/", frontend.Index)
	go api.router.HandleFunc("/create", frontend.Create)
	go api.router.HandleFunc("/guide", frontend.Guide)
	go api.router.Handle("../html/", http.StripPrefix("../html/", http.FileServer(http.Dir("../html/"))))
	go api.router.Handle("../html/assets/", http.StripPrefix("../html/assets", http.FileServer(http.Dir("../html/assets"))))

	//API endpoints
	go api.router.HandleFunc("/api/order/list", GetOrders)
	go api.router.HandleFunc("/api/order/create", CreateOrder)
	go api.router.HandleFunc("/api/order/update", UpdateOrder)
	go api.router.HandleFunc("/api/order/delete", DeleteOrder)

	//Auth Endpoints
	go api.router.HandleFunc("/api/users/login", auth.Login)
	go api.router.HandleFunc("/api/users/register", auth.Register)
	fmt.Println("API запущено")

	//Запролнения массивов с заказами хоть чем-то
	orders = append(orders, Order{ID: 1, Dish: "Пицца Пеперони", Quantity: 1, IsReady: true, IsServed: false}, Order{2, "Пельмени жаренные", 2, false, false})

}

// Метод запуска сервера для Api
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.address, api.router)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
