package sublist

import "reflect"

type Relation string

const (
	equal     Relation = "equal"
	unequal   Relation = "unequal"
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
)

func Sublist(a, b []int) Relation {
	if len(a) == 0 && len(b) == 0 {
		return equal
	} else if len(a) == 0 {
		return sublist
	} else if len(b) == 0 {
		return superlist
	}

	if reflect.DeepEqual(a, b) {
		return equal
	}
	if len(a) > len(b) && checkSub(b, a) {
		return superlist
	}
	if len(a) < len(b) && checkSub(a, b) {
		return sublist
	}
	return unequal
}

func checkSub(short, long []int) bool {
	for i, v := range long {
		if v == short[0] {
			if reflect.DeepEqual(short, long[i:(i+len(short))]) {
				return true
			}
		}
	}
	return false
}
