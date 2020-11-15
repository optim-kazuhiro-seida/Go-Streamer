package sample

import (
	"math"
	"reflect"
	"sort"
)

type SampleStream []Sample

func SampleStreamOf(arg ...Sample) SampleStream {
	return arg
}
func SampleStreamFrom(arg []Sample) SampleStream {
	return arg
}
func CreateSampleStream(arg ...Sample) *SampleStream {
	tmp := SampleStreamOf(arg...)
	return &tmp
}
func GenerateSampleStream(arg []Sample) *SampleStream {
	tmp := SampleStreamFrom(arg)
	return &tmp
}

func (self *SampleStream) Add(arg Sample) *SampleStream {
	return self.AddAll(arg)
}
func (self *SampleStream) AddAll(arg ...Sample) *SampleStream {
	*self = append(*self, arg...)
	return self
}
func (self *SampleStream) AddSafe(arg *Sample) *SampleStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SampleStream) AllMatch(fn func(Sample, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SampleStream) AnyMatch(fn func(Sample, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SampleStream) Clone() *SampleStream {
	temp := make([]Sample, self.Len())
	copy(temp, *self)
	return (*SampleStream)(&temp)
}
func (self *SampleStream) Copy() *SampleStream {
	return self.Clone()
}
func (self *SampleStream) Concat(arg []Sample) *SampleStream {
	return self.AddAll(arg...)
}
func (self *SampleStream) Contains(arg Sample) bool {
	return self.FindIndex(func(_arg Sample, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SampleStream) Clean() *SampleStream {
	*self = SampleStreamOf()
	return self
}
func (self *SampleStream) Delete(index int) *SampleStream {
	return self.DeleteRange(index, index)
}
func (self *SampleStream) DeleteRange(startIndex, endIndex int) *SampleStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SampleStream) Distinct() *SampleStream {
	caches := map[Sample]bool{}
	result := SampleStreamOf()
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
func (self *SampleStream) Each(fn func(Sample)) *SampleStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SampleStream) EachRight(fn func(Sample)) *SampleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SampleStream) Equals(arr []Sample) bool {
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
func (self *SampleStream) Filter(fn func(Sample, int) bool) *SampleStream {
	result := SampleStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SampleStream) FilterSlim(fn func(Sample, int) bool) *SampleStream {
	result := SampleStreamOf()
	caches := map[Sample]bool{}
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
func (self *SampleStream) Find(fn func(Sample, int) bool) *Sample {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SampleStream) FindOr(fn func(Sample, int) bool, or Sample) Sample {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SampleStream) FindIndex(fn func(Sample, int) bool) int {
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
func (self *SampleStream) First() *Sample {
	return self.Get(0)
}
func (self *SampleStream) FirstOr(arg Sample) Sample {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SampleStream) ForEach(fn func(Sample, int)) *SampleStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SampleStream) ForEachRight(fn func(Sample, int)) *SampleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SampleStream) GroupBy(fn func(Sample, int) string) map[string][]Sample {
	m := map[string][]Sample{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SampleStream) GroupByValues(fn func(Sample, int) string) [][]Sample {
	var tmp [][]Sample
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SampleStream) IndexOf(arg Sample) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SampleStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SampleStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SampleStream) Last() *Sample {
	return self.Get(self.Len() - 1)
}
func (self *SampleStream) LastOr(arg Sample) Sample {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SampleStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SampleStream) Limit(limit int) *SampleStream {
	self.Slice(0, limit)
	return self
}
func (self *SampleStream) Map(fn func(Sample, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Map2Int(fn func(Sample, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Map2Int32(fn func(Sample, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Map2Int64(fn func(Sample, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Map2Float32(fn func(Sample, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Map2Float64(fn func(Sample, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Map2Bool(fn func(Sample, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Map2Bytes(fn func(Sample, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Map2String(fn func(Sample, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleStream) Max(fn func(Sample, int) float64) *Sample {
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
func (self *SampleStream) Min(fn func(Sample, int) float64) *Sample {
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
func (self *SampleStream) NoneMatch(fn func(Sample, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SampleStream) Get(index int) *Sample {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SampleStream) GetOr(index int, arg Sample) Sample {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SampleStream) Peek(fn func(*Sample, int)) *SampleStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *SampleStream) Reduce(fn func(Sample, Sample, int) Sample) *SampleStream {
	return self.ReduceInit(fn, Sample{})
}
func (self *SampleStream) ReduceInit(fn func(Sample, Sample, int) Sample, initialValue Sample) *SampleStream {
	result := SampleStreamOf()
	self.ForEach(func(v Sample, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SampleStream) ReduceInterface(fn func(interface{}, Sample, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Sample{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SampleStream) ReduceString(fn func(string, Sample, int) string) []string {
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
func (self *SampleStream) ReduceInt(fn func(int, Sample, int) int) []int {
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
func (self *SampleStream) ReduceInt32(fn func(int32, Sample, int) int32) []int32 {
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
func (self *SampleStream) ReduceInt64(fn func(int64, Sample, int) int64) []int64 {
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
func (self *SampleStream) ReduceFloat32(fn func(float32, Sample, int) float32) []float32 {
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
func (self *SampleStream) ReduceFloat64(fn func(float64, Sample, int) float64) []float64 {
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
func (self *SampleStream) ReduceBool(fn func(bool, Sample, int) bool) []bool {
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
func (self *SampleStream) Reverse() *SampleStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SampleStream) Replace(fn func(Sample, int) Sample) *SampleStream {
	return self.ForEach(func(v Sample, i int) { self.Set(i, fn(v, i)) })
}
func (self *SampleStream) Set(index int, val Sample) *SampleStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SampleStream) Skip(skip int) *SampleStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SampleStream) SkippingEach(fn func(Sample, int) int) *SampleStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SampleStream) Slice(startIndex, n int) *SampleStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Sample{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SampleStream) Sort(fn func(i, j int) bool) *SampleStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SampleStream) Tail() *Sample {
	return self.Last()
}
func (self *SampleStream) TailOr(arg Sample) Sample {
	return self.LastOr(arg)
}
func (self *SampleStream) ToList() []Sample {
	return self.Val()
}
func (self *SampleStream) Unique() *SampleStream {
	return self.Distinct()
}
func (self *SampleStream) Val() []Sample {
	if self == nil {
		return []Sample{}
	}
	return *self.Copy()
}
func (self *SampleStream) While(fn func(Sample, int) bool) *SampleStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
