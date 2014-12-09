package trie

type (
	Trie     struct{}
	Iterator struct{}
)

func New() *Trie {
	return new(Trie)
}

func (t *Trie) Get(key []byte) interface{} {
	return nil
}

func (t *Trie) Set(key []byte, value interface{}) bool {
	return false
}

func (t *Trie) Delete(key []byte) bool {
	return false
}

func (t *Trie) Len() int {
	return 0
}

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
