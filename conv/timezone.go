package conv

import (
	"sync"
	"time"
)

var (
	defaultTimezone = time.UTC
	tzMutex         sync.RWMutex
)

func SetTimezone(value *time.Location) {
	if value == nil {
		return
	}
	
	tzMutex.Lock()
	
	defaultTimezone = value
	
	tzMutex.Unlock()
}
