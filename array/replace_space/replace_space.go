package replace_space

func replaceSpace(input string) string {
	length := len(input)
	for i := 0; i < length-1; i++ {
		if input[i] == ' ' {
			input += "  "
		}
	}

	i := length - 1
	j := len(input) - 1
	runes := []rune(input)
	for i >= 0 && i < j {
		if runes[i] != ' ' {
			runes[j] = runes[i]
			j--
		} else {
			runes[j] = '0'
			j--
			runes[j] = '2'
			j--
			runes[j] = '%'
			j--
		}
		i--
	}
	return string(runes)
}
