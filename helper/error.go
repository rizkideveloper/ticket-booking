package helper

//P nya besar biar bisa di eksport
func PanicIfError(err error)  {
	if err != nil {
		panic(err)
	}
}