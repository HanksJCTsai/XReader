package config

import "testing"

func TestLoadConfig(t *testing.T) {
	if err := LoadConfig(""); err != nil {
		t.Errorf("Failed to load default configuration: %s", err)
	}
	if Cfg.LogLevel != "info" {
		t.Errorf("Expected log_level to be 'info', but got: %s", Cfg.LogLevel)
	}
}
