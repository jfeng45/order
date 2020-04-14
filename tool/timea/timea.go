package timea

import "time"

// The following constants represents different data format layout
const (
	FORMAT_ISO8601_DATE                 = "2006-01-02"
	FORMAT_ISO8601_DATE_TIME            = "2006-01-02 15:04:05"
	FORMAT_ISO8601_DATE_TIME_MILLI      = "2006-01-02 15:04:05.000"
	FORMAT_ISO8601_DATE_TIME_MILLI_ZONE = "2006-01-02 15:04:05.000Z07:00"
	FORMAT_ISO8601_DATE_TIME_MICRO      = "2006-01-02 15:04:05.000000"
	FORMAT_ISO8601_DATE_TIME_MICRO_ZONE = "2006-01-02 15:04:05.000000Z07:00"
	FORMAT_ISO8601_DATE_TIME_NANO       = "2006-01-02 15:04:05.000000000"
	FORMAT_ISO8601_DATE_TIME_NANO_ZONE  = "2006-01-02 15:04:05.00000000007:00"

	TIME_NULL_STRING = "0000-00-00 00:00:00"

)

func Parse(layout, value string) (time.Time, error) {
	var t time.Time
	if value == TIME_NULL_STRING {
		return t, nil
	} else {
		return time.Parse(layout, value)
	}

}
