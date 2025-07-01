package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://youtube.com",
			val: []byte("Sweet cats video"),
		},
		{
			key: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			val: []byte("Click me!"),
		},
	}
	for i, cs := range cases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {
			cache := NewCache(5 * time.Second)
			cache.Add(cs.key, cs.val)
			val, ok := cache.Get(cs.key)
			if !ok {
				t.Errorf("expected to find a key")
				return
			}
			if string(val) != string(cs.val) {
				t.Errorf("expected to find a value")
				return
			}

		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestGetEmpty(t *testing.T) {
	cache := NewCache(5 * time.Second)
	keys := make([]string, 0, len(cache.cache))
	if len(keys) != 0 {
		t.Errorf("expected to not find any key")
		return
	}
}
