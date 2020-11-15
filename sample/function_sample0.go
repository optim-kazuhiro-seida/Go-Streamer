
package sample

import (
	"encoding/json"
	"reflect"
)

func UnmarshalSample0(byts []byte) (*Sample0, error) {
	sample := &Sample0{}
	if err := json.Unmarshal(byts, &sample); err != nil {
		return nil, err
	} else {
		return sample, nil
	}
}
func (self *Sample0) Copy(target *Sample0) (Sample0, error) {
	if byts, err := json.Marshal(*self); err != nil {
		return Sample0{}, err
	} else {
		return *target, json.Unmarshal(byts, target)
	}
}
func (self *Sample0) CopyOr(target *Sample0, _default Sample0) Sample0 {
	if _, err := self.Copy(target); err != nil {
		*target = _default
		return *target
	}
	return *target
}
func (self *Sample0) Equals(arg Sample0) bool {
	return reflect.DeepEqual(*self, arg)
}
func (self *Sample0) Fields() []reflect.StructField {
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

func (self *Sample0) MapFields(target interface{}) *Sample0 {
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
func (self *Sample0) MapJson(target interface{}) (*Sample0, error) {
	if byts, err := json.Marshal(*self); err != nil {
		return self, err
	} else {
		return self, json.Unmarshal(byts, target)
	}
}

func (self *Sample0) Marshal() ([]byte, error) {
	return json.Marshal(*self)
}
