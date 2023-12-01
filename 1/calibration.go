package main


import (
   "fmt"
    "log"
    "os"
    "bufio"
    "regexp"
    "strconv"
    
)


func main() {
    content, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
        return 
    }
    defer content.Close()
    scanner := bufio.NewScanner(content)
    sum := 0
    for scanner.Scan() {
        sum += getNumbers(scanner.Text())
        //        fmt.Printf("text: %v, numbers: %v \n", scanner.Text(), getNumbers(scanner.Text()))
    }
    fmt.Println(sum)
}

func getNumbers(txt string) int {
    ret := 0
    re := regexp.MustCompile("[0-9]")
    nums := re.FindAllString(txt, -1)
    numsLength := len(nums)
    if numsLength >= 1 {
        first, _ := strconv.Atoi(nums[0])
        second, _ :=strconv.Atoi(nums[numsLength - 1])
        ret = (first * 10) + second
    }
    return ret 
}

