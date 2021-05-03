package rangemap

import "time"

// ValueInTime Value that can change over time
type ValueInTime struct {
	From  time.Time
	To    time.Time
	Value interface{}
}

// MustValueInTime sugar function to create new item. Time have to be in YYYY-MM-dd format
func MustValueInTime(from, to string, value interface{}) ValueInTime {
	return ValueInTime{
		From:  MustCreateDate(from),
		To:    MustCreateDate(to),
		Value: value,
	}
}

func NewValueInTime(from, to time.Time, value interface{}) ValueInTime {
	return ValueInTime{
		From:  from,
		To:    to,
		Value: value,
	}
}
