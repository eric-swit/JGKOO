cd $HOME/go/src/swit/swit-grpc-channel-v5-golang 
git checkout release.v5 
git pull 
git tag release-v5.3.6
 
git push origin release-v5.3.6
 
 
cd $HOME/go/src/swit/swit-grpc-project-v5-golang 
git checkout release.v5 
git pull 
git tag release-v5.3.8
 
git push origin release-v5.3.8
 
 
cd $HOME/go/src/swit/swit-grpc-workspace-golang 
git checkout release.v4 
git pull 
git tag release-v4.3.9
 
git push origin release-v4.3.9
 
 
cd $HOME/go/src/swit/swit-openapi-golang 
git checkout master 
git pull 
git tag release-v1.3.6
 
git push origin release-v1.3.6
 
 
