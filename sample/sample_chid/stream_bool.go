/*
 * Collection utility of Bool Struct
 *
 * Generated by: Go Streamer
 */

package sample

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type BoolStream []Bool

func BoolStreamOf(arg ...Bool) BoolStream {
	return arg
}
func BoolStreamFrom(arg []Bool) BoolStream {
	return arg
}
func CreateBoolStream(arg ...Bool) *BoolStream {
	tmp := BoolStreamOf(arg...)
	return &tmp
}
func GenerateBoolStream(arg []Bool) *BoolStream {
	tmp := BoolStreamFrom(arg)
	return &tmp
}

func (self *BoolStream) Add(arg Bool) *BoolStream {
	return self.AddAll(arg)
}
func (self *BoolStream) AddAll(arg ...Bool) *BoolStream {
	*self = append(*self, arg...)
	return self
}
func (self *BoolStream) AddSafe(arg *Bool) *BoolStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *BoolStream) Aggregate(fn func(Bool, Bool) Bool) *BoolStream {
	result := BoolStreamOf()
	self.ForEach(func(v Bool, i int) {
		if i == 0 {
			result.Add(fn(Bool{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *BoolStream) AllMatch(fn func(Bool, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *BoolStream) AnyMatch(fn func(Bool, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *BoolStream) Clone() *BoolStream {
	temp := make([]Bool, self.Len())
	copy(temp, *self)
	return (*BoolStream)(&temp)
}
func (self *BoolStream) Copy() *BoolStream {
	return self.Clone()
}
func (self *BoolStream) Concat(arg []Bool) *BoolStream {
	return self.AddAll(arg...)
}
func (self *BoolStream) Contains(arg Bool) bool {
	return self.FindIndex(func(_arg Bool, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *BoolStream) Clean() *BoolStream {
	*self = BoolStreamOf()
	return self
}
func (self *BoolStream) Delete(index int) *BoolStream {
	return self.DeleteRange(index, index)
}
func (self *BoolStream) DeleteRange(startIndex, endIndex int) *BoolStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *BoolStream) Distinct() *BoolStream {
	caches := map[string]bool{}
	result := BoolStreamOf()
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[key] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *BoolStream) Each(fn func(Bool)) *BoolStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *BoolStream) EachRight(fn func(Bool)) *BoolStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *BoolStream) Equals(arr []Bool) bool {
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
func (self *BoolStream) Filter(fn func(Bool, int) bool) *BoolStream {
	result := BoolStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BoolStream) FilterSlim(fn func(Bool, int) bool) *BoolStream {
	result := BoolStreamOf()
	caches := map[string]bool{}
	for i, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v, i); caches[key] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *BoolStream) Find(fn func(Bool, int) bool) *Bool {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *BoolStream) FindOr(fn func(Bool, int) bool, or Bool) Bool {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *BoolStream) FindIndex(fn func(Bool, int) bool) int {
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
func (self *BoolStream) First() *Bool {
	return self.Get(0)
}
func (self *BoolStream) FirstOr(arg Bool) Bool {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *BoolStream) ForEach(fn func(Bool, int)) *BoolStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *BoolStream) ForEachRight(fn func(Bool, int)) *BoolStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *BoolStream) GroupBy(fn func(Bool, int) string) map[string][]Bool {
	m := map[string][]Bool{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *BoolStream) GroupByValues(fn func(Bool, int) string) [][]Bool {
	var tmp [][]Bool
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *BoolStream) IndexOf(arg Bool) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *BoolStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *BoolStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *BoolStream) Last() *Bool {
	return self.Get(self.Len() - 1)
}
func (self *BoolStream) LastOr(arg Bool) Bool {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *BoolStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *BoolStream) Limit(limit int) *BoolStream {
	self.Slice(0, limit)
	return self
}

func (self *BoolStream) Map(fn func(Bool, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Int(fn func(Bool, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Int32(fn func(Bool, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Int64(fn func(Bool, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Float32(fn func(Bool, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Float64(fn func(Bool, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Bool(fn func(Bool, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Bytes(fn func(Bool, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2String(fn func(Bool, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Max(fn func(Bool, int) float64) *Bool {
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
func (self *BoolStream) Min(fn func(Bool, int) float64) *Bool {
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
func (self *BoolStream) NoneMatch(fn func(Bool, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *BoolStream) Get(index int) *Bool {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *BoolStream) GetOr(index int, arg Bool) Bool {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *BoolStream) Peek(fn func(*Bool, int)) *BoolStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *BoolStream) Reduce(fn func(Bool, Bool, int) Bool) *BoolStream {
	return self.ReduceInit(fn, Bool{})
}
func (self *BoolStream) ReduceInit(fn func(Bool, Bool, int) Bool, initialValue Bool) *BoolStream {
	result := BoolStreamOf()
	self.ForEach(func(v Bool, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *BoolStream) ReduceInterface(fn func(interface{}, Bool, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Bool{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *BoolStream) ReduceString(fn func(string, Bool, int) string) []string {
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
func (self *BoolStream) ReduceInt(fn func(int, Bool, int) int) []int {
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
func (self *BoolStream) ReduceInt32(fn func(int32, Bool, int) int32) []int32 {
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
func (self *BoolStream) ReduceInt64(fn func(int64, Bool, int) int64) []int64 {
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
func (self *BoolStream) ReduceFloat32(fn func(float32, Bool, int) float32) []float32 {
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
func (self *BoolStream) ReduceFloat64(fn func(float64, Bool, int) float64) []float64 {
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
func (self *BoolStream) ReduceBool(fn func(bool, Bool, int) bool) []bool {
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
func (self *BoolStream) Reverse() *BoolStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *BoolStream) Replace(fn func(Bool, int) Bool) *BoolStream {
	return self.ForEach(func(v Bool, i int) { self.Set(i, fn(v, i)) })
}
func (self *BoolStream) Select(fn func(Bool) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *BoolStream) Set(index int, val Bool) *BoolStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *BoolStream) Skip(skip int) *BoolStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *BoolStream) SkippingEach(fn func(Bool, int) int) *BoolStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *BoolStream) Slice(startIndex, n int) *BoolStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Bool{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *BoolStream) Sort(fn func(i, j int) bool) *BoolStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *BoolStream) Tail() *Bool {
	return self.Last()
}
func (self *BoolStream) TailOr(arg Bool) Bool {
	return self.LastOr(arg)
}
func (self *BoolStream) ToList() []Bool {
	return self.Val()
}
func (self *BoolStream) Unique() *BoolStream {
	return self.Distinct()
}
func (self *BoolStream) Val() []Bool {
	if self == nil {
		return []Bool{}
	}
	return *self.Copy()
}
func (self *BoolStream) While(fn func(Bool, int) bool) *BoolStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *BoolStream) Where(fn func(Bool) bool) *BoolStream {
	result := BoolStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BoolStream) WhereSlim(fn func(Bool) bool) *BoolStream {
	result := BoolStreamOf()
	caches := map[string]bool{}
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v); caches[key] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
