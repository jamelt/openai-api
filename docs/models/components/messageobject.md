# MessageObject

Represents a message within a [thread](/docs/api-reference/threads).


## Fields

| Field                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ID`                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                           | The identifier, which can be referenced in API endpoints.                                                                                                                                                                                                    |
| `Object`                                                                                                                                                                                                                                                     | [components.MessageObjectObject](../../models/components/messageobjectobject.md)                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                           | The object type, which is always `thread.message`.                                                                                                                                                                                                           |
| `CreatedAt`                                                                                                                                                                                                                                                  | *int64*                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                           | The Unix timestamp (in seconds) for when the message was created.                                                                                                                                                                                            |
| `ThreadID`                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                           | The [thread](/docs/api-reference/threads) ID that this message belongs to.                                                                                                                                                                                   |
| `Status`                                                                                                                                                                                                                                                     | [components.MessageObjectStatus](../../models/components/messageobjectstatus.md)                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                           | The status of the message, which can be either `in_progress`, `incomplete`, or `completed`.                                                                                                                                                                  |
| `IncompleteDetails`                                                                                                                                                                                                                                          | [components.MessageObjectIncompleteDetails](../../models/components/messageobjectincompletedetails.md)                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                           | On an incomplete message, details about why the message is incomplete.                                                                                                                                                                                       |
| `CompletedAt`                                                                                                                                                                                                                                                | *int64*                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                           | The Unix timestamp (in seconds) for when the message was completed.                                                                                                                                                                                          |
| `IncompleteAt`                                                                                                                                                                                                                                               | *int64*                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                           | The Unix timestamp (in seconds) for when the message was marked as incomplete.                                                                                                                                                                               |
| `Role`                                                                                                                                                                                                                                                       | [components.MessageObjectRole](../../models/components/messageobjectrole.md)                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                           | The entity that produced the message. One of `user` or `assistant`.                                                                                                                                                                                          |
| `Content`                                                                                                                                                                                                                                                    | [][components.MessageObjectContent](../../models/components/messageobjectcontent.md)                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                           | The content of the message in array of text and/or images.                                                                                                                                                                                                   |
| `AssistantID`                                                                                                                                                                                                                                                | *string*                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                           | If applicable, the ID of the [assistant](/docs/api-reference/assistants) that authored this message.                                                                                                                                                         |
| `RunID`                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                           | The ID of the [run](/docs/api-reference/runs) associated with the creation of this message. Value is `null` when messages are created manually using the create message or create thread endpoints.                                                          |
| `Attachments`                                                                                                                                                                                                                                                | [][components.MessageObjectAttachments](../../models/components/messageobjectattachments.md)                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                           | A list of files attached to the message, and the tools they were added to.                                                                                                                                                                                   |
| `Metadata`                                                                                                                                                                                                                                                   | [components.MessageObjectMetadata](../../models/components/messageobjectmetadata.md)                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                           | Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.<br/> |