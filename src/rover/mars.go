package main


type MarseDimensions struct {
	 MaxX int;
	 MaxY int ;
}

func ( mars MarseDimensions) IsValid(x , y int) bool {
	if mars.MaxX >= x && mars.MaxY >= y {
		return true ;
	}
	return false ;
}

