/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package exec

import (
	"context"
	"fmt"
	"path"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

type ShowActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewShowActionSpec() spec.ExpActionCommandSpec {
	return &ShowActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "match",
					Desc: "test match",
				},
			},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:   "info",
					Desc:   "show info",
					NoArgs: false,
				},
			},
			ActionExecutor: &ShowActionExecutor{},
		},
	}
}

func (*ShowActionSpec) Name() string {
	return "fill"
}

func (*ShowActionSpec) Aliases() []string {
	return []string{}
}

func (*ShowActionSpec) ShortDesc() string {
	return "Show the specified directory path"
}

func (*ShowActionSpec) LongDesc() string {
	return "Show the specified directory path. If the path is not directory or does not exist, an error message will be returned."
}

type ShowActionExecutor struct {
	channel spec.Channel
}

func (*ShowActionExecutor) Name() string {
	return "fill"
}

var testshowBin = "chaos_testshow"

func (fae *ShowActionExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	if fae.channel == nil {
		return spec.ReturnFail(spec.Code[spec.ServerError], "channel is nil")
	}
	info := model.ActionFlags["info"]
	if info != "" {
		return fae.start(info, ctx)
	}
	if _, ok := spec.IsDestroy(ctx); ok {
                return fae.stop(ctx)
        } else {
                return fae.start(info, ctx)
        }	
}

func (fae *ShowActionExecutor) start(info string, ctx context.Context) *spec.Response {
	flags := fmt.Sprintf("--start --debug=%t --info=%s", util.Debug, info)
	return fae.channel.Run(ctx, path.Join(fae.channel.GetScriptPath(), testshowBin), flags)
}

func (fae *ShowActionExecutor) stop(ctx context.Context) *spec.Response {
	return fae.channel.Run(ctx, path.Join(fae.channel.GetScriptPath(), testshowBin),
		fmt.Sprintf("--stop --debug=%t", util.Debug))
}

func (fae *ShowActionExecutor) SetChannel(channel spec.Channel) {
	fae.channel = channel
}
