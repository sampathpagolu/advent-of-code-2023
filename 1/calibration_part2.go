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


func main() {
    //fmt.Println(getNumbersFromStringToo("275"))

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
        fmt.Printf("text: %v, numbers: %v \n", scanner.Text(), num)
        //fmt.Printf("text: %v: \t", scanner.Text())
        //fmt.Printf("numbers: %v \n", getNumbersFromStringToo(scanner.Text()))
    }
    fmt.Println(sum)
}


func getNumbersFromStringToo(txt string) int {
    numStrs := []string{
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
    numMap := map[string]int {
        "one": 1, 
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
        "zero": 10,
    }

    /*
    get the substrings of lengths 3,4,5 and check if a number is present if present return number. Else if it is one of the NumMap. Return the int value of num
    */
    lenTxt := len([]rune(txt))
    first, second := 100, 100
    //fmt.Println("finding first number")
    //// first number
    if lenTxt < 3 {
        intNum, success := getNumbers(txt, false)
        intNum2, success_ := getNumbers(txt, true)
        if success {
            v, _ := strconv.Atoi(intNum)
            first = v
        } else if success_ {
            
            v2, _ := strconv.Atoi(intNum2)
            second = v2
        }
    } else {

        firstNumFound := false
        for i := range(txt) {
            for j := 3; j <= 5; j++ {
                if  intNum, success := getNumbers(txt[i:j + i ], false); success {
                    val, _ := strconv.Atoi(intNum)
                    //fmt.Printf("found int, %d\n", val)
                    first = val
                    firstNumFound = true
                    break
                } else {
                    if (i+j) < lenTxt {
                        for _, numStr := range(numStrs){
                            // fmt.Println(numStr)
                            // fmt.Println(txt[i:i+j], numStr)
                            // fmt.Println(i, i+j)
                            // break
                            if strings.Contains(txt[i:j + i], numStr) {
                                first = numMap[numStr] 
                                firstNumFound = true
                                break
                            }
                        }
                        if  firstNumFound {
                            break
                        }
                    }
                    if  firstNumFound {
                        break
                    }
                }
                if  firstNumFound {
                    break
                }
            }
            if  firstNumFound {
                break
            }
        }

        //fmt.Println("finding second number")
        // second number
        secondNumFound := false
        for i:= range(txt) {
            for j := 3; j <= 5; j++ {
                reverseIndex := lenTxt  - i
                //fmt.Println(txt[reverseIndex -j:reverseIndex])
                if  intNum, success := getNumbers(txt[reverseIndex - j:reverseIndex], true); success {
                    val, _ := strconv.Atoi(intNum)
                    // fmt.Printf("found int, %d\n",val)
                    second = val
                    secondNumFound = true
                    break

                } else {
                    if  reverseIndex - j > 0  {
                        for _, numStr := range(numStrs){
                            //fmt.Println(txt[reverseIndex - j:reverseIndex], numStr)
                            //fmt.Println(reverseIndex - j , reverseIndex)
                            //break
                            if strings.Contains(txt[reverseIndex - j:reverseIndex], numStr) {
                                //fmt.Println("found the str")
                                second = numMap[numStr]
                                secondNumFound = true
                                break
                            }
                        }
                        if  secondNumFound {
                            break
                        }
                    }
                    if  secondNumFound {
                        break
                    }
                }
                if  secondNumFound {
                    break
                }
            }
            if  secondNumFound {
                break
            }
        }
   }
    if first <= 9 && second <= 9{
        return (first * 10) + second
    } else if first <= 9 && second > 9{
        return first * 10 + first 
    } else if second <= 9 && first > 9{
        return second * 10 + second
    } else {
        return 0
    }

}


func getNumbers(txt string, reverse bool) (string, bool) {
    //ret := 0
    re := regexp.MustCompile("[0-9]")
    nums := re.FindAllString(txt, -1)
    numsLength := len(nums)
    if numsLength >= 1 && reverse {
        return nums[len(nums)-1], true
    } else if numsLength >=1 && !reverse {
        return nums[0], true
    }
    return "a", false
}

