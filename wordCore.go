// ------------------------------------------------------ //
// ---------------------{ WORD CORE }-------------------- //
// -----------------{ Gestion des mots }----------------- //
// ------------------------------------------------------ //
package hangman

// ---------------------------------------------- //
// ---------------{ Importations }--------------- //
// ---------------------------------------------- //

import (
	"math/rand"
	"time"
)

// ---------------------------------------------- //
// --------------{ Vérifications }--------------- //
// ---------------------------------------------- //

// Vérifier que le string est en caractères minuscules
func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] <= 122 && s[i] >= 97) {
			return false
		}
	}
	return true
}

// Vérifier que le string ne contient que des caractères
func IsOnlyLetters(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] <= 90 && s[i] >= 65) && !(s[i] <= 122 && s[i] >= 97) {
			return false
		}
	}
	return true
}

// Vérifier qu'une lettre passée en argument est cachée
func IsHidden(lettre rune) bool {
	return lettre == '_'
}

// Vérifier qu'une lettre a déjà été utilisée
func IsAlreadyUsed(lettre rune) bool {
	for _, item := range JeuEnCours.LettresUtilisees {
		if item == lettre {
			return true
		}
	}
	return false
}

// Vérifier que la lettre est présente dans le mot
func IsInWord(lettre rune) []int {
	res := []int{}
	for index, item := range JeuEnCours.VraiMot {
		if item == lettre {
			res = append(res, index)
		}
	}
	return res
}

// Vérifier que le mot est complété
func IsFull() bool {
	for _, item := range JeuEnCours.MotDeJeu {
		if item == '_' {
			return false
		}
	}
	return true
}

// ---------------------------------------------- //
// ------------{ Gestion des mots }-------------- //
// ---------------------------------------------- //

// Mettre en caractères majuscules un string
func ToUpper(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if IsLower(string(s[i])) {
			result = result + string(s[i]-32)
		} else {
			result = result + string(s[i])
		}
	}
	return result
}

// Récupérer un mot aléatoire dans le fichier en argument
func GetRandomWord(fileName string) string {
	rand.Seed(time.Now().UnixNano())
	line := rand.Intn(len(GetFileContent(fileName)))
	return ToUpper(GetFileContent(fileName)[line])
}

// Cacher complètement le mot
func HideWord() {
	for range JeuEnCours.VraiMot {
		JeuEnCours.MotDeJeu = append(JeuEnCours.MotDeJeu, '_')
	}
}

// Afficher la lettre à l'index donné
func SetLetter(index int) {
	if IsHidden((JeuEnCours.MotDeJeu)[index]) {
		JeuEnCours.MotDeJeu[index] = rune((JeuEnCours.VraiMot)[index])
	}
}

// Afficher aléatoirement des lettres dans le mot
func SetRandomLetters() {
	rand.Seed(time.Now().UnixNano())
	numberOfLetters := len(JeuEnCours.MotDeJeu)/2 - 1
	var randomIndex int
	for i := 0; i < numberOfLetters; i++ {
		randomIndex = rand.Intn(len(JeuEnCours.MotDeJeu))
		SetLetter(randomIndex)
	}
}
