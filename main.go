package main

import (
	"fmt"
	"time"

	"github.com/akinobufujii/advent-calendar-counter/qiita"
)

// requestItem Qiita記事を非同期で取得
func requestItem(itemID string, ch chan qiita.Item) {
	data, err := qiita.RequestItem(itemID)

	if err != nil {
		fmt.Printf("Qiita記事取得失敗 %s\n", err)
		close(ch)
	}

	time.Sleep(time.Second * 3)

	ch <- data
}

func main() {
	fmt.Println("Start")

	fmt.Println("Request Start")
	ch := make(chan qiita.Item)
	go requestItem("50d605ecf22daa3309cb", ch)
	fmt.Println("Request Wait")
	item := <-ch

	fmt.Println("Request Done Show Info")
	fmt.Println(item)

	fmt.Println("End")
}
