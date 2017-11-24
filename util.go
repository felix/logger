package logger

import (
	"fmt"
	"strconv"
)

// ToString converts interface to string
func ToString(v interface{}) string {
	switch c := v.(type) {
	case string:
		return c
	case int:
		return strconv.FormatInt(int64(c), 10)
	case int64:
		return strconv.FormatInt(int64(c), 10)
	case int32:
		return strconv.FormatInt(int64(c), 10)
	case int16:
		return strconv.FormatInt(int64(c), 10)
	case int8:
		return strconv.FormatInt(int64(c), 10)
	case uint:
		return strconv.FormatUint(uint64(c), 10)
	case uint64:
		return strconv.FormatUint(uint64(c), 10)
	case uint32:
		return strconv.FormatUint(uint64(c), 10)
	case uint16:
		return strconv.FormatUint(uint64(c), 10)
	case uint8:
		return strconv.FormatUint(uint64(c), 10)
	default:
		return fmt.Sprintf("%v", c)
	}
}
