cd $HOME/go/src/swit/swit-grpc-channel-v5-golang 
git merge origin/express.v5 
git push origin release.v5 
 
cd $HOME/go/src/swit/swit-grpc-project-v5-golang 
git merge origin/express.v5 
git push origin release.v5 
 
cd $HOME/go/src/swit/swit-grpc-workspace-golang 
git merge origin/express.v4 
git push origin release.v4 
 
cd $HOME/go/src/swit/swit-openapi-golang 
git merge origin/express 
git push origin master 
 
