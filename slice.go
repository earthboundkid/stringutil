package stringutil

import "sort"

// Sort is a convenience alias so you don't have to import package sort.
var Sort = sort.Strings

type Slice = []string

func Copy(ss Slice) Slice {
	return append(Slice(nil), ss...)
}

func prealloc(dst Slice, size int) Slice {
	ss := dst
	if avail := cap(dst) - len(dst); avail < size {
		ss = append(dst, make(Slice, size)...)[:len(dst)]
	}
	return ss
}

func ToSlice(m Set) Slice {
	return Append(nil, m)
}

func Append(dst Slice, m Set) Slice {
	ss := prealloc(dst, len(m))
	for s, ok := range m {
		if ok {
			ss = append(ss, s)
		}
	}
	return ss
}

func UniqueVia(dst, ss Slice, m Set) Slice {
	m = Update(m, ss)
	r := prealloc(dst, len(m))
	for _, s := range ss {
		if m[s] {
			r = append(r, s)
			m[s] = false
		}
	}
	return r
}

func Unique(dst, ss Slice) Slice {
	return UniqueVia(dst, ss, nil)
}

func SymDiff(oldss, newss Slice) (added, removed Slice) {
	oldSet := ToSet(oldss)
	newSet := ToSet(newss)
	addedSet, removedSet := SymDiffSets(oldSet, newSet)
	added = ToSlice(addedSet)
	removed = ToSlice(removedSet)
	return
}

func EqSlice(a, b Slice) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
