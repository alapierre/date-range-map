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

// ValueInTime TODO: powinna być funkcja do tworzenia instancji (może do innego pliku?)
type ValueInTime struct {
	From  time.Time
	To    time.Time
	Value interface{}
}

type RangeMap interface {
	Get(key time.Time) (interface{}, bool)
}

type rangeMap struct {
	keys   []TimeRange
	values []interface{}
}

// New create new instance of RangeMap map. Keys have to be sorted!
func New(values []ValueInTime) (RangeMap, error) {

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
	i := sort.Search(count, func(i int) bool {
		return key.Before(rm.keys[i].L)
	})

	i -= 1
	if i >= 0 && i < count && !key.After(rm.keys[i].U) {
		return rm.values[i], true
	}
	return nil, false
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

func sortValueInTime(values []ValueInTime) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].From.Before(values[j].From)
	})
}
