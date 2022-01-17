cd $HOME/go/src/swit/swit-grpc-user-golang 
git checkout release.v4 
git pull 
git tag release-v4.3.2
git push origin release-v4.3.2
cd $HOME/go/src/swit/swit-grpc-workspace-golang
git checkout release.v4
git pull
git tag release-v4.3.6
git push origin release-v4.3.6
cd $HOME/go/src/swit/swit-grpc-switbot-golang
git checkout master
git pull
git tag release-v1.3.1
git push origin release-v1.3.1
cd $HOME/go/src/swit/swit-grpc-task-golang
git checkout release.v4
git pull
git tag release-v4.3.3
git push origin release-v4.3.3
cd $HOME/go/src/swit/swit-grpc-apps-golang
git checkout master
git pull
git tag release-v1.3.4
git push origin release-v1.3.4
cd $HOME/go/src/swit/swit-grpc-contents-golang
git checkout master
git pull
git tag release-v1.3.8
git push origin release-v1.3.8
cd $HOME/go/src/swit/hook
git checkout master
git pull
git tag release-v1.1.9
git push origin release-v1.1.9
cd $HOME/go/src/swit/swit-apiv1
git checkout master
git pull
git tag release-v1.3.9
git push origin release-v1.3.9
cd $HOME/go/src/swit/swit-api-golang
git checkout master
git pull
git tag release-v1.3.10
git push origin release-v1.3.10