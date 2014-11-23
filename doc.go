// Copyright (c) 2014, Tobias Weingartner
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package vendor provides methods to register vendor'd code information as
well as build information for a golang binary.

Within each vendor'd copy of a package a single file is created, usually
called '$dir/vendor.go' within the package of the vendor'd code that looks
something like:

	package vendor-code

	import "github.com/weingart/vendor"

	func init() {
		vendor.Add(&vendor.Info{
			"git-repo": "github.com/vendor/repo",
			"git-sha":  "9e69b5e2e8d4d042f67d9b24f0f69ffb3ab35687",
			"git-tag":  "sometag",
		})
	}

Within the packe github.com/weingart/vendor, there are variables that can
set during the build via the -ldflags flags.  These variables can be
overridden using something like:

	SHA=$(git rev-parse HEAD)
	go install -ldflags "-X github.com/weingart/vendor.buildSHA ${SHA}" example

The following variables are available to be set:

	buildSHA
	buildTag
	buildUser
	buildTime
	buildComment

In order to be able to introspect the values within a running server, the
information is registered with the expvar package.  They can also be retrieved
by using the appropriate vendor.Get*() functions.  For example:

	package main

	import (
		_ "expvar"
		"fmt"
		"github.com/weingart/vendor"
		"net/http"
	)

	func main() {
		fmt.Println(vendor.GetInfo())
		http.ListenAndServe(":8080", nil)
	}

The information can be found by going to: http://localhost:8080/debug/vars

*/
package vendor
