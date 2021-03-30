package stringutil

type Set = map[string]bool

func ToSet(ss Slice) Set {
	return Update(nil, ss)
}

func Update(m Set, ss Slice) Set {
	if m == nil {
		m = make(Set, len(ss))
	}
	for _, s := range ss {
		m[s] = true
	}
	return m
}

func Clone(m Set) Set {
	newm := make(Set, len(m))
	for k, v := range m {
		newm[k] = v
	}
	return newm
}

func Purge(m Set) {
	for k, v := range m {
		if !v {
			delete(m, k)
		}
	}
}

func Remove(m Set, ss Slice) {
	for _, s := range ss {
		delete(m, s)
	}
}

func Zero(m Set) {
	for k := range m {
		delete(m, k)
	}
}

func EqSet(a, b Set) bool {
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	// handle case of only one set having a false value
	for k := range b {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func DiffSets(oldSet, newSet Set) (added Set) {
	added = make(map[string]bool, len(newSet))
	for s := range newSet {
		if newSet[s] && !oldSet[s] {
			added[s] = true
		}
	}
	return
}

func SymDiffSets(oldSet, newSet Set) (added, removed Set) {
	added = DiffSets(oldSet, newSet)
	removed = DiffSets(newSet, oldSet)
	return
}
