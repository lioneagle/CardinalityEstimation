package murmur

import (
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"testing"
)

func TestMurmur3(t *testing.T) {
	text1 := []byte("The quick brown fox jumps over the lazy dog")
	expectedHash1 := uint32(0x78e69e27)
	actualHash1 := Murmur3(text1)
	if expectedHash1 != actualHash1 {
		t.Errorf("Expected hash value: %d, got %d", expectedHash1, actualHash1)
	}

	text2 := []byte("The quick brown fox jumps over the lazy cog")
	expectedHash2 := uint32(0xd5ece287)
	actualHash2 := Murmur3(text2)
	if expectedHash2 != actualHash2 {
		t.Errorf("Expected hash value: %d, got %d", expectedHash2, actualHash2)
	}
}

var testDataSize int64 = 15

func BenchmarkMurmur3(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(testDataSize)
	data := make([]byte, testDataSize)
	for i := range data {
		data[i] = byte(i + 'a')
	}

	b.StartTimer()
	for todo := b.N; todo != 0; todo-- {
		Murmur3(data)
	}
}

func BenchmarkFnv(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(testDataSize)
	data := make([]byte, testDataSize)
	for i := range data {
		data[i] = byte(i + 'a')
	}

	h := fnv.New32()
	//in := make([]byte, 0, h.Size())

	b.StartTimer()
	for todo := b.N; todo != 0; todo-- {
		h.Reset()
		h.Write(data)
		h.Sum32()
	}
}

func BenchmarkCrc32(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(testDataSize)
	data := make([]byte, testDataSize)
	for i := range data {
		data[i] = byte(i + 'a')
	}

	//h := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	h := crc32.NewIEEE()
	//in := make([]byte, 0, h.Size())

	b.StartTimer()
	for todo := b.N; todo != 0; todo-- {
		h.Reset()
		h.Write(data)
		h.Sum32()
	}
}

func BenchmarkAdler32(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(testDataSize)
	data := make([]byte, testDataSize)
	for i := range data {
		data[i] = byte(i + 'a')
	}

	//h := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	h := adler32.New()
	//in := make([]byte, 0, h.Size())

	b.StartTimer()
	for todo := b.N; todo != 0; todo-- {
		h.Reset()
		h.Write(data)
		h.Sum32()
	}
}