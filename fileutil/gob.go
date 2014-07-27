package file

import (
	"encoding/gob"
	"io"
)

func newGobEncoder(w io.Writer) Encoder { return gob.NewEncoder(w) }
func newGobDecoder(r io.Reader) Decoder { return gob.NewDecoder(r) }

func SaveGob(fname string, x interface{}) error {
	return Save(fname, newGobEncoder, x)
}

func LoadGob(fname string, x interface{}) error {
	return Load(fname, newGobDecoder, x)
}
