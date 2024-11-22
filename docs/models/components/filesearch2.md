# FileSearch2


## Fields

| Field                                                                                                                                                                                             | Type                                                                                                                                                                                              | Required                                                                                                                                                                                          | Description                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `VectorStoreIds`                                                                                                                                                                                  | []*string*                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                | The [vector store](/docs/api-reference/vector-stores/object) attached to this assistant. There can be a maximum of 1 vector store attached to the assistant.<br/>                                 |
| `VectorStores`                                                                                                                                                                                    | [][components.FileSearchVectorStores](../../models/components/filesearchvectorstores.md)                                                                                                          | :heavy_check_mark:                                                                                                                                                                                | A helper to create a [vector store](/docs/api-reference/vector-stores/object) with file_ids and attach it to this assistant. There can be a maximum of 1 vector store attached to the assistant.<br/> |