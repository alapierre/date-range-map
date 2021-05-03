package rangemap

import (
	"testing"
	"time"
)

func TestRangeMap_Get(t *testing.T) {

	type fields struct {
		Keys   []TimeRange
		Values []interface{}
	}

	type args struct {
		key time.Time
	}

	values := make([]interface{}, 0)
	values = append(values, "a", "b", "c", "d", "e")

	inputData := fields{
		Keys: []TimeRange{
			MustCreateRange("2021-01-01", "2021-01-31"),
			MustCreateRange("2021-03-01", "2021-04-30"),
			MustCreateRange("2021-05-15", "2021-05-31"),
			MustCreateRange("2021-06-15", "2021-06-20"),
			MustCreateRange("2021-07-15", "2021-09-20")},
		Values: values,
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		ok     bool
	}{
		{
			name:   "begin of date range in a middle",
			fields: inputData,
			args:   args{MustCreateDate("2021-03-01")},
			want:   "b",
			ok:     true},
		{
			name:   "inside of date range in a middle",
			fields: inputData,
			args:   args{MustCreateDate("2021-03-02")},
			want:   "b",
			ok:     true},
		{
			name:   "end of date range in a middle",
			fields: inputData,
			args:   args{MustCreateDate("2021-04-30")},
			want:   "b",
			ok:     true},
		{
			name:   "begin of date range in a first element",
			fields: inputData,
			args:   args{MustCreateDate("2021-01-01")},
			want:   "a",
			ok:     true},
		{
			name:   "begin of date range in a last element",
			fields: inputData,
			args:   args{MustCreateDate("2021-07-15")},
			want:   "e",
			ok:     true},
		{
			name:   "end of date range in a last element",
			fields: inputData,
			args:   args{MustCreateDate("2021-09-20")},
			want:   "e",
			ok:     true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := rangeMap{
				keys:   tt.fields.Keys,
				values: tt.fields.Values,
			}
			got, got1 := rm.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Get() ok = %v, want %v", got1, tt.ok)
			}
		})
	}
}
