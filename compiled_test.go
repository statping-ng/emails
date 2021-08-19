package emails

import (
	"fmt"
	"github.com/statping-ng/statping-ng/types/core"
	"github.com/statping-ng/statping-ng/types/failures"
	"github.com/statping-ng/statping-ng/types/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type replacer struct {
	Core    core.Core
	Service services.Service
	Failure failures.Failure
	Custom  map[string]string
	Email   string
}

func TestServiceOnline(t *testing.T) {
	example := services.Example(true)

	replaced := replacer{
		Core:    *core.Example(),
		Service: example,
		Failure: failures.Example(),
		Email:   "info@statping.com",
	}
	tmpl, err := Parse(Success, replaced)
	require.Nil(t, err)
	assert.Contains(t, tmpl, example.Name)
	assert.Contains(t, tmpl, fmt.Sprintf("%0.0f seconds", example.Downtime().Seconds()))
}

func TestServiceOffline(t *testing.T) {
	example := services.Example(false)
	failure := failures.Example()

	replaced := &replacer{
		Core:    *core.Example(),
		Service: example,
		Failure: failure,
		Email:   "info@statping.com",
	}
	tmpl, err := Parse(Failure, replaced)
	require.Nil(t, err)
	assert.Contains(t, tmpl, example.Name)
	assert.Contains(t, tmpl, fmt.Sprintf("%0.0f seconds", example.Downtime().Seconds()))
}
