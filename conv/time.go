package conv

import (
	"time"
)

const (
	layoutDateTime         = "2006-01-02 15:04:05"
	layoutDate             = "2006-01-02"
	layoutMonth            = "2006-01"
	layoutTime             = "15:04:05"
	layoutCalendarDateTime = "02 Jan 2006 15:04:05"
	layoutCalendarDate     = "02 Jan 2006"
	layoutISO              = time.RFC3339
)

func format(value time.Time, layout string) string {
	if value.IsZero() {
		return ""
	}
	
	tzMutex.RLock()
	
	tz := defaultTimezone
	
	tzMutex.RUnlock()
	
	return value.In(tz).Format(layout)
}

func formatPtr(value time.Time, layout string) *string {
	if value.IsZero() {
		return nil
	}
	
	result := format(value, layout)
	
	return &result
}

func formatFromPtr(value *time.Time, layout string) string {
	if value == nil || value.IsZero() {
		return ""
	}
	
	return format(*value, layout)
}

func formatPtrFromPtr(value *time.Time, layout string) *string {
	if value == nil || value.IsZero() {
		return nil
	}
	
	result := formatFromPtr(value, layout)
	
	return &result
}

func DateTime(value time.Time) string {
	return format(value, layoutDateTime)
}

func DateTimePtr(value time.Time) *string {
	return formatPtr(value, layoutDateTime)
}

func DateTimeFromPtr(value *time.Time) string {
	return formatFromPtr(value, layoutDateTime)
}

func DateTimePtrFromPtr(value *time.Time) *string {
	return formatPtrFromPtr(value, layoutDateTime)
}

func Date(value time.Time) string {
	return format(value, layoutDate)
}

func DatePtr(value time.Time) *string {
	return formatPtr(value, layoutDate)
}

func DateFromPtr(value *time.Time) string {
	return formatFromPtr(value, layoutDate)
}

func DatePtrFromPtr(value *time.Time) *string {
	return formatPtrFromPtr(value, layoutDate)
}

func Time(value time.Time) string {
	return format(value, layoutTime)
}

func TimePtr(value time.Time) *string {
	return formatPtr(value, layoutTime)
}

func TimeFromPtr(value *time.Time) string {
	return formatFromPtr(value, layoutTime)
}

func TimePtrFromPtr(value *time.Time) *string {
	return formatPtrFromPtr(value, layoutTime)
}

func Timestamp(value time.Time) int64 {
	if value.IsZero() {
		return 0
	}
	
	return value.Unix()
}

func TimestampPtr(value time.Time) *int64 {
	if value.IsZero() {
		return nil
	}
	
	result := Timestamp(value)
	
	return &result
}

func TimestampFromPtr(value *time.Time) int64 {
	if value == nil || value.IsZero() {
		return 0
	}
	
	return value.Unix()
}

func TimestampPtrFromPtr(value *time.Time) *int64 {
	if value == nil || value.IsZero() {
		return nil
	}
	
	result := TimestampFromPtr(value)
	
	return &result
}

func CalendarDateTime(value time.Time) string {
	return format(value, layoutCalendarDateTime)
}

func CalendarDateTimePtr(value time.Time) *string {
	return formatPtr(value, layoutCalendarDateTime)
}

func CalendarDateTimeFromPtr(value *time.Time) string {
	return formatFromPtr(value, layoutCalendarDateTime)
}

func CalendarDateTimePtrFromPtr(value *time.Time) *string {
	return formatPtrFromPtr(value, layoutCalendarDateTime)
}

func CalendarDate(value time.Time) string {
	return format(value, layoutCalendarDate)
}

func CalendarDatePtr(value time.Time) *string {
	return formatPtr(value, layoutCalendarDate)
}

func CalendarDateFromPtr(value *time.Time) string {
	return formatFromPtr(value, layoutCalendarDate)
}

func CalendarDatePtrFromPtr(value *time.Time) *string {
	return formatPtrFromPtr(value, layoutCalendarDate)
}

func ISO(value time.Time) string {
	return format(value, layoutISO)
}

func ISOPtr(value time.Time) *string {
	return formatPtr(value, layoutISO)
}

func ISOFromPtr(value *time.Time) string {
	return formatFromPtr(value, layoutISO)
}

func ISOPtrFromPtr(value *time.Time) *string {
	return formatPtrFromPtr(value, layoutISO)
}

func TimestampMilli(value time.Time) int64 {
	if value.IsZero() {
		return 0
	}
	
	return value.UnixMilli()
}

func TimestampMilliPtr(value time.Time) *int64 {
	if value.IsZero() {
		return nil
	}
	
	result := TimestampMilli(value)
	
	return &result
}

func TimestampMilliFromPtr(value *time.Time) int64 {
	if value == nil || value.IsZero() {
		return 0
	}
	
	return value.UnixMilli()
}

func TimestampMilliPtrFromPtr(value *time.Time) *int64 {
	if value == nil || value.IsZero() {
		return nil
	}
	
	result := TimestampMilliFromPtr(value)
	
	return &result
}

func ParseDateTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	
	result, err := time.Parse(layoutDate, value)
	if err != nil {
		return time.Time{}
	}
	
	if result.IsZero() {
		//nolint:revive
		return time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	
	return result
}

func ParseDateTimePtr(value string) *time.Time {
	if value == "" {
		return nil
	}
	
	result := ParseDateTime(value)
	
	return &result
}

func ParseDateTimeFromPtr(value *string) time.Time {
	if value == nil || *value == "" {
		return time.Time{}
	}
	
	return ParseDateTime(*value)
}

func ParseDateTimePtrFromPtr(value *string) *time.Time {
	if value == nil || *value == "" {
		return nil
	}
	
	result := ParseDateTimeFromPtr(value)
	
	return &result
}

func ParseDate(value string) time.Time {
	return ParseDateTime(value)
}

func StringToTimeDuration(value string) time.Duration {
	if value == "" {
		return time.Duration(0)
	}
	
	result, err := time.ParseDuration(value)
	if err != nil {
		return time.Duration(0)
	}
	
	return result
}

func StringToTimeDurationPtr(value string) *time.Duration {
	if value == "" {
		return nil
	}
	
	result := StringToTimeDuration(value)
	
	return &result
}

func StringPtrToTimeDuration(value *string) time.Duration {
	if value == nil || *value == "" {
		return time.Duration(0)
	}
	
	return StringToTimeDuration(*value)
}

func StringPtrToTimeDurationPtr(value *string) *time.Duration {
	if value == nil || *value == "" {
		return nil
	}
	
	result := StringPtrToTimeDuration(value)
	
	return &result
}

func LogTime() string {
	timezone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return ""
	}
	
	return time.Now().In(timezone).Format("2006/01/02 15:04:05")
}
