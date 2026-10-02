package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestGetMissing(t *testing.T) {
	cache := &Cache{cache: make(map[string]cacheEntry)}

	_, ok := cache.Get("https://example.com/missing")
	if ok {
		t.Errorf("expected key to be missing")
	}
}

func TestAddOverwritesExistingValue(t *testing.T) {
	cache := &Cache{cache: make(map[string]cacheEntry)}
	key := "https://example.com"
	cache.Add(key, []byte("old value"))
	cache.Add(key, []byte("new value"))

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected to find key")
	}
	if string(got) != "new value" {
		t.Errorf("got %q, want %q", got, "new value")
	}
}

func TestReapRemovesExpiredEntriesOnly(t *testing.T) {
	now := time.Now()
	interval := time.Minute
	cache := &Cache{cache: map[string]cacheEntry{
		"expired": {
			createdAt: now.Add(-2 * interval),
			val:       []byte("expired"),
		},
		"fresh": {
			createdAt: now.Add(-interval / 2),
			val:       []byte("fresh"),
		},
	}}

	cache.reap(now, interval)

	if _, ok := cache.Get("expired"); ok {
		t.Errorf("expected expired entry to be removed")
	}
	if got, ok := cache.Get("fresh"); !ok || string(got) != "fresh" {
		t.Errorf("expected fresh entry to remain, got %q, found %v", got, ok)
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
