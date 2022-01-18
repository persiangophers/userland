package protocol

type UUID interface {
	Create() ([16]byte, error)
	Delete(uuid []byte) error
}
