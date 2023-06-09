package utils

import (
	"math/rand"
	"path/filepath"

	oauthDto "edtech-app/internal/oauth/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RandString(number int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, number)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func Paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset

		// Jika page isinya kurang atau sama dengan 0 kita akan ganti menjadi 1
		if page <= 0 {
			page = 1
		}

		pageSize := limit

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(pageSize)

	}
}

func GetCurrentUser(ctx *gin.Context) *oauthDto.MapClaimsResponse {
	user, _ := ctx.Get("user")

	return user.(*oauthDto.MapClaimsResponse)
}

func GetFileName(fileName string) string {
	file := filepath.Base(fileName)

	return file[:len(file)-len(filepath.Ext(file))]
}
