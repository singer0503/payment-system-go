package main

import (
	"go-payment-system/controllers"
	"go-payment-system/database"
	"go-payment-system/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	//database.Connect("root:root@tcp(localhost:3306)/jwt_demo?parseTime=true") // mysql
	database.Connect("server=127.0.0.1;user id=sa;password=P@ssw0rd;port=1433;database=gp_db_tx;") // sqlserver
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	// api := router.Group("/api")
	// {
	// 	api.POST("/users/authenticate", controllers.GenerateToken)
	// 	api.POST("/user/register", controllers.RegisterUser)
	// 	secured := api.Group("/secured").Use(middlewares.Auth())
	// 	{
	// 		secured.GET("/ping", controllers.Ping)
	// 	}
	// }
	router.POST("/users/authenticate", controllers.GenerateToken)
	router.POST("/user/register", controllers.RegisterUser)
	routerAuth := router.Use(middlewares.Auth()) // routerAuth 是要登入才能使用的路由, 通過 Auth() 中介軟體驗證
	{
		routerAuth.GET("/ping", controllers.Ping)
		routerAuth.GET("/payment/payout", controllers.GetPayoutData)
	}
	return router
}
