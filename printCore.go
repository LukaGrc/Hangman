// ------------------------------------------------------ //
// --------------------{ PRINT CORE }-------------------- //
// ----------{ Gestion des affichages & inputs }--------- //
// ------------------------------------------------------ //
package hangman

// ---------------------------------------------- //
// ---------------{ Importations }--------------- //
// ---------------------------------------------- //

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------- //
// ------------{ Gestion des input }------------- //
// ---------------------------------------------- //

// Demande une lettre au joueur
func AskFor() string {
	fmt.Println("")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Choose: ")
	scanner.Scan()
	return scanner.Text()
}

// ---------------------------------------------- //
// ---------{ Gestion des affichages }----------- //
// ---------------------------------------------- //

// Afficher le mot du jeu de manière formatée (avec espaces ...)
func PrintMotDeJeu() {
	for index, item := range JeuEnCours.MotDeJeu {
		if index != 0 {
			fmt.Printf(" ")
		}
		fmt.Print(string(item))
	}
	fmt.Printf("\n")
}

// Affiche José et le nombre d'essais restants
func PrintPendu() {
	fmt.Println("Not present in the word, " + strconv.Itoa(10-(JeuEnCours.NbErrors+1)) + " attempts remaining")
	if JeuEnCours.NbErrors == 0 {
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("         ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 1 {
		fmt.Println("         ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 2 {
		fmt.Println("  +---+  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 3 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 4 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  o   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 5 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  o   |  ")
		fmt.Println("  |   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 6 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  o   |  ")
		fmt.Println(" /|   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 7 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  o   |  ")
		fmt.Println(" /|\\  |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 8 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  o   |  ")
		fmt.Println(" /|\\  |  ")
		fmt.Println(" /    |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	} else if JeuEnCours.NbErrors == 9 {
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  o   |  ")
		fmt.Println(" /|\\  |  ")
		fmt.Println(" / \\  |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
		fmt.Println("")
	}
}
