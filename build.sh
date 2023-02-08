git pull --rebase
git tag -d v1.6.0
git tag v1.6.0

make package
mv *-package/*.tar.gz ~/bin/gocron
cd ~/bin/gocron
tar -zxvf *.tar.gz
mv *-amd64/* ./
rm -rf *-amd64*

