package solution

func generateTheString(n int) string {
	res := make([]byte, n)

	for i := 0; i < n; i++ {
		res[i] = 'a'
	}

	if n%2 == 1 {
		return string(res)
	} else {
		res[n-1] = 'b'
		return string(res)
	}
}
