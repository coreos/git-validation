package subsystem

import (
	"regexp"
	"strings"

	"github.com/vbatts/git-validation/git"
	"github.com/vbatts/git-validation/validate"
)

var (
	subsystemPattern = regexp.MustCompile("^.+: .+$")

	// SubsystemRule is the rule being registered
	SubsystemRule = validate.Rule{
		Name:        "subsystem-in-subject",
		Description: "commit subjects are of the form '<subsystem>: <commit-msg>'",
		Run:         ValidateSubsystem,
	}
)

func init() {
	validate.RegisterRule(SubsystemRule)
}

// ValidateSubsystem checks that the commit's subject is of the form '<subsystem>: <commit-msg>'.
func ValidateSubsystem(c git.CommitEntry) (vr validate.Result) {
	vr.CommitEntry = c
	if len(strings.Split(c["parent"], " ")) > 1 {
		vr.Pass = true
		vr.Msg = "merge commits do not require subsystem in subject"
		return vr
	}

	if ok := subsystemPattern.MatchString(c["subject"]); !ok {
		vr.Pass = false
		vr.Msg = "commit subject should be of the form '<subsystem>: <commit-msg>'"

		return
	}

	vr.Pass = true
	vr.Msg = "commit subject is of the form '<subsystem>: <commit-msg>'"

	return
}
