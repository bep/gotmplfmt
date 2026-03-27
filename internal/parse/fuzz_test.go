package parse

import "testing"

func FuzzParseString(f *testing.F) {
	samples := []string{
		`{{}}`,
		`{{.}}`,
		`{{.Field}}`,
		`<ul>{{ range . }}<li>{{.Field}}</li>{{ end }}</ul>`,
		`{{.Field | printf "%q"}}`,
		`{{if .}}yes{{else}}no{{end}}`,
	}
	for _, s := range samples {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		root, err := Parse(s)
		if err != nil {
			return
		}
		out := root.String()
		root2, err := Parse(out)
		if err != nil {
			t.Fatalf("failed to parse output: %v", err)
		}
		out2 := root2.String()
		if out != out2 {
			t.Fatalf("round trip failure: %q != %q", out, out2)
		}
	})
}
