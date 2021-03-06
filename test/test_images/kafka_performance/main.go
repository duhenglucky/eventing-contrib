/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"

	"knative.dev/eventing/test/performance/infra"
)

var minWorkers uint64
var bootstrapUrl string
var topic string

func init() {
	infra.DeclareFlags()

	// Specific to http load generator
	flag.Uint64Var(&minWorkers, "min-workers", 10, "Number of vegeta workers")
	flag.StringVar(&bootstrapUrl, "bootstrap-url", "", "The bootstrap url for the event destination.")
	flag.StringVar(&topic, "topic", "", "The topic for the event destination.")
}

func main() {
	flag.Parse()

	infra.StartPerformanceImage(NewKafkaLoadGeneratorFactory(bootstrapUrl, topic, minWorkers), JsonTypeExtractor, JsonIdExtractor)
}
