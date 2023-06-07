package controllers

import (
	"github.com/revel/revel"
	"reflect"
	"testing"
)

func TestApp_Index(t *testing.T) {
	type fields struct {
		Controller *revel.Controller
	}
	tests := []struct {
		name   string
		fields fields
		want   revel.Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := App{
				Controller: tt.fields.Controller,
			}
			if got := c.Index(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}
