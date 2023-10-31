package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/common/expfmt"
)

func main() {
	_, body, err := HTTPGet(fmt.Sprintf("http://127.0.0.1:%d/stats/prometheus", 9901))
	if err != nil {
		fmt.Printf("get metric error: %v\n", err)
		return
	}
	reader := strings.NewReader(body)
	_, err = (&expfmt.TextParser{}).TextToMetricFamilies(reader)
	if err != nil {
		fmt.Printf("parse metric error: %v body: %s\n", err, body)
	}
}

func HTTPGet(url string) (code int, respBody string, err error) {
	log.Println("HTTP GET", url)
	client := &http.Client{Timeout: time.Minute}
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	respBody = string(body)
	code = resp.StatusCode
	return code, respBody, nil
}
