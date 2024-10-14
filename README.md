
# Hangman en Go

Ce projet est une implémentation du jeu **Hangman** en Go, avec plusieurs modes de difficulté et des easter eggs à découvrir !

## Fonctionnalités
- Plusieurs niveaux de difficulté : facile, moyen, difficile, et deux modes spéciaux : *yann* et *raciste*.
- Gestion de mots à partir de fichiers texte pour chaque niveau de difficulté.
- Interface en ligne de commande avec un affichage graphique pour le titre du jeu grâce à la bibliothèque [go-figure](https://github.com/common-nighthawk/go-figure).
- Système de vies basé sur la difficulté choisie.
- Possibilité d'annuler la partie à tout moment avec la commande `stop`.

## Installation

1. Clonez le dépôt :
   ```bash
   git clone https://github.com/votre-nom-utilisateur/hangman-go.git
   cd hangman-go
   ```

2. Installez la dépendance `go-figure` :
   ```bash
   go get github.com/common-nighthawk/go-figure
   ```

## Exécution

Pour exécuter le programme, utilisez la commande suivante :

```bash
go run . -diff <difficulte>
```

### Exemple de commandes :
- Mode facile :
  ```bash
  go run . -diff facile
  ```

- Mode moyen :
  ```bash
  go run . -diff moyen
  ```

- Mode difficile :
  ```bash
  go run . -diff difficile
  ```

- Mode "Yann" :
  ```bash
  go run . -diff yann
  ```

- Mode "Raciste" :
  ```bash
  go run . -diff raciste
  ```

- Obtenir de l'aide sur les commandes :
  ```bash
  go run . -diff h
  ```

## Fonctionnalités du Jeu

- Entrez une lettre ou un mot pour deviner le mot caché.
- Si vous devinez correctement, les lettres seront révélées. Si vous faites une erreur, vous perdrez des points de vie.
- Commandes spéciales :
  - `stop` : Arrête la partie.
  - `gay` : Affiche une phrase secrète.
  - `non` : Affiche une réponse humoristique.

## Modes de Difficulté

- **Facile** : Mots simples avec 10 vies.
- **Moyen** : Mots moyens avec 10 vies.
- **Difficile** : Mots difficiles avec 10 vies.
- **Yann** : Un mode secret.
- **Raciste** : Un autre mode spécial.

## Fichiers de mots

Les mots à deviner sont stockés dans des fichiers texte. Assurez-vous que chaque fichier est bien rempli avec des mots appropriés pour chaque niveau de difficulté.

- `motsFacile.txt` : Contient des mots faciles.
- `motsNormal.txt` : Contient des mots de difficulté moyenne.
- `motsHard.txt` : Contient des mots difficiles.
- `motsYann.txt` : Contient des mots du mode "Yann".
- `motsRaciste.txt` : Contient des mots du mode "Raciste".

## Structure du Code

Le code se compose principalement des éléments suivants :
- **Gestion des entrées utilisateur** : pour deviner des lettres ou des mots.
- **Système de vie** : Diminue à chaque erreur, basé sur la difficulté.
- **Affichage** : Utilisation de `go-figure` pour un titre stylisé du jeu.
- **Modes spéciaux** : Divers easter eggs et modes cachés, activés par des commandes spécifiques.

## Avertissement

Ce projet contient des références humoristiques qui pourraient être mal interprétées. Veuillez noter que les modes spéciaux et les commandes comme "yann" ou "raciste" sont des blagues et ne doivent pas être prises au sérieux.

---