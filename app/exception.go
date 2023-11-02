package app

func Must(e any) {
	if e != nil {
		panic(e)
	}
}
