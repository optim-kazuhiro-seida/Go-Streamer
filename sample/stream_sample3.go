package sample

import (
	"math"
	"reflect"
	"sort"
)
type Sample3Stream []Sample3
func Sample3StreamOf(arg ...Sample3) Sample3Stream {
	return arg
}
func Sample3StreamFrom(arg []Sample3) Sample3Stream {
	return arg
}
func CreateSample3Stream(arg ...Sample3) *Sample3Stream {
    tmp := Sample3StreamOf(arg...)
    return &tmp
}
func GenerateSample3Stream(arg []Sample3) *Sample3Stream {
    tmp := Sample3StreamFrom(arg)
    return &tmp
}

func (self *Sample3Stream) Add(arg Sample3) *Sample3Stream {
	return self.AddAll(arg)
}
func (self *Sample3Stream) AddAll(arg ...Sample3) *Sample3Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Sample3Stream) AddSafe(arg *Sample3) *Sample3Stream {
    if arg != nil {
        self.Add(*arg)
    }
	return self
}
func (self *Sample3Stream) AllMatch(fn func(Sample3, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Sample3Stream) AnyMatch(fn func(Sample3, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Sample3Stream) Clone() *Sample3Stream {
	temp := make([]Sample3, self.Len())
	copy(temp, *self)
	return (*Sample3Stream)(&temp)
}
func (self *Sample3Stream) Copy() *Sample3Stream {
	return self.Clone()
}
func (self *Sample3Stream) Concat(arg []Sample3) *Sample3Stream {
	return self.AddAll(arg...)
}
func (self *Sample3Stream) Contains(arg Sample3) bool {
	return self.FindIndex(func(_arg Sample3, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *Sample3Stream) Clean() *Sample3Stream {
    return CreateSample3Stream()
}
func (self *Sample3Stream) Delete(index int) *Sample3Stream {
	return self.DeleteRange(index, index)
}
func (self *Sample3Stream) DeleteRange(startIndex, endIndex int) *Sample3Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Sample3Stream) Distinct() *Sample3Stream {
	stack := Sample3StreamOf()
	return self.Filter(func(arg Sample3, _ int) bool {
		if !stack.Contains(arg) {
			stack.Add(arg)
			return true
		}
		return false
	})
}
func (self *Sample3Stream) Equals(arr []Sample3) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if !reflect.DeepEqual((*self)[i], arr[i]) {
			return false
		}
	}
	return true
}
func (self *Sample3Stream) Filter(fn func(Sample3, int) bool) *Sample3Stream {
	_array := Sample3StreamOf()
	self.ForEach(func(v Sample3, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}
func (self *Sample3Stream) Find(fn func(Sample3, int) bool) *Sample3 {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}
func (self *Sample3Stream) FindIndex(fn func(Sample3, int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *Sample3Stream) First() *Sample3 {
	return self.Get(0)
}
func (self *Sample3Stream) ForEach(fn func(Sample3, int)) *Sample3Stream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *Sample3Stream) ForEachRight(fn func(Sample3, int)) *Sample3Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Sample3Stream) GroupBy(fn func(Sample3, int) string) map[string][]Sample3 {
    m := map[string][]Sample3{}
    for i, v := range self.Val() {
        key := fn(v, i)
        m[key] = append(m[key], v)
    }
    return m
}
func (self *Sample3Stream) GroupByValues(fn func(Sample3, int) string) [][]Sample3 {
	tmp := [][]Sample3{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Sample3Stream) IndexOf(arg Sample3) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Sample3Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Sample3Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Sample3Stream) Last() *Sample3 {
	return self.Get(self.Len() - 1)
}
func (self *Sample3Stream) Len() int {
    if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Sample3Stream) Limit(limit int) *Sample3Stream {
	self.Slice(0, limit)
	return self
}
func (self *Sample3Stream) Map(fn func(Sample3, int) Sample3) *Sample3Stream {
	return self.ForEach(func(v Sample3, i int) { self.Set(i, fn(v, i)) })
}
func (self *Sample3Stream) MapAny(fn func(Sample3, int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Map2Int(fn func(Sample3, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Map2Int32(fn func(Sample3, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Map2Int64(fn func(Sample3, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Map2Float32(fn func(Sample3, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Map2Float64(fn func(Sample3, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Map2Bool(fn func(Sample3, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Map2Bytes(fn func(Sample3, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Map2String(fn func(Sample3, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample3Stream) Max(fn func(Sample3, int) float64) *Sample3 {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Max(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *Sample3Stream) Min(fn func(Sample3, int) float64) *Sample3 {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Min(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *Sample3Stream) NoneMatch(fn func(Sample3, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Sample3Stream) Get(index int) *Sample3 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
        return &tmp
	}
	return nil
}
func (self *Sample3Stream) Peek(fn func(*Sample3, int)) *Sample3Stream {
    for i, v := range *self {
        fn(&v, i)
        self.Set(i, v)
    }
    return self
}
func (self *Sample3Stream) Reduce(fn func(Sample3, Sample3, int) Sample3) *Sample3Stream {
	return self.ReduceInit(fn, Sample3{})
}
func (self *Sample3Stream) ReduceInit(fn func(Sample3, Sample3, int) Sample3, initialValue Sample3) *Sample3Stream {
	result :=Sample3StreamOf()
	self.ForEach(func(v Sample3, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Sample3Stream) ReduceInterface(fn func(interface{}, Sample3, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Sample3{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample3Stream) ReduceString(fn func(string, Sample3, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample3Stream) ReduceInt(fn func(int, Sample3, int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample3Stream) ReduceInt32(fn func(int32, Sample3, int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample3Stream) ReduceInt64(fn func(int64, Sample3, int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample3Stream) ReduceFloat32(fn func(float32, Sample3, int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample3Stream) ReduceFloat64(fn func(float64, Sample3, int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample3Stream) ReduceBool(fn func(bool, Sample3, int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample3Stream) Reverse() *Sample3Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Sample3Stream) Replace(fn func(Sample3, int) Sample3) *Sample3Stream {
	return self.Map(fn)
}
func (self *Sample3Stream) Set(index int, val Sample3) *Sample3Stream {
    if len(*self) > index {
        (*self)[index] = val
    }
    return self
}
func (self *Sample3Stream) Skip(skip int) *Sample3Stream {
	self.Slice(skip, self.Len()-skip)
	return self
}
func (self *Sample3Stream) SkippingEach(fn func(Sample3, int) int) *Sample3Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Sample3Stream) Slice(startIndex, n int) *Sample3Stream {
    last := startIndex+n
    if len(*self)-1 < startIndex {
        *self = []Sample3{}
    } else if len(*self) < last {
        *self = (*self)[startIndex:len(*self)]
    } else {
        *self = (*self)[startIndex:last]
    }
	return self
}
func (self *Sample3Stream) Sort(fn func(i, j int) bool) *Sample3Stream {
	sort.Slice(*self, fn)
	return self
}
func (self *Sample3Stream) SortStable(fn func(i, j int) bool) *Sample3Stream {
	sort.SliceStable(*self, fn)
	return self
}
func (self *Sample3Stream) ToList() []Sample3 {
	return self.Val()
}
func (self *Sample3Stream) Unique() *Sample3Stream {
	return self.Distinct()
}
func (self *Sample3Stream) Val() []Sample3 {
	if self == nil {
		return []Sample3{}
	}
	return *self.Copy()
}
func (self *Sample3Stream) While(fn func(Sample3, int) bool) *Sample3Stream {
    for i, v := range self.Val() {
        if !fn(v, i) {
            break
        }
    }
    return self
}
