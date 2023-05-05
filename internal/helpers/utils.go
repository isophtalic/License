package helpers

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"math"
	"math/big"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"
	"time"

	"git.cyradar.com/license-manager/backend/internal/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var timeLayout = "2006-01-02"

func TrimStruct(obj interface{}) interface{} {
	if reflect.ValueOf(obj).Kind() != reflect.Pointer {
		panic("In func TrimStruct: input must be a pointer")
	}

	if reflect.TypeOf(obj).Elem().Kind() != reflect.Struct {
		panic("In func TrimStruct: input must be a pointer of Struct")
	}

	types := reflect.TypeOf(obj).Elem()
	values := reflect.ValueOf(obj).Elem()
	for i := 0; i < types.NumField(); i++ {
		if types.Field(i).Type.Elem().Kind() == reflect.String && !values.Field(i).IsNil() {
			values.Field(i).Elem().Set(reflect.ValueOf(strings.TrimSpace(values.Field(i).Elem().String())))
		}
	}
	return obj
}

func ReadFile(file *multipart.FileHeader) string {
	f, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var content string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content += scanner.Text()
	}
	return content
}

/*
Convert string to time.Time
*/
func ParseDate(date string) time.Time {
	dt, err := time.Parse(timeLayout, date)
	if err != nil {
		panic(err)
	}
	return dt
}

func CreatePagination(perPage string, page string, sort string) *dto.PaginationDTO {
	pagination := &dto.PaginationDTO{}
	if perPage != "" {
		if l, err := strconv.Atoi(perPage); err == nil {
			pagination.PerPage = l
		}
	}

	if page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			pagination.Page = p
		}
	}

	if sort != "" {
		if sort[0] == '-' {
			sort = sort[1:] + " DESC"
		} else {
			sort = sort + " ASC"
		}
		pagination.Sort = sort
	}

	return pagination
}

func Paginate(model interface{}, pagination *dto.PaginationDTO, db *gorm.DB) func(*gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(model).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetPerPage())))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPerPage()).Order(pagination.GetSort())
	}
}

func GenerateSalt(saltSize int) string {
	var letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]rune, saltSize)
	for i := range b {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		b[i] = rune(letterRunes[randomIndex.Int64()])
	}
	return string(b)
}

func HashPassword(password string, salt string) (hashPassword ArrayBytes) {
	var passwordBytes = []byte(password)
	passwordBytes = append(passwordBytes, []byte(salt)...)

	hashPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return
}

func CompareAndHashPassword(hashedPassword string, password string, salt string) bool {
	passwordBytes := append([]byte(password), []byte(salt)...)
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), passwordBytes); err != nil {
		return false
	}

	return true
}

func StringArrayToBytes(stringData []string) ArrayBytes {
	// concatenate the strings into a single []byte slice
	var byteData []byte
	var buffer bytes.Buffer
	for _, str := range stringData {
		buffer.WriteString(str)
	}
	byteData = buffer.Bytes()

	return byteData
}

type ArrayBytes []byte

func (arr ArrayBytes) ToStringArray() []string {
	byteChunks := bytes.Split(arr, []byte(" "))

	// create a new []string slice and copy each chunk into it
	stringChunks := make([]string, len(byteChunks))
	for i, chunk := range byteChunks {
		stringChunks[i] = string(chunk)
	}
	return stringChunks
}
