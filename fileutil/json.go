package file

import (
	"encoding/json"
	"io"
)

func newJSONEncoder(w io.Writer) Encoder { return json.NewEncoder(w) }
func newJSONDecoder(r io.Reader) Decoder { return json.NewDecoder(r) }

func SaveJSON(fname string, x interface{}) error {
	return Save(fname, newJSONEncoder, x)
}

func LoadJSON(fname string, x interface{}) error {
	return Load(fname, newJSONDecoder, x)
}
