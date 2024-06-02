package main

import "testing"

const parallelFactor = 10_000

func Benchmark_NoMutex(b *testing.B) {
	b.Skip("panic in nomutex")
	//c := NewCache()
	//for i := 0; i < b.N; i++ {
	//	//emulateLoad(c, parallelFactor)
	//}
}

func Benchmark_Mutex_Balanced(b *testing.B) {
	c := NewCache()
	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

func Benchmark_RWMutex_Balanced(b *testing.B) {
	c := NewCache()
	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

func Benchmark_Mutex_ReadIntensiveLoad(b *testing.B) {
	c := NewCache()
	for i := 0; i < b.N; i++ {
		emulateReadIntensiveLoad(c, parallelFactor)
	}
}
