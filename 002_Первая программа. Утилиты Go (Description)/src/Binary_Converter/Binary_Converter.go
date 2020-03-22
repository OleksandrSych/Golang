package Binary_Converter

import (
	"strconv"
)

func Binary_Converter(number int64) string {
	return strconv.FormatInt(number, 2)
}
