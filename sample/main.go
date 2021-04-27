package main

import (
	"fmt"
	"github.com/alapierre/date-range-map/rangemap"
	"sort"
)

func main() {

	//keys := []rangemap.TimeRange{
	//	rangemap.CreateRange("2021-01-01", "2021-01-31"),
	//	rangemap.CreateRange("2021-03-01", "2021-04-30"),
	//	rangemap.CreateRange("2021-05-15", "2021-05-31"),
	//	rangemap.CreateRange("2021-06-15", "2021-06-20"),
	//	rangemap.CreateRange("2021-07-15", "2021-09-20")}
	//
	//rm, err := rangemap.New(keys, []string{"a","b","c","d","e"})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//v, found := rm.Get(rangemap.CreateDate("2021-03-01"))
	//
	//if found {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("not found")
	//}

	//t := rangemap.ValueInTime{
	//	From: rangemap.CreateDate("2021-01-01"),
	//	To: rangemap.CreateDate("2021-01-31"),
	//	Value: "some value"}

	x := []rangemap.ValueInTime{
		{
			From:  rangemap.CreateDate("2021-06-15"),
			To:    rangemap.CreateDate("2021-06-20"),
			Value: "Ala ma kota",
		},
		{
			From:  rangemap.CreateDate("2021-01-01"),
			To:    rangemap.CreateDate("2021-01-31"),
			Value: "kot ma Ale",
		},
	}

	fmt.Println(x)

	sort.Slice(x, func(i, j int) bool {
		return x[i].From.Before(x[j].From)
	})

	fmt.Println(x)

}
