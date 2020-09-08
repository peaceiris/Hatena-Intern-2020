package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fetch_Title(t *testing.T) {
	url := "https://example.com/"
	extected := "Example Domain"
	title, err := Fetch(context.Background(), url)
	assert.NoError(t, err)
	assert.Equal(t, extected, title)
}
