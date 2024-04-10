package dsl

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fhke/infrastructure-abstraction/sdk/client"
	"github.com/fhke/infrastructure-abstraction/util/git"
	"github.com/fhke/tpp"
	"github.com/samber/lo"
	"k8s.io/utils/ptr"
)

func Run(ctx context.Context, cl client.Client, sourceDir, stackName string) error {
	outDir := "./.generatedTerraform"
	if err := os.MkdirAll(outDir, 0750); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}
	if err := git.IgnoreDir(outDir); err != nil {
		return fmt.Errorf("error gitignoring output directory: %w", err)
	}

	repoName, err := git.GetCurrentRepoSlug()
	if err != nil {
		return fmt.Errorf("error getting current repo name: %w", err)
	}

	tfs, err := tpp.NewTerraformsForDir(sourceDir)
	if err != nil {
		return fmt.Errorf("error reading Terraform: %w", err)
	}

	origModNames, err := tfs.GetModuleSources()
	if err != nil {
		return fmt.Errorf("error getting module sources: %w", err)
	}

	bso, err := cl.BuildStack(ctx, stackName, repoName, origModNames)
	if err != nil {
		return fmt.Errorf("error from API for build stack: %w", err)
	}

	err = tfs.SetModuleSources(lo.MapValues(bso.Modules, func(mod client.BuildStackOutModule, _ string) tpp.ModuleSource {
		var ver *string
		if mod.Version != "" {
			ver = ptr.To(mod.Version)
		}
		return tpp.ModuleSource{
			Source:  mod.Source,
			Version: ver,
		}
	}))
	if err != nil {
		return fmt.Errorf("error replacing module sources: %w", err)
	}

	if err := tfs.WriteTo(outDir); err != nil {
		return fmt.Errorf("error writing Terraform to output directory: %w", err)
	}

	if err := runTerraform(outDir, "init", "-upgrade"); err != nil {
		return fmt.Errorf("error running Terraform init: %w", err)
	}
	if err := runTerraform(outDir, "apply"); err != nil {
		return fmt.Errorf("error running Terraform apply: %w", err)
	}
	return nil
}

func runTerraform(dir string, args ...string) error {
	log.Printf("Running Terraform in directory %s with args %+v", dir, args)
	cmd := exec.Command("terraform", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Dir = dir
	return cmd.Run()
}
