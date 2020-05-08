package yask

type Voice struct {
	NameEn  string `json:"name_en"`
	MameRu  string `json:"name_ru"`
	Voice   string `json:"voice"`
	Lang    string `json:"lang"`
	Male    bool   `json:"is_male"`
	Premium bool   `json:"is_premium"`
}

const (
	// YaSTTUrl is url for send speech to text requests
	YaSTTUrl = "https://stt.api.cloud.yandex.net/speech/v1/stt:recognize"

	// YaTTSUrl is url for send text to speech requests
	YaTTSUrl = "https://tts.api.cloud.yandex.net/speech/v1/tts:synthesize"

	// Formats of audio
	// FormatLPCM is PCM audio format (wav) without wav header (more details in https://en.wikipedia.org/wiki/Pulse-code_modulation)
	FormatLPCM = "lpcm"
	// FormatOgg is audio ogg format
	FormatOgg = "oggopus"

	// Sample rates

	// Rate8k is rate of 8kHz
	Rate8k int = 8000
	// Rate16k is rate of 16kHz
	Rate16k int = 16000
	// Rate48k is rate of 48kHz
	Rate48k int = 48000

	// Languages

	// LangRU is russian language
	LangRU = "ru-Ru"
	// LangEN is english language
	LangEN = "en-US"
	// LangTR is turkish language
	LangTR = "tr-TR"

	// Speed constants

	// SpeedStandard is standart speed of voice (1.0)
	SpeedStandard float32 = 1.0
	// SpeedMostFastest is maximum speed voice (3.0)
	SpeedMostFastest float32 = 3.0
	// SpeedSlowest is minimum speed of voice (0.1)
	SpeedSlowest float32 = 0.1

	// Voice speechs

	// VoiceOksana is Oksana voice (russian, female, standard)
	VoiceOksana = "oksana"
	// VoiceJane is Jane voice (russian, female, standard)
	VoiceJane = "jane"
	// VoiceOmazh is Omazh voice (russian, female, standard)
	VoiceOmazh = "omazh"
	// VoiceZahar is Zahar voice (russian, male, standard)
	VoiceZahar = "zahar"
	// VoiceErmil is Ermil voice (russian, male, standard)
	VoiceErmil = "ermil"
	// VoiceSilaerkan is Silaerkan voice (turkish, female, standard)
	VoiceSilaerkan = "silaerkan"
	// VoiceErkanyavas is Erkanyavas voice (turkish, male, standard)
	VoiceErkanyavas = "erkanyavas"
	// VoiceAlyss is Alyss voice (english, female, standard)
	VoiceAlyss = "alyss"
	// VoiceNick is Nick voice (engish, male, standard)
	VoiceNick = "nick"
	// VoiceAlena is Alena voice (russian, female, premium)
	VoiceAlena = "alena"
	// VoiceFilipp is Filipp voice (russian, male, premium)
	VoiceFilipp = "filipp"

	// Voice emotions
	// EmotionGood is good voice emotion
	EmotionGood = "good"
	// EmotionEvil is evil voice emotion
	EmotionEvil = "evil"
	// EmotionNeutral is neutral voice emotion
	EmotionNeutral = "neutral"

	// Models for speech recodnition

	// TopicGeneral is current version of voice model (available in all languages)
	TopicGeneral = "general"
	// TopicRC is experimental version of voice model (russian language)
	TopicGeneralRC = "general:rc"
	// TopicDeprecated is deprecated version of voice model (russian language)
	TopicGeneralDeprecated = "general:deprecated"
	// TopicMaps is model for addresses anc company names
	TopicMaps = "maps"

	// This constants for use in voice selection filter

	// SexAll is male and female
	SexAll = 0
	// SexMale is male
	SexMale = 1
	// SexFemale is female
	SexFemale = 2
)

var (
	// voices is list of voice params
	voices = []Voice{
		Voice{"Oksana", "Оксана", VoiceOksana, LangRU, false, false},
		Voice{"Jane", "Джейн", VoiceJane, LangRU, false, false},
		Voice{"Omazh", "Омаж", VoiceOmazh, LangRU, false, false},
		Voice{"Zahar", "Захар", VoiceZahar, LangRU, true, false},
		Voice{"Ermil", "Эрмил", VoiceErmil, LangTR, true, false},
		Voice{"Sila Erkan", "Сыла Эркан", VoiceSilaerkan, LangTR, false, false},
		Voice{"Alyss", "Элис", VoiceAlyss, LangTR, false, false},
		Voice{"Nick", "Ник", VoiceNick, LangTR, true, false},
		Voice{"Alena", "Алёна", VoiceNick, LangRU, false, true},
		Voice{"Filipp", "Филипп", VoiceNick, LangRU, true, true},
	}
)
