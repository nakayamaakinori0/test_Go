
package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fmt.Println("input:", s)
	
	wordsArr := strings.Split(s, " ")
	fmt.Println("wordsArr:", wordsArr)

	wordsMap := make(map[string]int)
	
	for i := 0; i < len(wordsArr); i++ {
		wordsMap[wordsArr[i]] = 0
	}
	fmt.Println("wordsMap:",wordsMap)
	for j := 0; j < len(wordsArr); j++{
		_, ok := wordsMap[wordsArr[j]]
		if(ok){
			wordsMap[wordsArr[j]] = wordsMap[wordsArr[j]] + 1
		}
	}
	fmt.Println("wordsMap",wordsMap)

	return wordsMap
}

func main() {
	wc.Test(WordCount)
}
