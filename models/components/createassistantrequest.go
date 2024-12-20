// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

type Two string

const (
	TwoGpt4o             Two = "gpt-4o"
	TwoGpt4o20240806     Two = "gpt-4o-2024-08-06"
	TwoGpt4o20240513     Two = "gpt-4o-2024-05-13"
	TwoGpt4oMini         Two = "gpt-4o-mini"
	TwoGpt4oMini20240718 Two = "gpt-4o-mini-2024-07-18"
	TwoGpt4Turbo         Two = "gpt-4-turbo"
	TwoGpt4Turbo20240409 Two = "gpt-4-turbo-2024-04-09"
	TwoGpt40125Preview   Two = "gpt-4-0125-preview"
	TwoGpt4TurboPreview  Two = "gpt-4-turbo-preview"
	TwoGpt41106Preview   Two = "gpt-4-1106-preview"
	TwoGpt4VisionPreview Two = "gpt-4-vision-preview"
	TwoGpt4              Two = "gpt-4"
	TwoGpt40314          Two = "gpt-4-0314"
	TwoGpt40613          Two = "gpt-4-0613"
	TwoGpt432k           Two = "gpt-4-32k"
	TwoGpt432k0314       Two = "gpt-4-32k-0314"
	TwoGpt432k0613       Two = "gpt-4-32k-0613"
	TwoGpt35Turbo        Two = "gpt-3.5-turbo"
	TwoGpt35Turbo16k     Two = "gpt-3.5-turbo-16k"
	TwoGpt35Turbo0613    Two = "gpt-3.5-turbo-0613"
	TwoGpt35Turbo1106    Two = "gpt-3.5-turbo-1106"
	TwoGpt35Turbo0125    Two = "gpt-3.5-turbo-0125"
	TwoGpt35Turbo16k0613 Two = "gpt-3.5-turbo-16k-0613"
)

func (e Two) ToPointer() *Two {
	return &e
}
func (e *Two) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gpt-4o":
		fallthrough
	case "gpt-4o-2024-08-06":
		fallthrough
	case "gpt-4o-2024-05-13":
		fallthrough
	case "gpt-4o-mini":
		fallthrough
	case "gpt-4o-mini-2024-07-18":
		fallthrough
	case "gpt-4-turbo":
		fallthrough
	case "gpt-4-turbo-2024-04-09":
		fallthrough
	case "gpt-4-0125-preview":
		fallthrough
	case "gpt-4-turbo-preview":
		fallthrough
	case "gpt-4-1106-preview":
		fallthrough
	case "gpt-4-vision-preview":
		fallthrough
	case "gpt-4":
		fallthrough
	case "gpt-4-0314":
		fallthrough
	case "gpt-4-0613":
		fallthrough
	case "gpt-4-32k":
		fallthrough
	case "gpt-4-32k-0314":
		fallthrough
	case "gpt-4-32k-0613":
		fallthrough
	case "gpt-3.5-turbo":
		fallthrough
	case "gpt-3.5-turbo-16k":
		fallthrough
	case "gpt-3.5-turbo-0613":
		fallthrough
	case "gpt-3.5-turbo-1106":
		fallthrough
	case "gpt-3.5-turbo-0125":
		fallthrough
	case "gpt-3.5-turbo-16k-0613":
		*e = Two(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Two: %v", v)
	}
}

type CreateAssistantRequestModelType string

const (
	CreateAssistantRequestModelTypeStr CreateAssistantRequestModelType = "str"
	CreateAssistantRequestModelTypeTwo CreateAssistantRequestModelType = "2"
)

// CreateAssistantRequestModel - ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models) for descriptions of them.
type CreateAssistantRequestModel struct {
	Str *string
	Two *Two

	Type CreateAssistantRequestModelType
}

func CreateCreateAssistantRequestModelStr(str string) CreateAssistantRequestModel {
	typ := CreateAssistantRequestModelTypeStr

	return CreateAssistantRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateAssistantRequestModelTwo(two Two) CreateAssistantRequestModel {
	typ := CreateAssistantRequestModelTypeTwo

	return CreateAssistantRequestModel{
		Two:  &two,
		Type: typ,
	}
}

func (u *CreateAssistantRequestModel) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateAssistantRequestModelTypeStr
		return nil
	}

	var two Two = Two("")
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = CreateAssistantRequestModelTypeTwo
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateAssistantRequestModel", string(data))
}

func (u CreateAssistantRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type CreateAssistantRequestModel: all fields are null")
}

type CreateAssistantRequestToolsType string

const (
	CreateAssistantRequestToolsTypeAssistantToolsCode       CreateAssistantRequestToolsType = "AssistantToolsCode"
	CreateAssistantRequestToolsTypeAssistantToolsFileSearch CreateAssistantRequestToolsType = "AssistantToolsFileSearch"
	CreateAssistantRequestToolsTypeAssistantToolsFunction   CreateAssistantRequestToolsType = "AssistantToolsFunction"
)

type CreateAssistantRequestTools struct {
	AssistantToolsCode       *AssistantToolsCode
	AssistantToolsFileSearch *AssistantToolsFileSearch
	AssistantToolsFunction   *AssistantToolsFunction

	Type CreateAssistantRequestToolsType
}

func CreateCreateAssistantRequestToolsAssistantToolsCode(assistantToolsCode AssistantToolsCode) CreateAssistantRequestTools {
	typ := CreateAssistantRequestToolsTypeAssistantToolsCode

	return CreateAssistantRequestTools{
		AssistantToolsCode: &assistantToolsCode,
		Type:               typ,
	}
}

func CreateCreateAssistantRequestToolsAssistantToolsFileSearch(assistantToolsFileSearch AssistantToolsFileSearch) CreateAssistantRequestTools {
	typ := CreateAssistantRequestToolsTypeAssistantToolsFileSearch

	return CreateAssistantRequestTools{
		AssistantToolsFileSearch: &assistantToolsFileSearch,
		Type:                     typ,
	}
}

func CreateCreateAssistantRequestToolsAssistantToolsFunction(assistantToolsFunction AssistantToolsFunction) CreateAssistantRequestTools {
	typ := CreateAssistantRequestToolsTypeAssistantToolsFunction

	return CreateAssistantRequestTools{
		AssistantToolsFunction: &assistantToolsFunction,
		Type:                   typ,
	}
}

func (u *CreateAssistantRequestTools) UnmarshalJSON(data []byte) error {

	var assistantToolsCode AssistantToolsCode = AssistantToolsCode{}
	if err := utils.UnmarshalJSON(data, &assistantToolsCode, "", true, true); err == nil {
		u.AssistantToolsCode = &assistantToolsCode
		u.Type = CreateAssistantRequestToolsTypeAssistantToolsCode
		return nil
	}

	var assistantToolsFileSearch AssistantToolsFileSearch = AssistantToolsFileSearch{}
	if err := utils.UnmarshalJSON(data, &assistantToolsFileSearch, "", true, true); err == nil {
		u.AssistantToolsFileSearch = &assistantToolsFileSearch
		u.Type = CreateAssistantRequestToolsTypeAssistantToolsFileSearch
		return nil
	}

	var assistantToolsFunction AssistantToolsFunction = AssistantToolsFunction{}
	if err := utils.UnmarshalJSON(data, &assistantToolsFunction, "", true, true); err == nil {
		u.AssistantToolsFunction = &assistantToolsFunction
		u.Type = CreateAssistantRequestToolsTypeAssistantToolsFunction
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateAssistantRequestTools", string(data))
}

func (u CreateAssistantRequestTools) MarshalJSON() ([]byte, error) {
	if u.AssistantToolsCode != nil {
		return utils.MarshalJSON(u.AssistantToolsCode, "", true)
	}

	if u.AssistantToolsFileSearch != nil {
		return utils.MarshalJSON(u.AssistantToolsFileSearch, "", true)
	}

	if u.AssistantToolsFunction != nil {
		return utils.MarshalJSON(u.AssistantToolsFunction, "", true)
	}

	return nil, errors.New("could not marshal union type CreateAssistantRequestTools: all fields are null")
}

type CreateAssistantRequestCodeInterpreter struct {
	// A list of [file](/docs/api-reference/files) IDs made available to the `code_interpreter` tool. There can be a maximum of 20 files associated with the tool.
	//
	FileIds []string `json:"file_ids,omitempty"`
}

func (o *CreateAssistantRequestCodeInterpreter) GetFileIds() []string {
	if o == nil {
		return nil
	}
	return o.FileIds
}

// CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType - Always `static`.
type CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType string

const (
	CreateAssistantRequestChunkingStrategyToolResourcesFileSearchTypeStatic CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType = "static"
)

func (e CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType) ToPointer() *CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType {
	return &e
}
func (e *CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		*e = CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType: %v", v)
	}
}

type ChunkingStrategyStatic struct {
	// The maximum number of tokens in each chunk. The default value is `800`. The minimum value is `100` and the maximum value is `4096`.
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens"`
	// The number of tokens that overlap between chunks. The default value is `400`.
	//
	// Note that the overlap must not exceed half of `max_chunk_size_tokens`.
	//
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens"`
}

func (o *ChunkingStrategyStatic) GetMaxChunkSizeTokens() int64 {
	if o == nil {
		return 0
	}
	return o.MaxChunkSizeTokens
}

func (o *ChunkingStrategyStatic) GetChunkOverlapTokens() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkOverlapTokens
}

type ChunkingStrategyStaticChunkingStrategy struct {
	// Always `static`.
	Type   CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType `json:"type"`
	Static ChunkingStrategyStatic                                            `json:"static"`
}

func (o *ChunkingStrategyStaticChunkingStrategy) GetType() CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType {
	if o == nil {
		return CreateAssistantRequestChunkingStrategyToolResourcesFileSearchType("")
	}
	return o.Type
}

func (o *ChunkingStrategyStaticChunkingStrategy) GetStatic() ChunkingStrategyStatic {
	if o == nil {
		return ChunkingStrategyStatic{}
	}
	return o.Static
}

// CreateAssistantRequestChunkingStrategyToolResourcesType - Always `auto`.
type CreateAssistantRequestChunkingStrategyToolResourcesType string

const (
	CreateAssistantRequestChunkingStrategyToolResourcesTypeAuto CreateAssistantRequestChunkingStrategyToolResourcesType = "auto"
)

func (e CreateAssistantRequestChunkingStrategyToolResourcesType) ToPointer() *CreateAssistantRequestChunkingStrategyToolResourcesType {
	return &e
}
func (e *CreateAssistantRequestChunkingStrategyToolResourcesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = CreateAssistantRequestChunkingStrategyToolResourcesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAssistantRequestChunkingStrategyToolResourcesType: %v", v)
	}
}

// ChunkingStrategyAutoChunkingStrategy - The default strategy. This strategy currently uses a `max_chunk_size_tokens` of `800` and `chunk_overlap_tokens` of `400`.
type ChunkingStrategyAutoChunkingStrategy struct {
	// Always `auto`.
	Type CreateAssistantRequestChunkingStrategyToolResourcesType `json:"type"`
}

func (o *ChunkingStrategyAutoChunkingStrategy) GetType() CreateAssistantRequestChunkingStrategyToolResourcesType {
	if o == nil {
		return CreateAssistantRequestChunkingStrategyToolResourcesType("")
	}
	return o.Type
}

type FileSearchChunkingStrategyType string

const (
	FileSearchChunkingStrategyTypeChunkingStrategyAutoChunkingStrategy   FileSearchChunkingStrategyType = "chunking_strategy_Auto Chunking Strategy"
	FileSearchChunkingStrategyTypeChunkingStrategyStaticChunkingStrategy FileSearchChunkingStrategyType = "chunking_strategy_Static Chunking Strategy"
)

// FileSearchChunkingStrategy - The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
type FileSearchChunkingStrategy struct {
	ChunkingStrategyAutoChunkingStrategy   *ChunkingStrategyAutoChunkingStrategy
	ChunkingStrategyStaticChunkingStrategy *ChunkingStrategyStaticChunkingStrategy

	Type FileSearchChunkingStrategyType
}

func CreateFileSearchChunkingStrategyChunkingStrategyAutoChunkingStrategy(chunkingStrategyAutoChunkingStrategy ChunkingStrategyAutoChunkingStrategy) FileSearchChunkingStrategy {
	typ := FileSearchChunkingStrategyTypeChunkingStrategyAutoChunkingStrategy

	return FileSearchChunkingStrategy{
		ChunkingStrategyAutoChunkingStrategy: &chunkingStrategyAutoChunkingStrategy,
		Type:                                 typ,
	}
}

func CreateFileSearchChunkingStrategyChunkingStrategyStaticChunkingStrategy(chunkingStrategyStaticChunkingStrategy ChunkingStrategyStaticChunkingStrategy) FileSearchChunkingStrategy {
	typ := FileSearchChunkingStrategyTypeChunkingStrategyStaticChunkingStrategy

	return FileSearchChunkingStrategy{
		ChunkingStrategyStaticChunkingStrategy: &chunkingStrategyStaticChunkingStrategy,
		Type:                                   typ,
	}
}

func (u *FileSearchChunkingStrategy) UnmarshalJSON(data []byte) error {

	var chunkingStrategyAutoChunkingStrategy ChunkingStrategyAutoChunkingStrategy = ChunkingStrategyAutoChunkingStrategy{}
	if err := utils.UnmarshalJSON(data, &chunkingStrategyAutoChunkingStrategy, "", true, true); err == nil {
		u.ChunkingStrategyAutoChunkingStrategy = &chunkingStrategyAutoChunkingStrategy
		u.Type = FileSearchChunkingStrategyTypeChunkingStrategyAutoChunkingStrategy
		return nil
	}

	var chunkingStrategyStaticChunkingStrategy ChunkingStrategyStaticChunkingStrategy = ChunkingStrategyStaticChunkingStrategy{}
	if err := utils.UnmarshalJSON(data, &chunkingStrategyStaticChunkingStrategy, "", true, true); err == nil {
		u.ChunkingStrategyStaticChunkingStrategy = &chunkingStrategyStaticChunkingStrategy
		u.Type = FileSearchChunkingStrategyTypeChunkingStrategyStaticChunkingStrategy
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for FileSearchChunkingStrategy", string(data))
}

func (u FileSearchChunkingStrategy) MarshalJSON() ([]byte, error) {
	if u.ChunkingStrategyAutoChunkingStrategy != nil {
		return utils.MarshalJSON(u.ChunkingStrategyAutoChunkingStrategy, "", true)
	}

	if u.ChunkingStrategyStaticChunkingStrategy != nil {
		return utils.MarshalJSON(u.ChunkingStrategyStaticChunkingStrategy, "", true)
	}

	return nil, errors.New("could not marshal union type FileSearchChunkingStrategy: all fields are null")
}

// CreateAssistantRequestFileSearchMetadata - Set of 16 key-value pairs that can be attached to a vector store. This can be useful for storing additional information about the vector store in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
type CreateAssistantRequestFileSearchMetadata struct {
}

type FileSearchVectorStores struct {
	// A list of [file](/docs/api-reference/files) IDs to add to the vector store. There can be a maximum of 10000 files in a vector store.
	//
	FileIds []string `json:"file_ids,omitempty"`
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
	ChunkingStrategy *FileSearchChunkingStrategy `json:"chunking_strategy,omitempty"`
	// Set of 16 key-value pairs that can be attached to a vector store. This can be useful for storing additional information about the vector store in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	//
	Metadata *CreateAssistantRequestFileSearchMetadata `json:"metadata,omitempty"`
}

func (o *FileSearchVectorStores) GetFileIds() []string {
	if o == nil {
		return nil
	}
	return o.FileIds
}

func (o *FileSearchVectorStores) GetChunkingStrategy() *FileSearchChunkingStrategy {
	if o == nil {
		return nil
	}
	return o.ChunkingStrategy
}

func (o *FileSearchVectorStores) GetMetadata() *CreateAssistantRequestFileSearchMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type FileSearch2 struct {
	// The [vector store](/docs/api-reference/vector-stores/object) attached to this assistant. There can be a maximum of 1 vector store attached to the assistant.
	//
	VectorStoreIds []string `json:"vector_store_ids,omitempty"`
	// A helper to create a [vector store](/docs/api-reference/vector-stores/object) with file_ids and attach it to this assistant. There can be a maximum of 1 vector store attached to the assistant.
	//
	VectorStores []FileSearchVectorStores `json:"vector_stores"`
}

func (o *FileSearch2) GetVectorStoreIds() []string {
	if o == nil {
		return nil
	}
	return o.VectorStoreIds
}

func (o *FileSearch2) GetVectorStores() []FileSearchVectorStores {
	if o == nil {
		return []FileSearchVectorStores{}
	}
	return o.VectorStores
}

// CreateAssistantRequestChunkingStrategyType - Always `static`.
type CreateAssistantRequestChunkingStrategyType string

const (
	CreateAssistantRequestChunkingStrategyTypeStatic CreateAssistantRequestChunkingStrategyType = "static"
)

func (e CreateAssistantRequestChunkingStrategyType) ToPointer() *CreateAssistantRequestChunkingStrategyType {
	return &e
}
func (e *CreateAssistantRequestChunkingStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		*e = CreateAssistantRequestChunkingStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAssistantRequestChunkingStrategyType: %v", v)
	}
}

type Static struct {
	// The maximum number of tokens in each chunk. The default value is `800`. The minimum value is `100` and the maximum value is `4096`.
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens"`
	// The number of tokens that overlap between chunks. The default value is `400`.
	//
	// Note that the overlap must not exceed half of `max_chunk_size_tokens`.
	//
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens"`
}

func (o *Static) GetMaxChunkSizeTokens() int64 {
	if o == nil {
		return 0
	}
	return o.MaxChunkSizeTokens
}

func (o *Static) GetChunkOverlapTokens() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkOverlapTokens
}

type CreateAssistantRequestChunkingStrategyStaticChunkingStrategy struct {
	// Always `static`.
	Type   CreateAssistantRequestChunkingStrategyType `json:"type"`
	Static Static                                     `json:"static"`
}

func (o *CreateAssistantRequestChunkingStrategyStaticChunkingStrategy) GetType() CreateAssistantRequestChunkingStrategyType {
	if o == nil {
		return CreateAssistantRequestChunkingStrategyType("")
	}
	return o.Type
}

func (o *CreateAssistantRequestChunkingStrategyStaticChunkingStrategy) GetStatic() Static {
	if o == nil {
		return Static{}
	}
	return o.Static
}

// ChunkingStrategyType - Always `auto`.
type ChunkingStrategyType string

const (
	ChunkingStrategyTypeAuto ChunkingStrategyType = "auto"
)

func (e ChunkingStrategyType) ToPointer() *ChunkingStrategyType {
	return &e
}
func (e *ChunkingStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = ChunkingStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChunkingStrategyType: %v", v)
	}
}

// AutoChunkingStrategy - The default strategy. This strategy currently uses a `max_chunk_size_tokens` of `800` and `chunk_overlap_tokens` of `400`.
type AutoChunkingStrategy struct {
	// Always `auto`.
	Type ChunkingStrategyType `json:"type"`
}

func (o *AutoChunkingStrategy) GetType() ChunkingStrategyType {
	if o == nil {
		return ChunkingStrategyType("")
	}
	return o.Type
}

type CreateAssistantRequestFileSearchChunkingStrategyType string

const (
	CreateAssistantRequestFileSearchChunkingStrategyTypeAutoChunkingStrategy                                         CreateAssistantRequestFileSearchChunkingStrategyType = "Auto Chunking Strategy"
	CreateAssistantRequestFileSearchChunkingStrategyTypeCreateAssistantRequestChunkingStrategyStaticChunkingStrategy CreateAssistantRequestFileSearchChunkingStrategyType = "CreateAssistantRequest_chunking_strategy_Static Chunking Strategy"
)

// CreateAssistantRequestFileSearchChunkingStrategy - The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
type CreateAssistantRequestFileSearchChunkingStrategy struct {
	AutoChunkingStrategy                                         *AutoChunkingStrategy
	CreateAssistantRequestChunkingStrategyStaticChunkingStrategy *CreateAssistantRequestChunkingStrategyStaticChunkingStrategy

	Type CreateAssistantRequestFileSearchChunkingStrategyType
}

func CreateCreateAssistantRequestFileSearchChunkingStrategyAutoChunkingStrategy(autoChunkingStrategy AutoChunkingStrategy) CreateAssistantRequestFileSearchChunkingStrategy {
	typ := CreateAssistantRequestFileSearchChunkingStrategyTypeAutoChunkingStrategy

	return CreateAssistantRequestFileSearchChunkingStrategy{
		AutoChunkingStrategy: &autoChunkingStrategy,
		Type:                 typ,
	}
}

func CreateCreateAssistantRequestFileSearchChunkingStrategyCreateAssistantRequestChunkingStrategyStaticChunkingStrategy(createAssistantRequestChunkingStrategyStaticChunkingStrategy CreateAssistantRequestChunkingStrategyStaticChunkingStrategy) CreateAssistantRequestFileSearchChunkingStrategy {
	typ := CreateAssistantRequestFileSearchChunkingStrategyTypeCreateAssistantRequestChunkingStrategyStaticChunkingStrategy

	return CreateAssistantRequestFileSearchChunkingStrategy{
		CreateAssistantRequestChunkingStrategyStaticChunkingStrategy: &createAssistantRequestChunkingStrategyStaticChunkingStrategy,
		Type: typ,
	}
}

func (u *CreateAssistantRequestFileSearchChunkingStrategy) UnmarshalJSON(data []byte) error {

	var autoChunkingStrategy AutoChunkingStrategy = AutoChunkingStrategy{}
	if err := utils.UnmarshalJSON(data, &autoChunkingStrategy, "", true, true); err == nil {
		u.AutoChunkingStrategy = &autoChunkingStrategy
		u.Type = CreateAssistantRequestFileSearchChunkingStrategyTypeAutoChunkingStrategy
		return nil
	}

	var createAssistantRequestChunkingStrategyStaticChunkingStrategy CreateAssistantRequestChunkingStrategyStaticChunkingStrategy = CreateAssistantRequestChunkingStrategyStaticChunkingStrategy{}
	if err := utils.UnmarshalJSON(data, &createAssistantRequestChunkingStrategyStaticChunkingStrategy, "", true, true); err == nil {
		u.CreateAssistantRequestChunkingStrategyStaticChunkingStrategy = &createAssistantRequestChunkingStrategyStaticChunkingStrategy
		u.Type = CreateAssistantRequestFileSearchChunkingStrategyTypeCreateAssistantRequestChunkingStrategyStaticChunkingStrategy
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateAssistantRequestFileSearchChunkingStrategy", string(data))
}

func (u CreateAssistantRequestFileSearchChunkingStrategy) MarshalJSON() ([]byte, error) {
	if u.AutoChunkingStrategy != nil {
		return utils.MarshalJSON(u.AutoChunkingStrategy, "", true)
	}

	if u.CreateAssistantRequestChunkingStrategyStaticChunkingStrategy != nil {
		return utils.MarshalJSON(u.CreateAssistantRequestChunkingStrategyStaticChunkingStrategy, "", true)
	}

	return nil, errors.New("could not marshal union type CreateAssistantRequestFileSearchChunkingStrategy: all fields are null")
}

// FileSearchMetadata - Set of 16 key-value pairs that can be attached to a vector store. This can be useful for storing additional information about the vector store in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
type FileSearchMetadata struct {
}

type VectorStores struct {
	// A list of [file](/docs/api-reference/files) IDs to add to the vector store. There can be a maximum of 10000 files in a vector store.
	//
	FileIds []string `json:"file_ids,omitempty"`
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
	ChunkingStrategy *CreateAssistantRequestFileSearchChunkingStrategy `json:"chunking_strategy,omitempty"`
	// Set of 16 key-value pairs that can be attached to a vector store. This can be useful for storing additional information about the vector store in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	//
	Metadata *FileSearchMetadata `json:"metadata,omitempty"`
}

func (o *VectorStores) GetFileIds() []string {
	if o == nil {
		return nil
	}
	return o.FileIds
}

func (o *VectorStores) GetChunkingStrategy() *CreateAssistantRequestFileSearchChunkingStrategy {
	if o == nil {
		return nil
	}
	return o.ChunkingStrategy
}

func (o *VectorStores) GetMetadata() *FileSearchMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type FileSearch1 struct {
	// The [vector store](/docs/api-reference/vector-stores/object) attached to this assistant. There can be a maximum of 1 vector store attached to the assistant.
	//
	VectorStoreIds []string `json:"vector_store_ids"`
	// A helper to create a [vector store](/docs/api-reference/vector-stores/object) with file_ids and attach it to this assistant. There can be a maximum of 1 vector store attached to the assistant.
	//
	VectorStores []VectorStores `json:"vector_stores,omitempty"`
}

func (o *FileSearch1) GetVectorStoreIds() []string {
	if o == nil {
		return []string{}
	}
	return o.VectorStoreIds
}

func (o *FileSearch1) GetVectorStores() []VectorStores {
	if o == nil {
		return nil
	}
	return o.VectorStores
}

type CreateAssistantRequestFileSearchType string

const (
	CreateAssistantRequestFileSearchTypeFileSearch1 CreateAssistantRequestFileSearchType = "file_search_1"
	CreateAssistantRequestFileSearchTypeFileSearch2 CreateAssistantRequestFileSearchType = "file_search_2"
)

type CreateAssistantRequestFileSearch struct {
	FileSearch1 *FileSearch1
	FileSearch2 *FileSearch2

	Type CreateAssistantRequestFileSearchType
}

func CreateCreateAssistantRequestFileSearchFileSearch1(fileSearch1 FileSearch1) CreateAssistantRequestFileSearch {
	typ := CreateAssistantRequestFileSearchTypeFileSearch1

	return CreateAssistantRequestFileSearch{
		FileSearch1: &fileSearch1,
		Type:        typ,
	}
}

func CreateCreateAssistantRequestFileSearchFileSearch2(fileSearch2 FileSearch2) CreateAssistantRequestFileSearch {
	typ := CreateAssistantRequestFileSearchTypeFileSearch2

	return CreateAssistantRequestFileSearch{
		FileSearch2: &fileSearch2,
		Type:        typ,
	}
}

func (u *CreateAssistantRequestFileSearch) UnmarshalJSON(data []byte) error {

	var fileSearch1 FileSearch1 = FileSearch1{}
	if err := utils.UnmarshalJSON(data, &fileSearch1, "", true, true); err == nil {
		u.FileSearch1 = &fileSearch1
		u.Type = CreateAssistantRequestFileSearchTypeFileSearch1
		return nil
	}

	var fileSearch2 FileSearch2 = FileSearch2{}
	if err := utils.UnmarshalJSON(data, &fileSearch2, "", true, true); err == nil {
		u.FileSearch2 = &fileSearch2
		u.Type = CreateAssistantRequestFileSearchTypeFileSearch2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateAssistantRequestFileSearch", string(data))
}

func (u CreateAssistantRequestFileSearch) MarshalJSON() ([]byte, error) {
	if u.FileSearch1 != nil {
		return utils.MarshalJSON(u.FileSearch1, "", true)
	}

	if u.FileSearch2 != nil {
		return utils.MarshalJSON(u.FileSearch2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateAssistantRequestFileSearch: all fields are null")
}

// CreateAssistantRequestToolResources - A set of resources that are used by the assistant's tools. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs.
type CreateAssistantRequestToolResources struct {
	CodeInterpreter *CreateAssistantRequestCodeInterpreter `json:"code_interpreter,omitempty"`
	FileSearch      *CreateAssistantRequestFileSearch      `json:"file_search,omitempty"`
}

func (o *CreateAssistantRequestToolResources) GetCodeInterpreter() *CreateAssistantRequestCodeInterpreter {
	if o == nil {
		return nil
	}
	return o.CodeInterpreter
}

func (o *CreateAssistantRequestToolResources) GetFileSearch() *CreateAssistantRequestFileSearch {
	if o == nil {
		return nil
	}
	return o.FileSearch
}

// CreateAssistantRequestMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
type CreateAssistantRequestMetadata struct {
}

type CreateAssistantRequest struct {
	// ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models) for descriptions of them.
	//
	Model CreateAssistantRequestModel `json:"model"`
	// The name of the assistant. The maximum length is 256 characters.
	//
	Name *string `json:"name,omitempty"`
	// The description of the assistant. The maximum length is 512 characters.
	//
	Description *string `json:"description,omitempty"`
	// The system instructions that the assistant uses. The maximum length is 256,000 characters.
	//
	Instructions *string `json:"instructions,omitempty"`
	// A list of tool enabled on the assistant. There can be a maximum of 128 tools per assistant. Tools can be of types `code_interpreter`, `file_search`, or `function`.
	//
	Tools []CreateAssistantRequestTools `json:"tools,omitempty"`
	// A set of resources that are used by the assistant's tools. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs.
	//
	ToolResources *CreateAssistantRequestToolResources `json:"tool_resources,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	//
	Metadata *CreateAssistantRequestMetadata `json:"metadata,omitempty"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	//
	Temperature *float64 `default:"1" json:"temperature"`
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or temperature but not both.
	//
	TopP *float64 `default:"1" json:"top_p"`
	// Specifies the format that the model must output. Compatible with [GPT-4o](/docs/models#gpt-4o), [GPT-4 Turbo](/docs/models#gpt-4-turbo-and-gpt-4), and all GPT-3.5 Turbo models since `gpt-3.5-turbo-1106`.
	//
	// Setting to `{ "type": "json_schema", "json_schema": {...} }` enables Structured Outputs which ensures the model will match your supplied JSON schema. Learn more in the [Structured Outputs guide](/docs/guides/structured-outputs).
	//
	// Setting to `{ "type": "json_object" }` enables JSON mode, which ensures the message the model generates is valid JSON.
	//
	// **Important:** when using JSON mode, you **must** also instruct the model to produce JSON yourself via a system or user message. Without this, the model may generate an unending stream of whitespace until the generation reaches the token limit, resulting in a long-running and seemingly "stuck" request. Also note that the message content may be partially cut off if `finish_reason="length"`, which indicates the generation exceeded `max_tokens` or the conversation exceeded the max context length.
	//
	ResponseFormat *AssistantsAPIResponseFormatOption `json:"response_format,omitempty"`
}

func (c CreateAssistantRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAssistantRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAssistantRequest) GetModel() CreateAssistantRequestModel {
	if o == nil {
		return CreateAssistantRequestModel{}
	}
	return o.Model
}

func (o *CreateAssistantRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateAssistantRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateAssistantRequest) GetInstructions() *string {
	if o == nil {
		return nil
	}
	return o.Instructions
}

func (o *CreateAssistantRequest) GetTools() []CreateAssistantRequestTools {
	if o == nil {
		return nil
	}
	return o.Tools
}

func (o *CreateAssistantRequest) GetToolResources() *CreateAssistantRequestToolResources {
	if o == nil {
		return nil
	}
	return o.ToolResources
}

func (o *CreateAssistantRequest) GetMetadata() *CreateAssistantRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CreateAssistantRequest) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}

func (o *CreateAssistantRequest) GetTopP() *float64 {
	if o == nil {
		return nil
	}
	return o.TopP
}

func (o *CreateAssistantRequest) GetResponseFormat() *AssistantsAPIResponseFormatOption {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}
