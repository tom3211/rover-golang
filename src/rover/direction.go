package main

const (
	NORTH = iota
	EAST = iota
	SOUTH = iota
	WEST = iota
)
func ConvertDirectionStrToInt(dirStr string) int {
	var dirInt int
	switch (dirStr) {
		case "N" :
			 dirInt = NORTH
			 case "S" :
			dirInt= SOUTH 
			 case "W" :
			dirInt= WEST 
			 case "E" :
			dirInt= EAST 
	}
	return dirInt
}
func ConvertDirectionIntToStr(dir int) string {
	var dirStr string
	switch (dir) {
		case NORTH :
			 dirStr = "N"
	    case SOUTH :
			 dirStr = "S"
		case WEST :
			 dirStr = "W" 
		case EAST :
			 dirStr = "E"
	}
	return dirStr
}
