package main

import (
	//"fmt"
	"golang.org/x/tour/wc"
	//"strings"
  "regexp"
)

func WordCount(s string) map[string]int{
  wordsMap := make(map[string]int)

  // \b 単語の境界を表す
  // + 1回以上の繰り返し
  // \w+ 「1つ以上の単語文字」を表します。単語文字とは、a〜z、A〜Z、0〜9、そしてアンダースコア _ のことです。例えば、"hello" や "world123" は \w+ にマッチします。
  // [!\?,.;'] 感嘆符 !, 疑問符 ?, カンマ ,, ピリオド ., セミコロン ;, アポストロフィ ' のいずれか1文字。
  // | または
  re := regexp.MustCompile(`\b\w+('\w+)*([!\?,.;']|\b)`)
  wordsArr := re.FindAllString(s, -1)

  for _, word := range wordsArr {
    wordsMap[word]++
  }
  return wordsMap
}

func main() {
	wc.Test(WordCount)
}
