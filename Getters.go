package rayUtils

func GetBracketOppositesMap() map[byte]byte {
	if rayUtilsGlobals.BracketOppositesMap == nil {
		out := make(map[byte]byte)
		out['"'] = '"'
		out['\''] = '\''
		out['['] = ']'
		out['{'] = '}'
		out['('] = ')'
		rayUtilsGlobals.BracketOppositesMap = out
	}
	return rayUtilsGlobals.BracketOppositesMap
}

func GetBooleanTrueTranslator() []string {
	if rayUtilsGlobals.BooleanTrueTranslator == nil {
		rayUtilsGlobals.BooleanTrueTranslator = []string{
			"true", "y", "yes", "1",
		}
	}
	return rayUtilsGlobals.BooleanTrueTranslator
}
