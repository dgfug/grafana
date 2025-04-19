//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by openapi-gen. DO NOT EDIT.

package v1beta1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.DescendantCounts": schema_pkg_apis_folder_v1beta1_DescendantCounts(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.Folder":           schema_pkg_apis_folder_v1beta1_Folder(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderAccessInfo": schema_pkg_apis_folder_v1beta1_FolderAccessInfo(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderInfo":       schema_pkg_apis_folder_v1beta1_FolderInfo(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderInfoList":   schema_pkg_apis_folder_v1beta1_FolderInfoList(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderJSONCodec":  schema_pkg_apis_folder_v1beta1_FolderJSONCodec(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderList":       schema_pkg_apis_folder_v1beta1_FolderList(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderMetadata":   schema_pkg_apis_folder_v1beta1_FolderMetadata(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderSpec":       schema_pkg_apis_folder_v1beta1_FolderSpec(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderStatus":     schema_pkg_apis_folder_v1beta1_FolderStatus(ref),
		"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.ResourceStats":    schema_pkg_apis_folder_v1beta1_ResourceStats(ref),
	}
}

func schema_pkg_apis_folder_v1beta1_DescendantCounts(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"counts": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.ResourceStats"),
									},
								},
							},
						},
					},
				},
				Required: []string{"counts"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.ResourceStats"},
	}
}

func schema_pkg_apis_folder_v1beta1_Folder(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Spec is the spec of the Folder",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderStatus"),
						},
					},
				},
				Required: []string{"metadata", "spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderSpec", "github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_folder_v1beta1_FolderAccessInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Access control information for the current user",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"canSave": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canEdit": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canAdmin": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canDelete": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
				},
				Required: []string{"canSave", "canEdit", "canAdmin", "canDelete"},
			},
		},
	}
}

func schema_pkg_apis_folder_v1beta1_FolderInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FolderInfo briefly describes a folder -- unlike a folder resource, this is a partial record of the folder metadata used for navigating parents and children",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the k8s name (eg, the unique identifier) for a folder",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"title": {
						SchemaProps: spec.SchemaProps{
							Description: "Title is the display value",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "The folder description",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"parent": {
						SchemaProps: spec.SchemaProps{
							Description: "The parent folder UID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"detached": {
						SchemaProps: spec.SchemaProps{
							Description: "This folder does not resolve",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "title"},
			},
		},
	}
}

func schema_pkg_apis_folder_v1beta1_FolderInfoList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FolderInfoList returns a list of folder references (parents or children) Unlike FolderList, each item is not a full k8s object",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"uid",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderInfo"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.FolderInfo", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_folder_v1beta1_FolderJSONCodec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FolderJSONCodec is an implementation of resource.Codec for kubernetes JSON encoding",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_folder_v1beta1_FolderList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.Folder"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/folder/pkg/apis/folder/v1beta1.Folder", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_folder_v1beta1_FolderMetadata(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "metadata contains embedded CommonMetadata and can be extended with custom string fields without external reference as using the CommonMetadata reference breaks thema codegen.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"updateTimestamp": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"createdBy": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"uid": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"creationTimestamp": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"deletionTimestamp": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"finalizers": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"resourceVersion": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"generation": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int64",
						},
					},
					"updatedBy": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"updateTimestamp", "createdBy", "uid", "creationTimestamp", "finalizers", "resourceVersion", "generation", "updatedBy", "labels"},
			},
		},
	}
}

func schema_pkg_apis_folder_v1beta1_FolderSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"title": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"title"},
			},
		},
	}
}

func schema_pkg_apis_folder_v1beta1_FolderStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Empty stub",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_folder_v1beta1_ResourceStats(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"count": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int64",
						},
					},
				},
				Required: []string{"group", "resource", "count"},
			},
		},
	}
}
