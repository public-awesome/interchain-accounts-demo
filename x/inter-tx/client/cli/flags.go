package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// The connection end identifier on the controller chain
	FlagConnectionID = "connection-id"
	// The controller chain channel version
	FlagVersion = "version"
	// The packet timeout period
	FlagTimeout = "timeout"
)

// common flagsets to add to various functions
var (
	fsConnectionID = flag.NewFlagSet("", flag.ContinueOnError)
	fsTimeout      = flag.NewFlagSet("", flag.ContinueOnError)
	fsVersion      = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsConnectionID.String(FlagConnectionID, "", "Connection ID")
	fsTimeout.String(FlagTimeout, "1h", "Timeout")
	fsVersion.String(FlagVersion, "", "Version")
}
