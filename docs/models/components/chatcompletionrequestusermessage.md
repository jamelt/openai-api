# ChatCompletionRequestUserMessage


## Fields

| Field                                                                                                                        | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `Content`                                                                                                                    | [components.ChatCompletionRequestUserMessageContent](../../models/components/chatcompletionrequestusermessagecontent.md)     | :heavy_check_mark:                                                                                                           | The contents of the user message.<br/>                                                                                       |
| `Role`                                                                                                                       | [components.ChatCompletionRequestUserMessageRole](../../models/components/chatcompletionrequestusermessagerole.md)           | :heavy_check_mark:                                                                                                           | The role of the messages author, in this case `user`.                                                                        |
| `Name`                                                                                                                       | **string*                                                                                                                    | :heavy_minus_sign:                                                                                                           | An optional name for the participant. Provides the model information to differentiate between participants of the same role. |