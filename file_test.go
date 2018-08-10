package cache

import (
	"os"
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {
	_, err := NewFile(".test", time.Second*1)
	defer os.RemoveAll(".test")

	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)
}

func TestFile_Set(t *testing.T) {
	r, err := NewFile(".test", time.Second*1)
	defer os.RemoveAll(".test")

	err = r.Set("test", []byte("test data"), time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)

	err = r.Set("test", []byte("test data"), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Set("test", []byte("test data"), time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)
}

func TestFile_Get(t *testing.T) {
	r, err := NewFile(".test", time.Second*1)
	defer os.RemoveAll(".test")

	testString := "test data"
	err = r.Set("test", []byte(testString), time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)

	d, err := r.Get("test")
	if err != nil {
		t.Fatal(err)
	}

	if string(d) != testString {
		t.Fatalf("Returned data not what we saved: %s != %s", string(d), testString)
	}
}

func TestFile_Load(t *testing.T) {
	r, err := NewFile(".test", time.Second*1)
	defer os.RemoveAll(".test")
	if err != nil {
		t.Fatal(err)
	}

	err = r.Set("test", []byte("test data"), time.Hour*1)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 1500)
	r.Stop()

	n, err := NewFile(".test", time.Second*1)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)

	err = n.Set("test2", []byte("test data"), time.Hour*1)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)
}
