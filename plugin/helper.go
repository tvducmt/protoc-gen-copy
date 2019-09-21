package plugin

import (
	"reflect"

	"github.com/golang/glog"
)

// TrimFirstRune ...
func TrimFirstRune(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return ""
}

// ExistInFromArr ...
func ExistInFromArr(input string, fromArg []interface{}, isFirst bool) ObjQueue {
	if isFirst {
		for i := 1; i < len(fromArg); i++ {
			if _, ok := fromArg[i].(string); ok {
				if input == fromArg[i] {
					glog.Infoln("input == v", input, fromArg[i])
					if val, ok := fromArg[0].(ObjQueue); ok {
						return val
					}
					return ObjQueue{} //fromArg[0].(string)
				}
			} else {
				s := reflect.ValueOf(fromArg[i])
				if s.Kind() != reflect.Slice {
					panic("InterfaceSlice() given a non-slice type")
				}

				ret := make([]interface{}, s.Len())

				for i := 0; i < s.Len(); i++ {
					ret[i] = s.Index(i).Interface()
				}
				// glog.Infoln()
				// glog.Infoln()
				// glog.Infoln()
				// glog.Infoln("sdfgret", ret)
				// glog.Infoln("input", input)
				return ExistInFromArr(input, ret, true)

			}
		}
		return ObjQueue{Name: ""}
	}
	for _, v := range fromArg {
		if _, ok := v.(string); ok {
			if input == v {
				glog.Infoln("input == v", input, v)
				return ObjQueue{Name: input}
			}
		} else {
			s := reflect.ValueOf(v)
			if s.Kind() != reflect.Slice {
				panic("InterfaceSlice() given a non-slice type")
			}

			ret := make([]interface{}, s.Len())

			for i := 0; i < s.Len(); i++ {
				ret[i] = s.Index(i).Interface()
			}
			// glog.Infoln()
			// glog.Infoln()
			// glog.Infoln()
			// glog.Infoln("sdfgret", ret)
			// glog.Infoln("input", input)
			return ExistInFromArr(input, ret, true)

		}
	}
	return ObjQueue{Name: ""}

}
