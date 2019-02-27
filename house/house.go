package house

import "strings"

var subject = []string{
	"the house",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}

var action = []string{
	"Jack built.",
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

// Verse generate the verse no v
func Verse(v int) string {
	var sb strings.Builder
	sb.WriteString("This is ")
	add := "\n"
	space := " "
	for i := v - 1; i >= 0; i-- {
		if i == 0 {
			add = " "
			space = ""
		}
		sb.WriteString(subject[i])
		sb.WriteString(add)
		sb.WriteString("that ")
		sb.WriteString(action[i])
		sb.WriteString(space)
	}
	return sb.String()
}

// Song Returns the song
func Song() string {
	var sb strings.Builder
	sb.WriteString(Verse(1))
	for i := 2; i <= 12; i++ {
		sb.WriteString("\n\n")
		sb.WriteString(Verse(i))
	}
	return sb.String()
}
