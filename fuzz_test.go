package jsoniter

import (
	stdjson "encoding/json"
	"testing"
	"unicode/utf8"
)

func FuzzJsonIter(f *testing.F) {
	var json = ConfigCompatibleWithStandardLibrary

	f.Fuzz(func(t *testing.T, s string) {
		bb := []byte(s)
		if !utf8.ValidString(s) {
			t.Skip("non utf")
		}
		if !stdjson.Valid(bb) {
			t.Skip("invalid json")
		}
		v := map[string]interface{}{}
		if err := stdjson.Unmarshal(bb, &v); err != nil {
			t.Skip("invalid json")
		}

		v2 := map[string]interface{}{}
		if err := json.Unmarshal(bb, &v2); err != nil {
			t.Fatalf("Failed to unmarshal: %s", err)
		}
		if len(v2) >= 5 {
			panic("did it!")
		}
	})
}
