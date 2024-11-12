package generator_test

import (
	"regexp"
	"testing"
	"uuid-generator/generator"
)

func TestUUID(t *testing.T) {
	pattern := "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
	want := regexp.MustCompile(pattern)
	g := generator.GeneratorV4{}
	uuid, err := g.UUID()
	if err != nil {
		t.Fatalf("UUIDv4 generator returned error: %v", err)
	}
	if !want.MatchString(g.String(uuid)) {
		t.Fatalf("UUID() string = %q, want match for #%q", uuid, want)
	}
}