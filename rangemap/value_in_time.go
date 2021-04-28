package rangemap

import "time"

// ValueInTime Value that can change over time
type ValueInTime struct {
	From  time.Time
	To    time.Time
	Value interface{}
}

// NewValueInTime sugar function to create new item. Time have to be in YYYY-MM-dd format
func NewValueInTime(from, to string, value interface{}) ValueInTime {
	return ValueInTime{
		From:  CreateDate(from),
		To:    CreateDate(to),
		Value: value,
	}
}
