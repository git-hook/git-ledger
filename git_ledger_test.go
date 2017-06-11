package ledger_test

import (
	. "github.com/git-hook/git-ledger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"fmt"
	"path"
)
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

// Create a dummy ledger-file to act as the device-under-test.
func createDummyLedger() {
	r = Record{Path: "/tmp", Slug: "pony/fill"}
	r.AddToLedger()
	records, _ := GetRecords()
	Expect(len(records) == 1)

	s = Record{Path: "/home", Slug: "badger/bear"}
	s.AddToLedger()
	records, _ = GetRecords()
	Expect(len(records) == 2)
}

var (
	r Record
	s Record
)

var _ = Describe("GitLedger", func() {

	BeforeEach(func() {
		pushGitLedger()
		createDummyLedger()
	})
	AfterEach(popGitLedger)

	// TODO: test that git-ledger silently creates the ledger if asked to read it

	Describe("CLI", func() {
		Context("add", func() {
			Specify("should accept fully-initialized records", func() {
				Record{Path: "/tmp", Slug: "pony/fill"}.AddToLedger()
				records, _ := GetRecords()
				Expect(len(records) == 1)

				Record{Path: "/home", Slug: "badger/bear"}.AddToLedger()
				records, _ = GetRecords()
				Expect(len(records) == 2)
			})
		})

		Context("remove", func() {
			Specify("should accept fully-initialized records", func() {
				s.RemoveFromLedger()
				records, _ := GetRecords()
				Expect(len(records) == 1)
				r.RemoveFromLedger()
				records, _ = GetRecords()
				Expect(len(records) == 0)
			})
		})

		// Context("find", func() {
		// 	var (
		// 		r Record
		// 		s Record
		// 	)

		// 	Specify("")
		// })
	})

	Context("The ledger", func() {
		Specify("should reside in the user's home directory", func() {
			Expect(path.Join(os.Getenv("HOME"), ".git-ledger") == Path())
		})
	})

	Describe("Records", func() {

		Context("When the ledger is empty", func() {
			It("should return 0 records", func() {
				records, _ := GetRecords()
				Expect(len(records) == 0)
			})
		})

		// fixme: make this format more awesome
		Specify("should print in an expected format", func() {
			path := "/path/to/nowhere"
			slug := "sclopio/peepio"
			record := Record{Path: path, Slug: slug}
			expected := fmt.Sprintf("[[Record]]\npath = \"%s\"\nslug = \"%s\"\n\n", record.Path, record.Slug)
			Expect(record.String() == expected)
		})

		Specify("should be index-able by slug", func() {
			rec, _ := GetBySlug("fill")
			Expect(rec == r)
			rec, _ = GetBySlug("pony/fill")
			Expect(rec == r)
			rec, _ = GetBySlug("badger")
			Expect(rec == s)
		})

		Specify("should be index-able by path", func() {
			rec, _ := GetBySlug("/tmp")
			Expect(rec == r)
			rec, _ = GetBySlug("/home")
			Expect(rec == s)
			_, err := GetBySlug("DNE")
			Expect(err != nil)
		})
	})

	Describe("Detecting git meta-data", func() {
		Context("Git repositories", func() {
			var (
				dir string
				gitdir string
			)

			BeforeEach(func() {
				dir = "/tmp"
				gitdir = path.Join(dir, ".git")
			})

			AfterEach(func() {
				os.Remove(gitdir)
			})

			Specify("should identify as such", func() {
				os.MkdirAll(gitdir, os.ModePerm)
				Expect(IsGitProject(dir))
			})

			Specify("should be differentiable from non-git repositories", func() {
				Expect(!IsGitProject(dir))
			})
		})
	})
})
