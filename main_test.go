package main

import (
	"testing"
)

const (
	defaultImportPath = "host.xz/path/to/repo.git"
)

func TestImportPath(t *testing.T) {
	ts := []struct {
		in  string
		out string
	}{
		// https://git-scm.com/docs/git-clone#URLS
		{"ssh://host.xz:22/path/to/repo.git/", defaultImportPath},
		{"ssh://user@host.xz:22/path/to/repo.git/", defaultImportPath},
		{"git://host.xz/path/to/repo.git/", defaultImportPath},
		{"git://host.xz:999/path/to/repo.git/", defaultImportPath},
		{"http://host.xz:80/path/to/repo.git/", defaultImportPath},
		{"https://host.xz:443/path/to/repo.git", defaultImportPath},
		{"user@host.xz:path/to/repo.git/", defaultImportPath},
		{"host.xz:path/to/repo.git/", defaultImportPath},

		{"host.xz:repo.git", "host.xz/repo.git"},
		{"/foo:bar/repo.git/", "localhost/foo:bar/repo.git"},
		{"/path/to/repo.git/", "localhost/path/to/repo.git"},
		{"file:///path/to/repo.git/", "localhost/path/to/repo.git"},

		// TODO - ssh://{startsb}user@{endsb}host.xz{startsb}:port{endsb}/~{startsb}user{endsb}/path/to/repo.git/
		// TODO - git://host.xz{startsb}:port{endsb}/~{startsb}user{endsb}/path/to/repo.git/
		// TODO - {startsb}user@{endsb}host.xz:/~{startsb}user{endsb}/path/to/repo.git/
	}

	for _, tt := range ts {
		got := importPath(tt.in)
		if got != tt.out {
			t.Errorf("importPath(%v) -> %v, expected %v\n", tt.in, got, tt.out)
		}
	}
}
