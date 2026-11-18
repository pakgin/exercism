package rnatranscription

import "strings"

func ToRNA(dna string) string {
	rna := strings.Builder{}
	for _, n := range dna {
		switch n {
		case 'G':
			rna.WriteRune('C')
		case 'C':
			rna.WriteRune('G')
		case 'T':
			rna.WriteRune('A')
		case 'A':
			rna.WriteRune('U')
		}
	}
	return rna.String()
}
