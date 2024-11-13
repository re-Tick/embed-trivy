package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

//go:embed assets/trivy
var embeddedBinary []byte

func main() {
	// Step 1: Create a temporary file to write the embedded binary
	tmpFile, err := os.Create("trivy-test-isolation")
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}
	// Make sure to close the file after writing
	// defer os.Remove(tmpFile.Name()) // Clean up the temp file when done

	// Step 2: Write the embedded binary to the temporary file
	_, err = io.Copy(tmpFile, bytes.NewReader(embeddedBinary))
	if err != nil {
		log.Fatalf("Failed to write to temporary file: %v", err)
	}

	// Step 3: Close the file before executing
	if err := tmpFile.Close(); err != nil {
		log.Fatalf("Failed to close temporary file: %v", err)
	}

	// Step 4: Make the file executable
	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		log.Fatalf("Failed to make file executable: %v", err)
	}

	// Step 5: Run the Trivy command using the temporary file
	// Example: Running "trivy --version"
	currentPAth, _ := os.Getwd()
	cmd := exec.Command(currentPAth+"/"+tmpFile.Name(), "--version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	fmt.Println("Trivy command executed successfully! the path to the tmp trivy file is: ", tmpFile.Name())
}
