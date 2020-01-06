package qiita

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RequestItem QiitaのアイテムをREST APIを使用して情報を取得する
func RequestItem(itemID string) (Item, error) {
	response, err := http.Get("https://qiita.com/api/v2/items/" + itemID)
	if err != nil {
		return Item{}, err
	}

	data, _ := ioutil.ReadAll(response.Body)

	var item Item
	err = json.Unmarshal(data, &item)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}
