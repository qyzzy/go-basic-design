package basic_map

const TableSize = 10

type HashTable struct {
	data [TableSize]*KeyValuePair
}

type KeyValuePair struct {
	key   string
	value interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (ht *HashTable) HashFunction(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % TableSize
}

func (ht *HashTable) Insert(key, value string) {
	index := ht.HashFunction(key)

	for ht.data[index] != nil {
		index = (index + 1) % TableSize
	}

	ht.data[index] = &KeyValuePair{key, value}
}

func (ht *HashTable) Search(key string) (interface{}, bool) {
	index := ht.HashFunction(key)

	for ht.data[index] != nil {
		if ht.data[index].key == key {
			return ht.data[index].value, true
		}
		index = (index + 1) % TableSize
	}

	return "", false
}
