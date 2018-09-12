package rayUtils

func ConvertInterfaceToStringArrays(array []interface{}) []string {
	var out []string
	for _, entry := range array {
		out = append(out, entry.(string))
	}
	return out
}

func IsContainedInStringArray(val string, array []string) bool {
	for _, entry := range array{
		if entry == val {return true}
	}
	return false
}

func IsContainedInArray(val interface{}, array []interface{}) bool {
	switch v := val.(type) {
	case string:
		return IsContainedInStringArray(v, ConvertInterfaceToStringArrays(array))
	case int, int8, int16, int32:
		//TODO
	case uint, uint8, uint16, uint32, uint64:
		//TODO
	case float32, float64:
		//TODO
	case complex64, complex128:
		//TODO
	}
	return false
}
