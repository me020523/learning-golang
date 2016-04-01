package algorithm

import "fmt"
func getLayer(offset int, numRows int) (layer int) {
	if(offset < 0){
		return 1
	}
	if offset < numRows {
		layer = offset
	} else {
		layer = (numRows - 1) - (offset - numRows + 1)
	}
	return
}
func Convert(s string, numRows int) string {
    size := len(s)
    if(size == 0 || numRows == 1){
    	return s
    }
    buf := make([]rune, size)
    remain := size % ((numRows - 1) << 1)
    cols := size / ((numRows - 1) << 1)
	last_char_layer := getLayer(remain - 1, numRows)

	midVal := ((numRows - 1) << 1)
	fmt.Println(last_char_layer)
    for index, char := range s {
    	var currentCol int = index / midVal
    	var offset int = index % midVal
    	var layer int = getLayer(offset, numRows)
    	var position int = 0
    	if layer == 0 {
    		position = currentCol
    	} else if layer == (numRows - 1) {
    		position = (size - 1) - (cols - currentCol - 1)
    		if(remain >= numRows){
    			position -= 1;
    		}
    	} else {
    		position = (layer << 1 - 1) * cols
    		
    		if(currentCol >= cols && cols != 0){
    			position += cols << 1;
    		} else {
    			position += currentCol << 1
    		}
    		if(remain <= numRows && remain > 0){
    			if(layer > last_char_layer){
    				position += (last_char_layer + 1)
    			} else {
    				position += layer
    			}
    		} else if (remain > numRows) {
    			if(layer > last_char_layer){
    				position += ((layer - last_char_layer) << 1 + last_char_layer)
    			} else {
    				position += layer
    			}
    		}
    		if(offset >= numRows){
    			position += 1
    		}
    	}
    	fmt.Println(index, layer, " ", position," ", string(char), offset, currentCol)
    	buf[position] = char
    }
    return string(buf)
}
