package cage

import (
	"reflect"
	"testing"
)

func TestNewContract(t *testing.T) {
	type args struct {
		name  string
		alias string
		descr string
	}
	tests := []struct {
		name string
		args args
		want Contract
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContract(tt.args.name, tt.args.alias, tt.args.descr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContract_AddSubject(t *testing.T) {
	type args struct {
		s *Subject
	}
	tests := []struct {
		name string
		c    *Contract
		args args
		want *Contract
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.AddSubject(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contract.AddSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContract_AddEPG(t *testing.T) {
	type args struct {
		e *EPG
	}
	tests := []struct {
		name string
		c    *Contract
		args args
		want *Contract
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.AddEPG(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contract.AddEPG() = %v, want %v", got, tt.want)
			}
		})
	}
}
