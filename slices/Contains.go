package main

func ContainsInt(xi []int, in int) (bool, int) {
	for i, v := range xi {
		if v == in {
			return true, i
		}
	}
	return false, -1
}
