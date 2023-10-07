package openapi31_test

import (
	"os"
	"testing"

	"github.com/STARRISEPH/openapi-go/openapi31"
	"github.com/stretchr/testify/require"
)

func TestSpec_UnmarshalYAML(t *testing.T) {
	bytes, err := os.ReadFile("testdata/albums_api.yaml")
	require.NoError(t, err)

	refl := openapi31.NewReflector()
	require.NoError(t, refl.Spec.UnmarshalYAML(bytes))
}
