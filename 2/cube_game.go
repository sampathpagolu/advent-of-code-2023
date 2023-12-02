package main


import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
    "strconv"
)

var debug bool = false

var TOTAL_RED, TOTAL_GREEN, TOTAL_BLUE int= 12,13,14


func main() {

    if !debug {
        content, err := os.Open("input.txt")
        if err != nil {
            log.Fatal(err)
            return 
        }
        defer content.Close()
        scanner := bufio.NewScanner(content)
        sum := 0
        for scanner.Scan() && len(scanner.Text()) > 0{
            gameId, red, green, blue :=   parseInputString(scanner.Text())
            validGameId, ok := isCubeGameValid(gameId, red, green, blue)
            if ok {
                fmt.Println("ValidGameID: ",validGameId)
                sum += validGameId
            }
            // fmt.Printf("text: %v, numbers: %v \n", scanner.Text(), num)
        }
        fmt.Println(sum)
    } else {
        //fmt.Println(
        gameId, red, green, blue := parseInputString("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 19 green; 5 green, 1 red")
        fmt.Println(isCubeGameValid(gameId, red, green, blue))


        //)
    }
}

func isCubeGameValid(gameId int, red []int, green []int, blue []int) (int, bool) { 
    
    redNotValid, greenNotValid, blueNotValid := false, false, false
    for _, r := range(red) {
        if r > TOTAL_RED {
            redNotValid = true
        }
        if redNotValid {
            break
        }
    }
    for _, g := range(green) {
        if g > TOTAL_GREEN {
            greenNotValid = true
        }
        if greenNotValid {
            break
        }
    }
    for _, b := range(blue) {
        if b > TOTAL_BLUE {
            blueNotValid = true
        }
        if blueNotValid {
            break
        }
    }
    if !redNotValid && !greenNotValid && !blueNotValid {
        return gameId, true
    }

    return 0, false
}

func parseInputString(txt string) (int, []int, []int, []int){
    red:= []int{}
    green:= []int{}
    blue := []int{}
    _1 := strings.Split(txt, ":")
    gameId_s := strings.Split(_1[0], " ")
    gameId, _ := strconv.Atoi(gameId_s[1])
    //Get Red blue green
    draws := strings.Split(_1[1], ";")
    for _, draw := range draws {
        cubes := strings.Split(draw, ",")

        for _, cube := range cubes {

            cube_split := strings.Split(cube, " ")
            c := []string{}
            for i := 0; i < len(cube_split); i++ {
                if len(cube_split[i]) > 0 {
                    c = append(c, cube_split[i])
                }
            }
            num, err := strconv.Atoi(c[0])
            if  err == nil {
                if strings.Contains(c[1], "red") {
                    red = append(red, num)
                }
                if strings.Contains(c[1],"green")  {
                    green = append(green, num)
                }
                if  strings.Contains(c[1],"blue"){
                    blue = append(blue, num)
                }
            }
        }
    }
   // fmt.Printf("gameId: %v, red: %v, green: %v, blue: %v\n", gameId, red, green, blue)
    return gameId, red, green, blue
}
