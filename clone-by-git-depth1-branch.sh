LOCALDIR=/tmp/repo-git-branch
REPO_URL=https://github.com/freebsd/freebsd-src.git
BRANCH=releng/10.0
rm -rf $LOCALDIR
git clone --depth=1 -b $BRANCH $REPO_URL $LOCALDIR
