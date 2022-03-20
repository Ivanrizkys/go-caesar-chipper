package main

import (
	"bufio"
	"caesar-chipper/helper"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	alg := flag.String("alg", "no-modulo", "Enter algoritma here")
	flag.Parse()

	Run(*alg)
}

func Run(alg string) {

	var str string
	var round int
	var method int
	var exit string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Pilih salah satu metode di bawah ini \n1.Encrypt \n2.Decrypt")
	fmt.Print("Pilihan anda: ")
	fmt.Scanln(&method)
	if method != 1 && method != 2 {
		fmt.Println("Input yang anda masukkan salah silahkan ulangi lagi :)")
		fmt.Println("====================")
		fmt.Println(" ")
		Run(alg)
		return
	}
	fmt.Print("Kata: ")
	scanner.Scan()
	str = scanner.Text()
	fmt.Print("Key: ")
	fmt.Scanln(&round)

	chipperText := Chipper(str, round, method, alg)
	fmt.Println("Hasil")
	fmt.Println(chipperText)
	fmt.Print("Exit (y/n): ")
	fmt.Scanln(&exit)
	if exitUpper := strings.ToUpper(exit); exitUpper != "Y" {
		fmt.Println("====================")
		fmt.Println(" ")
		Run(alg)
		return
	}
	fmt.Println("-------------------END-------------------")
}

func Chipper(str string, round int, method int, alg string) string {
	if algUpper := strings.ToUpper(alg); algUpper == "WITH-MODULO" {
		round %= 26
	}
	loweCaseString := strings.ToLower(str)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	arrAlphabet := strings.Split(alphabet, "")
	var newString string

	for i := 0; i < len(loweCaseString); i++ {
		currentLater := string(loweCaseString[i])
		// * jika character adalah spasi
		if currentLater == " " {
			newString += currentLater
			continue
		}
		// * mencari index dari character di dalam array alphabet
		currentIndex := helper.IndexOf(currentLater, arrAlphabet)
		var newIndex int
		if method == 1 {
			newIndex = currentIndex + round
		} else if method == 2 {
			newIndex = currentIndex - round
		} else {
			log.Fatal("Error sir")
		}
		// * jika index dari new index lebih dari jumlah angka di dalam alphabet maka dia akan membalikkanya (z -> d, round = 4)
		if newIndex > 25 {
			newIndex -= 26
		}
		// * jika index dari new index negatif maka akan ditambahkan dengan 26
		if newIndex < 0 {
			newIndex += 26
		}
		// * jika string huruf kapital
		if strUpper := strings.ToUpper(string(str[i])); string(str[i]) == strUpper {
			strUpperResult := strings.ToUpper(arrAlphabet[newIndex])
			newString += strUpperResult
			continue
		}
		newString += arrAlphabet[newIndex]
	}

	return newString
}
