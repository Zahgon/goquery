package goquery

import "golang.org/x/net/html"

type siblingType int

// Sibling type, used internally when iterating over children at the same
// level (siblings) to specify which nodes are requested.
const (
	siblingPrevUntil siblingType = iota - 3
	siblingPrevAll
	siblingPrev
	siblingAll
	siblingNext
	siblingNextAll
	siblingNextUntil
	siblingAllIncludingNonElements
)

// Find gets the descendants of each element in the current set of matched
// elements, filtered by a selector. It returns a new Selection object
// containing these matched elements.
//
// Note that as for all methods accepting a selector string, the selector is
// compiled and applied by the cascadia package and inherits its behavior and
// constraints regarding supported selectors. See the note on cascadia in
// the goquery documentation here:
// https://github.com/PuerkitoBio/goquery?tab=readme-ov-file#api
func (s *Selection) Find(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// FindMatcher gets the descendants of each element in the current set of matched
// elements, filtered by the matcher. It returns a new Selection object
// containing these matched elements.
func (s *Selection) FindMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// FindSelection gets the descendants of each element in the current
// Selection, filtered by a Selection. It returns a new Selection object
// containing these matched elements.
func (s *Selection) FindSelection(sel *Selection) *Selection { _ = "STUB: not implemented"; return nil }

// FindNodes gets the descendants of each element in the current
// Selection, filtered by some nodes. It returns a new Selection object
// containing these matched elements.
func (s *Selection) FindNodes(nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// Contents gets the children of each element in the Selection,
// including text and comment nodes. It returns a new Selection object
// containing these elements.
func (s *Selection) Contents() *Selection { _ = "STUB: not implemented"; return nil }

// ContentsFiltered gets the children of each element in the Selection,
// filtered by the specified selector. It returns a new Selection
// object containing these elements. Since selectors only act on Element nodes,
// this function is an alias to ChildrenFiltered unless the selector is empty,
// in which case it is an alias to Contents.
func (s *Selection) ContentsFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ContentsMatcher gets the children of each element in the Selection,
// filtered by the specified matcher. It returns a new Selection
// object containing these elements. Since matchers only act on Element nodes,
// this function is an alias to ChildrenMatcher.
func (s *Selection) ContentsMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// Children gets the child elements of each element in the Selection.
// It returns a new Selection object containing these elements.
func (s *Selection) Children() *Selection { _ = "STUB: not implemented"; return nil }

// ChildrenFiltered gets the child elements of each element in the Selection,
// filtered by the specified selector. It returns a new
// Selection object containing these elements.
func (s *Selection) ChildrenFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ChildrenMatcher gets the child elements of each element in the Selection,
// filtered by the specified matcher. It returns a new
// Selection object containing these elements.
func (s *Selection) ChildrenMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// Parent gets the parent of each element in the Selection. It returns a
// new Selection object containing the matched elements.
func (s *Selection) Parent() *Selection { _ = "STUB: not implemented"; return nil }

// ParentFiltered gets the parent of each element in the Selection filtered by a
// selector. It returns a new Selection object containing the matched elements.
func (s *Selection) ParentFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentMatcher gets the parent of each element in the Selection filtered by a
// matcher. It returns a new Selection object containing the matched elements.
func (s *Selection) ParentMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// Closest gets the first element that matches the selector by testing the
// element itself and traversing up through its ancestors in the DOM tree.
func (s *Selection) Closest(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// ClosestMatcher gets the first element that matches the matcher by testing the
// element itself and traversing up through its ancestors in the DOM tree.
func (s *Selection) ClosestMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// For each node in the selection, test the node itself, then each parent
// until a match is found.

// ClosestNodes gets the first element that matches one of the nodes by testing the
// element itself and traversing up through its ancestors in the DOM tree.
func (s *Selection) ClosestNodes(nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// For each node in the selection, test the node itself, then each parent
// until a match is found.

// ClosestSelection gets the first element that matches one of the nodes in the
// Selection by testing the element itself and traversing up through its ancestors
// in the DOM tree.
func (s *Selection) ClosestSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// Parents gets the ancestors of each element in the current Selection. It
// returns a new Selection object with the matched elements.
func (s *Selection) Parents() *Selection { _ = "STUB: not implemented"; return nil }

// ParentsFiltered gets the ancestors of each element in the current
// Selection. It returns a new Selection object with the matched elements.
func (s *Selection) ParentsFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsMatcher gets the ancestors of each element in the current
// Selection. It returns a new Selection object with the matched elements.
func (s *Selection) ParentsMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// ParentsUntil gets the ancestors of each element in the Selection, up to but
// not including the element matched by the selector. It returns a new Selection
// object containing the matched elements.
func (s *Selection) ParentsUntil(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// ParentsUntilMatcher gets the ancestors of each element in the Selection, up to but
// not including the element matched by the matcher. It returns a new Selection
// object containing the matched elements.
func (s *Selection) ParentsUntilMatcher(m Matcher) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsUntilSelection gets the ancestors of each element in the Selection,
// up to but not including the elements in the specified Selection. It returns a
// new Selection object containing the matched elements.
func (s *Selection) ParentsUntilSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsUntilNodes gets the ancestors of each element in the Selection,
// up to but not including the specified nodes. It returns a
// new Selection object containing the matched elements.
func (s *Selection) ParentsUntilNodes(nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsFilteredUntil is like ParentsUntil, with the option to filter the
// results based on a selector string. It returns a new Selection
// object containing the matched elements.
func (s *Selection) ParentsFilteredUntil(filterSelector, untilSelector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsFilteredUntilMatcher is like ParentsUntilMatcher, with the option to filter the
// results based on a matcher. It returns a new Selection object containing the matched elements.
func (s *Selection) ParentsFilteredUntilMatcher(filter, until Matcher) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsFilteredUntilSelection is like ParentsUntilSelection, with the
// option to filter the results based on a selector string. It returns a new
// Selection object containing the matched elements.
func (s *Selection) ParentsFilteredUntilSelection(filterSelector string, sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsMatcherUntilSelection is like ParentsUntilSelection, with the
// option to filter the results based on a matcher. It returns a new
// Selection object containing the matched elements.
func (s *Selection) ParentsMatcherUntilSelection(filter Matcher, sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsFilteredUntilNodes is like ParentsUntilNodes, with the
// option to filter the results based on a selector string. It returns a new
// Selection object containing the matched elements.
func (s *Selection) ParentsFilteredUntilNodes(filterSelector string, nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ParentsMatcherUntilNodes is like ParentsUntilNodes, with the
// option to filter the results based on a matcher. It returns a new
// Selection object containing the matched elements.
func (s *Selection) ParentsMatcherUntilNodes(filter Matcher, nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// Siblings gets the siblings of each element in the Selection. It returns
// a new Selection object containing the matched elements.
func (s *Selection) Siblings() *Selection { _ = "STUB: not implemented"; return nil }

// SiblingsFiltered gets the siblings of each element in the Selection
// filtered by a selector. It returns a new Selection object containing the
// matched elements.
func (s *Selection) SiblingsFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// SiblingsMatcher gets the siblings of each element in the Selection
// filtered by a matcher. It returns a new Selection object containing the
// matched elements.
func (s *Selection) SiblingsMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// Next gets the immediately following sibling of each element in the
// Selection. It returns a new Selection object containing the matched elements.
func (s *Selection) Next() *Selection { _ = "STUB: not implemented"; return nil }

// NextFiltered gets the immediately following sibling of each element in the
// Selection filtered by a selector. It returns a new Selection object
// containing the matched elements.
func (s *Selection) NextFiltered(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// NextMatcher gets the immediately following sibling of each element in the
// Selection filtered by a matcher. It returns a new Selection object
// containing the matched elements.
func (s *Selection) NextMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// NextAll gets all the following siblings of each element in the
// Selection. It returns a new Selection object containing the matched elements.
func (s *Selection) NextAll() *Selection { _ = "STUB: not implemented"; return nil }

// NextAllFiltered gets all the following siblings of each element in the
// Selection filtered by a selector. It returns a new Selection object
// containing the matched elements.
func (s *Selection) NextAllFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NextAllMatcher gets all the following siblings of each element in the
// Selection filtered by a matcher. It returns a new Selection object
// containing the matched elements.
func (s *Selection) NextAllMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// Prev gets the immediately preceding sibling of each element in the
// Selection. It returns a new Selection object containing the matched elements.
func (s *Selection) Prev() *Selection { _ = "STUB: not implemented"; return nil }

// PrevFiltered gets the immediately preceding sibling of each element in the
// Selection filtered by a selector. It returns a new Selection object
// containing the matched elements.
func (s *Selection) PrevFiltered(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// PrevMatcher gets the immediately preceding sibling of each element in the
// Selection filtered by a matcher. It returns a new Selection object
// containing the matched elements.
func (s *Selection) PrevMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// PrevAll gets all the preceding siblings of each element in the
// Selection. It returns a new Selection object containing the matched elements.
func (s *Selection) PrevAll() *Selection { _ = "STUB: not implemented"; return nil }

// PrevAllFiltered gets all the preceding siblings of each element in the
// Selection filtered by a selector. It returns a new Selection object
// containing the matched elements.
func (s *Selection) PrevAllFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevAllMatcher gets all the preceding siblings of each element in the
// Selection filtered by a matcher. It returns a new Selection object
// containing the matched elements.
func (s *Selection) PrevAllMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// NextUntil gets all following siblings of each element up to but not
// including the element matched by the selector. It returns a new Selection
// object containing the matched elements.
func (s *Selection) NextUntil(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// NextUntilMatcher gets all following siblings of each element up to but not
// including the element matched by the matcher. It returns a new Selection
// object containing the matched elements.
func (s *Selection) NextUntilMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// NextUntilSelection gets all following siblings of each element up to but not
// including the element matched by the Selection. It returns a new Selection
// object containing the matched elements.
func (s *Selection) NextUntilSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NextUntilNodes gets all following siblings of each element up to but not
// including the element matched by the nodes. It returns a new Selection
// object containing the matched elements.
func (s *Selection) NextUntilNodes(nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevUntil gets all preceding siblings of each element up to but not
// including the element matched by the selector. It returns a new Selection
// object containing the matched elements.
func (s *Selection) PrevUntil(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// PrevUntilMatcher gets all preceding siblings of each element up to but not
// including the element matched by the matcher. It returns a new Selection
// object containing the matched elements.
func (s *Selection) PrevUntilMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// PrevUntilSelection gets all preceding siblings of each element up to but not
// including the element matched by the Selection. It returns a new Selection
// object containing the matched elements.
func (s *Selection) PrevUntilSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevUntilNodes gets all preceding siblings of each element up to but not
// including the element matched by the nodes. It returns a new Selection
// object containing the matched elements.
func (s *Selection) PrevUntilNodes(nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NextFilteredUntil is like NextUntil, with the option to filter
// the results based on a selector string.
// It returns a new Selection object containing the matched elements.
func (s *Selection) NextFilteredUntil(filterSelector, untilSelector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NextFilteredUntilMatcher is like NextUntilMatcher, with the option to filter
// the results based on a matcher.
// It returns a new Selection object containing the matched elements.
func (s *Selection) NextFilteredUntilMatcher(filter, until Matcher) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NextFilteredUntilSelection is like NextUntilSelection, with the
// option to filter the results based on a selector string. It returns a new
// Selection object containing the matched elements.
func (s *Selection) NextFilteredUntilSelection(filterSelector string, sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NextMatcherUntilSelection is like NextUntilSelection, with the
// option to filter the results based on a matcher. It returns a new
// Selection object containing the matched elements.
func (s *Selection) NextMatcherUntilSelection(filter Matcher, sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NextFilteredUntilNodes is like NextUntilNodes, with the
// option to filter the results based on a selector string. It returns a new
// Selection object containing the matched elements.
func (s *Selection) NextFilteredUntilNodes(filterSelector string, nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// NextMatcherUntilNodes is like NextUntilNodes, with the
// option to filter the results based on a matcher. It returns a new
// Selection object containing the matched elements.
func (s *Selection) NextMatcherUntilNodes(filter Matcher, nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevFilteredUntil is like PrevUntil, with the option to filter
// the results based on a selector string.
// It returns a new Selection object containing the matched elements.
func (s *Selection) PrevFilteredUntil(filterSelector, untilSelector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevFilteredUntilMatcher is like PrevUntilMatcher, with the option to filter
// the results based on a matcher.
// It returns a new Selection object containing the matched elements.
func (s *Selection) PrevFilteredUntilMatcher(filter, until Matcher) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevFilteredUntilSelection is like PrevUntilSelection, with the
// option to filter the results based on a selector string. It returns a new
// Selection object containing the matched elements.
func (s *Selection) PrevFilteredUntilSelection(filterSelector string, sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevMatcherUntilSelection is like PrevUntilSelection, with the
// option to filter the results based on a matcher. It returns a new
// Selection object containing the matched elements.
func (s *Selection) PrevMatcherUntilSelection(filter Matcher, sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevFilteredUntilNodes is like PrevUntilNodes, with the
// option to filter the results based on a selector string. It returns a new
// Selection object containing the matched elements.
func (s *Selection) PrevFilteredUntilNodes(filterSelector string, nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrevMatcherUntilNodes is like PrevUntilNodes, with the
// option to filter the results based on a matcher. It returns a new
// Selection object containing the matched elements.
func (s *Selection) PrevMatcherUntilNodes(filter Matcher, nodes ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// Filter and push filters the nodes based on a matcher, and pushes the results
// on the stack, with the srcSel as previous selection.
func filterAndPush(srcSel *Selection, nodes []*html.Node, m Matcher) *Selection {
	_ = "STUB: not implemented"
	// Create a temporary Selection with the specified nodes to filter using winnow
	return nil
}

// Filter based on matcher and push on stack

// Internal implementation of Find that return raw nodes.
func findWithMatcher(nodes []*html.Node, m Matcher) []*html.Node {
	_ = "STUB: not implemented"
	// Map nodes to find the matches within the children of each node
	return nil
}

// Go down one level, becausejQuery's Find selects only within descendants

// Internal implementation to get all parent nodes, stopping at the specified
// node (or nil if no stop).
func getParentsNodes(nodes []*html.Node, stopm Matcher, stopNodes []*html.Node) []*html.Node {
	_ = "STUB: not implemented"
	return nil
}

// Internal implementation of sibling nodes that return a raw slice of matches.
func getSiblingNodes(nodes []*html.Node, st siblingType, untilm Matcher, untilNodes []*html.Node) []*html.Node {
	_ = "STUB: not implemented"
	return nil

	// If the requested siblings are ...Until, create the test function to
	// determine if the until condition is reached (returns true if it is)
}

// Matcher-based condition

// Nodes-based condition

// Gets the children nodes of each node in the specified slice of nodes,
// based on the sibling type request.
func getChildrenNodes(nodes []*html.Node, st siblingType) []*html.Node {
	_ = "STUB: not implemented"
	return nil
}

// Gets the children of the specified parent, based on the requested sibling
// type, skipping a specified node if required.
func getChildrenWithSiblingType(parent *html.Node, st siblingType, skipNode *html.Node,
	untilFunc func(*html.Node) bool) (result []*html.Node) {
	_ = "STUB: not implemented"

	// Create the iterator function
	return nil
}

// Based on the sibling type requested, iterate the right way

// First iteration, start with first child of parent
// Skip node if required

// Skip node if required

// Start with previous sibling of the skip node

// Start with next sibling of the skip node

// Not a valid node, try again from this one

// If this is an ...Until case, test before append (returns true
// if the until condition is reached)

// Only one node was requested (immediate next or previous), so exit

// Internal implementation of parent nodes that return a raw slice of Nodes.
func getParentNodes(nodes []*html.Node) []*html.Node { _ = "STUB: not implemented"; return nil }

// Internal map function used by many traversing methods. Takes the source nodes
// to iterate on and the mapping function that returns an array of nodes.
// Returns an array of nodes mapped by calling the callback function once for
// each node in the source nodes.
func mapNodes(nodes []*html.Node, f func(int, *html.Node) []*html.Node) (result []*html.Node) {
	_ = "STUB: not implemented"
	return nil
}
