package pxbakergo

import (
	"fmt"
	"log"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/rivanjuthani/pxokhttptls"
)

func main() {
	px := NewPerimeterX("", true, false)
	profile := pxokhttptls.PXTLSClientProfile()

	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profile),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		tls_client.WithProxyUrl(""),
	}

	var err error
	px.Client, err = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Fatalln("Failed to create client:", err)
	}

	result := px.SubmitSensor()
	fmt.Println(result)
}
