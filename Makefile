# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

default:

test:
	go vet ./...
	go fmt ./...
	go test ./...

dev:
	go run miniconfd.go

clean:
