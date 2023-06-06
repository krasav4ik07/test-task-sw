package usecases

import (
	"http-server/internal/core"
)

type SubstringUsecases struct {
}

func NewSubstring() *SubstringUsecases {
	return &SubstringUsecases{}
}

func (s *SubstringUsecases) SearchMaxUniqueSubstring(str string) (string, error) {
	maxSubstring := &core.Supstring{}
	runs := []rune(str)
	for num, _ := range runs {
		hashSubstring := make(map[rune]int)
		lenTotalSubstring := 0
		for i, run := range runs[num:] {
			if _, ok := hashSubstring[run]; ok {
				break
			}
			hashSubstring[run] = i
			lenTotalSubstring += 1
		}

		if lenTotalSubstring > maxSubstring.Len {
			maxSubstring.Start = num
			maxSubstring.End = num + lenTotalSubstring
			maxSubstring.Len = lenTotalSubstring
		}

		if maxSubstring.Len > len(runs)-(num+1) {
			break
		}
	}
	return string(runs[maxSubstring.Start:maxSubstring.End]), nil
}
