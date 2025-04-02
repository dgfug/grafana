//
// This file is generated by grafana-app-sdk
// DO NOT EDIT
//

package apis

import (
	"encoding/json"

	"github.com/grafana/grafana-app-sdk/app"
)

var (
	rawSchemaInvestigationv0alpha1          = []byte(`{"spec":{"description":"spec is the schema of our resource","properties":{"collectables":{"items":{"properties":{"createdAt":{"type":"string"},"datasource":{"properties":{"uid":{"type":"string"}},"required":["uid"],"type":"object"},"fieldConfig":{"type":"string"},"id":{"type":"string"},"logoPath":{"type":"string"},"note":{"type":"string"},"noteUpdatedAt":{"type":"string"},"origin":{"type":"string"},"queries":{"items":{"type":"string"},"type":"array"},"timeRange":{"properties":{"from":{"type":"string"},"raw":{"properties":{"from":{"type":"string"},"to":{"type":"string"}},"required":["from","to"],"type":"object"},"to":{"type":"string"}},"required":["from","to","raw"],"type":"object"},"title":{"type":"string"},"type":{"type":"string"},"url":{"type":"string"}},"required":["id","createdAt","title","origin","type","queries","timeRange","datasource","url","note","noteUpdatedAt","fieldConfig"],"type":"object"},"type":"array"},"createdByProfile":{"properties":{"gravatarUrl":{"type":"string"},"name":{"type":"string"},"uid":{"type":"string"}},"required":["uid","name","gravatarUrl"],"type":"object"},"hasCustomName":{"type":"boolean"},"isFavorite":{"type":"boolean"},"overviewNote":{"type":"string"},"overviewNoteUpdatedAt":{"type":"string"},"title":{"type":"string"},"viewMode":{"properties":{"mode":{"enum":["compact","full"],"type":"string"},"showComments":{"type":"boolean"},"showTooltips":{"type":"boolean"}},"required":["mode","showComments","showTooltips"],"type":"object"}},"required":["title","createdByProfile","hasCustomName","isFavorite","overviewNote","overviewNoteUpdatedAt","collectables","viewMode"],"type":"object"},"status":{"properties":{"additionalFields":{"description":"additionalFields is reserved for future use","type":"object","x-kubernetes-preserve-unknown-fields":true},"operatorStates":{"additionalProperties":{"properties":{"descriptiveState":{"description":"descriptiveState is an optional more descriptive state field which has no requirements on format","type":"string"},"details":{"description":"details contains any extra information that is operator-specific","type":"object","x-kubernetes-preserve-unknown-fields":true},"lastEvaluation":{"description":"lastEvaluation is the ResourceVersion last evaluated","type":"string"},"state":{"description":"state describes the state of the lastEvaluation.\nIt is limited to three possible states for machine evaluation.","enum":["success","in_progress","failed"],"type":"string"}},"required":["lastEvaluation","state"],"type":"object"},"description":"operatorStates is a map of operator ID to operator state evaluations.\nAny operator which consumes this kind SHOULD add its state evaluation information to this field.","type":"object"}},"type":"object"}}`)
	versionSchemaInvestigationv0alpha1      app.VersionSchema
	_                                       = json.Unmarshal(rawSchemaInvestigationv0alpha1, &versionSchemaInvestigationv0alpha1)
	rawSchemaInvestigationIndexv0alpha1     = []byte(`{"spec":{"properties":{"investigationSummaries":{"description":"Array of investigation summaries","items":{"properties":{"collectableSummaries":{"items":{"properties":{"id":{"type":"string"},"logoPath":{"type":"string"},"origin":{"type":"string"},"title":{"type":"string"}},"required":["id","title","logoPath","origin"],"type":"object"},"type":"array"},"createdByProfile":{"properties":{"gravatarUrl":{"type":"string"},"name":{"type":"string"},"uid":{"type":"string"}},"required":["uid","name","gravatarUrl"],"type":"object"},"hasCustomName":{"type":"boolean"},"isFavorite":{"type":"boolean"},"overviewNote":{"type":"string"},"overviewNoteUpdatedAt":{"type":"string"},"title":{"type":"string"},"viewMode":{"properties":{"mode":{"enum":["compact","full"],"type":"string"},"showComments":{"type":"boolean"},"showTooltips":{"type":"boolean"}},"required":["mode","showComments","showTooltips"],"type":"object"}},"required":["title","createdByProfile","hasCustomName","isFavorite","overviewNote","overviewNoteUpdatedAt","viewMode","collectableSummaries"],"type":"object"},"type":"array"},"owner":{"description":"The Person who owns this investigation index","properties":{"gravatarUrl":{"type":"string"},"name":{"type":"string"},"uid":{"type":"string"}},"required":["uid","name","gravatarUrl"],"type":"object"},"title":{"description":"Title of the index, e.g. 'Favorites' or 'My Investigations'","type":"string"}},"required":["title","owner","investigationSummaries"],"type":"object"},"status":{"properties":{"additionalFields":{"description":"additionalFields is reserved for future use","type":"object","x-kubernetes-preserve-unknown-fields":true},"operatorStates":{"additionalProperties":{"properties":{"descriptiveState":{"description":"descriptiveState is an optional more descriptive state field which has no requirements on format","type":"string"},"details":{"description":"details contains any extra information that is operator-specific","type":"object","x-kubernetes-preserve-unknown-fields":true},"lastEvaluation":{"description":"lastEvaluation is the ResourceVersion last evaluated","type":"string"},"state":{"description":"state describes the state of the lastEvaluation.\nIt is limited to three possible states for machine evaluation.","enum":["success","in_progress","failed"],"type":"string"}},"required":["lastEvaluation","state"],"type":"object"},"description":"operatorStates is a map of operator ID to operator state evaluations.\nAny operator which consumes this kind SHOULD add its state evaluation information to this field.","type":"object"}},"type":"object"}}`)
	versionSchemaInvestigationIndexv0alpha1 app.VersionSchema
	_                                       = json.Unmarshal(rawSchemaInvestigationIndexv0alpha1, &versionSchemaInvestigationIndexv0alpha1)
)

var appManifestData = app.ManifestData{
	AppName: "investigations",
	Group:   "investigations.grafana.app",
	Kinds: []app.ManifestKind{
		{
			Kind:       "Investigation",
			Scope:      "Namespaced",
			Conversion: false,
			Versions: []app.ManifestKindVersion{
				{
					Name:   "v0alpha1",
					Schema: &versionSchemaInvestigationv0alpha1,
				},
			},
		},

		{
			Kind:       "InvestigationIndex",
			Scope:      "Namespaced",
			Conversion: false,
			Versions: []app.ManifestKindVersion{
				{
					Name:   "v0alpha1",
					Schema: &versionSchemaInvestigationIndexv0alpha1,
				},
			},
		},
	},
}

func jsonToMap(j string) map[string]any {
	m := make(map[string]any)
	json.Unmarshal([]byte(j), &j)
	return m
}

func LocalManifest() app.Manifest {
	return app.NewEmbeddedManifest(appManifestData)
}

func RemoteManifest() app.Manifest {
	return app.NewAPIServerManifest("investigations")
}
