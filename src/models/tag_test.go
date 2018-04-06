// +build unit

package models

import (
	"reflect"
	"testing"
)

func TestNewTag(t *testing.T) {
	type args struct {
		tagName string
	}
	tests := []struct {
		name string
		args args
		want Tag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTag(tt.args.tagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag_AsPayLoadFormat(t *testing.T) {
	tests := []struct {
		name string
		t    *Tag
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.AsPayLoadFormat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tag.AsPayLoadFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
