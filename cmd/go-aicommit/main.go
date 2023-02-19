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

	cmd = exec.Command("git", "diff", "--cached", "--", ".", "':(exclude)yarn.lock'", "':(exclude)package-lock.json'")

	diff, err := cmd.Output()
	if err != nil {
		fmt.Println("Erreur lors de la récupération de git diff :", err)
		os.Exit(1)
	}

	if len(diff) == 0 {
		fmt.Println("Diff est vide, impossible de générer un message de commit.")
		os.Exit(1)
	}

	commit := aicommit.New(1, diff)

	fmt.Println("Génération du message de commit 🤖")
	message, err := commit.GenerateCommitMessage()

	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(message)
}
