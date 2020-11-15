package sample

import (
	"math"
	"reflect"
	"sort"
)

type Sample1Stream []Sample1

func Sample1StreamOf(arg ...Sample1) Sample1Stream {
	return arg
}
func Sample1StreamFrom(arg []Sample1) Sample1Stream {
	return arg
}
func CreateSample1Stream(arg ...Sample1) *Sample1Stream {
	tmp := Sample1StreamOf(arg...)
	return &tmp
}
func GenerateSample1Stream(arg []Sample1) *Sample1Stream {
	tmp := Sample1StreamFrom(arg)
	return &tmp
}

func (self *Sample1Stream) Add(arg Sample1) *Sample1Stream {
	return self.AddAll(arg)
}
func (self *Sample1Stream) AddAll(arg ...Sample1) *Sample1Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Sample1Stream) AddSafe(arg *Sample1) *Sample1Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *Sample1Stream) AllMatch(fn func(Sample1, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Sample1Stream) AnyMatch(fn func(Sample1, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Sample1Stream) Clone() *Sample1Stream {
	temp := make([]Sample1, self.Len())
	copy(temp, *self)
	return (*Sample1Stream)(&temp)
}
func (self *Sample1Stream) Copy() *Sample1Stream {
	return self.Clone()
}
func (self *Sample1Stream) Concat(arg []Sample1) *Sample1Stream {
	return self.AddAll(arg...)
}
func (self *Sample1Stream) Contains(arg Sample1) bool {
	return self.FindIndex(func(_arg Sample1, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *Sample1Stream) Clean() *Sample1Stream {
	*self = Sample1StreamOf()
	return self
}
func (self *Sample1Stream) Delete(index int) *Sample1Stream {
	return self.DeleteRange(index, index)
}
func (self *Sample1Stream) DeleteRange(startIndex, endIndex int) *Sample1Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Sample1Stream) Distinct() *Sample1Stream {
	caches := map[Sample1]bool{}
	result := Sample1StreamOf()
	for _, v := range *self {
		if f, ok := caches[v]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[v] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *Sample1Stream) Each(fn func(Sample1)) *Sample1Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *Sample1Stream) EachRight(fn func(Sample1)) *Sample1Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *Sample1Stream) Equals(arr []Sample1) bool {
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
func (self *Sample1Stream) Filter(fn func(Sample1, int) bool) *Sample1Stream {
	result := Sample1StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Sample1Stream) FilterSlim(fn func(Sample1, int) bool) *Sample1Stream {
	result := Sample1StreamOf()
	caches := map[Sample1]bool{}
	for i, v := range *self {
		if f, ok := caches[v]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[v] = fn(v, i); caches[v] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *Sample1Stream) Find(fn func(Sample1, int) bool) *Sample1 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *Sample1Stream) FindOr(fn func(Sample1, int) bool, or Sample1) Sample1 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *Sample1Stream) FindIndex(fn func(Sample1, int) bool) int {
	if self == nil {
		return -1
	}
	for i, v := range *self {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *Sample1Stream) First() *Sample1 {
	return self.Get(0)
}
func (self *Sample1Stream) FirstOr(arg Sample1) Sample1 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *Sample1Stream) ForEach(fn func(Sample1, int)) *Sample1Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Sample1Stream) ForEachRight(fn func(Sample1, int)) *Sample1Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Sample1Stream) GroupBy(fn func(Sample1, int) string) map[string][]Sample1 {
	m := map[string][]Sample1{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Sample1Stream) GroupByValues(fn func(Sample1, int) string) [][]Sample1 {
	var tmp [][]Sample1
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Sample1Stream) IndexOf(arg Sample1) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Sample1Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Sample1Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Sample1Stream) Last() *Sample1 {
	return self.Get(self.Len() - 1)
}
func (self *Sample1Stream) LastOr(arg Sample1) Sample1 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *Sample1Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Sample1Stream) Limit(limit int) *Sample1Stream {
	self.Slice(0, limit)
	return self
}
func (self *Sample1Stream) Map(fn func(Sample1, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Map2Int(fn func(Sample1, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Map2Int32(fn func(Sample1, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Map2Int64(fn func(Sample1, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Map2Float32(fn func(Sample1, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Map2Float64(fn func(Sample1, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Map2Bool(fn func(Sample1, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Map2Bytes(fn func(Sample1, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Map2String(fn func(Sample1, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample1Stream) Max(fn func(Sample1, int) float64) *Sample1 {
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
func (self *Sample1Stream) Min(fn func(Sample1, int) float64) *Sample1 {
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
func (self *Sample1Stream) NoneMatch(fn func(Sample1, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Sample1Stream) Get(index int) *Sample1 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Sample1Stream) GetOr(index int, arg Sample1) Sample1 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *Sample1Stream) Peek(fn func(*Sample1, int)) *Sample1Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Sample1Stream) Reduce(fn func(Sample1, Sample1, int) Sample1) *Sample1Stream {
	return self.ReduceInit(fn, Sample1{})
}
func (self *Sample1Stream) ReduceInit(fn func(Sample1, Sample1, int) Sample1, initialValue Sample1) *Sample1Stream {
	result := Sample1StreamOf()
	self.ForEach(func(v Sample1, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Sample1Stream) ReduceInterface(fn func(interface{}, Sample1, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Sample1{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample1Stream) ReduceString(fn func(string, Sample1, int) string) []string {
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
func (self *Sample1Stream) ReduceInt(fn func(int, Sample1, int) int) []int {
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
func (self *Sample1Stream) ReduceInt32(fn func(int32, Sample1, int) int32) []int32 {
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
func (self *Sample1Stream) ReduceInt64(fn func(int64, Sample1, int) int64) []int64 {
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
func (self *Sample1Stream) ReduceFloat32(fn func(float32, Sample1, int) float32) []float32 {
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
func (self *Sample1Stream) ReduceFloat64(fn func(float64, Sample1, int) float64) []float64 {
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
func (self *Sample1Stream) ReduceBool(fn func(bool, Sample1, int) bool) []bool {
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
func (self *Sample1Stream) Reverse() *Sample1Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Sample1Stream) Replace(fn func(Sample1, int) Sample1) *Sample1Stream {
	return self.ForEach(func(v Sample1, i int) { self.Set(i, fn(v, i)) })
}
func (self *Sample1Stream) Set(index int, val Sample1) *Sample1Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *Sample1Stream) Skip(skip int) *Sample1Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *Sample1Stream) SkippingEach(fn func(Sample1, int) int) *Sample1Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Sample1Stream) Slice(startIndex, n int) *Sample1Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Sample1{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *Sample1Stream) Sort(fn func(i, j int) bool) *Sample1Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Sample1Stream) Tail() *Sample1 {
	return self.Last()
}
func (self *Sample1Stream) TailOr(arg Sample1) Sample1 {
	return self.LastOr(arg)
}
func (self *Sample1Stream) ToList() []Sample1 {
	return self.Val()
}
func (self *Sample1Stream) Unique() *Sample1Stream {
	return self.Distinct()
}
func (self *Sample1Stream) Val() []Sample1 {
	if self == nil {
		return []Sample1{}
	}
	return *self.Copy()
}
func (self *Sample1Stream) While(fn func(Sample1, int) bool) *Sample1Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
