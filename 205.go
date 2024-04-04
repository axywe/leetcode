package main

func isIsomorphic(s string, t string) bool {
	smap := make(map[byte]byte)
	tmap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if (smap[s[i]] != 0 && smap[s[i]] != t[i]) || (tmap[t[i]] != 0 && tmap[t[i]] != s[i]) {
			return false
		}
		smap[s[i]] = t[i]
		tmap[t[i]] = s[i]
	}
	return true
}
