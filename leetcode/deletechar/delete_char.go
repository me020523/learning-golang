package deletechar

func DeleteChar(str []rune, c rune) []rune {
	var index int = 0
	for _, s := range str {
		if s == c {
			continue
		}
		str[index] = s
		index += 1
	}
	if index == 0 {
		return []rune{}
	} else {
		return str[0:index]
	}

}
