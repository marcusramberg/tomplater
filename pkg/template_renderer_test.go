package pkg

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	renderer := NewTemplateRenderer("a", "b", "c")
	require.Equal(t, renderer.templateFile, "a")
	require.Equal(t, renderer.dataFile, "b")
	require.Equal(t, renderer.configFile, "c")
}

func TestRender(t *testing.T) {
	renderer := NewTemplateRenderer("../testdata/example.txt.tmpl", "../testdata/data.toml", "../testdata/config.toml")
	require.NoError(t, renderer.Render())
	require.FileExists(t, "output.txt")
	content, err := ioutil.ReadFile("output.txt")
	require.NoError(t, err)
	require.Contains(t, string(content), "Joe Blow")
	os.Remove("output.txt")
}
