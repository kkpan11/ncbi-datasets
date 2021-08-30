// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: ncbi/datasets/v1alpha1/reports/fasta.proto

package reports

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Rendering will be snake-case of the term.
type FastaFieldNames int32

const (
	FastaFieldNames_ORGANISM FastaFieldNames = 0
	// Unique to pathogen defline
	FastaFieldNames_ELEMENT_RANGE  FastaFieldNames = 1
	FastaFieldNames_ELEMENT_NAME   FastaFieldNames = 2
	FastaFieldNames_ELEMENT_SYMBOL FastaFieldNames = 3
	FastaFieldNames_CONTIG         FastaFieldNames = 4
	// Unique to Prokaryote
	FastaFieldNames_GENE              FastaFieldNames = 5
	FastaFieldNames_PROTEIN_ACCESSION FastaFieldNames = 6
	FastaFieldNames_CHROMOSOME        FastaFieldNames = 7
	FastaFieldNames_NAME              FastaFieldNames = 8
	FastaFieldNames_COMPLETENESS      FastaFieldNames = 9
)

// Enum value maps for FastaFieldNames.
var (
	FastaFieldNames_name = map[int32]string{
		0: "ORGANISM",
		1: "ELEMENT_RANGE",
		2: "ELEMENT_NAME",
		3: "ELEMENT_SYMBOL",
		4: "CONTIG",
		5: "GENE",
		6: "PROTEIN_ACCESSION",
		7: "CHROMOSOME",
		8: "NAME",
		9: "COMPLETENESS",
	}
	FastaFieldNames_value = map[string]int32{
		"ORGANISM":          0,
		"ELEMENT_RANGE":     1,
		"ELEMENT_NAME":      2,
		"ELEMENT_SYMBOL":    3,
		"CONTIG":            4,
		"GENE":              5,
		"PROTEIN_ACCESSION": 6,
		"CHROMOSOME":        7,
		"NAME":              8,
		"COMPLETENESS":      9,
	}
)

func (x FastaFieldNames) Enum() *FastaFieldNames {
	p := new(FastaFieldNames)
	*p = x
	return p
}

func (x FastaFieldNames) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FastaFieldNames) Descriptor() protoreflect.EnumDescriptor {
	return file_ncbi_datasets_v1alpha1_reports_fasta_proto_enumTypes[0].Descriptor()
}

func (FastaFieldNames) Type() protoreflect.EnumType {
	return &file_ncbi_datasets_v1alpha1_reports_fasta_proto_enumTypes[0]
}

func (x FastaFieldNames) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FastaFieldNames.Descriptor instead.
func (FastaFieldNames) EnumDescriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDescGZIP(), []int{0}
}

var File_ncbi_datasets_v1alpha1_reports_fasta_proto protoreflect.FileDescriptor

var file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2f, 0x66, 0x61, 0x73, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6e, 0x63,
	0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2a, 0xb1, 0x01, 0x0a,
	0x0f, 0x46, 0x61, 0x73, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x53, 0x4d, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4e, 0x54, 0x49,
	0x47, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x45, 0x4e, 0x45, 0x10, 0x05, 0x12, 0x15, 0x0a,
	0x11, 0x50, 0x52, 0x4f, 0x54, 0x45, 0x49, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x52, 0x4f, 0x4d, 0x4f, 0x53, 0x4f,
	0x4d, 0x45, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x08, 0x12, 0x10,
	0x0a, 0x0c, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x4e, 0x45, 0x53, 0x53, 0x10, 0x09,
	0x42, 0x23, 0x5a, 0x1e, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDescOnce sync.Once
	file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDescData = file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDesc
)

func file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDescGZIP() []byte {
	file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDescOnce.Do(func() {
		file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDescData = protoimpl.X.CompressGZIP(file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDescData)
	})
	return file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDescData
}

var file_ncbi_datasets_v1alpha1_reports_fasta_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ncbi_datasets_v1alpha1_reports_fasta_proto_goTypes = []interface{}{
	(FastaFieldNames)(0), // 0: ncbi.datasets.v1alpha1.reports.FastaFieldNames
}
var file_ncbi_datasets_v1alpha1_reports_fasta_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ncbi_datasets_v1alpha1_reports_fasta_proto_init() }
func file_ncbi_datasets_v1alpha1_reports_fasta_proto_init() {
	if File_ncbi_datasets_v1alpha1_reports_fasta_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ncbi_datasets_v1alpha1_reports_fasta_proto_goTypes,
		DependencyIndexes: file_ncbi_datasets_v1alpha1_reports_fasta_proto_depIdxs,
		EnumInfos:         file_ncbi_datasets_v1alpha1_reports_fasta_proto_enumTypes,
	}.Build()
	File_ncbi_datasets_v1alpha1_reports_fasta_proto = out.File
	file_ncbi_datasets_v1alpha1_reports_fasta_proto_rawDesc = nil
	file_ncbi_datasets_v1alpha1_reports_fasta_proto_goTypes = nil
	file_ncbi_datasets_v1alpha1_reports_fasta_proto_depIdxs = nil
}