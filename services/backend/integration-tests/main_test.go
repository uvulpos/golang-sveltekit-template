package main

import (
	"os"
	"testing"

	"github.com/uvulpos/golang-sveltekit-template/integration-tests/helper/setup"
)

func TestMain(m *testing.M) {
	setup.MigrateDatabaseWithTestData()

	exitVal := m.Run()
	os.Exit(exitVal)
}
