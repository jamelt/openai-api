# Results


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Flagged`                                                                                    | *bool*                                                                                       | :heavy_check_mark:                                                                           | Whether any of the below categories are flagged.                                             |
| `Categories`                                                                                 | [components.Categories](../../models/components/categories.md)                               | :heavy_check_mark:                                                                           | A list of the categories, and whether they are flagged or not.                               |
| `CategoryScores`                                                                             | [components.CategoryScores](../../models/components/categoryscores.md)                       | :heavy_check_mark:                                                                           | A list of the categories along with their scores as predicted by model.                      |
| `CategoryAppliedInputTypes`                                                                  | [components.CategoryAppliedInputTypes](../../models/components/categoryappliedinputtypes.md) | :heavy_check_mark:                                                                           | A list of the categories along with the input type(s) that the score applies to.             |