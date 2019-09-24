package plugin

import (
	"reflect"
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
	for i := 1; i < len(fromArg); i++ {
		if _, ok := fromArg[i].(string); ok {
			if input == fromArg[i] {
				// glog.Infoln("input == v", input, fromArg[i])
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
			return checkIntoArrString(input, ret)

		}
	}
	return ObjQueue{Name: ""}
}

// ExistInFromArr ...
func ExistInFromArr(input string, fromArg []interface{}) ObjQueue {
	for _, v := range fromArg {

		// glog.Infoln("dsfds", v, input)
		if _, ok := v.(string); ok {
			if input == v {
				// glog.Infoln("input == v", input, v)
				return ObjQueue{Name: input}
			}
		} else {
			s := reflect.ValueOf(v)
			// glog.Infoln("s.Kind()", s.Kind(), s, input)
			if s.Kind() != reflect.Slice {
				// glog.Infoln("input", input)
				panic("InterfaceSlice() given a non-slice type1")
			}

			ret := make([]interface{}, s.Len())

			for i := 0; i < s.Len(); i++ {
				ret[i] = s.Index(i).Interface()
			}
			result := checkIntoArrString(input, ret)
			if result.Name == "" {

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
