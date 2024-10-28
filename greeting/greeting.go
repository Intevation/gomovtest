package greeting

func SignOfLife(n int) string {
	var beep string = "Be"
	for i := 1; i <= n; i++ {
		beep = beep + "e"
	}
	beep = beep + "p"

	return beep
}
