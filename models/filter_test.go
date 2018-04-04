package models

import (
	"reflect"
	"testing"
)

func TestNewFilter(t *testing.T) {
	type args struct {
		name  string
		alias string
		descr string
	}
	tests := []struct {
		name string
		args args
		want Filter
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilter(tt.args.name, tt.args.alias, tt.args.descr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_AddSubject(t *testing.T) {
	type args struct {
		s *Subject
	}
	tests := []struct {
		name string
		f    *Filter
		args args
		want *Filter
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.AddSubject(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter.AddSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}
