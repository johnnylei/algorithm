package review

import (
	"testing"
)

var myStack *MyStack

func init() {
	myStack = &MyStack{
		length: 0,
	}
}

func TestMyStack_Push(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		args   args
	}{
		{
			"test1",
			args{
				value: 2,
			},
		},
		{
			"test2",
			args{
				value: 1,
			},
		},
		{
			"test2",
			args{
				value: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myStack.Push(tt.args.value)
		})
	}
}

func TestMyStack_Pop(t *testing.T) {
	tests := []struct {
		name   string
		want   int
	}{
		{
			"test1",
			2,
		},
		{
			"test2",
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myStack.Pop(); got != tt.want {
				t.Errorf("MyStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStack_GetMin(t *testing.T) {
	tests := []struct {
		name   string
		want   int
	}{
		{
			"test1",
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myStack.GetMin(); got != tt.want {
				t.Errorf("MyStack.GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
