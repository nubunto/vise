package main

import (
	"os"
	"os/exec"
	"path"

	"fmt"
)

func main() {
	var dir string
	if len(os.Args) < 2 {
		dir = "./compiled/"
	} else {
		dir = os.Args[1]
	}
	archs := []string{"amd64", "386"}
	platforms := []string{"windows", "linux", "darwin"}
	for _, p := range platforms {
		for _, arch := range archs {
			name := fmt.Sprintf("vise-%s-%s", p, arch)
			if p == "windows" {
				name = name + ".exe"
			}
			os.Setenv("GOOS", p)
			os.Setenv("GOARCH", arch)
			_, err := exec.Command("go", "build", "-o", path.Join(dir, name), "github.com/nubunto/vise").CombinedOutput()
			if err != nil {
				fmt.Printf("ERROR (%s - %s): %v\n", p, arch, err)
				return
			}
			fmt.Printf("compiled: %s\n", name)
		}
	}
}
