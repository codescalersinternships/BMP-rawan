package pkg

import "os"
const (
	widthStart = 18
	heightStart = 22
	Size = 4
)

func readFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ExtractWH(path string) (int, int, error) {
	image, err := readFile(path)
	if err != nil {
		return -1, -1, err
	}
	
	widthBytes := image[widthStart:widthStart+Size]
	heightBytes := image[heightStart:heightStart+Size]

	//little endian
	width :=  int(widthBytes[3])<<24 | int(widthBytes[2])<<16 | int(widthBytes[1])<<8 | int(widthBytes[0])
	height := int(heightBytes[3])<<24 | int(heightBytes[2])<<16 | int(heightBytes[1])<<8 | int(heightBytes[0])
	
	return width, height, nil
}
