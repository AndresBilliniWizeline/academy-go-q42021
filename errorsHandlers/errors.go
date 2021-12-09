package errorsHandlers

// CheckNilErr - handle all kind of errors and stops everything with a panic
func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckFileErr - handle error of file exists
func CheckFileErr(err error) {
	if err != nil {
		panic("The file doesn't exist")
	}
}
