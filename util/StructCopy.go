package util

import (
	"errors"
	"reflect"
)

func Copy(dst, src interface{}) (err error) {
	d := reflect.ValueOf(dst)
	s := reflect.ValueOf(src)
	//判断是否为同一种类型
	if d.Kind() != s.Kind() {
		return errors.New("dst,src not the same type ")
	}
	st := s.Elem().Type()
	for i := 0; i < s.Elem().NumField(); i++ {
		name := st.Field(i).Name
		if ok := d.Elem().FieldByName(name).IsValid(); ok {
			d.Elem().FieldByName(name).Set(reflect.ValueOf(s.Elem().Field(i).Interface()))
		}
	}
	return err
}
