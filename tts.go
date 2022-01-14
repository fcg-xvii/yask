package yask

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TTSConfig is config for text to speeh method
type TTSConfig struct {
	Text       string
	SSML       string
	Lang       string
	Voice      string
	Emotion    string
	Speed      float32
	Format     string
	Rate       int
	YaFolderID string
	YaAPIKey   string
}

// ssml content checker
func (s *TTSConfig) isSSML() bool {
	return len(s.SSML) > 0
}

// returns TTSConfig with default parameters
func defaultTTSConfig(yaFolderID, yaAPIKey string) *TTSConfig {
	return &TTSConfig{
		Lang:       LangRU,
		Voice:      VoiceOksana,
		Emotion:    EmotionNeutral,
		Speed:      SpeedStandard,
		Format:     FormatLPCM,
		Rate:       Rate8k,
		YaFolderID: yaFolderID,
		YaAPIKey:   yaAPIKey,
	}
}

// TTSDefaultConfigText returns config with default parameters for raw text recognition and use in TextToSpeech method
func TTSDefaultConfigText(yaFolderID, yaAPIKey, text string) *TTSConfig {
	conf := defaultTTSConfig(yaFolderID, yaAPIKey)
	conf.Text = text
	return conf
}

// TTSDefaultConfigSSML returns config with default parameters for raw text recognition and use in TextToSpeech method
// more details of SSML language in https://cloud.yandex.ru/docs/speechkit/tts/ssml
func TTSDefaultConfigSSML(yaFolderID, yaAPIKey, SSML string) *TTSConfig {
	conf := defaultTTSConfig(yaFolderID, yaAPIKey)
	conf.SSML = SSML
	return conf
}

// TextToSpeech returns PCM or OGG sound stream using the service Yandex Speech Kit.
// Result PCM stream can be converted to Wav stream using EncodePCMToWav
func TextToSpeech(config *TTSConfig) (io.ReadCloser, error) {
	httpForm := url.Values{
		"lang":            []string{config.Lang},
		"voice":           []string{config.Voice},
		"emotion":         []string{config.Emotion},
		"speed":           []string{strconv.FormatFloat(float64(config.Speed), 'f', 1, 32)},
		"format":          []string{config.Format},
		"sampleRateHertz": []string{strconv.FormatInt(int64(config.Rate), 10)},
		"folderId":        []string{config.YaFolderID},
	}
	// emotion deprecated - only for omazh and jane
	if config.Voice != "omazh" && config.Voice != "jane" {
		delete(httpForm, "emotion")
	}
	// ssml
	if config.isSSML() {
		httpForm.Set("ssml", config.SSML)
	} else {
		httpForm.Set("text", config.Text)
	}

	request, err := http.NewRequest("POST", YaTTSUrl, strings.NewReader(httpForm.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", fmt.Sprintf("Api-Key %v", config.YaAPIKey))

	client := new(http.Client)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		err = unmarshallYaError(response.Body)
		response.Body.Close()
		return nil, err
	}

	return response.Body, nil
}
