func isValid(s string) bool {
	var S = make([]byte, 0, len(s))

	for _, char := range s {
		switch char {
		case '{', '[', '(':
			S = append(S, byte(char))
		case '}':
			if len(S) != 0 && S[len(S)-1] == '{' {
				S = S[:len(S)-1]
			} else {
				return false
			}

		case ']':
			if len(S) != 0 && S[len(S)-1] == '[' {
				S = S[:len(S)-1]
			} else {
				return false
			}
		case ')':
			if len(S) != 0 && S[len(S)-1] == '(' {
				S = S[:len(S)-1]
			} else {
				return false
			}

		default:
			return false
		}

	}

	fmt.Println(len(S))

	return len(S) == 0
}