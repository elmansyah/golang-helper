package conv

import (
	"strconv"
)

func stringToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value string) (T, error) {
	var bitSize int
	
	switch any(*new(T)).(type) {
	case int8:
		bitSize = 8
	case int16:
		bitSize = 16
	case int32:
		bitSize = 32
	case int64:
		bitSize = 64
	default:
		bitSize = strconv.IntSize
	}
	
	result, err := strconv.ParseInt(value, 10, bitSize)
	if err != nil {
		return 0, err
	}
	
	return T(result), nil
}

func StringToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value string) (T, error) {
	if value == "" {
		return 0, nil
	}
	
	return stringToInt[T](value)
}

func StringToIntPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value string) (*T, error) {
	if value == "" {
		return nil, nil
	}
	
	result, err := stringToInt[T](value)
	if err != nil {
		return nil, err
	}
	
	return &result, nil
}

func StringPtrToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value *string) (T, error) {
	if value == nil {
		return 0, nil
	}
	
	return StringToInt[T](*value)
}

func StringPtrToIntPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value *string) (*T, error) {
	if value == nil {
		return nil, nil
	}
	
	return StringToIntPtr[T](*value)
}

func floatToInt[T ~float32 | ~float64](value T) int {
	return int(value)
}

func FloatToInt[T ~float32 | ~float64](value T) int {
	return floatToInt[T](value)
}

func FloatToIntPtr[T ~float32 | ~float64](value *T) *int {
	if value == nil {
		return nil
	}
	
	result := FloatToInt[T](*value)
	
	return &result
}

func FloatPtrToInt[T ~float32 | ~float64](value *T) int {
	if value == nil {
		return 0
	}
	
	return FloatToInt[T](*value)
}

func FloatPtrToIntPtr[T ~float32 | ~float64](value *T) *int {
	if value == nil {
		return nil
	}
	
	result := FloatToInt[T](*value)
	
	return &result
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	
	return 0
}

func BoolToInt(value bool) int {
	return boolToInt(value)
}

func BoolToIntPtr(value bool) *int {
	if !value {
		return nil
	}
	
	result := 1
	
	return &result
}

func BoolPtrToInt(value *bool) int {
	if value == nil || !*value {
		return 0
	}
	
	return 1
}

func BoolPtrToIntPtr(value *bool) *int {
	if value == nil || !*value {
		return nil
	}
	
	result := 1
	
	return &result
}
