package helpers

import (
	"mongo-api/pkg/errs"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetParamId(r *http.Request, key string) (int, errs.MessageErr) {
	// Ambil path dari request
	path := r.URL.Path

	// Pisahkan path berdasarkan "/"
	parts := strings.Split(path, "/")

	// Temukan indeks parameter dalam path
	for i, part := range parts {
		if part == key {
			if i+1 < len(parts) {
				value := parts[i+1]
				id, err := strconv.Atoi(value)
				if err != nil {
					return 0, errs.NewBadRequest("invalid parameter id")
				}
				return id, nil
			}
			return 0, errs.NewBadRequest("parameter not found")
		}
	}
	return 0, errs.NewBadRequest("parameter not found")
}

func GetQueryFloat(ctx *gin.Context, key string) (float64, errs.MessageErr) {
	value := ctx.Query(key)

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, errs.NewBadRequest("invalid parameter float")
	}

	return floatValue, nil
}
