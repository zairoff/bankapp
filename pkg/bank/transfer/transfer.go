package transfer

func Total(summa int) int {
	if summa == 0 {
		return 0
	}
	summa += (summa * 5) / 1000
	return summa
}
