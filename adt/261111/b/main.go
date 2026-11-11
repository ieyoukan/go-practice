package main

import (
	"errors"
	"fmt"
)

func IsRated(x int, r int) (bool, error) {
	switch x {
	case 1:
		if 1600 <= x && x <= 2999 { //超うっかりミスでrなのにxって書いてる これを防ぐにはテスト書くのがいいかな。名前ちゃんとつけるとか
			return true, nil
		} else {
			return false, nil
		}
	case 2:
		if 1200 <= x && x <= 2399 {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, errors.New("")
}
func main() {
	var r, x int
	fmt.Scan(&r, &x)
	israted, err := IsRated(x, r)
	if err != nil {
		fmt.Println("エラー:", err)
	}
	if israted {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
