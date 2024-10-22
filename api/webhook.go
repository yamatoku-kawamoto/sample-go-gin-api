package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "ivs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Webhook(ctx HttpContext) {
	var requestJson = make(gin.H)
	ctx.Bind(&requestJson)
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(requestJson); err != nil {
		log.Print(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println(buf.String())
	ctx.Status(http.StatusNoContent)
}
