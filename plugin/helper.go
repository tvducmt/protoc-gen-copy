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

func checkIntoArrString(input string, fromArg []interface{}) ObjQueue {
	// duyet cac phan tu don
	for i := 1; i < len(fromArg); i++ {
		if val, ok := fromArg[i].(string); ok {
			if input == val {
				if val, ok := fromArg[0].(ObjQueue); ok {
					return val
				}
				return ObjQueue{}
			}
		}
	}

	//duyet cac phan tu trong mang
	for i := 1; i < len(fromArg); i++ {
		s := reflect.ValueOf(fromArg[i])
		if s.Kind() != reflect.Slice {
			continue
		}

		ret := make([]interface{}, s.Len())

		for i := 0; i < s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
		}
		rls := checkIntoArrString(input, ret)
		if rls.Name != "" {
			return rls
		}

		// }
	}
	return ObjQueue{Name: ""}
}

// ExistInFromArr ...
func ExistInFromArrString(input string, fromArg []interface{}) ObjQueue {

	for _, v := range fromArg {
		if _, ok := v.(string); ok {
			if input == v {
				// glog.Infoln("input == v", input, v)
				return ObjQueue{Name: input}
			}
		} else {
			s := reflect.ValueOf(v)
			if s.Kind() != reflect.Slice {
				glog.Infoln("s.Kind()", s.Kind(), s, input)
				panic("InterfaceSlice() given a non-slice type1")
			}

			ret := make([]interface{}, s.Len())

			for i := 0; i < s.Len(); i++ {
				ret[i] = s.Index(i).Interface()
			}
			result := checkIntoArrString(input, ret)
			if result.Name == "" {
				continue
			} else {
				return result
			}

		}
	}
	return ObjQueue{Name: ""}

}
