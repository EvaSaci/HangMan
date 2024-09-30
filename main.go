package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Déclaration de difficulte
type diff int

// Déclaration des constantes pour la difficulté
const (
	facile    diff = iota
	moyen     diff = iota
	difficile diff = iota
	yann      diff = iota
	raciste   diff = iota
)

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Gestion du flag pour la difficulté
	difficultyFlag := flag.String("diff", "moyen", "Choisir la difficulté (facile, moyen, difficile)")
	flag.Parse()

	var fichier *os.File
	var err error

	// Gestion de la difficulté
	switch *difficultyFlag {
	case "facile": // mode facile
		fichier, err = os.Open("mots/motsFacile.txt") // ouvre le fichier des mots facile
		fmt.Println("mode facile")
	case "moyen": // mode moyen
		fichier, err = os.Open("mots/motsNormal.txt") // ouvre le fichier des mots Moyen
		fmt.Println("mode Moyen")
	case "difficile": // mode difficile
		fichier, err = os.Open("mots/motsHard.txt") // ouvre le fichier des mots Difficile
		fmt.Println("mode Difficile")
	case "yann": // mode Yann
		fmt.Println("Mode yann")
		fichier, err = os.Open("mots/motsYann.txt") // ouvre le fichier des mots de Yann
	case "raciste": // mode raciste
		fmt.Println("Mode raciste")
		fichier, err = os.Open("mots/motsRaciste.txt") // ouvre le fichier des mots Raciste
	default:
		fmt.Println("Difficulté inconnue. Utilisation de la difficulté moyenne par défaut.") // difficulté par default = moyen
		fichier, err = os.Open("mots/motsNormal.txt")                                        // ouvre le fichier des mots moyen
	}

	// Vérification d'erreurs lors de l'ouverture du fichier
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer fichier.Close()

	// Lire les mots dans un tableau
	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	// Vérifier s'il y a des mots dans la liste
	if len(mots) == 0 {
		fmt.Println("Le fichier ne contient aucun mot")
		return
	}

	// Choisir un mot aléatoirement
	mot := mots[rand.Intn(len(mots))]

	// Définir les vies en fonction de la difficulté
	var pv int
	switch *difficultyFlag {
	case "facile":
		pv = 10
	case "moyen":
		pv = 10
	case "difficile":
		pv = 10
	case "yann":
		pv = 10
	case "raciste":
		pv = 10
	default:
		pv = 10 // par défaut moyen
	}

	var test string               // Stocke la lettre ou mot entrée par le joueur
	estla := make(map[rune]bool)  // Lettres correctes
	estpas := make(map[rune]bool) // Lettres incorrectes

	fmt.Println("Bienvenue dans le jeu Hangman !")
	fmt.Println("Prêt ?")
	fmt.Println("Tu peux annuler en écrivant 'non'")
	fmt.Println("(Si tu veux pas y jouer t'es gay)")
	fmt.Println("")

	for {
		fmt.Print("Entrez une lettre ou un mot: ")
		fmt.Scan(&test)
		if test == "gay" {
			fmt.Println("la beuteu a Yann")
			break
		}
		if test == "non" {
			fmt.Println("t'es gay")
			break
		}

		// Si le joueur entre un mot entier
		if len(test) != 1 && test != mot {
			pv--
			fmt.Printf("Mot incorrect ! Il vous reste %d chances\n", pv)
		} else if len(test) != 1 && test == mot {
			fmt.Printf("Bravo mec ! Tu as trouvé le mot : %s\n", mot)
			break
		}

		// Si le joueur entre une seule lettre
		if len(test) == 1 {
			lettre := rune(test[0]) // Convertir la chaîne en rune
			if strings.ContainsRune(mot, lettre) {
				estla[lettre] = true
			} else {
				if !estpas[lettre] {
					estpas[lettre] = true
					pv--
					fmt.Printf("La lettre %s n'est pas dans le mot\n", test)
					fmt.Printf("Il vous reste %d chances\n", pv)
				}
			}
		}

		// Afficher l'état actuel du mot
		fmt.Print("Mot à trouver : ")
		for _, char := range mot {
			if estla[char] {
				fmt.Printf("%c ", char)
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()

		// Vérifier la victoire
		if checkWin(mot, estla) {
			fmt.Printf("Bravo ! Tu as deviné le mot : %s\n", mot)
			break
		}

		// Si le joueur perd
		if pv == 0 {
			fmt.Println("Vous avez perdu")
			fmt.Println("Le mot à trouver :", mot) // bite
			fmt.Printf("   _________\n   ||/     |\n   ||      O\n   ||     /|\\\n   ||     / \\\n   ||\n__/||\\__________\n")
			break
		}
	}
}

// Fonction pour vérifier la victoire
func checkWin(mot string, estla map[rune]bool) bool {
	for _, char := range mot {
		if !estla[char] {
			return false
		}
	}
	return true
}
