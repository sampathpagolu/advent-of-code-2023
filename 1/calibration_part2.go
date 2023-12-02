package main


import (
    "fmt"
    "log"
    "os"
    "bufio"
    "regexp"
    "strconv"
    "strings"    
)

var numStrs = []string{
    "one", 
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "zero",
}

var debug bool
func main() {
    if debug {
        fmt.Println(getNumbersFromStringToo("274"))
    } else {
        content, err := os.Open("input.txt")
        if err != nil {
            log.Fatal(err)
            return 
        }
        defer content.Close()
        scanner := bufio.NewScanner(content)
        sum := 0
        for scanner.Scan() {
            num := getNumbersFromStringToo(scanner.Text())

            sum += num
           // fmt.Printf("text: %v, numbers: %v \n", scanner.Text(), num)
        }
        fmt.Println(sum)
    }
}

func getNumbersFromStringToo(txt string) int {

    /*
    get the substrings of lengths 3,4,5 and check if a number is present if present return number. Else if it is one of the NumMap. Return the int value of num
    */
    lenTxt := len([]rune(txt))
    first, firstOk := findFirstNumber(txt, lenTxt)
    second, secondOk := findSecondNumber(txt, lenTxt)
   

    if firstOk && secondOk {
        return (first * 10) + second
    } else if firstOk && !secondOk{
        return first * 10 + first 
    } else if secondOk && !firstOk{
        return second * 10 + second
    } else {
        return 0
    }

}

func findFirstNumber(txt string, length int)  (int, bool){
    if debug {fmt.Println("In findings first")}
    if length < 3 {
        return getNumbers(txt, false)
    }
    for i := 0; i < length; i++ {
        for j := 3; j <= 5 && i + j <= length; j++ {
            substr := txt[i:i+j]
            if  debug {fmt.Printf("first: %v\n", substr)}
            num, ok := getNumbers(substr, false)
            if ok {return num, ok}
            for num, numStr := range numStrs {
                if strings.Contains(substr, numStr) {
                    return num+1, true
                }
            }
        }
    }
    return 100, false
}

func findSecondNumber(txt string, length int)  (int, bool){
    if debug {fmt.Println("In findings second")}
    if length < 3 {
        return getNumbers(txt, true)
    }
    for i := length; i >= 0; i-- {
        for j := 3; j <= 5 && i - j >= 0; j++ {
            substr := txt[i-j:i]
            if debug {fmt.Printf("reverse substring: %v\n", substr)}
            num, ok := getNumbers(substr, true)
            if ok {return num, ok}
            for num, numStr := range numStrs {
                if strings.Contains(substr, numStr) {
                    return num+1, true
                }
            }
        }
    }
    return 100, false
}

func getNumbers(txt string, reverse bool) (int, bool) {
    //ret := 0
    if debug {fmt.Printf("In finding int: %v\n", txt)}
    re := regexp.MustCompile("[0-9]")
    nums := re.FindAllString(txt, -1)
    numsLength := len(nums)
    if numsLength >= 1 && reverse {
        r, err := strconv.Atoi(nums[numsLength-1])
        if err == nil {
            if debug {fmt.Printf("found %d\n", r)}
            return r, true
        }
    } else if numsLength >=1 && !reverse {
        r, err := strconv.Atoi(nums[0])
        if err == nil {
            if debug {fmt.Printf("found %d\n", r)}
            return r, true
        }
    }
    return 100, false
}

