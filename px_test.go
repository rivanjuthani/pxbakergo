package main

import (
	"strings"
	"testing"
)

type TestInstance struct {
	pxCliet *PerimeterX
	toDo    ToDoResponse
}

func getTestInstance() TestInstance {
	px := NewPerimeterX("", true)
	px.PxUuid = "9d92d005-5e00-11ef-90d2-1568ec2dc4f9"
	px.Device.BuildModel = "SM-S908N"

	todo := ToDoResponse{
		Do: []string{"sid|9e5f9442-5e00-11ef-b865-16b6a0e17b04", "vid|9e5f89ba-5e00-11ef-b865-3d46c922ea5a|31536000|false", "appc|1|1724054237200|a3120a645bc5fb1b1dcbb1008a1c0c89f5a470a5add56836939aa2b85dab1e1c|9e61df80-5e00-11ef-bf75-3114a1687ec5", "appc|2|1724054237200|330f82a5a57c7d7570c612ea178609ee4d6c9e36e12e8a0f92ac5532026b2a12,330f82a5a57c7d7570c612ea178609ee4d6c9e36e12e8a0f92ac5532026b2a12|1016|1126|2656|1173|2609|4037", "ipd|false"},
	}

	return TestInstance{
		pxCliet: px,
		toDo:    todo,
	}
}

func TestPxHello(t *testing.T) {
	testInstance := getTestInstance()
	px := testInstance.pxCliet

	result := px.PxHello("9d92d005-5e00-11ef-90d2-1568ec2dc4f9", "2")
	if result != "C1YLAFYCAgcfB1cCAh8DA1dUHwsCVgAfAwcECldRAFZRBlQL" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "C1YLAFYCAgcfB1cCAh8DA1dUHwsCVgAfAwcECldRAFZRBlQL")
	}
}

func TestBitwiseXOR(t *testing.T) {
	testInstance := getTestInstance()
	px := testInstance.pxCliet

	result := BitewiseXOR(px.Device.BuildModel, strings.Split("appc|2|1723615917295|ae848523a7ba5c53dd3704c8a1108dc0eab12cfb4ad2fb00f173dc343daeab59,ae848523a7ba5c53dd3704c8a1108dc0eab12cfb4ad2fb00f173dc343daeab59|1466|679|686|3843|2535|1073", "|"))
	if result != 1063303278 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1063303278)
	}
}

func TestTimestampUUID(t *testing.T) {
	testInstance := getTestInstance()
	px := testInstance.pxCliet

	result := px.TimestampUUID("1724732402055")

	if result != "00000191-920f-e987-0000-000000000001" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "00000191-920f-e987-0000-000000000001")
	}
}

func TestHashedID(t *testing.T) {
	testInstance := getTestInstance()
	px := testInstance.pxCliet

	timestampUUID := "00000191-69a3-b1ff-0000-000000000001"
	modelEncodedUUID := timestampUUID[:len(px.Device.BuildModel)]

	result := px.HashedID(px.Device.BuildModel + timestampUUID + modelEncodedUUID)
	if result != "1E884FACDB77E5D0A41C9F87DAA032ADDEF2FF3D" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "1E884FACDB77E5D0A41C9F87DAA032ADDEF2FF3D")
	}
}
