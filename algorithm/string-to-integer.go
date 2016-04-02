package algorithm
import (
	"math"
)
func MyAtoi(str string) int {
 	if(len(str) == 0){
 		return 0
 	}
 	symbols := []rune{}
 	digitals := []rune{}
 	numSys := []string{}
 	for index, char := range str {
 		if(string(char) == " "){
 			if(len(symbols) != 0 || len(digitals) != 0 || len(numSys) != 0) {
 				break
 			} 
 			continue
 		} else if(string(char) == "+" || string(char) == "-"){
 			if(len(symbols) != 0 || len(digitals) != 0 || len(numSys) != 0) {
 				break
 			} 
 			symbols = append(symbols, char)
 			
 		} else if string(char) == "x" {
 			if(index != 0 && str[index - 1:index + 1] == "0x"){
 				if(len(symbols) != 0 || len(digitals) != 0 || len(numSys) != 0) {
 					break
 				} 
 				numSys = append(numSys, "0x")
 			} else {
 				break
 			}
 		}else if (int(char) - int("0"[0])) >= 0 && (int(char) - int("0"[0])) <= 9{
 			digitals = append(digitals, char)
 		} else {
 			break
 		}
 	}
 	if(len(numSys) > 1 || len(symbols) > 1){
 		return 0
 	}
 	var max int = 0
 	size := len(digitals)

 	base := 10
 	if(len(numSys) != 0){
 		switch numSys[0]{
	 	case "0x":
	 		base = 16
	 	case "b":
	 		base = 2
	 	case "0":
	 		base = 8
	 	default:
	 		base = 10
	 	}
 	}
 	
 	for index, char := range digitals {
 		max += (int(char) - int("0"[0])) * int(math.Pow(float64(base), float64(size - index - 1)))
 	}
 	
 	if (len(symbols) == 0 || string(symbols[0:1]) == "+"){
 		if(max < 0){
 			max = 2147483647
 		}else{
 			max = max
 		}
 	} else if (len(symbols) != 0 && string(symbols[0:1]) == "-"){
 		if(max < 0){
 			max = -2147483648
 		}else{
 			max = -max
 		}
 	} else {
 		max = 0;
 	}
 	if max <= -2147483648{
		max = -2147483648
	} else if (max >= 2147483648){
		max = 2147483647
	}
	return max
}