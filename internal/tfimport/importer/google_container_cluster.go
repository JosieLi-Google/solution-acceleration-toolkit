/*
 * Copyright 2020 Google LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package importer

import (
	"fmt"

	"github.com/GoogleCloudPlatform/healthcare-data-protection-suite/internal/terraform"
)

// GKECluster defines a struct with the necessary information for a GKE cluster to be imported.
type GKECluster struct{}

// ImportID returns the GCP project and bucket name for use in importing.
func (b *GKECluster) ImportID(rc terraform.ResourceChange, pcv ProviderConfigMap) (string, error) {
	project := fromConfigValues("project", rc.Change.After, pcv)
	if project == nil {
		return "", fmt.Errorf("could not find project in resource change or provider config")
	}

	location := fromConfigValues("location", rc.Change.After, pcv)
	if project == nil {
		return "", fmt.Errorf("could not find location in resource change or provider config")
	}

	name := fromConfigValues("name", rc.Change.After, nil)
	if name == nil {
		return "", fmt.Errorf("could not find cluster name in resource change or provider config")
	}

	return fmt.Sprintf("projects/%v/locations/%v/clusters/%v", project, location, name), nil
}
