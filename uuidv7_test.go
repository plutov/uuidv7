package uuidv7

import "testing"

func BenchmarkUUIDV7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuidVal, err := UUIDV7()
		if err != nil {
			b.Fatal(err)
		}
		if len(uuidVal) != 16 {
			b.Fatalf("uuidVal length is %d, want 16", len(uuidVal))
		}
	}
}
