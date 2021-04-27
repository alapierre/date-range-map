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

type RangeMap struct {
	keys   []TimeRange
	values []string
}

// New create new instance of TimeRange map. Keys have to be sorted!
func New(ranges []TimeRange, values []string) (*RangeMap, error) {
	if len(ranges) != len(values) {
		return nil, fmt.Errorf("size of keys and values should be the same")
	}
	return &RangeMap{
		keys:   ranges,
		values: values,
	}, nil
}

// Get works like Guava RangeMap com.google.common.collect.ImmutableRangeMap
func (rm RangeMap) Get(key time.Time) (string, bool) {
	count := len(rm.keys)
	i := sort.Search(count, func(i int) bool {
		return key.Before(rm.keys[i].L)
	})

	i -= 1
	if i >= 0 && i < count && !key.After(rm.keys[i].U) {
		return rm.values[i], true
	}
	return "", false
}

func CreateDate(str string) time.Time {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		fmt.Printf("CreateDate: date in wrong format %s - expected YYYY-MM-dd\n", str)
	}
	return t
}

func CreateRange(from, to string) TimeRange {
	return TimeRange{
		L: CreateDate(from),
		U: CreateDate(to),
	}
}
