package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func shuffler(cardList []string) []string {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(len(cardList))
	shuffledCards := []string{}
	for i := 0; i < len(permutation); i++ {
		shuffledCards = append(shuffledCards, cardList[permutation[i]])
	}
	return shuffledCards
}

func arrangeCards(cardList []string) [52]string {
	arrangedCards := [52]string{}
	for i := 0; i < len(cardList); i++ {
		if strings.Contains(cardList[i], "Heart") {
			if !(strings.Contains(cardList[i], "J") || strings.Contains(cardList[i], "Q") || strings.Contains(cardList[i], "K") || strings.Contains(cardList[i], "A")) {
				y, _ := strconv.Atoi(string(cardList[i][6:]))
				arrangedCards[y-2] = cardList[i]
			} else if strings.Contains(cardList[i], "J") {
				arrangedCards[9] = cardList[i]
			} else if strings.Contains(cardList[i], "Q") {
				arrangedCards[10] = cardList[i]
			} else if strings.Contains(cardList[i], "K") {
				arrangedCards[11] = cardList[i]
			} else if strings.Contains(cardList[i], "A") {
				arrangedCards[12] = cardList[i]
			}
		} else if strings.Contains(cardList[i], "Daimond") {
			if !(strings.Contains(cardList[i], "J") || strings.Contains(cardList[i], "Q") || strings.Contains(cardList[i], "K") || strings.Contains(cardList[i], "A")) {
				y, _ := strconv.Atoi(string(cardList[i][8:]))
				arrangedCards[11+y] = cardList[i]
			} else if strings.Contains(cardList[i], "J") {
				arrangedCards[12+10] = cardList[i]
			} else if strings.Contains(cardList[i], "Q") {
				arrangedCards[13+10] = cardList[i]
			} else if strings.Contains(cardList[i], "K") {
				arrangedCards[14+10] = cardList[i]
			} else if strings.Contains(cardList[i], "A") {
				arrangedCards[15+10] = cardList[i]
			}
		} else if strings.Contains(cardList[i], "Spad") {
			if !(strings.Contains(cardList[i], "J") || strings.Contains(cardList[i], "Q") || strings.Contains(cardList[i], "K") || strings.Contains(cardList[i], "A")) {
				y, _ := strconv.Atoi(string(cardList[i][5:]))
				arrangedCards[24+y] = cardList[i]
			} else if strings.Contains(cardList[i], "J") {
				arrangedCards[25+10] = cardList[i]
			} else if strings.Contains(cardList[i], "Q") {
				arrangedCards[26+10] = cardList[i]
			} else if strings.Contains(cardList[i], "K") {
				arrangedCards[27+10] = cardList[i]
			} else if strings.Contains(cardList[i], "A") {
				arrangedCards[28+10] = cardList[i]
			}
		} else if strings.Contains(cardList[i], "Club") {
			if !(strings.Contains(cardList[i], "J") || strings.Contains(cardList[i], "Q") || strings.Contains(cardList[i], "K") || strings.Contains(cardList[i], "A")) {
				y, _ := strconv.Atoi(string(cardList[i][5:]))
				arrangedCards[37+y] = cardList[i]
			} else if strings.Contains(cardList[i], "J") {
				arrangedCards[38+10] = cardList[i]
			} else if strings.Contains(cardList[i], "Q") {
				arrangedCards[39+10] = cardList[i]
			} else if strings.Contains(cardList[i], "K") {
				arrangedCards[40+10] = cardList[i]
			} else if strings.Contains(cardList[i], "A") {
				arrangedCards[41+10] = cardList[i]
			}
		}
	}
	fmt.Println(arrangedCards)
	return arrangedCards
}

func main() {
	cardList := []string{"Club_9", "Club_10", "Club_Jack", "Spad_Ace", "Spad_2", "Spad_3", "Spad_4", "Spad_5", "Spad_6", "Spad_7", "Daimond_2", "Daimond_3", "Daimond_4", "Daimond_5", "Daimond_6", "Daimond_7", "Daimond_8", "Daimond_9", "Daimond_10", "Daimond_Jack", "Daimond_Queen", "Daimond_King", "Heart_Ace", "Heart_2", "Heart_3", "Heart_4", "Heart_5", "Heart_6", "Heart_7", "Heart_8", "Heart_9", "Heart_10", "Heart_Jack", "Heart_Queen", "Heart_King", "Club_Ace", "Club_2", "Club_3", "Club_4", "Club_5", "Daimond_Ace", "Club_7", "Club_8", "Club_6", "Spad_8", "Spad_9", "Spad_10", "Spad_Jack", "Spad_Queen", "Spad_King", "Club_Queen", "Club_King"}
	cardList = shuffler(cardList)

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println()
	fmt.Println(cardList)
	newCards := arrangeCards(cardList)
	for i := 0; i < len(cardList); i++ {
		//fmt.Println(cardList[rand.Intn(10)])
		fmt.Println(newCards[i])
	}
	// rand.Seed(time.Now().Unix())
	// permutation := rand.Perm(52)
	// fmt.Println(len(permutation))
}
