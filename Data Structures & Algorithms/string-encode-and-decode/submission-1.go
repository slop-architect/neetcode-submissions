import (
	"unicode/utf8"
)

const DELIMETER rune = '/'

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var resultStr string
	for _, val := range strs {
		partStr := fmt.Sprintf("%d", utf8.RuneCountInString(val)) + string(DELIMETER) + val
		resultStr += partStr
	}
	return resultStr
}

func (s *Solution) Decode(encoded string) []string {
	var numStr string
	var strs = []string{}

	runes := []rune(encoded)
	encodedLen := utf8.RuneCountInString(encoded)

	for i := 0; i < encodedLen; i++ {

		numStr += string(runes[i])
		if runes[i+1] != DELIMETER {
			continue
		} else {
			num, _ := strconv.Atoi(numStr)
			strs = append(strs, string(runes[i+2:i+2+num]))
			i += num + 1
			numStr = ""
		}
		if i >= encodedLen {
			break
		}
	}
	return strs
}
