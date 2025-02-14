package mgmt

import (
	"context"
	"net/http"
	"testing"

	"github.com/descope/go-sdk/descope/tests/helpers"
	"github.com/stretchr/testify/require"
)

func TestRoleCreateSuccess(t *testing.T) {
	mgmt := newTestMgmt(nil, helpers.DoOk(func(r *http.Request) {
		require.Equal(t, r.Header.Get("Authorization"), "Bearer a:key")
		req := map[string]any{}
		require.NoError(t, helpers.ReadBody(r, &req))
		require.Equal(t, "abc", req["name"])
		require.Equal(t, "description", req["description"])
		roleNames := req["permissionNames"].([]any)
		require.Len(t, roleNames, 1)
		require.Equal(t, "foo", roleNames[0])
	}))
	err := mgmt.Role().Create(context.Background(), "abc", "description", []string{"foo"})
	require.NoError(t, err)
}

func TestRoleCreateError(t *testing.T) {
	mgmt := newTestMgmt(nil, helpers.DoOk(nil))
	err := mgmt.Role().Create(context.Background(), "", "description", []string{"foo"})
	require.Error(t, err)
}

func TestRoleUpdateSuccess(t *testing.T) {
	mgmt := newTestMgmt(nil, helpers.DoOk(func(r *http.Request) {
		require.Equal(t, r.Header.Get("Authorization"), "Bearer a:key")
		req := map[string]any{}
		require.NoError(t, helpers.ReadBody(r, &req))
		require.Equal(t, "abc", req["name"])
		require.Equal(t, "def", req["newName"])
		require.Equal(t, "description", req["description"])
		roleNames := req["permissionNames"].([]any)
		require.Len(t, roleNames, 1)
		require.Equal(t, "foo", roleNames[0])
	}))
	err := mgmt.Role().Update(context.Background(), "abc", "def", "description", []string{"foo"})
	require.NoError(t, err)
}

func TestRoleUpdateError(t *testing.T) {
	mgmt := newTestMgmt(nil, helpers.DoOk(nil))
	err := mgmt.Role().Update(context.Background(), "", "def", "description", []string{"foo"})
	require.Error(t, err)
	err = mgmt.Role().Update(context.Background(), "abc", "", "description", []string{"foo"})
	require.Error(t, err)
}

func TestRoleDeleteSuccess(t *testing.T) {
	mgmt := newTestMgmt(nil, helpers.DoOk(func(r *http.Request) {
		require.Equal(t, r.Header.Get("Authorization"), "Bearer a:key")
		req := map[string]any{}
		require.NoError(t, helpers.ReadBody(r, &req))
		require.Equal(t, "abc", req["name"])
	}))
	err := mgmt.Role().Delete(context.Background(), "abc")
	require.NoError(t, err)
}

func TestRoleDeleteError(t *testing.T) {
	mgmt := newTestMgmt(nil, helpers.DoOk(nil))
	err := mgmt.Role().Delete(context.Background(), "")
	require.Error(t, err)
}

func TestRoleLoadSuccess(t *testing.T) {
	response := map[string]any{
		"roles": []map[string]any{{
			"name": "abc",
		}}}
	mgmt := newTestMgmt(nil, helpers.DoOkWithBody(func(r *http.Request) {
		require.Equal(t, r.Header.Get("Authorization"), "Bearer a:key")
	}, response))
	res, err := mgmt.Role().LoadAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Len(t, res, 1)
	require.Equal(t, "abc", res[0].Name)
}

func TestRoleLoadError(t *testing.T) {
	mgmt := newTestMgmt(nil, helpers.DoBadRequest(nil))
	res, err := mgmt.Role().LoadAll(context.Background())
	require.Error(t, err)
	require.Nil(t, res)
}
