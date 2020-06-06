package stringutil

// Fullname exported func	
func Fullname(f, l string )(string, int){
	full := f + " " + l
	length:=len(full)
	return full, length
}

// FullNameNakedReturn Naked func
func FullNameNakedReturn(f, l string)(full string, length int){
	full = f + " " + l
	length=len(full)
	return
}