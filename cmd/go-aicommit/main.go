package main

import (
	"fmt"
	"github.com/PaulRequillartDole/go-aicommit/internal/aicommit"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Erreur lors de la vérification du dépôt Git :", err)
		os.Exit(1)
	}
	if !strings.Contains(string(output), "true") {
		fmt.Println("Le répertoire actuel n'est pas un dépôt Git valide.")
		os.Exit(1)
	}

	cmd = exec.Command("git", "diff")
	diff, err := cmd.Output()
	if err != nil {
		fmt.Println("Erreur lors de la récupération de git diff :", err)
		os.Exit(1)
	}

	if diff == nil {
		fmt.Println("Diff est vide", err)
		os.Exit(1)
	}

	commit := aicommit.New(1, diff)
	message, err := commit.GenerateCommitMessage()

	if err != nil {
		log.Panic(err.Error())
	}

	println(message)
}
