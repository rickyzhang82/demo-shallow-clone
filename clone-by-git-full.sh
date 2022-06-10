LOCALDIR=/tmp/repo-git-full
REPO_URL=https://github.com/freebsd/freebsd-src.git
rm -rf $LOCALDIR
git clone $REPO_URL $LOCALDIR
