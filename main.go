package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
	"github.com/common-nighthawk/go-figure"
)

// Déclaration de difficulte
type diff int

// Déclaration des constantes pour la difficulté
const (
	facile    diff = iota
	moyen     diff = iota
	difficile diff = iota
	yann      diff = iota //
	raciste   diff = iota
	h              = iota //
)

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Gestion du flag pour la difficulté
	difficultyFlag := flag.String("diff", "moyen", "Choisir la difficulté (facile, moyen, difficile)")
	flag.Parse()

	var fichier *os.File
//	var f *os.File
	var err error

	// Gestion de la difficulté
	switch *difficultyFlag {
	case "facile": // mode facile
		fichier, err = os.Open("mots/motsFacile.txt") // ouvre le fichier des mots facile
		fmt.Println("mode de jeux : facile")
	case "moyen": // mode moyen
		fichier, err = os.Open("mots/motsNormal.txt") // ouvre le fichier des mots Moyen
		fmt.Println("mode de jeux : Moyen")
	case "difficile": // mode difficile
		fichier, err = os.Open("mots/motsHard.txt") // ouvre le fichier des mots Difficile
		fmt.Println("mode de jeux : Difficile")
	case "yann": // mode Yann
		fmt.Println("Mode de jeux : yann")
		fichier, err = os.Open("mots/motsYann.txt") // ouvre le fichier des mots de Yann
	case "raciste": // mode raciste
		fmt.Println("Mode de jeux : raciste")
		fichier, err = os.Open("mots/motsRaciste.txt") // ouvre le fichier des mots Raciste
	case "h": // mode help
		fmt.Println("mode HELP :")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("pour changer de mode : go run . -diff (facile; moyen; difficile)")
		fmt.Println("vous pouvez annulez le jeu a tous moment avec 'stop'")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("evidemment plusieurs easterEggs sont caché a vous de les trouvez !")
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
	VoF := true
	title := figure.NewFigure("Hangman !", "", true)
	fmt.Println("----------------------------------------------------------------")
	title.Print()
	fmt.Println("----------------------------------------------------------------")
	//fmt.Println("		Bienvenue dans le jeu Hangman !")
	//fmt.Println("----------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("plusieurs mode de jeux sont accessible")
	fmt.Println("--> go run . -diff h")
	fmt.Println("")
	fmt.Println("Prêt ?")
	fmt.Println("Tu peux annuler en écrivant 'stop'")
	fmt.Println("(Si tu veux pas y jouer t'es gay)")
	fmt.Println("----------------------------------------------------------------")

	for {
		cmt := 0
        cmt2 := 10 - pv + 1
        lignedebut := (10-pv)*8 + 1
        lignefin := 9 + lignedebut

		fmt.Print("Entrez une lettre ou un mot : ")
		//fmt.Println("")
		fmt.Scan(&test)

		// Vérification si tous les caractères sont valides (lettres et '-')
		if !valide(test) {
			fmt.Println("Mauvais caractères, veuillez réessayer") // Message affiché si caractère non valide
			continue
		}

		if test == "gay" {
			fmt.Println("la beuteu à Yann")
			break
		}
		if test == "non" {
			fmt.Println("t'es gay")
			break
		}
		if test == "stop" {
			fmt.Println("merci d'avoir joué !")
			fmt.Println("à bientôt !")
			break
		}

		// Si le joueur entre un mot entier
		if len(test) != 1 && test != mot {
			pv--
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
				VoF = false
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

		// Afficher les lettres incorrectes
		fmt.Print("Lettres incorrectes : ")
		for lettre := range estpas {
			fmt.Printf("%c ", lettre) 
		}
		fmt.Println() 

		// Vérifier la victoire
		if checkWin(mot, estla) {
			for pv >= 7 {
				fmt.Println("double Monstre pour avoir échoué moins de 3 fois !")
				break
			}
			for pv == 6 || pv == 5 {
				fmt.Println("en vrais GG mais prochaine fois fait mieux stp !")
				break
			}
			for pv == 4 || pv == 3 {
				fmt.Println("Aie, Aie, Aie, c'était chaud la non ? Bravo quand même")
				break
			}
			for pv == 2 || pv == 1 {
				fmt.Println("ouais bon la soit t'es nul soit le mot été dur... GG quand même")
				break
			}
			fmt.Printf("Bravo ! Tu as deviné le mot : %s\n", mot)
			break
		}
		f, err := os.Open("O.txt")
        if err != nil {
            //log.Fatal(err)
        }
		defer f.Close()
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            cmt++

            if cmt >= lignedebut && cmt <= lignefin {
                fmt.Println(scanner.Text())
            }
            if cmt%(8*cmt2) == 0 {
                if VoF {
                    cmt2++
                }
                break
            }
        }
		if err := scanner.Err(); err != nil {
            //log.Fatal(err)
        }
		// Si le joueur perd
		if pv == 0 {
			fmt.Println("Vous avez perdu")
			fmt.Println("Le mot à trouver :", mot) // bite
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

// Fonction pour valider l'entrée (seulement lettres de a à z et le tiret '-')
func valide(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) && char != '-' {
			return false
		}
	}
	return true
}
