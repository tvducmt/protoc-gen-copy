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
		if _, ok := fromArg[i].(string); ok {
			if input == fromArg[i] {
				if val, ok := fromArg[0].(ObjQueue); ok {
					// glog.Infoln("input == v", val)
					return val
				}
				return ObjQueue{} //fromArg[0].(string)
			}
		}
	}

	//duyet cac phan tu trong mang
	for i := 1; i < len(fromArg); i++ {
		s := reflect.ValueOf(fromArg[i])
		if s.Kind() != reflect.Slice {
			continue
			// panic("InterfaceSlice() given a non-slice type")
		}

		ret := make([]interface{}, s.Len())

		for i := 0; i < s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
		}
		if input == ".ItemCount" {
			glog.Infoln("len(ret)", ret, "  len: ", len(ret))
		}
		return checkIntoArrString(input, ret)

		// }
	}
	return ObjQueue{Name: ""}
}

// ExistInFromArr ...
func ExistInFromArr(input string, fromArg []interface{}) ObjQueue {
	// if input == ".ItemCount" {
	// 	glog.Infoln("lenfromArg", len(fromArg))
	// }
	for _, v := range fromArg {

		// glog.Infoln("input", input)
		if _, ok := v.(string); ok {
			if input == v {
				// glog.Infoln("input == v", input, v)
				return ObjQueue{Name: input}
			}
		} else {
			s := reflect.ValueOf(v)
			if s.Kind() != reflect.Slice {
				glog.Infoln("s.Kind()", s.Kind(), s, input)
				// glog.Infoln("input", input)
				panic("InterfaceSlice() given a non-slice type1")
			}

			ret := make([]interface{}, s.Len())

			for i := 0; i < s.Len(); i++ {
				ret[i] = s.Index(i).Interface()
			}
			result := checkIntoArrString(input, ret)
			if result.Name == "" {
				// if input == ".ItemCount" {
				// 	glog.Infoln("dsfgfsgsfgsfgsfgfsgfsg")
				// }
				// glog.Infoln("dsfgfsgsfgsfgsfgfsgfsg")
				continue
			} else {
				return result
			}
			// return ExistInFromArr(input, ret)

		}
	}
	return ObjQueue{Name: ""}

}
