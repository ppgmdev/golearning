package iomanager

type IOManager interface { // go will check for methods with the signatures
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
