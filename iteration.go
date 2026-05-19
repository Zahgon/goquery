package goquery

import "iter"

// Each iterates over a Selection object, executing a function for each
// matched element. It returns the current Selection object. The function
// f is called for each element in the selection with the index of the
// element in that selection starting at 0, and a *Selection that contains
// only that element.
func (s *Selection) Each(f func(int, *Selection)) *Selection { _ = "STUB: not implemented"; return nil }

// EachIter returns an iterator that yields the Selection object in order.
// The implementation is similar to Each, but it returns an iterator instead.
func (s *Selection) EachIter() iter.Seq2[int, *Selection] { _ = "STUB: not implemented"; return nil }

// EachWithBreak iterates over a Selection object, executing a function for each
// matched element. It is identical to Each except that it is possible to break
// out of the loop by returning false in the callback function. It returns the
// current Selection object.
func (s *Selection) EachWithBreak(f func(int, *Selection) bool) *Selection {
	_ = "STUB: not implemented"
	return nil
}

// Map passes each element in the current matched set through a function,
// producing a slice of string holding the returned values. The function
// f is called for each element in the selection with the index of the
// element in that selection starting at 0, and a *Selection that contains
// only that element.
func (s *Selection) Map(f func(int, *Selection) string) (result []string) {
	_ = "STUB: not implemented"

	// Map is the generic version of Selection.Map, allowing any type to be
	// returned.
	return nil
}

func Map[E any](s *Selection, f func(int, *Selection) E) (result []E) {
	_ = "STUB: not implemented"
	return nil
}
