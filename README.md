# Hangman - Documentation

Ce projet Hangman Classic a été entretenu par Luka GARCIA, Eric MAGNE et Victor DALAMEL DE BOURNET.

## Sommaire

- [I. A propos du projet](#i-a-propos-du-projet)
- [II. Utilisation de Hangman](#ii-utilisation-de-hangman)
- [III. Réalisation des consignes & bonus](#iii-réalisation-des-consignes-bonus)
    - [A. Consignes requises](#a-consignes-requises)
    - [B. Bonus 1 : advanced-feature](#b-bonus-1-advanced-feature)
    - [C. Bonus 2 : start-and-stop](#c-bonus-2-start-and-stop)
- [IV. Crédits](#iv-crédits)


## I. A propos du projet

Le projet Hangman Classic est un projet dont le but est de recréer le jeu du pendu en CLI (Command Line Interface) en utilisant un langage imposé : le [Golang](https://go.dev/).

Nous avons eu un autre projet dans lequel nous avons réutilisé cet algorithmen pour le mettre sous forme d'interface web : [hangman web](https://github.com/LukaGrc/Hangman-Web).

## II. Utilisation de Hangman

Pour lancer le projet, vous avez seulement besoin de cloner ou télécharger ce repository, ouvrir un terminal et vous diriger dans le dossier contenant le projet.
Ici, vous n'avez plus qu'à exécuter la commande `go run main/main.go [words.txt/words2.txt/words.txt]`.

## III. Réalisation des consignes & bonus

### A. Consignes requises

La consigne principale était de recréer le jeu du pendu en interface de commandes en utilisant le langage de programmation Golang.

---

### B. Bonus 1 : advanced-feature

Le but ici est d'implémenter de nouvelles fonctionnalités au projet :

D'une part, que le joueur puisse suggérer à la fois un mot, et à la fois une lettre (un mot est une chaîne de caractères d'au moins deux lettres de long) : si le mot est correct la partie se termine, sinon le compteur d'essais restants diminue de 2.

D'autre part, que les lettres déjà utilisées soient stockées de telle sorte à ce que le joueur ne puisse pas proposer deux fois la même chose.

---

### C. Bonus 2 : start-and-stop

Le bonus "start-and-stop" est un moyen de stopper le jeu et sauvagarder la progression (avec possibilité de relancer à tout moment la partie en cours).

Il nous a été demandé ici d'implémenter un mot-clé `STOP` lors des parties. Ce dernier arrête et quitte la partie tout en la sauvegardant dans un fichier au format `json`.

**À noter :** pour relancer une partie à partir d'une sauvegarde, remplacez la commande de lancement par : `go run main/main.go --startWith [fichierDeSauvegarde]`.

## IV. Crédits

Le projet Hangman Web a été conçu par Luka GARCIA, Éric MAGNE et Victor DALAMEL DE BOURNET (avec l'aide précieuse des mentors qui nous ont accompagnés).