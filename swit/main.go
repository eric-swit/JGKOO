package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	Hook             = "hook"
	ApiV3            = "swit-api-golang"
	ApiV1            = "swit-apiv1"
	ApiV5            = "swit-apiv5-golang"
	Badge            = "swit-badge-golang"
	Storage          = "swit-gcs-file-golang"
	Activity         = "swit-grpc-activity-golang"
	Alltask          = "swit-grpc-alltask-golang"
	Approval         = "swit-grpc-approval-golang"
	Apps             = "swit-grpc-apps-golang"
	AppVersion       = "swit-grpc-appVersion-golang"
	Asset            = "swit-grpc-asset-golang"
	Auth             = "swit-grpc-auth-golang"
	Board            = "swit-grpc-board-golang"
	Channel          = "swit-grpc-channel-golang"
	Company          = "swit-grpc-company-golang"
	Contents         = "swit-grpc-contents-golang"
	Coupon           = "swit-grpc-coupon-golang"
	Elastic          = "swit-grpc-elastic"
	Email            = "swit-grpc-email"
	Emoji            = "swit-grpc-emoji"
	EventSub         = "swit-grpc-eventsub-golang"
	Export           = "swit-grpc-export-golang"
	Import           = "swit-grpc-import"
	Mention          = "swit-grpc-mention-golang"
	Message          = "swit-grpc-message-golang"
	Pay              = "swit-grpc-pay-golang"
	Project          = "swit-grpc-project-golang"
	Sfdc             = "swit-grpc-sfdc-golang"
	Stats            = "swit-grpc-stats"
	SwitBot          = "swit-grpc-switbot-golang"
	Task             = "swit-grpc-task-golang"
	User             = "swit-grpc-user-golang"
	Workspace        = "swit-grpc-workspace-golang"
	HttpOauth2       = "swit-http-oauth2-golang"
	ImportSql        = "swit-import-sql"
	InviteServer     = "swit-invite-server"
	MentionSubscribe = "swit-mention-subscribe-nodejs"
	Notification     = "swit-notification"
	Openapi          = "swit-openapi-golang"
	Saml             = "swit-saml-sp"
	Socket           = "swit-socket-api"
	SubscribeAgent   = "swit-subscribe-agent-golang"
	SupportApi       = "swit-support-api-golang"
	SupportFront     = "swit-support-front-angular"
	SupportGrpc      = "swit-support-grpc-golang"
	Batch            = "swit-batch"
	Scheduler        = "swit-scheduler-api-golang"

	V1 = "v1"
	V2 = "v2"
	V3 = "v3"
	V4 = "v4"
	V5 = "v5"
)

type Data struct {
	Distribution []*Distribution
}

type Distribution struct {
	Path    string
	Version string
}

var basePath = "$HOME/go/src/swit/"

// release-v1.0.0
// release-v4.0.0
// release-v5.0.0

func (a *Data) create() error {
	var str, branch string
	for _, t := range a.Distribution {
		switch t.Version {
		case V4:
			branch = "release.v4"
		case V5:
			branch = "release.v5"
		default:
			branch = "master"
		}

		str += fmt.Sprintf("cd %s \n", basePath+t.Path)
		//str += fmt.Sprintf("git checkout %s \n", branch)
		//str += fmt.Sprintf("git pull \n")

		cmd := exec.Command("git", "checkout", branch)
		cmd.Dir = "/Users/erickoo/go/src/swit/" + t.Path
		cmd.Run()

		cmd = exec.Command("git", "pull", branch)
		cmd.Dir = "/Users/erickoo/go/src/swit/" + t.Path
		cmd.Run()

		cmd = exec.Command("git", "describe", "--tags", "--abbrev=0")
		cmd.Dir = "/Users/erickoo/go/src/swit/" + t.Path
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		tag := string(out)

		str += fmt.Sprintf("git tag %s \n", tag)
		str += fmt.Sprintf("git push origin %s \n \n", tag)
	}

	b := []byte(str)
	r := time.Now().Format("2006-01-02")
	f1, err := os.Create("prod_" + r + ".sh")
	checkError(err)
	defer f1.Close()
	fmt.Fprintf(f1, string(b))

	return nil
}

func (a *Data) merge() error {
	var str, branch, mergeBranch string
	for _, t := range a.Distribution {
		switch t.Version {
		case V4:
			branch = "release.v4"
			mergeBranch = "express.v4"
		case V5:
			branch = "release.v5"
			mergeBranch = "express.v5"
		default:
			branch = "master"
			mergeBranch = "express"
		}

		str += fmt.Sprintf("cd %s \n", basePath+t.Path)
		//str += fmt.Sprintf("git checkout %s \n", branch)
		//str += fmt.Sprintf("git pull \n")
		str += fmt.Sprintf("git merge origin/%s \n", mergeBranch)
		str += fmt.Sprintf("git push origin %s \n \n", branch)
	}

	b := []byte(str)
	r := time.Now().Format("2006-01-02")
	f1, err := os.Create("prod_merge_" + r + ".sh")
	checkError(err)
	defer f1.Close()
	fmt.Fprintf(f1, string(b))

	return nil
}

func New(v []*Distribution) *Data {
	return &Data{
		Distribution: v,
	}
}

func main() {
	var s []*Distribution

	s = append(s, &Distribution{Path: Message, Version: V5})
	s = append(s, &Distribution{Path: Task, Version: V5})
	s = append(s, &Distribution{Path: ApiV1, Version: V1})
	s = append(s, &Distribution{Path: ApiV5, Version: V1})
	s = append(s, &Distribution{Path: Channel, Version: V4})
	s = append(s, &Distribution{Path: Project, Version: V4})
	s = append(s, &Distribution{Path: Contents, Version: V1})
	s = append(s, &Distribution{Path: Asset, Version: V1})
	s = append(s, &Distribution{Path: Storage, Version: V1})
	s = append(s, &Distribution{Path: Channel, Version: V5})
	s = append(s, &Distribution{Path: Project, Version: V5})
	s = append(s, &Distribution{Path: Openapi, Version: V1})

	d := New(s)
	d.merge()
	d.create()
}

// 마지막 태그명 복붙
// git describe --tags $(git rev-list --tags --max-count=1) | pbcopy
// pbpaste
