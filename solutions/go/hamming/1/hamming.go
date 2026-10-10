package hamming

import "errors"

func Distance(a, b string) (int, error) {
    mismatches := 0
	runesA := []rune(a)
    runesB := []rune(b)

	if len(runesA) != len(runesB) {
        return 0, errors.New("Must contain same-length strings")
    }
    
	for i := 0; i < len(runesA); i++ {
        if runesA[i] != runesB[i] {
            mismatches++
        }
    }

    return mismatches, nil
}
