package internal

import (
	"fmt"
	"strconv"
	"time"
)

// ToString converts interface to string
func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch c := v.(type) {
	case string:
		return c
	case *string:
		return *c
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
	case float32:
		return strconv.FormatFloat(float64(c), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(c, 'g', -1, 64)
	case bool:
		return strconv.FormatBool(c)
	case *bool:
		return strconv.FormatBool(*c)
	case *time.Time:
		return fmt.Sprintf("%s", c)
	case fmt.Stringer:
		return c.String()
	default:
		return fmt.Sprintf("%v", c)
	}
}
