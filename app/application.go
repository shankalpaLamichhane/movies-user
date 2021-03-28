package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var (
	router = gin.Default()
)

func StartApplication(){
	mapUrls()
	log.Info().Msg("starting the application now")
	router.Run(":8080")
}