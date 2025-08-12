# kcfg

kcfg is a kubeconfig files manager.

## Idea behind

Save each kubeconfig as a single file inside a folder (~/.kube/kcfg/) and auto merge them into ~/.kube/config.

use a deamon to do that ?

commands:
- list (kubectl config get-contexts)
- use (kubectl config use-context)
- add (copy config file into ~/.kube/kcfg/)
- merge (merge all config inside ~/.kube/kcfg/ into ~/.kube/config)
- clean (interrogate all clusters and delete old clusters)

maybe combine list and use inside root cmd (as ktx did)

> Workspace concept: sometimes there is similar cluster names but in differents networks, it can be useful to have a workspace system to add 1 level whre put config files and be able to switch workspace.

> Autofetcher system: implement auto finding of kubeconfig files on popular kubernetes distribution 