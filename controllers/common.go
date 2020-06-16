package controllers

import (
	"auuuth/models"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

const (
	minUnicodeRuneValue = 0 //U+0000
	maxUnicodeRuneValue = utf8.MaxRune
)

func okAndData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    models.OK,
		"message": "OK",
		"data":    data,
	})
	return
}

func errAndInfo(c *gin.Context, errCode models.ERRCODE) {
	c.JSON(200, gin.H{
		"code":    errCode,
		"message": models.ErrMap[errCode],
		"data":    nil,
	})
	return
}

func checkKey(key string) bool {
	if !utf8.ValidString(key) {
		return false
	}

	for _, runeValue := range key {
		if runeValue == minUnicodeRuneValue || runeValue == maxUnicodeRuneValue {
			return false
		}
	}

	return true
}
