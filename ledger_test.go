package ledger

import (
	"os"
	"fmt"
	"path"
	"testing"
)

func TestPath(t *testing.T) {
	expected := path.Join(os.Getenv("HOME"), ".git-ledger")
	path := Path()
	if path != expected {
		t.Error(
			"For", "Path()",
			"expected", expected,
			"got", path,
		)
	}
}

func TestIsGitProject(t *testing.T) {
	dir := "/tmp"
	gitdir := path.Join(dir, ".git")
	os.MkdirAll(gitdir, os.ModePerm)
	result := IsGitProject(dir)
	if result != true {
		t.Error(
			"For", "IsGitProject(dir)",
			"expected", true,
			"got", result,
		)
	}


	// remove dir and test for false
	os.Remove(gitdir)
	result = IsGitProject(dir)
	if result != false {
		t.Error(
			"For", "IsGitProject(dir)",
			"expected", false,
			"got", result,
		)
	}
}

// TODO: add test for nil fields, see how the toml parses?
func TestRecordString(t *testing.T) {
	path := "/path/to/nowhere"
	slug := "sclopio/peepio"
	record := Record{Path: path, Slug: slug}
	expected := fmt.Sprintf("[[Record]]\npath = \"%s\"\nslug = \"%s\"\n\n", record.Path, record.Slug)
	if record.String() != expected {
		t.Error(
			"For", "Record.String()",
			"expected", expected,
			"got", record.String(),
		)
	}
}

// Move the current git-ledger to a new location so testing can create
// a temporary one
func pushGitLedger() {
	os.Rename(Path(), fmt.Sprintf("%s.backup", Path()))
	f, _ := os.Create(Path())
	defer f.Close()
}

// Remove the temporary git-ledger and replace it with the backup
func popGitLedger() {
	os.Rename(fmt.Sprintf("%s.backup", Path()), Path())
}

func TestGetRecordsEmpty( t *testing.T) {
	pushGitLedger()
	records, _ := GetRecords()
	if len(records) != 0 {
		t.Error(
			"For", "len(GetRecords())",
			"expected", 0,
			"got", len(records),
		)
	}
	popGitLedger()
}

func TestAddToLedger(t *testing.T) {
	pushGitLedger()
	r := Record{Path: "/tmp", Slug: "pony/fill"}
	r.AddToLedger()
	records, _ := GetRecords()
	if len(records) != 1 {
		t.Error(
			"For", "len(GetRecords())",
			"expected", 1,
			"got", len(records),
		)
	}
	s := Record{Path: "/home", Slug: "badger/bear"}
	s.AddToLedger()
	records, _ = GetRecords()
	if len(records) != 2 {
		t.Error(
			"For", "len(GetRecords())",
			"expected", 2,
			"got", len(records),
		)
	}
	popGitLedger()
}

func TestRemoveFromLedger(t *testing.T) {
	pushGitLedger()
	r := Record{Path: "/tmp", Slug: "pony/fill"}
	r.AddToLedger()
	records, _ := GetRecords()
	if len(records) != 1 {
		t.Error(
			"For", "len(GetRecords())",
			"expected", 1,
			"got", len(records),
		)
	}
	s := Record{Path: "/home", Slug: "badger/bear"}
	s.AddToLedger()
	records, _ = GetRecords()
	if len(records) != 2 {
		t.Error(
			"For", "len(GetRecords())",
			"expected", 2,
			"got", len(records),
		)
	}
	s.RemoveFromLedger()
	records, _ = GetRecords()
	if len(records) != 1 {
		t.Error(
			"For", "len(GetRecords())",
			"expected", 1,
			"got", len(records),
		)
	}
	r.RemoveFromLedger()
	records, _ = GetRecords()
	if len(records) != 0 {
		t.Error(
			"For", "len(GetRecords())",
			"expected", 0,
			"got", len(records),
		)
	}
	popGitLedger()
}

func TestGetBySlug(t *testing.T) {
	pushGitLedger()
	r := Record{Path: "/tmp", Slug: "pony/fill"}
	r.AddToLedger()
	s := Record{Path: "/home", Slug: "badger/bear"}
	s.AddToLedger()
	rec, err := GetBySlug("fill")
	if rec != r {
		t.Error(
			"For", "GetBySlug(fill)",
			"expected", r,
			"got", rec,
		)
	}
	rec, err = GetBySlug("pony/fill")
	if rec != r || err != nil {
		t.Error(
			"For", "GetBySlug(pony/fill)",
			"expected", r,
			"got", rec,
		)
	}
	rec2, err2 := GetBySlug("badger")
	if rec2 != s || err2 != nil {
		t.Error(
			"For", "GetBySlug(badger)",
			"expected", s,
			"got", rec2,
		)
	}
	popGitLedger()
}

func TestGetByPath(t *testing.T) {
	pushGitLedger()
	r := Record{Path: "/tmp", Slug: "pony/fill"}
	r.AddToLedger()
	s := Record{Path: "/home", Slug: "badger/bear"}
	s.AddToLedger()
	rec, err := GetByPath("/tmp")
	if rec != r {
		t.Error(
			"For", "GetByPath(/tmp)",
			"expected", r,
			"got", rec,
		)
	}
	rec, err = GetByPath("/home")
	if rec != s || err != nil {
		t.Error(
			"For", "GetByPath(/home)",
			"expected", s,
			"got", rec,
		)
	}
	rec2, err2 := GetByPath("DNE")
	if err2 == nil {
		t.Error(
			"For", "GetByPath(DNE)",
			"expected", "an error",
			"got", rec2,
		)
	}
	popGitLedger()
}
