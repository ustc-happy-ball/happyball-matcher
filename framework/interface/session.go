package _interface

type Session interface {
	Read() ([]byte, error)
	Write(buff []byte) error
}
