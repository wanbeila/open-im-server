// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package push

import "github.com/OpenIMSDK/Open-IM-Server/pkg/errs"

func (x *PushMsgReq) Check() error {
	if x.MsgData == nil {
		return errs.ErrArgs.Wrap("MsgData is empty")
	}
	if err := x.MsgData.Check(); err != nil {
		return err
	}
	if x.ConversationID == "" {
		return errs.ErrArgs.Wrap("ConversationID is empty")
	}
	return nil
}

func (x *DelUserPushTokenReq) Check() error {
	if x.UserID == "" {
		return errs.ErrArgs.Wrap("UserID is empty")
	}
	if x.PlatformID < 1 || x.PlatformID > 9 {
		return errs.ErrArgs.Wrap("PlatformID is invalid")
	}
	return nil
}