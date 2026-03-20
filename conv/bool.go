package conv

import (
	"strconv"
)

func stringToBool(value string) bool {
	if value == "" {
		return false
	}
	
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	
	return result
}

func StringToBool(value string) bool {
	return stringToBool(value)
}

func StringToBoolPtr(value string) *bool {
	if value == "" {
		return nil
	}
	
	result, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}
	
	return &result
}

func StringPtrToBool(value *string) bool {
	if value == nil {
		return false
	}
	
	return StringToBool(*value)
}

func StringPtrToBoolPtr(value *string) *bool {
	if value == nil {
		return nil
	}
	
	return StringToBoolPtr(*value)
}
