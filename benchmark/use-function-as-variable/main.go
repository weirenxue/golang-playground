package main

func stringCombineGlobal(str1, str2 string) string {
	return str1 + str2
}

func StringCombine1(str1, str2 string) string {
	return stringCombineGlobal(str1, str2)
}

func StringCombine2(str1, str2 string) string {
	c := stringCombineGlobal
	return c(str1, str2)
}

func main() {
}
