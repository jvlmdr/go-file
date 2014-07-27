package fileutil

import (
	"errors"
	"path"
	"strings"
)

// SaveExt encodes to a file, determining encoder by file extension.
func SaveExt(fname string, x interface{}) error {
	switch ext := strings.ToLower(path.Ext(fname)); ext {
	case ".gob":
		return SaveGob(fname, x)
	case ".json":
		return SaveJSON(fname, x)
	case "":
		return errors.New("save: no extension: " + fname)
	default:
		return errors.New("save: unknown extension: " + ext)
	}
}

// LoadExt decodes from a file, determining decoder by file extension.
func LoadExt(fname string, x interface{}) error {
	switch ext := strings.ToLower(path.Ext(fname)); ext {
	case ".gob":
		return LoadGob(fname, x)
	case ".json":
		return LoadJSON(fname, x)
	case "":
		return errors.New("load: no extension: " + fname)
	default:
		return errors.New("load: unknown extension: " + ext)
	}
}
