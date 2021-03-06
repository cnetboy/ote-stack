/*
Copyright 2019 Baidu, Inc.

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

package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/baidu/ote-stack/pkg/clustermessage"
)

func TestResponse (t *testing.T) {
	head := &clustermessage.MessageHead{
		Command:	clustermessage.CommandType_ControlReq,
	}
	resp := Response(nil, head)
	assert.NotNil(t, resp)
}

func TestControlTaskResponse(t *testing.T) {
	resp := ControlTaskResponse(200, "")
	assert.NotNil(t, resp)
}

func TestGetControllerTask(t *testing.T) {
	ret := GetControllerTaskFromClusterMessage(nil)
	assert.Nil(t, ret)

	msg := &clustermessage.ClusterMessage{
		Body:	[]byte{1},
	}
	ret = GetControllerTaskFromClusterMessage(msg)
	assert.Nil(t, ret)
}

func TestGetControlMultiTask(t *testing.T) {
	ret := GetControlMultiTaskFromClusterMessage(nil)
	assert.Nil(t, ret)

	msg := &clustermessage.ClusterMessage{
		Body:	[]byte{1},
	}
	ret = GetControlMultiTaskFromClusterMessage(msg)
	assert.Nil(t, ret)
}