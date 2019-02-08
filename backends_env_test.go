// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package libconfd

import (
	"os"
	"strings"
)

// tEnvClient provides a shell for the env client
type tEnvClient struct{}

func init() {
	RegisterBackendClient(
		(*tEnvClient)(nil).Type(),
		func(cfg *BackendConfig) (BackendClient, error) {
			p := tNewEnvClient()
			return p, nil
		},
	)
}

// tNewEnvClient returns a new client
func tNewEnvClient() BackendClient {
	return new(tEnvClient)
}

func (_ *tEnvClient) Type() string {
	return "libconfd-backend-internal-env"
}

func (_ *tEnvClient) Close() error {
	return nil
}

func (_ *tEnvClient) WatchEnabled() bool {
	return false
}

// GetValues queries the environment for keys
func (_ *tEnvClient) GetValues(keys []string) (map[string]string, error) {
	allEnvVars := os.Environ()
	envMap := make(map[string]string)
	for _, e := range allEnvVars {
		index := strings.Index(e, "=")
		envMap[e[:index]] = e[index+1:]
	}

	vars := make(map[string]string)

	transform := func(key string) string {
		var replacer = strings.NewReplacer("/", "_")

		k := strings.TrimPrefix(key, "/")
		return strings.ToUpper(replacer.Replace(k))
	}
	clean := func(key string) string {
		var cleanReplacer = strings.NewReplacer("_", "/")

		newKey := "/" + key
		return cleanReplacer.Replace(strings.ToLower(newKey))
	}

	for _, key := range keys {
		k := transform(key)
		for envKey, envValue := range envMap {
			if strings.HasPrefix(envKey, k) {
				vars[clean(envKey)] = envValue
			}
		}
	}

	return vars, nil
}

func (_ *tEnvClient) WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error) {
	<-stopChan
	return 0, nil
}
