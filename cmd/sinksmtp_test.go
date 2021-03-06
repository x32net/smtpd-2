//
// Basic testing for file loading, which is stuck in sinksmtp.go for
// reasons that are partly historical.

package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func isPresent(a []string, p string) bool {
	i := sort.SearchStrings(a, p)
	return i < len(a) && a[i] == p
}

// Test readList() by loading a basic list and seeing if the entries we
// expect are present and have been properly lower-cased.
func TestLoader(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(basiclist))
	a, err := readList(reader)
	if err != nil {
		t.Fatalf("Error during read: %v", err)
	}
	sort.Strings(a)
	for _, p := range present {
		if !isPresent(a, p) {
			t.Fatalf("Missing a record: %s", p)
		}
	}
	if isPresent(a, "") {
		t.Fatalf("Blank line present in address list.")
	}
	if isPresent(a, "# t") {
		t.Fatalf("Comment present in address list.")
	}
}

var basiclist = `# This is a comment
INFO@FBI.GOV
root@

@example.com
postmaster@Example.Org
@.barney.net
# t
`

var present = []string{
	"info@fbi.gov", "root@", "@example.com", "postmaster@example.org",
	"@.barney.net",
}
