package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	dealerThreshold = 16 //Determines when the dealer will stick
	bustThreshold   = 21 //Determines when the player will bust
	randomThreshold = 12 //Determines max values for random number generation
)

var numberChips int
var playerName string
var playerScore int
var dealerScore int

func main() {
	fmt.Printf("Please enter your name:\n")
	fmt.Scan(&playerName)
	blackJack()
}

func blackJack() {
	numberChips = 100
	playerResponse := "Y"
	fmt.Printf("\nWelcome " + playerName + ", are you ready to play Blackjack ? (Y/N) = ")
	fmt.Scan(&playerResponse)
	if strings.EqualFold(playerResponse, "Y") {
		startGame()
	} else if strings.EqualFold(playerResponse, "N") {
		exitGame()
	} else {
		fmt.Printf("Please, enter 'Y' for yes or 'N' for no\n")
		blackJack()
	}
}

func exitGame() {
	fmt.Printf("\nGoodbye " + playerName + " ! See you soon ;)\n")
	os.Exit(0)
}

func startGame() {
	playerScore = 0
	dealerScore = 0
	valueBet := placeBet()
	if valueBet == -1 {
		color.Red("\nUnlucky, You lost all your chips...\n")
		exitGame()
	}
	if playersTurn() {
		if dealerTurn() {
			color.Red("\n\nYou lost, you had " + strconv.Itoa(playerScore) + " against " + strconv.Itoa(dealerScore) + "...\n")
			numberChips = numberChips - valueBet
		} else {
			color.Green("\n\nCongrats " + playerName + ", you won !\n")
			numberChips = numberChips + valueBet
		}
	} else {
		numberChips = numberChips - valueBet
	}
	startGame()
}

func placeBet() int {
	newBet := "0"
	bet := 0
	if numberChips == 0 {
		return -1
	}
	color.Green("\nYou have " + strconv.Itoa(numberChips) + " coins\n")
	fmt.Printf("Please, place bet :\n")
	fmt.Scan(&newBet)
	bet, err := strconv.Atoi(newBet)
	checkBet(err)
	if bet > numberChips || bet < 0 {
		fmt.Printf("\nYou can't bet that amount of chips...\n")
		placeBet()
	}
	return bet
}

func checkBet(e error) {
	if e != nil {
		fmt.Printf("\nPlease enter a valid amount of chips !\n")
		placeBet()
	}
}

func newCardGenerator() int {
	var randomNbr int
	randomSeed := rand.NewSource(time.Now().UnixNano())
	seededRandomGen := rand.New(randomSeed)
	randomNbr = seededRandomGen.Intn(randomThreshold) + 1
	if randomNbr > 10 {
		randomNbr = 10
	}
	return randomNbr
}

//return false if the player has busted
func playersTurn() bool {
	hit := "h"
	fmt.Printf("\n")
	for userRequestCard := true; userRequestCard; userRequestCard = (!strings.EqualFold(hit, "s")) {
		currentCard := newCardGenerator()
		playerScore += currentCard
		if playerScore > bustThreshold {
			fmt.Print("You busted...\n")
			hit = "s"
			return false
		}
		fmt.Print("You have " + strconv.Itoa(playerScore) + " would you like to (s)tick or (h)it ? ")
		fmt.Scan(&hit)
	}
	return true
}

//return true if the dealer wins
func dealerTurn() bool {
	currentCard := newCardGenerator()
	dealerScore += currentCard
	fmt.Printf("\nDealer has " + strconv.Itoa(dealerScore) + " ;")
	if checkWinners() {
		return true
	}
	for dealerScore < dealerThreshold {
		fmt.Printf(" dealer hits\n")
		currentCard = newCardGenerator()
		dealerScore += currentCard
		if dealerScore > 21 {
			fmt.Printf("Dealer bust !")
			return false
		}
		fmt.Printf("Dealer has " + strconv.Itoa(dealerScore) + " ;")
		if checkWinners() {
			return true
		}
	}
	if checkWinners() {
		return true
	}
	return false
}

//return true if the dealer wins
func checkWinners() bool {
	if dealerScore > playerScore {
		return true
	} else if dealerScore == playerScore {
		return true
	} else {
		return false
	}
}
