package vendor

import (
	"expvar"
)

// These vendor package level variables are here to allow build script to
// modify them via:
//		go install -ldflags \
//			'-X github.com/weingart/vendor.buildUser ${USER}' \
//		...
// These are all strings and default to "".
var (
	buildSHA     string = "" // SHA (or build revision)
	buildTag     string = "" // Tag of this build
	buildUser    string = "" // User that built the artifact
	buildTime    string = "" // Unix time since EPOCH (string, in seconds)
	buildComment string = "" // Random comment
)

type BuildInfo struct {
	SHA     string `json:"sha"`
	Tag     string `json:"tag"`
	User    string `json:"user"`
	Time    string `json:"time"`
	Comment string `json:"comment"`
}

func GetBuildInfo() *BuildInfo {
	return &BuildInfo{
		SHA:     buildSHA,
		Tag:     buildTag,
		User:    buildUser,
		Time:    buildTime,
		Comment: buildComment,
	}
}

func init() {
	expvar.Publish("build", expvar.Func(func() interface{} {
		return GetBuildInfo()
	}))
}
