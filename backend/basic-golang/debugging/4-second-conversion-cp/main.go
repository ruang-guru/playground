package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Convert the given second to 00:00:00 hour minute second format

		Example Input/Output
		30 -> 00:00:30
		70 -> 00:01:10
		67812 -> 18:50:12
		678120 -> 188:22:00

	*/
	res := ConvertSecondToTimeString(30)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ConvertSecondToTimeStringCorrect(arr)
	// fmt.Println(resCorrect)
}

func ConvertSecondToTimeString(second int) string {
	hours := second / 3600
	minutes := second / 60

	timeString := fmt.Sprintf("%d:%d:%d", hours, minutes, second)
	return timeString
}

func ConvertSecondToTimeStringCorrect(second int) string {
	if second < 0 {
		return ""
	}
	var hours, minutes int
	if second >= 3600 {
		hours = second / 3600
		second -= hours * 3600
	}

	if second >= 60 {
		minutes = second / 60
		second -= minutes * 60
	}

	timeString := fmt.Sprintf("%v:%v:%v", addZero(hours), addZero(minutes), addZero(second))
	return timeString // TODO: replace this
}

func addZero(timeInt int) string {
	if timeInt < 10 {
		return "0" + strconv.Itoa(timeInt)
	}
	return strconv.Itoa(timeInt)
}
