## MODULE VERSIONING

When a module go.mod file is not found on the root of the repository but in a subfolder, then taging for this module must include the sufolder name as a tag prefix

for example if logger is the subfolder containg the go.mod file, tagging for this module would be
**git tag logger/v0.0.1**

to push changes
**git push origin logger/v0.0.1**

or
git push --tags

if folder is the root so repository name becomse logger, then tagging would be
**git tag v0.0.1**



source: https://go.dev/ref/mod#modules-overview
source: https://go.dev/ref/mod#vcs-version
