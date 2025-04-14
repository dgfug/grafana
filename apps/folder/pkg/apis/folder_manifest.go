//
// This file is generated by grafana-app-sdk
// DO NOT EDIT
//

package apis

import (
	"encoding/json"

	"github.com/grafana/grafana-app-sdk/app"
)

var appManifestData = app.ManifestData{
	AppName: "folder",
	Group:   "folder.grafana.app",
	Kinds: []app.ManifestKind{
		{
			Kind:       "Folder",
			Scope:      "Namespaced",
			Conversion: false,
			Versions: []app.ManifestKindVersion{
				{
					Name: "v1beta1",
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
	return app.NewAPIServerManifest("folder")
}
