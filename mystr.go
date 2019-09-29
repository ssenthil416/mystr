package mystr


func StrMirror(inStr string) bool {
	inStrLen := len(inStr)
	j := inStrLen-1
	for i:=0;  i<j; i++ {
		if inStr[i] != inStr[j] {
			return false
		}
		j = j - 1
	} 
	return true
}
