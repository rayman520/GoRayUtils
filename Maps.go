package rayUtils

func InvertMapOfBytes(in map[byte]byte) map[byte]byte {
	out := make(map[byte]byte)
	for k, v := range in {
		out[v] = k
	}
	return out
}
