package main

import (
	"math/rand"
	"strconv"
)

type PhoneSensor struct {
	AndroidId                  string
	ScreenWidth                int
	ScreenHeight               int
	ScreenBrightness           int
	HasSimCard                 bool
	AndroidSDK_INT             string
	OSKernelVersion            string
	BuildModel                 string
	BuildBrand                 string
	BoardBrand                 string
	Init_Timestamp             int64
	OSName                     string
	HasGPS                     bool
	HasGyroscope               bool
	HasAccelerometer           bool
	HasEthernet                bool
	HasTouchscreen             bool
	HasNFC                     bool
	HasWifi                    bool
	HardcodedTrue              string
	IsRooted                   string
	AllowTouchDetection        string
	AllowDeviceMotionDetection string
	NetworkStatus              string
	NetworkCarrier             string
	Locale                     string
	BatteryHealth              string
	BatteryPowerType           string
	BatteryStatus              string
	BatteryType                string
	BatteryTemp                float64
	BatteryVoltage             float64
}

func NewPhoneSensor(initTS int64) *PhoneSensor {
	return &PhoneSensor{
		AndroidId:                  RandomHex(16),
		ScreenWidth:                900,
		ScreenHeight:               1600,
		ScreenBrightness:           rand.Intn(51) + 100,
		HasSimCard:                 true,
		AndroidSDK_INT:             "28",
		OSKernelVersion:            "4.4.146",
		BuildModel:                 "SM-S908N",
		BuildBrand:                 "samsung",
		BoardBrand:                 "gracelte",
		Init_Timestamp:             initTS,
		OSName:                     "Android",
		HasGPS:                     true,
		HasGyroscope:               true,
		HasAccelerometer:           true,
		HasEthernet:                true,
		HasTouchscreen:             true,
		HasNFC:                     false,
		HasWifi:                    true,
		HardcodedTrue:              strconv.FormatBool(true),
		IsRooted:                   strconv.FormatBool(false),
		AllowTouchDetection:        strconv.FormatBool(true),
		AllowDeviceMotionDetection: strconv.FormatBool(true),
		NetworkStatus:              "WiFi",
		NetworkCarrier:             "Verizon",
		Locale:                     "[\"en_US\"]",
		BatteryHealth:              "good",
		BatteryPowerType:           "AC",
		BatteryStatus:              "charging",
		BatteryType:                "Li-ion",
		BatteryTemp:                Round(rand.Float64()*(35-32)+32, 1),
		BatteryVoltage:             3.7,
	}
}

func (ps *PhoneSensor) BuildPXDevicePayload() map[string]interface{} {
	return map[string]interface{}{
		"PX1214":  ps.AndroidId,
		"PX91":    ps.ScreenWidth,
		"PX92":    ps.ScreenHeight,
		"PX21215": ps.ScreenBrightness,
		"PX316":   ps.HasSimCard,
		"PX318":   ps.AndroidSDK_INT,
		"PX319":   ps.OSKernelVersion,
		"PX320":   ps.BuildModel,
		"PX339":   ps.BuildBrand,
		"PX321":   ps.BoardBrand,
		"PX323":   ps.Init_Timestamp,
		"PX322":   ps.OSName,
		"PX337":   ps.HasGPS,
		"PX336":   ps.HasGyroscope,
		"PX335":   ps.HasAccelerometer,
		"PX334":   ps.HasEthernet,
		"PX333":   ps.HasTouchscreen,
		"PX331":   ps.HasNFC,
		"PX332":   ps.HasWifi,
		"PX421":   ps.HardcodedTrue,
		"PX442":   ps.IsRooted,

		// Included here for structure
		"PX21218": "[]",
		"PX21217": "[]",

		"PX21224": ps.AllowTouchDetection,
		"PX21221": ps.AllowDeviceMotionDetection,
		"PX317":   ps.NetworkStatus,
		"PX344":   ps.NetworkCarrier,
		"PX347":   ps.Locale,

		// Included here for structure
		"PX343": "Unknown",
		"PX415": 100, // FirebaseAnalytics.Param.LEVEL

		"PX413": ps.BatteryHealth,
		"PX416": ps.BatteryPowerType,
		"PX414": ps.BatteryStatus,
		"PX419": ps.BatteryType,
		"PX418": ps.BatteryTemp,
		"PX420": ps.BatteryVoltage,
	}
}
