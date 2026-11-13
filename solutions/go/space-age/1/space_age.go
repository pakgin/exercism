package space

type Planet string

const (
	Earth   Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

func Age(seconds float64, planet Planet) float64 {
	year := seconds / (60 * 60 * 24 * 365.256)
	switch planet {
	case Earth:
		return year
	case Mercury:
		return year / 0.2408467
	case Venus:
		return year / 0.61519726
	case Mars:
		return year / 1.8808158
	case Jupiter:
		return year / 11.862615
	case Saturn:
		return year / 29.447498
	case Uranus:
		return year / 84.016846
	case Neptune:
		return year / 164.79132
	default:
		return -1.0
	}
}
