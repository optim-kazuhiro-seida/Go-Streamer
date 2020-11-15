package sample

import (
	"math"
	"reflect"
	"sort"
)

type Sample0Stream []Sample0

func Sample0StreamOf(arg ...Sample0) Sample0Stream {
	return arg
}
func Sample0StreamFrom(arg []Sample0) Sample0Stream {
	return arg
}
func CreateSample0Stream(arg ...Sample0) *Sample0Stream {
	tmp := Sample0StreamOf(arg...)
	return &tmp
}
func GenerateSample0Stream(arg []Sample0) *Sample0Stream {
	tmp := Sample0StreamFrom(arg)
	return &tmp
}

func (self *Sample0Stream) Add(arg Sample0) *Sample0Stream {
	return self.AddAll(arg)
}
func (self *Sample0Stream) AddAll(arg ...Sample0) *Sample0Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Sample0Stream) AddSafe(arg *Sample0) *Sample0Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *Sample0Stream) AllMatch(fn func(Sample0, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Sample0Stream) AnyMatch(fn func(Sample0, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Sample0Stream) Clone() *Sample0Stream {
	temp := make([]Sample0, self.Len())
	copy(temp, *self)
	return (*Sample0Stream)(&temp)
}
func (self *Sample0Stream) Copy() *Sample0Stream {
	return self.Clone()
}
func (self *Sample0Stream) Concat(arg []Sample0) *Sample0Stream {
	return self.AddAll(arg...)
}
func (self *Sample0Stream) Contains(arg Sample0) bool {
	return self.FindIndex(func(_arg Sample0, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *Sample0Stream) Clean() *Sample0Stream {
	*self = Sample0StreamOf()
	return self
}
func (self *Sample0Stream) Delete(index int) *Sample0Stream {
	return self.DeleteRange(index, index)
}
func (self *Sample0Stream) DeleteRange(startIndex, endIndex int) *Sample0Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Sample0Stream) Distinct() *Sample0Stream {
	caches := map[Sample0]bool{}
	result := Sample0StreamOf()
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
func (self *Sample0Stream) Each(fn func(Sample0)) *Sample0Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *Sample0Stream) EachRight(fn func(Sample0)) *Sample0Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *Sample0Stream) Equals(arr []Sample0) bool {
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
func (self *Sample0Stream) Filter(fn func(Sample0, int) bool) *Sample0Stream {
	result := Sample0StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Sample0Stream) FilterSlim(fn func(Sample0, int) bool) *Sample0Stream {
	result := Sample0StreamOf()
	caches := map[Sample0]bool{}
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
func (self *Sample0Stream) Find(fn func(Sample0, int) bool) *Sample0 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *Sample0Stream) FindOr(fn func(Sample0, int) bool, or Sample0) Sample0 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *Sample0Stream) FindIndex(fn func(Sample0, int) bool) int {
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
func (self *Sample0Stream) First() *Sample0 {
	return self.Get(0)
}
func (self *Sample0Stream) FirstOr(arg Sample0) Sample0 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *Sample0Stream) ForEach(fn func(Sample0, int)) *Sample0Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Sample0Stream) ForEachRight(fn func(Sample0, int)) *Sample0Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Sample0Stream) GroupBy(fn func(Sample0, int) string) map[string][]Sample0 {
	m := map[string][]Sample0{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Sample0Stream) GroupByValues(fn func(Sample0, int) string) [][]Sample0 {
	var tmp [][]Sample0
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Sample0Stream) IndexOf(arg Sample0) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Sample0Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Sample0Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Sample0Stream) Last() *Sample0 {
	return self.Get(self.Len() - 1)
}
func (self *Sample0Stream) LastOr(arg Sample0) Sample0 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *Sample0Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Sample0Stream) Limit(limit int) *Sample0Stream {
	self.Slice(0, limit)
	return self
}
func (self *Sample0Stream) Map(fn func(Sample0, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Map2Int(fn func(Sample0, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Map2Int32(fn func(Sample0, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Map2Int64(fn func(Sample0, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Map2Float32(fn func(Sample0, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Map2Float64(fn func(Sample0, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Map2Bool(fn func(Sample0, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Map2Bytes(fn func(Sample0, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Map2String(fn func(Sample0, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample0Stream) Max(fn func(Sample0, int) float64) *Sample0 {
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
func (self *Sample0Stream) Min(fn func(Sample0, int) float64) *Sample0 {
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
func (self *Sample0Stream) NoneMatch(fn func(Sample0, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Sample0Stream) Get(index int) *Sample0 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Sample0Stream) GetOr(index int, arg Sample0) Sample0 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *Sample0Stream) Peek(fn func(*Sample0, int)) *Sample0Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Sample0Stream) Reduce(fn func(Sample0, Sample0, int) Sample0) *Sample0Stream {
	return self.ReduceInit(fn, Sample0{})
}
func (self *Sample0Stream) ReduceInit(fn func(Sample0, Sample0, int) Sample0, initialValue Sample0) *Sample0Stream {
	result := Sample0StreamOf()
	self.ForEach(func(v Sample0, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Sample0Stream) ReduceInterface(fn func(interface{}, Sample0, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Sample0{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample0Stream) ReduceString(fn func(string, Sample0, int) string) []string {
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
func (self *Sample0Stream) ReduceInt(fn func(int, Sample0, int) int) []int {
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
func (self *Sample0Stream) ReduceInt32(fn func(int32, Sample0, int) int32) []int32 {
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
func (self *Sample0Stream) ReduceInt64(fn func(int64, Sample0, int) int64) []int64 {
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
func (self *Sample0Stream) ReduceFloat32(fn func(float32, Sample0, int) float32) []float32 {
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
func (self *Sample0Stream) ReduceFloat64(fn func(float64, Sample0, int) float64) []float64 {
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
func (self *Sample0Stream) ReduceBool(fn func(bool, Sample0, int) bool) []bool {
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
func (self *Sample0Stream) Reverse() *Sample0Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Sample0Stream) Replace(fn func(Sample0, int) Sample0) *Sample0Stream {
	return self.ForEach(func(v Sample0, i int) { self.Set(i, fn(v, i)) })
}
func (self *Sample0Stream) Set(index int, val Sample0) *Sample0Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *Sample0Stream) Skip(skip int) *Sample0Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *Sample0Stream) SkippingEach(fn func(Sample0, int) int) *Sample0Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Sample0Stream) Slice(startIndex, n int) *Sample0Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Sample0{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *Sample0Stream) Sort(fn func(i, j int) bool) *Sample0Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Sample0Stream) Tail() *Sample0 {
	return self.Last()
}
func (self *Sample0Stream) TailOr(arg Sample0) Sample0 {
	return self.LastOr(arg)
}
func (self *Sample0Stream) ToList() []Sample0 {
	return self.Val()
}
func (self *Sample0Stream) Unique() *Sample0Stream {
	return self.Distinct()
}
func (self *Sample0Stream) Val() []Sample0 {
	if self == nil {
		return []Sample0{}
	}
	return *self.Copy()
}
func (self *Sample0Stream) While(fn func(Sample0, int) bool) *Sample0Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
