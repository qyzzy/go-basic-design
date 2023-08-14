package basic_map

import (
	"reflect"
	"testing"
)

func TestHashTable_HashFunction(t *testing.T) {
	type fields struct {
		data [TableSize]*KeyValuePair
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := &HashTable{
				data: tt.fields.data,
			}
			if got := ht.HashFunction(tt.args.key); got != tt.want {
				t.Errorf("HashFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Insert(t *testing.T) {
	type fields struct {
		data [TableSize]*KeyValuePair
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := &HashTable{
				data: tt.fields.data,
			}
			ht.Insert(tt.args.key, tt.args.value)
		})
	}
}

func TestHashTable_Search(t *testing.T) {
	type fields struct {
		data [TableSize]*KeyValuePair
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := &HashTable{
				data: tt.fields.data,
			}
			got, got1 := ht.Search(tt.args.key)
			if got != tt.want {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewHashTable(t *testing.T) {
	tests := []struct {
		name string
		want *HashTable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
