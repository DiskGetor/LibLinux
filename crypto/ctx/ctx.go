package ctx

type (
	IoInfo struct {
		In      []byte
		Key     []byte
		Out     []byte
		Err     error
		ErrInfo string
	}
)

const (
	Encode = iota
	Decode
)
