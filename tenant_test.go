package main

import (
	"reflect"
	"testing"
)

func TestNewTenant(t *testing.T) {
	// tenant := NewTenant("test", "testing-alias", "short description")
}

func TestTenant_AddVRF(t *testing.T) {
	type args struct {
		v *VRF
	}
	tests := []struct {
		name string
		t    *Tenant
		args args
		want *Tenant
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.AddVRF(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tenant.AddVRF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenant_AddBridgeDomain(t *testing.T) {
	type args struct {
		bd *BridgeDomain
	}
	tests := []struct {
		name string
		t    *Tenant
		args args
		want *Tenant
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.AddBridgeDomain(tt.args.bd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tenant.AddBridgeDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenant_AddAppProfile(t *testing.T) {
	type args struct {
		ap *AppProfile
	}
	tests := []struct {
		name string
		t    *Tenant
		args args
		want *Tenant
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.AddAppProfile(tt.args.ap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tenant.AddAppProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenant_AddContract(t *testing.T) {
	type args struct {
		c *Contract
	}
	tests := []struct {
		name string
		t    *Tenant
		args args
		want *Tenant
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.AddContract(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tenant.AddContract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenant_AddFilter(t *testing.T) {
	type args struct {
		f *Filter
	}
	tests := []struct {
		name string
		t    *Tenant
		args args
		want *Tenant
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.AddFilter(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tenant.AddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
