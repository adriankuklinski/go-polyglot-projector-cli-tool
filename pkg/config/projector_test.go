package config_test

import (
	"testing"

	"github.com/adriankuklinski/go-polyglot-projector-cli-tool/pkg/config"
)

func getData() *config.Data {
	return &config.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"baz": "zap",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *config.Data) *config.Projector {
	return config.CreateProjector(
		&config.Config{
			Args:      []string{},
			Operation: config.Print,
			Pwd:       pwd,
			Config:    "Hello, Squirrel",
		},
		data,
	)
}

func test(t *testing.T, proj *config.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%+v\"", value)
	}

	if value != v {
		t.Errorf("expected to find %v but recieved %+v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
}
