package commands

import (
	"io/ioutil"
	"reflect"
	"regexp"
	"testing"

	"github.com/hobeone/gonab/types"
	. "github.com/onsi/gomega"
)

func TestNewzNabFileParse(t *testing.T) {
	RegisterTestingT(t)

	content, err := ioutil.ReadFile("testdata/latestregex.sql")
	if err != nil {
		t.Fatalf("Error parsing: %v", err)
	}
	regexes, err := parseNewzNabRegexes(content)
	if err != nil {
		t.Fatalf("Error parsing: %v", err)
	}
	if len(regexes) != 5 {
		t.Fatalf("Expected 5 regexes, got %d", len(regexes))
	}
}

func TestNewsNabToRegex(t *testing.T) {
	RegisterTestingT(t)

	p := []string{
		"",
		"1",
		"misc.test",
		"/^(?P<name>.*?)\\s==\\s\\((?P<parts>\\d{1,3}\\/\\d{1,3})/iS ",
		"150",
		"1",
		"",
		"NULL",
	}
	reg, err := newzNabRegexToRegex(p)
	if err != nil {
		t.Fatalf("Error parsing: %v", err)
	}
	expected := &types.Regex{
		ID:          1,
		Regex:       `(?i)^(?P<name>.*?)\s==\s\((?P<parts>\d{1,3}\/\d{1,3})`,
		Description: "",
		Status:      true,
		Ordinal:     150,
		GroupName:   "misc.test",
	}
	Expect(reg).Should(Equal(expected))
	if !reflect.DeepEqual(reg, expected) {
		t.Fatalf("Unexpected parse result %#v != %#v", reg, expected)
	}
	_, err = regexp.Compile(reg.Regex)
	if err != nil {
		t.Fatalf("Returned non compilable regex: %v", err)
	}
}
