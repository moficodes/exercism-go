package twelve

import "fmt"

var days = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Verse return one line for the given input
func Verse(i int) string {
	chorus := "On the %s day of Christmas my true love gave to me: "
	out := fmt.Sprintf(chorus, days[i])
	extra := ""
	// in the first verse we do not need 'and'
	if i != 1 {
		extra = "and "
	}
	for j := i - 1; j > 0; j-- {
		out += gifts[j] + ", "
	}
	out += extra + gifts[0] + "."
	return out
}

// Song returns the entire song
func Song() string {
	out := ""
	for i := 1; i <= 12; i++ {
		out += Verse(i) + "\n"
	}
	return out
}
