package main

func Parse(word string) string {
	if len(word) == 0 {
		return ""
	}

	if len(word) == 1 {
		return word
	}

	var i int

	result := word

	for {
		switch result[i : i+2] {
		case "AB", "BA", "CD", "DC":
			if len(result) == 2 {
				return ""
			}

			result = result[:i] + result[i+2:]

			i = 0
		}

		if len(result) == 2 {
			switch result {
			case "AB", "BA", "CD", "DC":
				return ""
			default:
				return result
			}
		}

		if len(result) < 2 || i >= len(result)-2 {
			break
		}

		i++
	}

	return result
}
