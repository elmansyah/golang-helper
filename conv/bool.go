package conv

import (
	"strconv"
)

func convStringToBool(value string, def bool) bool {
	if value == "" {
		return def
	}
	
	result, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	
	return result
}

func StringToBool(value string, def bool) bool {
	return convStringToBool(value, def)
}

func StringToBoolPtr(value string, def bool) *bool {
	if value == "" {
		return &def
	}
	
	result, err := strconv.ParseBool(value)
	if err != nil {
		return &def
	}
	
	return &result
}

func StringPtrToBool(value *string, def bool) bool {
	if value == nil {
		return def
	}
	
	return StringToBool(*value, def)
}

func StringPtrToBoolPtr(value *string, def bool) *bool {
	if value == nil {
		return nil
	}
	
	return StringToBoolPtr(*value, def)
}
