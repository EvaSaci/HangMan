package main

import (
    "fmt"
)

func main() {
    // Définition des variables
    pv := 5         // Nombre de vies(5)
    mot := "hello"  // Le mot à deviner
    var test string // Stocke la lettre ou mot entrée par le joueur
    var estla bool  // Indique si la lettre est présente dans le mot

    fmt.Println("Bienvenue dans le jeu Hangman !\n")
    fmt.Println("tu veut y jouer ?\n")
    fmt.Println("(si tu veut pas t'es gay)\n")
    for {
        // Réinitialiser le flag estla à false avant chaque tentative
        estla = false
        fmt.Print("Entrez une lettre ou un mot: ")
        fmt.Scan(&test)
        if test == "gay" {

            fmt.Printf("la beuteu de Yann \n")
            break
        }
        if len(test) != 1 && test != mot {
            pv--
            fmt.Printf("Il vous reste %d chances \n", pv)
        }
        if len(test) != 1 && test == mot {
            fmt.Printf("GG ! Vous avez trouvé le mot : %s \n", mot)
            break
        }
        // Boucle à travers chaque lettre du mot
        for _, i := range mot {
            if string(i) == test {
                fmt.Printf("%s", test)
                estla = true
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Print("\n")
        if !estla {
            pv--
            fmt.Printf("La lettre %s n'est pas dans le mot\n", test)
            fmt.Printf("Il vous reste %d chances \n", pv)
        }
        if pv == 0 {
            fmt.Println("Vous avez perdu")
            break
        }

    }
}