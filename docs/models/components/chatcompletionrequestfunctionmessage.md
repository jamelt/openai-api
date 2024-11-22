# ~~ChatCompletionRequestFunctionMessage~~

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.


## Fields

| Field                                                                                                                      | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `Role`                                                                                                                     | [components.ChatCompletionRequestFunctionMessageRole](../../models/components/chatcompletionrequestfunctionmessagerole.md) | :heavy_check_mark:                                                                                                         | The role of the messages author, in this case `function`.                                                                  |
| `Content`                                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The contents of the function message.                                                                                      |
| `Name`                                                                                                                     | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The name of the function to call.                                                                                          |