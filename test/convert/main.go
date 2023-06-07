package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func parseInterfaceToString(val interface{}) string {
	switch v := val.(type) {
	case nil:
		return ""
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case *int:
		return strconv.Itoa(*v)
	case *float64:
		return strconv.FormatFloat(*v, 'f', -1, 64)
	case *string:
		return *v
	default:
		if str, ok := v.(string); ok {
			return str
		}
		return fmt.Sprintf("%v", val)
	}
}
