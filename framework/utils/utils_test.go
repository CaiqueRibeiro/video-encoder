package utils_test

import (
	"encoder/framework/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
						"id": "11111111",
						"file_path": "teste.mp4",
						"status": "pending"
				}`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = "teste"
	err = utils.IsJson(json)
	require.Error(t, err)
}
