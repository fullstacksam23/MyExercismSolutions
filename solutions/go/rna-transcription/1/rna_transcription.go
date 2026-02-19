package strand


func ToRNA(dna string) string {
    comp := map[byte]byte{
        'G': 'C',
        'C': 'G',
        'T': 'A',
        'A': 'U',
    }
	res := make([]byte, len(dna))
    for i, _ := range dna {
        res[i]=comp[dna[i]]
    }
    return string(res)
}
