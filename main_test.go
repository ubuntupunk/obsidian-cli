package main

import (
    "os"
    "os/exec"
    "testing"
)

func TestBuildCommand(t *testing.T) {
    platforms := []struct {
        GOOS      string
        GOARCH    string
        BinaryName string
        SourceFile string
    }{
        {"darwin", "amd64", "test_binary_darwin", "main.go"},
        {"linux", "amd64", "test_binary_linux", "main.go"},
        {"windows", "amd64", "test_binary_windows.exe", "main.go"},
        {"linux", "arm64", "test_binary_rpi", "main.go"},
    }

    for _, p := range platforms {
        // Set the environment variables for GOOS and GOARCH
        cmd := exec.Command("go", "build", "-o", p.BinaryName, p.SourceFile)
        cmd.Env = append(os.Environ(), "GOOS="+p.GOOS, "GOARCH="+p.GOARCH)

        err := cmd.Run()
        if err != nil {
            t.Fatalf("Failed to build for %s/%s: %v", p.GOOS, p.GOARCH, err)
        }

        // Check if the binary was created
        if _, err := os.Stat(p.BinaryName); os.IsNotExist(err) {
            t.Errorf("Expected binary '%s' not found", p.BinaryName)
        }

        // Clean up the created binary
        os.Remove(p.BinaryName)
    }
}