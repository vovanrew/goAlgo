package structures

type HashTableInt struct {
	array []List
	Size  int
}

func (h *HashTableInt) hash(value int) int {
	return h.Size % value;
}

func (h *HashTableInt) Find(key int) (int, error) {
	hashingValue := h.hash(key)
	return h.array[hashingValue].FindNodeWithKey(key)
}

func (h *HashTableInt) Add(key int, value int) (error) {
	hashedValue := h.hash(key)
	node := &Node{key, value, nil, nil}
	h.array[hashedValue].AddNode(node)
}

func (h *HashTableInt) Remove(key int) {
	hashingValue := h.hash(key)
	h.array[hashingValue].RemoveNodeWithKey(key)
}