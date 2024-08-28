package pxbakergo

import (
	"reflect"
	"testing"
)

func TestDeviceSensorTimestamp(t *testing.T) {
	testInstance := GetTestInstance()
	px := testInstance.pxCliet

	result := reflect.TypeOf(px.Device.Init_Timestamp).Kind()
	if result != reflect.Int64 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, reflect.Int64)
	}
}
