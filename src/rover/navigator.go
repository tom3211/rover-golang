package main


import "fmt"
import "os"
import "strings"
import "log"
import "bufio"
import "strconv"
import "flag"
var loggingEnabled = true 
type argError struct {
    arg  int
    prob string
}
func handleError(err error) {
	if err != nil {
		fmt.Println("Error in processing %s ", err) ;
		log.Fatal("Error ")
	}
}
func initalizeMars(line string) MarseDimensions  {
	dimensions := strings.Split(line, " ")
	marsX, err := strconv.Atoi(dimensions[0])
	handleError(err) 
	marsY, err := strconv.Atoi(dimensions[1])
	handleError(err) 
	mars := MarseDimensions{marsX,marsY}
	return mars 
}
func initalizeRoverLocation(line string) RoverLocation  {
	roverParams := strings.Split(line, " ")
	roverX, err := strconv.Atoi(roverParams[0])
	handleError(err) 
	roverY, err := strconv.Atoi(roverParams[1])
	handleError(err) 
	dirInt := ConvertDirectionStrToInt(roverParams[2])
	
	roverLocation := RoverLocation{roverX,roverY, dirInt}
	return roverLocation 
}
func logData(line string) {
	if loggingEnabled == true {
		fmt.Printf("%s\n" , line)
	}
}

func main() {
		loggingEnabled = *flag.Bool("loggingEnabled", false, "a bool")
		scanner := bufio.NewScanner(os.Stdin)	
		scanner.Scan() 
		line := scanner.Text() ;
		logData(line)
		mars := initalizeMars(line)
	for scanner.Scan() {
		line := scanner.Text()
		logData(line)
		if strings.EqualFold(line, "Q") {
			fmt.Println("Done process")
			break ;
		}
		roverLocation := initalizeRoverLocation(line)
	//	fmt.Printf("rover x %d y %d dir %d", roverLocation.XPosition,
	//						roverLocation.YPosition, roverLocation.Direction)
		scanner.Scan()
		line = scanner.Text()
		logData(line)
		for _, r := range line {
		   if r == 'L' {
		   	 roverLocation.TurnLeft()
		   	//fmt.Printf("main TurnLeft %d ", roverLocation.Direction) 
		   } else if r == 'R' {
		   	  roverLocation.TurnRight()
		   	  // fmt.Printf("main TurnRight %d ", roverLocation.Direction) 
		   } else if r == 'M' {
			   x, y := roverLocation.NextMovePosition() ;
			   valid := mars.IsValid(x,y) 
			   if valid == true {
				   roverLocation.Move() ;
				    //fmt.Println("main x y %d %d ", roverLocation.XPosition,roverLocation.YPosition) 
			   } else {
			   	log.Fatal(" Invalid move command %c " , r)
			   }
		   } else {
		   	  log.Fatal(" Invalid command %d " , r)
		   }
	 }
		fmt.Printf("%d %d %s\n", roverLocation.XPosition, 
			roverLocation.YPosition, ConvertDirectionIntToStr(roverLocation.Direction)) ;
	}
}
