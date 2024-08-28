package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func Bencodejson(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	base64Encoded := base64.StdEncoding.EncodeToString(jsonData)
	return base64Encoded
}

func Bdencodejson(data string) string {
	base64Decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)
	}
	return string(base64Decoded)
}

func LogicSwitch(i11, i12, i13, i14 int) int {
	i16 := i14 % 10
	var i15 int
	if i16 != 0 {
		i15 = i13 % i16
	} else {
		i15 = i13 % 10
	}

	i17 := i11 * i11
	i18 := i12 * i12

	switch i15 {
	case 0:
		return i12 + i17
	case 1:
	case 2:
		return i12 * i17
	case 3:
		return i12 ^ i11
	case 4:
		return i11 - i18
	case 5:
		i19 := i11 + 783
		i11 = i19 * i19
	case 6:
		return i12 + (i11 ^ i12)
	case 7:
		return i17 - i18
	case 8:
		return i12 * i11
	case 9:
		return (i12 * i11) - i11
	default:
		return -1
	}

	return i11 + i18
}

func BitewiseXOR(deviceModel string, doString []string) int {
	charset := []byte(deviceModel)
	if len(charset) < 4 {
		fmt.Println("Error: deviceModel is too short")
		return -1
	}
	bytesInt := int(charset[0])<<24 | int(charset[1])<<16 | int(charset[2])<<8 | int(charset[3])

	p1, _ := strconv.Atoi(string(doString[6]))
	p2, _ := strconv.Atoi(string(doString[7]))
	p3, _ := strconv.Atoi(string(doString[4]))
	p4, _ := strconv.Atoi(string(doString[9]))

	pp1 := LogicSwitch(p1, p2, p3, p4)
	pp2, _ := strconv.Atoi(string(doString[8]))
	pp3, _ := strconv.Atoi(string(doString[5]))
	pp4 := p4

	xorInt := LogicSwitch(pp1, pp2, pp3, pp4)

	return bytesInt ^ xorInt
}

func RandomHex(length int) string {
	const hexChars = "0123456789abcdef"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = hexChars[rand.Intn(len(hexChars))]
	}
	return string(result)
}

func Round(val float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(val*factor) / factor
}
