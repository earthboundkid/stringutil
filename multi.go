package stringutil

type MultiObject = map[string][]string

func ToMultiObject(obj Object) MultiObject {
	m := make(MultiObject, len(obj))
	for k, v := range obj {
		m[k] = []string{v}
	}
	return m
}

func Merge(dst, a, b MultiObject) MultiObject {
	if dst == nil {
		dst = make(MultiObject, len(a)+len(b))
	}
	for k, v := range a {
		dst[k] = v
	}
	for k, v := range b {
		if v2, ok := dst[k]; ok && len(v) != 0 && len(v2) != 0 {
			if !EqSlice(v, v2) {
				v = append(v2[:len(v2):len(v2)], v...)
			}
		}
		dst[k] = v
	}
	return dst
}
