package easy

var opp = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}'}

func isValid(s string) bool {
	open := make([]byte, 0)
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
