
package sample

import (
	"encoding/json"
	"reflect"
)

func UnmarshalSample1(byts []byte) (*Sample1, error) {
	sample := &Sample1{}
	if err := json.Unmarshal(byts, &sample); err != nil {
		return nil, err
	} else {
		return sample, nil
	}
}
func (self *Sample1) Copy(target *Sample1) (Sample1, error) {
	if byts, err := json.Marshal(*self); err != nil {
		return Sample1{}, err
	} else {
		return *target, json.Unmarshal(byts, target)
	}
}
func (self *Sample1) CopyOr(target *Sample1, _default Sample1) Sample1 {
	if _, err := self.Copy(target); err != nil {
		*target = _default
		return *target
	}
	return *target
}
func (self *Sample1) Equals(arg Sample1) bool {
	return reflect.DeepEqual(*self, arg)
}
func (self *Sample1) Fields() []reflect.StructField {
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

func (self *Sample1) MapFields(target interface{}) *Sample1 {
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
func (self *Sample1) MapJson(target interface{}) (*Sample1, error) {
	if byts, err := json.Marshal(*self); err != nil {
		return self, err
	} else {
		return self, json.Unmarshal(byts, target)
	}
}

func (self *Sample1) Marshal() ([]byte, error) {
	return json.Marshal(*self)
}
