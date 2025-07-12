package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/Lesnoi-Dmytro/drank-dwarves-api/internal/template/feat"
	"github.com/Lesnoi-Dmytro/drank-dwarves-api/pkg/exec"
)

func main() {
	root := "internal"

	name, err := readName(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Creating feature", name+"...")

	dir := filepath.Join(root, "core", name)
	err = os.Mkdir(dir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", dir, err)
		return
	}

	ch := make(chan error, 4)
	go writeFile(dir, "dto.go", name, feat.DtoFileTemplate, ch)
	go writeFile(dir, "service.go", name, feat.ServiceFileTemplate, ch)
	go writeFile(dir, "handlers.go", name, feat.HandlersFileTemplate, ch)
	go writeFile(dir, "router.go", name, feat.RouterFileTemplate, ch)

	hasError := false
	for i := 0; i < 4; i++ {
		err := <-ch
		if err != nil {
			hasError = true
		}
	}
	if hasError {
		deleteDir(dir)
		return
	}

	err = updateRouter(root, name)
	if err != nil {
		fmt.Println("Error updating router config: ", err)
		deleteDir(dir)
		return
	}

	exec.ExecuteCommand("goimports", "-w", filepath.Join("internal", "config", "router"))

	fmt.Printf("Feature %s generation completed", name)
}

func readName(root string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Name of a feature: ")
	scanner.Scan()
	name := scanner.Text()

	if len(name) < 2 {
		return "", errors.New("feature name should have at least 2 letters")
	}
	for _, r := range name {
		if !unicode.IsLower(r) && r != '-' {
			return "", errors.New("feature name should only contain lower case latin letters and '-' symbol to separate words")
		}
	}

	stats, err := os.Stat(filepath.Join(root, name))
	if err != nil {
		if !os.IsNotExist(err) {
			return "", fmt.Errorf("error reading %s directory: %v", root, err)
		}
	} else if stats.IsDir() {
		return "", fmt.Errorf("feature directory %s already exists", name)
	}

	return name, nil
}

func writeFile(dir string, fileName string, name string, template string, ch chan error) {
	filePath := filepath.Join(dir, fileName)

	err := os.WriteFile(filePath, feat.FileContent(name, template), os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
	} else {
		fmt.Printf("File %s created\n", filePath)
	}

	ch <- err
}

func deleteDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Printf("Error deleting directory %s: %v\n", dir, err)
	} else {
		fmt.Printf("Directory %s deleted\n", dir)
	}
}

func updateRouter(root string, name string) error {
	filePath := filepath.Join(root, "config", "router", "router.go")

	content, err := os.ReadFile(filepath.Join(root, "config", "router", "router.go"))
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}
	strContent := string(content)

	newRouterStr := fmt.Sprintf("r.Mount(\"/%[1]s\", %[1]s.Router())\n\t\t", name)
	newImport := fmt.Sprintf("\"github.com/Lesnoi-Dmytro/drank-dwarves-api/internal/core/%s\"\n\t", name)

	importsStr := "\"github.com/Lesnoi-Dmytro/"
	importsStartIdx := strings.Index(strContent, importsStr)
	routesEndIdx := strings.Index(strContent, "r.Get(\"/swagger\",")

	updatedContent := strContent[:importsStartIdx] + newImport +
		strContent[importsStartIdx:routesEndIdx] + newRouterStr + strContent[routesEndIdx:]

	err = os.WriteFile(filePath, []byte(updatedContent), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing file %s: %v", filePath, err)
	}

	return nil
}
