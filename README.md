# Date Range map inspired by Guava RangeMap

Very early stage of project

It is useful in work with values changing in time - like e.g. person salary

from 2020-01-01 he or she was earning 100.00 but after 2020-09-01 he or she will be earning 200.00

| range                   | value |
|-------------------------|-------|
| 2020-01-01 - 2020-09-01 | 100   |
| 2020-10-01 - 2999-13-31 | 200   |

````go
package main

import (
	"fmt"
	"github.com/alapierre/date-range-map/rangemap"
)

func main()  {

	values := []rangemap.ValueInTime{
		rangemap.NewValueInTime("2021-03-01", "2021-04-30", "b"),
		rangemap.NewValueInTime("2021-01-01", "2021-01-31", "a"),
		rangemap.NewValueInTime("2021-05-15", "2021-05-31", "c"),
		rangemap.NewValueInTime("2021-06-15", "2021-06-20", "d"),
		rangemap.NewValueInTime("2021-07-15", "2021-09-20", "e"),
	}

    // values will be sorted by begin date
	rm, err := rangemap.New(values)

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
````