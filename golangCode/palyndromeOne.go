package main

package palyndrome

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	var wordOne = "tenet"

	wordToCheck := toSplit(wordOne)
	isPalyndrome(wordToCheck)
	// splitted := strings.Split(wordOne, "")
	// fmt.Println(splitted)
	// fmt.Println(reflect.TypeOf(splitted))
}

func toSplit(word string) []string {
	wordToSplit := strings.Split(word, "")
	// fmt.Println(wordToSplit[0])
	fmt.Println(wordToSplit)
	return wordToSplit

}

func isPalyndrome(wordToCheck []string) bool {
	for i := 0; i < len(wordToCheck)+1; i++ {
		if reflect.DeepEqual(wordToCheck[i:], wordToCheck[:1]) {
			// fmt.Println(wordToCheck)
			fmt.Println("true")
			return true
		}
	}
	fmt.Println("false")
	return false
}
