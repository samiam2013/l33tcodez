package easy

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	opp := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}'}
	open := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		r := s[i]
		switch {
		case r == '(' || r == '[' || r == '{':
			open = append(open, r)
		case len(open) > 0:
			if r != opp[open[len(open)-1]] {
				return false
			}
			open = open[:len(open)-1]
		default:
			return false
		}
	}
	return len(open) == 0
}
