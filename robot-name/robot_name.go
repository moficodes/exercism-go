package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const maxNamesAvailable = 26 * 26 * 1000

var namesInUse = make(map[string]bool)

// Robot with a name
type Robot struct {
	name string
}

// Name returns the name of the Robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		e := r.Reset()
		if e != nil {
			return "", e
		}
	}
	return r.name, nil
}

// Reset the robot to factory settings
func (r *Robot) Reset() error {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(namesInUse) >= maxNamesAvailable {
		return fmt.Errorf("Could not get new name for robot, all available names was in use")
	}

	var name string
	for namesInUse[name] || name == "" {
		a := string(rand.Intn(26) + 'A')
		b := string(rand.Intn(26) + 'A')
		num := rand.Intn(1000)

		name = fmt.Sprintf("%s%s%03d", a, b, num)
	}
	namesInUse[name] = true
	r.name = name
	return nil
}
