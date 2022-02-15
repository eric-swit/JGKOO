cd $HOME/go/src/swit/swit-grpc-message-golang 
git merge origin/express.v5 
git push origin release.v5 
 
cd $HOME/go/src/swit/swit-grpc-apps-golang 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-grpc-user-golang 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-grpc-sfdc-golang 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-grpc-channel-golang 
git merge origin/express.v4 
git push origin release.v4 
 
cd $HOME/go/src/swit/swit-grpc-company-golang 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-api-golang 
git merge origin/express 
git push origin master 
 
