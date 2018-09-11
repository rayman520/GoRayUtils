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
