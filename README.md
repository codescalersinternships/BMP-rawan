# BMP Width/Height Extractor 

This repository implements a BMP Width/Height Extractor that reads BMP image and extracts its width/height from the header

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/codescalersinternships/BMP-rawan.git
   ```
## APIs

- `ExtractWH(path string) (int, int, error)` : Responsible for parsing the BMP file an returning its width and height

## Usage

```go
    w,h,err:= bmp.ExtractWH("testdata/sample1.bmp")
	if err!=nil{
		panic(err)
	}
	println("Width:",w,"Height:",h)
```