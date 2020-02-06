package main

import (
	"testing"
)

func TestProviderMapping(t *testing.T) {
	provider := Provider()
	if _, ok := provider.DataSourcesMap["assert_test"]; !ok {
		t.Error("\"assert_test\" should be included in the DataSourcesMap")
	}
}
