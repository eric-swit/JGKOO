cd $HOME/go/src/swit/swit-apiv1 
git checkout master 
git pull 
git tag release-v1.3.19

git push release-v1.3.19


cd $HOME/go/src/swit/swit-api-golang
git checkout prod
git pull
git tag release-v1.3.21

git push origin release-v1.3.21

 
cd $HOME/go/src/swit/swit-socket-api
git checkout master
git pull
git tag release-v1.3.3

git push origin release-v1.3.3

cd $HOME/go/src/swit/swit-gcs-file-golang 
git checkout master 
git pull 
git tag release-v1.4.3
 
git push origin release-v1.4.3
 
 
