package prpc

import (
	"testing"

	"github.com/oim/common/config"

	ptrace "github.com/oim/common/prpc/trace"
	"github.com/stretchr/testify/assert"
)

func TestNewPClient(t *testing.T) {
	config.Init("../../oim.yaml")
	ptrace.StartAgent()
	defer ptrace.StopAgent()

	_, err := NewPClient("oim_server")
	assert.NoError(t, err)
}
