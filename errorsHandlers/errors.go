package errorsHandlers

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckFileErr(err error) {
	if err != nil {
		panic("The file doesn't exist")
	}
}
