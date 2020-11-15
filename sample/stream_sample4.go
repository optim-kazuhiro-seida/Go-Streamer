package sample

import (
	"math"
	"reflect"
	"sort"
)

type Sample4Stream []Sample4

func Sample4StreamOf(arg ...Sample4) Sample4Stream {
	return arg
}
func Sample4StreamFrom(arg []Sample4) Sample4Stream {
	return arg
}
func CreateSample4Stream(arg ...Sample4) *Sample4Stream {
	tmp := Sample4StreamOf(arg...)
	return &tmp
}
func GenerateSample4Stream(arg []Sample4) *Sample4Stream {
	tmp := Sample4StreamFrom(arg)
	return &tmp
}

func (self *Sample4Stream) Add(arg Sample4) *Sample4Stream {
	return self.AddAll(arg)
}
func (self *Sample4Stream) AddAll(arg ...Sample4) *Sample4Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Sample4Stream) AddSafe(arg *Sample4) *Sample4Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *Sample4Stream) AllMatch(fn func(Sample4, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Sample4Stream) AnyMatch(fn func(Sample4, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Sample4Stream) Clone() *Sample4Stream {
	temp := make([]Sample4, self.Len())
	copy(temp, *self)
	return (*Sample4Stream)(&temp)
}
func (self *Sample4Stream) Copy() *Sample4Stream {
	return self.Clone()
}
func (self *Sample4Stream) Concat(arg []Sample4) *Sample4Stream {
	return self.AddAll(arg...)
}
func (self *Sample4Stream) Contains(arg Sample4) bool {
	return self.FindIndex(func(_arg Sample4, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *Sample4Stream) Clean() *Sample4Stream {
	*self = Sample4StreamOf()
	return self
}
func (self *Sample4Stream) Delete(index int) *Sample4Stream {
	return self.DeleteRange(index, index)
}
func (self *Sample4Stream) DeleteRange(startIndex, endIndex int) *Sample4Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Sample4Stream) Distinct() *Sample4Stream {
	caches := map[Sample4]bool{}
	result := Sample4StreamOf()
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
func (self *Sample4Stream) Each(fn func(Sample4)) *Sample4Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *Sample4Stream) EachRight(fn func(Sample4)) *Sample4Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *Sample4Stream) Equals(arr []Sample4) bool {
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
func (self *Sample4Stream) Filter(fn func(Sample4, int) bool) *Sample4Stream {
	result := Sample4StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Sample4Stream) FilterSlim(fn func(Sample4, int) bool) *Sample4Stream {
	result := Sample4StreamOf()
	caches := map[Sample4]bool{}
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
func (self *Sample4Stream) Find(fn func(Sample4, int) bool) *Sample4 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *Sample4Stream) FindOr(fn func(Sample4, int) bool, or Sample4) Sample4 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *Sample4Stream) FindIndex(fn func(Sample4, int) bool) int {
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
func (self *Sample4Stream) First() *Sample4 {
	return self.Get(0)
}
func (self *Sample4Stream) FirstOr(arg Sample4) Sample4 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *Sample4Stream) ForEach(fn func(Sample4, int)) *Sample4Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Sample4Stream) ForEachRight(fn func(Sample4, int)) *Sample4Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Sample4Stream) GroupBy(fn func(Sample4, int) string) map[string][]Sample4 {
	m := map[string][]Sample4{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Sample4Stream) GroupByValues(fn func(Sample4, int) string) [][]Sample4 {
	var tmp [][]Sample4
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Sample4Stream) IndexOf(arg Sample4) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Sample4Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Sample4Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Sample4Stream) Last() *Sample4 {
	return self.Get(self.Len() - 1)
}
func (self *Sample4Stream) LastOr(arg Sample4) Sample4 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *Sample4Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Sample4Stream) Limit(limit int) *Sample4Stream {
	self.Slice(0, limit)
	return self
}
func (self *Sample4Stream) Map(fn func(Sample4, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Map2Int(fn func(Sample4, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Map2Int32(fn func(Sample4, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Map2Int64(fn func(Sample4, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Map2Float32(fn func(Sample4, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Map2Float64(fn func(Sample4, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Map2Bool(fn func(Sample4, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Map2Bytes(fn func(Sample4, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Map2String(fn func(Sample4, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Sample4Stream) Max(fn func(Sample4, int) float64) *Sample4 {
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
func (self *Sample4Stream) Min(fn func(Sample4, int) float64) *Sample4 {
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
func (self *Sample4Stream) NoneMatch(fn func(Sample4, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Sample4Stream) Get(index int) *Sample4 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Sample4Stream) GetOr(index int, arg Sample4) Sample4 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *Sample4Stream) Peek(fn func(*Sample4, int)) *Sample4Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Sample4Stream) Reduce(fn func(Sample4, Sample4, int) Sample4) *Sample4Stream {
	return self.ReduceInit(fn, Sample4{})
}
func (self *Sample4Stream) ReduceInit(fn func(Sample4, Sample4, int) Sample4, initialValue Sample4) *Sample4Stream {
	result := Sample4StreamOf()
	self.ForEach(func(v Sample4, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Sample4Stream) ReduceInterface(fn func(interface{}, Sample4, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Sample4{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Sample4Stream) ReduceString(fn func(string, Sample4, int) string) []string {
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
func (self *Sample4Stream) ReduceInt(fn func(int, Sample4, int) int) []int {
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
func (self *Sample4Stream) ReduceInt32(fn func(int32, Sample4, int) int32) []int32 {
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
func (self *Sample4Stream) ReduceInt64(fn func(int64, Sample4, int) int64) []int64 {
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
func (self *Sample4Stream) ReduceFloat32(fn func(float32, Sample4, int) float32) []float32 {
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
func (self *Sample4Stream) ReduceFloat64(fn func(float64, Sample4, int) float64) []float64 {
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
func (self *Sample4Stream) ReduceBool(fn func(bool, Sample4, int) bool) []bool {
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
func (self *Sample4Stream) Reverse() *Sample4Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Sample4Stream) Replace(fn func(Sample4, int) Sample4) *Sample4Stream {
	return self.ForEach(func(v Sample4, i int) { self.Set(i, fn(v, i)) })
}
func (self *Sample4Stream) Set(index int, val Sample4) *Sample4Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *Sample4Stream) Skip(skip int) *Sample4Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *Sample4Stream) SkippingEach(fn func(Sample4, int) int) *Sample4Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Sample4Stream) Slice(startIndex, n int) *Sample4Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Sample4{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *Sample4Stream) Sort(fn func(i, j int) bool) *Sample4Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Sample4Stream) Tail() *Sample4 {
	return self.Last()
}
func (self *Sample4Stream) TailOr(arg Sample4) Sample4 {
	return self.LastOr(arg)
}
func (self *Sample4Stream) ToList() []Sample4 {
	return self.Val()
}
func (self *Sample4Stream) Unique() *Sample4Stream {
	return self.Distinct()
}
func (self *Sample4Stream) Val() []Sample4 {
	if self == nil {
		return []Sample4{}
	}
	return *self.Copy()
}
func (self *Sample4Stream) While(fn func(Sample4, int) bool) *Sample4Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
