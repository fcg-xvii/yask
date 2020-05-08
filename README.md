<p align="center">
  <span>English</span> |
  <a href="README.ru.md">Русский</a>
</p>

# yask
> Tools for work with the synthesis and speech recognition service <b>Yandex Speech Kit</b> (more about in <a href="https://cloud.yandex.ru/docs/speechkit/" target="_blank">https://cloud.yandex.ru/docs/speechkit/</a>) for <b>golang</b> programming language. Used to synthesize speech from text and recognize text from a sound stream.

Before start to use, you must register at <a href="https://cloud.yandex.ru/" target="_blank">https://cloud.yandex.ru/</a> to get the API key and directory identifier (more about <a href="https://cloud.yandex.ru/docs" target="_blank">https://cloud.yandex.ru/docs</a>).

### Audio stream formats
<ul>
    <li><b>OGG</b> <a href="https://ru.wikipedia.org/wiki/Ogg" target="_blank">https://en.wikipedia.org/wiki/Ogg</a></li>
    <li><b>PCM</b> <a href="https://en.wikipedia.org/wiki/Pulse-code_modulation" target="_blank">https://en.wikipedia.org/wiki/Pulse-code_modulation</a> (when recognizing text in the lpcm format parameter, a wav format stream can be used</li>
</ul>

### Speech synthesis from text
> As a result of the example, get a file in wav format, ready for playback in any player program. The default bitrate is 8000.
```golang
import (
	"log"
	"os"

	"github.com/fcg-xvii/go-tools/speech/yask"
)

func main() {
	yaFolderID := "b1g..."    // yandex folder id
	yaAPIKey := "AQVNy..."    // yandex api yandex
	text := "Hi It's test of speech synthesis" // text for synthesis

	// init config for synthesis (по умоланию установлен формат lpcm)
	config := yask.TTSDefaultConfigText(yaFolderID, yaAPIKey, text)

    // By default language in config russian. For english must setup 
    // english language and voice
    config.Lang = yask.LangEN
    config.Voice = yask.VoiceNick


	// speech synthesis
	r, err := yask.TextToSpeech(config)
	if err != nil {
		log.Println(err)
		return
	}

    // open file for save result
	f, err := os.OpenFile("tts.wav", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0655)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

    // lpcm encoding to wav format
	if err := yask.EncodePCMToWav(r, f, config.Rate, 16, 1); err != nil {
		log.Println(err)
		return
	}
}
```

### Speech recognition to text
> Example of recognition of short audio. The example uses a wav file that can be used with a configuration format value of <b>lpcm</b>

```golang
package main

import (
	"log"
	"os"

	"github.com/fcg-xvii/go-tools/speech/yask"
)

func main() {
	yaFolderID := "b1g4..." // yandex folder id
	yaAPIKey := "AQVNyr..." // yandex api key
	dataFileName := "data.wav" // audio file in wav format for recodnition to text

    // open audio file
	f, err := os.Open(dataFileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

    // init config for recodnition
	config := yask.STTConfigDefault(yaFolderID, yaAPIKey, f)

    // setup english language
    config.Lang = yask.LangEN

    // recodnition speech to text
	text, err := yask.SpeechToTextShort(config)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(text)
}
```