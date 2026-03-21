package conv

import (
	"strconv"
)

func stringToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value string, def T) T {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return def
	}
	
	result := T(val)
	if int64(result) != val {
		return def
	}
	
	return result
}

func convStringToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value string, def T) T {
	if value == "" {
		return def
	}
	
	return stringToInt(value, def)
}

func StringToIntPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value string, def T) *T {
	if value == "" {
		return &def
	}
	
	var zero T
	
	result := convStringToInt(value, zero)
	
	return &result
}

func StringPtrToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value *string, def T) T {
	if value == nil {
		return def
	}
	
	return convStringToInt(*value, def)
}

func StringPtrToIntPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value *string, def T) *T {
	if value == nil {
		return &def
	}
	
	var zero T
	
	result := convStringToInt(*value, zero)
	
	return &result
}

func convFloatToInt[T ~float32 | ~float64](value T) int {
	return int(value)
}

func FloatToInt[T ~float32 | ~float64](value T) int {
	return convFloatToInt[T](value)
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

func convBoolToInt(value bool) int {
	if value {
		return 1
	}
	
	return 0
}

func BoolToInt(value bool) int {
	return convBoolToInt(value)
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
