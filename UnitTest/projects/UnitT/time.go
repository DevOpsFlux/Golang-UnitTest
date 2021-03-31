package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		now := time.Now()

		var secs int64
		secs = now.Unix()
		fmt.Println("secs : ", secs)
		t := time.Unix(secs, 0)
		fmt.Println(t)
		secsStr := strconv.FormatInt(secs, 10)
		fmt.Println("secsStr : ", secsStr)

		secsI, _ := strconv.Atoi(secsStr)
		//secsI, _ := strconv.ParseInt(secsStr, 10, 64)
		fmt.Println("secsI : ", secsI)

		then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		fmt.Println("then : ", then)

		//diff := now.Sub(then)
		//diff := now.Sub(then)
		//fmt.Println("diff : ", diff)

		t1 := time.Now()
		t2 := t1.Add(time.Second * 341)

		fmt.Println(t1)
		fmt.Println(t2)

		diff := t2.Sub(t1)
		fmt.Println(diff)

		//In hours
		fmt.Printf("Hours: %f\n", diff.Hours())

		//In minutes
		fmt.Printf("Minutes: %f\n", diff.Minutes())

		//In seconds
		fmt.Printf("Seconds: %f\n", diff.Seconds())

		//In nanoseconds
		fmt.Printf("Nanoseconds: %d\n", diff.Nanoseconds())

		chkMinVal := int(diff.Minutes())
		chkSecVal := int(diff.Seconds())

		fmt.Println("chkMinVal : ", chkMinVal)
		fmt.Println("chkSecVal : ", chkSecVal)
	*/
	dateformat := time.Now()
	fmt.Println("dateformat : ", dateformat)

	dateformatUTC := time.Now().UTC()
	fmt.Println("dateformatUTC : ", dateformatUTC)

	/*
			currentTime := time.Now()
		    oldTime := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
		    diff := currentTime.Sub(oldTime)

		    //In hours
		    fmt.Printf("Hours: %f\n", diff.Hours())

		    //In minutes
		    fmt.Printf("Minutes: %f\n", diff.Minutes())

		    //In seconds
		    fmt.Printf("Seconds: %f\n", diff.Seconds())

		    //In nanoseconds
		    fmt.Printf("Nanoseconds: %d\n", diff.Nanoseconds())
	*/

}
