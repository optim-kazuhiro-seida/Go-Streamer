package sample

import (
	"math"
	"reflect"
	"sort"
)
type Sample2Stream []Sample2
func Sample2StreamOf(arg ...Sample2) Sample2Stream {
	return arg
}
func Sample2StreamFrom(arg []Sample2) Sample2Stream {
	return arg
}
func CreateSample2Stream(arg ...Sample2) *Sample2Stream {
    tmp := Sample2StreamOf(arg...)
    return &tmp
}
func GenerateSample2Stream(arg []Sample2) *Sample2Stream {
    tmp := Sample2StreamFrom(arg)
    return &tmp
}

func (self *Sample2Stream) Add(arg Sample2) *Sample2Stream {
	return self.AddAll(arg)
}
func (self *Sample2Stream) AddAll(arg ...Sample2) *Sample2Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Sample2Stream) AddSafe(arg *Sample2) *Sample2Stream {
    if arg != nil {
        self.Add(*arg)
    }
	return self
}
func (self *Sample2Stream) AllMatch(fn func(Sample2, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Sample2Stream) AnyMatch(fn func(Sample2, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Sample2Stream) Clone() *Sample2Stream {
	temp := make([]Sample2, self.Len())
	copy(temp, *self)
	return (*Sample2Stream)(&temp)
}
func (self *Sample2Stream) Copy() *Sample2Stream {
	return self.Clone()
}
func (self *Sample2Stream) Concat(arg []Sample2) *Sample2Stream {
	return self.AddAll(arg...)
}
func (self *Sample2Stream) Contains(arg Sample2) bool {
	return self.FindIndex(func(_arg Sample2, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *Sample2Stream) Clean() *Sample2Stream {
    return CreateSample2Stream()
}
func (self *Sample2Stream) Delete(index int) *Sample2Stream {
	return self.DeleteRange(index, index)
}
func (self *Sample2Stream) DeleteRange(startIndex, endIndex int) *Sample2Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Sample2Stream) Distinct() *Sample2Stream {
	stack := Sample2StreamOf()
	return self.Filter(func(arg Sample2, _ int) bool {
		if !stack.Contains(arg) {
			stack.Add(arg)
			return true
		}
		return false
	})
}
func (self *Sample2Stream) Equals(arr []Sample2) bool {
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
func (self *Sample2Stream) Filter(fn func(Sample2, int) bool) *Sample2Stream {
	_array := Sample2StreamOf()
	self.ForEach(func(v Sample2, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}
func (self *Sample2Stream) Find(fn func(Sample2, int) bool) *Sample2 {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}
func (self *Sample2Stream) FindIndex(fn func(Sample2, int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *Sample2Stream) First() *Sample2 {
	return self.Get(0)
}
func (self *Sample2Stream) ForEach(fn func(Sample2, int)) *Sample2Stream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *Sample2Stream) ForEachRight(fn func(Sample2, int)) *Sample2Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Sample2Stream) GroupBy(fn func(Sample2, int) string) map[string][]Sample2 {
    m := map[string][]Sample2{}
    for i, v := range self.Val() {
        key := fn(v, i)
        m[key] = append(m[key], v)
    }
    return m
}
func (self *Sample2Stream) GroupByValues(fn func(Sample2, int) string) [][]Sample2 {
	tmp := [][]Sample2{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Sample2Stream) IndexOf(arg Sample2) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Sample2Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Sample2Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Sample2Stream) Last() *Sample2 {
	return self.Get(self.Len() - 1)
}
func (self *Sample2Stream) Len() int {
    if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Sample2Stream) Limit(limit int) *Sample2Stream {
	self.Slice(0, limit)
	return self
}
func (self *Sample2Stream) Map(fn func(Sample2, int) Sample2) *Sample2Stream {
	return self.ForEach(func(v Sample2, i int) { self.Set(i, fn(v, i)) })
}
func (self *Sample2Stream) MapAny(fn func(Sample2, int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Map2Int(fn func(Sample2, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Map2Int32(fn func(Sample2, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Map2Int64(fn func(Sample2, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Map2Float32(fn func(Sample2, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Map2Float64(fn func(Sample2, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Map2Bool(fn func(Sample2, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Map2Bytes(fn func(Sample2, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Map2String(fn func(Sample2, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample2Stream) Max(fn func(Sample2, int) float64) *Sample2 {
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
func (self *Sample2Stream) Min(fn func(Sample2, int) float64) *Sample2 {
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
func (self *Sample2Stream) NoneMatch(fn func(Sample2, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Sample2Stream) Get(index int) *Sample2 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
        return &tmp
	}
	return nil
}
func (self *Sample2Stream) Peek(fn func(*Sample2, int)) *Sample2Stream {
    for i, v := range *self {
        fn(&v, i)
        self.Set(i, v)
    }
    return self
}
func (self *Sample2Stream) Reduce(fn func(Sample2, Sample2, int) Sample2) *Sample2Stream {
	return self.ReduceInit(fn, Sample2{})
}
func (self *Sample2Stream) ReduceInit(fn func(Sample2, Sample2, int) Sample2, initialValue Sample2) *Sample2Stream {
	result :=Sample2StreamOf()
	self.ForEach(func(v Sample2, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Sample2Stream) ReduceInterface(fn func(interface{}, Sample2, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Sample2{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample2Stream) ReduceString(fn func(string, Sample2, int) string) []string {
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
func (self *Sample2Stream) ReduceInt(fn func(int, Sample2, int) int) []int {
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
func (self *Sample2Stream) ReduceInt32(fn func(int32, Sample2, int) int32) []int32 {
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
func (self *Sample2Stream) ReduceInt64(fn func(int64, Sample2, int) int64) []int64 {
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
func (self *Sample2Stream) ReduceFloat32(fn func(float32, Sample2, int) float32) []float32 {
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
func (self *Sample2Stream) ReduceFloat64(fn func(float64, Sample2, int) float64) []float64 {
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
func (self *Sample2Stream) ReduceBool(fn func(bool, Sample2, int) bool) []bool {
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
func (self *Sample2Stream) Reverse() *Sample2Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Sample2Stream) Replace(fn func(Sample2, int) Sample2) *Sample2Stream {
	return self.Map(fn)
}
func (self *Sample2Stream) Set(index int, val Sample2) *Sample2Stream {
    if len(*self) > index {
        (*self)[index] = val
    }
    return self
}
func (self *Sample2Stream) Skip(skip int) *Sample2Stream {
	self.Slice(skip, self.Len()-skip)
	return self
}
func (self *Sample2Stream) SkippingEach(fn func(Sample2, int) int) *Sample2Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Sample2Stream) Slice(startIndex, n int) *Sample2Stream {
    last := startIndex+n
    if len(*self)-1 < startIndex {
        *self = []Sample2{}
    } else if len(*self) < last {
        *self = (*self)[startIndex:len(*self)]
    } else {
        *self = (*self)[startIndex:last]
    }
	return self
}
func (self *Sample2Stream) Sort(fn func(i, j int) bool) *Sample2Stream {
	sort.Slice(*self, fn)
	return self
}
func (self *Sample2Stream) SortStable(fn func(i, j int) bool) *Sample2Stream {
	sort.SliceStable(*self, fn)
	return self
}
func (self *Sample2Stream) ToList() []Sample2 {
	return self.Val()
}
func (self *Sample2Stream) Unique() *Sample2Stream {
	return self.Distinct()
}
func (self *Sample2Stream) Val() []Sample2 {
	if self == nil {
		return []Sample2{}
	}
	return *self.Copy()
}
func (self *Sample2Stream) While(fn func(Sample2, int) bool) *Sample2Stream {
    for i, v := range self.Val() {
        if !fn(v, i) {
            break
        }
    }
    return self
}
