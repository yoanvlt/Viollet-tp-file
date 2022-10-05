package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"strconv"
	"math/rand"
	"time"
)	

func main(){
	firstWord()
	lastWord()
	afterWord("before")
	dernier("now ")
	random()
	
}


func firstWord() {
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		fmt.Println(scan.Text())
		break
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}

func lastWord() {
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	var dernier_mot string
	for scan.Scan() {
		dernier_mot = scan.Text()
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dernier_mot)
}

func afterWord(mot string) {
    var afterWord string
	nmot := ""
    file, err := os.Open("Hangman.txt")
    if err != nil {
        log.Fatalf("Error when opening file: %s", err)
    }
    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        if fileScanner.Text() == mot {
            fileScanner.Scan()
            afterWord = fileScanner.Text()
			i,_ := strconv.Atoi(afterWord)
			nmot = nthWord(i)
            break
        }
    }
    if err := fileScanner.Err(); err != nil {
        log.Fatalf("Error while reading file: %s", err)
    }
    file.Close()
    fmt.Println(nmot)
}

func nthWord(n int) string{
	var nthWord string
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if n == 1 {
			nthWord = fileScanner.Text()
			break
		}
		n--
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()
	return nthWord
}

func beforeWord(mot string) string{
    var beforeWord string
    nmot := ""
    file, err := os.Open("Hangman.txt")
    if err != nil {
        log.Fatalf("Error when opening file: %s", err)
    }
    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        if fileScanner.Text() == mot {
            beforeWord = nmot
            break
        }
        nmot = fileScanner.Text()
    }
    if err := fileScanner.Err(); err != nil {
        log.Fatalf("Error while reading file: %s", err)
    }
    file.Close()
    return beforeWord
}

func dernier(mot string){
	nmot := beforeWord(mot)
	r := []rune(nmot)
	char := r[1]
	res := int(char/38)
	newmot := nthWord(res)
	fmt.Println(newmot)
}


func random() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	fmt.Println(i)
}
	