package maps

type Map map[interface{}]interface{}

func (this Map) Keys() []interface{} {
	ret := make([]string, 0, len(this))
	for k := range this {
		ret = append(ret, k)
	}
	return ret
}
