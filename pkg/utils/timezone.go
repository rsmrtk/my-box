package utils

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

func ChangeTZ(timestamp int64, from_tz string, to_tz string) int64 {
	// convert timestamp from_tz to to_tz
	from_loc, err := time.LoadLocation(from_tz)
	if err != nil {
		fmt.Println(err)
	}
	to_loc, err := time.LoadLocation(to_tz)
	if err != nil {
		fmt.Println(err)
	}

	if from_tz == "UTC" {
		timeInToLocation := time.Now().In(to_loc)
		_, offset := timeInToLocation.Zone() // seconds

		return timestamp + int64(offset*1000) // ms
	} else {
		timeInToLocation := time.Now().In(from_loc)
		_, offset := timeInToLocation.Zone() // seconds

		return timestamp - int64(offset*1000) // ms
	}
}
