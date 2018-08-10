package cache

import (
	"testing"
	"time"
)

func TestNewInmem(t *testing.T) {
	NewInmem()
}

func TestInmem_Set(t *testing.T) {
	r := NewInmem()

	err := r.Set("test", []byte("test data"), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInmem_Get(t *testing.T) {
	r := NewInmem()

	testString := "test data"
	err := r.Set("test", []byte(testString), time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second)

	d, err := r.Get("test")
	if err != nil {
		t.Fatal(err)
	}

	if string(d) != testString {
		t.Fatalf("Returned data not what we saved: %s != %s", string(d), testString)
	}
}
