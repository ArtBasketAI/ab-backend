package main

import (
	"fmt"
	"os"

	"github.com/artbasketai/ab-backend/cmd/server/realmain"
)

// version is git commit at the time of build of this service binary. version is overwritten by -ldflags
// while building.
//nolint:varcheck,deadcode,unused // will be later used for logging
var version string

// serviceName is service name of this microservices. serviceName
// is overwritten using -ldflags while building.
//nolint:varcheck,deadcode,unused // will be later used for logging
var serviceName string

func main() {
	_, _ = fmt.Fprintf(os.Stdout, "[INFO] Starting service\n")
	os.Exit(realmain.Execute())
}
