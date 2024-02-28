package main

import (
	"fmt"
	"strings"
	"math/rand"
)

const original = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetter(key int,letter string) (result string){
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOverLetters := string(runes[0:len(letter)-key])
	return fmt.Sprintf(`%s%s`,lastLetterKey,leftOverLetters)
}

func encrypt(key int,plainText string ) (result string){
	hashletter := hashLetter(key,original)
	var hashstring = ""
	findOne := func(r rune) rune {
		pos := strings.Index(original,string([]rune{r}))
		if pos != -1{
			letterPosition := (pos + len(original)) % len(original)
			hashstring = hashstring + string(hashletter[letterPosition])
			return r
		}
		return r
	}
		strings.Map(findOne,plainText)
		return hashstring
}

func decrypt(key int,encryptedText string) (result string){
	hashletter := hashLetter(key,original)
	var hashString = ""

	findOne := func(r rune) rune{
		pos := strings.Index(hashletter,string([]rune{r}))
		if pos != -1{
			letterPosition := (pos + len(original)) % len(original)
			hashString = hashString + string(original[letterPosition]) 
			return r
		}
		return r
	}

	strings.Map(findOne,encryptedText)
	return hashString
}

func main(){
	var plainText string 
	fmt.Println("Enter a plain text")
	fmt.Scan(&plainText)
	text := strings.ToUpper(plainText)
	fmt.Println("Plain Text",text)
	min := 0
	max := 9
	key := rand.Intn(max-min)+min;
	fmt.Println("Key value is :",key)
	encrypted := encrypt(key,text)
	fmt.Println("Encrypted Text :",encrypted)
	decrypted := decrypt(key,encrypted)
	fmt.Println("Decrypted Text :",decrypted)
}