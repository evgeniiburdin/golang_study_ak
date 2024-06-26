package hashmap

import (
	"hash/crc32"
	"hash/crc64"
	"time"
)

// HashMaper - интерфейс для хэш-карты
type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

// HashFunction - тип для хэш-функции
type HashFunction func(string) uint64

// HashMap - структура хэш-карты
type HashMap struct {
	data     map[uint64]interface{}
	hashFunc HashFunction
}

// NewHashMap - конструктор хэш-карты с опциями
func NewHashMap(opts ...func(*HashMap)) *HashMap {
	hm := &HashMap{
		data:     make(map[uint64]interface{}),
		hashFunc: defaultHashFunc,
	}
	for _, opt := range opts {
		opt(hm)
	}
	return hm
}

// Set - установка значения по ключу
func (hm *HashMap) Set(key string, value interface{}) {
	hash := hm.hashFunc(key)
	hm.data[hash] = value
}

// Get - получение значения по ключу
func (hm *HashMap) Get(key string) (interface{}, bool) {
	hash := hm.hashFunc(key)
	value, exists := hm.data[hash]
	return value, exists
}

// WithHashCRC64 - опция для использования хэш-функции CRC64
func WithHashCRC64() func(*HashMap) {
	return func(hm *HashMap) {
		hm.hashFunc = hashCRC64
	}
}

// WithHashCRC32 - опция для использования хэш-функции CRC32
func WithHashCRC32() func(*HashMap) {
	return func(hm *HashMap) {
		hm.hashFunc = hashCRC32
	}
}

// defaultHashFunc - стандартная хэш-функция
func defaultHashFunc(s string) uint64 {
	return hashCRC64(s)
}

// hashCRC64 - хэш-функция CRC64
func hashCRC64(s string) uint64 {
	return uint64(crc64.Checksum([]byte(s), crc64.MakeTable(crc64.ISO)))
}

// hashCRC32 - хэш-функция CRC32
func hashCRC32(s string) uint64 {
	return uint64(crc32.ChecksumIEEE([]byte(s)))
}

// MeassureTime - измерение времени выполнения функции
func MeassureTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
