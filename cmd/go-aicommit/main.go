package main

import (
	"fmt"
	"github.com/PaulRequillartDole/go-aicommit/internal/aicommit"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	output, err := cmd.Output()

	if !strings.Contains(string(output), "true") || err != nil {
		fmt.Println("Le répertoire actuel n'est pas un dépôt Git valide.")
		os.Exit(1)
	}

	cmd = exec.Command("git", "diff", "--cached", "--", ".", "':(exclude)yarn.lock'", "':(exclude)package-lock.json'")
	output, err = cmd.Output()

	if len(output) == 0 || err != nil {
		fmt.Println("Diff est vide, impossible de générer un message de commit.")
		os.Exit(1)
	}

	commit := aicommit.New(1, output)

	fmt.Println("Génération du message de commit 🤖")
	message, err := commit.GenerateCommitMessage()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(message)
}
