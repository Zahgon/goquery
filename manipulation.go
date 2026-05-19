package goquery

import (
	"golang.org/x/net/html"
)

// After applies the selector from the root document and inserts the matched elements
// after the elements in the set of matched elements.
//
// If one of the matched elements in the selection is not currently in the
// document, it's impossible to insert nodes after it, so it will be ignored.
//
// This follows the same rules as Selection.Append.
func (s *Selection) After(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// AfterMatcher applies the matcher from the root document and inserts the matched elements
// after the elements in the set of matched elements.
//
// If one of the matched elements in the selection is not currently in the
// document, it's impossible to insert nodes after it, so it will be ignored.
//
// This follows the same rules as Selection.Append.
func (s *Selection) AfterMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// AfterSelection inserts the elements in the selection after each element in the set of matched
// elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) AfterSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// AfterHtml parses the html and inserts it after the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) AfterHtml(htmlStr string) *Selection { _ = "STUB: not implemented"; return nil }

// AfterNodes inserts the nodes after each element in the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) AfterNodes(ns ...*html.Node) *Selection { _ = "STUB: not implemented"; return nil }

// Append appends the elements specified by the selector to the end of each element
// in the set of matched elements, following those rules:
//
// 1) The selector is applied to the root document.
//
// 2) Elements that are part of the document will be moved to the new location.
//
// 3) If there are multiple locations to append to, cloned nodes will be
// appended to all target locations except the last one, which will be moved
// as noted in (2).
func (s *Selection) Append(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// AppendMatcher appends the elements specified by the matcher to the end of each element
// in the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) AppendMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// AppendSelection appends the elements in the selection to the end of each element
// in the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) AppendSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// AppendHtml parses the html and appends it to the set of matched elements.
func (s *Selection) AppendHtml(htmlStr string) *Selection { _ = "STUB: not implemented"; return nil }

// AppendNodes appends the specified nodes to each node in the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) AppendNodes(ns ...*html.Node) *Selection { _ = "STUB: not implemented"; return nil }

// Before inserts the matched elements before each element in the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) Before(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// BeforeMatcher inserts the matched elements before each element in the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) BeforeMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// BeforeSelection inserts the elements in the selection before each element in the set of matched
// elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) BeforeSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// BeforeHtml parses the html and inserts it before the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) BeforeHtml(htmlStr string) *Selection { _ = "STUB: not implemented"; return nil }

// BeforeNodes inserts the nodes before each element in the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) BeforeNodes(ns ...*html.Node) *Selection { _ = "STUB: not implemented"; return nil }

// Clone creates a deep copy of the set of matched nodes. The new nodes will not be
// attached to the document.
func (s *Selection) Clone() *Selection { _ = "STUB: not implemented"; return nil }

// Empty removes all children nodes from the set of matched elements.
// It returns the children nodes in a new Selection.
func (s *Selection) Empty() *Selection { _ = "STUB: not implemented"; return nil }

// Prepend prepends the elements specified by the selector to each element in
// the set of matched elements, following the same rules as Append.
func (s *Selection) Prepend(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// PrependMatcher prepends the elements specified by the matcher to each
// element in the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) PrependMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// PrependSelection prepends the elements in the selection to each element in
// the set of matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) PrependSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// PrependHtml parses the html and prepends it to the set of matched elements.
func (s *Selection) PrependHtml(htmlStr string) *Selection { _ = "STUB: not implemented"; return nil }

// PrependNodes prepends the specified nodes to each node in the set of
// matched elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) PrependNodes(ns ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// sn.FirstChild may be nil, in which case this functions like
// sn.AppendChild()

// Remove removes the set of matched elements from the document.
// It returns the same selection, now consisting of nodes not in the document.
func (s *Selection) Remove() *Selection { _ = "STUB: not implemented"; return nil }

// RemoveFiltered removes from the current set of matched elements those that
// match the selector filter. It returns the Selection of removed nodes.
//
// For example if the selection s contains "<h1>", "<h2>" and "<h3>"
// and s.RemoveFiltered("h2") is called, only the "<h2>" node is removed
// (and returned), while "<h1>" and "<h3>" are kept in the document.
func (s *Selection) RemoveFiltered(selector string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// RemoveMatcher removes from the current set of matched elements those that
// match the Matcher filter. It returns the Selection of removed nodes.
// See RemoveFiltered for additional information.
func (s *Selection) RemoveMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// ReplaceWith replaces each element in the set of matched elements with the
// nodes matched by the given selector.
// It returns the removed elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) ReplaceWith(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// ReplaceWithMatcher replaces each element in the set of matched elements with
// the nodes matched by the given Matcher.
// It returns the removed elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) ReplaceWithMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// ReplaceWithSelection replaces each element in the set of matched elements with
// the nodes from the given Selection.
// It returns the removed elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) ReplaceWithSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceWithHtml replaces each element in the set of matched elements with
// the parsed HTML.
// It returns the removed elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) ReplaceWithHtml(htmlStr string) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceWithNodes replaces each element in the set of matched elements with
// the given nodes.
// It returns the removed elements.
//
// This follows the same rules as Selection.Append.
func (s *Selection) ReplaceWithNodes(ns ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// SetHtml sets the html content of each element in the selection to
// specified html string.
func (s *Selection) SetHtml(htmlStr string) *Selection { _ = "STUB: not implemented"; return nil }

// SetText sets the content of each element in the selection to specified content.
// The provided text string is escaped.
func (s *Selection) SetText(text string) *Selection { _ = "STUB: not implemented"; return nil }

// Unwrap removes the parents of the set of matched elements, leaving the matched
// elements (and their siblings, if any) in their place.
// It returns the original selection.
func (s *Selection) Unwrap() *Selection { _ = "STUB: not implemented"; return nil }

// For some reason, jquery allows unwrap to remove the <head> element, so
// allowing it here too. Same for <html>. Why it allows those elements to
// be unwrapped while not allowing body is a mystery to me.

// Wrap wraps each element in the set of matched elements inside the first
// element matched by the given selector. The matched child is cloned before
// being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) Wrap(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// WrapMatcher wraps each element in the set of matched elements inside the
// first element matched by the given matcher. The matched child is cloned
// before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// WrapSelection wraps each element in the set of matched elements inside the
// first element in the given Selection. The element is cloned before being
// inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapSelection(sel *Selection) *Selection { _ = "STUB: not implemented"; return nil }

// WrapHtml wraps each element in the set of matched elements inside the inner-
// most child of the given HTML.
//
// It returns the original set of elements.
func (s *Selection) WrapHtml(htmlStr string) *Selection { _ = "STUB: not implemented"; return nil }

// WrapNode wraps each element in the set of matched elements inside the inner-
// most child of the given node. The given node is copied before being inserted
// into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapNode(n *html.Node) *Selection { _ = "STUB: not implemented"; return nil }

func (s *Selection) wrapNodes(ns ...*html.Node) *Selection { _ = "STUB: not implemented"; return nil }

// WrapAll wraps a single HTML structure, matched by the given selector, around
// all elements in the set of matched elements. The matched child is cloned
// before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapAll(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// WrapAllMatcher wraps a single HTML structure, matched by the given Matcher,
// around all elements in the set of matched elements. The matched child is
// cloned before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapAllMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// WrapAllSelection wraps a single HTML structure, the first node of the given
// Selection, around all elements in the set of matched elements. The matched
// child is cloned before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapAllSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// WrapAllHtml wraps the given HTML structure around all elements in the set of
// matched elements. The matched child is cloned before being inserted into the
// document.
//
// It returns the original set of elements.
func (s *Selection) WrapAllHtml(htmlStr string) *Selection { _ = "STUB: not implemented"; return nil }

func (s *Selection) wrapAllNodes(ns ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// WrapAllNode wraps the given node around the first element in the Selection,
// making all other nodes in the Selection children of the given node. The node
// is cloned before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapAllNode(n *html.Node) *Selection { _ = "STUB: not implemented"; return nil }

// WrapInner wraps an HTML structure, matched by the given selector, around the
// content of element in the set of matched elements. The matched child is
// cloned before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapInner(selector string) *Selection { _ = "STUB: not implemented"; return nil }

// WrapInnerMatcher wraps an HTML structure, matched by the given selector,
// around the content of element in the set of matched elements. The matched
// child is cloned before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapInnerMatcher(m Matcher) *Selection { _ = "STUB: not implemented"; return nil }

// WrapInnerSelection wraps an HTML structure, matched by the given selector,
// around the content of element in the set of matched elements. The matched
// child is cloned before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapInnerSelection(sel *Selection) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// WrapInnerHtml wraps an HTML structure, matched by the given selector, around
// the content of element in the set of matched elements. The matched child is
// cloned before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapInnerHtml(htmlStr string) *Selection { _ = "STUB: not implemented"; return nil }

// WrapInnerNode wraps an HTML structure, matched by the given selector, around
// the content of element in the set of matched elements. The matched child is
// cloned before being inserted into the document.
//
// It returns the original set of elements.
func (s *Selection) WrapInnerNode(n *html.Node) *Selection { _ = "STUB: not implemented"; return nil }

func (s *Selection) wrapInnerNodes(ns ...*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}

func parseHtml(h string) []*html.Node {
	_ = "STUB: not implemented"
	// Errors are only returned when the io.Reader returns any error besides
	// EOF, but strings.Reader never will
	return nil
}

func parseHtmlWithContext(h string, context *html.Node) []*html.Node {
	_ = "STUB: not implemented"
	// Errors are only returned when the io.Reader returns any error besides
	// EOF, but strings.Reader never will
	return nil
}

// Get the first child that is an ElementNode
func getFirstChildEl(n *html.Node) *html.Node { _ = "STUB: not implemented"; return nil }

// Deep copy a slice of nodes.
func cloneNodes(ns []*html.Node) []*html.Node { _ = "STUB: not implemented"; return nil }

// Deep copy a node. The new node has clones of all the original node's
// children but none of its parents or siblings.
func cloneNode(n *html.Node) *html.Node { _ = "STUB: not implemented"; return nil }

func (s *Selection) manipulateNodes(ns []*html.Node, reverse bool,
	f func(sn *html.Node, n *html.Node)) *Selection {
	_ = "STUB: not implemented"
	return nil

	// net.Html doesn't provide document fragments for insertion, so to get
	// things in the correct order with After() and Prepend(), the callback
	// needs to be called on the reverse of the nodes.
}

// eachNodeHtml parses the given html string and inserts the resulting nodes in the dom with the mergeFn.
// The parsed nodes are inserted for each element of the selection.
// isParent can be used to indicate that the elements of the selection should be treated as the parent for the parsed html.
// A cache is used to avoid parsing the html multiple times should the elements of the selection result in the same context.
func (s *Selection) eachNodeHtml(htmlStr string, isParent bool, mergeFn func(n *html.Node, nodes []*html.Node)) *Selection {
	_ = "STUB: not implemented"
	// cache to avoid parsing the html for the same context multiple times
	return nil
}
