package iteration

func Iteration(s string, times int) string {
	result := ""
	for range times {
		result += s
	}
	return result
}
