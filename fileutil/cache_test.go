package file

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var (
	mockErr error  = errors.New("mock error")
	fname   string = "tmp.json"
)

func TestCache(t *testing.T) {
	// Ensure that file does not exist.
	os.Remove(fname)
	if _, err := os.Stat(fname); err == nil {
		t.Fatalf("could not remove file")
	} else if !os.IsNotExist(err) {
		t.Fatalf("could not check presence of file: %v", err)
	}

	var (
		err error
		y   int
	)
	err = Cache(&y, fname, func() (int, error) { return 42, mockErr })
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != mockErr {
		t.Fatalf("wrong error: %v", err)
	}

	err = Cache(&y, fname, func() (int, error) { return 42, nil })
	if err != nil {
		t.Fatal(err)
	}
	if y != 42 {
		t.Fatal("wrong value")
	}

	var z int
	err = Cache(&z, fname, func() (int, error) { return 21, errors.New("ignored") })
	// Error should not be noticed since function is not called.
	if err != nil {
		t.Fatal(err)
	}
	if z != 42 {
		t.Fatal("cache was not used")
	}
}

func ExampleCache() {
	var (
		x int = 5
		y int
	)
	err := Cache(&y, "square.json", func() int {
		return x * x
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Output:
}
