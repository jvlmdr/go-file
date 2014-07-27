package file

import "io"

type Encoder interface {
	Encode(v interface{}) error
}

type Decoder interface {
	Decode(v interface{}) error
}

type (
	NewEncoder func(w io.Writer) Encoder
	NewDecoder func(r io.Reader) Decoder
)
