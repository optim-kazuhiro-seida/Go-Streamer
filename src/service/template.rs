pub static TEMPLATE: &str =
    "package {{.PackageName}}

import (
	\"math\"
	\"reflect\"
	\"sort\"
)

type {{.TypeName}}Stream []{{.TypeName}}

func {{.TypeName}}StreamOf(arg ...{{.TypeName}}) {{.TypeName}}Stream {
	return arg
}
func {{.TypeName}}StreamFrom(arg []{{.TypeName}}) {{.TypeName}}Stream {
	return arg
}
func Create{{.TypeName}}Stream(arg ...{{.TypeName}}) *{{.TypeName}}Stream {
	tmp := {{.TypeName}}StreamOf(arg...)
	return &tmp
}
func Generate{{.TypeName}}Stream(arg []{{.TypeName}}) *{{.TypeName}}Stream {
	tmp := {{.TypeName}}StreamFrom(arg)
	return &tmp
}

func (self *{{.TypeName}}Stream) Add(arg {{.TypeName}}) *{{.TypeName}}Stream {
	return self.AddAll(arg)
}
func (self *{{.TypeName}}Stream) AddAll(arg ...{{.TypeName}}) *{{.TypeName}}Stream {
	*self = append(*self, arg...)
	return self
}
func (self *{{.TypeName}}Stream) AddSafe(arg *{{.TypeName}}) *{{.TypeName}}Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *{{.TypeName}}Stream) Aggregate(fn func({{.TypeName}}, {{.TypeName}}) {{.TypeName}}) *{{.TypeName}}Stream {
	result := {{.TypeName}}StreamOf()
	self.ForEach(func(v {{.TypeName}}, i int) {
		if i == 0 {
			result.Add(fn({{.TypeName}}{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *{{.TypeName}}Stream) AllMatch(fn func({{.TypeName}}, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *{{.TypeName}}Stream) AnyMatch(fn func({{.TypeName}}, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *{{.TypeName}}Stream) Clone() *{{.TypeName}}Stream {
	temp := make([]{{.TypeName}}, self.Len())
	copy(temp, *self)
	return (*{{.TypeName}}Stream)(&temp)
}
func (self *{{.TypeName}}Stream) Copy() *{{.TypeName}}Stream {
	return self.Clone()
}
func (self *{{.TypeName}}Stream) Concat(arg []{{.TypeName}}) *{{.TypeName}}Stream {
	return self.AddAll(arg...)
}
func (self *{{.TypeName}}Stream) Contains(arg {{.TypeName}}) bool {
	return self.FindIndex(func(_arg {{.TypeName}}, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *{{.TypeName}}Stream) Clean() *{{.TypeName}}Stream {
	*self = {{.TypeName}}StreamOf()
	return self
}
func (self *{{.TypeName}}Stream) Delete(index int) *{{.TypeName}}Stream {
	return self.DeleteRange(index, index)
}
func (self *{{.TypeName}}Stream) DeleteRange(startIndex, endIndex int) *{{.TypeName}}Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *{{.TypeName}}Stream) Distinct() *{{.TypeName}}Stream {
	caches := map[{{.TypeName}}]bool{}
	result := {{.TypeName}}StreamOf()
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
func (self *{{.TypeName}}Stream) Each(fn func({{.TypeName}})) *{{.TypeName}}Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *{{.TypeName}}Stream) EachRight(fn func({{.TypeName}})) *{{.TypeName}}Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *{{.TypeName}}Stream) Equals(arr []{{.TypeName}}) bool {
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
func (self *{{.TypeName}}Stream) Filter(fn func({{.TypeName}}, int) bool) *{{.TypeName}}Stream {
	result := {{.TypeName}}StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *{{.TypeName}}Stream) FilterSlim(fn func({{.TypeName}}, int) bool) *{{.TypeName}}Stream {
	result := {{.TypeName}}StreamOf()
	caches := map[{{.TypeName}}]bool{}
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
func (self *{{.TypeName}}Stream) Find(fn func({{.TypeName}}, int) bool) *{{.TypeName}} {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *{{.TypeName}}Stream) FindOr(fn func({{.TypeName}}, int) bool, or {{.TypeName}}) {{.TypeName}} {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *{{.TypeName}}Stream) FindIndex(fn func({{.TypeName}}, int) bool) int {
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
func (self *{{.TypeName}}Stream) First() *{{.TypeName}} {
	return self.Get(0)
}
func (self *{{.TypeName}}Stream) FirstOr(arg {{.TypeName}}) {{.TypeName}} {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *{{.TypeName}}Stream) ForEach(fn func({{.TypeName}}, int)) *{{.TypeName}}Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *{{.TypeName}}Stream) ForEachRight(fn func({{.TypeName}}, int)) *{{.TypeName}}Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *{{.TypeName}}Stream) GroupBy(fn func({{.TypeName}}, int) string) map[string][]{{.TypeName}} {
	m := map[string][]{{.TypeName}}{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *{{.TypeName}}Stream) GroupByValues(fn func({{.TypeName}}, int) string) [][]{{.TypeName}} {
	var tmp [][]{{.TypeName}}
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *{{.TypeName}}Stream) IndexOf(arg {{.TypeName}}) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *{{.TypeName}}Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *{{.TypeName}}Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *{{.TypeName}}Stream) Last() *{{.TypeName}} {
	return self.Get(self.Len() - 1)
}
func (self *{{.TypeName}}Stream) LastOr(arg {{.TypeName}}) {{.TypeName}} {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *{{.TypeName}}Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *{{.TypeName}}Stream) Limit(limit int) *{{.TypeName}}Stream {
	self.Slice(0, limit)
	return self
}

func (self *{{.TypeName}}Stream) Map(fn func({{.TypeName}}, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Map2Int(fn func({{.TypeName}}, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Map2Int32(fn func({{.TypeName}}, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Map2Int64(fn func({{.TypeName}}, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Map2Float32(fn func({{.TypeName}}, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Map2Float64(fn func({{.TypeName}}, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Map2Bool(fn func({{.TypeName}}, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Map2Bytes(fn func({{.TypeName}}, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Map2String(fn func({{.TypeName}}, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Max(fn func({{.TypeName}}, int) float64) *{{.TypeName}} {
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
func (self *{{.TypeName}}Stream) Min(fn func({{.TypeName}}, int) float64) *{{.TypeName}} {
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
func (self *{{.TypeName}}Stream) NoneMatch(fn func({{.TypeName}}, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *{{.TypeName}}Stream) Get(index int) *{{.TypeName}} {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *{{.TypeName}}Stream) GetOr(index int, arg {{.TypeName}}) {{.TypeName}} {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *{{.TypeName}}Stream) Peek(fn func(*{{.TypeName}}, int)) *{{.TypeName}}Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *{{.TypeName}}Stream) Reduce(fn func({{.TypeName}}, {{.TypeName}}, int) {{.TypeName}}) *{{.TypeName}}Stream {
	return self.ReduceInit(fn, {{.TypeName}}{})
}
func (self *{{.TypeName}}Stream) ReduceInit(fn func({{.TypeName}}, {{.TypeName}}, int) {{.TypeName}}, initialValue {{.TypeName}}) *{{.TypeName}}Stream {
	result := {{.TypeName}}StreamOf()
	self.ForEach(func(v {{.TypeName}}, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *{{.TypeName}}Stream) ReduceInterface(fn func(interface{}, {{.TypeName}}, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn({{.TypeName}}{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *{{.TypeName}}Stream) ReduceString(fn func(string, {{.TypeName}}, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(\"\", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *{{.TypeName}}Stream) ReduceInt(fn func(int, {{.TypeName}}, int) int) []int {
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
func (self *{{.TypeName}}Stream) ReduceInt32(fn func(int32, {{.TypeName}}, int) int32) []int32 {
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
func (self *{{.TypeName}}Stream) ReduceInt64(fn func(int64, {{.TypeName}}, int) int64) []int64 {
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
func (self *{{.TypeName}}Stream) ReduceFloat32(fn func(float32, {{.TypeName}}, int) float32) []float32 {
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
func (self *{{.TypeName}}Stream) ReduceFloat64(fn func(float64, {{.TypeName}}, int) float64) []float64 {
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
func (self *{{.TypeName}}Stream) ReduceBool(fn func(bool, {{.TypeName}}, int) bool) []bool {
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
func (self *{{.TypeName}}Stream) Reverse() *{{.TypeName}}Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *{{.TypeName}}Stream) Replace(fn func({{.TypeName}}, int) {{.TypeName}}) *{{.TypeName}}Stream {
	return self.ForEach(func(v {{.TypeName}}, i int) { self.Set(i, fn(v, i)) })
}
func (self *{{.TypeName}}Stream) Select(fn func({{.TypeName}}) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *{{.TypeName}}Stream) Set(index int, val {{.TypeName}}) *{{.TypeName}}Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *{{.TypeName}}Stream) Skip(skip int) *{{.TypeName}}Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *{{.TypeName}}Stream) SkippingEach(fn func({{.TypeName}}, int) int) *{{.TypeName}}Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *{{.TypeName}}Stream) Slice(startIndex, n int) *{{.TypeName}}Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []{{.TypeName}}{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *{{.TypeName}}Stream) Sort(fn func(i, j int) bool) *{{.TypeName}}Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *{{.TypeName}}Stream) Tail() *{{.TypeName}} {
	return self.Last()
}
func (self *{{.TypeName}}Stream) TailOr(arg {{.TypeName}}) {{.TypeName}} {
	return self.LastOr(arg)
}
func (self *{{.TypeName}}Stream) ToList() []{{.TypeName}} {
	return self.Val()
}
func (self *{{.TypeName}}Stream) Unique() *{{.TypeName}}Stream {
	return self.Distinct()
}
func (self *{{.TypeName}}Stream) Val() []{{.TypeName}} {
	if self == nil {
		return []{{.TypeName}}{}
	}
	return *self.Copy()
}
func (self *{{.TypeName}}Stream) While(fn func({{.TypeName}}, int) bool) *{{.TypeName}}Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *{{.TypeName}}Stream) Where(fn func({{.TypeName}}) bool) *{{.TypeName}}Stream {
	result := {{.TypeName}}StreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *{{.TypeName}}Stream) WhereSlim(fn func({{.TypeName}}) bool) *{{.TypeName}}Stream {
	result := {{.TypeName}}StreamOf()
	caches := map[{{.TypeName}}]bool{}
	for _, v := range *self {
		if f, ok := caches[v]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[v] = fn(v); caches[v] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
";
pub static TEMPLATE_FUNC: &str ="package {{.PackageName}}

import (
	\"encoding/json\"
	\"reflect\"
)

func Unmarshal{{.TypeName}}(byts []byte) (*{{.TypeName}}, error) {
	sample := &{{.TypeName}}{}
	if err := json.Unmarshal(byts, &sample); err != nil {
		return nil, err
	} else {
		return sample, nil
	}
}
func (self *{{.TypeName}}) Copy(target *{{.TypeName}}) ({{.TypeName}}, error) {
	if byts, err := json.Marshal(*self); err != nil {
		return {{.TypeName}}{}, err
	} else {
		return *target, json.Unmarshal(byts, target)
	}
}
func (self *{{.TypeName}}) CopyOr(target *{{.TypeName}}, _default {{.TypeName}}) {{.TypeName}} {
	if _, err := self.Copy(target); err != nil {
		*target = _default
		return *target
	}
	return *target
}
func (self *{{.TypeName}}) Equals(arg {{.TypeName}}) bool {
	return reflect.DeepEqual(*self, arg)
}
func (self *{{.TypeName}}) Fields() []reflect.StructField {
	var (
		fs []reflect.StructField
		v  = reflect.Indirect(reflect.ValueOf(self))
	)
	for i, t := 0, v.Type(); i < t.NumField(); i++ {
		if field := t.Field(i); field.PkgPath == \"\" {
			fs = append(fs, field)
		}
	}
	return fs
}

func (self *{{.TypeName}}) MapFields(target interface{}) *{{.TypeName}} {
	var (
		_self   = reflect.Indirect(reflect.ValueOf(self))
		_target = reflect.ValueOf(target)
	)

	if _target.Kind() == reflect.Invalid || _target.Kind() != reflect.Ptr {
		return self
	}
	if _target = reflect.Indirect(_target); _target.Type().Kind() != reflect.Struct {
		return self
	}
	for i, t := 0, _self.Type(); i < t.NumField(); i++ {
		if field := t.Field(i); field.PkgPath == \"\" {
			_targetField := _target.FieldByName(field.Name)
			if f, ok := _target.Type().FieldByName(field.Name); ok && f.Type == field.Type && f.PkgPath == \"\" {
				_targetField.Set(_self.FieldByName(field.Name))
			}
		}
	}
	return self
}
func (self *{{.TypeName}}) MapJson(target interface{}) (*{{.TypeName}}, error) {
	if byts, err := json.Marshal(*self); err != nil {
		return self, err
	} else {
		return self, json.Unmarshal(byts, target)
	}
}

func (self *{{.TypeName}}) Marshal() ([]byte, error) {
	return json.Marshal(*self)
}
";
