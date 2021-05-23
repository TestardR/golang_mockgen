package uuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	generator := New()
	uuid := generator.Generate()
	err := generator.Parse(uuid)

	require.NoError(t, err)
}

func TestParse(t *testing.T) {
	generator := New()
	id := uuid.New().String()

	err := generator.Parse(id)
	require.NoError(t, err)

	err = generator.Parse("test")
	require.Error(t, err)
}
