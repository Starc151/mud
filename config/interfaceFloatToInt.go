package conf

func InterfaceFloatToInt(in interface{}) int {
	if in == nil {
		return 0
	}
	tmp := in.(float64)
	return int(tmp)
}
