package ledger

import (
	"os"
	"fmt"
	"path"
	"errors"
	"strings"
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

func GetRecords() ([]Record, error) {
	var ledger Records

	b, err := ioutil.ReadFile(Path())
	if err != nil {
		return ledger.Record, err
	}
	str := string(b)

	if _, err := toml.Decode(str, &ledger); err != nil {
		return ledger.Record, err
	}
	return ledger.Record, err
}

func GetBySlug(slug string) (Record, error) {
	var match Record
	var ledger Records

	// TODO: use getrecords
	b, err := ioutil.ReadFile(Path())
	if err != nil {
		return match, err
	}
	str := string(b)

	if _, err := toml.Decode(str, &ledger); err != nil {
		return match, err
	}

	for _, r := range ledger.Record {
		if strings.Contains(r.Slug, slug) {
			return r, nil
		}
	}
	return match, errors.New(fmt.Sprintf("Unknown project: %s", slug))
}

// fixme: remove duplicated code
func (r Record) String() string {
	return fmt.Sprintf("[[Record]]\npath = \"%s\"\nslug = \"%s\"\n\n", r.Path, r.Slug)
}

// comparison by Path
func (r Record) RemoveFromLedger() error {
	b, err := ioutil.ReadFile(Path())
	if err != nil {
		return err
	}
	str := string(b)

	var ledger Records
	if _, err := toml.Decode(str, &ledger); err != nil {
		return err
	}

	var content string
	for _, s := range ledger.Record {
		if s.Path != r.Path {
			content = content + s.String() + "\n"
		}
	}

	ioutil.WriteFile(Path(), []byte(content), 0644)
	return nil
}

func (r Record) AddToLedger() {
	f, err := os.OpenFile(Path(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, r.String())
}
