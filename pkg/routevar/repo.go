package routevar

import (
	"regexp"
	"strings"

	"sourcegraph.com/sourcegraph/sourcegraph/go-sourcegraph/sourcegraph"
)

// A RepoRev specifies a repo at a revision. Unlike
// sourcegraph.RepoRevSpec, the revision need not be an absolute
// commit ID. This RepoRev type is appropriate for user input (e.g.,
// from a URL), where it is convenient to allow users to specify
// non-absolute commit IDs that the server can resolve.
type RepoRev struct {
	sourcegraph.RepoSpec
	Rev string // a VCS revision specifier (branch, "master~7", commit ID, etc.)
}

// A TreeEntry specifies a tree/blob by path in a repo at a
// revision. It is analogous to sourcegraph.TreeEntrySpec, but like
// RepoRev, it allows for a non-absolute commit ID.
type TreeEntry struct {
	RepoRev
	Path string // path to the VCS tree/blob
}

var (
	Repo = `{Repo:` + namedToNonCapturingGroups(RepoPattern) + `}`
	Rev  = `{Rev:` + namedToNonCapturingGroups(RevPattern) + `}`

	RepoRevSuffix = `{Rev:` + namedToNonCapturingGroups(`(?:@`+RevPattern+`)?`) + `}`
)

const (
	// RepoPattern is the regexp pattern that matches RepoSpec strings
	// ("repo" or "domain.com/repo" or "domain.com/path/to/repo").
	RepoPattern = `(?P<repo>(?:` + pathComponentNotDelim + `/)*` + pathComponentNotDelim + `)`

	RepoPathDelim         = "-"
	pathComponentNotDelim = `(?:[^@/` + RepoPathDelim + `]|(?:[^/@]{2,}))`

	// RevPattern is the regexp pattern that matches a VCS revision
	// specifier (e.g., "master" or "my/branch~1", or a full 40-char
	// commit ID).
	RevPattern = `(?P<rev>(?:` + pathComponentNotDelim + `/)*` + pathComponentNotDelim + `)`

	// CommitPattern is the regexp pattern that matches absolute
	// (40-character) hexidecimal commit IDs.
	CommitPattern = `(?P<commit>[[:xdigit:]]{40})`
)

var (
	repoPattern = regexp.MustCompile("^" + RepoPattern + "$")
	revPattern  = regexp.MustCompile("^" + RevPattern + "$")
)

// ParseRepo parses a RepoSpec string. If spec is invalid, an
// InvalidError is returned.
func ParseRepo(spec string) (repo string, err error) {
	if m := repoPattern.FindStringSubmatch(spec); len(m) > 0 {
		repo = m[0]
		return
	}
	return "", InvalidError{"RepoSpec", spec, nil}
}

// RepoRouteVars returns route variables for constructing repository
// routes.
func RepoRouteVars(s sourcegraph.RepoSpec) map[string]string {
	return map[string]string{"Repo": s.URI}
}

// ToRepoRev marshals a map containing route variables
// generated by (RepoRevSpec).RouteVars() and returns the equivalent
// RepoRevSpec struct.
func ToRepoRev(routeVars map[string]string) RepoRev {
	repo := ToRepoSpec(routeVars)
	rrspec := RepoRev{RepoSpec: repo}
	if revStr := routeVars["Rev"]; revStr != "" {
		if !strings.HasPrefix(revStr, "@") {
			panic("Rev should have had '@' prefix from route")
		}
		rrspec.Rev = strings.TrimPrefix(revStr, "@")
	}
	if _, ok := routeVars["CommitID"]; ok {
		panic("unexpected CommitID route var; was removed in the simple-routes branch")
	}
	return rrspec
}

// ToRepoSpec marshals a map containing route variables
// generated by (*RepoSpec).RouteVars() and returns the
// equivalent RepoSpec struct.
func ToRepoSpec(routeVars map[string]string) sourcegraph.RepoSpec {
	return sourcegraph.RepoSpec{URI: routeVars["Repo"]}
}

// RepoRevRouteVars returns route variables for constructing routes to a
// repository revision.
func RepoRevRouteVars(s RepoRev) map[string]string {
	m := RepoRouteVars(s.RepoSpec)
	var rev string
	if s.Rev != "" {
		rev = "@" + s.Rev
	}
	m["Rev"] = rev
	return m
}

func TreeEntryRouteVars(s TreeEntry) map[string]string {
	m := RepoRevRouteVars(s.RepoRev)
	m["Path"] = s.Path
	return m
}
