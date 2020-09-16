package review

import (
	"testing"
)

const DefaultValue = "DefaultValue"

var myMap *MyMap

func init() {
	myMap = &MyMap{
		bucket: make(map[string][]interface{}, 0),
	}
}

func TestMyMap_PutAll(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		args   args
	}{
		{
			"test1",
			args{
				DefaultValue,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myMap.PutAll(tt.args.value)
		})
	}
}

func TestMyMap_Put(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		args   args
	}{
		{
			"test1",
			args{
				"test1",
				"test1",
			},
		},
		{
			"test2",
			args{
				"test2",
				"test2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myMap.Put(tt.args.key, tt.args.value)
		})
	}
}

func TestMyMap_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		want   string
		want1  bool
	}{
		{
			"test1",
			args{
				"test1",
			},
			"test1",
			true,
		},
		{
			"test2",
			args{
				"test21",
			},
			DefaultValue,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := myMap
			got, got1 := m.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("MyMap.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MyMap.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}


