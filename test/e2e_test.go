package test

import (
	"context"
	"testing"

	"github.com/Masterminds/semver"
	"github.com/fhke/infrastructure-abstraction/client/sdk/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	dynamoTableNameStacks  = "stacks"
	dynamoTableNameModules = "modules"
	serverAddr             = "127.0.0.1:9001"
)

func TestE2E(t *testing.T) {
	ctx, can := context.WithCancel(context.Background())
	defer can()

	runServer(ctx, t)
	cl := client.New("http://" + serverAddr)

	// check that server responds correctly when no modules exist
	assertModuleNames(ctx, t, cl)

	// add 2 versions for 2 modules
	assertAddModuleVersion(ctx, t, cl, "test1", "test-repo.acme.com/test1", "1.1.0")
	assertAddModuleVersion(ctx, t, cl, "test1", "test-repo.acme.com/test1", "1.2.0")
	assertAddModuleVersion(ctx, t, cl, "test2", "other-test-repo.acme.com/test2", "1.5.0")
	assertAddModuleVersion(ctx, t, cl, "test2", "test-repo.acme.com/test2", "2.2.0")

	// check that modules now exist
	assertModuleNames(ctx, t, cl, "test1", "test2")

	// build a stack
	assertBuildStack(
		ctx,
		t,
		cl,
		"testing",
		"acme/test-repo",
		[]string{"test1", "test2"},
		client.Stack{
			Modules: map[string]client.BuildStackOutModule{
				"test1": {
					Version: "1.2.0",
					Source:  "test-repo.acme.com/test1",
				},
				"test2": {
					Version: "2.2.0",
					Source:  "test-repo.acme.com/test2",
				},
			},
		},
	)

	// patch stack
	assertSetStackModules(
		ctx,
		t,
		cl,
		"testing",
		"acme/test-repo",
		map[string]string{"test2": "1.5.0"},
		client.Stack{
			Modules: map[string]client.BuildStackOutModule{
				"test1": {
					Version: "1.2.0",
					Source:  "test-repo.acme.com/test1",
				},
				"test2": {
					Version: "1.5.0",
					Source:  "other-test-repo.acme.com/test2",
				},
			},
		},
	)

	// rebuild stack
	assertBuildStack(
		ctx,
		t,
		cl,
		"testing",
		"acme/test-repo",
		[]string{"test1", "test2"},
		client.Stack{
			Modules: map[string]client.BuildStackOutModule{
				"test1": {
					Version: "1.2.0",
					Source:  "test-repo.acme.com/test1",
				},
				"test2": {
					Version: "1.5.0",
					Source:  "other-test-repo.acme.com/test2",
				},
			},
		},
	)

}

func assertSetStackModules(ctx context.Context, t *testing.T, cl client.Client, stackName, repo string, moduleVersions map[string]string, assertOut client.Stack) {
	st, err := cl.SetStackModules(ctx, stackName, repo, moduleVersions)
	require.NoError(t, err, "It should set modules")
	assert.Equal(t, assertOut, st, "It should match assertion")
}

func assertBuildStack(ctx context.Context, t *testing.T, cl client.Client, stackName, repo string, moduleNames []string, assertOut client.Stack) {
	actualStack, err := cl.BuildStack(ctx, stackName, repo, moduleNames)
	require.NoError(t, err, "It should build stack")
	assert.Equal(t, assertOut, actualStack, "Output should match assertion")
}

func assertAddModuleVersion(ctx context.Context, t *testing.T, cl client.Client, modName, modSrc, modVersion string) {
	err := cl.CreateModuleVersion(ctx, modName, modSrc, semver.MustParse(modVersion))
	require.NoError(t, err, "It should add module version")
}

func assertModuleNames(ctx context.Context, t *testing.T, cl client.Client, moduleNames ...string) {
	actualNames, err := cl.GetModuleNames(ctx)
	require.NoError(t, err, "It should not return an error")
	assert.Equal(
		t,
		sets.New(moduleNames...),
		sets.New(actualNames...),
		"Module names should match",
	)

}
