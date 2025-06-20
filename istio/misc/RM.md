Modified version of https://github.com/istio/istio/issues/55908 (1.26)

---

## Introduction

This issue tracks all steps needed to prepare the 1.27 release branches/builds.

This issue was created to help document the automation steps as well as any ordering changes, etc. Please update/add additional steps taken here so we can document for future releases.

**Note that the actual coding steps done by the automation steps are left in this issue for reference only. You only need to run the automated steps, and not the individual instructions. Also note that the legacy reference steps may be INCOMPLETE as the automation is updated, but the automation steps needed to actually run the release should always be updated to be accurate.**

## The Automation

The release-builder automation can be found in the master branch of https://github.com/istio/release-builder. It's best if you fork this repo into your own org, as you will need to submit PRs into the repo to trigger the automation to run on GitHub rather than locally.

It is also worth forking all of the other Istio repos specified in [`release/branch.sh`](https://github.com/istio/release-builder/blob/master/release/branch.sh) into your own organization, as that way you can dry-run against it.

### Important things to know about the automation

First, spend a little time understanding the release-builder [README](https://github.com/istio/release-builder/blob/master/README.md).  Don't worry if you don't understand all of it, its not really necessary to in order to work through the rest of these steps, but having some understanding will help.  Here are the important things you must know:

The release directory contains three config files:
   * trigger-branch - this contains the version and step for the automation to work on
   * trigger-build - this contains the tag for the release-builder to use for whatever its building
   * trigger-publish - this contains the tag that the release-builder will publish.

Before you start testing the automation with dry runs, update those files to whatever you need them to be.

Each time you start the build container with `make shell` you need to:
* Run `gh auth login` (watch for errors, you may need to run this command with `sudo`
You also need to remember that for not-dry-runs you need to ensure that the GITHUB_TOKEN environment variable is set **within** the build container.  The automation does not care at all if you've logged in with the `gh` CLI.  The token that you use has to have `repo` permissions.  If you don't do this, the automation will fail when trying to interact with github by pushing, creating PRs etc.

### How to run it

Update the relevant files to specify what you want, than:

```
gh auth login
GITHUB_TOKEN=$(gh auth token) make shell
REPO_ORG=istio STEP=x ./release/branch.sh
```

This generates the changes for you to review (look at `git status` or `git diff` to see them).  When you are satisfied with what you see, run

```
REPO_ORG=istio STEP=x DRY_RUN=false ./release/branch.sh 
```

This will cause the actual step to happen and PRs to be raised in whichever github org you specified in `REPO_ORG`

Note that the automated steps 3 and on require the new release's branch in the various repos which are created by automated step 2.  

#### On a Mac
* ~~Run `make shell` with `sudo`~~
* Later versions of Docker Desktop for Mac may require the [`DOCKER_HOST`](https://docs.docker.com/desktop/setup/install/mac-permission-requirements/#installing-symlinks) environment variable set so the build container can communicate with Docker on the host.
```
DOCKER_HOST="unix:///var/run/docker.sock" make shell
```
* Run`brew install docker-credential-helper` first to avoid difficult-to-spot errors in the scripts
* Edit `~/.docker/config.json`, change the line `credsStore` (two s) to `credstore`

#### Handling common problems

If you hit CLA issues from the automation omitting your email address in the commit (possibly due to [configuring GitHub to keep your email address private](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-personal-account-on-github/managing-email-preferences/setting-your-commit-email-address)?), you may need to checkout the PR branch locally and manually update the author, then force push the commit back to the branch.
```
gh pr checkout [PR_NUMBER]
git commit --amend --reset-author
git push --force
```

## Manual / Pre-req steps
### When to start: 2 weeks before branch cut date
- [ ] Set up release manager team - Ensure the new team has write access to the repositories. (see [teams.yaml](https://github.com/istio/community/blob/master/org/teams.yaml)) 
  - ~~https://github.com/istio/community/pull/1597~~
- [ ] Create a slack channel for the [1.27 release](). Edit the description/topic of the channel to include RM's for the release, expected release date, etc.
- [ ] Add new label: cherrypick/release-1.27 PR
- [ ] Create a 1.28 milestone:
- [ ] Send an announcement on [GitHub Discussions](https://github.com/istio/istio/discussions) pre-announcing branch cut
    - ~~https://github.com/istio/istio/discussions/55912~~
- [ ] Send an announcement on [GitHub Discussions](https://github.com/istio/istio/discussions) announcing the start of branch cut and to stop merging of PRs
- [ ] Fork (and update) istio repos mentioned in Step2. Example of forking a single repo from cli
```
gh repo fork --clone --default-branch-only https://github.com/istio/common-files --fork-name istio-common-files
```

## Step 0 - Alpha build

Build and release 1.27-alpha.0 from master branch.
- [ ] ~~https://github.com/istio/release-builder/pull/2082~~
- [ ] Find the triggered [`build-release_release-builder_postsubmit`](https://prow.istio.io/view/gs/istio-prow/logs/build-release_release-builder_postsubmit/1881824336185135104) job on https://prow.istio.io/ and wait for it to complete successfully.
- [ ] ~~https://github.com/istio/release-builder/pull/2083~~

**Because the alpha build is tagged from the master branch, merging these PRs requires approval from [WG - Test and Release Maintainers](https://github.com/orgs/istio/teams/wg-test-and-release-maintainers), not just leads for the specific release.**

## Automation Step 1 - Update Dependencies (2 weeks out)

This step updates dependencies in master prior to the release branch being cut

### Command to run

```
CONDITIONAL_HOST_MOUNTS=" --mount type=bind,source=${HOME}/.ssh,destination=/home/ubuntu/.ssh,readonly " GITHUB_TOKEN=$(gh auth token) make shell
REPO_ORG=istio STEP=1 ./release/branch.sh
```
- [ ] PR after the automation step 1, legacy manual steps should generate the same (~~https://github.com/istio/istio/pull/55972~~)

### Legacy Manual Steps
- (**Automation step=1**) Update dependencies. This must be done **before branching** so that PRs can be easily backported later. [PR](https://github.com/istio/istio/pull/44436)
  - Run `go get github.com/envoyproxy/go-control-plane@main` to update (Added in 1.17 in https://github.com/istio/release-builder/pull/1332).
  - Run the following command in istio/istio `UPDATE_BRANCH=master ./bin/update_deps.sh; make gen` ([example](https://github.com/istio/istio/pull/36812))
  - [ ] Merge PR in istio/istio. Wait for it to complete. 
  - [ ] _Need to verify git code to allow commits and pulls to work for upcoming 1.27 branch cut._

## Automation Step 2 - Create The Release Branches
### When to start: 1 week before branch cut date

This step creates release-1.27 branches in the necessary repos. Please run it against the istio org since there are no PRs for this step 

### Command to run
```
make shell
REPO_ORG=(myorg|istio) VERSION=1.27 STEP=2 ./release/branch.sh
```

### Legacy Manual Steps
- (**Automation step=2**)  (new step because we are waiting on prior change to master before branching) Create a `release-1.27` branch in every Istio repo - run commands `export org=xyz ; export repo=xyz ; (git clone git@github.com:${org}/${repo}.git && cd $repo && git checkout -b release-1.27 && git push --set-upstream origin release-1.27)`.
  - [ ] istio/istio
  - [ ] istio/api
  - [ ] istio/ztunnel
  - [ ] istio/proxy
  - [ ] istio/client-go
  - [ ] istio/tools
  - [ ] istio/common-files
  - [ ] istio/release-builder
  - [ ] istio/enhancements
  - **Explicitly skipped:** istio-releases/pipeline, istio/tests, istio/istio.io, istio/test-infra
  - No PRs to merge

### Post-automation steps
- [ ] Set up branch protection. Modify prow/config.yaml in test-infra to add release managers as owners for a branch. (https://github.com/istio/test-infra/pull/5646)

## Automation Step 3 - Set Up Prow On Release Branches

**You will need a T&R maintainer who has the ability to tag new images in the istio-testing Google Artifact Registry repo (e.g. John Howard or Daniel Hawton) to assist with this step.**

### Command to run
```
make shell
REPO_ORG=myorg VERSION=1.27 STEP=3 ./release/branch.sh
```
Running this command with `DRY_RUN=false` will open a PR like https://github.com/istio/test-infra/pull/5598, **but because this automation is still incomplete, a few additional manual steps will be needed.**

The output logs will tell you the necessary commands for retagging image and should look something like as follows:
```
2025/02/06 17:24:01 Please find a maintainer with sufficient permissions and have them run `gcloud container image add-tag gcr.io/istio-testing/build-tools:master-6de2ce5813071b34e0ca033dbac7f79ffc1644be gcr.io/istio-testing/build-tools:release-1.27-6de2ce5813071b34e0ca033dbac7f79ffc1644be`
2025/02/06 17:24:01 Please find a maintainer with sufficient permissions and have them run `gcloud container image add-tag gcr.io/istio-testing/build-tools-proxy:master-6de2ce5813071b34e0ca033dbac7f79ffc1644be gcr.io/istio-testing/build-tools-proxy:release-1.27-6de2ce5813071b34e0ca033dbac7f79ffc1644be`
```

You can query tag like this:
```
gcloud container images list-tags gcr.io/istio-testing/build-tools | grep master-latest
```

Checkout your fork of the test-infra project, make sure its up to date with `master`, then create a new branch and do the following:
- [ ] Manually update [testgrid/config.yaml](https://github.com/istio/test-infra/blob/master/testgrid/config.yaml) by adding in the relevant entries for 1.27.  Its pretty obvious once you see the file what you have to do.
- [ ] Manually update [prow/config/private-presets.yaml](https://github.com/istio/test-infra/blob/master/prow/config/private-presets.yaml). Again, this should be pretty obvious, just add new entries similar to the last release but for the new branch.
- [ ] Cherry-pick the commit from the automation into this branch, **close the automation PR** - CI checks depend on the manual changes you'll need to add to pass.
- [ ] Remove duplicated `env` entries of the following in the configs
    ```
    - name: BUILD_WITH_CONTAINER
      value: "0"
    ```
  - **TODO:** I believe this should be harmless and is _mostly_ automated for the generated output files in `prow/cluster` after https://github.com/istio/test-infra/pull/5592, but similar env var filtering should likely be added to `tools/prowtrans/cmd/prowtrans/main.go` to rewrite the source configs too to make this step unnecessary.
- [ ] Manually update the `VERSION` env var in `prow/config/jobs/release-builder-1.27.yaml` from `master` to `"1.27"`
  - **TODO:** This could likely be automated too, quotes are necessary to unmarshal as string type.
- [ ] Run `make gen`
- [ ] Merge PR in istio/test-infra. Wait for all postsubmit jobs (visible at https://prow.istio.io/) to complete. (https://github.com/istio/test-infra/pull/5647)

## Automation Step 4

**Do not start step 4 until the post submit jobs in step 3 are completed!**

This step has a two steps:

* Run the automation
* Merge the PRs created by it, EXCEPT FOR ONE! As part of the istio/istio pull request post-submits, a new PR will be created in the istio/common-files repo with a title like: `Automator: update build-tools image@release-1.27 in istio/common-files@release-1.27`. The work in this PR is only a portion of Step 5, and merging it will cause issues with reverting the repos to the main branch common-files. **CLOSE THIS PR**.

### Command To Run
```
make shell
REPO_ORG=istio VERSION=1.27 STEP=4 ./release/branch.sh
```

These are the PRs created by the automation for 1.27:

- https://github.com/istio/proxy/pull/6245
- https://github.com/istio/release-builder/pull/2087
- https://github.com/istio/enhancements/pull/194
- https://github.com/istio/istio/pull/55982
- https://github.com/istio/api/pull/3492
- https://github.com/istio/ztunnel/pull/1528
- https://github.com/istio/client-go/pull/2430
- https://github.com/istio/tools/pull/3192

### Legacy Manual Instructions
- (**Automation step=4**) (new step since we want the automation from prior step to actually create the image after PR merges) PRS: 
  - Updates istio/tools to build new release-1.27 build image (update BRANCH in docker/build-tools/build-and-push.sh. PR postsubmit will create new container images who's name will be used in the next step)  
  - Update common-files in _new release_. You first have to manually update the _common/Makefile.common.mk_ *update-common* target to point to the _new release_  (but not in common-files). In Step=5, the prow automation will actually call `make update-common` in these repos to do the actual `make update-common` against the new common-files release branch
    - istio/common-files - skipped. Covered in next step
    - istio/istio 
    - istio/pkg 
    - istio/api
    - istio/client-go
    - istio/proxy 
    - istio/release-builder 
    - istio/tools 
  - Update `CODEOWNERS` to contain only the release managers for this release. Command: `export org=xyz ; export repo=xyz ; (git clone git@github.com:${org}/${repo}.git && cd $repo && git checkout release-1.27 && git checkout -b release-1.27-codeowners && echo '* @istio/release-managers-1.27' > CODEOWNERS && git add CODEOWNERS && git commit -m 'Set release managers as CODEOWNERS for release-1.27' && git push --set-upstream origin release-1.27-codeowners)`.
    - istio/common-files  - skipped. Covered in next step
    - istio/istio 
    - istio/pkg 
    - istio/api 
    - istio/proxy
    - istio/client-go 
    - istio/tools
    - istio/release-builder
    -  istio/enhancements
  - Stop publishing `latest` tags
  - Update istio/release-builder branch changes from `master` to new release in build.sh, publish.sh and manifests. 
  
 - [ ] Merge PRs from STEP 4. Wait for the new build images to be created at https://gcr.io/istio-testing/build-tools.
**WARNING -- DO NOT MERGE** Merge the PRs created by it, EXCEPT FOR ONE! As part of the istio/istio pull request post-submits, a new PR will be created in the istio/common-files repo with a title like: `Automator: update build-tools image@release-1.27 in istio/common-files@release-1.27`. The work in this PR is only a portion of Step 5, and merging it will cause issues with reverting the repos to the main branch common-files. **CLOSE THIS PR**. (Closed https://github.com/istio/common-files/pull/1174)

## Automation Step 5

**Important**: DO NOT START STEP 5 UNTIL ALL POSTSUBMIT JOBS FROM STEP 4 ARE COMPLETED, CHECK https://prow.istio.io/ FOR STATUS.

### Command to run

```
make shell
REPO_ORG=istio VERSION=1.27 STEP=5 ./release/branch.sh
```

- https://github.com/istio/common-files/pull/1175

**Wait until step 5 PR merges and all the postsubmit generated PRs (to update common files in repos) merge.**

### Legacy Manual Steps
- (**Automation step=5**) (new step since we need image from prior step)
  - Update istio/common-files to set release-1.27 build image (Update the UPDATE_BRANCH in files/common/Makefile.common.mk to be the new release name
  - Update IMAGE_VERSION in files/common/scripts/setup_env.sh  to be the new build image from prior step (found at https://gcr.io/istio-testing/build-tools).
  - Also Update `CODEOWNERS` to contain only the release managers for this release. 
 - [ ] Merge PR from STEP 5. Wait until step 5 PR merges and all the postsubmit generated PRs (to update common files in repos) merge.

## Remaining Manual Steps

After the automation is complete, there are some manual steps that need to be completed to finish the branch cut.

- [ ] Fix up test-infra proxy automated job to pull from most recent stable Envoy release branch if (set `UPDATE_BRANCH`)
  - Check https://www.envoyproxy.io/docs/envoy/latest/version_history/version_history to find the latest stable minor version of Envoy.
  - https://github.com/istio/test-infra/pull/5649
- [ ] Verify that the github.com/envoyproxy/go-control-plane Go module dependencies match the latest commit before github.com/envoyproxy/envoy was branched.
  - Find the new Envoy branch in GitHub (https://github.com/envoyproxy/envoy/tree/release/v1.34 in this case) and click on `commits behind main` to find the commit where the branch happened (it should appear first in in the chronologically-sorted commit list).
  - Click the sha to view the commit (which should look similar to https://github.com/envoyproxy/envoy/commit/a5cf609225dfd223ec734cdc2d9a2cb33e58cacc), then click the sha next to `1 parent` (or manually search the commit history on the main branch), which should take you to the last commit before the branch, which should look similar to https://github.com/envoyproxy/envoy/commit/b0f43d67aa25c1b03c97186a200cc187f4c22db3).
  - Find this commit on the Envoy release branch commit history (https://github.com/envoyproxy/envoy/commits/release/v1.33 in this case), then search https://github.com/envoyproxy/go-control-plane/commits/main for either this commit or the most recent before it - not all commits will have a corresponding mirror commit). In this example, https://github.com/envoyproxy/go-control-plane/commit/4eb1955954fa7607752b3957a1885ce9fe6cf7e8 is the go-control-plane commit we've chosen, which corresponds to https://github.com/envoyproxy/envoy/commit/078dae3549912e632c3776a5e9a4679226093276.
  - Update the istio/istio and istio/proxy repos to use this commit for the github.com/envoyproxy/go-control-plane Go module dependency with `go get github.com/envoyproxy/go-control-plane@GITSHA` where GITSHA is the go-control-plane commit you've selected.
  - (https://github.com/istio/istio/pull/56001)
  - (https://github.com/istio/proxy/pull/6249)
- [ ] Update UPDATE_BRANCH in `bin/update_deps.sh` and `bin/update_ztunnel.sh` on release branch, then run those scripts and commit the changes.
  - (https://github.com/istio/istio/pull/56001)
- [ ] Bump 1.27 to 1.28 in the `master` branch
  - Updates 1.27.0 to 1.28.0, 1.27-dev to 1.28-dev, etc
  - https://github.com/istio/istio/pull/56000
- [ ] Post on [GitHub Discussions](https://github.com/istio/istio/discussions) announcing branch cut complete and PRs can be merged again.
- [ ] Ask istio.io team to run job to update to use the 1.27 branch (probably Daniel Hawton). Verify the EOL date for the n-2 release as 6 weeks after the planned release date.
  - The `make` command documented at https://github.com/istio/istio.io/?tab=readme-ov-file#when-istio-source-code-is-branched can generate a PR similar to below. `make prepare-1.27.0`
  - https://github.com/istio/istio.io/pull/16420

## Publishing a release
The next step is to publish a release.  You will need to follow these steps for each beta and release candidate and finally for the final release build.

- [ ] Run `make shell` then `./bin/update_deps.sh` on the release branch in istio/istio. PR and merge the resulting diff.
- [ ] Trigger a beta or release candidate build
- [ ] Some verification to verify build is good
  - You can download the build tar from the appropriate directory here: https://gcsweb.istio.io/gcs/istio-prerelease/prerelease/
  - untar the file and install, specifying the image repository. (ex: `./bin/istioctl install --set profile=demo -y --set hub=gcr.io/istio-prerelease-testing`).
  - Run bookinfo from within the untar'd directory
- [ ] Publish the beta or release candidate
- [ ] Post on [GitHub Discussions](https://github.com/istio/istio/discussions) announcing availability of beta or release candidate. Release managers may not have permission to post in the Announcements category on GitHub Discussions, get a TOC or Steering member to change the category after posting.
  - [ ] Share this post in the #announcements and #release-1.27 channels on Istio Slack

## Generating a release candidate

### When to start: 10-14 days before the release date

- [ ] Verify and update the min K8s version supported for the release.
  - RMs usually determine the range of k8s versions and it appears on the TOC agenda at times. The max version is usually is the current K8s version and the min version includes 3 past versions, so for a 1.32 max version, the min version would be 1.29 as an example.

Use the Publishing a Release instructions above to update_deps, trigger the build, do verification and publish the release

Next, contact Sergii Shapar on Istio Slack and ask him to run the long running tests on the build, and any subsequent release candidate builds if there winds up being more then one.

## Preparing for final release

### Generating the release notes
- [ ] Prepare the release notes for the release.
- [ ] Update istio.io docs to new version following steps at https://github.com/istio/istio.io?tab=readme-ov-file#creating-a-majorminor-release

Preparing the release notes for a major release is time-consuming; allow 4-5 hours across several days to factor in time for reviews. **Publishing requires approval from [istio/wg-docs-maintainers-infra](https://github.com/orgs/istio/teams/wg-docs-maintainers-infra), _not_ release managers, a maintainer with admin privileges to handle the GitHub branch switch, and someone with Netlify and Google Custom Search Engine privileges (currently Craig Box) to cut over the release branch.** An example PR for the 1.18 release notes is [here](https://github.com/istio/istio.io/pull/13269)

To generate the release notes, first make sure you have up-to-date forks of istio, API and tools repositories checked out to the same folder.  Specifically, they must contain the release-N and release-N-1 branches (where N is the version that your working on, and N-1 is the previous one), so you must not have only forked the repos' master branches. 

To generate release notes, run the command `./gen-release-notes --notes ../../../istio --oldBranch 1.17.0 --newBranch release-1.18` (against the istio/istio and istio/api repos). Follow the instructions from [here](https://github.com/istio/tools/tree/master/cmd/gen-release-notes).  This will get you text files that contain the contents of what will be the change-notes part of you're PR (see the 1.18 PR above for what I mean)

The output of the tool will fail linting.  Once you've copied the content into the right place in you're fork of the isio.io repo, run `make lint` than go get a cup of coffee.  Five to ten minutes later, you'll have a list of things that are wrong with the automatically generated content that you'll need to work through.  The biggest source of errors is that there cannot be more then one blank line between any line in the file, headers must be surrounded by blank lines, all istio.io links must be relative and that there will be trailing whitespace that will need removing.  To fix the last of these, run `sed -i 's/[[:space:]]*$//' file.txt`; this will clean up the trailing whitespace that will otherwise drive you slightly mad trying to find (since its not always quite on the line indicated by the tool)

- [ ] Each time you publish a beta/rc/official release build, make sure that the build/helm charts/images are published correctly to:
  - GitHub releases https://github.com/istio/istio/releases
  - Helm repo https://gcr.io/istio-release/charts/base
  - Google Container Registry repo https://console.cloud.google.com/gcr/images/istio-release
- [ ] After official release, send out an announcement on `announcement` and release slack channel for Istio. Currently Craig Box handles all the twitter announcements related to Istio and he usually sends release related announcements on twitter. Coordinate with him. [discuss](https://discuss.istio.io/t/istio-1-17-is-out/14951), [twitter](https://twitter.com/IstioMesh/status/1625654199336968197?cxt=HHwWioDR8bL4vY8tAAAA)
- [ ] Update the description in the release slack channel to reflect the official release date.   