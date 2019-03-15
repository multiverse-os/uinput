package uinput

//func typeString(barcode string, done chan<- bool) {
//
//	rawRunes := []rune(barcode)
//	var keys []string
//
//	for _, r := range rawRunes {
//		if unicode.IsUpper(r) && !unicode.IsDigit(r) {
//			keys = append(keys, "shift", strings.ToLower(string(r)))
//			continue
//		}
//
//		keys = append(keys, string(r))
//
//	}
//
//	for _, key := range keys {
//		robotgo.KeyTap(key)
//	}
//
//	done <- true
//}
