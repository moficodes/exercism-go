package space

// Planet type alias for string
type Planet string

const (
	earth   float64 = 31557600.0
	mercury float64 = 0.2408467 * earth
	venus   float64 = 0.61519726 * earth
	mars    float64 = 1.8808158 * earth
	jupiter float64 = 11.862615 * earth
	saturn  float64 = 29.447498 * earth
	uranus  float64 = 84.016846 * earth
	neptune float64 = 164.79132 * earth
)

// Age calculate the age of someone in earth years
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / earth
	case "Mercury":
		return seconds / mercury
	case "Venus":
		return seconds / venus
	case "Mars":
		return seconds / mars
	case "Jupiter":
		return seconds / jupiter
	case "Saturn":
		return seconds / saturn
	case "Uranus":
		return seconds / uranus
	case "Neptune":
		return seconds / neptune
	default:
		return 0.0
	}
}
