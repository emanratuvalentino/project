package helper

import (
	"code-be-docudigital/config"

	"code-be-docudigital/storage"
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	uuid "github.com/satori/go.uuid"
)

/*-------------------------------------------------
	HELPER FUNCTION
	This is list of global function that used
	commonly. Refactor your function here only if
	needed and gave it a comment about its function!
--------------------------------------------------*/

// GetRSAPublicKey for read public key in config
func GetRSAPublicKey() (*rsa.PublicKey, error) {
	configuration := config.GetConfig()
	keyData, err := ioutil.ReadFile(configuration.PUBLIC_KEY)

	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}

func FailOnError(err error, msgerr string, msgsuc string) {
	if err != nil {
		log.Fatalf("%s: %s", msgerr, err)

	} else {
		fmt.Printf("%s\n", msgsuc)
	}

}

func CustomError(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}

				break
			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
}
func FindIn(what interface{}, where []interface{}) (idx int) {
	for i, v := range where {
		if v == what {
			return i
		}
	}
	return -1
}

func GenerateUUIDV5(namespace string) uuid.UUID {
	id, _ := uuid.NewV4()
	u1 := uuid.NewV5(id, namespace)

	return u1
}

func GetObject(objectName string) string {

	minioClient := storage.InitMinio()
	config := config.GetConfig()

	err := minioClient.FGetObject(context.Background(), config.MINIO_BUCKET_NAME, objectName, objectName, minio.GetObjectOptions{})
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		return errResponse.Code
	}

	return ""
}

func GetCurl(url string) map[string]interface{} {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var Result map[string]interface{}
	json.Unmarshal([]byte(data), &Result)
	return Result
}

func RemoveObject(objectName string) string {

	minioClient := storage.InitMinio()
	config := config.GetConfig()

	err := minioClient.RemoveObject(context.Background(), config.MINIO_BUCKET_NAME, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		return errResponse.Code
	}

	return ""
}

func PutObject(filename string, path string, file *multipart.FileHeader) string {

	config := config.GetConfig()

	ext := filepath.Ext(file.Filename)

	fTemp, err := ioutil.TempFile("", "tmp_files")
	if err != nil {
		panic(err)
	}

	src, _ := file.Open()

	defer src.Close()

	io.Copy(fTemp, src)

	minioClient := storage.InitMinio()

	n, err := minioClient.FPutObject(context.Background(), config.MINIO_BUCKET_NAME, path+filename+ext, fTemp.Name(), minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(n.Bucket)

	return path + filename + ext
}


