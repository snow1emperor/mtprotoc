/*
 *  Copyright (c) 2017, https://github.com/snow1emperor
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/snow1emperor/telegramd/baselib/logger"
	"github.com/snow1emperor/telegramd/grpc_util"
	"github.com/snow1emperor/telegramd/mtproto"
	"golang.org/x/net/context"
)

// channels.getFullChannel#8736a09 channel:InputChannel = messages.ChatFull;
func (s *ChannelsServiceImpl) ChannelsGetFullChannel(ctx context.Context, request *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ChannelsGetFullChannel - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl ChannelsGetFullChannel logic

	return nil, fmt.Errorf("Not impl ChannelsGetFullChannel")
}
