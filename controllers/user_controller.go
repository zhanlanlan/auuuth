package controllers

import (
	"auuuth/models"
	"auuuth/services"
	"auuuth/utils/logger"

	"github.com/gin-gonic/gin"
)

// AddUser ...
func AddUser(c *gin.Context) {
	// get api input and unmarshal to json
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error(err)
		errAndInfo(c, models.BadRequestErr)
		return
	}

	// check if key is valid
	if !checkKey(user.ID) {
		logger.Error(models.ErrMap[models.InvalidKeyErr])
		errAndInfo(c, models.BadRequestErr)
		return
	}

	// store in leveldb
	if err := services.AddUser(user); err != nil {
		// 无法存入数据库就是服务器内部错误
		logger.Error(err)
		errAndInfo(c, models.InternalErr)
		return
	}

	okButNoData(c)
	return
}

// DeleteUSer ...
func DeleteUSer(c *gin.Context) {

}
