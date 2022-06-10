What
====

I did a clone performance comparison between `git binary version 2.30.1` and `go git v5.4.2`.

The cloned Git repo is the public FreeBSD source repo -- `https://github.com/freebsd/freebsd-src.git`. The shallow clone sets depth as `1`. Use either the branch `releng/10.0` or the tag `release/10.0.0`

| Metrics                                     |                              git binary |           go git |
|:--------------------------------------------|----------------------------------------:|-----------------:|
| Full Cloning time                           |                                2m5.153s | 10m38.560722612s |
| Shallow Cloning by Branch                   |                               0m42.007s |  3m52.281577264s |
| Shallow Cloning by Tag                      |                               0m41.349s |  3m52.043277082s |
| .git directory size full clone              |                                 2.58 GB |          2.58 GB |
| .git directory size shallow clone by Branch |                               198.57 MB |          1.01 GB |
| .git directory size shallow clone by Tag    |                               198.54 MB |          1.01 GB |


Below is the output from `run-time-me.sh`

```azure
***************************************************************************************
 
 
Full clone by git binary
 
 
***************************************************************************************
Cloning into '/tmp/repo-git-full'...
remote: Enumerating objects: 4689957, done.
remote: Counting objects: 100% (875/875), done.
remote: Compressing objects: 100% (581/581), done.
remote: Total 4689957 (delta 349), reused 689 (delta 286), pack-reused 4689082
Receiving objects: 100% (4689957/4689957), 2.26 GiB | 31.71 MiB/s, done.
Resolving deltas: 100% (3241391/3241391), done.
Updating files: 100% (88650/88650), done.

real	2m5.153s
user	6m25.011s
sys	0m48.245s
***************************************************************************************
 
 
Shallow clone by branch by git binary
 
 
***************************************************************************************
Cloning into '/tmp/repo-git-branch'...
remote: Enumerating objects: 60954, done.
remote: Counting objects: 100% (60954/60954), done.
remote: Compressing objects: 100% (51621/51621), done.
Receiving objects: 100% (60954/60954), 181.85 MiB | 11.98 MiB/s, done.
remote: Total 60954 (delta 12594), reused 32727 (delta 6409), pack-reused 0
Resolving deltas: 100% (12594/12594), done.
Updating files: 100% (57174/57174), done.

real	0m42.007s
user	0m13.148s
sys	0m4.617s
***************************************************************************************
 
 
Shallow clone by tag by git binary
 
 
***************************************************************************************
Cloning into '/tmp/repo-git-tag'...
remote: Enumerating objects: 60979, done.
remote: Counting objects: 100% (60979/60979), done.
remote: Compressing objects: 100% (51645/51645), done.
Receiving objects: 100% (60979/60979), 181.83 MiB | 12.07 MiB/s, done.
remote: Total 60979 (delta 12594), reused 32819 (delta 6410), pack-reused 0
Resolving deltas: 100% (12594/12594), done.
Note: switching to 'e3ff8788aa4381f45f5c39efd99c9cbf1b174b5a'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by switching back to a branch.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -c with the switch command. Example:

  git switch -c <new-branch-name>

Or undo this operation with:

  git switch -

Turn off this advice by setting config variable advice.detachedHead to false

Updating files: 100% (57197/57197), done.

real	0m41.349s
user	0m13.584s
sys	0m5.087s
***************************************************************************************
 
 
Clone by go git
 
 
***************************************************************************************
2022/06/10 11:18:13.153555 demo-shallow-clone/main.go:60: Go Git shallow clone by branch
2022/06/10 11:18:13.153736 demo-shallow-clone/main.go:31: Cloning repo https://github.com/freebsd/freebsd-src.git into local directory /tmp/repo-go-git-branch with reference refs/heads/releng/10.0 ...
2022/06/10 11:22:05.435351 demo-shallow-clone/main.go:49: Cloning by refs/heads/releng/10.0 is done! It took 3m52.281577264s
2022/06/10 11:22:05.435383 demo-shallow-clone/main.go:66: Go Git shallow clone by tag
2022/06/10 11:22:05.435395 demo-shallow-clone/main.go:31: Cloning repo https://github.com/freebsd/freebsd-src.git into local directory /tmp/repo-go-git-tag with reference refs/tags/release/10.0.0 ...
2022/06/10 11:25:57.478685 demo-shallow-clone/main.go:49: Cloning by refs/tags/release/10.0.0 is done! It took 3m52.043277082s
2022/06/10 11:25:57.478712 demo-shallow-clone/main.go:71: Go Git shallow clone by tag
2022/06/10 11:25:57.478723 demo-shallow-clone/main.go:31: Cloning repo https://github.com/freebsd/freebsd-src.git into local directory /tmp/repo-go-git-full with reference  ...
2022/06/10 11:36:36.039461 demo-shallow-clone/main.go:49: Cloning by  is done! It took 10m38.560722612s
```