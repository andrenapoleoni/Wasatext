package utils

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

func SaveAndCrop(filename string, w int, h int) error {
	// Apri il file immagine
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decodifica l'immagine in un oggetto image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	// Ridimensiona l'immagine usando imaging
	resizedImg := imaging.Resize(img, w, h, imaging.Lanczos)

	// Crea un nuovo file per salvare l'immagine ridimensionata
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Codifica e salva l'immagine ridimensionata
	if err := jpeg.Encode(out, resizedImg, nil); err != nil {
		return err
	}

	return nil
}
