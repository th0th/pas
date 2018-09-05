package main

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/putdotio/pas/internal/pas"
	"github.com/stretchr/testify/assert"
)

func TestConfigUnmarshalEvents(t *testing.T) {
	s := `
	[events]
	[events.test_event]
	test_property = "string"
	[events.test_event2]
	test_property2 = "string"
	`
	var c Config
	err := toml.Unmarshal([]byte(s), &c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(c.Events), 2)
	assert.Equal(t, c.Events[pas.EventName("test_event")][pas.PropertyName("test_property")], pas.PropertyType("string"))
}
