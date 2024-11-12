package cmd_test

import (
	"regexp"
	"testing"
	"uuid-generator/cmd"
)

// TestVersion check whether version command have returned the correct version or not.
func TestVersion(t *testing.T) {
	ver := "v4"
	want := regexp.MustCompile(`\b`+ver+`\b`)
	v := cmd.GetVersion()
	if !want.MatchString(v) {
		t.Fatalf(`GetVersion() = %q, want match for %#q, nil`, v, want)
	}
}