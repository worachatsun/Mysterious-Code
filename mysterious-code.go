package main

import (
    b64 "encoding/base64"
    "fmt"
)

func main() {
    secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
    sd, _ := b64.StdEncoding.DecodeString(secret)
    stringSecret := string(sd)
    left := 0
    right := len(stringSecret)-1
    var firstHalf string = ""
    var secondHalf string = ""
	for left <= right {
            if (left == right) {
                firstHalf += string(stringSecret[right])
                break
            }
	    firstHalf += string(stringSecret[right])
	    secondHalf = string(stringSecret[left]) + secondHalf
	    left++
	    right--
	}
    fmt.Println(firstHalf + secondHalf)
}
