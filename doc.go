// Copyright 2020 fcg-xvii. All rights reserved.
// Use of this source code is governed by a MIT license
// license that can be found in the LICENSE file.

// Tools for work with the synthesis and speech recognition service Yandex Speech Kit (more about in https://cloud.yandex.ru/docs/speechkit/). 
// Can be used to synthesize speech from text and recognize text from a sound stream.
//
// Before start to use, you must register at https://cloud.yandex.ru/ to get the API key and directory identifier (more about https://cloud.yandex.ru/docs).
//
// Audio stream formats
//    OGG https://en.wikipedia.org/wiki/Ogg
//    PCM https://en.wikipedia.org/wiki/Pulse-code_modulation (when recognizing text in the lpcm format parameter, a wav format stream can be used
//
//
package yask