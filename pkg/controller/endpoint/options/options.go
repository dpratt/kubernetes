/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"github.com/spf13/pflag"
)

type EndpointControllerOptions struct {
	ConcurrentEndpointSyncs int
}

func NewEndpointControllerOptions() EndpointControllerOptions {
	return EndpointControllerOptions{
		ConcurrentEndpointSyncs: 5,
	}
}

func (o *EndpointControllerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&o.ConcurrentEndpointSyncs, "concurrent-endpoint-syncs", o.ConcurrentEndpointSyncs,
		"The number of endpoint syncing operations that will be done concurrently. Larger number = faster endpoint updating, but more CPU (and network) load")
}
