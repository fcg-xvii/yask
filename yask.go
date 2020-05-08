package yask

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

// EncodePCMToWav encode input stream of pcm audio format to wav and write to out stream
func EncodePCMToWav(in io.Reader, out io.WriteSeeker, sampleRate, bitDepth, numChans int) error {
	encoder := wav.NewEncoder(out, sampleRate, bitDepth, numChans, 1)

	audioBuf := &audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: numChans,
			SampleRate:  sampleRate,
		},
	}

	for {
		var sample int16
		if err := binary.Read(in, binary.LittleEndian, &sample); err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		audioBuf.Data = append(audioBuf.Data, int(sample))
	}

	if err := encoder.Write(audioBuf); err != nil {
		return err
	}

	return encoder.Close()
}

func unmarshallYaError(r io.Reader) (err error) {
	var data []byte
	if data, err = ioutil.ReadAll(r); err != nil {
		return
	}
	mErr := make(map[string]interface{})
	if err = json.Unmarshal(data, &mErr); err == nil {
		err = fmt.Errorf("Yandex request error: %v", mErr["error_message"])
	}
	return
}

// Voises returns slice of available vioces
// lang: empty (all alngs) ru-RU, en-EN, tr-TR
// sex: 0 - all, 1 - male, 2 - female
// premium: 0 - all, 1 - standard only, 2 - premium only
func Voices(lang string, sex, premium int) (res []Voice) {
	for _, voice := range voices {
		if len(lang) > 0 && voice.Lang != lang {
			continue
		}
		if sex != 0 && (sex == 1 && !voice.Male || sex == 2 && voice.Male) {
			continue
		}
		if premium != 0 && (voice.Premium && premium != 2 || !voice.Premium && premium != 1) {
			continue
		}
		res = append(res, voice)
	}
	return
}
