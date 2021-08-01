package tests

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDockerfiles(t *testing.T) {
	workspaces, err := discoverWorkspaces()
	require.NoError(t, err)
	require.Len(t, workspaces, 1) // TODO: Antes de fechar esse pull request precisamos testar todos os projetos.

	workspaces = workspaces[:1]
	require.Len(t, workspaces, 1)

	for _, workspace := range workspaces {
		t.Run(workspace.directory, func(t *testing.T) {
			shutdownfn := setupServer(t, workspace)
			defer func() {
				require.NoError(t, shutdownfn())
			}()

			// TODO: Encontrar uma forma mais elegante de descobrir que o servidor
			// está funcionando.
			<-time.After(2 * time.Second)

			resp, err := http.Get("http://0.0.0.0:" + serverPort)
			require.NoError(t, err)

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)

			require.Equal(t, `{"msg":"error in path"}`, string(body))
			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
	}
}
