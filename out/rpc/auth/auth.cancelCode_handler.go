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

// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (s *AuthServiceImpl) AuthCancelCode(ctx context.Context, request *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthCancelCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AuthCancelCode logic

	return nil, fmt.Errorf("Not impl AuthCancelCode")
}
