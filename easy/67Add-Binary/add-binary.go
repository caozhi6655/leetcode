package easy

func addBinary(a string, b string) string {
	zeros := ""
	if len(a) > len(b) {
		for i := 0; i < len(a)-len(b); i++ {
			zeros += "0"
		}
		b = zeros + b
	} else {
		for i := 0; i < len(b)-len(a); i++ {
			zeros += "0"
		}
		a = zeros + a
	}
	byteA := []byte(a)
	byteB := []byte(b)
	var result []byte
	var up byte = 0
	for i := len(byteB) - 1; i >= 0; i-- {
		if byteB[i] == 48 && byteA[i] == 48 {
			if up == 0 {
				result = append(result, 48)
			} else {
				result = append(result, 49)
			}
			up = 0

		} else if byteB[i] == 48 && byteA[i] == 49 {
			if up == 0 {
				result = append(result, 49)
				up = 0
			} else {
				result = append(result, 48)
				up = 1
			}

		} else if byteB[i] == 49 && byteA[i] == 48 {
			if up == 0 {
				result = append(result, 49)
				up = 0
			} else {
				result = append(result, 48)
				up = 1
			}
		} else if byteB[i] == 49 && byteA[i] == 49 {
			if up == 0 {
				result = append(result, 48)
			} else {
				result = append(result, 49)
			}
			up = 1
		}
	}

	reverse := make([]byte, len(result))
	for i := 0; i < len(result); i++ {
		reverse[len(result)-i-1] = result[i]
	}
	s := string(reverse)
	if up == 1 {
		s = "1" + s
	}
	return s
}
