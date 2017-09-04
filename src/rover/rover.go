package main

type RoverLocation struct {
	XPosition int 
	YPosition int 
	Direction int
}

func (rl *RoverLocation) Move() {
	switch(rl.Direction) {
		case NORTH :
			rl.YPosition +=1
		case SOUTH :
			rl.YPosition -=1
		case WEST :
			rl.XPosition -=1
		case EAST :
			rl.XPosition +=1
	}
	//fmt.Println("RoverLocation Move %d %d", rl.XPosition, rl.YPosition) 
}

func (rl RoverLocation) NextMovePosition() (x, y int ){
	 localX := rl.XPosition 
	 localY := rl.YPosition 
	switch(rl.Direction) {
		case NORTH :
			localY +=1
		case SOUTH :
			localY -=1
		case WEST :
			localX -=1
		case EAST :
			localX +=1
		
	}
	return localX, localY 
}

func (rl *RoverLocation) TurnLeft() {
	rl.Direction = (rl.Direction + (4 -1)) %4 ;
	//fmt.Println("RoverLocation TurnLeft %d ", rl.Direction) 
}

func (rl *RoverLocation) TurnRight() {
	rl.Direction = (rl.Direction + (4 +1)) %4 ;
	//fmt.Println("RoverLocation TurnRight %d ", rl.Direction) 
}
