cd $HOME/go/src/swit/swit-scheduler-api-golang
git checkout master
git pull
git merge origin/express
git push origin master

cd $HOME/go/src/swit/swit-grpc-task-golang
git checkout release.v4
git pull
git merge origin/express.v4
git push origin release.v4

cd $HOME/go/src/swit/swit-grpc-task-golang
git checkout release.v5
git pull
git merge origin/express.v5
git push origin release.v5