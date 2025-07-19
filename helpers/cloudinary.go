package helpers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"pelaporan_keuangan/config" // Menggunakan path config Anda
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(file *multipart.FileHeader, folder string) (*uploader.UploadResult, error) {
	cfg := config.LoadBucketConfig()

	// 2. PERBAIKAN PENTING: Validasi bahwa kredensial tidak kosong
	if cfg.CLOUDINARY_CLOUD_NAME == "" || cfg.CLOUDINARY_API_KEY == "" || cfg.CLOUDINARY_API_SECRET == "" {
		log.Println("Kesalahan Kritis: Kredensial Cloudinary tidak lengkap. Pastikan CLOUDINARY_CLOUD_NAME, CLOUDINARY_API_KEY, dan CLOUDINARY_API_SECRET telah diatur di file .env dan dimuat dengan benar.")
		return nil, errors.New("konfigurasi Cloudinary tidak lengkap di server")
	}
	// 3. Buat instance Cloudinary dari konfigurasi yang dimuat
	cld, err := cloudinary.NewFromParams(
		cfg.CLOUDINARY_CLOUD_NAME,
		cfg.CLOUDINARY_API_KEY,
		cfg.CLOUDINARY_API_SECRET,
	)
	if err != nil {
		log.Printf("Gagal menginisialisasi Cloudinary: %v", err)
		return nil, fmt.Errorf("gagal menginisialisasi cloudinary: %w", err)
	}

	ctx := context.Background()
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("gagal membuka file: %w", err)
	}
	defer src.Close()

	// Menggunakan logika resourceType Anda yang lebih baik
	var resourceType string
	if strings.HasPrefix(file.Header.Get("Content-Type"), "image") {
		resourceType = "image"
	} else if strings.HasPrefix(file.Header.Get("Content-Type"), "video") {
		resourceType = "video"
	} else {
		resourceType = "raw"
	}

	uploadParams := uploader.UploadParams{
		PublicID:     file.Filename,
		Folder:       folder,
		ResourceType: resourceType,
	}

	uploadResult, err := cld.Upload.Upload(ctx, src, uploadParams)
	if err != nil {
		log.Printf("Error saat mengunggah ke Cloudinary: %v", err)
		return nil, fmt.Errorf("gagal mengunggah file: %w", err)
	}

	log.Printf("File berhasil diunggah ke Cloudinary. URL: %s", uploadResult.SecureURL)
	return uploadResult, nil
}
