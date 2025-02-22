package utils

import (
	"fmt"
	"time"
)

func GenerateCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:5]
}
