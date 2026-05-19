package goquery

import "golang.org/x/net/html"

// Filter reduces the set of matched elements to those that match the selector string.
// It returns a new Selection object for this subset of matching elements.
func (s *Selection) Filter(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// FilterMatcher reduces the set of matched elements to those that match
// the given matcher. It returns a new Selection object for this subset
// of matching elements.
func (s *Selection) FilterMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// Not removes elements from the Selection that match the selector string.
// It returns a new Selection object with the matching elements removed.
func (s *Selection) Not(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// NotMatcher removes elements from the Selection that match the given matcher.
// It returns a new Selection object with the matching elements removed.
func (s *Selection) NotMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// FilterFunction reduces the set of matched elements to those that pass the function's test.
// It returns a new Selection object for this subset of elements.
func (s *Selection) FilterFunction(f func(int, *Selection) bool) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NotFunction removes elements from the Selection that pass the function's test.
// It returns a new Selection object with the matching elements removed.
func (s *Selection) NotFunction(f func(int, *Selection) bool) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// FilterNodes reduces the set of matched elements to those that match the specified nodes.
// It returns a new Selection object for this subset of elements.
func (s *Selection) FilterNodes(nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NotNodes removes elements from the Selection that match the specified nodes.
// It returns a new Selection object with the matching elements removed.
func (s *Selection) NotNodes(nodes ...*html.Node) *Selection { _ = "STUB: not implemented"; return nil }

// FilterSelection reduces the set of matched elements to those that match a
// node in the specified Selection object.
// It returns a new Selection object for this subset of elements.
func (s *Selection) FilterSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NotSelection removes elements from the Selection that match a node in the specified
// Selection object. It returns a new Selection object with the matching elements removed.
func (s *Selection) NotSelection(sel *Selection) *Selection { _ = "STUB: not implemented"; return nil }

// Intersection is an alias for FilterSelection.
func (s *Selection) Intersection(sel *Selection) *Selection { _ = "STUB: not implemented"; return nil }

// Has reduces the set of matched elements to those that have a descendant
// that matches the selector.
// It returns a new Selection object with the matching elements.
func (s *Selection) Has(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// HasMatcher reduces the set of matched elements to those that have a descendant
// that matches the matcher.
// It returns a new Selection object with the matching elements.
func (s *Selection) HasMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// HasNodes reduces the set of matched elements to those that have a
// descendant that matches one of the nodes.
// It returns a new Selection object with the matching elements.
func (s *Selection) HasNodes(nodes ...*html.Node) *Selection { _ = "STUB: not implemented"; return nil }

// Add all nodes that contain one of the specified nodes

// HasSelection reduces the set of matched elements to those that have a
// descendant that matches one of the nodes of the specified Selection object.
// It returns a new Selection object with the matching elements.
func (s *Selection) HasSelection(sel *Selection) *Selection { _ = "STUB: not implemented"; return nil }

// End ends the most recent filtering operation in the current chain and
// returns the set of matched elements to its previous state.
func (s *Selection) End() *Selection { _ = "STUB: not implemented"; return nil }

// Filter based on the matcher, and the indicator to keep (Filter) or
// to get rid of (Not) the matching elements.
func winnow(sel *Selection, m Matcher, keep bool) []*html.Node {
	_ = "STUB: not implemented"
	// Optimize if keep is requested
	return nil
}

// Use grep

// Filter based on an array of nodes, and the indicator to keep (Filter) or
// to get rid of (Not) the matching elements.
func winnowNodes(sel *Selection, nodes []*html.Node, keep bool) []*html.Node {
	_ = "STUB: not implemented"
	return nil
}

// Filter based on a function test, and the indicator to keep (Filter) or
// to get rid of (Not) the matching elements.
func winnowFunction(sel *Selection, f func(int, *Selection) bool, keep bool) []*html.Node {
	_ = "STUB: not implemented"
	return nil
}
