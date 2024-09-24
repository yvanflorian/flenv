package utils

import (
	"fmt"
	"log"
	"os/exec"
)

// Register Encrypter
func EncryptConfigFile() error {
	log.Println("Encryting Config File...")
	mFilePath, err := GetConfigPath()
	if err != nil {
		return err
	}

	cmd := exec.Command("ansible-vault", "encrypt", mFilePath)
	output, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("encrypt output: %v", err)
	}
	fmt.Println("command output", string(output))
	return nil
}
