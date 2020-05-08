package yask

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// STTConfigDefault returns STTConfig with default parameters
func STTConfigDefault(yaFolderID, yaAPIKey string, data io.Reader) *STTConfig {
	return &STTConfig{
		Lang:            "ru-RU",
		Topic:           "general",
		ProfanityFilter: false,
		Format:          FormatLPCM,
		Rate:            Rate8k,
		YaFolderID:      yaFolderID,
		YaAPIKey:        yaAPIKey,
		Data:            data,
	}
}

// STTConfig is config for speech to text methods
type STTConfig struct {
	Lang            string
	Topic           string
	ProfanityFilter bool
	Format          string
	Rate            int
	YaFolderID      string
	YaAPIKey        string
	Data            io.Reader
}

// uri returns url with get parameters for http request
func (s *STTConfig) uri() string {
	vars := url.Values{
		"lang":            []string{s.Lang},
		"topic":           []string{s.Topic},
		"profanityFilter": []string{strconv.FormatBool(s.ProfanityFilter)},
		"format":          []string{s.Format},
		"sampleRateHertz": []string{strconv.FormatInt(int64(s.Rate), 10)},
		"folderId":        []string{s.YaFolderID},
	}

	url := fmt.Sprintf("%v?%v", YaSTTUrl, vars.Encode())
	return url
}

// SpeechToTextShort returns text from a PCM or OGG sound stream using the service Yandex Speech Kit
func SpeechToTextShort(conf *STTConfig) (string, error) {
	req, err := http.NewRequest(
		"POST",
		conf.uri(),
		conf.Data,
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Transfer-encoding", "chunked")
	req.Header.Set("Authorization", fmt.Sprintf("Api-Key %v", conf.YaAPIKey))

	cl := new(http.Client)

	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", unmarshallYaError(resp.Body)
	}

	rSource, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	m := make(map[string]interface{})
	if err = json.Unmarshal(rSource, &m); err != nil {
		return "", err
	}

	result := fmt.Sprint(m["result"])
	return result, nil
}
