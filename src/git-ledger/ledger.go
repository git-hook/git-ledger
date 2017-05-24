package ledger

import (
	"path"
	"github.com/jmalloc/grit/src/grit/pathutil"
)

type Record struct {
	Path string
	Slug string
}

func Path() string {
	home, _ := pathutil.HomeDir()
	return path.Join(home, ".git-ledger")
}
