package internal

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ConvertImageToBase64(imgPath string) string {
	imageFilePath := imgPath

	base64String, err := imageToBase64(imageFilePath)
	if err != nil {
		log.Fatal(err)
	}

	return base64String
}

func imageToBase64(filePath string) (string, error) {
	imageFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	imageData, err := io.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(imageData)
	return base64String, nil
}

func ExtractImageName(filePath string) string {
	// Find the last index of "/"
	lastIndex := strings.LastIndex(filePath, "/")
	// Find the index of ".jpg"
	dotIndex := strings.Index(filePath, ".")

	if lastIndex < 0 || dotIndex < 0 || lastIndex >= dotIndex {
		return ""
	}

	// Extract the image name
	imgName := filePath[lastIndex+1 : dotIndex]
	return imgName
}

func ExtractImageFromJSON(input string) string {
	// Find the index of "image:"
	imageIndex := strings.Index(input, "image:")

	if imageIndex == -1 {
		fmt.Println("Value 'image:' not found.")
		return ""
	}

	// Find the next space character after "image:"
	spaceIndex := strings.Index(input[imageIndex:], " ")

	if spaceIndex == -1 {
		fmt.Println("Space character after 'image:' not found.")
		return ""
	}

	// Extract the value between "image:" and the space character
	imageValue := input[imageIndex+6 : imageIndex+spaceIndex]

	return imageValue
}
