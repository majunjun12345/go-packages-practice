package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"sync"
	"testing"
)

// ---------------------------------------------------------------- json decoder
type TestingData struct {
	Data string `json:"data"`
	Key  string `json:"key"`
}

var responsePool = sync.Pool{
	New: func() interface{} {
		return new(TestingData)
	},
}

func getRespnse() *TestingData {
	return responsePool.Get().(*TestingData)
}

func PutRespnse(buf *TestingData) {
	responsePool.Put(buf)
}

// BenchmarkJsonDecodeWithPool-12    	 1000000	      1421 ns/op	    1152 B/op	      10 allocs/op   2.394
func BenchmarkJsonDecodeWithPool(b *testing.B) {
	b.N = 1000000
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data := bytes.NewReader([]byte("{\"data\":\"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie.\",\"key\":\"Lorem\"}"))
		response := getRespnse()
		json.NewDecoder(data).Decode(&response)
		PutRespnse(response)
	}
}

// BenchmarkJsonDecodeWithoutPool-12    	 1000000	      1439 ns/op	    1184 B/op	      11 allocs/op   1.829
func BenchmarkJsonDecodeWithoutPool(b *testing.B) {
	b.N = 1000000
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data := bytes.NewReader([]byte("{\"data\":\"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie.\",\"key\":\"Lorem\"}"))
		response := &TestingData{}
		json.NewDecoder(data).Decode(&response)
	}
}

// ---------------------------------------------------------------- buffer

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func PutBuffer(b *bytes.Buffer) {
	bufferPool.Put(b)
}

func BenchmarkReadStreamWithPool(b *testing.B) {
	data := TestingData{
		Data: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie.",
		Key:  "Lorem",
	}

	b.N = 3000000
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := GetBuffer()

		json.NewEncoder(buf).Encode(data)

		io.Copy(ioutil.Discard, buf)
		PutBuffer(buf)
	}
}

func BenchmarkReadStreamWithoutPool(b *testing.B) {
	data := TestingData{
		Data: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie.",
		Key:  "Lorem",
	}

	b.N = 3000000
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(data)
		io.Copy(ioutil.Discard, buf)
	}
}
