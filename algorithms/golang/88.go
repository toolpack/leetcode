package golang

func merge(a []int, m int, b []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if a[i] > b[j] {
			a[k] = a[i]
			k--
			i--
		} else {
			a[k] = b[j]
			k--
			j--
		}
	}
	for j >= 0 {
		a[k] = b[j]
		k--
		j--
	}
	return
}
