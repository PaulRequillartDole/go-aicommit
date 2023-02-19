# go-aicommit

Le projet go-aicommit est un utilitaire de ligne de commande écrit en Go qui permet de générer automatiquement un message de commit en fonction des changements apportés à votre projet.

## Installation

Pour installer go-aicommit, vous devez avoir Go installé sur votre machine. Ensuite, vous pouvez exécuter la commande suivante :

```shell
go install github.com/PaulRequillartDole/go-aicommit@latest
```

## Utilisation

Une fois installé, vous pouvez utiliser go-aicommit en exécutant la commande suivante à la racine de votre projet :

```shell
git add .
go-aicommit
```

Cette commande va analyser les fichiers modifiés dans votre projet et générer un message de commit basé sur ces modifications. Le message de commit sera alors affiché dans la console.

## Contribution

Les contributions sont les bienvenues ! Si vous souhaitez contribuer à ce projet, veuillez créer une pull request sur GitHub.

## Licence

Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus d'informations.


