package ledger

import (
	"os"
	"fmt"
	"path"
	"io/ioutil"

	"github.com/jmalloc/grit/src/grit/pathutil"
	"github.com/BurntSushi/toml"
)

type Record struct {
	Path string
	Slug string
}

type Records struct {
	Record []Record
}

func Path() string {
	home, _ := pathutil.HomeDir()
	return path.Join(home, ".git-ledger")
}

func (r Record) String() string {
	return fmt.Sprintf("[[Record]]\npath = \"%s\"\nslug = \"%s\"\n\n", r.Path, r.Slug)
}

func (r Record) RemoveFromLedger() {
	b, err := ioutil.ReadFile(Path())
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)

	var ledger Records
	if _, err := toml.Decode(str, &ledger); err != nil {
		panic(err)
	}

	var content string
	for _, s := range ledger.Record {
		if s.Path != r.Path {
			content = content + s.String() + "\n"
		}
	}

	ioutil.WriteFile(Path(), []byte(content), 0644)

}

func (r Record) WriteToLedger() {
	f, err := os.OpenFile(Path(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, r.String())
}
