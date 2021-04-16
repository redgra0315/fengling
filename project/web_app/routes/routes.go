package routes

import (
	"github.com/fengling/project/log_demo/project/web_app/logger"
	"github.com/fengling/project/log_demo/project/web_app/settings"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Setup(cfg *settings.AppConfig) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": viper.GetString("app.version"), "status": http.StatusOK})
	})

	return r
}
