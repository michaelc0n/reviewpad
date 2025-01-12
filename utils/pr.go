// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package utils

import (
	"strings"

	"github.com/reviewpad/reviewpad/v3/handler"
)

func IsPullRequestReadyForReportMetrics(eventData *handler.EventData) bool {
	return eventData != nil && eventData.EventName == "pull_request" && eventData.EventAction == "closed"
}

func IsReviewpadCommand(eventData *handler.EventData) bool {
	return eventData != nil &&
		eventData.EventName == "issue_comment" &&
		eventData.EventAction == "created" &&
		eventData.Comment.Body != nil &&
		strings.HasPrefix(*eventData.Comment.Body, "/reviewpad")
}
