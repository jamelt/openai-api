# Input

Input text to embed, encoded as a string or array of tokens. To embed multiple inputs in a single request, pass an array of strings or array of token arrays. The input must not exceed the max input tokens for the model (8192 tokens for `text-embedding-ada-002`), cannot be an empty string, and any array must be 2048 dimensions or less. [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken) for counting tokens.



## Supported Types

### 

```go
input := components.CreateInputStr(string{/* values here */})
```

### 

```go
input := components.CreateInputArrayOfStr([]string{/* values here */})
```

### 

```go
input := components.CreateInputArrayOfInteger([]int64{/* values here */})
```

### 

```go
input := components.CreateInputArrayOfArrayOfInteger([][]int64{/* values here */})
```

