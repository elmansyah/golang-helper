package conv

import (
	"fmt"
	"strconv"
)

func convIntToFloat[T ~int | ~int8 | ~int16 | ~int32 | ~int64, Result ~float32 | ~float64](value T) Result {
	return Result(value)
}

func IntToFloat[T ~int | ~int8 | ~int16 | ~int32 | ~int64, Result ~float32 | ~float64](value T) Result {
	if value == 0 {
		return 0
	}
	
	return convIntToFloat[T, Result](value)
}

func IntToFloatPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64, Result ~float32 | ~float64](value T) *Result {
	if value == 0 {
		return nil
	}
	
	result := convIntToFloat[T, Result](value)
	
	return &result
}

func IntPtrToFloat[T ~int | ~int8 | ~int16 | ~int32 | ~int64, Result ~float32 | ~float64](value *T) Result {
	if value == nil {
		return 0
	}
	
	return convIntToFloat[T, Result](*value)
}

func IntPtrToFloatPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64, Result ~float32 | ~float64](value *T) *Result {
	if value == nil {
		return nil
	}
	
	result := convIntToFloat[T, Result](*value)
	
	return &result
}

func convStringToFloat[T ~float32 | ~float64](value string) (T, error) {
	var bitSize int
	
	switch any(*new(T)).(type) {
	case float32:
		bitSize = 32
	default:
		bitSize = 64
	}
	
	result, err := strconv.ParseFloat(value, bitSize)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	
	return T(result), nil
}

func StringToFloat[T ~float32 | ~float64](value string) (T, error) {
	if value == "" {
		return 0, nil
	}
	
	return convStringToFloat[T](value)
}

func StringToFloatPtr[T ~float32 | ~float64](value string) (*T, error) {
	if value == "" {
		return nil, nil
	}
	
	result, err := convStringToFloat[T](value)
	if err != nil {
		return nil, err
	}
	
	return &result, nil
}

func StringPtrToFloat[T ~float32 | ~float64](value *string) (T, error) {
	if value == nil {
		return 0, nil
	}
	
	return StringToFloat[T](*value)
}

func StringPtrToFloatPtr[T ~float32 | ~float64](value *string) (*T, error) {
	if value == nil {
		return nil, nil
	}
	
	return StringToFloatPtr[T](*value)
}

func convBoolToFloat[T ~float32 | ~float64](value bool) T {
	if value {
		return 1
	}
	
	return 0
}

func BoolToFloat[T ~float32 | ~float64](value bool) T {
	return convBoolToFloat[T](value)
}

func BoolToFloatPtr[T ~float32 | ~float64](value bool) *T {
	if !value {
		return nil
	}
	
	result := BoolToFloat[T](value)
	
	return &result
}

func BoolPtrToFloat[T ~float32 | ~float64](value *bool) T {
	if value == nil || !*value {
		return 0
	}
	
	return BoolToFloat[T](*value)
}

func BoolPtrToFloatPtr[T ~float32 | ~float64](value *bool) *T {
	if value == nil || !*value {
		return nil
	}
	
	result := BoolToFloat[T](*value)
	
	return &result
}
