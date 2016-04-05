package testutils

import "testing"

func ExpectUint8SlicesEqual(t *testing.T, a1, a2 []uint8) {
	if len(a1) != len(a2) {
		t.Errorf("[]uint8 slice 1 had length %d, slice 2 had length %d", a1, a2)
	}

	differences := []uint8{}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			differences = append(differences, a1[i], a2[i])
		}
	}

	if len(differences) != 0 {
		for i := 0; i < len(differences)/2; i++ {
			t.Logf("Index %d\t differs: %d\t%d", i, differences[2*i], differences[2*i+1])
		}
		t.Errorf("[]unit8 slice 1 and slice 2 were different")
	}
}
