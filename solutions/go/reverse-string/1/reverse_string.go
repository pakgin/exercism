package reverse

func Reverse(input string) string {
	output := []rune(input)
	start, end := 0, len(output)-1

	for start < end {
		output[start], output[end] = output[end], output[start]

		start++
		end--
	}

	return string(output)
}
