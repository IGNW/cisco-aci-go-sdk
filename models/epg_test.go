package models

import (
	"reflect"
	"testing"
)

func TestNewEPG(t *testing.T) {
	type args struct {
		name  string
		alias string
		descr string
	}
	tests := []struct {
		name string
		args args
		want EPG
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEPG(tt.args.name, tt.args.alias, tt.args.descr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEPG() = %v, want %v", got, tt.want)
			}
		})
	}
}
