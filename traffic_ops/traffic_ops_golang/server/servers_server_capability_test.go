package server

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"testing"

	"github.com/apache/trafficcontrol/traffic_ops/traffic_ops_golang/api"
)

func TestInterfaces(t *testing.T) {
	var i interface{}
	i = &TOServerServerCapability{}

	if _, ok := i.(api.Creator); !ok {
		t.Errorf("ServerServerCapability must be Creator")
	}
	if _, ok := i.(api.Reader); !ok {
		t.Errorf("ServerServerCapability must be Reader")
	}
	if _, ok := i.(api.Deleter); !ok {
		t.Errorf("ServerServerCapability must be Deleter")
	}
	if _, ok := i.(api.Identifier); !ok {
		t.Errorf("ServerServerCapability must be Identifier")
	}
}
