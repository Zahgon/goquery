package goquery

import (
	"io"

	"golang.org/x/net/html"
)

// used to determine if a set (map[*html.Node]bool) should be used
// instead of iterating over a slice. The set uses more memory and
// is slower than slice iteration for small N.
const minNodesForSet = 1000

var nodeNames = []string{
	html.ErrorNode:    "#error",
	html.TextNode:     "#text",
	html.DocumentNode: "#document",
	html.CommentNode:  "#comment",
}

// NodeName returns the node name of the first element in the selection.
// It tries to behave in a similar way as the DOM's nodeName property
// (https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeName).
//
// Go's net/html package defines the following node types, listed with
// the corresponding returned value from this function:
//
//	ErrorNode : #error
//	TextNode : #text
//	DocumentNode : #document
//	ElementNode : the element's tag name
//	CommentNode : #comment
//	DoctypeNode : the name of the document type
func NodeName(s *Selection) string { _ = "STUB: not implemented"; return "" }

// nodeName returns the node name of the given html node.
// See NodeName for additional details on behaviour.
func nodeName(node *html.Node) string { _ = "STUB: not implemented"; return "" }

// Render renders the HTML of the first item in the selection and writes it to
// the writer. It behaves the same as OuterHtml but writes to w instead of
// returning the string.
func Render(w io.Writer, s *Selection) error { _ = "STUB: not implemented"; return nil }

// OuterHtml returns the outer HTML rendering of the first item in
// the selection - that is, the HTML including the first element's
// tag and attributes.
//
// Unlike Html, this is a function and not a method on the Selection,
// because this is not a jQuery method (in javascript-land, this is
// a property provided by the DOM).
func OuterHtml(s *Selection) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Loop through all container nodes to search for the target node.
func sliceContains(container []*html.Node, contained *html.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// Checks if the contained node is within the container node.
func nodeContains(container *html.Node, contained *html.Node) bool {
	_ = "STUB: not implemented"
	// Check if the parent of the contained node is the container node, traversing
	// upward until the top is reached, or the container is found.
	return false
}

// Checks if the target node is in the slice of nodes.
func isInSlice(slice []*html.Node, node *html.Node) bool { _ = "STUB: not implemented"; return false }

// Returns the index of the target node in the slice, or -1.
func indexInSlice(slice []*html.Node, node *html.Node) int { _ = "STUB: not implemented"; return 0 }

// Appends the new nodes to the target slice, making sure no duplicate is added.
// There is no check to the original state of the target slice, so it may still
// contain duplicates. The target slice is returned because append() may create
// a new underlying array. If targetSet is nil, a local set is created with the
// target if len(target) + len(nodes) is greater than minNodesForSet.
func appendWithoutDuplicates(target []*html.Node, nodes []*html.Node, targetSet map[*html.Node]bool) []*html.Node {
	_ = "STUB: not implemented"
	// if there are not that many nodes, don't use the map, faster to just use nested loops
	// (unless a non-nil targetSet is passed, in which case the caller knows better).
	return nil
}

// if a targetSet is passed, then assume it is reliable, otherwise create one
// and initialize it with the current target contents.

// Loop through a selection, returning only those nodes that pass the predicate
// function.
func grep(sel *Selection, predicate func(i int, s *Selection) bool) (result []*html.Node) {
	_ = "STUB: not implemented"
	return nil
}

// Creates a new Selection object based on the specified nodes, and keeps the
// source Selection object on the stack (linked list).
func pushStack(fromSel *Selection, nodes []*html.Node) *Selection {
	_ = "STUB: not implemented"
	return nil
}
