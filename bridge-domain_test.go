package cage

import (
	"reflect"
	"testing"
)

func TestNewBridgeDomain(t *testing.T) {
	type args struct {
		name  string
		alias string
		descr string
	}
	tests := []struct {
		name string
		args args
		want BridgeDomain
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBridgeDomain(tt.args.name, tt.args.alias, tt.args.descr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBridgeDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBridgeDomain_AddSubnet(t *testing.T) {
	type args struct {
		s *Subnet
	}
	tests := []struct {
		name string
		bd   *BridgeDomain
		args args
		want *BridgeDomain
	}{
	// @TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bd.AddSubnet(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BridgeDomain.AddSubnet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBridgeDomain_AddEPG(t *testing.T) {
	type args struct {
		e *EPG
	}
	tests := []struct {
		name string
		bd   *BridgeDomain
		args args
		want *BridgeDomain
	}{
	// @TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bd.AddEPG(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BridgeDomain.AddEPG() = %v, want %v", got, tt.want)
			}
		})
	}
}
