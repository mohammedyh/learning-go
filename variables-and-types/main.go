package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int16
	var conditionGrade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	conditionGrade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "released in", year, "number of pages", pageNumber, "condition grade", conditionGrade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	conditionGrade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "released in", year, "number of pages", pageNumber, "condition grade", conditionGrade)
}
