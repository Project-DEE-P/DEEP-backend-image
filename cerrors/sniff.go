package cerrors

func Sniff(err error) {
	if err != nil {
		panic(err)
	}
}
