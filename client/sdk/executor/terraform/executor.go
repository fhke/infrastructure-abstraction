package terraform

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fhke/infrastructure-abstraction/sdk/executor"
	"github.com/fhke/infrastructure-abstraction/sdk/render/terraform"
	"github.com/fhke/infrastructure-abstraction/util/git"
)

const (
	TerraformCommand = "terraform"
	TerraformFile    = "generated.tf.json"
)

type terraformExecutor struct {
	terraformDir string
}

func New(terrformDir string) (executor.Executor[terraform.TerraformStack], error) {
	if err := os.MkdirAll(terrformDir, 0750); err != nil {
		return nil, fmt.Errorf("error creating terraform dir: %w", err)
	}
	if err := git.IgnoreDir(terrformDir); err != nil {
		return nil, fmt.Errorf("error configuring gitignore: %w", err)
	}

	return &terraformExecutor{
		terraformDir: terrformDir,
	}, nil
}

func (t *terraformExecutor) Exec(stack terraform.TerraformStack) error {
	terraformFile := filepath.Join(t.terraformDir, TerraformFile)

	tfJSON, err := json.Marshal(stack)
	if err != nil {
		return fmt.Errorf("error encoding module to JSON: %w", err)
	}

	if err := os.WriteFile(terraformFile, tfJSON, 0640); err != nil {
		return fmt.Errorf("error writing Terraform file: %w", err)
	}

	if err := t.execCmd(TerraformCommand, "init", "-upgrade"); err != nil {
		return fmt.Errorf("init error: %w", err)
	}
	if err := t.execCmd(TerraformCommand, "apply"); err != nil {
		return fmt.Errorf("apply error: %w", err)
	}
	return nil
}

func (t *terraformExecutor) execCmd(cmd string, args ...string) error {
	log.Printf("Executing command %q in dir %q with args: %+v", cmd, t.terraformDir, args)
	eCmd := exec.Command(cmd, args...)
	eCmd.Dir = t.terraformDir
	eCmd.Stdout = os.Stdout
	eCmd.Stderr = os.Stderr
	eCmd.Stdin = os.Stdin
	return eCmd.Run()
}
