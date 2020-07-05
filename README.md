- *Go语言QQ群: 102319854, 1055927514*
- *凹语言(凹读音“Wa”)(The Wa Programming Language): https://github.com/wa-lang/wa*

----

# libconfd

[![Build Status](https://travis-ci.org/chai2010/libconfd.svg)](https://travis-ci.org/chai2010/libconfd)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/libconfd)](https://goreportcard.com/report/github.com/chai2010/libconfd)
[![GoDoc](https://godoc.org/github.com/chai2010/libconfd?status.svg)](https://godoc.org/github.com/chai2010/libconfd)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/chai2010/libconfd/blob/master/LICENSE)

mini confd lib, based on [confd](https://github.com/kelseyhightower/confd)/[memkv](https://github.com/kelseyhightower/memkv)/[secconf](https://github.com/xordataexchange/crypt)/[logger](https://github.com/chai2010/logger).


## Example

```go
package main

import (
	"github.com/chai2010/libconfd"
)

func main() {
	cfg := libconfd.MustLoadConfig("./confd.toml")

	backendConfig := libconfd.MustLoadBackendConfig("./confd-backend.toml")
	backendClient := libconfd.MustNewBackendClient(backendConfig)

	libconfd.NewProcessor().Run(cfg, backendClient)
}
```

## miniconfd (only support toml/etcd backend)

```
$ go run miniconfd.go -h
```

See [etcdv3/miniconfd-etcdv3/miniconfd.go](etcdv3/miniconfd-etcdv3/miniconfd.go)
