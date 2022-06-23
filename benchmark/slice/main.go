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
// and assign value by index
func SetSlice03(n int) {
	s := make([]Dummy, n)
	for i := 0; i < n; i++ {
		s[i] = Dummy{
			Value: 1,
		}
	}
}
