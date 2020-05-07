package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
)

const (
	outFile         = "/tmp/veyil.out"
	envZipCode      = "VEYIL_ZIP_CODE"
	envAPIKey       = "VEYIL_API_KEY"
	defaultInterval = 0
)

type Veyil struct {
	zipCode  uint
	interval uint
	apiKey   string
	outFile  string
}

// Parse user-given arguments in to an instance of Veyil and return a
// pointer to it.
func parseArgs(progname string, args []string) (*Veyil, string, error) {
	var v Veyil
	var buf bytes.Buffer

	progname = path.Base(progname)

	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	flags.SetOutput(&buf)

	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s -- weather data in terminal\n", progname)
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", progname)
		flags.PrintDefaults()
	}

	flags.UintVar(
		&v.zipCode,
		"z",
		0,
		"`ZIP-CODE` to fetch weather data for",
	)
	flags.UintVar(
		&v.interval,
		"i",
		defaultInterval,
		"fetch weather data every `INTERVAL` seconds",
	)
	flags.StringVar(
		&v.apiKey,
		"k",
		"",
		"OpenWeatherMap `API-KEY`",
	)
	flags.StringVar(
		&v.outFile,
		"f",
		outFile,
		"write weather data to `FILE`",
	)

	err := flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}

	// If ZIP code is not specified in cmd line, try fetching it from
	// environ variable.
	if v.zipCode == 0 {
		zip := os.Getenv(envZipCode)
		if len(zip) == 0 {
			return nil,
				buf.String(),
				fmt.Errorf("-z ZIP-CODE or %s environ var is required",
					envZipCode)
		}
		zip64, err := strconv.ParseUint(zip, 10, 32)
		if err != nil {
			return nil,
				buf.String(),
				fmt.Errorf("invalid ZIP code, %s, in environ var %s",
					zip, envZipCode)
		}
		v.zipCode = uint(zip64)
	}

	// If API key is not specified in cmd line, try fetching it from
	// environ variable.
	if len(v.apiKey) == 0 {
		v.apiKey = os.Getenv(envAPIKey)
		if len(v.apiKey) == 0 {
			return nil,
				buf.String(),
				fmt.Errorf("-k API-KEY or %s environ var is required",
					envAPIKey)
		}
	}

	return &v, buf.String(), nil
}

func main() {
	v, output, err := parseArgs(os.Args[0], os.Args[1:])

	if err == flag.ErrHelp {
		fmt.Fprintln(os.Stderr, output)
		os.Exit(0)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", v)
}
