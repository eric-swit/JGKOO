cd $HOME/go/src/swit/hook 
git checkout master 
git pull 
git tag release-v1.1.11
 
git push origin release-v1.1.11
 
 
cd $HOME/go/src/swit/swit-grpc-workspace-golang
git checkout release.v4
git pull
git tag release-v4.3.8

git push origin release-v4.3.8


cd $HOME/go/src/swit/swit-grpc-apps-golang
git checkout master
git pull
git tag release-v1.3.12

git push origin release-v1.3.12


cd $HOME/go/src/swit/swit-gcs-file-golang
git checkout master
git pull
git tag release-v1.3.10

git push origin release-v1.3.10


cd $HOME/go/src/swit/swit-grpc-asset-golang
git checkout master
git pull
git tag release-v1.3.7

git push origin release-v1.3.7


cd $HOME/go/src/swit/swit-grpc-message-golang
git checkout release.v5
git pull
git tag release-v5.3.18

git push origin release-v5.3.18


cd $HOME/go/src/swit/swit-grpc-approval-golang
git checkout master
git pull
git tag release-v1.3.6

git push origin release-v1.3.6


cd $HOME/go/src/swit/swit-grpc-activity-golang
git checkout master
git pull
git tag release-v1.3.9

git push origin release-v1.3.9


cd $HOME/go/src/swit/swit-notification
git checkout master
git pull
git tag release-v1.3.10

git push origin release-v1.3.10


cd $HOME/go/src/swit/swit-grpc-contents-golang
git checkout master
git pull
git tag release-v1.3.15

git push origin release-v1.3.15


cd $HOME/go/src/swit/swit-grpc-user-golang
git checkout master
git pull
git tag release-v1.3.7

git push origin release-v1.3.7


cd $HOME/go/src/swit/swit-grpc-task-golang
git checkout release.v4
git pull
git tag release-v4.3.7

git push origin release-v4.3.7


cd $HOME/go/src/swit/swit-apiv1
git checkout master
git pull
git tag release-v1.3.18

git push origin release-v1.3.18