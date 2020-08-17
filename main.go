package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	acceptedLevels = map[string]struct{}{
		"patch": struct{}{},
		"minor": struct{}{},
		"major": struct{}{},
	}
)

func main() {
	var level string
	flag.StringVar(&level, "level", "patch", "which level to bump: patch, minor or major")
	flag.Parse()

	if _, ok := acceptedLevels[level]; !ok {
		fmt.Fprintln(os.Stderr, "unknown level provided:", level)
		os.Exit(1)
		return
	}

	var version string
	fmt.Scanln(&version)

	if version == "" {
		fmt.Fprintln(os.Stdout, "1.0.0")
		return
	}

	var bumpedVersion strings.Builder
	if version[0] == 'v' {
		bumpedVersion.WriteString("v")
		version = version[1:]
	}

	versionParts := strings.Split(version, ".")

	var err error
	switch level {
	case "patch":
		var patchVersion int64
		patchVersion, err = strconv.ParseInt(versionParts[2], 10, 32)
		patchVersion++
		versionParts[2] = strconv.Itoa(int(patchVersion))
	case "minor":
		var minorVersion int64
		minorVersion, err = strconv.ParseInt(versionParts[1], 10, 32)
		minorVersion++
		versionParts[1] = strconv.Itoa(int(minorVersion))
		versionParts[2] = "0"
	case "major":
		var majorVersion int64
		majorVersion, err = strconv.ParseInt(versionParts[0], 10, 32)
		majorVersion++
		versionParts[0] = strconv.Itoa(int(majorVersion))
		versionParts[1] = "0"
		versionParts[2] = "0"
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		return
	}

	bumpedVersion.WriteString(strings.Join(versionParts, "."))

	fmt.Fprintln(os.Stdout, bumpedVersion.String())
}
