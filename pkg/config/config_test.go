package config_test

import (
	"reflect"
	"testing"

	"github.com/adriankuklinski/go-polyglot-projector-cli-tool/pkg/config"
)

func getOpts(args []string) *config.Opts {
	opts := &config.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}

	return opts
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation config.Operation) {
	opts := getOpts(args)
	config, err := config.NewConfig(opts)

	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("expected args to be %+v but got %+v", expectedArgs, config.Args)
	}

	if config.Operation != operation {
		t.Errorf("operation expect was %v but got %v", operation, config.Operation)
	}
}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, config.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, config.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, config.Add)
}

func TestConfigRemoveKeyValue(t *testing.T) {
	testConfig(t, []string{"rm", "foo"}, []string{"foo"}, config.Remove)
}
