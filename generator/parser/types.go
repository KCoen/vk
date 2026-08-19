package parser

import "encoding/xml"

// Registry represents the root <registry> element of vk.xml.
type Registry struct {
	XMLName           xml.Name          `xml:"registry"`
	Comments          []string          `xml:"comment"`
	Platforms         Platforms         `xml:"platforms"`
	Tags              Tags              `xml:"tags"`
	Types             Types             `xml:"types"`
	Enums             []Enums           `xml:"enums"`
	Commands          Commands          `xml:"commands"`
	Features          []Feature         `xml:"feature"`
	Extensions        Extensions        `xml:"extensions"`
	Formats           Formats           `xml:"formats"`
	Sync              Sync              `xml:"sync"`
	SpirvExtensions   SpirvExtensions   `xml:"spirvextensions"`
	SpirvCapabilities SpirvCapabilities `xml:"spirvcapabilities"`
}

// Platforms is a list of <platform> elements.
type Platforms struct {
	Platforms []Platform `xml:"platform"`
}

// Platform represents a <platform> element.
type Platform struct {
	Name    string `xml:"name,attr"`
	Protect string `xml:"protect,attr"`
	Comment string `xml:"comment,attr"`
}

// Tags is a list of <tag> elements.
type Tags struct {
	Tags []Tag `xml:"tag"`
}

// Tag represents a <tag> element.
type Tag struct {
	Name    string `xml:"name,attr"`
	Author  string `xml:"author,attr"`
	Contact string `xml:"contact,attr"`
	Comment string `xml:"comment,attr"`
}

// Types is a list of <type> elements.
type Types struct {
	Types []Type `xml:"type"`
}

// Type represents a <type> element in <types>.
type Type struct {
	Name            string   `xml:"name,attr"`
	Category        string   `xml:"category,attr"`
	Requires        string   `xml:"requires,attr"`
	Bitvalues       string   `xml:"bitvalues,attr"`
	Parent          string   `xml:"parent,attr"`
	ObjTypeEnum     string   `xml:"objtypeenum,attr"`
	StructExtends   string   `xml:"structextends,attr"`
	ReturnedOnly    string   `xml:"returnedonly,attr"`
	AllowDuplicate  string   `xml:"allowduplicate,attr"`
	Alias           string   `xml:"alias,attr"`
	Api             string   `xml:"api,attr"`
	Comment         string   `xml:"comment,attr"`
	Members         []Member `xml:"member"`
	InnerXML        string   `xml:",innerxml"`
	NameTag         string   `xml:"name"`
	TypeTag         string   `xml:"type"`
	Proto           Proto    `xml:"proto"`
	Params          []Param  `xml:"param"`
}

// Member represents a <member> element in struct/union <type>.
type Member struct {
	Values                string `xml:"values,attr"`
	Len                   string `xml:"len,attr"`
	AltLen                string `xml:"altlen,attr"`
	Optional              string `xml:"optional,attr"`
	NoAutoValidity        string `xml:"noautovalidity,attr"`
	ValidExtensionStructs string `xml:"validextensionstructs,attr"`
	Deprecated            string `xml:"deprecated,attr"`
	Api                   string `xml:"api,attr"`
	CommentAttr           string `xml:"comment,attr"`
	CommentTag            string `xml:"comment"`
	InnerXML              string `xml:",innerxml"`
	Name                  string `xml:"name"`
	Type                  string `xml:"type"`
	Enum                  string `xml:"enum"`
}

// Enums represents an <enums> group element.
type Enums struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr"`
	Comment  string `xml:"comment,attr"`
	Bitwidth string `xml:"bitwidth,attr"`
	Api      string `xml:"api,attr"`
	Enums    []Enum `xml:"enum"`
}

// Enum represents an <enum> element.
type Enum struct {
	Name       string `xml:"name,attr"`
	Value      string `xml:"value,attr"`
	Bitpos     string `xml:"bitpos,attr"`
	Extends    string `xml:"extends,attr"`
	ExtNumber  string `xml:"extnumber,attr"`
	Offset     string `xml:"offset,attr"`
	Dir        string `xml:"dir,attr"`
	Alias      string `xml:"alias,attr"`
	Comment    string `xml:"comment,attr"`
	Protect    string `xml:"protect,attr"`
	Type       string `xml:"type,attr"`
	Api        string `xml:"api,attr"`
	Deprecated string `xml:"deprecated,attr"`
}

// Commands represents the <commands> element.
type Commands struct {
	Commands []Command `xml:"command"`
}

// Command represents a <command> element.
type Command struct {
	Name         string   `xml:"name,attr"`
	Alias        string   `xml:"alias,attr"`
	Api          string   `xml:"api,attr"`
	SuccessCodes string   `xml:"successcodes,attr"`
	ErrorCodes   string   `xml:"errorcodes,attr"`
	Queues       string   `xml:"queues,attr"`
	RenderPass   string   `xml:"renderpass,attr"`
	CmdBufferLvl string   `xml:"cmdbufferlevel,attr"`
	Tasks        string   `xml:"tasks,attr"`
	Comment      string   `xml:"comment,attr"`
	Proto        Proto    `xml:"proto"`
	Params       []Param  `xml:"param"`
	ImplicitExt  string   `xml:"implicitexternsyncparams,attr"`
}

// Proto represents the <proto> element of a <command>.
type Proto struct {
	Type     string `xml:"type"`
	Name     string `xml:"name"`
	InnerXML string `xml:",innerxml"`
}

// Param represents a <param> element of a <command>.
type Param struct {
	Len                   string `xml:"len,attr"`
	AltLen                string `xml:"altlen,attr"`
	Optional              string `xml:"optional,attr"`
	ExternSync            string `xml:"externsync,attr"`
	NoAutoValidity        string `xml:"noautovalidity,attr"`
	ValidExtensionStructs string `xml:"validextensionstructs,attr"`
	Stride                string `xml:"stride,attr"`
	Api                   string `xml:"api,attr"`
	Type                  string `xml:"type"`
	Name                  string `xml:"name"`
	InnerXML              string `xml:",innerxml"`
}

// Feature represents a <feature> element.
type Feature struct {
	Api      string    `xml:"api,attr"`
	Name     string    `xml:"name,attr"`
	Number   string    `xml:"number,attr"`
	Comment  string    `xml:"comment,attr"`
	Depends  string    `xml:"depends,attr"`
	Requires []Require `xml:"require"`
	Removes  []Remove  `xml:"remove"`
}

// Require represents a <require> element inside <feature> or <extension>.
type Require struct {
	Comment   string       `xml:"comment,attr"`
	Depends   string       `xml:"depends,attr"`
	Api       string       `xml:"api,attr"`
	Types     []ReqType    `xml:"type"`
	Enums     []Enum       `xml:"enum"`
	Commands  []ReqCommand `xml:"command"`
}

// Remove represents a <remove> element inside <feature> or <extension>.
type Remove struct {
	Comment  string       `xml:"comment,attr"`
	Types    []ReqType    `xml:"type"`
	Enums    []Enum       `xml:"enum"`
	Commands []ReqCommand `xml:"command"`
}

// ReqType represents a <type> inside <require> or <remove>.
type ReqType struct {
	Name    string `xml:"name,attr"`
	Api     string `xml:"api,attr"`
	Comment string `xml:"comment,attr"`
}

// ReqCommand represents a <command> inside <require> or <remove>.
type ReqCommand struct {
	Name    string `xml:"name,attr"`
	Api     string `xml:"api,attr"`
	Comment string `xml:"comment,attr"`
}

// Extensions is a list of <extension> elements.
type Extensions struct {
	Extensions []Extension `xml:"extension"`
}

// Extension represents an <extension> element.
type Extension struct {
	Name         string    `xml:"name,attr"`
	Number       string    `xml:"number,attr"`
	Type         string    `xml:"type,attr"`
	Author       string    `xml:"author,attr"`
	Contact      string    `xml:"contact,attr"`
	Supported    string    `xml:"supported,attr"`
	PromotedTo   string    `xml:"promotedto,attr"`
	DeprecatedBy string    `xml:"deprecatedby,attr"`
	ObsoletedBy  string    `xml:"obsoletedby,attr"`
	Depends      string    `xml:"depends,attr"`
	Platform     string    `xml:"platform,attr"`
	SpecialUse   string    `xml:"specialuse,attr"`
	Ratified     string    `xml:"ratified,attr"`
	Comment      string    `xml:"comment,attr"`
	Requires     []Require `xml:"require"`
	Removes      []Remove  `xml:"remove"`
}

// Formats represents the <formats> element.
type Formats struct {
	Formats []Format `xml:"format"`
}

// Format represents a <format> element.
type Format struct {
	Name             string      `xml:"name,attr"`
	Class            string      `xml:"class,attr"`
	BlockSize        string      `xml:"blockSize,attr"`
	TexelsPerBlock   string      `xml:"texelsPerBlock,attr"`
	Packed           string      `xml:"packed,attr"`
	Chroma           string      `xml:"chroma,attr"`
	Components       []Component `xml:"component"`
	Planes           []Plane     `xml:"plane"`
	SpirvImageFormat string      `xml:"spirvimageformat"`
}

// Component represents a <component> element inside <format>.
type Component struct {
	Name          string `xml:"name,attr"`
	Bits          string `xml:"bits,attr"`
	NumericFormat string `xml:"numericFormat,attr"`
}

// Plane represents a <plane> element inside <format>.
type Plane struct {
	Index        string `xml:"index,attr"`
	WidthDivisor string `xml:"widthDivisor,attr"`
	HeightDivisor string `xml:"heightDivisor,attr"`
	Compatible   string `xml:"compatible,attr"`
}

// Sync represents the <sync> element.
type Sync struct {
	SyncStages     []SyncStage     `xml:"syncstage"`
	SyncAccesses   []SyncAccess    `xml:"syncaccess"`
	SyncPipelines  []SyncPipeline  `xml:"syncpipeline"`
}

// SyncStage represents a <syncstage> element.
type SyncStage struct {
	Name    string `xml:"name,attr"`
	Alias   string `xml:"alias,attr"`
	Support string `xml:"support,attr"`
}

// SyncAccess represents a <syncaccess> element.
type SyncAccess struct {
	Name    string `xml:"name,attr"`
	Alias   string `xml:"alias,attr"`
	Support string `xml:"support,attr"`
}

// SyncPipeline represents a <syncpipeline> element.
type SyncPipeline struct {
	Name   string      `xml:"name,attr"`
	Stages []SyncStage `xml:"syncstage"`
}

// SpirvExtensions represents the <spirvextensions> element.
type SpirvExtensions struct {
	Extensions []SpirvExtension `xml:"spirvextension"`
}

// SpirvExtension represents a <spirvextension> element.
type SpirvExtension struct {
	Name    string        `xml:"name,attr"`
	Enables []SpirvEnable `xml:"enable"`
}

// SpirvCapabilities represents the <spirvcapabilities> element.
type SpirvCapabilities struct {
	Capabilities []SpirvCapability `xml:"spirvcapability"`
}

// SpirvCapability represents a <spirvcapability> element.
type SpirvCapability struct {
	Name    string        `xml:"name,attr"`
	Enables []SpirvEnable `xml:"enable"`
}

// SpirvEnable represents an <enable> element inside spirvextension / spirvcapability.
type SpirvEnable struct {
	Version   string `xml:"version,attr"`
	Extension string `xml:"extension,attr"`
	Struct    string `xml:"struct,attr"`
	Feature   string `xml:"feature,attr"`
	Property  string `xml:"property,attr"`
	Member    string `xml:"member,attr"`
	Value     string `xml:"value,attr"`
}
