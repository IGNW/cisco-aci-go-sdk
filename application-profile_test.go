package main

import (
	"reflect"
	"testing"
)

func TestNewAppProfile(t *testing.T) {
	type args struct {
		name  string
		alias string
		descr string
	}
	tests := []struct {
		name string
		args args
		want AppProfile
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAppProfile(tt.args.name, tt.args.alias, tt.args.descr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAppProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppProfile_AddEPG(t *testing.T) {
	type args struct {
		e *EPG
	}
	tests := []struct {
		name string
		ap   *AppProfile
		args args
		want *AppProfile
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ap.AddEPG(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppProfile.AddEPG() = %v, want %v", got, tt.want)
			}
		})
	}
}
