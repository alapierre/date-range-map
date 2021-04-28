package main

import (
	"fmt"
	"github.com/alapierre/date-range-map/rangemap"
)

func main() {

	values := []rangemap.ValueInTime{
		rangemap.NewValueInTime("2021-03-01", "2021-04-30", "b"),
		rangemap.NewValueInTime("2021-01-01", "2021-01-31", "a"),
		rangemap.NewValueInTime("2021-05-15", "2021-05-31", "c"),
		rangemap.NewValueInTime("2021-06-15", "2021-06-20", "d"),
		rangemap.NewValueInTime("2021-07-15", "2021-09-20", "e"),
	}

	rm, err := rangemap.New(values)

	fmt.Println(values)

	if err != nil {
		panic(err)
	}

	v, found := rm.Get(rangemap.CreateDate("2021-03-01"))

	if found {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}

}
