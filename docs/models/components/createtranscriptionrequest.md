# CreateTranscriptionRequest


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                                                                                                                                | Example                                                                                                                                                                                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `File`                                                                                                                                                                                                                                                                                                                                                     | [components.File](../../models/components/file.md)                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                         | The audio file object (not file name) to transcribe, in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm.<br/>                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                            |
| `Model`                                                                                                                                                                                                                                                                                                                                                    | [components.CreateTranscriptionRequestModel](../../models/components/createtranscriptionrequestmodel.md)                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                         | ID of the model to use. Only `whisper-1` (which is powered by our open source Whisper V2 model) is currently available.<br/>                                                                                                                                                                                                                               | whisper-1                                                                                                                                                                                                                                                                                                                                                  |
| `Language`                                                                                                                                                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                         | The language of the input audio. Supplying the input language in [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) format will improve accuracy and latency.<br/>                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                            |
| `Prompt`                                                                                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                         | An optional text to guide the model's style or continue a previous audio segment. The [prompt](/docs/guides/speech-to-text#prompting) should match the audio language.<br/>                                                                                                                                                                                |                                                                                                                                                                                                                                                                                                                                                            |
| `ResponseFormat`                                                                                                                                                                                                                                                                                                                                           | [*components.AudioResponseFormat](../../models/components/audioresponseformat.md)                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                         | The format of the output, in one of these options: `json`, `text`, `srt`, `verbose_json`, or `vtt`.<br/>                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                            |
| `Temperature`                                                                                                                                                                                                                                                                                                                                              | **float64*                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                         | The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use [log probability](https://en.wikipedia.org/wiki/Log_probability) to automatically increase the temperature until certain thresholds are hit.<br/> |                                                                                                                                                                                                                                                                                                                                                            |
| `TimestampGranularities`                                                                                                                                                                                                                                                                                                                                   | [][components.TimestampGranularities](../../models/components/timestampgranularities.md)                                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                         | The timestamp granularities to populate for this transcription. `response_format` must be set `verbose_json` to use timestamp granularities. Either or both of these options are supported: `word`, or `segment`. Note: There is no additional latency for segment timestamps, but generating word timestamps incurs additional latency.<br/>              |                                                                                                                                                                                                                                                                                                                                                            |