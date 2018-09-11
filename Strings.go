package rayUtils

import "strings"

func FindNextBracketIndex(str []byte, endchar byte) int {
	i := 0
	len := len(str)
	bracketOpposites := GetBracketOppositesMap()
	for i < len {
		if str[i] == endchar {
			return i + 1
		}
		if bracketOpposites[str[i]] != 0 {
			i += FindNextBracketIndex(str[i + 1:], bracketOpposites[str[i]])
		}
		i++
	}
	return 0
}

func RemoveSubstringsFromString(str string, substrs []string) string {
	var replacerArg []string
	for _, sub := range substrs {
		replacerArg = append(replacerArg, sub, "")
	}
	replacer := strings.NewReplacer(replacerArg...)
	return replacer.Replace(str)
}

func SplitIgnoreBrackets(rawString string, char byte) []string {
	//TODO Clean me
	out := []string{}
	str := []byte(rawString)
	len := len(str)
	startIndex := 0
	CurrentIndex := 0
	BracketOpposites := GetBracketOppositesMap()
	BracketOppositesOpposites := InvertMapOfBytes(GetBracketOppositesMap())
	for CurrentIndex < len {
		if str[CurrentIndex] == char {
			outstr := str[startIndex : CurrentIndex]
			if outstr[0] == char {outstr = outstr[1:]}
			out = append(out, string(outstr))
			startIndex = CurrentIndex
		}
		if BracketOpposites[str[CurrentIndex]] != 0 {
			CurrentIndex += FindNextBracketIndex(str[CurrentIndex + 1:], BracketOpposites[str[CurrentIndex]])
		} else if BracketOppositesOpposites[str[CurrentIndex]] != 0{
			outstr := str[startIndex : CurrentIndex]
			if outstr[0] == char {outstr = outstr[1:]}
			out = append(out, string(outstr))
			startIndex = CurrentIndex
		}
		CurrentIndex++
	}
	outstr := str[startIndex : CurrentIndex]
	if outstr[0] == char {outstr = outstr[1:]}
	out = append(out, string(outstr))
	return out
}
