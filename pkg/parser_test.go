package pkg

import (
	"os"
	"testing"

	"golang.org/x/image/bmp"
)

func TestExtractWH(t *testing.T) {
	w, h, err := ExtractWH("../testdata/sample1.bmp")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	file, err := os.Open("../testdata/sample1.bmp")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer file.Close()

	img, err := bmp.Decode(file)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedWidth := img.Bounds().Dx()
	expectedHeight := img.Bounds().Dy()
	
	if w != expectedWidth || h != expectedHeight {
		t.Errorf("Expected width: %d, height: %d; got width: %d, height: %d", expectedWidth, expectedHeight, w, h)
	}

}
