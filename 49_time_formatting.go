package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	t := time.Now()
	// We can format time simply using build in formats
	p(t.Format(time.RFC3339))
	// Time parsing takes a time and converts it to the given format
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	// Format and Parse use example based layouts. Layouts MUST use the reference time
	// "Mon Jan 2 15:04:05 MST 2006" to show the pattern with which to parse the string
	p(t.Format("3:04PM"))
	p(t.Format("Mon jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
	// Typical, boring way
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// Using Parse on a malformed inputs returns an error
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}