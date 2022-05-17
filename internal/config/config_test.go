package config

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func TestInit(t *testing.T) {
	err := Init(&pflag.FlagSet{}) //nolint

	if err != nil {
		t.Errorf("unexpected error initializing config: %s", err.Error())
	}
}

func TestReadConfigFile(t *testing.T) {
	Init(&pflag.FlagSet{}) //nolint

	viper.Set("config-path", "../../test/test-config.yml")

	err := ReadConfigFile() //nolint
	if err != nil {
		t.Errorf("unexpected error reading config: %s", err.Error())
	}
}

func TestReadConfigFile_Error(t *testing.T) {
	Init(&pflag.FlagSet{}) //nolint

	viper.Set("config-path", "notafile.json")

	err := ReadConfigFile() //nolint
	if err == nil {
		t.Errorf("expected error reading config: %s", err.Error())
	} else if err.Error() != "open notafile.json: no such file or directory" {
		t.Errorf("unexpected error reading config, got: '%s', want: 'open notafile.json: no such file or directory'", err.Error())
	}
}
