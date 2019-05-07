package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
)

var description = "Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\n"
var appname = os.Args[0:1][0]

var objectTests = []struct {
	in  []string
	out string
}{
	{[]string{appname, "objects", "https://app.argu.co/argu/u/joep", "http://schema.org/description"}, description},
}

func TestObjectParser(t *testing.T) {
	for _, tt := range objectTests {
		fmt.Print(tt.in[1:])
		out := capturer.CaptureStdout(func() {
			run(tt.in)
		})
		if tt.out != out {
			t.Error(fmt.Sprintf("Expected: \n%vGot:\n%v", tt.out, out))
		} else {
			fmt.Print("PASS\n")
		}
	}
}