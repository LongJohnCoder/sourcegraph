package reader

import "github.com/sourcegraph/sourcegraph/lib/codeintel/lsif/protocol/reader"

// LineContext holds a line index and the element parsed from that line.
type LineContext struct {
	Index   int
	Element reader.Element
}
