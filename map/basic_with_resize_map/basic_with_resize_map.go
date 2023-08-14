package basic_with_resize_map

const (
	InitialTableSize = 10
	LoadFactor       = 0.7
)

type HashTable struct {
	data     []*KeyValuePair
	size     int
	capacity int
}

type KeyValuePair struct {
	key   string
	value string
}

func NewHashTable() *HashTable {
	return &HashTable{
		data:     make([]*KeyValuePair, InitialTableSize),
		size:     0,
		capacity: InitialTableSize,
	}
}

func (ht *HashTable) HashFunction(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % ht.capacity
}

func (ht *HashTable) Insert(key, value string) {
	if float64(ht.size)/float64(ht.capacity) >= LoadFactor {
		ht.resize()
	}

	index := ht.HashFunction(key)

	for ht.data[index] != nil {
		if ht.data[index].key == key {
			ht.data[index].value = value
			return
		}
		index = (index + 1) % ht.capacity
	}

	ht.data[index] = &KeyValuePair{key, value}
	ht.size++
}

func (ht *HashTable) Search(key string) (string, bool) {
	index := ht.HashFunction(key)

	for ht.data[index] != nil {
		if ht.data[index].key == key {
			return ht.data[index].value, true
		}
		index = (index + 1) % ht.capacity
	}

	return "", false
}

func (ht *HashTable) resize() {
	ht.capacity *= 2
	newData := make([]*KeyValuePair, ht.capacity)
	for _, kvp := range ht.data {
		if kvp != nil {
			newIndex := ht.HashFunction(kvp.key)
			for newData[newIndex] != nil {
				newIndex = (newIndex + 1) % ht.capacity
			}
			newData[newIndex] = kvp
		}
	}
	ht.data = newData
}
