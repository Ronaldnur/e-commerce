package handler

import (
	"mongo-api/infra/database"
	"mongo-api/repository/analyticsreport_repository/analyticsreport_mg"
	"mongo-api/repository/feedback_repository/feedback_mg"
	"mongo-api/repository/order_repository/order_mg"
	"mongo-api/repository/product_repository/product_mg"
	"mongo-api/repository/promotion_repository/promotion_mg"
	"mongo-api/repository/review_repository/review_mg"
	"mongo-api/repository/transaction_repository/transaction_mg"
	"mongo-api/repository/user_repository/user_mg"
	"mongo-api/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {

	var port = ":8080"

	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	productRepo := product_mg.NewProductMg(db)

	productService := service.NewProductService(productRepo)

	productHandler := NewProductHandler(productService)

	userRepo := user_mg.NewUserMg(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	orderRepo := order_mg.NewOrderMg(db)

	orderService := service.NewOrderService(orderRepo, productRepo)

	orderHandler := NeworderHandler(orderService)

	analyticsRepo := analyticsreport_mg.NewAnalyticsMg(db)

	analyticsService := service.NewAnalyticsService(analyticsRepo)

	analyticsHandler := NewAnalyticsHandler(analyticsService)

	feedbackRepo := feedback_mg.NewFeedbackMg(db)

	feedbackService := service.NewFeedbackService(feedbackRepo)

	feedbackHandler := NewFeedbackHandler(feedbackService)

	transactionRepo := transaction_mg.NewTransactionMg(db)

	transactionService := service.NewTransactionService(transactionRepo)

	transactionHandler := NewTransactionHandler(transactionService)

	promotionRepo := promotion_mg.NewPromotionMg(db)

	promotionService := service.NewPromotionService(promotionRepo, productRepo)

	promotionHanlder := NewPromotionHandler(promotionService)

	reviewRepo := review_mg.NewReviewMg(db)

	reviewService := service.NewReviewService(reviewRepo)

	reviewHandler := NewReviewHandler(reviewService)

	authService := service.NewAuthService(userRepo, productRepo)
	route := gin.Default()

	productRoute := route.Group("/products")

	{
		productRoute.Use(authService.Authentitaction())
		productRoute.POST("/", productHandler.MakeProduct)
		productRoute.GET("/", productHandler.GetAllData)
		productRoute.GET("/:productId", productHandler.GetOneProduct)
		productRoute.PUT("/:productId", productHandler.UpdateProductById)
		productRoute.DELETE("/:productId", productHandler.DeleteProduct)
		productRoute.PATCH("/:productId", productHandler.UpdateStock)
	}

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.RegisterGin)
		userRoute.POST("/login", userHandler.LoginGin)
	}

	orderRoute := route.Group("/order")
	{
		orderRoute.POST("/", orderHandler.MakeOrder)
		orderRoute.GET("/", orderHandler.GetOrder)
		orderRoute.PATCH("/:orderId", orderHandler.UpdateStatus)

	}

	analyticsRoute := route.Group("/sales")
	{
		analyticsRoute.POST("/", analyticsHandler.CreateAnalytics)
		analyticsRoute.GET("/", authService.Authentitaction(), analyticsHandler.GetSellerReport)
	}

	feedbackRoute := route.Group("/feedback")
	{
		feedbackRoute.Use(authService.Authentitaction())
		feedbackRoute.POST("/", feedbackHandler.FeedbackSupport)
		feedbackRoute.GET("/", feedbackHandler.GetFeedback)
	}

	transactionRoute := route.Group("/transaction")
	{
		transactionRoute.Use(authService.Authentitaction())
		transactionRoute.POST("/payment", transactionHandler.CreatePayment)
		transactionRoute.GET("/payment", transactionHandler.GetPaymentSeller)
		transactionRoute.GET("balance", transactionHandler.GetBalanceSeller)
		transactionRoute.POST("/withdraw", transactionHandler.Withdraw)
		transactionRoute.POST("/balance", transactionHandler.CreateBalance)
	}

	promotionRoute := route.Group("/promotion")
	{
		promotionRoute.POST("/", promotionHanlder.MakePromotion)
		promotionRoute.Use(authService.Authentitaction())
		promotionRoute.GET("/", promotionHanlder.GetPromotion)
		promotionRoute.POST("apply", promotionHanlder.ApplyPromotion)
	}

	reviewRoute := route.Group("/review")
	{
		reviewRoute.POST("/", reviewHandler.CreateReview)
		reviewRoute.GET("/", authService.Authentitaction(), reviewHandler.GetAllDataReview)
	}
	route.Run(port)

	// http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

	// 	if r.Method == http.MethodGet {
	// 		productHandler.HttpGetAllProduct(w, r)
	// 		return
	// 	}

	// 	if r.Method == http.MethodPost {
	// 		productHandler.HttpMakeProduct(w, r)
	// 		return
	// 	}

	// })

	// http.HandleFunc("/products/:productId", func(w http.ResponseWriter, r *http.Request) {

	// 	if r.Method == http.MethodGet {
	// 		productHandler.HttpGetOneProduct(w, r)
	// 		return
	// 	}

	// 	if r.Method == http.MethodPost {
	// 		productHandler.HttpUpdateProduct(w, r)
	// 		return
	// 	}

	// 	if r.Method == http.MethodPost {
	// 		productHandler.HttpDeleteProduct(w, r)
	// 		return
	// 	}

	// })

	// http.HandleFunc("/users/register", func(w http.ResponseWriter, r *http.Request) {

	// 	if r.Method == http.MethodPost {
	// 		userHandler.Register(w, r)
	// 		return
	// 	}

	// })
	// http.HandleFunc("/users/login", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodPost {
	// 		userHandler.Login(w, r)
	// 		return
	// 	}
	// })
	// http.ListenAndServe(port, nil)
}
