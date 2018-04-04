package models

import (
	"reflect"
	"testing"

	"github.com/Jeffail/gabs"
)

func TestResourceAttributes_GetAPIPayload(t *testing.T) {
	tests := []struct {
		name string
		r    *ResourceAttributes
		want *gabs.Container
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetAPIPayload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceAttributes.GetAPIPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceAttributes_CreateDefaultPayload(t *testing.T) {
	tests := []struct {
		name string
		r    *ResourceAttributes
		want *gabs.Container
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.CreateDefaultPayload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceAttributes.CreateDefaultPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceAttributes_CreateEmptyJSONContainer(t *testing.T) {
	tests := []struct {
		name    string
		r       *ResourceAttributes
		want    *gabs.Container
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CreateEmptyJSONContainer()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceAttributes.CreateEmptyJSONContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceAttributes.CreateEmptyJSONContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceAttributes_AddDefaultPropsToPayload(t *testing.T) {
	type args struct {
		data *gabs.Container
	}
	tests := []struct {
		name string
		r    *ResourceAttributes
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.AddDefaultPropsToPayload(tt.args.data)
		})
	}
}

func TestResourceAttributes_AddTagsToPayload(t *testing.T) {
	type args struct {
		data *gabs.Container
	}
	tests := []struct {
		name string
		r    *ResourceAttributes
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.AddTagsToPayload(tt.args.data)
		})
	}
}

func TestResourceAttributes_AddTag(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		r    *ResourceAttributes
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.AddTag(tt.args.name)
		})
	}
}

func TestResourceAttributes_SetParent(t *testing.T) {
	type args struct {
		parent ResourceInterface
	}
	tests := []struct {
		name string
		r    *ResourceAttributes
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.SetParent(tt.args.parent)
		})
	}
}
