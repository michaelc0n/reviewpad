// Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package plugins_aladino_functions

import (
	"fmt"
	"strconv"

	"github.com/reviewpad/reviewpad/v3/handler"
	"github.com/reviewpad/reviewpad/v3/lang/aladino"
)

func ToNumber() *aladino.BuiltInFunction {
	return &aladino.BuiltInFunction{
		Type:           aladino.BuildFunctionType([]aladino.Type{aladino.BuildStringType()}, aladino.BuildIntType()),
		Code:           toNumberCode,
		SupportedKinds: []handler.TargetEntityKind{handler.PullRequest, handler.Issue},
	}
}

func toNumberCode(e aladino.Env, args []aladino.Value) (aladino.Value, error) {
	str := args[0].(*aladino.StringValue).Val

	num, err := strconv.Atoi(str)
	if err != nil {
		return nil, fmt.Errorf(`error converting "%s" to number: %w`, str, err)
	}

	return aladino.BuildIntValue(num), nil
}
