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
	ChannelV5        = "swit-grpc-channel-v5-golang"
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
	ProjectV5        = "swit-grpc-project-v5-golang"
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
	BatchGrpc        = "swit-grpc-batch"
	Scheduler        = "swit-scheduler-api-golang"

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
		} else if version == V5 {
			branch = ReleaseV5
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
		} else if version == V5 {
			branch = ReleaseV5
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
			panic(err)
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
	s = append(s, &Distribution{Path: ApiV1, Version: V1})
	s = append(s, &Distribution{Path: Workspace, Version: V4})
	s = append(s, &Distribution{Path: Apps, Version: V1})
	s = append(s, &Distribution{Path: Storage, Version: V1})
	s = append(s, &Distribution{Path: Asset, Version: V1})
	s = append(s, &Distribution{Path: Message, Version: V5})
	s = append(s, &Distribution{Path: Approval, Version: V1})
	s = append(s, &Distribution{Path: Activity, Version: V1})
	s = append(s, &Distribution{Path: Notification, Version: V1})
	s = append(s, &Distribution{Path: Contents, Version: V1})
	s = append(s, &Distribution{Path: User, Version: V1})
	s = append(s, &Distribution{Path: Task, Version: V4})
	s = append(s, &Distribution{Path: Hook, Version: V1})
	//1차 배포
	//s = append(s, &Distribution{Path: Badge, Version: V1})
	//s = append(s, &Distribution{Path: Activity, Version: V1})
	//s = append(s, &Distribution{Path: Message, Version: V5})
	//s = append(s, &Distribution{Path: ProjectV5, Version: V5})
	//s = append(s, &Distribution{Path: Task, Version: V4})
	//s = append(s, &Distribution{Path: Task, Version: V5})
	//s = append(s, &Distribution{Path: ApiV3, Version: V1})
	//s = append(s, &Distribution{Path: ApiV1, Version: V1})

	//마이그레이션

	//2차 배포
	//s = append(s, &Distribution{Path: Badge, Version: V1})
	//s = append(s, &Distribution{Path: Notification, Version: V1})
	//s = append(s, &Distribution{Path: Activity, Version: V1})
	//s = append(s, &Distribution{Path: Message, Version: V5})
	//s = append(s, &Distribution{Path: Project, Version: V4})
	//s = append(s, &Distribution{Path: ProjectV5, Version: V5})
	//s = append(s, &Distribution{Path: Task, Version: V4})
	//s = append(s, &Distribution{Path: Task, Version: V5})
	//s = append(s, &Distribution{Path: Workspace, Version: V1})
	//s = append(s, &Distribution{Path: Channel, Version: V4})
	//s = append(s, &Distribution{Path: Channel, Version: V1})
	//s = append(s, &Distribution{Path: ChannelV5, Version: V5})
	//s = append(s, &Distribution{Path: Board, Version: V1})
	//s = append(s, &Distribution{Path: Openapi, Version: V1})
	//s = append(s, &Distribution{Path: Workspace, Version: V4})
	//s = append(s, &Distribution{Path: Contents, Version: V1})
	//s = append(s, &Distribution{Path: ApiV5, Version: V5})
	//s = append(s, &Distribution{Path: ApiV3, Version: V1})
	//s = append(s, &Distribution{Path: ApiV1, Version: V1})

	//s = append(s, &Distribution{Path: ChannelV5, Version: V5})
	//s = append(s, &Distribution{Path: Board, Version: V1})
	//s = append(s, &Distribution{Path: ApiV1, Version: V1})
	//s = append(s, &Distribution{Path: ProjectV5, Version: V5})
	//s = append(s, &Distribution{Path: Badge, Version: V1})
	//s = append(s, &Distribution{Path: Notification, Version: V1})
	//s = append(s, &Distribution{Path: Activity, Version: V1})
	//s = append(s, &Distribution{Path: Message, Version: V5})
	//s = append(s, &Distribution{Path: Project, Version: V4})
	//s = append(s, &Distribution{Path: Task, Version: V4})
	//s = append(s, &Distribution{Path: Task, Version: V5})
	//s = append(s, &Distribution{Path: Workspace, Version: V1})
	//s = append(s, &Distribution{Path: Channel, Version: V4})
	//s = append(s, &Distribution{Path: ApiV3, Version: V1})
	//s = append(s, &Distribution{Path: Contents, Version: V1})
	//s = append(s, &Distribution{Path: ApiV5, Version: V5})
	//s = append(s, &Distribution{Path: Approval, Version: V1})
	//s = append(s, &Distribution{Path: Company, Version: V1})
	//s = append(s, &Distribution{Path: Apps, Version: V1})
	//s = append(s, &Distribution{Path: Mention, Version: V1})
	//s = append(s, &Distribution{Path: EventSub, Version: V1})
	//s = append(s, &Distribution{Path: Emoji, Version: V1})
	//s = append(s, &Distribution{Path: Elastic, Version: V1})
	//s = append(s, &Distribution{Path: Alltask, Version: V1})
	//s = append(s, &Distribution{Path: Activity, Version: V1})
	//s = append(s, &Distribution{Path: SwitBot, Version: V1})
	//s = append(s, &Distribution{Path: Workspace, Version: V4})
	//s = append(s, &Distribution{Path: AppVersion, Version: V1})
	//s = append(s, &Distribution{Path: Stats, Version: V1})
	//s = append(s, &Distribution{Path: SubscribeAgent, Version: V1})

	//
	//
	//s = append(s, &Distribution{Path: ApiV1, Version: V1})
	//s = append(s, &Distribution{Path: ApiV3, Version: V1})
	//s = append(s, &Distribution{Path: ApiV5, Version: V5})
	//s = append(s, &Distribution{Path: ChannelV5, Version: V5})
	//s = append(s, &Distribution{Path: Board, Version: V1})
	//s = append(s, &Distribution{Path: ProjectV5, Version: V5})
	//s = append(s, &Distribution{Path: Badge, Version: V1})
	//s = append(s, &Distribution{Path: Project, Version: V4})
	//s = append(s, &Distribution{Path: Task, Version: V4})
	//s = append(s, &Distribution{Path: Task, Version: V5})
	//s = append(s, &Distribution{Path: Workspace, Version: V1})
	//s = append(s, &Distribution{Path: Channel, Version: V4})
	//s = append(s, &Distribution{Path: Contents, Version: V1})
	//s = append(s, &Distribution{Path: Approval, Version: V1})
	//s = append(s, &Distribution{Path: Company, Version: V1})
	//s = append(s, &Distribution{Path: Message, Version: V5})
	//s = append(s, &Distribution{Path: Activity, Version: V1})
	//s = append(s, &Distribution{Path: Notification, Version: V1})
	//s = append(s, &Distribution{Path: Apps, Version: V5})
	//
	//
	////gorm
	//s = append(s, &Distribution{Path: Mention, Version: V1})
	//s = append(s, &Distribution{Path: EventSub, Version: V1})
	//s = append(s, &Distribution{Path: Emoji, Version: V1})
	//s = append(s, &Distribution{Path: Elastic, Version: V1})
	//s = append(s, &Distribution{Path: Alltask, Version: V1})
	//s = append(s, &Distribution{Path: SwitBot, Version: V1})
	//s = append(s, &Distribution{Path: Workspace, Version: V4})
	//s = append(s, &Distribution{Path: AppVersion, Version: V1})
	//s = append(s, &Distribution{Path: Stats, Version: V1})
	//s = append(s, &Distribution{Path: SubscribeAgent, Version: V1})

	//s = append(s, &Distribution{Path: Channel, Version: V5})
	//s = append(s, &Distribution{Path: Channel, Version: V4})
	//s = append(s, &Distribution{Path: Approval, Version: V1})
	//s = append(s, &Distribution{Path: Mention, Version: V1})
	//s = append(s, &Distribution{Path: EventSub, Version: V1})
	//s = append(s, &Distribution{Path: Emoji, Version: V1})
	//s = append(s, &Distribution{Path: Elastic, Version: V1})
	//s = append(s, &Distribution{Path: Company, Version: V1})
	//s = append(s, &Distribution{Path: Contents, Version: V1})
	//s = append(s, &Distribution{Path: Board, Version: V1})
	//s = append(s, &Distribution{Path: Alltask, Version: V1})
	//s = append(s, &Distribution{Path: Asset, Version: V1})
	//s = append(s, &Distribution{Path: Activity, Version: V1})
	//s = append(s, &Distribution{Path: Badge, Version: V1})
	//s = append(s, &Distribution{Path: Notification, Version: V1})
	//s = append(s, &Distribution{Path: Project, Version: V1})
	//s = append(s, &Distribution{Path: Project, Version: V4})
	//s = append(s, &Distribution{Path: Project, Version: V5})
	//s = append(s, &Distribution{Path: SwitBot, Version: V1})
	//s = append(s, &Distribution{Path: Workspace, Version: V4})
	//s = append(s, &Distribution{Path: AppVersion, Version: V1})
	//s = append(s, &Distribution{Path: Stats, Version: V1})
	//s = append(s, &Distribution{Path: SubscribeAgent, Version: V1})
	//
	//s = append(s, &Distribution{Path: Task, Version: V5})
	//s = append(s, &Distribution{Path: Message, Version: V1})
	//s = append(s, &Distribution{Path: User, Version: V1})
	//s = append(s, &Distribution{Path: Apps, Version: V1})
	//
	//s = append(s, &Distribution{Path: ApiV3, Version: V1})
	//s = append(s, &Distribution{Path: ApiV5, Version: V1})
	//s = append(s, &Distribution{Path: ApiV1, Version: V1})

	d := New(s)
	_ = d.merge()
	_ = d.create()
}

// 마지막 태그명 복붙
// git describe --tags $(git rev-list --tags --max-count=1) | pbcopy
// pbpaste
