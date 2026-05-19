package goquery

import (
	"regexp"

	"golang.org/x/net/html"
)

var rxClassTrim = regexp.MustCompile("[\t\r\n]")

// Attr gets the specified attribute's value for the first element in the
// Selection. To get the value for each element individually, use a looping
// construct such as Each or Map method.
func (s *Selection) Attr(attrName string) (val string, exists bool) {
	_ = "STUB: not implemented"
	return "", false
}

// AttrOr works like Attr but returns default value if attribute is not present.
func (s *Selection) AttrOr(attrName, defaultValue string) string {
	_ = "STUB: not implemented"
	return ""
}

// RemoveAttr removes the named attribute from each element in the set of matched elements.
func (s *Selection) RemoveAttr(attrName string) *Selection { _ = "STUB: not implemented"; return nil }

// SetAttr sets the given attribute on each element in the set of matched elements.
func (s *Selection) SetAttr(attrName, val string) *Selection { _ = "STUB: not implemented"; return nil }

// Text gets the combined text contents of each element in the set of matched
// elements, including their descendants.
func (s *Selection) Text() string { _ = "STUB: not implemented"; return "" }

// Slightly optimized vs calling Each: no single selection object created

// Keep newlines and spaces, like jQuery

// Size is an alias for Length.
func (s *Selection) Size() int {
	_ = "STUB: not implemented"

	// Length returns the number of elements in the Selection object.
	return 0
}

func (s *Selection) Length() int { _ = "STUB: not implemented"; return 0 }

// Html gets the HTML contents of the first element in the set of matched
// elements. It includes text and comment nodes.
func (s *Selection) Html() (ret string, e error) {
	_ = "STUB: not implemented"
	// Since there is no .innerHtml, the HTML content must be re-created from
	// the nodes using html.Render.
	return "", nil
}

// AddClass adds the given class(es) to each element in the set of matched elements.
// Multiple class names can be specified, separated by a space or via multiple arguments.
func (s *Selection) AddClass(class ...string) *Selection { _ = "STUB: not implemented"; return nil }

// HasClass determines whether any of the matched elements are assigned the
// given class.
func (s *Selection) HasClass(class string) bool { _ = "STUB: not implemented"; return false }

// RemoveClass removes the given class(es) from each element in the set of matched elements.
// Multiple class names can be specified, separated by a space or via multiple arguments.
// If no class name is provided, all classes are removed.
func (s *Selection) RemoveClass(class ...string) *Selection { _ = "STUB: not implemented"; return nil }

// ToggleClass adds or removes the given class(es) for each element in the set of matched elements.
// Multiple class names can be specified, separated by a space or via multiple arguments.
func (s *Selection) ToggleClass(class ...string) *Selection { _ = "STUB: not implemented"; return nil }

func getAttributePtr(attrName string, n *html.Node) *html.Attribute {
	_ = "STUB: not implemented"
	return nil
}

// Private function to get the specified attribute's value from a node.
func getAttributeValue(attrName string, n *html.Node) (val string, exists bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Get and normalize the "class" attribute from the node.
func getClassesAndAttr(n *html.Node, create bool) (classes string, attr *html.Attribute) {
	_ = "STUB: not implemented"
	// Applies only to element nodes
	return "", nil
}

func getClassesSlice(classes string) []string { _ = "STUB: not implemented"; return nil }

func removeAttr(n *html.Node, attrName string) { _ = "STUB: not implemented"; return }

func setClasses(n *html.Node, attr *html.Attribute, classes string) {
	_ = "STUB: not implemented"
	return
}
