package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"

	"github.com/zalando/go-keyring"
)

func generateRandomKey() (string, error) {
	// 128bits
	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func InitKeyRing() error {
	k, err := generateRandomKey()
	if err != nil {
		return fmt.Errorf("failed generating random key:", err)
	}

	err = keyring.Set(FLENV_KEYRING_SERVICENAME, FLENV_KEYRING_SERVICEUSER, k)
	if err != nil {
		return fmt.Errorf("Failed to setting secret in the OS KeyRing. Error: %v", err)
	}
	return nil
}

func GetSecret() (string, error) {
	pwd, err := keyring.Get(FLENV_KEYRING_SERVICENAME, FLENV_KEYRING_SERVICEUSER)
	if err != nil {
		return "", fmt.Errorf("Failed Retrieving secret from OS KeyRing. Error: %v", err)
	}
	return pwd, nil
}

func EncryptFlenvConfigFile(filename, passphrase string) error {
	tempFile := filename + ".gpg"
	cmd := exec.Command(
		"gpg",
		"--batch",
		"--yes",
		"--passphrase",
		passphrase,
		"-c",
		"--output",
		tempFile,
		filename,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("encryption failed: %v, output: %s", err, output)
	}

	// Remove the original file
	err = os.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to remove original file: %w", err)
	}

	// Rename the encrypted file to the original filename
	err = os.Rename(tempFile, filename)
	if err != nil {
		return fmt.Errorf("failed to rename encrypted file: %w", err)
	}

	return nil
}

func DecryptFlenvConfigFile(filename, passphrase string) error {
	// Create a temporary file for the decrypted content
	tempFile := filename + ".tmp"

	// Decrypt the file to the temporary location
	cmd := exec.Command(
		"gpg",
		"--batch",
		"--yes",
		"--passphrase",
		passphrase,
		"-d",
		"--output",
		tempFile,
		filename,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("decryption failed: %v, output: %s", err, output)
	}

	// Remove the original encrypted file
	err = os.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to remove encrypted file: %w", err)
	}

	// Rename the decrypted file to the original filename
	err = os.Rename(tempFile, filename)
	if err != nil {
		return fmt.Errorf("failed to rename decrypted file: %w", err)
	}

	return nil
}
