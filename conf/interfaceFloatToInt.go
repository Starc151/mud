package conf

func InterfaceFloatTouint16(in interface{}) uint16 {
	if in == nil {
		return 0
	}
	tmp := in.(float64)
	return uint16(tmp)
}
