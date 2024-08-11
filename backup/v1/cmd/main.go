package main

import (
	"SoCC/pkg/api"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func runServer() {
	fmt.Println("Подгрузка конфигов")
	if err := initConfig(); err != nil {
		logrus.Fatalf("error: failed config read attempt: %s", err.Error())
	}
	/*if err := godotenv.Load(); err != nil {
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
	postgres.Initial(db)*/
	port := viper.Get("port").(string)
	api := api.New("localhost:"+port, http.NewServeMux())
	api.FillEndpoints()

	fmt.Println("Сервер запущен на порту " + port)
	logrus.Fatal(api.ListenAndServe())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	runServer()
}
