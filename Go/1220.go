package main

func countVowelPermutation(n int) int {
	const MOD = 1e9 + 7
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 1; j < n; j++ {
		a_next := e
		e_next := (a + i) % MOD
		i_next := (a + e + o + u) % MOD
		o_next := (i + u) % MOD
		u_next := a

		a, e, i, o, u = a_next, e_next, i_next, o_next, u_next
	}

	return (a + e + i + o + u) % MOD
}
