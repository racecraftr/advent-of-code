package conv

import (
	"adventOfCode/util"
	"fmt"
	"reflect"
	"strconv"
)

func ToInt(v any) int {
	switch v := v.(type) {
	case int:
		return v
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		i, err := strconv.Atoi(v)
		util.Check(err)
		return i
	case rune:
		if v <= '0' && v >= '9' {
			return int(v - '0')
		}
	}
	panic(fmt.Sprintf("Could not parse value of type %v to int", reflect.TypeOf(v)))
}

func ToString(v any) string {
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case byte:
		return string(v)
	case rune:
		return string(v)
	}
	return fmt.Sprintf("%v", v)
}
