package product

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	// import service

	userservices "github.com/piriyapong39/market-platform/services/user-services"
)

func createProduct(c *fiber.Ctx) error {

	// declare main variable
	newPictureName := uuid.New()
	var picFilesArr []string

	// verify identity
	token := c.Get("Authorization")
	userData, err := userservices.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// request from form-data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// main picture
	picMain := form.File["main_image"]
	if len(picMain) != 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "only 1 main picture",
		})
	}
	// fmt.Println(picMain)
	mainPicName := strings.Split(picMain[0].Filename, ".")
	fileExtension := mainPicName[len(mainPicName)-1]
	if fileExtension != "jpg" && fileExtension != "png" && fileExtension != "jpeg" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid main picture format. Only jpg, png, and jpeg are allowed.",
		})
	}
	mainPicNameFormatted := fmt.Sprintf("main.%s", fileExtension)
	picFilesArr = append(picFilesArr, mainPicNameFormatted)

	// picture files
	picFiles := form.File["images"]
	if len(picFiles) == 0 || len(picFiles) > 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid number of files",
		})
	}
	for _, file := range picFiles {
		re := regexp.MustCompile(`\s+`)
		joinName := re.ReplaceAllString(file.Filename, "")
		splitName := strings.Split(joinName, ".")
		extFile := splitName[len(splitName)-1]
		if extFile != "png" && extFile != "jpeg" && extFile != "jpg" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid picture format. Only jpg, png, and jpeg are allowed.",
			})
		}
		file.Filename = fmt.Sprintf("%s.%s", newPictureName.String(), extFile)
		picFilesArr = append(picFilesArr, file.Filename)
	}

	// product name
	productName := form.Value["name"]

	// product description
	productDescription := form.Value["description"]

	// product stock
	productStock := form.Value["stock"]
	intStock, err := strconv.Atoi(productStock[0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid stock value",
		})
	}
	if intStock < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Stock value must more than 1",
		})
	}

	// product price
	productPrice := form.Value["price"]
	floatPrice, err := strconv.ParseFloat(productPrice[0], 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid price value",
		})
	}
	if floatPrice < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product price must more than 0",
		})
	}

	// product category
	productCategoryID := form.Value["category_id"]
	intCategoryID, err := strconv.Atoi(productCategoryID[0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	// Set product request data using struct initialization
	productRequest := &Product{
		Name:        productName[0],
		Description: productDescription[0],
		Stock:       intStock,
		Price:       floatPrice,
		CategoryID:  intCategoryID,
		PicPath:     picFilesArr,
		User_id:     userData.Id,
	}

	// Create product
	picturesNameArr, err := _createProduct(*productRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// create folder and upload files

	mainnnPic, err := c.FormFile("main_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(mainnnPic)

	// result, err := productService.SavePictureFiles([]*multipart.FileHeader{picMain})
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": err,
	// 	})
	// }
	// fmt.Println(result)

	// Create picture
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": picturesNameArr,
	})
}
