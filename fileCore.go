// ------------------------------------------------------ //
// ---------------------{ FILE CORE }-------------------- //
// -----------{ Gestion des fichiers & données}---------- //
// ------------------------------------------------------ //
package hangman

// ---------------------------------------------- //
// ---------------{ Importations }--------------- //
// ---------------------------------------------- //

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

// ---------------------------------------------- //
// -------------{ Structure de Jeu }------------- //
// ---------------------------------------------- //

type Game struct {
	VraiMot          string
	MotDeJeu         []rune
	NbErrors         int
	LettresUtilisees []rune
}

var JeuEnCours Game

// ---------------------------------------------- //
// -------{ Gestion des fichiers de mots }------- //
// ---------------------------------------------- //

// Récupérer le contenu du fichier (liste de strings : chaque ligne est un élément)
func GetFileContent(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// ---------------------------------------------- //
// ----{ Gestion des fichiers de sauvegarde }---- //
// ---------------------------------------------- //

// Récupérer le nombre de sauvegardes existantes
func NbOfSaves() int {
	files, err := os.ReadDir("saves")
	if err != nil {
		return 0
	}
	return len(files)
}

// Créer une nouvelle sauvegarde
func SaveToFile() string {
	saveNumber := strconv.Itoa(NbOfSaves() + 1)
	file, _ := os.Create("saves/save-" + saveNumber + ".txt")

	save, _ := json.Marshal(JeuEnCours)

	save = append(save, '\n')

	file.Write(save)
	file.Close()

	return saveNumber
}

// Récupérer et lire une sauvegarde
func LoadSave(fileName string) {
	content, err := os.ReadFile("saves/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(content, &JeuEnCours)
}
