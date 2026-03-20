package conv

import (
	"strconv"
)

func stringToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value string, def T) T {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return def
	}
	
	result := T(v)
	if int64(result) != v {
		return def
	}
	
	return result
}

func StringToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value string, def T) T {
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
	
	result := StringToInt(value, zero)
	
	return &result
}

func StringPtrToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value *string, def T) T {
	if value == nil {
		return def
	}
	
	return StringToInt(*value, def)
}

func StringPtrToIntPtr[T ~int | ~int8 | ~int16 | ~int32 | ~int64](value *string, def T) *T {
	if value == nil {
		return &def
	}
	
	var zero T
	
	result := StringToInt(*value, zero)
	
	return &result
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
