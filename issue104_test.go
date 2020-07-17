package mergo

import (
	"fmt"
	"reflect"
	"testing"
)

type Record struct {
	Data    map[string]interface{}
	Mapping map[string]string
}

func StructToRecord(in interface{}) *Record {
	rec := Record{}
	rec.Data = make(map[string]interface{})
	rec.Mapping = make(map[string]string)
	typ := reflect.TypeOf(in)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		dbFieldName := field.Tag.Get("db")

		fmt.Printf("%d %v, tags: %v\n", i, field.Name, dbFieldName)
		if dbFieldName != "" {
			rec.Mapping[field.Name] = dbFieldName
		}
	}

	Map(&rec.Data, in)
	return &rec
}

func TestStructToRecord(t *testing.T) {
	type A struct {
		Name string `json:"name" db:"name"`
		CIDR string `json:"cidr" db:"cidr"`
	}
	type Record struct {
		Data    map[string]interface{}
		Mapping map[string]string
	}
	a := A{Name: "David", CIDR: "10.0.0.0/8"}
	rec := StructToRecord(a)
	fmt.Printf("rec: %+v\n", rec)
	if len(rec.Mapping) < 2 {
		t.Fatalf("struct to record failed, no mapping, struct missing tags?, rec: %+v, a: %+v ", rec, a)
	}
}