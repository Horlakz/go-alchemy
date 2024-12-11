package services

import (
	"os"

	"github.com/samber/lo"
)

func GetEnv[T interface{}](key string, fallback T) T {
	var v interface{} = os.Getenv(key)
	var _v string = v.(string)
	if lo.IsEmpty(_v) {
		return fallback
	}

	return v.(T)
}
