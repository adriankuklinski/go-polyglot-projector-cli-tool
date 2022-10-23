package config_test

import (
	"testing"
)

func getData() *projector.Data {
	return &projector.Data{
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

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Args:      []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "Hello, Squirrel",
		}, data)
}

func test(t *testing.T, proj *projector.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Error("expected to find value \"%v\"")
	}

	if value != v {
		t.Error("expected to find %v but recieved %v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("foo/bar", data)
	value, ok := proj.GetValue("foo")

	test(t, proj, "foo", "bar3")
}
