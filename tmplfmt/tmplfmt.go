package tmplfmt

import (
	"github.com/bep/gotmplfmt/internal/parse"
)

func Format(text string) (string, error) {
	root, err := parse.Parse(text)
	if err != nil {
		return "", err
	}
	// TODO: probably want to move all the printing logic out of the nodes
	// and into something more flexible here.
	return root.String(), nil
}
