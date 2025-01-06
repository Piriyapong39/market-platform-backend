package productService

// Handle single file upload
// func SavePictureFiles(mainPic *multipart.FileHeader) (string, error) {
// 	// Generate unique filename
// 	filename := time.Now().Format("20060102150405") + filepath.Ext(mainPic.Filename)
// 	uploadPath := "uploads/products/" + filename

// 	// Create directory if not exists
// 	if err := os.MkdirAll("uploads/products", 0755); err != nil {
// 		return "", err
// 	}

// 	// Save file
// 	if err := mainPic.Save(uploadPath); err != nil {
// 		return "", err
// 	}

// 	return filename, nil
// }
