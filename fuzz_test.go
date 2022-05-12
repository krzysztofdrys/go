package jsoniter

import (
	stdjson "encoding/json"
	"testing"
	"unicode/utf8"
)

func FuzzJsonIter(f *testing.F) {
	var json = ConfigCompatibleWithStandardLibrary

	f.Add(`{"a":1,"b":2,"c":3}`)
	f.Add(`{"a":[1,2,3]}`)
	f.Add(`{"a":{"b":{"c":1}}}`)

	f.Fuzz(func(t *testing.T, s string) {
		bb := []byte(s)
		if !utf8.ValidString(s) {
			t.Skip("non utf")
		}
		if !stdjson.Valid(bb) {
			t.Skip("invalid json")
		}

		v2 := map[string]interface{}{}
		if err := json.Unmarshal(bb, &v2); err != nil {
			t.Fatalf("Failed to unmarshal: %s", err)
		}
		if len(v2) >= 3 {
			panic("did it!")
		}
	})
}
