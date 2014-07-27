package file

import (
	"log"
	"os"
	"reflect"
)

// Cache caches the result of a function call in the specified file.
// If the file already exists, then it is loaded and the function is not called.
//
// The function must take no arguments and have one or two return values.
// The destination must be a pointer.
//
// If the file exists but the result cannot be loaded, an error is returned.
//
// The method of serialization is determined by file extension.
// Refer to github.com/jvlmdr/go-enc/enc.SaveExt for more details.
func Cache(dst interface{}, fname string, f interface{}) error {
	// Check if the file exists.
	if _, err := os.Stat(fname); err == nil {
		// File exists. Load from cache or fail.
		log.Println("load from cache:", fname)
		return LoadExt(fname, dst)
	} else if !os.IsNotExist(err) {
		// Stat failed for some reason other than file not existing.
		return err
	}

	y, err := call(reflect.ValueOf(f))
	if err != nil {
		return err
	}
	// Save first output argument.
	if err := SaveExt(fname, y.Interface()); err != nil {
		return err
	}
	// Assign first output argument to dst.
	reflect.ValueOf(dst).Elem().Set(y)
	return nil
}

func call(f reflect.Value) (reflect.Value, error) {
	// Stat failed because file does not exist.
	// Call function with no arguments and save result.
	out := f.Call(nil)
	if len(out) > 1 {
		// Assert second output argument is an error and examine it.
		var err error
		reflect.ValueOf(&err).Elem().Set(out[1])
		if err != nil {
			// Function returned an error.
			return reflect.Value{}, err
		}
	}
	return out[0], nil
}
