package git_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/fhke/infrastructure-abstraction/client/util/git"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCurrentRepoSlug__HTTPS(t *testing.T) {
	tmp := t.TempDir()
	require.NoError(t, os.Chdir(tmp))

	mustExec(t, "git", "init", ".")
	mustExec(t, "git", "remote", "add", "origin", "https://github.com/sampleorg/myrepo.git")

	slug, err := git.GetCurrentRepoSlug()
	require.NoError(t, err, "It should not error")
	assert.Equal(t, "sampleorg/myrepo", slug, "It should return repo slug")
}

func TestGetCurrentRepoSlug__SSH(t *testing.T) {
	tmp := t.TempDir()
	require.NoError(t, os.Chdir(tmp))

	mustExec(t, "git", "init", ".")
	mustExec(t, "git", "remote", "add", "origin", "git@github.com:othersampleorg/other.repo.git")

	slug, err := git.GetCurrentRepoSlug()
	require.NoError(t, err, "It should not error")
	assert.Equal(t, "othersampleorg/other.repo", slug, "It should return repo slug")
}

func TestGetCurrentRepoSlug__NestedDir(t *testing.T) {
	tmp := t.TempDir()
	nestedTmp := filepath.Join(tmp, "nested", "dir")
	require.NoError(t, os.MkdirAll(nestedTmp, 0750))
	require.NoError(t, os.Chdir(tmp))
	mustExec(t, "git", "init", ".")
	mustExec(t, "git", "remote", "add", "origin", "https://gitlab.org/sample.org/myrepo.git")
	require.NoError(t, os.Chdir(nestedTmp))

	slug, err := git.GetCurrentRepoSlug()
	require.NoError(t, err, "It should not error")
	assert.Equal(t, "sample.org/myrepo", slug, "It should return repo slug")
}

func mustExec(t *testing.T, cmd string, args ...string) {
	require.NoError(
		t,
		exec.Command(cmd, args...).Run(),
	)
}
