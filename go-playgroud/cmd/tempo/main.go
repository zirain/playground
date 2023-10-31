package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-logfmt/logfmt"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/grafana/tempo/pkg/tempopb"
)

func main() {
	host := "172.30.255.202:3100"

	total, err := QueryTraceFromTempo(host, map[string]string{
		"component":    "proxy",
		"service.name": "gateway-conformance-infra-same-namespace",
		// "k8s.cluster.name": "envoy-gateway",
		// "k8s.namespace.name": "envoy-gateway-system",
	})
	if err != nil {
		log.Fatalf("failed to query trace from tempo, err=%v", err)
		return
	}
	log.Printf("total traces in tempo: %d\n", total)
}

func QueryTraceFromTempo(host string, tags map[string]string) (int, error) {
	tagsQueryParam, err := createTagsQueryParam(tags)
	if err != nil {
		return -1, err
	}

	tempoURL := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   "/api/search",
	}
	query := tempoURL.Query()
	query.Add("start", fmt.Sprintf("%d", time.Now().Add(-24*time.Hour).Unix()))
	query.Add("end", fmt.Sprintf("%d", time.Now().Unix()))
	query.Add("tags", tagsQueryParam)
	tempoURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", tempoURL.String(), nil)
	if err != nil {

		return -1, err
	}

	log.Printf("send request to %s\n", tempoURL.String())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}

	if res.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("failed to query tempo, url=%s, status=%s", tempoURL.String(), res.Status)
	}

	tempoResponse := &tempopb.SearchResponse{}
	if err := jsonpb.Unmarshal(res.Body, tempoResponse); err != nil {
		return -1, err
	}

	// TODO: looks like there's some unmarshal issue with gogo/protobuf,
	// but it's fine for now cause we only need the total count.
	total := int(tempoResponse.Metrics.InspectedTraces)
	tracesLen := len(tempoResponse.Traces)
	log.Printf("get response from tempo, InspectedTraces=%d, Traces=%d \n ", total, tracesLen)
	return total, nil
}

// copy from https://github.com/grafana/tempo/blob/c0127c78c368319433c7c67ca8967adbfed2259e/cmd/tempo-query/tempo/plugin.go#L361
func createTagsQueryParam(tags map[string]string) (string, error) {
	tagsBuilder := &strings.Builder{}
	tagsEncoder := logfmt.NewEncoder(tagsBuilder)
	for k, v := range tags {
		err := tagsEncoder.EncodeKeyval(k, v)
		if err != nil {
			return "", err
		}
	}
	return tagsBuilder.String(), nil
}
