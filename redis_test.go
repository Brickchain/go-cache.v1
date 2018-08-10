package cache

import (
	"os"
	"testing"
	"time"
)

func TestNewRedisCache(t *testing.T) {
	if os.Getenv("REDIS") == "" {
		t.Skip("No REDIS environment variable set")
	}
	NewRedisCache(os.Getenv("REDIS"))
}

func TestRedisCache_Set(t *testing.T) {
	if os.Getenv("REDIS") == "" {
		t.Skip("No REDIS environment variable set")
	}

	r := NewRedisCache(os.Getenv("REDIS"))

	err := r.Set("test", []byte("test data"), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisCache_Get(t *testing.T) {
	if os.Getenv("REDIS") == "" {
		t.Skip("No REDIS environment variable set")
	}

	r := NewRedisCache(os.Getenv("REDIS"))

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
