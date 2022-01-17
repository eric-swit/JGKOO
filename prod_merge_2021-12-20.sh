cd $HOME/go/src/swit/swit-grpc-user-golang 
git checkout release.v4 
git pull 
git merge origin/express.v4 
git push origin release.v4 
 
cd $HOME/go/src/swit/swit-grpc-workspace-golang 
git checkout release.v4 
git pull 
git merge origin/express.v4 
git push origin release.v4 
 
cd $HOME/go/src/swit/swit-grpc-switbot-golang 
git checkout master 
git pull 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-grpc-task-golang 
git checkout release.v4 
git pull 
git merge origin/express.v4 
git push origin release.v4 
 
cd $HOME/go/src/swit/swit-grpc-apps-golang 
git checkout master 
git pull 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-grpc-contents-golang 
git checkout master 
git pull 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-grpc-activity-golang 
git checkout master 
git pull 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/hook 
git checkout master 
git pull 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-apiv1 
git checkout master 
git pull 
git merge origin/express 
git push origin master 
 
cd $HOME/go/src/swit/swit-api-golang 
git checkout master 
git pull 
git merge origin/express 
git push origin master 
 
