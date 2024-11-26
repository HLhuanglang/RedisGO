package config

import (
	"testing"
)

const (
	debugCfgPath = "redis_test.conf"
)

func TestSetup(t *testing.T) {
	SetupConfig(debugCfgPath)
}
