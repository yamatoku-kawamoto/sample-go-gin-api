package main

import (
	"fmt"
	"ivs/api"
	"log"
	"os"
	"strconv"

	. "ivs"

	"github.com/gin-gonic/gin"
)

func main() {
	e := newEngine()
	initRoutes(e)
	addr := address(isLocal())
	log.Print(addr)
	if err := e.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func newEngine() Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	return e
}

func address(localonly ...bool) string {
	// TODO: 環境変数でPORTを取得する
	const port uint16 = 10943
	// TODO: 環境変数でHOSTTを取得する
	var host string
	if len(localonly) > 0 && localonly[0] {
		host = "localhost"
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func isLocal() bool {
	v := os.Getenv("LOCALONLY")
	if v == "" {
		return false
	}
	local, err := strconv.ParseBool(v)
	return local || err != nil
}

func initRoutes(e Engine) Engine {
	e.GET("/", api.HandleHelthCheck)
	e.POST("/webhook", api.Webhook)
	e.POST("/api/v1/verifications", api.HandleVerificationRequest)
	return e
}
