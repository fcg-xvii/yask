package yask

import (
	"os"
	"testing"

	"github.com/fcg-xvii/go-tools/text/config"
)

var (
	yaFolderID, yaAPIKey string
)

func init() {
	if f, err := os.Open("test_data/ya.config"); err == nil {
		config.SplitToVals(f, "::", &yaFolderID, &yaAPIKey)
		f.Close()
	}
}

func TestTextToSpeech(t *testing.T) {
	if len(yaFolderID) == 0 {
		t.Log("ya config 'test_data/ya.config' not parsed. format 'ya_folder_id::ya_api_key")
		return
	}

	// init request config
	config := TTSDefaultConfigText(yaFolderID, yaAPIKey, "Привет, это тест синтеза речи с помощью сервиса Яндекса")

	r, err := TextToSpeech(config)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.OpenFile("test_data/tts.wav", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0655)
	if err != nil {
		t.Fatal(err)
	}

	if err := EncodePCMToWav(r, f, config.Rate, 16, 1); err != nil {
		t.Fatal(err)
	}

	r.Close()
	f.Close()
}

func TestSpeechToTextShort(t *testing.T) {
	if len(yaFolderID) == 0 {
		t.Log("ya config 'test_data/ya.config' not parsed. format 'ya_folder_id::ya_api_key")
		return
	}

	f, err := os.Open("test_data/test_sound.wav")
	if err != nil {
		t.Fatal(err)
	}

	conf := STTConfigDefault(yaFolderID, yaAPIKey, f)

	text, err := SpeechToTextShort(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(text)
}

func TestVoiseFilter(t *testing.T) {
	// Get all voices
	items := Voices("", 0, 0)
	t.Log(len(voices), len(items), items)

	// Get only russian standard females
	items = Voices(LangRU, 2, 1)
	t.Log(len(voices), len(items), items)
}
