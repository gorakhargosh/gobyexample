package trie

type (
	Trie struct {
		root   node
		length int
	}
	Iterator struct{}
	node     struct {
		value    interface{}
		children [256]*node
		length   byte
	}
)

func New() *Trie {
	return new(Trie)
}

func (t *Trie) Get(key []byte) interface{} {
	n := &t.root
	for _, b := range key {
		n = n.children[b]
		if n == nil {
			return nil
		}
	}
	return n.value
}

func (t *Trie) Set(key []byte, value interface{}) bool {
	n := &t.root
	// get to the leaf node (creating any branches along the way)
	for _, b := range key {
		c := n.children[b]
		if c == nil {
			c = new(node)
			n.children[b] = c
			n.length++
		}
		n = c
	}
	added := n.value == nil
	n.value = value
	if added {
		t.length++
	}
	return added
}

func (t *Trie) Delete(key []byte) bool {
	path := make([]*node, len(key))
	n := &t.root
	// build a path of nodes
	for i, b := range key {
		path[i] = n
		n = n.children[b]
		// if the node doesn't exist then the key must not be in the trie.
		if n == nil {
			return false
		}
	}

	n.value = nil
	// if this is a leaf node, we need to remove it from its parent (and so on up
	// the line for unvalued nodes that become leaves)
	if n.length == 0 {
		for i := len(key) - 1; i >= 0; i-- {
			path[i].children[key[i]] = nil
			path[i].length--
			if path[i].value == nil && path[i].length > 0 {
				break
			}
		}
	}
	t.length--
	return true
}

// Len returns the number of key value pairs in the trie.
func (t *Trie) Len() int {
	return t.length - 1
}

// Iterator returns an iterator that can be used to traverse the trie in sorted
// order. The iterator starts before the first key value pair, so Next() should
// be called before Key() and Value().
func (t *Trie) Iterator() *Iterator {
	return &Iterator{}
}

func (i *Iterator) Next() bool {
	return false
}

func (i *Iterator) Key() []byte {
	return nil
}

func (i *Iterator) Value() interface{} {
	return nil
}
