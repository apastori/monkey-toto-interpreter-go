package tokenizer

func isTwoCharCandidate(ch byte) bool {
	switch ch {
	case '=', '!', '<', '>', '|', '&':
		return true
	}
	return false
}
