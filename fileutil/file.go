package fileutil

import "os"

func Save(fname string, enc NewEncoder, x interface{}) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()
	return enc(file).Encode(x)
}

func Load(fname string, dec NewDecoder, x interface{}) error {
	file, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()
	return dec(file).Decode(x)
}
