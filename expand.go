package goquery

import "golang.org/x/net/html"

// Add adds the selector string's matching nodes to those in the current
// selection and returns a new Selection object.
// The selector string is run in the context of the document of the current
// Selection object.
func (s *Selection) Add(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// AddMatcher adds the matcher's matching nodes to those in the current
// selection and returns a new Selection object.
// The matcher is run in the context of the document of the current
// Selection object.
func (s *Selection) AddMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// AddSelection adds the specified Selection object's nodes to those in the
// current selection and returns a new Selection object.
func (s *Selection) AddSelection(sel *Selection) *Selection { _ = "STUB: not implemented"; return nil }

// Union is an alias for AddSelection.
func (s *Selection) Union(sel *Selection) *Selection { _ = "STUB: not implemented"; return nil }

// AddNodes adds the specified nodes to those in the
// current selection and returns a new Selection object.
func (s *Selection) AddNodes(nodes ...*html.Node) *Selection { _ = "STUB: not implemented"; return nil }

// AndSelf adds the previous set of elements on the stack to the current set.
// It returns a new Selection object containing the current Selection combined
// with the previous one.
// Deprecated: This function has been deprecated and is now an alias for AddBack().
func (s *Selection) AndSelf() *Selection {
	_ = "STUB: not implemented"

	// AddBack adds the previous set of elements on the stack to the current set.
	// It returns a new Selection object containing the current Selection combined
	// with the previous one.
	return nil
}

func (s *Selection) AddBack() *Selection { _ = "STUB: not implemented"; return nil }

// AddBackFiltered reduces the previous set of elements on the stack to those that
// match the selector string, and adds them to the current set.
// It returns a new Selection object containing the current Selection combined
// with the filtered previous one
func (s *Selection) AddBackFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// AddBackMatcher reduces the previous set of elements on the stack to those that match
// the matcher, and adds them to the current set.
// It returns a new Selection object containing the current Selection combined
// with the filtered previous one
func (s *Selection) AddBackMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }
