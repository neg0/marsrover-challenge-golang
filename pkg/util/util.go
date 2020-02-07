package util

func TernaryInt(cond bool, value int, finally int) int {
	if cond {
		return value
	}
	return finally
}
