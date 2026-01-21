// @title Simple Procurement System
// @version 1.0
// @description REST API untuk Simple Procurement System
// @termsOfService http://swagger.io/terms/

// @contact.name Restu Adi Wahyujati
// @contact.email adijati1029@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8888
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" + a space + your JWT token.

package main

func main()  {
	server := NewServer()
	server.Run()
}