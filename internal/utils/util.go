package utils

import (
	cryptorand "crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var LicenseText = "Duranz API"

type JsonWrappedContent struct {
	License    string      `json:"license"`
	StatusCode int         `json:"statusCode"`
	Content    interface{} `json:"content"`
}

func CleanText(text string, lowerCase bool) string {
	sanitizedText := text
	if lowerCase {
		sanitizedText = strings.ToLower(sanitizedText)
	}

	return strings.TrimSpace(sanitizedText)
}

// PrintJSON prints a json object nicely
func PrintJSON(j interface{}) error {
	var out []byte
	var err error

	out, err = json.MarshalIndent(j, "", "    ")

	if err == nil {
		fmt.Println(string(out))
	}

	return err
}

// JSONMessageObj returns an encoded JSON of the object provided.
func JSONMessageObj(obj interface{}) []byte {

	result, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	return result
}

// WebResponseJSONObject is a wrapper function that returns an already prepared JSON object as web response.
func WebResponseJSONObject(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(obj.([]byte))
}

// JSONMessageObj returns an encoded JSON of the object provided.
func JSONMessageWrappedObj(code int, obj interface{}) []byte {
	jsonString := JsonWrappedContent{
		License:    LicenseText,
		StatusCode: code,
		Content:    obj,
	}

	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	return result
}

// makeRandomString : Make the randon string of certain length
func makeRandomString(strlen int, stringType string) string {

	//rand.Seed(time.Now().UTC().UnixNano())
	var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	switch strings.ToLower(stringType) {
	case "number":
		chars = "1234567890"
		break
	}

	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		max := *big.NewInt(int64(len(chars)))
		n, err := cryptorand.Int(cryptorand.Reader, &max)
		if err != nil {
			fmt.Println(err)
		}

		result[i] = chars[int(n.Int64())]
		//result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// RandString - Generates a randomized string of specified length and case
// length: what length of random strings to produce,  casing: 0 for lowercase, 1 for uppercase, 2 for mixed case
func RandString(length int, casing int) string {
	result := makeRandomString(length, "string")

	if casing == 0 {
		result = strings.ToLower(result)
	}
	if casing == 1 {
		result = strings.ToUpper(result)
	}
	time.Sleep(time.Nanosecond * 248)
	return result
}

// RandNumString - Generates a numeric string of specified length
func RandNumString(length int) string {
	result := makeRandomString(length, "number")
	return result
}

// IsNumeric : check if the string is numeric or not
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Round : Round the number
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// RequestAPIData - Calls a (API) URL and return the data from the request.
func RequestAPIData(method, url, postdata string, headers map[string]string) ([]byte, int, error) {

	req, err := http.NewRequest(method, url, strings.NewReader(postdata))
	if err != nil {
		return nil, 500, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	statusCode := resp.StatusCode
	if err != nil {
		return nil, statusCode, err
	}

	// read body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, statusCode, err
	}
	return body, statusCode, nil
}

// CheckNull : check null value
func CheckNull(rowData string) *int {
	if rowData != "" && rowData != "NULL" {
		p1total, _ := strconv.Atoi(rowData)
		return &p1total
	}
	return nil
}

// SecondsToMinutes :
func SecondsToMinutes(inSeconds int64) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d", minutes) + `:` + fmt.Sprintf("%02d", seconds)
	return str
}
