package homeworks

type MyHashSet struct {
	Obj []int
}

func (this MyHashSet) Add(value int) MyHashSet {
	this.Obj = append(this.Obj, value)
	return this
}

func (this MyHashSet) Contains(value int) bool {

	for _, v := range this.Obj {
		if v == value {
			return true
		}
	}
	return false
}

func (this MyHashSet) Remove(index int) MyHashSet {
	newArr := []int{}
	for i, _ := range this.Obj {
		if i != index {
			newArr = append(newArr, this.Obj[i])
		}
	}
	this.Obj = newArr
	return this
}
