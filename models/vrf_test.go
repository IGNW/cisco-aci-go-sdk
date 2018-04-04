package models

import (
	"reflect"
	"testing"
)

func TestNewVRF(t *testing.T) {
	type args struct {
		name  string
		alias string
		descr string
	}
	tests := []struct {
		name string
		args args
		want VRF
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVRF(tt.args.name, tt.args.alias, tt.args.descr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVRF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVRF_AddBridgeDomain(t *testing.T) {
	type args struct {
		bd *BridgeDomain
	}
	tests := []struct {
		name string
		v    *VRF
		args args
		want *VRF
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.AddBridgeDomain(tt.args.bd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VRF.AddBridgeDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
