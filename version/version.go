package version

// these variables should be set during compile them with -ldflags
//  https://polyverse.com/blog/how-to-embed-versioning-information-in-go-applications-f76e2579b572/

var (
	BuildVersion string = ""
	BuildTime    string = ""
	BuildSha     string = ""
	BuildHost    string = ""
	Author       string = "Antti Peltonen <antti.peltonen@iki.fi>"
)

// eof
