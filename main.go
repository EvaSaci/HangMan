package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Ouvrir le fichier mots.txt
	fichier, err := os.Open("mots/mots.txt")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer fichier.Close() // Fermer le fichier à la fin du programme

	// Lire les mots dans un tableau
	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	// Vérifier s'il y a des mots dans la liste
	if len(mots) == 0 {
		fmt.Println("Le fichier ne contient rien")
		return
	}

	// Choisir un mot aléatoirement
	mot := mots[rand.Intn(len(mots))]

	pv := 10 //Nombre de vies(10)
	//mot := "evaprime"             //Le mot à deviner
	var test string               //Stocke la lettre ou mot entrée par le joueur
	estla := make(map[rune]bool)  //La bonne lettre
	estpas := make(map[rune]bool) //La mauvaiseeee

	fmt.Println("Bienvenue dans le jeu Hangman !")
	fmt.Println(" ")
	fmt.Println("Prêt ?")
	fmt.Println(" ")
	fmt.Println("tu peut annuler en écrivant 'non' ")
	fmt.Println(" ")
	fmt.Println("(Si tu veux pas y jouer t'es gay)")
	fmt.Println(" ")
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
		// ça c'est pour le mot entier
		if len(test) != 1 && test != mot {
			pv--
			fmt.Printf("Mot incorrect ! Il vous reste %d chances\n", pv)
		} else if len(test) != 1 && test == mot {
			fmt.Printf("bravo mec ! t'a trouvé le mot : %s\n", mot)
			break
		}
		// Si le mec entre une seule lettre
		if len(test) == 1 {
			lettre := rune(test[0]) // Convertir la chaîne en rune
			// Vérifier si la lettre est dans le mot
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
		// victoire t'a mère
		if checkWin(mot, estpas) {
			fmt.Printf("Bravo ! tu as deviné le mot : %s\n", mot)
			break
		}
		if pv == 0 {
			fmt.Println("Vous avez perdu")
			fmt.Println("le mot à trouver :", mot)
			fmt.Printf("   _________\n   ||/     |\n   ||      O\n   ||     /|\\\n   ||     /'\\\n   ||\n   ||\n__/||\\__________\n")
			break
		}
		if pv == 9 {
			fmt.Println("__/||\\__________")
		}
		if pv == 8 {
			fmt.Printf("   ||/     \n   ||      \n   ||     \n   ||     \n   ||\n   ||\n__/||\\__________\n")
		}
		if pv == 7 {
			fmt.Printf("   _________\n   ||/     \n   ||      \n   ||     \n   ||     \n   ||\n   ||\n__/||\\__________\n")
		}
		if pv == 6 {
			fmt.Printf("   _________\n   ||/     |\n   ||      O\n   ||     \n   ||     \n   ||\n   ||\n__/||\\__________\n")
		}
		if pv == 5 {
			fmt.Printf("   _________\n   ||/     |\n   ||      O\n   ||      |\n   ||      \n   ||\n   ||\n__/||\\__________\n")
		}
		if pv == 4 {
			fmt.Printf("   _________\n   ||/     |\n   ||      O\n   ||      |\\\n   ||      \n   ||\n   ||\n__/||\\__________\n")
		}
		if pv == 3 {
			fmt.Printf("   _________\n   ||/     |\n   ||      O\n   ||     /|\\\n   ||     \n   ||\n   ||\n__/||\\__________\n")
		}
		if pv == 2 {
			fmt.Printf("   _________\n   ||/     |\n   ||      O\n   ||     /|\\\n   ||      '\n   ||\n   ||\n__/||\\__________\n")
		}
		if pv == 1 {
			fmt.Printf("   _________\n   ||/     |\n   ||      O\n   ||     /|\\\n   ||      '\\\n   ||\n   ||\n__/||\\__________\n")
		}
	}
}
func checkWin(mot string, estla map[rune]bool) bool {
	for _, char := range mot {
		if !estla[char] {
			return false
		}
	}
	return true
}
