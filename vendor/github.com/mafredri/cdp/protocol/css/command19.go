// +build go1.9

// Code generated by cdpgen. DO NOT EDIT.

package css

import (
	"github.com/mafredri/cdp/protocol/page"
)

// CreateStyleSheetArgs represents the arguments for CreateStyleSheet in the CSS domain.
type CreateStyleSheetArgs struct {
	FrameID page.FrameID `json:"frameId"` // Identifier of the frame where "via-inspector" stylesheet should be created.
}

// NewCreateStyleSheetArgs initializes CreateStyleSheetArgs with the required arguments.
func NewCreateStyleSheetArgs(frameID page.FrameID) *CreateStyleSheetArgs {
	args := new(CreateStyleSheetArgs)
	args.FrameID = frameID
	return args
}