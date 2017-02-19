package secret

import "strconv"

const testVersion = 1

var secretHandshakes = []struct {
	code  int
	event string
}{
	{10000, "Reverse the order of the operations in the secret handshake"},
	{1000, "jump"},
	{100, "close your eyes"},
	{10, "double blink"},
	{1, "wink"},
}

func Handshake(code uint) []string {
	codeInBinaryStr := strconv.FormatInt(int64(code), 2)
	codeInBinary, err := strconv.Atoi(codeInBinaryStr)
	var binaryElements []int
	if err == nil {
		binaryElements = GetBinaryElements(codeInBinary)
		return ConvertToEvents(binaryElements)
	}
	return nil
}
func IsPositiveBinary(number int) bool {
	s := strconv.Itoa(number)
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil || i < 0 {
		return false
	}
	return true
}
func GetBinaryElements(binary int) []int {
	var elements []int
	for _, secretHandshake := range secretHandshakes {
		binary -= secretHandshake.code
		if IsPositiveBinary(binary) {
			elements = append(elements, secretHandshake.code)
		} else {
			binary += secretHandshake.code
		}
		if binary == 0 {
			break
		}
	}
	return elements
}
func ConvertToEvents(binaries []int) []string {
	var events []string
	isReverse := 1
	for i := 0; i < len(binaries); i++ {
		for _, secretHandshake := range secretHandshakes {
			if binaries[i] != 10000 && secretHandshake.code == binaries[i] {
				events = append(events, secretHandshake.event)
			} else if binaries[i] == 10000 {
				isReverse = 0
			}
		}
	}
	if isReverse == 1 {
		return Reverse(events)
	} else {
		return events
	}

}
func Reverse(array []string) []string {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
