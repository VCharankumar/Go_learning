package iomanager

type IOManager interface {
	ReadData() ([]string, error)
	WriteResult(data interface{}) error
}
