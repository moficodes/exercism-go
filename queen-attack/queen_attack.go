package queenattack

import "errors"

// CanQueenAttack checks if a queen can attack the other queen
func CanQueenAttack(w, b string) (bool, error) {
	if len(w) != 2 || len(b) != 2 {
		return false, errors.New("not a valid position")
	}
	wx, wy := w[0]-'a', w[1]-'1'
	bx, by := b[0]-'a', b[1]-'1'
	if wx < 0 || wx > 7 || wy < 0 || wy > 7 || bx < 0 || bx > 7 || by < 0 || by > 7 {
		return false, errors.New("outside of board")
	}

	if wx == bx && wy == by {
		// same square
		return false, errors.New("same square")
	}
	if wx == bx || wy == by {
		// same row
		return true, nil
	}
	if wx-wy == bx-by || wx+wy == bx+by {
		// diagonals
		return true, nil
	}
	return false, nil
}
