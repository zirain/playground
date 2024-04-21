package jsonpatch_test

import (
	"testing"

	jsonpatchv5 "github.com/evanphx/json-patch/v5"
	"github.com/stretchr/testify/require"
)

func TestPatch(t *testing.T) {
	input := `[{ "op": "add", "path": "/filter_chains/0/filters/0/typed_config/preserve_external_request_id", "value": true }]`
	patchObj, err := jsonpatchv5.DecodePatch([]byte(input))
	require.NoError(t, err)
	resourceJSON := `{"name":"first-listener","address":{"socket_address":{"address":"0.0.0.0","port_value":10080}},"filter_chains":[{"filters":[{"name":"envoy.filters.network.http_connection_manager","typed_config":{"@type":"type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager","stat_prefix":"https","rds":{"config_source":{"ads":{},"resource_api_version":"V3"},"route_config_name":"first-listener"},"http_filters":[{"name":"envoy.filters.http.router","typed_config":{"@type":"type.googleapis.com/envoy.extensions.filters.http.router.v3.Router","suppress_envoy_headers":true}}],"common_http_protocol_options":{"headers_with_underscores_action":"REJECT_REQUEST"},"http2_protocol_options":{"max_concurrent_streams":100,"initial_stream_window_size":65536,"initial_connection_window_size":1048576},"server_header_transformation":"PASS_THROUGH","use_remote_address":true,"normalize_path":true,"merge_slashes":true,"path_with_escaped_slashes_action":"UNESCAPE_AND_REDIRECT"}}],"transport_socket":{"name":"envoy.transport_sockets.tls","typed_config":{"@type":"type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext","common_tls_context":{"tls_certificate_sds_secret_configs":[{"name":"secret-1","sds_config":{"ads":{},"resource_api_version":"V3"}},{"name":"secret-2","sds_config":{"ads":{},"resource_api_version":"V3"}}],"alpn_protocols":["h2","http/1.1"]}}}}],"per_connection_buffer_limit_bytes":32768,"drain_type":"MODIFY_ONLY"}`
	// Apply patch
	opts := jsonpatchv5.NewApplyOptions()
	opts.EnsurePathExistsOnAdd = true
	modifiedJSON, err := patchObj.ApplyWithOptions([]byte(resourceJSON), opts)
	require.NoError(t, err)
	expected := `{"name":"first-listener","address":{"socket_address":{"address":"0.0.0.0","port_value":10080}},"filter_chains":[{"filters":[{"name":"envoy.filters.network.http_connection_manager","typed_config":{"@type":"type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager","stat_prefix":"https","rds":{"config_source":{"ads":{},"resource_api_version":"V3"},"route_config_name":"first-listener"},"http_filters":[{"name":"envoy.filters.http.router","typed_config":{"@type":"type.googleapis.com/envoy.extensions.filters.http.router.v3.Router","suppress_envoy_headers":true}}],"common_http_protocol_options":{"headers_with_underscores_action":"REJECT_REQUEST"},"http2_protocol_options":{"max_concurrent_streams":100,"initial_stream_window_size":65536,"initial_connection_window_size":1048576},"server_header_transformation":"PASS_THROUGH","use_remote_address":true,"normalize_path":true,"merge_slashes":true,"path_with_escaped_slashes_action":"UNESCAPE_AND_REDIRECT","preserve_external_request_id":true}}],"transport_socket":{"name":"envoy.transport_sockets.tls","typed_config":{"@type":"type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext","common_tls_context":{"tls_certificate_sds_secret_configs":[{"name":"secret-1","sds_config":{"ads":{},"resource_api_version":"V3"}},{"name":"secret-2","sds_config":{"ads":{},"resource_api_version":"V3"}}],"alpn_protocols":["h2","http/1.1"]}}}}],"per_connection_buffer_limit_bytes":32768,"drain_type":"MODIFY_ONLY"}`
	require.Equal(t, expected, string(modifiedJSON))
}
