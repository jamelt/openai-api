# RunStepObjectLastError

The last error associated with this run step. Will be `null` if there are no errors.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Code`                                                                       | [components.RunStepObjectCode](../../models/components/runstepobjectcode.md) | :heavy_check_mark:                                                           | One of `server_error` or `rate_limit_exceeded`.                              |
| `Message`                                                                    | *string*                                                                     | :heavy_check_mark:                                                           | A human-readable description of the error.                                   |