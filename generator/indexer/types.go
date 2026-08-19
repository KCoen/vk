package indexer

import (
	"go.cld.moe/vk_google/generator/docparser"
	"go.cld.moe/vk_google/generator/parser"
)

// HandleKind represents dispatchable vs non-dispatchable handles.
type HandleKind int

const (
	HandleDispatchable HandleKind = iota
	HandleNonDispatchable
)

// HandleInfo holds metadata for a Vulkan handle.
type HandleInfo struct {
	Name        string
	GoName      string
	Parent      string
	ObjTypeEnum string
	Kind        HandleKind
	Comment     string
	DocDesc     string
	ShortDesc   string
	Api         string
}

// EnumValue holds metadata for an individual enum or bitmask bit.
type EnumValue struct {
	Name        string
	GoName      string
	Value       int64
	ValueStr    string
	ValueInt    int64
	Bitpos      int
	ExtNumber   int
	Offset      int
	Dir         string
	IsAlias     bool
	AliasOf     string
	ExtendsEnum string
	Comment     string
	DocDesc     string
	IsBit64     bool
	Api         string
}

// EnumGroup holds metadata for an enum group (<enums> tag).
type EnumGroup struct {
	Name      string
	GoName    string
	Type      string // "enum" or "bitmask"
	Bitwidth  int    // 32 or 64
	Comment   string
	DocDesc   string
	ShortDesc string
	Values    []EnumValue
	Api       string
}

// BitmaskInfo holds metadata for a bitmask type (<type category="bitmask">).
type BitmaskInfo struct {
	Name      string
	GoName    string
	BitsType  string // Name of the associated Vk...FlagBits enum group, if any
	Bitwidth  int    // 32 or 64
	Comment   string
	DocDesc   string
	ShortDesc string
	Api       string
}

// StructMember holds metadata for a single field in a struct/union.
type StructMember struct {
	Name              string
	GoName            string
	CType             string
	GoRawType         string
	GoHighType        string
	IsPointer         bool
	IsDoublePointer   bool
	IsConst           bool
	IsSlice           bool
	IsString          bool
	IsStringSlice     bool
	IsFuncPointer     bool
	SliceCountMember  string // name of the member holding the count for this slice
	LenAttr           string // length attribute from XML
	ArrayDimensions   []string
	ArraySizeInt      int
	DefaultValue      string // values="..." attribute (e.g. VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO)
	GoDefaultValue    string // GoName for default value
	Comment           string
	DocDesc           string
	Optional          bool
	ReturnedOnly      bool
	Api               string
}

// StructInfo holds metadata for a struct or union.
type StructInfo struct {
	Name           string
	GoName         string
	Category       string // "struct" or "union"
	Comment        string
	DocDesc        string
	ShortDesc      string
	StructExtends  []string
	ReturnedOnly   bool
	Members        []StructMember
	HasSType       bool
	HasNext        bool
	STypeDefault   string
	GoSTypeDefault string
	IsAlias        bool
	AliasOf        string
	Api            string
}

// CommandParam holds metadata for a command parameter.
type CommandParam struct {
	Name             string
	GoName           string
	CType            string
	GoRawType        string
	GoHighType       string
	IsPointer        bool
	IsDoublePointer  bool
	IsConst          bool
	IsSlice          bool
	IsString         bool
	IsStringSlice    bool
	IsCountParam     bool // whether this param is purely a count for another slice param
	PairedSliceParam string
	IsOutputQuery    bool // whether this is part of 2-call query (e.g. pCount, pItems)
	IsOutputSingle   bool // whether this is a single output pointer (e.g. *Instance)
	LenAttr          string
	Optional         bool
	Comment          string
	DocDesc          string
	Api              string
}

// CommandInfo holds metadata for a Vulkan command.
type CommandInfo struct {
	Name             string
	GoName           string
	ReturnCType      string
	ReturnGoType     string
	SuccessCodes     []string
	ErrorCodes       []string
	Queues           string
	RenderPass       string
	CmdBufferLvl     string
	Comment          string
	DocDesc          string
	ShortDesc        string
	ParamDocs        map[string]string
	Params           []CommandParam
	HighParams       []CommandParam // params exposed in the high-level Go API
	HighReturnTypes  []string       // return types in high-level Go API
	IsOutputQuery    bool           // 2-call query pattern
	OutputQueryParam *CommandParam
	OutputCountParam *CommandParam
	IsAlias          bool
	AliasOf          string
	IsInstanceCmd    bool
	IsDeviceCmd      bool
	IsGlobalCmd      bool
	Api              string
}

// ConstantInfo holds metadata for an API constant.
type ConstantInfo struct {
	Name    string
	GoName  string
	Value   string
	Comment string
	Type    string
	Api     string
}

// ApiBranchInfo holds commands and metadata for an API branch (vulkan, vulkanbase, vulkansc).
type ApiBranchInfo struct {
	Name            string // "vulkan", "vulkanbase", "vulkansc"
	PkgName         string // "vulkan", "vulkanbase", "vulkansc"
	Title           string // "Vulkan Core API", "Vulkan Base API", "Vulkan SC API"
	Commands        []string
	BranchCommands  map[string]*CommandInfo
	Handles         map[string]HandleInfo
	Bitmasks        map[string]BitmaskInfo
	Structs         map[string]*StructInfo
	EnumGroups      map[string]*EnumGroup
	Constants       map[string]ConstantInfo
	TypeAliases     map[string]TypeAliasInfo
	FuncPointerDefs map[string]*FuncPointerInfo
}

// VersionInfo holds items introduced in a Vulkan core version.
type VersionInfo struct {
	Number   string // "1.0", "1.1", "1.2", "1.3", "1.4"
	PkgName  string // "v1_0", "v1_1", "v1_2", "v1_3", "v1_4"
	Types    []string
	Enums    []string
	Commands []string
}

// ExtensionInfo holds metadata for a Vulkan extension.
type ExtensionInfo struct {
	Name         string
	GoName       string
	PkgName      string
	Number       int
	Type         string // "instance" or "device"
	Author       string
	Contact      string
	Supported    string
	PromotedTo   string
	DeprecatedBy string
	Depends      string
	Platform     string
	Comment      string
	DocDesc      string
	ShortDesc    string
	Types        []string
	Enums        []string
	Commands     []string
}

// FormatInfo holds format metadata from <formats>.
type FormatInfo struct {
	Name           string
	GoName         string
	Class          string
	BlockSize      string
	TexelsPerBlock string
	Packed         string
	Components     []parser.Component
	Planes         []parser.Plane
}

// SyncStageInfo holds sync stage metadata from <sync>.
type SyncStageInfo struct {
	Name    string
	GoName  string
	Alias   string
	Support string
}

// SyncAccessInfo holds sync access metadata from <sync>.
type SyncAccessInfo struct {
	Name    string
	GoName  string
	Alias   string
	Support string
}

// FuncPointerParam represents a parameter of a callback function pointer.
type FuncPointerParam struct {
	Name      string
	GoName    string
	Type      string
	GoType    string
	IsPointer bool
	IsConst   bool
}

// FuncPointerInfo holds metadata for a callback function pointer type.
type FuncPointerInfo struct {
	Name         string
	GoName       string
	ReturnType   string
	ReturnGoType string
	Params       []FuncPointerParam
	Comment      string
	Api          string
}

// TypeAliasInfo holds metadata for a type alias.
type TypeAliasInfo struct {
	Name     string
	GoName   string
	AliasOf  string
	GoAlias  string
	Category string // "enum", "bitmask", "handle", "struct"
	Comment  string
	Api      string
}

// Index holds the fully indexed and resolved Vulkan registry.
type Index struct {
	Platforms        map[string]parser.Platform
	Tags             map[string]parser.Tag
	Constants        map[string]ConstantInfo
	Handles          map[string]HandleInfo
	BaseTypes        map[string]string // e.g. "VkDeviceSize" -> "uint64"
	EnumGroups       map[string]*EnumGroup
	Bitmasks         map[string]BitmaskInfo
	Structs          map[string]*StructInfo
	TypeAliases      map[string]TypeAliasInfo
	FuncPointers     map[string]string
	FuncPointerDefs  map[string]*FuncPointerInfo
	Commands         map[string]*CommandInfo
	ApiBranches      map[string]*ApiBranchInfo
	Versions         map[string]*VersionInfo
	Extensions       map[string]*ExtensionInfo
	Formats          map[string]FormatInfo
	SyncStages       map[string]SyncStageInfo
	SyncAccesses     map[string]SyncAccessInfo
	SpirvExts        []parser.SpirvExtension
	SpirvCaps        []parser.SpirvCapability
	StructExtensions map[string][]string
	DocIndex         *docparser.DocIndex
}
