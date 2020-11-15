
package sample

import (
	"encoding/json"
	"reflect"
)

func UnmarshalSample2(byts []byte) (*Sample2, error) {
	sample := &Sample2{}
	if err := json.Unmarshal(byts, &sample); err != nil {
		return nil, err
	} else {
		return sample, nil
	}
}
func (self *Sample2) Copy(target *Sample2) (Sample2, error) {
	if byts, err := json.Marshal(*self); err != nil {
		return Sample2{}, err
	} else {
		return *target, json.Unmarshal(byts, target)
	}
}
func (self *Sample2) CopyOr(target *Sample2, _default Sample2) Sample2 {
	if _, err := self.Copy(target); err != nil {
		*target = _default
		return *target
	}
	return *target
}
func (self *Sample2) Equals(arg Sample2) bool {
	return reflect.DeepEqual(*self, arg)
}
func (self *Sample2) Fields() []reflect.StructField {
	var (
		fs []reflect.StructField
		v  = reflect.Indirect(reflect.ValueOf(self))
	)
	for i, t := 0, v.Type(); i < t.NumField(); i++ {
		if field := t.Field(i); field.PkgPath == "" {
			fs = append(fs, field)
		}
	}
	return fs
}

func (self *Sample2) MapFields(target interface{}) *Sample2 {
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
		if field := t.Field(i); field.PkgPath == "" {
			_targetField := _target.FieldByName(field.Name)
			if f, ok := _target.Type().FieldByName(field.Name); ok && f.Type == field.Type && f.PkgPath == "" {
				_targetField.Set(_self.FieldByName(field.Name))
			}
		}
	}
	return self
}
func (self *Sample2) MapJson(target interface{}) (*Sample2, error) {
	if byts, err := json.Marshal(*self); err != nil {
		return self, err
	} else {
		return self, json.Unmarshal(byts, target)
	}
}

func (self *Sample2) Marshal() ([]byte, error) {
	return json.Marshal(*self)
}
