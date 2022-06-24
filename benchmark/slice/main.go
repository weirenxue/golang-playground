package main

type Dummy struct {
	Value int
}

// no initial len and cap
func SetSlice01(n int) {
	s := []Dummy{}
	for i := 0; i < n; i++ {
		s = append(s, Dummy{
			Value: 1,
		})
	}
}

// initialize with empty array
func SetSlice02(n int) {
	s := make([]Dummy, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, Dummy{
			Value: 1,
		})
	}
}

// initialize with full zero value array
// and assign new struct by index access
func SetSlice03(n int) {
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
func SetSlice04(n int) {
	s := make([]Dummy, n)
	for i := 0; i < n; i++ {
		s[i].Value = 1
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

// declare a variable by assigning type
func InitSlice01(n int) {
	for i := 0; i < n; i++ {
		var _ []Dummy
	}
}

// declare a variable by type inference
func InitSlice02(n int) {
	for i := 0; i < n; i++ {
		_ = []Dummy{}
	}
}
