package main

import (
	"flag"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

const (
	dummyZipCode = 12345
	dummyAPIKey  = "1234qwerasdfzxcv"
)

func TestArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		expected *Veyil
	}{
		{"all args",
			"-z 27606 -k 12345 -f /tmp/veyil_out -i 33",
			&Veyil{
				27606,
				33,
				"12345",
				"/tmp/veyil_out",
			},
		},
		{"all env", "", &Veyil{dummyZipCode, defaultInterval, dummyAPIKey, outFile}},
		{"zip", "-z 95134", &Veyil{95134, defaultInterval, dummyAPIKey, outFile}},
		{"api", "-k abcdefg", &Veyil{dummyZipCode, defaultInterval, "abcdefg", outFile}},
		{"interval", "-i 456", &Veyil{dummyZipCode, 456, dummyAPIKey, outFile}},
	}

	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if tt.expected.zipCode == dummyZipCode {
				os.Setenv(envZipCode, strconv.Itoa(dummyZipCode))
			}

			if tt.expected.apiKey == dummyAPIKey {
				os.Setenv(envAPIKey, dummyAPIKey)
			}

			actual, output, err := parseArgs(
				"prog", strings.Split(tt.args, " "))
			if err != nil {
				t.Fatalf("test %s: parsing args failed: %s\n",
					tt.name, err)
			}

			if output != "" {
				t.Fatalf("test %s: got %q, want empty\n", tt.name, output)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("test %s: actual{%+v} expected{%+v}\n",
					tt.name, actual, tt.expected)
			}
		})
	}
}

func TestArgsHelp(t *testing.T) {
	helpArgs := []string{"-h", "--help", "-help"}

	for _, arg := range helpArgs {
		t.Run(arg, func(t *testing.T) {
			v, _, err := parseArgs("prog", []string{arg})

			// We should only receive 'flag.ErrHelp' as err when help is
			// requested.
			if err != flag.ErrHelp {
				t.Fatalf("TestArgsHelp: got err %v, want flag.ErrHelp", err)
			}

			// We should never get a pointer to Veyil struct for help.
			if v != nil {
				t.Fatalf("TestArgsHelp: got veyil %v, want nil", v)
			}
		})
	}
}
