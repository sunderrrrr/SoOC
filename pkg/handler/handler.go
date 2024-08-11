package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	auth := router.Group("/api/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signUp)
	}

	api := router.Group("/api/")
	{
		orders := api.Group("/order")
		{
			orders.POST("/create", h.createOrder)
			orders.GET("/view/", h.listOrder)
			orders.DELETE("/delete/:id", h.deleteOrder)
			orders.PUT("/update/:id", h.updateOrder)

			dishes := orders.Group("/dishes")
			{
				dishes.POST("/create", h.createDish)
				dishes.GET("/list/", h.listDishes)
				dishes.GET("/contain/:id", h.InsideDish)
				dishes.DELETE("/delete/:id", h.deleteDish)
				dishes.PUT("/update/:id", h.updateDish)
			}
		}

	}
	return router
}
