package cache

import (
	"os"
	"testing"
	"time"
)

func TestNewMemcached(t *testing.T) {
	if os.Getenv("MEMCACHED") == "" {
		t.Skip("No MEMCACHED environment variable set")
	}

	NewMemcached(os.Getenv("MEMCACHED"))
}

func TestMemcached_Set(t *testing.T) {
	if os.Getenv("MEMCACHED") == "" {
		t.Skip("No MEMCACHED environment variable set")
	}

	r := NewMemcached(os.Getenv("MEMCACHED"))

	err := r.Set("test", []byte("test data"), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMemcached_Get(t *testing.T) {
	if os.Getenv("MEMCACHED") == "" {
		t.Skip("No MEMCACHED environment variable set")
	}

	r := NewMemcached(os.Getenv("MEMCACHED"))

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
