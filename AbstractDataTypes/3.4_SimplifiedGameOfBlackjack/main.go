package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var suits = []rune{'\u2660', '\u2661', '\u2662', '\u2663'}

type Card struct {
	Rank string
	Suit string
}

type Hand struct {
	Cards []Card
}

type Deck struct {
	Cards []Card
}

func (hand Hand) value() int {
	result := 0
	numberAces := 0
	for index := 0; index < len(hand.Cards); index++ {
		if hand.Cards[index].Rank != "A" &&
			hand.Cards[index].Rank != "K" &&
			hand.Cards[index].Rank != "Q" &&
			hand.Cards[index].Rank != "J" {
			intVal, _ := strconv.Atoi(hand.Cards[index].Rank)
			result += intVal
		} else if hand.Cards[index].Rank == "J" ||
			hand.Cards[index].Rank == "Q" ||
			hand.Cards[index].Rank == "K" {
			result += 10
		} else if hand.Cards[index].Rank == "A" {
			result += 11
			numberAces += 1
		}
	}

	if result > 21 && numberAces > 1 {
		result -= 10 * numberAces
	}

	return result
}

func (hand *Hand) addCard(card Card) {
	hand.Cards = append(hand.Cards, card)
}

func (hand Hand) Display() {
	fmt.Println("\n")
	for _, card := range hand.Cards {
		fmt.Print(card.Rank + card.Suit + " ")
	}
}

func (deck *Deck) dealCard() Card {
	result := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return result
}

func (deck *Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

func (deck *Deck) initializeDeck() Deck {
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.Cards = append(deck.Cards, Card{rank, string(suit)})
		}
	}
	deck.shuffle()
	return *deck
}

func (deck Deck) display() {
	for _, card := range deck.Cards {
		fmt.Print(card.Rank + card.Suit + " ")
	}
}

func main() {
	gameOver := false
	myDeck := Deck{}
	myDeck.initializeDeck()
	houseHand := Hand{}
	playerHand := Hand{}

	for i := 1; i <= 2; i++ {
		card := myDeck.dealCard()
		houseHand.addCard(card)

		card = myDeck.dealCard()
		playerHand.addCard(card)
	}
	playerHand.Display()
	fmt.Println("     Do you want to be hit (y/n)?")
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := reader.ReadRune()
	for {
		if res != 'y' {
			break
		}
		card := myDeck.dealCard()
		playerHand.addCard(card)
		playerHand.Display()

		if playerHand.value() > 21 {
			fmt.Println("PLAYER'S SCORE EXCEEDS 21. GAME OVER. HOUSE WINS!")
			gameOver = true
			break
		}
		fmt.Println("     Do you want to be hit (y/n)?")
		reader = bufio.NewReader(os.Stdin)
		res, _, _ = reader.ReadRune()
	}

	if !gameOver {
		for {
			if houseHand.value() > 21 {
				fmt.Println("HOUSE SCORE EXCEEDS 21. GAME OVER. PLAYER WINS!")
				gameOver = true
				break
			}
			if houseHand.value() < 17 {
				card := myDeck.dealCard()
				houseHand.addCard(card)
			} else {
				break
			}
		}
	}
	if !gameOver {
		if playerHand.value() > houseHand.value() {
			fmt.Println("PLAYER SCORE EXCEEDS HOUSE SCORE. GAME OVER. PLAYER WINS!")
		} else if playerHand.value() == houseHand.value() {
			fmt.Println("PLAYER SCORE EQUALS HOUSE SCORE. GAME OVER. TIE GAME!")
		} else {
			fmt.Println("HOUSE SCORE EXCEEDS PLAYER SCORE. GAME OVER. HOUSE WINS!")
		}
	}
}
