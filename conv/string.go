package conv

import (
	"strconv"
)

var (
	trueString = "true"
)

func convIntToString[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value T) string {
	return strconv.FormatInt(int64(value), 10)
}

func IntToString[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value T) string {
	if value == 0 {
		return ""
	}
	
	return convIntToString(value)
}

func IntToStringPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value T) *string {
	if value == 0 {
		return nil
	}
	
	result := convIntToString(value)
	
	return &result
}

func IntPtrToString[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value *T) string {
	if value == nil {
		return ""
	}
	
	result := convIntToString(*value)
	
	return result
}

func IntPtrToStringPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value *T) *string {
	if value == nil {
		return nil
	}
	
	result := convIntToString(*value)
	
	return &result
}

func convUintToString[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](value T) string {
	return strconv.FormatUint(uint64(value), 10)
}

func UintToString[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](value T) string {
	if value == 0 {
		return ""
	}
	
	return convUintToString(value)
}

func UintToStringPtr[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](value T) *string {
	if value == 0 {
		return nil
	}
	
	result := convUintToString(value)
	
	return &result
}

func UintPtrToString[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](value *T) string {
	if value == nil {
		return ""
	}
	
	return convUintToString(*value)
}

func UintPtrToStringPtr[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](value *T) *string {
	if value == nil {
		return nil
	}
	
	result := convUintToString(*value)
	
	return &result
}

func convFloatToString[T ~float32 | ~float64](value T) string {
	var bitSize int
	
	switch any(value).(type) {
	case float32:
		bitSize = 32
	default:
		bitSize = 64
	}
	
	return strconv.FormatFloat(float64(value), 'f', -1, bitSize)
}

func FloatToString[T ~float32 | ~float64](value T) string {
	if value == 0 {
		return ""
	}
	
	return convFloatToString(value)
}

func FloatToStringPtr[T ~float32 | ~float64](value T) *string {
	if value == 0 {
		return nil
	}
	
	result := convFloatToString(value)
	
	return &result
}

func FloatPtrToString[T ~float32 | ~float64](value *T) string {
	if value == nil {
		return ""
	}
	
	return convFloatToString(*value)
}

func FloatPtrToStringPtr[T ~float32 | ~float64](value *T) *string {
	if value == nil {
		return nil
	}
	
	result := convFloatToString(*value)
	
	return &result
}

func convBoolToString(value bool) string {
	if !value {
		return ""
	}
	
	return trueString
}

func BoolToString(value bool) string {
	return convBoolToString(value)
}

func BoolToStringPtr(value bool) *string {
	if !value {
		return nil
	}
	
	result := "true"
	
	return &result
}

func BoolPtrToString(value *bool) string {
	if value == nil {
		return ""
	}
	
	return convBoolToString(*value)
}

func BoolPtrToStringPtr(value *bool) *string {
	if value == nil || !*value {
		return nil
	}
	
	result := "true"
	
	return &result
}
