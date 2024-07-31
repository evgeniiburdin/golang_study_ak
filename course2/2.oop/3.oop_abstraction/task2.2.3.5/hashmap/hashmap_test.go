package hashmap

import (
	"testing"
)

func TestHashMap_SetGet(t *testing.T) {
	m := NewHashMap(WithHashCRC64())
	m.Set("key1", "value1")
	value, ok := m.Get("key1")
	if !ok || value != "value1" {
		t.Errorf("expected value1, got %v", value)
	}

	m = NewHashMap(WithHashCRC32())
	m.Set("key2", "value2")
	value, ok = m.Get("key2")
	if !ok || value != "value2" {
		t.Errorf("expected value2, got %v", value)
	}
}

func BenchmarkHashMap_SetGet_CRC64(b *testing.B) {
	m := NewHashMap(WithHashCRC64())
	for i := 0; i < b.N; i++ {
		m.Set("key", "value")
		m.Get("key")
	}
}

func BenchmarkHashMap_SetGet_CRC32(b *testing.B) {
	m := NewHashMap(WithHashCRC32())
	for i := 0; i < b.N; i++ {
		m.Set("key", "value")
		m.Get("key")
	}
}
