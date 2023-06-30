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

package user

import (
	"context"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/errs"
	pbuser "github.com/OpenIMSDK/Open-IM-Server/pkg/proto/user"
	"time"
)

func (s *userServer) UserRegisterCount(ctx context.Context, req *pbuser.UserRegisterCountReq) (*pbuser.UserRegisterCountResp, error) {
	if req.Start > req.End {
		return nil, errs.ErrArgs.Wrap("start > end")
	}
	total, err := s.CountTotal(ctx)
	if err != nil {
		return nil, err
	}
	count, err := s.CountRangeEverydayTotal(ctx, time.UnixMilli(req.Start), time.UnixMilli(req.End))
	if err != nil {
		return nil, err
	}
	return &pbuser.UserRegisterCountResp{Total: total, Count: count}, nil
}