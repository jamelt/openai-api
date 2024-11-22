# Function

The function that the model called.


## Fields

| Field                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Name`                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                                                | The name of the function to call.                                                                                                                                                                                                                                                 |
| `Arguments`                                                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                                                | The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function. |