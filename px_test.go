package pxbakergo

import (
	"strings"
	"testing"
)

func TestPxHello(t *testing.T) {
	testInstance := GetTestInstance()
	px := testInstance.pxCliet

	result := px.PxHello("9d92d005-5e00-11ef-90d2-1568ec2dc4f9", "2")
	if result != "C1YLAFYCAgcfB1cCAh8DA1dUHwsCVgAfAwcECldRAFZRBlQL" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "C1YLAFYCAgcfB1cCAh8DA1dUHwsCVgAfAwcECldRAFZRBlQL")
	}
}

func TestBitwiseXOR(t *testing.T) {
	testInstance := GetTestInstance()
	px := testInstance.pxCliet

	result := BitewiseXOR(px.Device.BuildModel, strings.Split("appc|2|1723615917295|ae848523a7ba5c53dd3704c8a1108dc0eab12cfb4ad2fb00f173dc343daeab59,ae848523a7ba5c53dd3704c8a1108dc0eab12cfb4ad2fb00f173dc343daeab59|1466|679|686|3843|2535|1073", "|"))
	if result != 1063303278 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1063303278)
	}
}

func TestTimestampUUID(t *testing.T) {
	testInstance := GetTestInstance()
	px := testInstance.pxCliet

	result := px.TimestampUUID("1724732402055")

	if result != "00000191-920f-e987-0000-000000000001" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "00000191-920f-e987-0000-000000000001")
	}
}

func TestHashedID(t *testing.T) {
	testInstance := GetTestInstance()
	px := testInstance.pxCliet

	timestampUUID := "00000191-69a3-b1ff-0000-000000000001"
	modelEncodedUUID := timestampUUID[:len(px.Device.BuildModel)]

	result := px.HashedID(px.Device.BuildModel + timestampUUID + modelEncodedUUID)
	if result != "1E884FACDB77E5D0A41C9F87DAA032ADDEF2FF3D" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "1E884FACDB77E5D0A41C9F87DAA032ADDEF2FF3D")
	}
}
