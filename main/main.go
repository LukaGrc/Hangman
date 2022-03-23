// ------------------------------------------------------ //
// ---------------------{ HANGMAN }---------------------- //
// ---------------{ Ynov Informatique B1 }--------------- //
// ------------------------------------------------------ //
// -----------{ Eric }---{ Victor }---{ Luka }----------- //
// ------------------------------------------------------ //

package main

// ---------------------------------------------- //
// ---------------{ Importations }--------------- //
// ---------------------------------------------- //

import (
	"fmt"
	"hangman"
	"os"
	"strconv"
)

// ---------------------------------------------- //
// -----------{ Fonctions principales }---------- //
// ---------------------------------------------- //

func main() {
	args := os.Args

	if len(args) == 2 {
		// Initialisation du jeu et des mots
		hangman.JeuEnCours = hangman.Game{
			VraiMot:          hangman.GetRandomWord(args[1]),
			NbErrors:         0,
			LettresUtilisees: []rune{},
		}
		hangman.HideWord()
		hangman.SetRandomLetters() // On met quelques lettres dans notre mot pour aider

		LaunchGame()
	} else if len(args) == 3 && args[1] == "--startWith" {
		if args[2][:6] == "saves/" {
			hangman.LoadSave(args[2][6:])
		} else if args[2][:7] == "/saves/" {
			hangman.LoadSave(args[2][7:])
		} else {
			hangman.LoadSave(args[2])
		}
		LaunchGame()
	} else {
		fmt.Println("⚠️  Erreur : Arguments invalides")
		fmt.Println("Pour lancer un jeu, vous devez faire : go run main/main.go (fichier_de_mots)")
		fmt.Println("Pour ouvrir une sauvegarde, vous devez faire : go run main/main.go --startWith (fichier_de_sauvegarde)")
	}
}

func LaunchGame() {
	// Début du jeu

	fmt.Println("\033[2J")
	fmt.Println(" //===================\\\\")
	fmt.Println("||       HANGMAN       ||")
	fmt.Println(" \\\\===================//")
	fmt.Println("")
	fmt.Println("Good luck, you have " + strconv.Itoa(10-hangman.JeuEnCours.NbErrors) + " attempts.")
	hangman.PrintMotDeJeu()
	rec := Recursive()
	if rec == 3 {
		fmt.Println("")
		saveId := hangman.SaveToFile()
		fmt.Println("🎮 Game Saved in saves/save-" + saveId + ".txt 🎮")
		fmt.Println("To launch the save : use --startWith save-" + saveId + ".txt")
		fmt.Println("")
	} else if rec == 1 {
		fmt.Println("")
		fmt.Println("Congrats !")
		fmt.Println("")
	} else if rec == 2 {
		fmt.Println("")
		fmt.Println("Plus de tentatives. You lose !")
		fmt.Println("Le mot était : " + hangman.JeuEnCours.VraiMot)
		fmt.Println("")
	}
}

// Fonction de jeu principale
// Retours possibles :
// 1 = Le mot est complet, c'est gagné ! We did it !
// 2 = Le nombre d'erreurs max est atteint, c'est perdu !
// 3 = Le jeu a été quitté et sauvegardé. A bientôt !
func Recursive() int {
	if hangman.IsFull() {
		return 1
	}
	if hangman.JeuEnCours.NbErrors > 9 {
		return 2
	} else {
		reponse := hangman.ToUpper(string(hangman.AskFor())) // On demande une entrée (mot/lettre) au joueur
		if len(reponse) == 0 {
			fmt.Println("Erreur : Vous devez entrer une lettre ou un mot.")
			fmt.Println("")
		} else if len(reponse) > 1 { // Si c'est un mot
			if hangman.IsOnlyLetters(reponse) {
				if reponse == "STOP" {
					return 3
				} else if reponse == hangman.JeuEnCours.VraiMot {
					return 1
				} else {
					if hangman.JeuEnCours.NbErrors < 9 {
						hangman.JeuEnCours.NbErrors++
					}
					hangman.PrintPendu()
					hangman.JeuEnCours.NbErrors++
				}
			} else {
				fmt.Println("Erreur : Vous devez entrer une lettre ou un mot uniquement.")
				fmt.Println("⚠️  Les accents ne comptent pas ! ⚠️")
				fmt.Println("")
			}
		} else { // Si c'est une lettre
			if hangman.IsOnlyLetters(reponse) {
				letter := rune(reponse[0])
				if !(hangman.IsAlreadyUsed(letter)) { // La lettre n'a jamais été utilisée ?
					if len(hangman.IsInWord(letter)) > 0 { // La lettre existe-t-elle dans le mot ?
						for _, index := range hangman.IsInWord(letter) {
							hangman.SetLetter(index)
							hangman.JeuEnCours.LettresUtilisees = append(hangman.JeuEnCours.LettresUtilisees, letter)
						}
					} else {
						hangman.PrintPendu()
						hangman.JeuEnCours.LettresUtilisees = append(hangman.JeuEnCours.LettresUtilisees, letter)
						hangman.JeuEnCours.NbErrors++
					}
				} else {
					fmt.Println("Erreur : lettre déjà testée.")
					fmt.Println("")
				}
			} else {
				fmt.Println("Erreur : Vous devez entrer une lettre ou un mot uniquement.")
				fmt.Println("⚠️  Les accents ne comptent pas ! ⚠️")
				fmt.Println("")
			}
		}

		hangman.PrintMotDeJeu()
		return Recursive()
	}
}
