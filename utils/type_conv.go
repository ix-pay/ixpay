package utils

import "strconv"

// StringToInt 字符串转int（失败返回默认值0）
func StringToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

// StringToInt64 字符串转int64（失败返回默认值0）
func StringToInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

// IntToString 整数转字符串（无失败情况）
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToString int64转字符串（无失败情况）
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Float64ToString float64转字符串（无失败情况）
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToFloat64 字符串转float64（失败返回默认值0）
func StringToFloat64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

// InterfaceToInt 接口转int（支持string/float64/uint等类型）
func InterfaceToString(i interface{}) string {
	switch v := i.(type) {
	case int:
		return IntToString(v)
	case string:
		return v
	case float64:
		return Float64ToString(v)
	default:
		return ""
	}
}

// InterfaceToInt 接口转int（支持string/float64/uint等类型）
func InterfaceToInt(i interface{}) int {
	switch v := i.(type) {
	case int:
		return v
	case string:
		return StringToInt(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}

// InterfaceToInt64 接口转int64（支持string/float64等类型）
func InterfaceToInt64(i interface{}) int64 {
	switch v := i.(type) {
	case int64:
		return v
	case string:
		return StringToInt64(v)
	case float64:
		return int64(v)
	default:
		return 0
	}
}
