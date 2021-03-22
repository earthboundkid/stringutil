package stringutil

type Object = map[string]string

func Clear(obj Object) {
	for k, v := range obj {
		if v == "" {
			delete(obj, k)
		}
	}
}

func ToObject(m MultiObject) Object {
	o := make(Object, len(m))
	for k, v := range m {
		o[k] = Last(v...)
	}
	return o
}
