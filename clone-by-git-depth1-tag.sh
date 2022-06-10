LOCALDIR=/tmp/repo-git-tag
REPO_URL=https://github.com/freebsd/freebsd-src.git
TAG=release/10.0.0
rm -rf $LOCALDIR
git clone --depth=1 -b $TAG $REPO_URL $LOCALDIR
