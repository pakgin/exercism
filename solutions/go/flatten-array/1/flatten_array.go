package flattenarray

func Flatten(nested any) []any {
	output := []any{}
	switch v := nested.(type) {
	case nil:
		return output
	case []any:
		for _, inner := range v {
			output = append(output, Flatten(inner)...)
		}
	default:
		output = append(output, v)
	}
	return output
}
