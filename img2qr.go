package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {

	// Define command-line flags for width and height
	width := flag.Int("width", 256, "Width of the QR code")
	height := flag.Int("height", 256, "Height of the QR code")
	colorFlag := flag.String("color", "black", "Color of the QR code (options: black, green, red, blue)")

	flag.Parse()

	// Define the directory to search for images
	dir := "."

	// Supported image file extensions
	imageExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	// Walk through the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is an image
		if !info.IsDir() && imageExtensions[filepath.Ext(path)] {
			log.Printf("Found image: %s", path)
			imageBytes, err := os.ReadFile(path)
			if err != nil {
				log.Fatalf("Failed to read image file: %v", err)
			}

			// Encode the image bytes to Base64
			encodedString := base64.StdEncoding.EncodeToString(imageBytes)

			// Print or use the Base64 encoded string
			//fmt.Println(encodedString)
			generateQRCode(path, encodedString, *width, *height, *colorFlag)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through directory: %v", err)
	}

	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}

func generateQRCode(filePath string, encodedString string, width int, height int, colorFlag string) {

	// Create a QR code writer
	qrWriter := qrcode.NewQRCodeWriter()

	// Encode the data into a QR code matrix
	qrCode, err := qrWriter.Encode(encodedString, gozxing.BarcodeFormat_QR_CODE, width, height, nil)
	if err != nil {
		log.Fatalf("Failed to encode QR code: %v", err)
	}

	// Create a blank image with the same size as the QR code matrix
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Set the colors for the QR code
	var qrColor color.RGBA
	switch colorFlag {
	case "green":
		qrColor = color.RGBA{0, 255, 0, 255} // Green color
	case "red":
		qrColor = color.RGBA{255, 0, 0, 255} // Red color
	case "blue":
		qrColor = color.RGBA{0, 0, 255, 255} // Blue color
	default:
		qrColor = color.RGBA{0, 0, 0, 255} // Black color
	}
	white := color.RGBA{255, 255, 255, 255}

	// Transfer the QR code matrix to the image
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if qrCode.Get(x, y) {
				img.Set(x, y, qrColor)
			} else {
				img.Set(x, y, white)
			}
		}
	}

	// Open a file to save the QR code image
	outputFilePath := filePath + "_QR_Code.png"
	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Failed to create QR code file: %v", err)
	}
	defer file.Close()

	// Encode the image as a PNG and save it to the file
	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("Failed to encode QR code image: %v", err)
	}

	log.Printf("QR code generated and saved to %s", outputFilePath)
}
