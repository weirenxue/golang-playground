package main

type Dummy struct {
	Value int
}

// initialize with full zero value array
// and assign new struct by index access
func SetSlice01(n int) {
	s := make([]Dummy, n)
	for i := 0; i < n; i++ {
		s[i] = Dummy{
			Value: 1,
		}
	}
}

// initialize with full zero value array
// and assign value to original struct in array
// by index access
func SetSlice02(n int) {
	s := make([]Dummy, n)
	for i := 0; i < n; i++ {
		s[i].Value = 1
	}
}

// initialize with empty array
func SetSlice03(n int) {
	s := make([]Dummy, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, Dummy{
			Value: 1,
		})
	}
}

// no initial len and cap
func SetSlice04(n int) {
	s := []Dummy{}
	for i := 0; i < n; i++ {
		s = append(s, Dummy{
			Value: 1,
		})
	}
}

// slice of pointer
func SetSlice05(n int) {
	s := make([]*Dummy, n)
	for i := 0; i < n; i++ {
		s[i] = &Dummy{
			Value: 1,
		}
	}
}
