package rangemap

import (
	"fmt"
	"sort"
	"time"
)

type TimeRange struct {
	L time.Time
	U time.Time
}

type RangeMap interface {
	Get(key time.Time) (interface{}, bool)
}

type rangeMap struct {
	keys   []TimeRange
	values []interface{}
}

// New create new instance of RangeMap map. Values will be sorted!!!
func New(values []ValueInTime) (RangeMap, error) {

	sortValueInTime(values)

	result := rangeMap{}

	for _, value := range values {
		result.keys = append(result.keys, TimeRange{L: value.From, U: value.To})
		result.values = append(result.values, value.Value)
	}

	return &result, nil
}

// Get works like Guava RangeMap com.google.common.collect.ImmutableRangeMap
func (rm rangeMap) Get(key time.Time) (interface{}, bool) {
	count := len(rm.keys)
	// look for 'next' period
	i := sort.Search(count, func(i int) bool {
		return key.Before(rm.keys[i].L)
	})

	i-- // if we found something it will be 'next' period, so we need take previews
	if i >= 0 && i < count && !key.After(rm.keys[i].U) {
		return rm.values[i], true
	}
	return nil, false
}

// MustCreateDate parse time in YYYY-MM-dd format into time.Time. Parsing errors will panic!
func MustCreateDate(str string) time.Time {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		panic(fmt.Errorf("MustCreateDate: date in wrong format %s - expected YYYY-MM-dd, %v", str, err))
	}
	return t
}

// MustCreateRange parse time in YYYY-MM-dd format into TimeRange. Parsing errors will panic!
func MustCreateRange(from, to string) TimeRange {
	return TimeRange{
		L: MustCreateDate(from),
		U: MustCreateDate(to),
	}
}

func sortValueInTime(values []ValueInTime) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].From.Before(values[j].From)
	})
}
