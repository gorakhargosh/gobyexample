package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := New()

	cases := []struct{ key, value string }{
		{"car", "white"},
		{"cat", "black"},
	}

	for _, c := range cases {
		if !trie.Set([]byte(c.key), c.value) {
			t.Error("expected trie to not already contain", c.key)
		}
	}

	for _, c := range cases {
		if trie.Set([]byte(c.key), c.value) {
			t.Error("expected trie to already contain", c.key)
		}
	}

	if trie.Len() != 2 {
		t.Error("expected trie to contain 2 items.")
	}

	for _, c := range cases {
		v, ok := trie.Get([]byte(c.key)).(string)
		if !ok || v != c.value {
			t.Errorf("expected trie to have %v, got %v", c.value, v)
		}
	}

	for _, c := range cases {
		if !trie.Delete([]byte(c.key)) {
			t.Error("expected trie to delete", c.key)
		}
	}

	for _, c := range cases {
		if trie.Get([]byte(c.key)) != nil {
			t.Error("expected trie to no longer contain", c.key)
		}
	}

	for trie.Len() != 0 {
		t.Error("expected trie to contain 0 items")
	}

	for _, c := range cases {
		trie.Set([]byte(c.key), c.value)
	}

	iterator := trie.Iterator()
	for _, c := range cases {
		if !iterator.Next() {
			t.Error("expected another item")
		}
		if string(iterator.Key()) != c.key {
			t.Errorf("expected %v, got %v", c.key, iterator.Key())
		}
		if str, ok := iterator.Value().(string); !ok || str != c.value {
			t.Errorf("expected %v, got %v", c.value, iterator.Value())
		}
	}

	if iterator.Next() {
		t.Error("expected only", len(cases), "entries")
	}
}
