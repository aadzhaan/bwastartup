package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// input
// handler mapping input ke struct
// service mapping ke struct (ex:User)
// repository save struct ke db
// db

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	router.Run()

}

//MIDDLEWARE
//ambil nilai header Authorization: Bearer token
//dari header Authorization, kita ambil nilai tokennya saja
//kita validasi token
//ambil nilai user_id (yang ada di payload dari jwt)
//ambil user dari db berdasarkan user_id lewat service
//kita set context isinya user

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			//menggunakan abort karena middleware ada ditengah2, biar tidak lanjut ke proses selanjutnya jika tidak benar
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Bearer token
		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			//menggunakan abort karena middleware ada ditengah2, biar tidak lanjut ke proses selanjutnya jika tidak benar
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			//menggunakan abort karena middleware ada ditengah2, biar tidak lanjut ke proses selanjutnya jika tidak benar
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userId)

		if err != nil {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			//menggunakan abort karena middleware ada ditengah2, biar tidak lanjut ke proses selanjutnya jika tidak benar
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
