package sample

import (
	"math"
	"reflect"
	"sort"
)

type Sample5Stream []Sample5

func Sample5StreamOf(arg ...Sample5) Sample5Stream {
	return arg
}
func Sample5StreamFrom(arg []Sample5) Sample5Stream {
	return arg
}
func CreateSample5Stream(arg ...Sample5) *Sample5Stream {
	tmp := Sample5StreamOf(arg...)
	return &tmp
}
func GenerateSample5Stream(arg []Sample5) *Sample5Stream {
	tmp := Sample5StreamFrom(arg)
	return &tmp
}

func (self *Sample5Stream) Add(arg Sample5) *Sample5Stream {
	return self.AddAll(arg)
}
func (self *Sample5Stream) AddAll(arg ...Sample5) *Sample5Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Sample5Stream) AddSafe(arg *Sample5) *Sample5Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *Sample5Stream) AllMatch(fn func(Sample5, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Sample5Stream) AnyMatch(fn func(Sample5, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Sample5Stream) Clone() *Sample5Stream {
	temp := make([]Sample5, self.Len())
	copy(temp, *self)
	return (*Sample5Stream)(&temp)
}
func (self *Sample5Stream) Copy() *Sample5Stream {
	return self.Clone()
}
func (self *Sample5Stream) Concat(arg []Sample5) *Sample5Stream {
	return self.AddAll(arg...)
}
func (self *Sample5Stream) Contains(arg Sample5) bool {
	return self.FindIndex(func(_arg Sample5, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *Sample5Stream) Clean() *Sample5Stream {
	*self = Sample5StreamOf()
	return self
}
func (self *Sample5Stream) Delete(index int) *Sample5Stream {
	return self.DeleteRange(index, index)
}
func (self *Sample5Stream) DeleteRange(startIndex, endIndex int) *Sample5Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Sample5Stream) Distinct() *Sample5Stream {
	caches := map[Sample5]bool{}
	result := Sample5StreamOf()
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
func (self *Sample5Stream) Each(fn func(Sample5)) *Sample5Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *Sample5Stream) EachRight(fn func(Sample5)) *Sample5Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *Sample5Stream) Equals(arr []Sample5) bool {
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
func (self *Sample5Stream) Filter(fn func(Sample5, int) bool) *Sample5Stream {
	result := Sample5StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Sample5Stream) FilterSlim(fn func(Sample5, int) bool) *Sample5Stream {
	result := Sample5StreamOf()
	caches := map[Sample5]bool{}
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
func (self *Sample5Stream) Find(fn func(Sample5, int) bool) *Sample5 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *Sample5Stream) FindOr(fn func(Sample5, int) bool, or Sample5) Sample5 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *Sample5Stream) FindIndex(fn func(Sample5, int) bool) int {
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
func (self *Sample5Stream) First() *Sample5 {
	return self.Get(0)
}
func (self *Sample5Stream) FirstOr(arg Sample5) Sample5 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *Sample5Stream) ForEach(fn func(Sample5, int)) *Sample5Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Sample5Stream) ForEachRight(fn func(Sample5, int)) *Sample5Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Sample5Stream) GroupBy(fn func(Sample5, int) string) map[string][]Sample5 {
	m := map[string][]Sample5{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Sample5Stream) GroupByValues(fn func(Sample5, int) string) [][]Sample5 {
	var tmp [][]Sample5
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Sample5Stream) IndexOf(arg Sample5) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Sample5Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Sample5Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Sample5Stream) Last() *Sample5 {
	return self.Get(self.Len() - 1)
}
func (self *Sample5Stream) LastOr(arg Sample5) Sample5 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *Sample5Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Sample5Stream) Limit(limit int) *Sample5Stream {
	self.Slice(0, limit)
	return self
}
func (self *Sample5Stream) Map(fn func(Sample5, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Map2Int(fn func(Sample5, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Map2Int32(fn func(Sample5, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Map2Int64(fn func(Sample5, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Map2Float32(fn func(Sample5, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Map2Float64(fn func(Sample5, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Map2Bool(fn func(Sample5, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Map2Bytes(fn func(Sample5, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Map2String(fn func(Sample5, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample5Stream) Max(fn func(Sample5, int) float64) *Sample5 {
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
func (self *Sample5Stream) Min(fn func(Sample5, int) float64) *Sample5 {
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
func (self *Sample5Stream) NoneMatch(fn func(Sample5, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Sample5Stream) Get(index int) *Sample5 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Sample5Stream) GetOr(index int, arg Sample5) Sample5 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *Sample5Stream) Peek(fn func(*Sample5, int)) *Sample5Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Sample5Stream) Reduce(fn func(Sample5, Sample5, int) Sample5) *Sample5Stream {
	return self.ReduceInit(fn, Sample5{})
}
func (self *Sample5Stream) ReduceInit(fn func(Sample5, Sample5, int) Sample5, initialValue Sample5) *Sample5Stream {
	result := Sample5StreamOf()
	self.ForEach(func(v Sample5, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Sample5Stream) ReduceInterface(fn func(interface{}, Sample5, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Sample5{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample5Stream) ReduceString(fn func(string, Sample5, int) string) []string {
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
func (self *Sample5Stream) ReduceInt(fn func(int, Sample5, int) int) []int {
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
func (self *Sample5Stream) ReduceInt32(fn func(int32, Sample5, int) int32) []int32 {
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
func (self *Sample5Stream) ReduceInt64(fn func(int64, Sample5, int) int64) []int64 {
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
func (self *Sample5Stream) ReduceFloat32(fn func(float32, Sample5, int) float32) []float32 {
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
func (self *Sample5Stream) ReduceFloat64(fn func(float64, Sample5, int) float64) []float64 {
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
func (self *Sample5Stream) ReduceBool(fn func(bool, Sample5, int) bool) []bool {
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
func (self *Sample5Stream) Reverse() *Sample5Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Sample5Stream) Replace(fn func(Sample5, int) Sample5) *Sample5Stream {
	return self.ForEach(func(v Sample5, i int) { self.Set(i, fn(v, i)) })
}
func (self *Sample5Stream) Set(index int, val Sample5) *Sample5Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *Sample5Stream) Skip(skip int) *Sample5Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *Sample5Stream) SkippingEach(fn func(Sample5, int) int) *Sample5Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Sample5Stream) Slice(startIndex, n int) *Sample5Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Sample5{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *Sample5Stream) Sort(fn func(i, j int) bool) *Sample5Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Sample5Stream) Tail() *Sample5 {
	return self.Last()
}
func (self *Sample5Stream) TailOr(arg Sample5) Sample5 {
	return self.LastOr(arg)
}
func (self *Sample5Stream) ToList() []Sample5 {
	return self.Val()
}
func (self *Sample5Stream) Unique() *Sample5Stream {
	return self.Distinct()
}
func (self *Sample5Stream) Val() []Sample5 {
	if self == nil {
		return []Sample5{}
	}
	return *self.Copy()
}
func (self *Sample5Stream) While(fn func(Sample5, int) bool) *Sample5Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
