package set

type Set map[interface{}]struct{}

func NewSet(arr ...interface{}) Set {
	s := make(Set, len(arr))
	for _, v := range arr {
		s[v] = struct{}{}
	}
	return s
}

func (this Set) Add(e interface{}) Set {
	this[e] = true
	return this
}

func (this Set) AddArr(arr []interface{}) Set {
	for _, e := range arr {
		this.Add(e)
	}
	return this
}

func (this Set) Elements() []interface{} {
	s := make([]interface{}, len(this))
	for e := range this {
		s = append(s, e)
	}
	return s
}

func (this Set) Remove(e interface{}) Set {
	delete(this, e)
	return this
}
func (this Set) RemoveArr(arr []interface{}) Set {
	for _, e := range arr {
		this.Remove(e)
	}
	return this
}

func (this Set) Has(e interface{}) bool {
	_, ok := this[e]
	return ok
}
