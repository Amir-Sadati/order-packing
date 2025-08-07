package router

import (
	"net/http"
	"time"

	_ "github.com/Amir-Sadati/order-packing/docs"
	"github.com/Amir-Sadati/order-packing/internal/handler/api"
	"github.com/Amir-Sadati/order-packing/internal/handler/api/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func New(
	packHandler *api.PackHandler,
) *gin.Engine {
	r := gin.New()
	r.Use(globalRecover())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // or specific frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Serve demo UI at root
	r.GET("/", func(c *gin.Context) {
		c.File("demo.html")
	})

	v1 := r.Group("/api/v1")
	//************** Pack Routes **************
	orderRoutes := v1.Group("/packs")
	orderRoutes.GET("/calculate", packHandler.CalculatePack)
	orderRoutes.GET("/sizes", packHandler.GetPackSizes)
	orderRoutes.POST("/sizes", packHandler.AddPackSize)
	orderRoutes.DELETE("/sizes", packHandler.RemovePackSize)

	//************** swagger Route **************
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}

func globalRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				response.WriteFailNoData(c.Writer, http.StatusInternalServerError, "Internal Server Error", "An unexpected error occurred")
			}
		}()
		c.Next()
	}
}
