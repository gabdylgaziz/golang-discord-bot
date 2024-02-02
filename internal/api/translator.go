package api

import (
	"discord/internal/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Detect(text string) (result string) {
	url := "https://google-translate1.p.rapidapi.com/language/translate/v2/detect"

	payload := strings.NewReader(fmt.Sprintf("q=%s", text))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("TranslatorAPI"))
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var lang entity.SourceLanguage

	err := json.Unmarshal(body, &lang)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	result = lang.Data.Detections[0][0].Language
	return
}

func Translate(text string, target string, source string) (result string) {
	url := "https://google-translate1.p.rapidapi.com/language/translate/v2"
	payload := strings.NewReader(fmt.Sprintf("q=%s&target=%s&source=%s", text, target, source))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("TranslatorAPI"))
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var lang entity.TranslatedText

	err := json.Unmarshal(body, &lang)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	result = lang.Data.Translations[0].TranslatedText
	return
}
