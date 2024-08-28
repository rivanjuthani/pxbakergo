package pxbakergo

/* TEST INSTANCE COPYPASTA -- START */

type TestInstance struct {
	pxCliet *PerimeterX
	toDo    ToDoResponse
}

func GetTestInstance() TestInstance {
	px := NewPerimeterX("", true, true)
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

/* TEST INSTANCE COPYPASTA -- END */
