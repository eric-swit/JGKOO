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

const (
	Hook             = "hook"
	ApiV3            = "switbe-api-golang"
	ApiV1            = "switbe-apiv1"
	ApiV5            = "switbe-apiv5-golang"
	Badge            = "switbe-badge-golang"
	Storage          = "switbe-gcs-file-golang"
	Activity         = "switbe-grpc-activity-golang"
	Alltask          = "switbe-grpc-alltask-golang"
	Approval         = "switbe-grpc-approval-golang"
	Apps             = "switbe-grpc-apps-golang"
	AppVersion       = "switbe-grpc-appVersion-golang"
	Asset            = "switbe-grpc-asset-golang"
	Auth             = "switbe-grpc-auth-golang"
	Board            = "switbe-grpc-board-golang"
	Channel          = "switbe-grpc-channel-golang"
	ChannelV5        = "switbe-grpc-channel-v5-golang"
	Company          = "switbe-grpc-company-golang"
	Contents         = "switbe-grpc-contents-golang"
	Coupon           = "switbe-grpc-coupon-golang"
	Elastic          = "switbe-grpc-elastic"
	Email            = "switbe-grpc-email"
	Emoji            = "switbe-grpc-emoji"
	EventSub         = "switbe-grpc-eventsub-golang"
	Export           = "switbe-grpc-export-golang"
	Import           = "switbe-grpc-import"
	Mention          = "switbe-grpc-mention-golang"
	Message          = "switbe-grpc-message-golang"
	Pay              = "switbe-grpc-pay-golang"
	Project          = "switbe-grpc-project-golang"
	ProjectV5        = "switbe-grpc-project-v5-golang"
	Sfdc             = "switbe-grpc-sfdc-golang"
	Stats            = "switbe-grpc-stats"
	SwitBot          = "switbe-grpc-switbot-golang"
	Task             = "switbe-grpc-task-golang"
	User             = "switbe-grpc-user-golang"
	Workspace        = "switbe-grpc-workspace-golang"
	HttpOauth2       = "switbe-http-oauth2-golang"
	ImportSql        = "switbe-import-sql"
	InviteServer     = "switbe-invite-server"
	MentionSubscribe = "switbe-mention-subscribe-nodejs"
	Notification     = "switbe-notification"
	Openapi          = "switbe-openapi-golang"
	Saml             = "switbe-saml-sp"
	Socket           = "switbe-socket-api"
	SubscribeAgent   = "switbe-subscribe-agent-golang"
	SupportApi       = "switbe-support-api-golang"
	SupportFront     = "switbe-support-front-angular"
	SupportGrpc      = "switbe-support-grpc-golang"
	Batch            = "switbe-batch"
	BatchGrpc        = "switbe-grpc-batch"
	Scheduler        = "switbe-scheduler-api-golang"

	ReleaseV4 = "release.v4"
	ReleaseV5 = "release.v5"
	Main      = "main"
	Master    = "master"
	V1        = "v1"
	V4        = "v4"
	V5        = "v5"
)

func (a *Data) getBranch(str, version string) string {
	branch := Master
	switch str {
	case Hook:
	case ApiV3:
	case ApiV1:
	case ApiV5:
		branch = Main
	case Badge:
	case Storage:
	case Activity:
	case Alltask:
	case Approval:
	case Apps:
	case AppVersion:
	case Asset:
	case Auth:
		if version == V4 {
			branch = ReleaseV4
		}
	case Board:
	case Channel:
		if version == V4 {
			branch = ReleaseV4
		}
	case ChannelV5:
		branch = ReleaseV5
	case Company:
	case Contents:
	case Coupon:
	case Elastic:
	case Email:
	case Emoji:
	case EventSub:
		branch = Main
	case Export:
	case Import:
	case Mention:
		branch = ReleaseV4
	case Message:
		branch = ReleaseV5
	case Pay:
	case Project:
		if version == V4 {
			branch = ReleaseV4
		}
	case ProjectV5:
		branch = ReleaseV5
	case Sfdc:
	case Stats:
	case SwitBot:
	case Task:
		if version == V4 {
			branch = ReleaseV4
		} else if version == V5 {
			branch = ReleaseV5
		}
	case User:
		if version == V4 {
			branch = ReleaseV4
		}
	case Workspace:
		if version == V4 {
			branch = ReleaseV4
		}
	case HttpOauth2:
	case ImportSql:
	case InviteServer:
	case MentionSubscribe:
	case Notification:
	case Openapi:
	case Saml:
	case Socket:
	case SubscribeAgent:
	case SupportApi:
	case SupportFront:
		if version == V4 {
			branch = ReleaseV4
		}
	case SupportGrpc:
	case Batch:
	case BatchGrpc:
	case Scheduler:
	}

	return branch
}

func (a *Data) create() error {
	var str, branch string
	for _, t := range a.Distribution {
		branch = a.getBranch(t.Path, t.Version)

		str += fmt.Sprintf("cd %s \n", basePath+t.Path)
		str += fmt.Sprintf("git checkout %s \n", branch)
		str += fmt.Sprintf("git pull \n")

		cmd := exec.Command("git", "checkout", branch)
		cmd.Dir = "/Users/erickoo/go/src/swit/" + t.Path
		err := cmd.Run()
		if err != nil {
			fmt.Println("err : ", t.Path, " ", t.Version)
			continue
		}

		cmd = exec.Command("git", "pull", branch)
		cmd.Dir = "/Users/erickoo/go/src/swit/" + t.Path
		_ = cmd.Run()

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
	defer func(f1 *os.File) {
		_ = f1.Close()
	}(f1)
	_, err = fmt.Fprintf(f1, string(b))
	if err != nil {
		return err
	}

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
	defer func(f1 *os.File) {
		_ = f1.Close()
	}(f1)
	_, err = fmt.Fprintf(f1, string(b))
	if err != nil {
		panic(err)
	}

	return nil
}

func New(v []*Distribution) *Data {
	return &Data{
		Distribution: v,
	}
}

func main() {
	var s []*Distribution

	//배포
	s = append(s, &Distribution{Path: Company, Version: V1})
	s = append(s, &Distribution{Path: Asset, Version: V1})
	s = append(s, &Distribution{Path: Auth, Version: V1})
	s = append(s, &Distribution{Path: Auth, Version: V4})
	s = append(s, &Distribution{Path: User, Version: V1})
	s = append(s, &Distribution{Path: User, Version: V4})
	s = append(s, &Distribution{Path: SupportGrpc, Version: V1})
	s = append(s, &Distribution{Path: Storage, Version: V1})
	s = append(s, &Distribution{Path: Emoji, Version: V1})
	s = append(s, &Distribution{Path: Approval, Version: V1})
	s = append(s, &Distribution{Path: Contents, Version: V1})
	s = append(s, &Distribution{Path: Notification, Version: V1})
	s = append(s, &Distribution{Path: Task, Version: V4})
	s = append(s, &Distribution{Path: Task, Version: V5})

	//Tech-Switness/switbe-grpc-company-golang
	//Tech-Switness/switbe-grpc-asset-golang
	//Tech-Switness/switbe-grpc-auth-golang
	//Tech-Switness/switbe-grpc-user-golang v1,v4
	//swit-be/swit-support-grpc-golang
	//Tech-Switness/switbe-grpc-company-golang
	//switbe-apiv1
	//switbe-gcs-file-golang
	//switbe-grpc-emoji
	//switbe-grpc-approval-golang
	//switbe-grpc-contents-golang
	//Tech-Switness/switbe-notification
	//switbe-grpc-task-golang

	d := New(s)
	_ = d.merge()
	_ = d.create()
}

// 마지막 태그명 복붙
// git describe --tags $(git rev-list --tags --max-count=1) | pbcopy
// pbpaste
