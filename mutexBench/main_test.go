package mutexBench

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
)

type Set struct {
	sync.Mutex
	m map[float64]struct{}
}

type SetRW struct {
	sync.RWMutex
	m map[float64]struct{}
}


func BenchmarkSet10Read90Write(b *testing.B) {
	s := &Set {
		m: map[float64]struct{}{},
	}
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 1; i <= 9; i += 1 {
					func() {
						s.Lock()
						defer s.Unlock()
						s.m[3.5] = struct{}{}
					}()
				}
				for i := 1; i <= 1; i += 1 {
					func() bool {
						s.Lock()
						defer s.Unlock()
						_, ok := s.m[3.5]
						return ok
					}()
				}
			}
		})
	})
}

func BenchmarkSetRW10Read90Write(b *testing.B) {
	s := &SetRW {
		m: map[float64]struct{}{},
	}
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 1; i <= 9; i += 1 {
					func() {
						s.Lock()
						defer s.Unlock()
						s.m[3.5] = struct{}{}
					}()
				}
				for i := 1; i <= 1; i += 1 {
					func() bool {
						s.Lock()
						defer s.Unlock()
						_, ok := s.m[3.5]
						return ok
					}()
				}
			}
		})
	})
}

func BenchmarkSet50Read50Write(b *testing.B) {
	s := &Set {
		m: map[float64]struct{}{},
	}
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 1; i <= 5; i += 1 {
					func() {
						s.Lock()
						defer s.Unlock()
						s.m[3.5] = struct{}{}
					}()
				}
				for i := 1; i <= 5; i += 1 {
					func() bool {
						s.Lock()
						defer s.Unlock()
						_, ok := s.m[3.5]
						return ok
					}()
				}
			}
		})
	})
}


func BenchmarkSetRW50Read50Write(b *testing.B) {
	s := &SetRW {
		m: map[float64]struct{}{},
	}
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 1; i <= 5; i += 1 {
					func() {
						s.Lock()
						defer s.Unlock()
						s.m[3.5] = struct{}{}
					}()
				}
				for i := 1; i <= 5; i += 1 {
					func() bool {
						s.Lock()
						defer s.Unlock()
						_, ok := s.m[3.5]
						return ok
					}()
				}
			}
		})
	})
}

func BenchmarkSet90Read10Write(b *testing.B) {
	s := &Set {
		m: map[float64]struct{}{},
	}
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 1; i <= 1; i += 1 {
					func() {
						s.Lock()
						defer s.Unlock()
						s.m[3.5] = struct{}{}
					}()
				}
				for i := 1; i <= 9; i += 1 {
					func() bool {
						s.Lock()
						defer s.Unlock()
						_, ok := s.m[3.5]
						return ok
					}()
				}
			}
		})
	})
}

func BenchmarkSetRW90Read10Write(b *testing.B) {
	s := &SetRW {
		m: map[float64]struct{}{},
	}
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 1; i <= 1; i += 1 {
					func() {
						s.Lock()
						defer s.Unlock()
						s.m[3.5] = struct{}{}
					}()
				}
				for i := 1; i <= 9; i += 1 {
					func() bool {
						s.Lock()
						defer s.Unlock()
						_, ok := s.m[3.5]
						return ok
					}()
				}
			}
		})
	})
}