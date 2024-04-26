package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/muharik19/shop-backend/models"
	logger "github.com/muharik19/shop-backend/pkg/logging"
)

// Response setting gin.JSON
func Response(c *gin.Context, req interface{}, res models.Response) {
	// LOGGER
	reqByte, _ := json.Marshal(req)
	resByte, _ := json.Marshal(res)
	logger.Infof("[shop-backend:log] [RequestURL] : %s, [RequestMethod] : %s, [RequestBody] : %s, [ResponseData] : %s", c.Request.RequestURI, c.Request.Method, string(reqByte), string(resByte))

	c.JSON(res.Code, res)
}
