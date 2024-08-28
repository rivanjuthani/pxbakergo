package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/google/uuid"
	"github.com/rivanjuthani/pxokhttptls"
)

type PerimeterX struct {
	SDKVersion            string
	PxUuid                string
	AppId                 string
	Tag                   string
	FTag                  string
	P1                    string
	P2                    string
	P3                    string
	PackageName           string
	IsInstantApp          bool
	timestamp             float64
	Device                *PhoneSensor
	client                tls_client.HttpClient
	sensorUrl             string
	DEBUG                 bool
	internalPayloadHeader map[string]interface{}
	internalPayloadP1     map[string]interface{}
	internalPayloadP2     map[string]interface{}
	internalPayloadP3     map[string]interface{}
	internalPayloadP4     map[string]interface{}
	Sid                   string
	Vid                   string
}

func NewPerimeterX(PROXY string, DEBUG bool) *PerimeterX {
	profile := pxokhttptls.PXTLSClientProfile()

	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profile),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		tls_client.WithProxyUrl(PROXY),
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return &PerimeterX{}
	}

	return &PerimeterX{
		SDKVersion:   "v3.2.1",
		PxUuid:       uuid.NewString(),
		AppId:        "PXaOtQIWNf",
		Tag:          "mobile",
		FTag:         "22",
		P1:           "Android",
		P2:           "Chegg Study",
		P3:           "15.15.0",
		PackageName:  "com.chegg",
		IsInstantApp: false,
		timestamp:    float64(time.Now().UnixNano()) / 1e9,
		Device:       NewPhoneSensor(time.Now().UnixNano()),
		client:       client,
		sensorUrl:    fmt.Sprintf("https://collector-%s.perimeterx.net/api/v1/collector/mobile", "pxaotqiwnf"),
		DEBUG:        DEBUG,
	}
}

func (px *PerimeterX) GetContainer(t string, d interface{}) []map[string]interface{} {
	return []map[string]interface{}{
		{"t": t, "d": d},
	}
}

func (px *PerimeterX) TimestampUUID(customCurrentTimeMillis string) string {
	currentTimeMillis := time.Now().UnixNano() / int64(time.Millisecond)

	if customCurrentTimeMillis != "" {
		currentTimeMillis, _ = strconv.ParseInt(customCurrentTimeMillis, 10, 64)
	}

	uuidValue := uuid.New()
	uuidBytes := uuidValue[:]

	for i := 0; i < 8; i++ {
		uuidBytes[i] = byte(currentTimeMillis >> (8 * (7 - i)))
	}

	uuidBytes[8] = 0
	uuidBytes[9] = 0
	uuidBytes[10] = 0
	uuidBytes[11] = 0
	uuidBytes[12] = 0
	uuidBytes[13] = 0
	uuidBytes[14] = 0
	uuidBytes[15] = 1

	modifiedUUID, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		fmt.Println("Error creating UUID:", err)
		return ""
	}

	return modifiedUUID.String()
}

func (px *PerimeterX) HashedID(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	hashValue := fmt.Sprintf("%X", hash.Sum(nil))
	return hashValue
}

func (px *PerimeterX) BuildInitialPayload() string {
	device := px.Device

	px.internalPayloadHeader = map[string]interface{}{"PX330": "new_session"}
	px.internalPayloadP1 = device.BuildPXDevicePayload()

	timestampUUID := px.TimestampUUID("")
	modelEncodedUUID := timestampUUID[:len(device.BuildModel)]

	hashValue := px.HashedID(device.BuildModel + timestampUUID + modelEncodedUUID)

	px.internalPayloadP2 = map[string]interface{}{
		"PX340":  px.SDKVersion,
		"PX342":  px.P3,
		"PX341":  fmt.Sprintf("\"%s\"", px.P2),
		"PX348":  px.PackageName,
		"PX1159": px.IsInstantApp,
		"PX345":  1,
		"PX351":  rand.Intn(41) + 280,
		"PX326":  timestampUUID,
		"PX327":  modelEncodedUUID,
		"PX328":  hashValue,
	}

	px.internalPayloadP3 = map[string]interface{}{
		"PX1208":  "[]",
		"PX21219": "{}",
	}

	builtPayload := MergeMaps(px.internalPayloadHeader, px.internalPayloadP1, px.internalPayloadP2, px.internalPayloadP3)
	payloadBytes := Bencodejson(px.GetContainer("PX315", builtPayload))
	return payloadBytes
}

func (px *PerimeterX) BuildFinalPayload(do ToDoResponse) string {
	px.internalPayloadP4 = px.ParseDo(do)

	builtPayload := MergeMaps(px.internalPayloadHeader, px.internalPayloadP1, px.internalPayloadP2, px.internalPayloadP4, px.internalPayloadP3)
	payloadBytes := Bencodejson(px.GetContainer("PX329", builtPayload))
	return payloadBytes
}

func (px *PerimeterX) ParseDo(do ToDoResponse) map[string]interface{} {
	data := make(map[string]interface{})

	for _, item := range do.Do {
		todo := item
		split := strings.Split(todo, "|")
		doData := split[1:]

		switch split[0] {
		case "sid":
			px.Sid = doData[0]
		case "vid":
			px.Vid = doData[0]
		case "appc":
			if doData[0] == "2" {
				fmt.Println(split)
				data["PX259"] = doData[1]
				data["PX256"] = doData[2]
				data["PX257"] = BitewiseXOR(px.Device.BuildModel, split)
			}
		}
	}
	return data
}

func (px *PerimeterX) SubmitSensor() map[string]interface{} {
	pxPayload := px.BuildInitialPayload()

	data := url.Values{
		"payload": {pxPayload},
		"uuid":    {px.PxUuid},
		"appId":   {px.AppId},
		"tag":     {px.Tag},
		"ftag":    {px.FTag},
		"p3":      {px.P3},
		"p1":      {px.P1},
		// "p2":      {px.P2}, Removed because of bug, added after encoding
	}

	encodedData := data.Encode() + "&p2=" + px.P2

	if px.DEBUG {
		fmt.Println("first payload:", encodedData)
	}

	req, err := http.NewRequest(http.MethodPost, px.sensorUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		log.Println(err)
	}

	req.Header = http.Header{
		"User-Agent":     {"PerimeterX Android SDK/3.2.1"},
		"Accept-Charset": {"UTF-8"},
		"Accept":         {"*/*"},
		"Content-Type":   {"application/x-www-form-urlencoded; charset=utf-8"},
		http.HeaderOrderKey: {
			"user-agent",
			"accept-charset",
			"accept",
			"content-type",
		},
	}

	resp, err := px.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	readBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if px.DEBUG {
		fmt.Println("Response:", string(readBytes))
	}

	var todoList ToDoResponse
	err = json.Unmarshal(readBytes, &todoList)
	if err != nil {
		fmt.Println(err)
	}

	if px.DEBUG {
		fmt.Println("Response Parsed:", todoList)
	}

	if len(todoList.Do) < 3 {
		fmt.Println("error with payload")
	}

	pxPayload = px.BuildFinalPayload(todoList)

	data = url.Values{
		"payload": {pxPayload},
		"uuid":    {px.PxUuid},
		"appId":   {px.AppId},
		"tag":     {px.Tag},
		"ftag":    {px.FTag},
		"sid":     {px.Sid},
		"vid":     {px.Vid},
		"p3":      {px.P3},
		"p1":      {px.P1},
		// "p2":      {px.P2}, Removed because of bug, added after encoding
	}

	encodedData = data.Encode() + "&p2=" + px.P2

	if px.DEBUG {
		fmt.Println("final payload:", encodedData)
	}

	req, err = http.NewRequest(http.MethodPost, px.sensorUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		log.Println(err)
	}

	req.Header = http.Header{
		"User-Agent":     {"PerimeterX Android SDK/3.2.1"},
		"Accept-Charset": {"UTF-8"},
		"Accept":         {"*/*"},
		"Content-Type":   {"application/x-www-form-urlencoded; charset=utf-8"},
		http.HeaderOrderKey: {
			"user-agent",
			"accept-charset",
			"accept",
			"content-type",
		},
	}

	resp, err = px.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	readBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if px.DEBUG {
		fmt.Println("Response:", string(readBytes))
	}

	err = json.Unmarshal(readBytes, &todoList)
	if err != nil {
		fmt.Println(err)
	}

	if px.DEBUG {
		fmt.Println("Response Parsed:", todoList)
	}

	bakedData := strings.Split(todoList.Do[0], "|")

	if bakedData[0] != "bake" {
		fmt.Println("error with baking!")
	}

	expiry, _ := strconv.Atoi(bakedData[2])
	return map[string]interface{}{
		"cookie": bakedData[3],
		"expiry": expiry,
	}
}

func (px *PerimeterX) PxHello(pxuuid string, str2 string) string {
	var bytes1 []byte
	if pxuuid == "" {
		bytes1 = []byte(px.PxUuid)
	} else {
		bytes1 = []byte(pxuuid)
	}
	bytes2 := []byte(str2)
	length := len(bytes1)
	bArr := make([]byte, length)

	for i := 0; i < length; i++ {
		bArr[i] = bytes1[i] ^ bytes2[i%len(bytes2)]
	}

	encoded := base64.StdEncoding.EncodeToString(bArr)
	return encoded
}
