package main

import (
	"context"
	"fmt"

	"github.com/aina-saa/json2pubsub/process"
	"github.com/aina-saa/json2pubsub/version"
	"github.com/alecthomas/kong"
)

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	info := fmt.Sprintf("%s ver. %s\n %s built at %s on %s\n Author: %s", app.Model.Name, version.BuildVersion, version.BuildSha, version.BuildTime, version.BuildHost, version.Author)
	fmt.Println(info)
	app.Exit(0)
	return nil
}

var CLI struct {
	Project string            `help:"Google Cloud Platform project id where the Pub/Sub topics in mappings are located in." required:"true"`
	Mapping map[string]string `help:"Format: VALUE=json.field:my-topic-for-value VALUE2=json.field:my-topic-for-value2" arg:""`
	File    string            `help:"Input file or '-' for stdin" name:"file" type:"existingfile" short:"f" default:"-"`
	Version VersionFlag       `name:"version" help:"Print version information and quit"`
	Quiet   bool              `help:"Be quiet." negatable:"true" default:"false"`
}

func main() {
	ctx := context.Background()

	kong_ctx := kong.Parse(&CLI,
		kong.Name("json2pubsub"),
		kong.Description("Reads JSON object (stream) from file/stdin and routes it/them to GCP Pub/Sub topics."),
	)
	switch kong_ctx.Command() {
	default:
		// we dont have a subcommand so we always drop into default
		process.Process(ctx, CLI.Project, CLI.File, CLI.Mapping, CLI.Quiet)
	}
}

// eof
