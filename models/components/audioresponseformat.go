// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AudioResponseFormat - The format of the output, in one of these options: `json`, `text`, `srt`, `verbose_json`, or `vtt`.
type AudioResponseFormat string

const (
	AudioResponseFormatJSON        AudioResponseFormat = "json"
	AudioResponseFormatText        AudioResponseFormat = "text"
	AudioResponseFormatSrt         AudioResponseFormat = "srt"
	AudioResponseFormatVerboseJSON AudioResponseFormat = "verbose_json"
	AudioResponseFormatVtt         AudioResponseFormat = "vtt"
)

func (e AudioResponseFormat) ToPointer() *AudioResponseFormat {
	return &e
}
func (e *AudioResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "text":
		fallthrough
	case "srt":
		fallthrough
	case "verbose_json":
		fallthrough
	case "vtt":
		*e = AudioResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AudioResponseFormat: %v", v)
	}
}
