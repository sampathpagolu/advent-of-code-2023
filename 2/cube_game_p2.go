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
            red, green, blue :=   parseInputString(scanner.Text())
            sum += powerOfFewestNumberOfCubes(red, green, blue)
        }
        fmt.Println(sum)
    } else {
        red, green, blue := parseInputString("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
        fmt.Println(powerOfFewestNumberOfCubes(red, green, blue))
    }
}
func powerOfFewestNumberOfCubes(red int, green int, blue int) int{
    
    return red * blue * green
}

func parseInputString(txt string) (int, int, int){
    biggestRed, biggestGreen, biggestBlue := 0,0,0
    // split by ":" to get the game Id and get the cubes
    _1 := strings.Split(txt, ":")
    // split each draw
    draws := strings.Split(_1[1], ";")
    for _, draw := range draws {
        // split each cube color
        cubes := strings.Split(draw, ",")

        for _, cube := range cubes {
            // split the number and letter
            cube_split := strings.Split(cube, " ")
            c := []string{}
            for i := 0; i < len(cube_split); i++ {
                // remove empty strings
                if len(cube_split[i]) > 0 {
                    c = append(c, cube_split[i])
                }
            }
            num, err := strconv.Atoi(c[0])
            if  err == nil {
                if strings.Contains(c[1], "red") {
                    if num > biggestRed {
                        biggestRed = num
                    }
                }
                if strings.Contains(c[1],"green")  {
                    if num > biggestGreen {
                        biggestGreen = num
                    }
                }
                if  strings.Contains(c[1],"blue"){
                    if num > biggestBlue {
                        biggestBlue = num
                    }
                }
            }
        }
    }
   // fmt.Printf("gameId: %v, red: %v, green: %v, blue: %v\n", gameId, red, green, blue)
    return biggestRed, biggestGreen, biggestBlue 
}

