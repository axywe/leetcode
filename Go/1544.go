package main

func makeGood(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if i+1 != len(b) && (b[i]+32 == b[i+1] || b[i]-32 == b[i+1]) {
			b = append(b[:i], b[i+2:]...)
			i = -1
		}
	}
	return string(b)
}
