# SavannahTech

## Description
This is a project that retrieves commit data about a specified github repository, the application provides three main features for retrieving repository information. 
## Key Features

&nbsp;
__Get Top N Commit Authors__: This feature returns the top N Committers for the given Github Repository.\
&nbsp;
__Get All Commits for Repository__: This feature retrieves all commits for the given Github Repository. \
&nbsp;
__Get Commits From a given Date__: This feature retrieves commits from the given date to the current time. 

## Environment Variables

To run this project, you will need to add the following environment variables to your example.env file

`DATABASE_HOST`
`DATABASE_USERNAME`
`DATABASE_PASSWORD`
`DATABASE_NAME`
`DATABASE_PORT`
`REPO_OWNER`
`REPO_NAME`
`AUTH_TOKEN` (github access token)
`FETCH_DATE_SINCE`(Format: MM-DD-YYYY)



## Getting Started
1. To Run the project, navigate to the projects root directory and execute the 'go run main.go' command in the projects root directory
## Usage
To Run the project, navigate to the projects root directory and do the following:
1. Run 'docker network create fullstack' before executing the application the first time.
2. Run 'docker-compose up --build' to build the application and execute the tests.
3. Run 'docker-compose run api' to start the application.

## Testing
- The test suites for all the components of the application are located in the controllers, repositories and services directories in the api directory. To run the tests, navigate to the projects root directory and execute one of the following commands:
- go test .\services .\repositories .\controllers
- 'docker-compose up --build'
## Technologies Used
Golang
## API Documentation

### Commits

#### GET http://localhost:8082/commits?size=3

- Get Top N commits
**Request Sample**
```shell
curl --location 'localhost:8082/commits?size=3' \
--header 'Content-Type: application/json' \'

```
**Response Examples:**
```json
[
  {
    "username": "stephenmieeyttecgruer",
    "email": "chroroll@email.com",
    "commit_count": 94
  },
  {
    "username": "colnbunieyuy",
    "email":"autoroll@email.com",
    "commit_count": 42
  },
  {
    "username": "colnbunuy",
    "email": "tch@google.com",
    "commit_count": 6
  },
  {
    "username": "colnbun",
    "email": "blul@chromium.org",
    "commit_count": 6
  },
  {
    "username": "stmcgrer",
    "email": "eyrgruer@chromium.org",
    "commit_count": 4
  }
]
```


#### POST localhost:8080/auth/login

- Get Commits From Date Given
  **Request Sample**
```shell
curl -X 'GET' \
  'http://localhost:8082/api/v1/commits/since?since=07-26-2024' \
  -H 'accept: application/json'

```
**Response Examples:**
```json
[
  {
    "id": 555,
    "message": "Re-enable flaky test AttributionUponAttestationsLoading\n\nThis cl fixed the attestations file version due to recent pre-installed\nattestations change.\n\nWe also registered the source using attributionsrc instead of navigation\nto hopefully reduce the test time.\n\nBug: 355215758\nChange-Id: I9a511c0f0dc2b77b45a0d66594ea5c76701dfd48\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740360\nReviewed-by: Andrew Paseltiner <apaseltiner@chromium.org>\nCommit-Queue: Nan Lin <linnan@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333529}",
    "author": "Nan Lin",
    "author_email": "linnan@chromium.org",
    "date": "2024-07-26T15:08:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f20f839616dc2eadb38af2b4b25ab2f9319c806f"
  },
  {
    "id": 556,
    "message": "Mark long-press-drag-drop-touch-editing test as slow\n\nBug: 354579691\nChange-Id: I9e5432d4a3ebb87dfb2cc7571bb78386481bcaf7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738755\nReviewed-by: Robert Flack <flackr@chromium.org>\nCommit-Queue: Kevin Ellis <kevers@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333528}",
    "author": "Kevin Ellis",
    "author_email": "kevers@google.com",
    "date": "2024-07-26T15:05:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/138e9c403ed7c57dcdec9a1f5eb1c7cc108ed93a"
  },
  {
    "id": 557,
    "message": "Reset PrefChangeRegistrar in UpgradeDetector::Shutdown.\n\nUpgradeDetector::pref_change_registrar_ keeps a raw pointer to the\nlocal state's PrefService. To avoid triggering detection of a leaked\ndangling pointer, reset it during UpgradeDetector::Shutdown.\n\nBug: 355610030\nChange-Id: I40bbdc4005205f35e345508de8a10a1e3801b20f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5742343\nReviewed-by: Greg Thompson <grt@chromium.org>\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nCr-Commit-Position: refs/heads/main@{#1333527}",
    "author": "Jan Keitel",
    "author_email": "jkeitel@google.com",
    "date": "2024-07-26T15:03:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2bfc82bd28e6b99c1c646d7577f4630cdbb02f35"
  },
  {
    "id": 558,
    "message": "spanification: Add `#pragma allow_unsafe_buffers` to ui/*\n\nSee `//docs/unsafe_buffers.md`\n\nThis is a preparation to fix each files.\nThis CL has no behavior changes.\n\nThis patch was fully automated using script:\nhttps://paste.googleplex.com/5614491201175552\n\nNote that in patchset2, change to:\n/build/config/unsafe_buffers_paths.txt\nwas reverted. Indeed, running too many (~3) CQ run touching this file is\nmaking the builder cache much slower. I will bundle every change to this\nfile in a subsequent CL. I will limit myself to 1-2 CQ run per day.\n\nSee internal doc about it:\nhttps://docs.google.com/document/d/1erdcokeh6rfBqs_h0drHqSLtbDbB61j7j3O2Pz8NH78/edit?resourcekey=0-hNe6w1hYAYyVXGEpWI7HVA&tab=t.0\n\nAX-Relnotes: n/a.\nBug: 40285824\nChange-Id: I008d48abd0b756a5c93a6c937b00679e5d2194f1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5717833\nReviewed-by: Nico Weber <thakis@chromium.org>\nAuto-Submit: Arthur Sonzogni <arthursonzogni@chromium.org>\nCommit-Queue: Arthur Sonzogni <arthursonzogni@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333526}",
    "author": "Arthur Sonzogni",
    "author_email": "arthursonzogni@chromium.org",
    "date": "2024-07-26T15:00:54+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/24d53e3e9a861c76d0e1626522c93e371f5d9223"
  },
  {
    "id": 559,
    "message": "Adjust destruction order of ProfileImpl members.\n\nPrefChangeRegistrar keeps a raw pointer to the PrefService it is\nobserving. As a result, it should be destroyed prior to the PrefService.\n\nThis CL does not yet modify the dangling annotation of the\nPrefRegistrar member because there seem to be (several) additional\ninstances in which the member becomes dangling.\n\nBug: 355610030\nChange-Id: I9f516e983b4c757b78d2dc302e0236ca1a12ecce\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743593\nReviewed-by: David Roger <droger@chromium.org>\nReviewed-by: Dominic Battré <battre@chromium.org>\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nCr-Commit-Position: refs/heads/main@{#1333525}",
    "author": "Jan Keitel",
    "author_email": "jkeitel@google.com",
    "date": "2024-07-26T15:00:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/bcb1754d6dd82ee448b3ce6ede8755c314959de6"
  },
  {
    "id": 560,
    "message": "Roll ios_internal from 501e67df31c7 to 5e4d598fc9be\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/501e67df31c7..5e4d598fc9be\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,joemerramos@google.com,rkgibson@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: I74be47390446ff9a9d7efcbdfe63b0c5e223fb48\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5742522\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333524}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-26T14:57:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/719ca2d9993093caa31282269af19be49b3f9993"
  },
  {
    "id": 561,
    "message": "[lensoverlay] Show searchbox in lens search bubble.\n\nSelecting matches in the searchbox is non functional, however the\nsearchbox input and dropdown are functional. Feature flag is in\nlens_features.cc.\n\nScreencast: screencast/cast/NTQ1NTE0OTcxMDQ0MjQ5Nnw3ZjIxZTg3Mi1hYg\n\nBug: 328632272\nChange-Id: I6b37ff47b66f1ca14c2b8f9933262f685ab8d9cb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5463167\nReviewed-by: Moe Ahmadi <mahmadi@chromium.org>\nCommit-Queue: Nihar Majmudar <niharm@google.com>\nCr-Commit-Position: refs/heads/main@{#1333523}",
    "author": "Nihar Majmudar",
    "author_email": "niharm@google.com",
    "date": "2024-07-26T14:56:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9c8e9823e7586361bec781dec3c40dad050e6399"
  },
  {
    "id": 562,
    "message": "Roll Chrome Linux PGO Profile\n\nRoll Chrome Linux PGO profile from chrome-linux-main-1721973548-72ec7edd41800a61cf7fbce7b7d39033ee7c32d2-321aaaf7ae9ca72b98adf17ed74aafafaf276da2.profdata to chrome-linux-main-1721995131-f4680ba138ea4278f4a631b7af1f07ed8e5366c5-1a765d8f98899564966915d2451bc73189822e5d.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-linux-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I8df00fc1c7e8cfac386a77c4a0ca5f9a40ce32f9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743422\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333522}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-26T14:56:01+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/72cb923a154ef11d9417e18b0f8754c59e934adf"
  },
  {
    "id": 563,
    "message": "Append OAuth client ID to IssueToken's request headers\n\nX-OAuth-Client-ID: <OAuth2 client ID>\n\nFixed: b:341371239\nChange-Id: I04e4b6cefbba5cf9563215ead7c46a82e94d0a70\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741345\nAuto-Submit: Alex Ilin <alexilin@chromium.org>\nReviewed-by: David Roger <droger@chromium.org>\nCommit-Queue: David Roger <droger@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333521}",
    "author": "Alex Ilin",
    "author_email": "alexilin@chromium.org",
    "date": "2024-07-26T14:51:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b4ced426ad28fe4336a1f3fc809daeb206654edd"
  },
  {
    "id": 564,
    "message": "Cleanup TrackingProtectionOnboardingDelegate\n\nThis CL removes the TrackingProtectionOnboardingDelegate, as we don't\nneed browser sides methods after the recent announcement.\n\nBug: b:355002866\n\nChange-Id: I2868d007c6b3a3343622eddafe2ef821a5e3e7a9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738417\nReviewed-by: Andrey Zaytsev <andzaytsev@google.com>\nCommit-Queue: Aashna Sheth <aashnas@google.com>\nReviewed-by: Abe Boujane <boujane@google.com>\nCr-Commit-Position: refs/heads/main@{#1333520}",
    "author": "Aashna Sheth",
    "author_email": "aashnas@google.com",
    "date": "2024-07-26T14:50:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a7443d0942350783db65e3d4c13027f1c36ea5e2"
  },
  {
    "id": 565,
    "message": "Roll Perfetto from e9b88b1f4968 to 439b00c39be4 (2 revisions)\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/e9b88b1f4968..439b00c39be4\n\n2024-07-26 ivankc@google.com Merge \"bigtrace: Add headless service minikube worker and modify cluster script\" into main\n2024-07-26 ivankc@google.com Merge \"bigtrace: Buffer results from workers instead of writing directly\" into main\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-chromium-autoroll\nPlease CC perfetto-bugs@google.com,primiano@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-perfetto-rel\nBug: None\nTbr: perfetto-bugs@google.com\nChange-Id: I33faf74a756c3ec6fa9ea2b599408dab5d2c5345\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743419\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333519}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-26T14:50:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fc0061d8dd6dcfc042a7000a133d32b0b4001b17"
  },
  {
    "id": 566,
    "message": "🧇 Fix crash when \"More\" is clicked multiple times\n\nIf the user manages to click the \"More\" button on the choice screen\nmore than once before we scroll to the end and disable it, we end\nup attempting to run a base::OnceCallback more than once, which\ncauses a crash. With this fix, we will ignore the subsequent calls\nto the callback, which only results in not logging the fact that\n\"More\" was tapped more than once per choice screen.\n\nBug: 353573024\nChange-Id: Id6bb06b640c38455a2743ac82eef98d2da01de53\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741468\nCommit-Queue: Nicolas Dossou-Gbété <dgn@chromium.org>\nAuto-Submit: Nicolas Dossou-Gbété <dgn@chromium.org>\nReviewed-by: David Roger <droger@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333518}",
    "author": "Nicolas Dossou-Gbete",
    "author_email": "dgn@chromium.org",
    "date": "2024-07-26T14:48:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9386833aaf2e29dfa5407fc1a381ce29d9849593"
  },
  {
    "id": 567,
    "message": "Fix test flake in animation-event-destroy-renderer.\n\nThe ?animationiteration variant of the test could fail due to a stall\non the testing machine causing to animation to be finished before\nanimationiteration is fired. An animationiteration event does not fire\non the last iteration. Switched to a long duration animation with\nsetting the current time to trigger each event.\n\nTested locally with hundreds of repeats in virtual/threaded.\n\nBug: 355327241\nChange-Id: I823ea9c3cb22c299e2f8c78ed0e051305827f404\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740864\nReviewed-by: Robert Flack <flackr@chromium.org>\nCommit-Queue: Kevin Ellis <kevers@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333517}",
    "author": "Kevin Ellis",
    "author_email": "kevers@google.com",
    "date": "2024-07-26T14:46:33+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/546b4c5975d5a31c3638e049f3ee6e862185b3ab"
  },
  {
    "id": 568,
    "message": "Roll ANGLE from 48d19755a37e to a0a832de0a0c (7 revisions)\n\nhttps://chromium.googlesource.com/angle/angle.git/+log/48d19755a37e..a0a832de0a0c\n\n2024-07-26 geofflang@chromium.org Revert \"GL: Forward client-side arrays to the driver when possible\"\n2024-07-26 angle-autoroll@skia-public.iam.gserviceaccount.com Roll vulkan-deps from 49b5e420b19a to 068847956e41 (2 revisions)\n2024-07-26 angle-autoroll@skia-public.iam.gserviceaccount.com Roll Chromium from 15e30105c06e to 2e8f18571b3e (242 revisions)\n2024-07-26 mark@lunarg.com Tests: Add Shovel Knight Pocket Dungeon trace\n2024-07-25 angle-autoroll@skia-public.iam.gserviceaccount.com Manual roll Chromium from 99513aa4b41c to 15e30105c06e (813 revisions)\n2024-07-25 geofflang@chromium.org GL: Forward client-side arrays to the driver when possible\n2024-07-25 angle-autoroll@skia-public.iam.gserviceaccount.com Manual roll Chromium from 834617e372fd to 99513aa4b41c (90 revisions)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/angle-chromium-autoroll\nPlease CC angle-team@google.com,solti@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in ANGLE: https://bugs.chromium.org/p/angleproject/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86\nBug: None\nTbr: solti@google.com\nChange-Id: Ie9dcc8a476184d81d64e0591f999fc24a12e1590\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743417\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333516}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-26T14:44:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c477e74aec2b9257a9f1fa6893976cf5e6fd340e"
  },
  {
    "id": 569,
    "message": "Delete Full3pcd Clank Onboarding Logic\n\nBug: 355002762\nChange-Id: I299f3350c6de77c60385337d7684fffc4a9d44b3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740817\nAuto-Submit: Abe Boujane <boujane@google.com>\nReviewed-by: Calder Kitagawa <ckitagawa@chromium.org>\nReviewed-by: Matt Dembski <dembski@google.com>\nCommit-Queue: Calder Kitagawa <ckitagawa@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333515}",
    "author": "Abe Boujane",
    "author_email": "boujane@google.com",
    "date": "2024-07-26T14:40:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/089fe853bf036a4d7d64290c71923e5d5fe88307"
  },
  {
    "id": 578,
    "message": "[AutoReload] Add url paramter for autoreload\n\nThis CL appends a url paramter to the gaia url when an automatic reload of the authentication flow is executed.\nAs of now no specific functionality utilizes this url parameter but it will be used for future debugging purposes and for future proofing.\n\nBug: b/352025280\nChange-Id: I36cdf534ea9c54e6288374fa8909f4cd54223c3a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5729977\nCommit-Queue: Aya Elgendy‎ <ayag@chromium.org>\nReviewed-by: Maksim Ivanov <emaxx@chromium.org>\nReviewed-by: Andrey Davydov <andreydav@google.com>\nCr-Commit-Position: refs/heads/main@{#1333506}",
    "author": "Aya El Gendy",
    "author_email": "ayag@chromium.org",
    "date": "2024-07-26T14:07:08+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/589fb91ec454bb08f2309f40bd7a0141784f1bfa"
  },
  {
    "id": 570,
    "message": "Roll Chrome Win ARM64 PGO Profile\n\nRoll Chrome Win ARM64 PGO profile from chrome-win-arm64-main-1721973548-f166d4d2c53038380101e46c634993efba9bca99-321aaaf7ae9ca72b98adf17ed74aafafaf276da2.profdata to chrome-win-arm64-main-1721995131-f779e870da14a04a56a23f0dee44cb6c7e21bff7-1a765d8f98899564966915d2451bc73189822e5d.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I8185782182109799c875a002baa01dac40a9697f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743285\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333514}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-26T14:38:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/192a406ccc3f73be2a785c3da12cdfb784ad414c"
  },
  {
    "id": 571,
    "message": "Extend safe-browsing-hash-prefix flag\n\nWebView has not rolled out yet.\n\nFixed: 355308112\nChange-Id: I9b9c4e5ac519bc24342e9916ac9611fedbf04b5f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740507\nAuto-Submit: Xinghui Lu <xinghuilu@chromium.org>\nReviewed-by: thefrog <thefrog@chromium.org>\nCommit-Queue: thefrog <thefrog@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333513}",
    "author": "Xinghui Lu",
    "author_email": "xinghuilu@chromium.org",
    "date": "2024-07-26T14:35:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4ef6847efb49a366322c2ce3bfc95ae7e876ac67"
  },
  {
    "id": 572,
    "message": "Switch riscv64 build to use stable r27 NDK\n\nBug: 353699308\nTest: CQ\nChange-Id: Ic87b2877e69ef6936cbb9043ef28053524a4581a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740908\nReviewed-by: Andrew Grieve <agrieve@chromium.org>\nReviewed-by: Nico Weber <thakis@chromium.org>\nAuto-Submit: Prashanth Swaminathan <prashanthsw@google.com>\nCommit-Queue: Nico Weber <thakis@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333512}",
    "author": "Prashanth Swaminathan",
    "author_email": "prashanthsw@google.com",
    "date": "2024-07-26T14:32:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/dc1044af4484b79db65f3cfb53f4b488f55e70d7"
  },
  {
    "id": 573,
    "message": "Field trial config for AutofillI18nAddressModelNewCountries.\n\nBug: b/41489892\nChange-Id: I45a7b1f892adf26670ea2a38213b6699d73d9fa8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741284\nReviewed-by: Jan Keitel <jkeitel@google.com>\nCommit-Queue: Norge Vizcay <vizcay@google.com>\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nAuto-Submit: Norge Vizcay <vizcay@google.com>\nCr-Commit-Position: refs/heads/main@{#1333511}",
    "author": "Norge Vizcay",
    "author_email": "vizcay@google.com",
    "date": "2024-07-26T14:31:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/419c75b3b4c79e1095f516e9ee12e21e2a39c992"
  },
  {
    "id": 574,
    "message": "[Mullhet] Split kGrouped MatchedFormType.\n\nBefore this CL, PasswordManager.PotentialBestMatchFormType didn't\ndifferentiate between grouped Android and other credentials.\nThis CL splits the MatchedFormType::kGrouped bucket into 2\nseparate buckets. The existing bucket gets deprecated.\n\nBug: 332673987\nChange-Id: I461b9e752c7f10cd3bcabcb62276a27719234889\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735375\nReviewed-by: Vasilii Sukhanov <vasilii@chromium.org>\nCommit-Queue: Timofey Chudakov <tchudakov@google.com>\nCr-Commit-Position: refs/heads/main@{#1333510}",
    "author": "Timofey Chudakov",
    "author_email": "tchudakov@google.com",
    "date": "2024-07-26T14:29:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/554c994dd740d5345ae6f279a4ef258dfe5a6b54"
  },
  {
    "id": 575,
    "message": "DocumentRules: Handle base URL change in ScriptLoader\n\nMoves some of the logic handling base URL changes to ScriptLoader, to\nallow it to update |speculation_rule_set_| to the new rule set\ncreated with the updated base URL.\n\nBug: 353099744\nChange-Id: I3e0ec98f9a57ccac98555a21d9fcf03e913506b9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738414\nReviewed-by: Jeremy Roman <jbroman@chromium.org>\nCommit-Queue: Adithya Srinivasan <adithyas@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333509}",
    "author": "Adithya Srinivasan",
    "author_email": "adithyas@chromium.org",
    "date": "2024-07-26T14:26:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/65a74c3b16d7638e62a4e16fb5676b191265c195"
  },
  {
    "id": 576,
    "message": "Roll devtools-internal from 894b13578bf0 to 5425428d1f4b (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/894b13578bf0..5425428d1f4b\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/devtools/devtools-frontend.git/+log/dee5906e292296195923cfc3f69cc6cdb8c58829..2c1d92ad7b1a4b5247b09ae1709a2d99da76e6af\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: Icaf4cb9356e17ba0e9e7138238b9ea4bb83debca\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5742968\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333508}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-26T14:23:43+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4b29d609a8be45f439bb2ea7d370ff1ebcb61ac6"
  },
  {
    "id": 577,
    "message": "[Passwords] Extend expiring histograms.\n\nFixed: 350504411, 354713400, 354713638, 354713483\nChange-Id: Ifc1364dd4788912fe26d793ffbc00b45d1c0465c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743615\nCommit-Queue: Maria Kazinova <kazinova@google.com>\nAuto-Submit: Maria Kazinova <kazinova@google.com>\nReviewed-by: Ioana Pandele <ioanap@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333507}",
    "author": "Maria Kazinova",
    "author_email": "kazinova@google.com",
    "date": "2024-07-26T14:21:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9b07a944b4784ac593f8d6deb63b270ef9f01d9d"
  },
  {
    "id": 579,
    "message": "[Android] Rename address sheet keyboard accessory actions.\n\nBug: 327838324\nChange-Id: I96d8476e8f0872dfb8d7e5ca1c9d6cbda59c3ca9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741546\nCommit-Queue: Timofey Chudakov <tchudakov@google.com>\nReviewed-by: Friedrich Horschig <fhorschig@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333505}",
    "author": "Timofey Chudakov",
    "author_email": "tchudakov@google.com",
    "date": "2024-07-26T14:01:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3a0e070ba2732e0664f64de691d0af941a5b1ed0"
  },
  {
    "id": 580,
    "message": "[Gardener] Disable flaky UkmBrowserTest\n\nUkmBrowserTest.RegularPlusIncognitoCheck has caused failure on\nandroid-oreo-x86-rel bot.\nUkmBrowserTest.InstallDateProviderPopulatesSystemProfile has caused\nfailure on android-bfcache-rel bot.\n\nBug: 355609356\nChange-Id: I6bfec433d155f12fe20fe4df2134a23510509cd1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743594\nAuto-Submit: Anna Tsvirchkova <atsvirchkova@google.com>\nCommit-Queue: Anna Tsvirchkova <atsvirchkova@google.com>\nReviewed-by: Adem Derinel <derinel@google.com>\nOwners-Override: Anna Tsvirchkova <atsvirchkova@google.com>\nCommit-Queue: Adem Derinel <derinel@google.com>\nCr-Commit-Position: refs/heads/main@{#1333504}",
    "author": "Anna Tsvirchkova",
    "author_email": "atsvirchkova@google.com",
    "date": "2024-07-26T14:01:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e4622e3f77d6e92c40edd01ff2e4c981985ef354"
  },
  {
    "id": 581,
    "message": "[ios] Use a single gn variable to control use of copy_xctrunner_app\n\nBuilding XCTests requires copying XCTRunner.app which is part of the iOS\nSDK (and shipped inside Xcode.app) into the application. When using the\nsystem installation of Xcode, those files are outside of the checkout.\nUsing absolute path works with gn, however the distributed build system\nrequires that all paths are relative to the checkout. This is faked by\nusing symbolic links to the SDK inside of Xcode. Additionally, each build\ndirectory may use a distinct version of Xcode (e.g. to build with beta),\nso the symlink needs to be present in the $root_build_dir. However, when\ndoing that, we need to list inputs pointing to file in $root_build_dir,\nand gn requires all failes in $root_build_dir to be listed as outputs of\nanother target.\n\nTo fullfill all of those requirements, we 1. create symlinks pointing to\nthe SDK files in Xcode, 2. declare a target listing the files as outputs\n(the target is a script that does nothing, it only pretends to created\nthe files but they already exists).\n\nThis works, but results in some files in $root_build_dir being links to\nfiles outside of the build directory. Running `ninja -t clean` will try\nto delete those files breaking Xcode installation. The recommendation is\nto use `gn clean` or `ninja -t cleandead` instead.\n\nAdd a single variable to control whether the workaround is needed.\n\nBug: none\nChange-Id: I0686232ebf59e9f39dec7cd73e464969479f5122\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741540\nCommit-Queue: Sylvain Defresne <sdefresne@chromium.org>\nReviewed-by: Rohit Rao <rohitrao@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333503}",
    "author": "Sylvain Defresne",
    "author_email": "sdefresne@chromium.org",
    "date": "2024-07-26T13:56:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/948b92504b84cf9b3f091052e708ace2cfae5651"
  },
  {
    "id": 582,
    "message": "[sync] Remove CreateDataTypeManager\n\nNo behavioral changes outside tests, as it always instantiates\nDataTypeManagerImpl.\n\nIn tests, before this class, a subclass was instantiated for the purpose\nof accessing some internal state. Instead, SyncEngine can be used for\nsimilar purposes, and everything else isn't externally visible and\narguably shouldn't be verified in tests.\n\nOne benefit is that SyncApiComponentFactory has a better-defined\nscope, which is dealing with SyncEngine instances. A TODO is added\nto find a better name for this class and make it less abstract.\n\nChange-Id: Ia54821245f07f09c49bb0c3d5dc595d1ac61bf0a\nBug: 335688372\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741644\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Mikel Astiz <mastiz@chromium.org>\nReviewed-by: Marc Treib <treib@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333502}",
    "author": "Mikel Astiz",
    "author_email": "mastiz@chromium.org",
    "date": "2024-07-26T13:55:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d66d47c65b5180387e321d05bffcf37be1d9112a"
  },
  {
    "id": 583,
    "message": "[Profiles] Profile picker no longer navigates in browser being destroyed\n\nBug: 40064092, 40242414\nChange-Id: Id8283b435a99254788225748800d7fec409fb9c6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741701\nReviewed-by: Greg Thompson <grt@chromium.org>\nCommit-Queue: David Roger <droger@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333501}",
    "author": "David Roger",
    "author_email": "droger@chromium.org",
    "date": "2024-07-26T13:48:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1cd71739a1661436a24c9b8ea057dc9061e73ef0"
  },
  {
    "id": 584,
    "message": "[High5] ActiveSessionAuthController unittests\n\nWe add several unittests that test the behavior of\n`ActiveSessionAuthController`. We assert that it behaves correctly in\nthe case of correct and wrong password/pin inputs, and in the case of\ncanceling the dialog.\n\nBug: b:352238958, b:348326316\nChange-Id: I141d45f932ad9884253480e578c413ec61d948ab\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735972\nReviewed-by: Xiyuan Xia <xiyuan@chromium.org>\nReviewed-by: Maksim Ivanov <emaxx@chromium.org>\nReviewed-by: Hardik Goyal <hardikgoyal@chromium.org>\nCommit-Queue: Elie Maamari <emaamari@google.com>\nCr-Commit-Position: refs/heads/main@{#1333500}",
    "author": "Elie Maamari",
    "author_email": "emaamari@google.com",
    "date": "2024-07-26T13:41:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3d5950913dbbd130539cca48ada2812498e5cf48"
  }
]
```

#### GET http://localhost:8082/api/v1/commits
- Get Commits for for specified repository

```shell
curl -X 'GET' \
  'http://localhost:8082/api/v1/commits/chromium' \
  -H 'accept: application/json'
```

**Response Examples:**
```json
[
  {
    "id": 1,
    "message": "[Android] Use @JniType for list of bottom sheet credentials.\n\nBefore this CL, all password bottom sheet credentials were\nstored in an array. This CL modifies this logic by passing a\nlist of credentials directly from the C++ to the Java side.\n\nBug: 327838324\nChange-Id: I87b37eea4ca23f029fa04081a2b59788e8b2de28\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5730054\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Friedrich Horschig <fhorschig@chromium.org>\nReviewed-by: Jan Keitel <jkeitel@google.com>\nCommit-Queue: Timofey Chudakov <tchudakov@google.com>\nCr-Commit-Position: refs/heads/main@{#1331708}",
    "author": "Timofey Chudakov",
    "author_email": "tchudakov@google.com",
    "date": "2024-07-23T14:31:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cc20ef5e0ee580a8bbb75f38dc0bef4c2ed68095"
  },
  {
    "id": 2,
    "message": "Replace std::unique_ptr<T[]> with HeapArray for: gpu/command_buffer/service/program_manager.cc\n\nWe need to replace occurrences of std::unique_ptr<T[]> with base::HeapArray<T>.\nUsing std::unique_ptr<T[]> does not automatically preserve the size of the\nallocation. This forces a need for ad-hoc bounds checks,\nleading to bugs which attackers use to compromise our users.\n\nhttps://docs.google.com/document/d/1YsPR8GoN8VTP1ABKCISaQkuBif1Cn80cTxTjsM8QT4s/edit\n\nFixed: 326459213, 326458934, 326458734\nChange-Id: I85e54039c6b68488e75124978a4fccffadb39482\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5730313\nCommit-Queue: Geoff Lang <geofflang@chromium.org>\nAuto-Submit: Ari Chivukula <arichiv@chromium.org>\nReviewed-by: Geoff Lang <geofflang@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331707}",
    "author": "Ari Chivukula",
    "author_email": "arichiv@chromium.org",
    "date": "2024-07-23T14:29:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c45a06670e214e51b65b7ed02b76965067e31e76"
  },
  {
    "id": 3,
    "message": "Determine User Bypass UI via TP content setting feature\n\nSee this flowchart:\nhttps://screenshot.googleplex.com/6AxQc4cY2jabTR6.png\nTL;DR: We want to show the new UI iff the user is controlling the TP\ncontent setting via UB. Showing the old UI with the TP content setting\nis misleading, as is showing the new UI with the COOKIES content\nsetting. The easiest way to accomplish this is to have the feature\ncontrolling the content setting also control the respective UX.\n\nBug: b:334970346\nChange-Id: I1cf5bb2a89b59495d3f08c29bcb171065c915c22\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5721315\nReviewed-by: Kevin Graney <kmg@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Christian Dullweber <dullweber@chromium.org>\nCommit-Queue: Fiona Macintosh <fmacintosh@google.com>\nCr-Commit-Position: refs/heads/main@{#1331706}",
    "author": "Fiona Macintosh",
    "author_email": "fmacintosh@google.com",
    "date": "2024-07-23T14:28:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/aa6069dd76a17ddf45581904de0dd00ed32187d5"
  },
  {
    "id": 4,
    "message": "[Autofill] Invoke voting callbacks in ~BAM()\n\nBefore this CL, the voting callbacks were only called in BAM::Reset(),\nnot in ~BAM(). The argument was that AutofillCrowdsourcingManager\nis destroyed in ~BAM() anyway, but that's not the case anymore since\ncrbug.com/40248764.\n\nBug: 40248764, 354649269\nChange-Id: I29a71db44eae8b1c1c64841fe7bdd2a304a563d6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5725359\nReviewed-by: Dominic Battré <battre@chromium.org>\nCommit-Queue: Christoph Schwering <schwering@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331705}",
    "author": "Christoph Schwering",
    "author_email": "schwering@google.com",
    "date": "2024-07-23T14:28:09+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7c4c65df61129b8d73645feeac8031e6fa4a0e31"
  },
  {
    "id": 5,
    "message": "Remove the kJelly feature flag in ash/components/arc\n\nFeature is launched and can be removed.\n\nThis CL was uploaded by git cl split.\n\nR=yhanada@chromium.org\n\nBug: None\nChange-Id: Iaef650ab915e20cf1f4fa86649b0776ef8c8fe42\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5731982\nAuto-Submit: Sean Kau <skau@chromium.org>\nReviewed-by: Yuichiro Hanada <yhanada@chromium.org>\nCommit-Queue: Yuichiro Hanada <yhanada@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331704}",
    "author": "Sean Kau",
    "author_email": "skau@chromium.org",
    "date": "2024-07-23T14:22:07+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/03422d897bffe9df71468f29208bf6c2f7fa4c5b"
  },
  {
    "id": 6,
    "message": "Support fixed-span construction from modern range\n\nWe have a ctor that works for legacy ranges, but we missed modern\nranges. It is explicit if the input is dynamic, like for legacy ranges.\n\nR=lukasza@chromium.org\n\nBug: 40284755\nChange-Id: Ib3b201267721fe59928b0e5f40150d8d50549da1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5731592\nReviewed-by: Łukasz Anforowicz <lukasza@chromium.org>\nCommit-Queue: danakj <danakj@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331703}",
    "author": "danakj",
    "author_email": "danakj@chromium.org",
    "date": "2024-07-23T14:20:57+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d689b7f7634926139ecdcd93f738617c9359fea0"
  },
  {
    "id": 7,
    "message": "[task] Adjust behavior of align wakeup for battery saver mode.\n\nThis CL tweaks effect of align wakeup\n-  GetMaxHighResolutionInterval in DOMTimer is controlled by GetAlignWakeUpsEnabled()\n- Timer slack on mac is controlled by GetAlignWakeUpsEnabled()\n\nSince align wakeups is to be enabled only in battery saver mode, this\nis a reliable signal.\nOn windows, high resolution timer (ExplicitHighResolutionTimerWin)\nruns independently from GetAlignWakeUpsEnabled() and also needs\ntask leeway.\n\nBug: 40158967\nChange-Id: I60f37b87a0fb10b4f93a61098806808a05bf7791\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5730813\nReviewed-by: Daniel Cheng <dcheng@chromium.org>\nReviewed-by: Francois Pierre Doray <fdoray@chromium.org>\nCommit-Queue: Etienne Pierre-Doray <etiennep@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331702}",
    "author": "Etienne Pierre-doray",
    "author_email": "etiennep@chromium.org",
    "date": "2024-07-23T14:20:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4a90aed3daf7bdec3d4dd7d06be7a93db780de2b"
  },
  {
    "id": 8,
    "message": "Roll ios_internal from e25e6f884d5e to a35dca29a231\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/e25e6f884d5e..a35dca29a231\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC ajuma@google.com,chrome-brapp-engprod@google.com,eic@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: I246271fdb90210540d9a4a69fd23b4bbf57483d8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734132\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331701}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-23T14:15:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cd95afce8d7bf4b578a994ac9acb9c02465463e0"
  },
  {
    "id": 9,
    "message": "Create ConnectorsServiceBase\n\nThis CL creates a new `ConnectorsServiceBase` class in components/ to\nstart moving code to allow its reuse on mobile platforms. The following\nrefactors are done to start supporting URLF on mobile:\n- `GetDMTokenForRealTimeUrlCheck()` and `GetAppliedRealTimeUrlCheck()`\n  are move from `ConnectorsService` to `ConnectorsServiceBase`.\n- Some internal methods/structs of `ConnectorsService` are moved to\n  support the above two URLF methods.\n- `GetPrefs()` methods are added to simplify accessing prefs from\n  components.\n\nBug: b/318842720\nChange-Id: I4d476de0c95cb4a60d0dad228f0465b51923cad6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5730314\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: thefrog <thefrog@chromium.org>\nReviewed-by: Nancy Xiao <nancylanxiao@google.com>\nCommit-Queue: Dominique Fauteux-Chapleau <domfc@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331700}",
    "author": "Dominique Fauteux-Chapleau",
    "author_email": "domfc@chromium.org",
    "date": "2024-07-23T14:15:15+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1c069074ec92d61305edc302c8c6236eeb3ad95f"
  },
  {
    "id": 10,
    "message": "Roll Perfetto from d072b30d58d1 to 5831e80e13fc (1 revision)\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/d072b30d58d1..5831e80e13fc\n\n2024-07-23 ivankc@google.com Merge \"ui: Change cpu labels to use android_cpu_cluster_mapping\" into main\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-chromium-autoroll\nPlease CC perfetto-bugs@google.com,primiano@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-perfetto-rel\nBug: None\nTbr: perfetto-bugs@google.com\nChange-Id: I895a0095092fa457f6bceafaa21c2461d92dd191\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734033\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331699}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-23T14:10:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7a19f1ecd90bb44256ef3d5a943b416c2ec07232"
  },
  {
    "id": 11,
    "message": "Roll Perfetto Trace Processor Win from 077d742d9f97 to d072b30d58d1\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/077d742d9f97..d072b30d58d1\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-trace-processor-win-chromium\nPlease CC chrometto-team@google.com,perfetto-bugs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: perfetto-bugs@google.com\nChange-Id: Id976caf723575cfc084d0c34d0db10b47e4b9007\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5733493\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331698}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-23T14:07:36+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0b1150f5663f20c4cf9711af104705e234c4b216"
  },
  {
    "id": 12,
    "message": "Roll Chrome Android ARM32 PGO Profile\n\nRoll Chrome Android ARM32 PGO profile from chrome-android32-main-1721713346-650cf2482a79ca2d479831eccbfa780ac404bfa5-60b3996af2dea122de6ca00475bc4e81ed95ee76.profdata to chrome-android32-main-1721735477-d2687032cc2bc3da1300271cbfb0b9f7b33a3ccd-cb4015de717cbd121d806e76c88726fc11c90aa7.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-android-arm32-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I9f40e9f324b3a70c64e1f17f8cc75c07162695e4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5732285\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331697}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-23T14:05:53+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ca69aa6639b3e5e9b87aa10df57a163bfa30732d"
  },
  {
    "id": 13,
    "message": "Revert \"[Interop] Coalesced/predicted event targets in untrusted events.\"\n\nThis reverts commit 2c7bc5f3422f4f43cca3f9ebe1300cce86b60249.\n\nCQ_INCLUDE_TRYBOTS=luci.chrome.try:chromeos-betty-chrome\n\nReason for revert: Suspect breaking input.VirtualKeyboardHandWriting tast tests; makes handwriting glyph weird.\n\nOriginal change's description:\n> [Interop] Coalesced/predicted event targets in untrusted events.\n>\n> Only trusted PointerEvents should update the targets in coalesced and predicted event lists.\n>\n> PSA: https://groups.google.com/a/chromium.org/g/blink-dev/c/5UbeICCy4wY\n>\n> Fixed: 353538500\n> Change-Id: I1188f85f6b2ab90c00cfa6fbf935f58593f4fba4\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709870\n> Commit-Queue: Robert Flack <flackr@chromium.org>\n> Auto-Submit: Mustaq Ahmed <mustaq@chromium.org>\n> Reviewed-by: Robert Flack <flackr@chromium.org>\n> Cr-Commit-Position: refs/heads/main@{#1331066}\n\nBug: b/354863416\nChange-Id: I6f44f26f275a8ea880b7b922d502418608bb2e31\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5733635\nCommit-Queue: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCommit-Queue: Eriko Kurimoto <elkurin@chromium.org>\nOwners-Override: Eriko Kurimoto <elkurin@chromium.org>\nBot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nReviewed-by: Eriko Kurimoto <elkurin@chromium.org>\nAuto-Submit: Qijiang Fan <fqj@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331696}",
    "author": "Qijiang Fan",
    "author_email": "fqj@chromium.org",
    "date": "2024-07-23T14:05:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/134086a139f9610b0395f0b50241b15d4f7f5216"
  },
  {
    "id": 87,
    "message": "Validate iterator before dereferencing in media engagement observer\n\nAdds a CHECK(it != _.end()) between it = _.find() and it->_;\n\nBug: 351745839\nChange-Id: I896f2999e5b304c7089e4564def5418378c1909f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738819\nReviewed-by: Ted (Chromium) Meyer <tmathmeyer@chromium.org>\nCommit-Queue: Alex Gough <ajgo@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332605}",
    "author": "Alex Gough",
    "author_email": "ajgo@chromium.org",
    "date": "2024-07-24T22:22:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c40fd5d9d6f124e826a45d05c2fff3171227e7f0"
  },
  {
    "id": 14,
    "message": "[SHoA] Add magic stack dismiss logic\n\n- Dismiss the magic stack when the user visits the Safety Hub page.\n- Dismiss the safe browsing magic stack module when safe browsing is not\ndisabled anymore.\n\n\nBug: 354724065, 324562205\nChange-Id: I6cc14ae1507ffe81e389e9152facea76f9bf60c1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5729757\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Zaina Al-Mashni <zalmashni@google.com>\nCommit-Queue: Rubin Deliallisi <rubindl@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331695}",
    "author": "Rubin Deliallisi",
    "author_email": "rubindl@chromium.org",
    "date": "2024-07-23T14:00:15+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/30f82194812d3ca2a66edbc0fab4d78919902fcd"
  },
  {
    "id": 15,
    "message": "Roll Chrome Win ARM64 PGO Profile\n\nRoll Chrome Win ARM64 PGO profile from chrome-win-arm64-main-1721713346-54c9ac371dfde2dc51b4847a87f0ea6e98a683ee-60b3996af2dea122de6ca00475bc4e81ed95ee76.profdata to chrome-win-arm64-main-1721735477-6b712aa84de200446a70f31a6c25ed2fab6bff0b-cb4015de717cbd121d806e76c88726fc11c90aa7.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I3b64fdc08dd9eb69de37db57a5e979eefeb94c2c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734136\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331694}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-23T13:58:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b02df56e6f045ed49874f0fa3fc1494e13c68e79"
  },
  {
    "id": 16,
    "message": "dEQP bots: Disable reduced libc++ headers\n\n...by looking at a global define that happens to be set when\nbuild_angle_deqp_tests is true. (build_angle_deqp_tests itself\nis not available in buildtools/.)\n\nBug: 354741410,354621206\nChange-Id: Ibb68d52fa701fdca11d688694e927574ca4fcb3d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5733205\nCommit-Queue: Sylvain Defresne <sdefresne@chromium.org>\nReviewed-by: Sylvain Defresne <sdefresne@chromium.org>\nAuto-Submit: Nico Weber <thakis@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331693}",
    "author": "Nico Weber",
    "author_email": "thakis@chromium.org",
    "date": "2024-07-23T13:55:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/04132ee8f8a5b483485580cab5b3f425bf07416c"
  },
  {
    "id": 17,
    "message": "Roll Chrome Mac Arm PGO Profile\n\nRoll Chrome Mac Arm PGO profile from chrome-mac-arm-main-1721728792-96e1c21567117f92c72e37b4e00da234b78f8a70-43dc94dcea8998c81b7dd05513da88e9a59c493c.profdata to chrome-mac-arm-main-1721735477-c04d159a9b9df3f2784bfd2c35dec78fab4f1efa-cb4015de717cbd121d806e76c88726fc11c90aa7.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-mac-arm-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:mac-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I84ef07a23f2aec71c6251d377ccf89e940c3613a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734135\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331692}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-23T13:54:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fa7ac77169ece72263fdee6e2d919765034853f2"
  },
  {
    "id": 18,
    "message": "Roll clank/internal/apps from 10c9872f8356 to 9832e7516a9a (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/10c9872f8356..9832e7516a9a\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,wbjacksonjr@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: wbjacksonjr@google.com\nNo-Try: true\nChange-Id: I134b14bf1ae956243088a509b2316bfea0c0eb15\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5733709\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331691}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-23T13:54:17+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8a11efb7d8a98bf355e88791ea362860df39dde9"
  },
  {
    "id": 19,
    "message": "Move ProfileTokenQuality::disable_randomization_for_testing() to TestApi\n\nBug: 40100455\nChange-Id: Ia6b30f14a8ab5bc21fd8ae35d4e2fed4b951c953\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734113\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nReviewed-by: Jan Keitel <jkeitel@google.com>\nAuto-Submit: Florian Leimgruber <fleimgruber@google.com>\nCommit-Queue: Florian Leimgruber <fleimgruber@google.com>\nCr-Commit-Position: refs/heads/main@{#1331690}",
    "author": "Florian Leimgruber",
    "author_email": "fleimgruber@google.com",
    "date": "2024-07-23T13:52:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b6ecabb73b5bf48ef256c85ad24517092176f0bd"
  },
  {
    "id": 27,
    "message": "Fix crash on shutdown while reporting UMA\n\nUpon destruction, PredictorDatabaseInternal is destroyed but keeps the\ntimer alive, causing an edge case in which the timer triggers a call\nto ReportUMA, which causes a use of db_path_ (which is corrupted in\nthis scenario).\nThis patch makes ReportUMA static and removes this risk.\n\nBug: 352467342\nChange-Id: I37febbf92b28889e6ad956380478234582f5ed21\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5728413\nCommit-Queue: Egor Pasko <pasko@chromium.org>\nReviewed-by: Egor Pasko <pasko@chromium.org>\nReviewed-by: Ryan Sturm <ryansturm@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331682}",
    "author": "m3rbi",
    "author_email": "yuval.meerbaum@island.io",
    "date": "2024-07-23T13:25:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cd075289746dc956f592bed0e1c174fa12260498"
  },
  {
    "id": 492,
    "message": "Reland \"Clear BrowserUserData before destroying Browser\"\n\nThis is a reland of commit 55002611925e501981acf15e3409cff49561108b\nThis reland contains an additional line to reset `features_` as well, to\nfix a lifetime dependency issue exposed by https://crrev.com/c/5612895\n(which will, in turn, be fixed by https://crrev.com/c/5612905).\n\nOriginal change's description:\n> Clear BrowserUserData before destroying Browser\n>\n> The Browser class allows arbitrary classes that inherit from\n> BrowserUserData to attach themselves to a Browser instance. However,\n> BrowserUserData instances are destroyed after Browser's fields are\n> already destroyed. If a BrowserUserData subclass tries to use the\n> Browser in its destructor, this is a use-after-destroy. An example of\n> such an error caught by MSan's use-after-dtor check:\n>\n>     #0 BrowserView::GetBrowserViewForBrowser(Browser const*) chrome/browser/ui/views/frame/browser_view.cc:1103:26\n>     #1 ReadAnythingCoordinator::~ReadAnythingCoordinator() chrome/browser/ui/views/side_panel/read_anything/read_anything_coordinator.cc:176:31\n>     #2 ReadAnythingCoordinator::~ReadAnythingCoordinator() chrome/browser/ui/views/side_panel/read_anything/read_anything_coordinator.cc:161:53\n>     #3 operator() third_party/libc++/src/include/__memory/unique_ptr.h:67:5\n>     #4 reset third_party/libc++/src/include/__memory/unique_ptr.h:278:7\n>     #5 ~unique_ptr third_party/libc++/src/include/__memory/unique_ptr.h:248:71\n>     #6 ~pair third_party/libc++/src/include/__utility/pair.h:63:29\n>     <...>\n>     #18 0x559686690082 in base::SupportsUserData::~SupportsUserData() base/supports_user_data.cc:122:1\n>     #19 0x55968264e074 in Browser::~Browser() chrome/browser/ui/browser.cc:726:1\n>     #20 0x55968264eb53 in Browser::~Browser() chrome/browser/ui/browser.cc:638:21\n>\n> Browser now explicitly destroys BrowserUserData instances as the first\n> step of its destruction, to avoid BrowserUserData instances from using a\n> partially-destroyed Browser.\n>\n> However, this reveals an additional issue in ReadAnythingCoordinator,\n> which registers itself as a BrowserList observer, but never unregisters\n> itself. It is unclear what the semantics should be for something that is\n> both a BrowserUserData and a BrowserList observer; fortunately,\n> ReadAnythingCoordinator does not care about the OnBrowserRemoved()\n> event, so this particular question can be punted down the road...\n>\n> Fixing this second issue requires removing the early return in the\n> ReadAnythingCoordinator destructor to avoid skipping observer\n> unregistrations; unfortunately, this reveals one last bug in\n> ReadAnythingCoordinator: it tries to use SidePanelRegistry from its\n> destructor, but both ReadAnythingCoordinator and SidePanelRegistry are\n> BrowserUserData instances.\n>\n> In an ancestor CL, SupportsUserData has been updated so Browser now\n> consistently appears to have no attached BrowserUserData when clearing\n> BrowserUserData instances, whether that's during destruction or during\n> an explicit `ClearAllUserData()` call. While it is an antipattern for a\n> BrowserUserData's destructor to try using another BrowserUserData,\n> fixing that is out of scope of this CL—so just add a null check for now.\n>\n> Bug: 40222690\n> Change-Id: I3144ca6225f0c739365d065db8bfe1db55db7560\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5590095\n> Reviewed-by: Abigail Klein <abigailbklein@google.com>\n> Reviewed-by: Mike Wasserman <msw@chromium.org>\n> Commit-Queue: Daniel Cheng <dcheng@chromium.org>\n> Reviewed-by: Erik Chen <erikchen@chromium.org>\n> Cr-Commit-Position: refs/heads/main@{#1313680}\n\nBug: 40222690\nChange-Id: Ia6472c5f669c49c9979365b8ebab4f90223ae686\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5623448\nOwners-Override: Daniel Cheng <dcheng@chromium.org>\nReviewed-by: Erik Chen <erikchen@chromium.org>\nCommit-Queue: Daniel Cheng <dcheng@chromium.org>\nAuto-Submit: Daniel Cheng <dcheng@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313824}",
    "author": "Daniel Cheng",
    "author_email": "dcheng@chromium.org",
    "date": "2024-06-12T05:31:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2c6ff93a67bbcbfc936358ba8dbca856a03d1d83"
  },
  {
    "id": 20,
    "message": "Reland \"ash: Move //c/b/ash:zeroconf_printer_detector_fuzzer to printing subdir\"\n\nThis is a reland of commit fff3254cab99523d41ffe914d3e46397d543a936\n\nThe reverted cl didn't update relative path correctly.\nThe diff from the reverted cl can be seen as Patch1 and Patch2 diff.\n\nOriginal change's description:\n> ash: Move //c/b/ash:zeroconf_printer_detector_fuzzer to printing subdir\n>\n> Bug: b:335294351\n> Change-Id: If1bd22cf9133daa12f430cf8001b006169c058b2\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5724778\n> Commit-Queue: Eriko Kurimoto <elkurin@chromium.org>\n> Reviewed-by: Hidehiko Abe <hidehiko@chromium.org>\n> Cr-Commit-Position: refs/heads/main@{#1331574}\n\nBug: b:335294351\nChange-Id: I44ad240370c1d79d3711731fd1beac1712890d66\nCq-Include-Trybots: luci.chromium.try:chromeos-libfuzzer-asan-rel\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5730557\nCommit-Queue: Eriko Kurimoto <elkurin@chromium.org>\nReviewed-by: Hidehiko Abe <hidehiko@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331689}",
    "author": "Eriko Kurimoto",
    "author_email": "elkurin@chromium.org",
    "date": "2024-07-23T13:48:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d75c7a04afce0aab71607f0d33d97051d81bdfce"
  },
  {
    "id": 21,
    "message": "[Quick Delete] Move survey trigger behind a separate flag\n\nMove survey trigger behind a separate flag that is disabled by default.\n\nBug: 353684166\nChange-Id: I96a416ad791e1d5adef1a9e48c7687b94822b793\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5730335\nReviewed-by: Christian Dullweber <dullweber@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Zaina Al-Mashni <zalmashni@google.com>\nCr-Commit-Position: refs/heads/main@{#1331688}",
    "author": "Zaina Al-Mashni",
    "author_email": "zalmashni@google.com",
    "date": "2024-07-23T13:44:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3981f9fa620cb2f805422581ed1ea7edf5aa9df3"
  },
  {
    "id": 22,
    "message": "Roll devtools-internal from 354b363e126b to dedfbac82b9b (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/354b363e126b..dedfbac82b9b\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/devtools/devtools-frontend.git/+log/625c4bcd65120c8143fd4bbf13402a3b052a7abb..7e8f399c9979addea625d96840179422d274ebd8\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: I468b72a3dde4b136ba50561b9a48c62721e89d3d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5733571\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331687}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-23T13:39:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/58c096449ef90425f305e391377f5e6462237026"
  },
  {
    "id": 23,
    "message": "[Passwords] Extend UPM histograms\n\nFixed: 354713473, 354713630, 354713593\nChange-Id: Ie06649f09cfe85ea8dcdac6389afc4829ca65a6c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734092\nReviewed-by: Ivana Žužić <izuzic@google.com>\nAuto-Submit: Ioana Pandele <ioanap@chromium.org>\nCommit-Queue: Ioana Pandele <ioanap@chromium.org>\nCommit-Queue: Ivana Žužić <izuzic@google.com>\nCr-Commit-Position: refs/heads/main@{#1331686}",
    "author": "Ioana Pandele",
    "author_email": "ioanap@chromium.org",
    "date": "2024-07-23T13:37:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ac424c06f399d442fe5cd7236b9e6b91ffdd6f7a"
  },
  {
    "id": 24,
    "message": "Roll Platform Experience Win from 6e4fd80536f5 to c8ae639df9eb (1 revision)\n\nhttps://chrome-internal.googlesource.com/chrome/browser/platform_experience/win.git/+log/6e4fd80536f5..c8ae639df9eb\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/platform-experience-win-chromium\nPlease CC chrome-platform-experience-win@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: chrome-platform-experience-win@google.com\nChange-Id: Id731a29cf84fc09470e7a1492f17b844a4a6cc97\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5733936\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331685}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-23T13:28:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/685f509c0cf85941b9d7a070fd49d30f8c7aac82"
  },
  {
    "id": 25,
    "message": "Roll clank/internal/apps from 869a2da3a058 to 10c9872f8356 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/869a2da3a058..10c9872f8356\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,wbjacksonjr@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: wbjacksonjr@google.com\nNo-Try: true\nChange-Id: Id363f021c02355f3b220fcb479793ea35c1b4a38\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734133\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331684}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-23T13:27:31+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a47782563947c56fff8fbed79729a730c63abf37"
  },
  {
    "id": 26,
    "message": "Rationalization of street-address field in Germany.\n\nBug: 41489892, 354888225\nChange-Id: I99dfb24316fe6b95205f34700b073f5f3c5df14c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5729762\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Norge Vizcay <vizcay@google.com>\nReviewed-by: Dominic Battré <battre@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331683}",
    "author": "Norge Vizcay",
    "author_email": "vizcay@google.com",
    "date": "2024-07-23T13:26:18+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2651af1da5744b47d498bb7d715ab1962a852f78"
  },
  {
    "id": 28,
    "message": "Roll WebRTC from f90a3ad3b334 to 7fe62f25d14a (1 revision)\n\nhttps://webrtc.googlesource.com/src.git/+log/f90a3ad3b334..7fe62f25d14a\n\n2024-07-23 meemetao@gmail.com Reland \"Fix 'Image will be cropped if WindowCapturerWinGdi used'\"\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/webrtc-chromium-autoroll\nPlease CC webrtc-chromium-sheriffs-robots@google.com,webrtc-infra@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in WebRTC: https://bugs.chromium.org/p/webrtc/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: webrtc-chromium-sheriffs-robots@google.com\nChange-Id: I61fc813ebceff9930c91f1b76073ba48a9cc5299\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5733935\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331681}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-23T13:25:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ad9460d9ae30b07882142839ccf1abb1c5f29bb8"
  },
  {
    "id": 29,
    "message": "[Viz] Add memory of compositor GraphiteImageProvider to background dumps\n\nBug: 330806170\nChange-Id: Iada23c52e6ed228b134e71daba738e310195c40d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5729333\nReviewed-by: Stephen Nusko <nuskos@chromium.org>\nReviewed-by: Sunny Sachanandani <sunnyps@chromium.org>\nCommit-Queue: Colin Blundell <blundell@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1331680}",
    "author": "Colin Blundell",
    "author_email": "blundell@chromium.org",
    "date": "2024-07-23T13:21:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/267f39e2e92527640ebb92d68dc323a89fb2ad0b"
  },
  {
    "id": 30,
    "message": "Roll Chrome Win64 PGO Profile\n\nRoll Chrome Win64 PGO profile from chrome-win64-main-1721713346-d08b548ddaa030def1febd5d689eea3913449b78-60b3996af2dea122de6ca00475bc4e81ed95ee76.profdata to chrome-win64-main-1721725046-c0b47e24106c82e8214af74054e72abeb5bf9ff0-8db9268777ca59b438ba0c3f9254cf54b0ff0a21.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:win64-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I9552b34d2d7d282a31a7fce75d2dcc9cfd58bafe\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5733706\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1331679}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-23T13:13:18+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/facdaef1d91467a8c11678fb1496fa0e028f4c56"
  },
  {
    "id": 32,
    "message": "Roll PDFium from 20c007ea4c98 to ba49b29cb699 (12 revisions)\n\nhttps://pdfium.googlesource.com/pdfium.git/+log/20c007ea4c98..ba49b29cb699\n\n2024-07-24 thestig@chromium.org Consistently check for `DeviceType::kPrinter`\n2024-07-24 thestig@chromium.org Add experimental FPDFFont_GetBaseName() API\n2024-07-24 thestig@chromium.org Add enum class ProgressiveDecoder::TransferMethod\n2024-07-24 thestig@chromium.org Change FPDFFont_GetFamilyName() to take/return a size_t\n2024-07-24 thestig@chromium.org Templatize SpanFromFPDFApiArgs()\n2024-07-24 thestig@chromium.org Remove useless CGdiDisplayDriver::StretchDIBits() params\n2024-07-24 thestig@chromium.org Further simplify OutputImage() inside cgdi_plus_ext.cpp\n2024-07-24 tsepez@chromium.org Spanify more of CFGAS_TxtBreak.\n2024-07-24 tsepez@chromium.org Pass spans to GetAlphaWithSrc().\n2024-07-24 thestig@chromium.org Use EmbedderTest::LoadScopedPage() in fpdf_text_embeddertest.cpp\n2024-07-24 tsepez@chromium.org Resolve UNSAFE_TODO() in CFGAS_DefaultFontManager::GetFont()\n2024-07-24 thestig@chromium.org Fix FXDIB_Format::kRgb32 handling when drawing with Skia\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pdfium-autoroll\nPlease CC andyphan@google.com,dhoss@chromium.org,thestig@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in PDFium: https://bugs.chromium.org/p/pdfium/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: chromium:351796783,chromium:353746891,chromium:354025368,chromium:42271176,chromium:42271776\nTbr: andyphan@google.com\nChange-Id: I8fe9e75e54021d9b4c05313ddd3ac4035a09e574\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739536\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332660}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-25T00:05:53+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/315f140879eec7acce594a07e5265036cbe4d0f8"
  },
  {
    "id": 33,
    "message": "DeprecatedLayoutImmediately() -> InvalidateLayout() for DeskBarViewBase.\n\nIt may save several milliseconds (see measurements in linked bug), and\nit addresses technical debt since we're using a deprecated method. The\ntrace is also cleaner and easier to follow afterwards since it reduces\nthe number of layout operations.\n\nBug: b:349872368\nChange-Id: I2830b251ec46a2ae99c3356fadd4eabc0d18164a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5672820\nReviewed-by: Xiyuan Xia <xiyuan@chromium.org>\nCommit-Queue: Eric Sum <esum@google.com>\nCr-Commit-Position: refs/heads/main@{#1332659}",
    "author": "Eric Sum",
    "author_email": "esum@google.com",
    "date": "2024-07-25T00:05:06+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ad5fa21efdf89b1b04b1c70b43892546ab0b404c"
  },
  {
    "id": 61,
    "message": "[lensoverlay] Disable tab contents web view when overlay is showing.\n\nBug: b:341329672\nChange-Id: Ie49e10d1ed8099af4c536345dc80a00fcd845718\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5731630\nReviewed-by: Erik Chen <erikchen@chromium.org>\nCommit-Queue: Bryan Nguyen <nguyenbryan@google.com>\nCr-Commit-Position: refs/heads/main@{#1332631}",
    "author": "Bryan Nguyen",
    "author_email": "nguyenbryan@google.com",
    "date": "2024-07-24T23:14:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/81249ea31cb9adb6eeb6211c24c4d4d736bffe74"
  },
  {
    "id": 34,
    "message": "Roll Skia from 267c7c56c6e1 to 746d444f3efd (3 revisions)\n\nhttps://skia.googlesource.com/skia.git/+log/267c7c56c6e1..746d444f3efd\n\n2024-07-24 johnstiles@google.com Avoid using optional<> for ModuleType.\n2024-07-24 johnstiles@google.com Add Analysis::FindSpecializedArgumentsForCall helper function.\n2024-07-24 jvanverth@google.com Update comment in patheffects.cpp.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/skia-autoroll\nPlease CC jamesgk@google.com,skiabot@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Skia: https://bugs.chromium.org/p/skia/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux-blink-rel;luci.chromium.try:linux-chromeos-compile-dbg;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nCq-Do-Not-Cancel-Tryjobs: true\nBug: None\nTbr: jamesgk@google.com\nChange-Id: Id9c320dd55b173ad5b83c52746db70bf9b1dbc30\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739614\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332658}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-25T00:04:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d3fd174ec3f597570dc18a4bc19557eb29f8dca5"
  },
  {
    "id": 35,
    "message": "[TabGroupSync] Update Everything Menu to use Wrapper Service\n\nTo migrate from using the SavedTabGroupKeyedService directly we update the Everything Menu entry point to use the Wrapper Service which manages the service (TabGroupSyncService or SavedTabGroupKeyedService) we use to query our data from when the everything menu is shown.\n\nChange-Id: Ie1f2afe640933ef080a9c96fb909a4b3f1db2c48\nBug: 350514491\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735344\nReviewed-by: David Pennington <dpenning@chromium.org>\nCommit-Queue: Darryl James <dljames@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332657}",
    "author": "dljames",
    "author_email": "dljames@chromium.org",
    "date": "2024-07-25T00:02:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0bfcb8a25eef0d66b4f97625efda90eae0099a31"
  },
  {
    "id": 36,
    "message": "Re-enable final thumbnail test on all platforms\n\nTest was migrated to kombucha (which has helped stability) in:\nhttps://chromium-review.googlesource.com/c/chromium/src/+/5715213\n\nBug: 40883117\nChange-Id: If56ac4e3d2189e0354a094a2d4cf8f808d1f0a96\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735339\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Dana Fried <dfried@chromium.org>\nCommit-Queue: Alison Gale <agale@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332656}",
    "author": "Alison Gale",
    "author_email": "agale@chromium.org",
    "date": "2024-07-25T00:01:54+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cc65351de89bd4eb44a26233774b81144a4401f4"
  },
  {
    "id": 37,
    "message": "Disable BiddingAndAuctionSerializerTest.SerializeWithSmallRequestSize\n\ntest is flaky.\n\nBug: 355013095\nChange-Id: I548c2c6161bcbc2033a130d8cd844d8bc95476fc\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738820\nAuto-Submit: Evan Stade <estade@chromium.org>\nCommit-Queue: Caleb Raitto <caraitto@chromium.org>\nReviewed-by: Caleb Raitto <caraitto@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332655}",
    "author": "Evan Stade",
    "author_email": "estade@chromium.org",
    "date": "2024-07-24T23:56:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/21e100fe9a2ef551833e783b3ee903c4cb69e7d5"
  },
  {
    "id": 38,
    "message": "Switch TsSection::Parse to consume a span\n\nTsSectionPsi still needs a rewrite to fix the CRC checksum failure.\n\nBug: 40057824\nChange-Id: I2b545bdd925b452e9149e9e07a041029216d913d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5725787\nReviewed-by: Eugene Zemtsov <eugene@chromium.org>\nCommit-Queue: Ted (Chromium) Meyer <tmathmeyer@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332654}",
    "author": "Ted Meyer",
    "author_email": "tmathmeyer@chromium.org",
    "date": "2024-07-24T23:56:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7748da360f9ef34acf58d108406ffeefd24715af"
  },
  {
    "id": 39,
    "message": "Roll V8 from 3311e78008a9 to 6fb1add67b08 (2 revisions)\n\nhttps://chromium.googlesource.com/v8/v8.git/+log/3311e78008a9..6fb1add67b08\n\n2024-07-24 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Version 12.9.43\n2024-07-24 mliedtke@chromium.org [fuzzer][wasm] Deopt fuzzer: Add values to value stack before deopt\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/v8-chromium-autoroll\nPlease CC liviurau@google.com,machenbach@google.com,v8-waterfall-gardener@grotations.appspotmail.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-blink-rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:dawn-linux-x64-deps-rel\nChange-Id: I1e747df2586d770a4eb72ac39a6f88e6ec9c8736\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739340\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332653}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T23:55:34+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/be7d2440bc51a6326ab9bc4be9b042c23bab9de9"
  },
  {
    "id": 40,
    "message": "Fix NPE thrown while retrieving state from strip visibility state\nsupplier\n\nCurrently, the supplier could be null after instantiation only when #destroy is invoked. Following trace from crash, speculating that\nCompositorViewHolder#handleWindowInsetChanged() is invoked soon after\nCompositorViewHolder#shutDown() (that invokes\nStripLayoutHelperManager#destroy()) during an activity destruction for\nsome reason, by which time the supplier is already null.\n\nThis CL removes resetting the supplier to null in #destroy(), so it\nwill continue to hold the last state after #destroy() is invoked.\n\nBug: 354822304\nChange-Id: I524f898942eca987bf957e7c4a7e48ef49d60949\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734892\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Sirisha Kavuluru <skavuluru@google.com>\nCommit-Queue: Aishwarya Rajesh <aishwaryarj@google.com>\nCr-Commit-Position: refs/heads/main@{#1332652}",
    "author": "Aishwarya Rajesh",
    "author_email": "aishwaryarj@google.com",
    "date": "2024-07-24T23:50:05+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0576a2fabf340a9457b888e8517990c40eaff239"
  },
  {
    "id": 41,
    "message": "Roll Chrome Mac Arm PGO Profile\n\nRoll Chrome Mac Arm PGO profile from chrome-mac-arm-main-1721851124-dafd418f852199ec81f2729a576ae2adfacaed88-c113d06ffb53ec81c836f2558e9534f6e66740a3.profdata to chrome-mac-arm-main-1721857641-9a9260f7b0a0f7c46962dcae1a6879940c98e512-c318fc5940761a783848b6a8feafe9024a72ba11.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-mac-arm-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:mac-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I3ccfe0499466777418fb1e40b22dbb33056deb0c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739624\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332651}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T23:49:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b434a15b363697110ad6a72450a4a39e946a5b26"
  },
  {
    "id": 42,
    "message": "[FPF] Add Fingerprinting Protection CNAME alias checking\n\nAdds alias checking for CNAMEs for the FingerprintingProtectionChildNavigationThrottle.\n\nAlias checks are now done once any enabling feature flag is turned on.\n\nMade feature flag binding easier for specific ChildFrameNavigationFilteringThrottles.\n\nAlso binds a feature flag for the alias checking logic specifically for\nFingerprinting Protection: \"UseCnameAliasesForFingerprintingProtectionFilter\", and moved \"SendCnameAliasesToSubresourceFilterFromBrowser\" feature flag to SafeBrowsingChildNavigationThrottle.\n\nBug: 353717668\nChange-Id: I35da1f71c926e4128775c395437401dc5c39d69f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5718729\nReviewed-by: Martin Verde <thesalsa@google.com>\nCommit-Queue: John Kim <johnykim@google.com>\nReviewed-by: Alex Turner <alexmt@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332650}",
    "author": "John Kim",
    "author_email": "johnykim@google.com",
    "date": "2024-07-24T23:46:40+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5abf9081a45deef2846d12311055c5f0ed83759c"
  },
  {
    "id": 43,
    "message": "Add IsActive() to BrowserWindowInterface.\n\nThis method tracks widget activity.\n\nChange-Id: Ibfe6eeae9078ed952f2a17cd40d9969e4f7844e4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735900\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Erik Chen <erikchen@chromium.org>\nReviewed-by: Scott Violet <sky@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332649}",
    "author": "Erik Chen",
    "author_email": "erikchen@chromium.org",
    "date": "2024-07-24T23:46:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8ddfb8b48a36edeb798fe0c648fbd33cc615c5b2"
  },
  {
    "id": 44,
    "message": "Roll src-internal from 960d1a5e5519 to 62b9d31f2498 (2 revisions)\n\nhttps://chrome-internal.googlesource.com/chrome/src-internal.git/+log/960d1a5e5519..62b9d31f2498\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/src-internal-chromium-autoroll\nPlease CC chrome-browser-infra-team,estade@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: chromium:349117847,chromium:349117870,chromium:352025642,chromium:355225022\nTbr: estade@google.com\nChange-Id: I90f1a3cb15fadd4e72b9ec2cd4a31b80f45e8389\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739336\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332648}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-24T23:45:17+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/bffb50e180958f591e258abf6b831545ae0b91ce"
  },
  {
    "id": 45,
    "message": "Roll Depot Tools from 1d5771b0fd3b to b5029f0194c5 (1 revision)\n\nhttps://chromium.googlesource.com/chromium/tools/depot_tools.git/+log/1d5771b0fd3b..b5029f0194c5\n\n2024-07-24 ayatane@chromium.org [git_cl] Add NO_AUTH enum\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/depot-tools-chromium-autoroll\nPlease CC chops-source-team@google.com,ddoman@google.com,gavinmak@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: ddoman@google.com,gavinmak@google.com\nChange-Id: I0f186b894d689f4dec3c05ac2f9b369d510b6b67\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739337\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332647}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T23:44:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f92e7e8f190c380f6f86c9d2fba0a11c785d61f0"
  },
  {
    "id": 46,
    "message": "Roll clank/internal/apps from 486256e47bc2 to c36258f93f4b (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/486256e47bc2..c36258f93f4b\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC aishwaryarj@google.com,chrome-brapp-engprod@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: aishwaryarj@google.com\nNo-Try: true\nChange-Id: I483ebe9f6b8dfb18efaaf6fc585b792da5cfd55c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739818\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332646}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-24T23:39:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/bab512492f1d04023a0fda10404d95391b0f1bf7"
  },
  {
    "id": 47,
    "message": "birch: Tweak suggestion chip icons\n\nUse the \"app\" icon size for file attachments and file suggestions.\nChange app icons to 16x16 for consistency with favicons. Change the\nbackground for favicons to a rounded-rect with 8 pixel corners, instead\nof a circle.\n\nKeep favicon icons at 16x16, since some web sites only provide us that\nresolution and it looks bad when enlarged.\n\nhttps://screenshot.googleplex.com/BW2yV4voDpsGxGq\n\nBug: b:354043357\nTest: visual\nChange-Id: I5aa4bc7c30de18139548900648ce02031da7ee18\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739134\nCommit-Queue: James Cook <jamescook@chromium.org>\nReviewed-by: Matthew Mourgos <mmourgos@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332645}",
    "author": "James Cook",
    "author_email": "jamescook@chromium.org",
    "date": "2024-07-24T23:36:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/58c77bf2b533d2f7b0abc308302a08ec868ec830"
  },
  {
    "id": 48,
    "message": "Roll ANGLE from 84e54d885f87 to 4f498eaa1426 (1 revision)\n\nhttps://chromium.googlesource.com/angle/angle.git/+log/84e54d885f87..4f498eaa1426\n\n2024-07-24 liza@chromium.org WebGPU: Add more format support\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/angle-chromium-autoroll\nPlease CC angle-team@google.com,solti@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in ANGLE: https://bugs.chromium.org/p/angleproject/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86\nBug: None\nTbr: solti@google.com\nChange-Id: Ib24ab627e434ba91bf4415d829708f0298d7f8c4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739373\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332644}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T23:35:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cced2f9208511d9e066daefa0c557df0b75e3a6c"
  },
  {
    "id": 49,
    "message": "updater: fix the needsadmin prefers scenario\n\nThe needsadmin=prefers scenario stopped working after crrev/1325550, and\nthis CL fixes it such that if any error happens when trying to install\nas admin, the \"prefers\" flag will attempt to install as user.\n\nBug: 354847245\nChange-Id: Id26ccc242fed475ed319c59a375897452aa6b1bf\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739333\nAuto-Submit: S Ganesh <ganesh@chromium.org>\nReviewed-by: Sorin Jianu <sorin@chromium.org>\nCommit-Queue: Sorin Jianu <sorin@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332643}",
    "author": "S. Ganesh",
    "author_email": "ganesh@chromium.org",
    "date": "2024-07-24T23:29:40+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fda67236655d2d48cf99d93c530519f6a3859134"
  },
  {
    "id": 50,
    "message": "[Birch] Add SecondaryIconType for Lost Media Items.\n\nThis CL adds the SecondaryIconType field to Lost Media Items, which prepares the item to display secondary icons. There are three types of secondary icons available for Lost Media Items: Audio icon, Video icon, or Video Conference Icon.\n\nBug: b/354043357\nChange-Id: I6d1d8644f910e0173c2f6b89020d2726c4652670\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5732067\nReviewed-by: Sammie Quon <sammiequon@chromium.org>\nReviewed-by: Matthew Mourgos <mmourgos@chromium.org>\nCommit-Queue: Owen Zhang <owenzhang@google.com>\nCr-Commit-Position: refs/heads/main@{#1332642}",
    "author": "Owen Zhang",
    "author_email": "owenzhang@google.com",
    "date": "2024-07-24T23:28:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/815f3699b9d9d1eecd3ff7581c676f30e00c1109"
  },
  {
    "id": 51,
    "message": "iwa: Polish chrome://web-app-internals UI\n\nThis updates the IWA portion of chrome://web-app-internals to not show\nDev Mode headers if dev mode is disabled, and to hide the IWA section\nentirely if dev mode is disabled on ChromeOS, as the only non-dev mode\nUI is force policy refreshes, which aren't applicable on WML.\n\nThis also renames HTML elements to have a more consistent id scheme,\nadds a loading message to the dev mode app list, and makes the list\nauto-update when a new app is installed.\n\nBug: 354941904\nChange-Id: Ibbceab991b6582c4107e8e63b80c36de7b969c55\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734585\nCommit-Queue: Robbie McElrath <rmcelrath@chromium.org>\nReviewed-by: Dibyajyoti Pal <dibyapal@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332641}",
    "author": "Robbie McElrath",
    "author_email": "rmcelrath@chromium.org",
    "date": "2024-07-24T23:28:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d79763cdd531cc7ccd6b53463ca98dc10395a6a4"
  },
  {
    "id": 52,
    "message": "Unify extracode1 behavior between update and download pings\n\nWhen errorcodes set by HTTP status codes are sent back in ping back\nrequests, Update pings that are sent in the same ping back request as the\ncorresponding Download pings end up propagating the same download\nerrorcode, but have an extracode1 set as 0 / unset, even if the\nextracode1 is non-zero in the Download event.\n\nAdds the extracode1 field to the CrxDownloader::Result struct and\npropagates that field through to PingBackRequests\n\nFixed: 350407903\nChange-Id: Idb18de6770e7bf1c6435d0b35d4157708305e31f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735962\nReviewed-by: Joshua Pawlicki <waffles@chromium.org>\nReviewed-by: S Ganesh <ganesh@chromium.org>\nCommit-Queue: S Ganesh <ganesh@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332640}",
    "author": "Micah Hutchens",
    "author_email": "mhutchens@dropbox.com",
    "date": "2024-07-24T23:24:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d1af4757b215ae87447a5d13625dfd874b9a4c72"
  },
  {
    "id": 53,
    "message": "Roll clank/internal/apps from 607214bd9ff4 to 486256e47bc2 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/607214bd9ff4..486256e47bc2\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC aishwaryarj@google.com,chrome-brapp-engprod@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: chromium:345584324\nTbr: aishwaryarj@google.com\nNo-Try: true\nChange-Id: I6762fb5595ccc6dcc5b4ef842e69318bbb571239\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739815\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332639}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-24T23:24:17+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2e7737c1e29fef1e1519203033f0ef1da4f7c35e"
  },
  {
    "id": 54,
    "message": "Roll vulkan-deps from 8e90204125ac to 968ae65f4dbe (2 revisions)\n\nhttps://chromium.googlesource.com/vulkan-deps.git/+log/8e90204125ac..968ae65f4dbe\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/external/github.com/KhronosGroup/SPIRV-Tools/+log/a0817526b8e391732632e6a887134be256a20a18..81a116002b2847c1c86e54f3418a0480b984e6b6\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/vulkan-deps-chromium-autoroll\nPlease CC angle-team@google.com,lokokung@google.com,radial-bots+chrome-roll@google.com,radial-bots@google.com,solti@google.com,webgpu-developers@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86;luci.chromium.try:dawn-linux-x64-deps-rel;luci.chromium.try:win-asan;luci.chromium.try:linux_chromium_cfi_rel_ng\nBug: None\nTbr: lokokung@google.com,radial-bots+chrome-roll@google.com,solti@google.com\nChange-Id: Ib46266197b1e8b52b6312cf1ef09ac7e3fb4897a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738086\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332638}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T23:19:56+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9cab11e77b36741c65d52ed685e418ca919098bb"
  },
  {
    "id": 55,
    "message": "overview: Remove IsOverviewNewFocusEnabled flag (4/N)\n\nRemoves `OverviewFocusableView` override in a couple classes. This is\npreparation to move the class entirely. This will mean if the flag is\ndisabled, these elements won't be focusable. However, the flag has been\nenabled for a while and cannot be disabled and will be removed soon, so\nthis is ok.\n\nTest: ash_unittests\nBug: b/325335020\nChange-Id: Ifb314a7f9c433fe95945283d9a59962e52932e6f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5731955\nReviewed-by: Xiyuan Xia <xiyuan@chromium.org>\nReviewed-by: Daniel Andersson <dandersson@chromium.org>\nCommit-Queue: Sammie Quon <sammiequon@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332637}",
    "author": "Sammie Quon",
    "author_email": "sammiequon@chromium.org",
    "date": "2024-07-24T23:17:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/44f6ad97a32b74a4e09a2f1a7af79938fafe4bcd"
  },
  {
    "id": 56,
    "message": "Enable PRE tests on content_browsertests on Android 14 with Pixel 7\n\nTo align with the current android-pie-arm64-rel\n\nBug: 352811552\nChange-Id: I48bfd421751bfb93f13af0fd0bffc23c457a54de\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738277\nReviewed-by: James Shen <zhiyuans@google.com>\nCommit-Queue: Haiyang Pan <hypan@google.com>\nCr-Commit-Position: refs/heads/main@{#1332636}",
    "author": "Haiyang Pan",
    "author_email": "hypan@google.com",
    "date": "2024-07-24T23:16:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/27d9fb647715ad1478d7e46088ebd1e929e4f1c9"
  },
  {
    "id": 57,
    "message": "usb: Enable unrestricted WebUSB access in isolated apps\n\nEnables the policy-controlled feature \"usb-unrestricted\" for isolated\ncontexts. Isolated apps with this permission can access USB devices\nnormally blocked by WebUSB API.\n\nhttps://chromestatus.com/feature/5106506475503616\n\nIntent to ship:\nhttps://groups.google.com/a/chromium.org/g/blink-dev/c/hXgwRCYta-k/m/NsLkciF1BwAJ\n\nFixed: 40783010\nChange-Id: Ic7310a078be4b1c93072a447237a21843982eccb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5684474\nCommit-Queue: Matt Reynolds <mattreynolds@chromium.org>\nReviewed-by: Kent Tamura <tkent@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332635}",
    "author": "Matt Reynolds",
    "author_email": "mattreynolds@google.com",
    "date": "2024-07-24T23:16:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c5c917d467789775d64d48426f6a6a0eeac08b9d"
  },
  {
    "id": 58,
    "message": "Roll Website from ed19767d81d4 to b3cd959c0e80 (1 revision)\n\nhttps://chromium.googlesource.com/website.git/+log/ed19767d81d4..b3cd959c0e80\n\n2024-07-24 briannorris@chromium.org servo: Reintroduce most schematics, docs\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/website-chromium\nPlease CC dpranke@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Website: https://bugs.chromium.org/p/chromium/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: dpranke@google.com\nChange-Id: I31a8f93a2cd67e2b711d3e986bda9cadb3327634\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5737790\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332634}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T23:15:48+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/746e0e41a9734bc79ff6e4b88f10f5c934294dc9"
  },
  {
    "id": 59,
    "message": "Roll Perfetto from 3177f68ba7e1 to 2961ee996e95 (1 revision)\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/3177f68ba7e1..2961ee996e95\n\n2024-07-24 mayzner@google.com Revert \"Add trace_file metadata table\"\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-chromium-autoroll\nPlease CC perfetto-bugs@google.com,primiano@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-perfetto-rel\nBug: None\nTbr: perfetto-bugs@google.com\nChange-Id: I062933f0e8e3cb177873840d4524b6aec9e8427c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739114\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332633}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T23:15:36+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b1c0163f34ad65b5cf49759cd2caf5ae91426869"
  },
  {
    "id": 60,
    "message": "Followup fix to crrev.com/c/5723176.\n\nAndroid cronet bot (https://ci.chromium.org/ui/p/chromium/builders/ci/android-cronet-mainline-clang-arm64-dbg/37238/overview) fails after that patch, due to the compiler-rt builtins library not being found.\n\nBug: 353167520\nChange-Id: I89c26f03302ffcd161ed73128bbb17d8e5f3c725\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738716\nReviewed-by: Alan Zhao <ayzhao@google.com>\nCommit-Queue: Amy Huang <akhuang@google.com>\nCr-Commit-Position: refs/heads/main@{#1332632}",
    "author": "Amy Huang",
    "author_email": "akhuang@google.com",
    "date": "2024-07-24T23:14:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8ee98a8a0bea04c3225d653531f0ed33fe8f50b6"
  },
  {
    "id": 62,
    "message": "[Price insights clank] Move price insights CPA to the highest position.\n\nThe price insights still has flag behind it. https://source.corp.google.com/h/chrome-internal/codesearch/chrome/src/+/main:chrome/browser/ui/android/toolbar/java/src/org/chromium/chrome/browser/toolbar/adaptive/AdaptiveToolbarFeatures.java;l=233?q=chrome%2Fbrowser%2Fui%2Fandroid%2Ftoolbar%2Fjava%2Fsrc%2Forg%2Fchromium%2Fchrome%2Fbrowser%2Ftoolbar%2Fadaptive%2FAdaptiveToolbarFeatures.java&sq=git:chrome-internal%2Fcodesearch%2Fchrome%2Fsrc@main\n\nBug: 336825251\nChange-Id: Ifc98c754bd1ebf67f868f82222a02f35e3cbeaed\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735899\nReviewed-by: Salvador Guerrero Ramos <salg@google.com>\nReviewed-by: Siddhartha S <ssid@chromium.org>\nCommit-Queue: Qi Bao <qib@google.com>\nCr-Commit-Position: refs/heads/main@{#1332630}",
    "author": "Qi Bao",
    "author_email": "qib@google.com",
    "date": "2024-07-24T23:13:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6c2434598966bac070dc13919a7fcde98ea3f5b3"
  },
  {
    "id": 63,
    "message": "preload-topchrome: resize host before showing preloaded WebUIs\n\nThis patch resizes the top-chrome WebUI host when the host is set for\nWebUIContentsWrapper. This fixes the visual glitches caused by immediate\nsize change after the host is shown.\n\nFor preloaded WebUIs, the resize event ResizeDueToAutoResize originates\nasynchronously from viz surface when the WebUI is attached to a\nviews::Widget. Effectively this means the resize would happen after the\nShowUI() call.\n\nThis patch sets the host size to the render frame size. When the auto-\nresize is enabled, the render frame will be in sync with the layout\nsize.\n\nAny size update between SetHost() and ShowUI() should be captured by\nthe existing WebUIContents::ResizeDueToAutoResize() and passed to the\nhost.\n\nBug: 40168622\nChange-Id: Iaa9b72e240e8ee18a8110b149193ff9c855e6146\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735037\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Thomas Lukaszewicz <tluk@chromium.org>\nReviewed-by: Avi Drissman <avi@chromium.org>\nCommit-Queue: Keren Zhu <kerenzhu@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332629}",
    "author": "Keren Zhu",
    "author_email": "kerenzhu@chromium.org",
    "date": "2024-07-24T23:13:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f5e21e78ae0469878d9a71447c7f60e367931967"
  },
  {
    "id": 64,
    "message": "Implement SodaSpeechRecognizerImpl\n\nThis CL implements the SodaSpeechRecognizerImpl, a version of the speech recognizer that will be used to run the Web Speech API from the speech recognition service process using the Speech On-Device API (SODA).\n\nDesign doc: go/webspeech+mediastreamtrack\n\nBug: 1495388\nChange-Id: Id81b7276daa6770bd9254a9c6c650ab2678af7b8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5668589\nReviewed-by: Dale Curtis <dalecurtis@chromium.org>\nCommit-Queue: Evan Liu <evliu@google.com>\nCr-Commit-Position: refs/heads/main@{#1332628}",
    "author": "Evan Liu",
    "author_email": "evliu@google.com",
    "date": "2024-07-24T23:12:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1f45824c6d941515af1da2558363d02504da5839"
  },
  {
    "id": 65,
    "message": "[OOPIF PDF] Record PDF.LoadStatus metric later in PDF load\n\nInstead of recording the PDF.LoadStatus metric when the PDF extension\nframe is about to navigate to the PDF extension URL, record it after the\nPDF finishes loading and sets up postMessage.\n\nThere are some cases of other PDF viewer extensions that don't fully\nload until after the OOPIF PDF extension frame performs this navigation.\nThe metric gets recorded, even if the user is not using the Chrome PDF\nviewer. GuestView PDF viewer does not record the metric in this case, so\nmake OOPIF PDF viewer have the same behavior as GuestView.\n\nBug: 355016909\nChange-Id: I9dd112efbbdc8fc518eba9aa338310b51014889f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736234\nCommit-Queue: Andy Phan <andyphan@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Lei Zhang <thestig@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332627}",
    "author": "Andy Phan",
    "author_email": "andyphan@chromium.org",
    "date": "2024-07-24T23:12:12+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cbe08df4529a36d2f1004c999fee8c8c25fa8332"
  },
  {
    "id": 66,
    "message": "Use zero raster-inducing scroll offsets in DisplayItemList::AddToValue()\n\nTo avoid crash when there are DrawScrollingContentsOps.\n\nChange-Id: I106a61b93b30f0f93185d3ccdac5ec6228fe4d17\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738822\nReviewed-by: Philip Rogers <pdr@chromium.org>\nCommit-Queue: Xianzhu Wang <wangxianzhu@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332626}",
    "author": "Xianzhu Wang",
    "author_email": "wangxianzhu@chromium.org",
    "date": "2024-07-24T23:11:12+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c375fd23f705f44755522741a5b9f966e4c750e9"
  },
  {
    "id": 67,
    "message": "Structured Metrics: Reduce upload cadence to 10 minutes\n\nChange-Id: Ia87f032c1b1c1cad662bf7716df3671a74fbe360\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738304\nReviewed-by: Jong Ahn <jongahn@chromium.org>\nCommit-Queue: Andrew Bregger <andrewbregger@google.com>\nCr-Commit-Position: refs/heads/main@{#1332625}",
    "author": "Andrew Bregger",
    "author_email": "andrewbregger@google.com",
    "date": "2024-07-24T23:11:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b58472831985e2b758196ec5df6549b5d4ba4690"
  },
  {
    "id": 68,
    "message": "Record metrics related to the management toolbar button\n\nBug: 347245819\nChange-Id: I55b8d773d05b960a93f4b17e0f3f8bae81dedeb6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735513\nReviewed-by: Owen Min <zmin@chromium.org>\nCommit-Queue: Yann Dago <ydago@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332624}",
    "author": "Yann Dago",
    "author_email": "ydago@chromium.org",
    "date": "2024-07-24T23:10:02+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ef1effc877b66bcd7b375deef76ce8cee8341a4b"
  },
  {
    "id": 69,
    "message": "Queue on demand callbacks when requesting product Info\n\nOn startup, there are could exists multiple on demand callbacks.\nThis CL queue all those callbacks so that we don't fail the latter\ncallbacks after the first one.\n\nBug: 352349972\nChange-Id: I55db7deec3c778fded0667bff56761592c0de9b4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736497\nCommit-Queue: Min Qin <qinmin@chromium.org>\nReviewed-by: Matthew Jones <mdjones@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332623}",
    "author": "Min Qin",
    "author_email": "qinmin@chromium.org",
    "date": "2024-07-24T23:09:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/454b572933303118a22d3237e43a5e78f5b91b2d"
  },
  {
    "id": 86,
    "message": "Add retry to oauth token fetcher.\n\nAlso change fetch mode to kWaitUntilRefreshTokenAvailable instead\nof retrying on USER_NOT_SIGNED_UP for early requests.\n\nBug: b:353974384\nChange-Id: I4957451ab10b4563cf61f9f1bc4a44a196e0ae3e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738335\nCommit-Queue: Ahmed Nasr <anasr@google.com>\nReviewed-by: Benjamin Zielinski <bzielinski@google.com>\nCr-Commit-Position: refs/heads/main@{#1332606}",
    "author": "Ahmed Nasr",
    "author_email": "anasr@google.com",
    "date": "2024-07-24T22:24:52+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/11ceebad952349524f3950b14619aa85384180ea"
  },
  {
    "id": 70,
    "message": "Roll clank/internal/apps from 349eeac8c645 to 607214bd9ff4 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/349eeac8c645..607214bd9ff4\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC aishwaryarj@google.com,chrome-brapp-engprod@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: chromium:349641010\nTbr: aishwaryarj@google.com\nNo-Try: true\nChange-Id: Ib9c0e43a694f66ce8eedcddc377699174cc441c8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739535\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332622}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-24T23:08:53+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0983affe4bd270b29ea4f7bc5e8d96842d58acfc"
  },
  {
    "id": 71,
    "message": "Roll src/components/accessibility/internal/ 3c105c674..76d64799f (3 commits)\n\nhttps://chrome-internal.googlesource.com/chrome-accessibility.git/+log/3c105c674952..76d64799fe09\n\nCreated with:\n  roll-dep src/components/accessibility/internal\n\nBug: N/A\nChange-Id: Id93282f99463bd9b446acd06b1a8f090075e37a1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738664\nReviewed-by: Mark Schillaci <mschillaci@google.com>\nCommit-Queue: Abigail Klein <abigailbklein@google.com>\nCr-Commit-Position: refs/heads/main@{#1332621}",
    "author": "Abigail Klein",
    "author_email": "abigailbklein@google.com",
    "date": "2024-07-24T23:08:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a3205006a18e4180f23fe7dc377bb27d9e4fadad"
  },
  {
    "id": 72,
    "message": "[lensoverlay] Add translate option to detected text context menu.\n\nDemo: https://screencast.googleplex.com/cast/NjYzMjc5MDQxNzIxMTM5MnwwMjg2MTQ5Zi1mYQ\n\nBug: b:354941145\nChange-Id: I558a595c34558edb7d6531300569cf6db63141dd\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5737499\nReviewed-by: Duncan Mercer <mercerd@google.com>\nCommit-Queue: Bryan Nguyen <nguyenbryan@google.com>\nCr-Commit-Position: refs/heads/main@{#1332620}",
    "author": "Bryan Nguyen",
    "author_email": "nguyenbryan@google.com",
    "date": "2024-07-24T23:07:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/77bebb152fe10c87668818b68a8da12e5521d201"
  },
  {
    "id": 73,
    "message": "webcodecs: Temporally disable WebCodecs_Terminate_Worker on windows\n\nBug: 355256365\nChange-Id: I4a8ed0008dcc49dcafefd8ebef8bda98e2b296d3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738712\nCommit-Queue: vikas soni <vikassoni@chromium.org>\nReviewed-by: vikas soni <vikassoni@chromium.org>\nCommit-Queue: Eugene Zemtsov <eugene@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332619}",
    "author": "Eugene Zemtsov",
    "author_email": "eugene@chromium.org",
    "date": "2024-07-24T23:07:33+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9492a812113d5730eb63064fe6a5c489c5284050"
  },
  {
    "id": 74,
    "message": "Add new field trustedSignalsCoordinator to WorkletKey\n\nAdd a new field trustedSignalsCoordinator to WorkletKey in the auction\nmanager and update the related call flow to prepare for fetching the\ntrusted signals KVv2 key.\n\nBug: 337917489\nChange-Id: Ifa75f63c752ba45c594e4fe93417f9a3abcd1b87\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5689154\nCommit-Queue: Tianyang Xu <xtlsheep@google.com>\nReviewed-by: mmenke <mmenke@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332618}",
    "author": "Tianyang Xu",
    "author_email": "xtlsheep@google.com",
    "date": "2024-07-24T22:47:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c318fc5940761a783848b6a8feafe9024a72ba11"
  },
  {
    "id": 75,
    "message": "[Read Anything] Allow downloading of images with ax node id\n\n- This CL adds a function to web contents and image downloader that will\n  allow downloads using the AX node id of the image. This allows us to\n  prevent a compromised reading mode from telling the main renderer to\n  download some arbitrary image and decode it. Similarly this whole task\n  exists to send a safe version of an image to the reading mode renderer\n  without compromising site isolation.\n\n  Bug: 328148648\n\nChange-Id: I1006aac372a9492cd35fcf9960c45d1a59032fa9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5719489\nReviewed-by: Abigail Klein <abigailbklein@google.com>\nReviewed-by: Dmitry Gozman <dgozman@chromium.org>\nReviewed-by: Tom Sepez <tsepez@chromium.org>\nCommit-Queue: Jacob Francis <francisjp@google.com>\nCr-Commit-Position: refs/heads/main@{#1332617}",
    "author": "Jacob Francis",
    "author_email": "francisjp@google.com",
    "date": "2024-07-24T22:46:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0043f24e822485750c0c1bc167aa5211c0994fca"
  },
  {
    "id": 76,
    "message": "Roll Skia from 1561e2127c6f to 267c7c56c6e1 (4 revisions)\n\nhttps://skia.googlesource.com/skia.git/+log/1561e2127c6f..267c7c56c6e1\n\n2024-07-24 skia-autoroll@skia-public.iam.gserviceaccount.com Roll skottie-base from 52a246a737a2 to a6d9a7de5704\n2024-07-24 robertphillips@google.com [graphite] Add comments to public API re Graphite-specific deprecations\n2024-07-24 johnstiles@google.com Allow `writeFunctionCallArgument` to add multiple arguments.\n2024-07-24 skia-autoroll@skia-public.iam.gserviceaccount.com Roll vulkan-deps from 8e90204125ac to 560809a2d33f (3 revisions)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/skia-autoroll\nPlease CC jamesgk@google.com,skiabot@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Skia: https://bugs.chromium.org/p/skia/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux-blink-rel;luci.chromium.try:linux-chromeos-compile-dbg;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nCq-Do-Not-Cancel-Tryjobs: true\nBug: None\nTbr: jamesgk@google.com\nChange-Id: Idaf49fbf0867cc991e11a663b03ef2852b1afdf9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738465\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332616}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T22:46:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f06cc09c58e07e48f87120158d5653ed94e79ef6"
  },
  {
    "id": 77,
    "message": "Fix import onc override policy issue\n\nCurrently, the \"import ONC\" in chrome://network will override the\nexisting policy network configuration. This CL uses the\nManagedNetworkConfigurationHandler to avoid override policy configs.\n\nBug: b/340398740\nChange-Id: Iff1febb7506d76e65e326a6a717727237437ff02\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735627\nReviewed-by: Steven Bennetts <stevenjb@chromium.org>\nCommit-Queue: Jason Zhang <jiajunz@google.com>\nCr-Commit-Position: refs/heads/main@{#1332615}",
    "author": "Jason Zhang",
    "author_email": "jiajunz@google.com",
    "date": "2024-07-24T22:40:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5eeedd7f193d7244e3690facd0bc63f5a35bba30"
  },
  {
    "id": 78,
    "message": "Roll V8 from 70845a08d32a to 3311e78008a9 (15 revisions)\n\nhttps://chromium.googlesource.com/v8/v8.git/+log/70845a08d32a..3311e78008a9\n\n2024-07-24 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Version 12.9.42\n2024-07-24 olivf@chromium.org [runtime] Fix Object.assign fastcase\n2024-07-24 zhaojiazhong-hf@loongson.cn [loong64][mips64][wasm] Introducing F16x8 binary operations opcodes\n2024-07-24 dmercadier@chromium.org [turboshaft] Use Float64/Float32 in ConstantOp instead of double/float\n2024-07-24 irezvov@chromium.org [wasm] Introducing F16x8 conversion operations opcodes]\n2024-07-24 olivf@chromium.org [maglev] Refactor bytecode analysis\n2024-07-24 olivf@chromium.org [maglev] Improvements to loop speeling state merging\n2024-07-24 saelo@chromium.org Disable Runtime::kClearFunctionFeedback for fuzzing\n2024-07-24 dmercadier@chromium.org [turboshaft] Add .CanAllocate() to JS loop stack check effects\n2024-07-24 marja@chromium.org [maglev] Fix another bug in extending properties backing stores\n2024-07-24 nikolaos@chromium.org [heap][test] Fix threaded cctests disabling CSS\n2024-07-24 mliedtke@chromium.org [fuzzer][wasm] Deopt fuzzer: Early return if no deopt was executed\n2024-07-24 omerkatz@chromium.org Reland \"[heap] Prepare tests for sticky markbits\"\n2024-07-24 pthier@chromium.org [turbofan] Prepare for moving RegExp data to trusted space\n2024-07-24 dlehmann@chromium.org [regalloc] Avoid repeated lookups in verifier\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/v8-chromium-autoroll\nPlease CC liviurau@google.com,machenbach@google.com,v8-waterfall-gardener@grotations.appspotmail.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-blink-rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:dawn-linux-x64-deps-rel\nChange-Id: I8ab93b87822ed5dce2359bc5c4511e6b52331224\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738462\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332614}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T22:38:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a3571102507b52e181fe06533d00a3a28a1fd220"
  },
  {
    "id": 79,
    "message": "Update OpenH264 dependency\n\nDiff:\nhttps://github.com/cisco/openh264/compare/09a4f3ec842a8932341b195c5b01e141c8a16eb7..478e5ab3eca30e600006d5a0a08b176fd34d3bd1\n\nChangelog (including some fuzzing fixes):\n478e5ab3 Change picOrderCnt to 2 (#3752)\nde026176 Remove unused pSpatialLayer variable in DetermineTemporalSettings (#3761)\n1c238875 Fix race issue among decoding threads which causes broken frame. (#3735)\n1debdaec Fix crash on multi-thread decoding. (#3736)\nf86f0e47 Revert \"Fix WelsTraceCallback conversion. (#3722)\" (#3739)\n28b533af Fix WelsTraceCallback conversion. (#3722)\nc0e5ea28 Fix regression in PR#3707 for multi-thread decoding (#3734)\nc59550a2 update release note (#3726)\nb29fd81e update openh264 version to v2.4.1 (#3724)\n4f01c15b Fix glitches that sometimes happen near the scene changes. (#3707)\ncfbd5896  Add security policy (#3708)\nfd66e67a Fix off by one regression in decoder (#3704)\n34a0d2d3 update release note (#3703)\n34e14eab update openh264 version to v2.4.0 (#3691)\n15d02fc4 Use Ant compile demo and UT projects when NDK version lower than r18 (#3690)\nb3feec25 emscripten processor count (#3667)\n859a08c3 A fuzzer found a null pointer dereference in welsDecoderExt.c (#3686)\n986bd65b Fix for out of bounds read issue by introducing bounds check on iPicBuffIdx in CWelsDecoder::ReleaseBufferedReadyPictureReorder (#3685)\n008465e9 compatible build with andrandroid-ndk-r18 (#3672)\n8f7fc380 compatible build with android-ndk-r18\n98660664 add simd optimizations and fix for loongarch (#3649)\n6967c09a Avoid providing a warning if the requested number of threads is the (#3618)\n17721565 Bump the API version to add support for new requirements from Firefox. (#3648)\n\nBUG=chromium:350938995\n\nChange-Id: I31d27e91d3e1788a3529c0e605754862f8049bc8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5676860\nReviewed-by: Sergey Silkin <ssilkin@chromium.org>\nCommit-Queue: Philipp Hancke <philipp.hancke@googlemail.com>\nReviewed-by: Dale Curtis <dalecurtis@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332613}",
    "author": "Philipp Hancke",
    "author_email": "philipp.hancke@googlemail.com",
    "date": "2024-07-24T22:35:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9d7bae44ba5835e40c248a818673b7c94901c67e"
  },
  {
    "id": 80,
    "message": "[iOS] Finchable image presence in rich in-product help bubbles\n\nMade images in in-product help rich bubbles have Finchable presence,\nand updated tests. If the no-image case is launched, the images support\nwill be cleaned up at a later date.\n\nBug: b:351004118\nChange-Id: I88fd87ce69a01abcbce52034cf9ad9e6b5f06476\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735522\nCommit-Queue: Nicolas MacBeth <nicolasmacbeth@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Ginny Huang <ginnyhuang@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332612}",
    "author": "Nicolas MacBeth",
    "author_email": "nicolasmacbeth@google.com",
    "date": "2024-07-24T22:30:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/270fce061c5e40aedec376e5d95f1c19c17a1800"
  },
  {
    "id": 81,
    "message": "Roll src-internal from 2459e0fb45e6 to 960d1a5e5519 (1 revision)\n\nhttps://chrome-internal.googlesource.com/chrome/src-internal.git/+log/2459e0fb45e6..960d1a5e5519\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/src-internal-chromium-autoroll\nPlease CC avi@google.com,chrome-browser-infra-team,estade@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: avi@google.com,estade@google.com\nChange-Id: Id8f8d537ae97044ec4f98c179c258f738b9818a7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5737789\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332611}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-24T22:29:56+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/80d245ff56b4f05069b5889be2096a40126f5b4d"
  },
  {
    "id": 82,
    "message": "[iOS] Report OS Notif Auth Status to NAU on every foreground.\n\nThis CL reports the OS notification status to NAUs on every foreground\nfor Content Notifications.\n\nBug: 355057017\n\nChange-Id: I504c28a0d3299de3df70c0659a725a3cbd4c119d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736455\nReviewed-by: Sergio Collazos <sczs@chromium.org>\nCommit-Queue: Guillem Perez <guiperez@google.com>\nCr-Commit-Position: refs/heads/main@{#1332610}",
    "author": "Guillem Perez",
    "author_email": "guiperez@google.com",
    "date": "2024-07-24T22:29:02+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/83cecab236d995f6749b202e135a440a8d7b6dbe"
  },
  {
    "id": 83,
    "message": "app-controls: Fix letter spacing in Pin Keyboard placeholder\n\nThe placeholder of the parental controls pin setup and verify dialogs\nshould include letter-spacing. I'm removing the CSS that forces the\nspacing to be normal.\n\nVerify pin dialog: https://screenshot.googleplex.com/8UivKKWXvWRc8mb\nSetup pin dialog: https://screenshot.googleplex.com/AJ4ZgpPjGpfLwwf\n\nBug: b:355160627\nChange-Id: Iac0efaea8e4579aabd590a48106c628a0d471c5e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738774\nReviewed-by: Aga Wronska <agawronska@chromium.org>\nReviewed-by: Xiyuan Xia <xiyuan@chromium.org>\nCommit-Queue: Jesulayomi Kupoluyi <lkupo@google.com>\nCr-Commit-Position: refs/heads/main@{#1332609}",
    "author": "LK Kupoluyi",
    "author_email": "lkupo@chromium.org",
    "date": "2024-07-24T22:27:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1846f003701ffb88565b2c2c1a6ef70d0608025f"
  },
  {
    "id": 84,
    "message": "[SafeArea] Add about flag entries for DynamicSafeAreaInsets\n\nBug: 324436581\nChange-Id: I93eaf14b3c621259db247e46121064ae3cbdac45\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5699426\nCommit-Queue: Charles Hager <clhager@google.com>\nAuto-Submit: Wenyu Fu <wenyufu@chromium.org>\nReviewed-by: Charles Hager <clhager@google.com>\nCr-Commit-Position: refs/heads/main@{#1332608}",
    "author": "Wenyu Fu",
    "author_email": "wenyufu@chromium.org",
    "date": "2024-07-24T22:26:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/351ce33cf7c31b4c195750ea8bfc823c00b75668"
  },
  {
    "id": 85,
    "message": "IDB WPTs: Expand IDBObjectStore.createIndex() to run on workers\n\nThe change updates IDBobjectstore.createIndex() related WPTs to also run on dedicated, service, and shared workers. Currently, these only run in window environments.\n\nBug: 41455766\nChange-Id: Icc51731881db1350c9cc04213cbf141fb24f570b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5676928\nReviewed-by: Rahul Singh <rahsin@microsoft.com>\nReviewed-by: Evan Stade <estade@chromium.org>\nAuto-Submit: Garima Chadha <garimachadha@microsoft.com>\nCommit-Queue: Evan Stade <estade@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332607}",
    "author": "Garima Chadha",
    "author_email": "garimachadha@microsoft.com",
    "date": "2024-07-24T22:26:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f3cfc7ba09e7edbf314ae25768d84fabf69af991"
  },
  {
    "id": 88,
    "message": "Roll GoogleTest from 57e107a10ea4 to 352788321faa (1 revision)\n\nhttps://chromium.googlesource.com/external/github.com/google/googletest.git/+log/57e107a10ea4..352788321faa\n\n2024-07-24 jacobsa@google.com gmock-actions: make DoAll convert to OnceAction via custom conversions.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/googletest-chromium-autoroll\nPlease CC asully@google.com,jonathanjlee@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:linux_chromium_cfi_rel_ng;luci.chrome.try:win-chrome\nTbr: asully@google.com,jonathanjlee@google.com\nChange-Id: I08c359c9cf24a1d2adb2b507b5fd1e95465b3bf9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734991\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332604}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T22:18:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ba1b748ed89c67b44bb29aa62ae19f2f1457c298"
  },
  {
    "id": 89,
    "message": "[Omnibox] Apply vertical card pardding to base view, not answer view\n\nThis allows the decoration to remain aligned with the answer text. Right\npadding is left unchanged since we don't want that to apply to the\naction chip carousel.\n\nBug: 355043942\nChange-Id: I8eebc8b304ced364613fbbd6a650d136b9500088\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735621\nReviewed-by: Tomasz Wiszkowski <ender@google.com>\nCommit-Queue: Patrick Noland <pnoland@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332603}",
    "author": "Patrick Noland",
    "author_email": "pnoland@google.com",
    "date": "2024-07-24T22:17:53+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ca06f28c48a3c8f4fea20a3c4a7d15dc7aaae18d"
  },
  {
    "id": 90,
    "message": "Disable new GWSAbandonedPageLoadMetricsObserverBrowserTest on ASAN.\n\nBug: 352578800\nChange-Id: I21fe7ac517ccc6232fa03fc99127dea5b94ca592\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5737569\nCommit-Queue: Evan Stade <estade@chromium.org>\nReviewed-by: Avi Drissman <avi@chromium.org>\nCommit-Queue: Avi Drissman <avi@chromium.org>\nAuto-Submit: Evan Stade <estade@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332602}",
    "author": "Evan Stade",
    "author_email": "estade@chromium.org",
    "date": "2024-07-24T22:16:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b2e1b7ef519ef75630849f4e58bc51fa3ff2b5b9"
  },
  {
    "id": 91,
    "message": "Add histogram for recording the profile management state.\n\nA similar histogram was recently added in https://crrev.com/c/5564406\nbut will be removed as it's recorded lower in the policy stack and\ntherefore doesn't have the same granularity.\n\nBug: b/340852535\nChange-Id: Ia67c3a96cdd8369b46e946a905f359db18e67eb3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5721995\nReviewed-by: Owen Min <zmin@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Igor Ruvinov <igorruvinov@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332601}",
    "author": "Igor Ruvinov",
    "author_email": "igorruvinov@chromium.org",
    "date": "2024-07-24T22:15:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fd7aa7ae855f83fbc046c9966e9b86dd29fadfe7"
  },
  {
    "id": 92,
    "message": "Specs: Implement disclosure actions\n\n- Create a new enum to identify the FRE disclosure's version.\n- Expose a set API through shopping_service_handler for the new pref.\n- Add decline button to the disclosure.\n\nScreenshot: screenshot/3bbckjVobHFJkCT\nBug: b:353825838\nChange-Id: Id19969e86834d1914169f3fc832c8c1441dea7dc\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5719948\nAuto-Submit: Ayman Almadhoun <ayman@chromium.org>\nReviewed-by: Yue Zhang <yuezhanggg@chromium.org>\nReviewed-by: Paul Adedeji <pauladedeji@google.com>\nReviewed-by: Alex Gough <ajgo@chromium.org>\nCommit-Queue: Ayman Almadhoun <ayman@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332600}",
    "author": "Ayman Almadhoun",
    "author_email": "aymana@google.com",
    "date": "2024-07-24T22:15:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/23de5136a5dea65fafc2fa791e430e43b4c4f484"
  },
  {
    "id": 93,
    "message": "overview: Remove feedback button\n\nTest: manual\nChange-Id: I7d67261e2914fa40a01c6ffb411e0714dbf7739e\nFixed: b/353583772\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735622\nCommit-Queue: Sammie Quon <sammiequon@chromium.org>\nReviewed-by: Elijah Hewer <hewer@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332599}",
    "author": "Sammie Quon",
    "author_email": "sammiequon@chromium.org",
    "date": "2024-07-24T22:14:08+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/151dcfd1a5c7e6b694f1bb1c9f4a66c42c2195c3"
  },
  {
    "id": 94,
    "message": "Update chromium_config for \"ci/android-14-arm64-rel\"\n\nMake it the same as \"ci/android-pie-arm64-rel\" to be ready for CQ\n\nBug: 352811552\nChange-Id: Ief1d29d2db96cda28fc85df7444de966f7600f5d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736820\nAuto-Submit: Haiyang Pan <hypan@google.com>\nReviewed-by: James Shen <zhiyuans@google.com>\nCommit-Queue: James Shen <zhiyuans@google.com>\nCr-Commit-Position: refs/heads/main@{#1332598}",
    "author": "Haiyang Pan",
    "author_email": "hypan@google.com",
    "date": "2024-07-24T22:13:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0d0973641d3684cfa90d4aef4155577034730474"
  },
  {
    "id": 95,
    "message": "[omnibox][ml] Update chrome://flags to allow testing of Android models.\n\nBug: 40062540\nChange-Id: Ia56207b4aa5cf280c7591dc213bb6a6851653607\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735964\nReviewed-by: manuk hovanesian <manukh@chromium.org>\nCommit-Queue: Khalid Peer <khalidpeer@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332597}",
    "author": "Khalid Peer",
    "author_email": "khalidpeer@chromium.org",
    "date": "2024-07-24T22:12:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2a1c63771fc0542a5706328355e23143bd7d067b"
  },
  {
    "id": 96,
    "message": "viz: FrameIntervalDecider\n\nThis is a rewrite of FrameRateDecider system. Overall design is linked\nfrom bug, but in general, clients do less \"aggregation\" and just passes\ninformation to viz. Viz uses a list of independent and stateless\nmatchers. A feature is added to control the new code.\n\nAdd new FrameIntervalInputs IPC struct which is added to\nCompositorFrameMetadata for new information from clients. Add enough\ninformation here to largely allow replicating the behavior of\nFrameRateDecider.\n\nframe_interval_decider.cc/h and frame_interval_matcher.cc/h contain most\nof the logic and shared among all platforms. Platform-specific code is\nin display.cc that initializes decider, and deals with how to handle the\nmatched result. For this CL,  only android is correctly hooked up\nend-to-end.\n\nResult is a variant that specifically allows \"FrameIntervalClass\"\nthat's more descriptive, instead of a specific frame interval. The\ninterval class allows more flexibility and that can map to some platform\nAPIs directly.\n\nOverall the system is more easily extensible and new matchers is\nrelatively safe to existing behavior.\n\nSome key differences from FrameRateDecider:\n* Two clear distinct modes where either system supports setting\n  arbitrary frame rates, or a fixed set of supported intervals.\n* FrameRateDecider sometimes uses only frame sinks that updated for a\n  frame and excluding ones that did not draw. FrameIntervalDecider uses\n  difference in frame time instead.\n* Video information is only implemented for VideoFrameSubmitter for now.\n  Only Android WebView still uses the path where video lives in cc.\n* The heuristic for avoiding quick frame interval no longer requires the\n  computed frame rate stays consistent and only delays lowering frame\n  rate.\n\nBug: 346732738\nChange-Id: Ie4c4fa4440ad5eed394472499dd4fcbef87cbb9e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5704115\nReviewed-by: Dale Curtis <dalecurtis@chromium.org>\nReviewed-by: Jonathan Ross <jonross@chromium.org>\nReviewed-by: Joe Mason <joenotcharles@google.com>\nCommit-Queue: Bo Liu <boliu@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332596}",
    "author": "Bo Liu",
    "author_email": "boliu@chromium.org",
    "date": "2024-07-24T22:10:16+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e6d49fe923d1621ccd4a0917fe85a344f387340a"
  },
  {
    "id": 139,
    "message": "Use weak ptr on for DistillerViewer::OnDistillerFinished\n\nWhile the current implementation prevents UAF in that case, a change in\nDistillerPage could lead to issues.\nUse a WeakPtr to avoid issues in the future.\n\nFixed: 40849562\nChange-Id: I09e723708f1cf931351dd7bab28e6c21b6798f4d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712688\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Quentin Pubert <qpubert@google.com>\nCommit-Queue: Olivier Robin <olivierrobin@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328153}",
    "author": "Olivier ROBIN",
    "author_email": "olivierrobin@google.com",
    "date": "2024-07-16T15:26:09+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0df7384c0e26b727535f7208f515c8f4e09b0130"
  },
  {
    "id": 97,
    "message": "Roll devtools-internal from a61739884538 to ce6bd7bd11d9 (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/a61739884538..ce6bd7bd11d9\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: I047fcc426aa441f3957925c468b43f23cf4d7f4e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738460\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332595}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-24T22:03:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0417b4717ef60d4bab91a5b3bb86a8155b69980c"
  },
  {
    "id": 98,
    "message": "[omnibox] Remove fallthroughs in `AutocompleteMatch::GetVectorIcon()`\n\nAvoids `ABSL_FALLTHROUGH_INTENDED` (i.e. `case` statements that don't\n`break` or `return` and therefore fall through to the next `case`):\nReplaces:\n\n```\nswitch(x) {\n  case 1:\n    if (predicate);\n      return 1;\n    ABSL_FALLTHROUGH_INTENDED;\n  case 2:\n    return 2;\n}\n```\n\nwith:\n\n```\nswitch(x) {\n  case 1:\n    return predicate ? 1 : 2;\n  case 2:\n    return 2;\n}\n```\n\nChange-Id: Ib7a58e2e55b1293cc8bdf7de251cbb5631a7ab62\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5719020\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Angela Yoeurng <yoangela@chromium.org>\nAuto-Submit: manuk hovanesian <manukh@chromium.org>\nReviewed-by: Angela Yoeurng <yoangela@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332594}",
    "author": "manukh",
    "author_email": "manukh@chromium.org",
    "date": "2024-07-24T22:03:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/89661b22439e27f3d1be4f0c06d1fe6b67ddc132"
  },
  {
    "id": 99,
    "message": "Fix CanvasPath BoundingRect calculation\n\nWhen RectF::BoundingRect is created with extreme large width or height,\nit's possible that (x + width) or (y + height) cannot have a precise\nfloat presentation. So that right and bottom are not included in the\nRectF.\n\nIn the bug, the user set y = -10^10 and height = 10^10+100, so that the\nbottom of the bounding rect suppose to be at 100. Since 10^10+100\ncannot be presented by float (float has 7 digit of precision), height\nis rounded to 10^10. Hence, the new bottom becomes 0.\n\nIn this cl: If the exact height or width cannot be presented by float,\nthen it's always rounded up to the nearest float value.\n\nBug: 40848283\n\nChange-Id: I5c3db827ebd8da255019fb35d99a03e955fc82cd\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5698290\nReviewed-by: danakj <danakj@chromium.org>\nCommit-Queue: Yi Xu <yiyix@chromium.org>\nReviewed-by: Andres Ricardo Perez <andresrperez@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332593}",
    "author": "Yi Xu",
    "author_email": "yiyix@chromium.org",
    "date": "2024-07-24T22:00:34+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/360e602ccc1d7f03aeb004bb23c65203c95d0de7"
  },
  {
    "id": 100,
    "message": "[SAA] Fix remote disconnection issue with GC & handles\n\nThere's a bug where if the SAA handle falls out of scope, connections\nbetween other objects (like Broadcast Channels) allocated from the\nhandle cease to work.\nhttps://github.com/privacysandbox/privacy-sandbox-dev-support/issues/398\n\nLet's add a test for the impacts of GC and transfer ownership of key\nremotes/objects to a per-window-global to prevent disconnection issues.\n\nI have confirmed this test crashes due to disconnection on trunk, and\nsucceeds here.\n\nFixed: 355018914\nChange-Id: I1e760964a9927293b8a21565afe55360a100e8aa\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736256\nReviewed-by: Ben Kelly <wanderview@chromium.org>\nCommit-Queue: Ayu Ishii <ayui@chromium.org>\nReviewed-by: Chris Fredrickson <cfredric@chromium.org>\nAuto-Submit: Ari Chivukula <arichiv@chromium.org>\nReviewed-by: Ayu Ishii <ayui@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332592}",
    "author": "Ari Chivukula",
    "author_email": "arichiv@chromium.org",
    "date": "2024-07-24T21:59:34+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4862b4da110b51f193b033d7943be47e15437bfc"
  },
  {
    "id": 101,
    "message": "[iOS] Refactor Safety Check Autorun to IOSChromeSafetyCheckManager.\n\nThis change moves the responsibility of Safety Check autoruns from the\nMagic Stack UI to the IOSChromeSafetyCheckManager. This ensures that\nSafety Checks run reliably even when the Magic Stack UI is not active or\nvisible.\n\nKey benefits:\n\n- Improved reliability: Safety Checks are no longer dependent on the\n  state of the Magic Stack UI.\n- Maintain user experience: Existing loading states and user\n  interactions remain unchanged as the Magic Stack module continues to\nobserve Safety Check Manager changes.\n\nChange-Id: Ib837125aee909b37d3a782309aa1dd6b3dde84e0\nFixed: 347975185\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5722908\nCommit-Queue: Benjamin Williams <bwwilliams@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Chris Lu <thegreenfrog@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332591}",
    "author": "Benjamin Williams",
    "author_email": "bwwilliams@google.com",
    "date": "2024-07-24T21:59:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d7cf668c6421a6c06192c61b73c73ded8be93f74"
  },
  {
    "id": 102,
    "message": "Tab Organization: Invalidate group with 0 new tabs\n\nAddresses a crash when accepting a group with no new tabs. These groups\nshould not be considered valid and therefore shouldn't get to the point\nof being accepted.\n\nBug: 354850235\nChange-Id: I667a96db26abb677fc9c79963a215358c60212fa\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734499\nAuto-Submit: Emily Shack <emshack@chromium.org>\nReviewed-by: David Pennington <dpenning@chromium.org>\nCommit-Queue: David Pennington <dpenning@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332590}",
    "author": "Emily Shack",
    "author_email": "emshack@chromium.org",
    "date": "2024-07-24T21:56:57+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/351208ee3e9583c9b645adf06fb13e161be7ec2a"
  },
  {
    "id": 140,
    "message": "Use inline constexpr for constant strings in chrome/updater\n\nThis allows for better optimization and use of compile-time tools.\n\nBug: 350944663\nChange-Id: I79f0b2ce020bb4f6d590f542a132735ad5de5270\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709248\nReviewed-by: Joshua Pawlicki <waffles@chromium.org>\nAuto-Submit: Avi Drissman <avi@chromium.org>\nCommit-Queue: Joshua Pawlicki <waffles@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328152}",
    "author": "Avi Drissman",
    "author_email": "avi@chromium.org",
    "date": "2024-07-16T15:25:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4263064297b23e5a597f8608a2621d493beac52e"
  },
  {
    "id": 103,
    "message": "brightness: Finalize strings\n\nIDS_SETTINGS_KEYBOARD_COLORS\nIDS_SETTINGS_KEYBOARD_BRIGHTNESS_LABEL\nIDS_SETTINGS_KEYBOARD_AUTO_BRIGHNTESS_ENABLE_LABEL\nIDS_SETTINGS_KEYBOARD_AUTO_BRIGHNTESS_ENABLE_SUB_LABEL\n - hash: ef1de73b97c16d0e6c78346e2534eaf426a9b003\n - screenshot: http://screen/BuyYXDh5Wh5382C\n\nIDS_SETTINGS_DISPLAY_BRIGHTNESS_LABEL\nIDS_SETTINGS_DISPLAY_AUTO_BRIGHTNESS_TOGGLE_LABEL\nIDS_SETTINGS_DISPLAY_AUTO_BRIGHTNESS_TOGGLE_SUBTITLE\n - hash: f3d8a8b052da00754cccd858d7e14feae16c4581\n - screenshot: http://screen/76vecFye4MPPfLr\n\nBug: b/329861823\nTest: CQ\nChange-Id: I7280ce498bfd90a13684ad2fda10ca183218ae68\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734691\nReviewed-by: Wes Okuhara <wesokuhara@google.com>\nCommit-Queue: Longbo Wei <longbowei@google.com>\nCr-Commit-Position: refs/heads/main@{#1332589}",
    "author": "Longbo Wei",
    "author_email": "longbowei@google.com",
    "date": "2024-07-24T21:55:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7a9f06e30bdac4b51b9d259ce1e3b2c9b2b435cf"
  },
  {
    "id": 104,
    "message": "Roll Chrome Mac Arm PGO Profile\n\nRoll Chrome Mac Arm PGO profile from chrome-mac-arm-main-1721843954-ce6d2842ed04eb5765c4c42eaffdfb4d0ab167be-3f75879e72af7aa4330999a92b65a0caa78fee24.profdata to chrome-mac-arm-main-1721851124-dafd418f852199ec81f2729a576ae2adfacaed88-c113d06ffb53ec81c836f2558e9534f6e66740a3.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-mac-arm-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:mac-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I9ff7aa860bac6cf6250c752170ad0ad8a4fd093c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738237\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332588}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T21:55:49+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e31e331f517bd52679de67bd3f8aa7bcd911f35d"
  },
  {
    "id": 105,
    "message": "Roll clank/internal/apps from 06a8b27e8cd6 to 349eeac8c645 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/06a8b27e8cd6..349eeac8c645\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC aishwaryarj@google.com,chrome-brapp-engprod@google.com,lizeb@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: aishwaryarj@google.com,lizeb@google.com\nNo-Try: true\nChange-Id: Ib218e646de8433b0a00361de9e754c83c22501b5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5737791\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332587}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-24T21:52:54+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d19c5defb3504ca1063cfc65dd779959325d0887"
  },
  {
    "id": 106,
    "message": "[iOS] Add metric for Contextual Panel model response time\n\nThis metric tracks how long it takes each info block's model to respond.\n\nBug: 349839698\nChange-Id: I82c19f4d2632a42b7ca2891b027f8268c684fa47\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736356\nAuto-Submit: Robbie Gibson <rkgibson@google.com>\nReviewed-by: Sebastien Seguin-Gagnon <sebsg@chromium.org>\nCommit-Queue: Sebastien Seguin-Gagnon <sebsg@chromium.org>\nReviewed-by: Nicolas MacBeth <nicolasmacbeth@google.com>\nCr-Commit-Position: refs/heads/main@{#1332586}",
    "author": "Robbie Gibson",
    "author_email": "rkgibson@google.com",
    "date": "2024-07-24T21:51:34+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a89163ecc0dbbd1de62aea8880bff9d3fd9f4e87"
  },
  {
    "id": 107,
    "message": "[ToolbarPinning3] Add the name of the party managing the NewTabPage.\n\nAdds a method that update the CustomizeChromePage whenever the default\nsearch engine gets updated.\n\nChange-Id: I8a5c10b4c5e2e634626e6541a1aa49846809502b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5719118\nReviewed-by: Tibor Goldschwendt <tiborg@chromium.org>\nReviewed-by: Chris Bookholt <bookholt@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Chris Bookholt <bookholt@chromium.org>\nCommit-Queue: David Pennington <dpenning@chromium.org>\nReviewed-by: Caroline Rising <corising@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332585}",
    "author": "David Pennington",
    "author_email": "dpenning@chromium.org",
    "date": "2024-07-24T21:50:52+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/94dd586296da5b76887695d2232f646b77dd186a"
  },
  {
    "id": 108,
    "message": "Move AutocompleteInputConnection to a dedicated file\n\nMove AutocompleteInputConnection out of\nSpannableAutocompleteEditTextModel, and also move InputDelegate out of\nAutocompleteEditTextModelBase, and put them into\nAutocompleteInputConnection.java.\n\nChange-Id: If5cb7d7b607744660fafdfa733f7384b7184ec0b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738173\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Tomasz Wiszkowski <ender@google.com>\nCommit-Queue: Gang Wu <gangwu@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332584}",
    "author": "Gang Wu",
    "author_email": "gangwu@chromium.org",
    "date": "2024-07-24T21:49:56+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c6a97a708e195fe026a0a41ad0a9e94fe2ab9001"
  },
  {
    "id": 109,
    "message": "Cast pointer diff in WTF::Vector::InsertAt to wtf_size_t\n\nWTF::Vector::InsertAt(Iterator, U&&) does a pointer diff and pass it to\ninsert(wtf_size_t,U&&). We need to do a cast from pointer diff type to\nwtf_size_t to use this function.\n\nWithout this, calling Vector::InsertAt() fails with:\nerror: implicit conversion loses integer precision: 'long' to 'wtf_size_t' (aka 'unsigned int') [-Werror,-Wshorten-64-to-32]\n 2129 |   insert(position - begin(), val);\n      |   ~~~~~~ ~~~~~~~~~^~~~~~~~~\nPreviously, nothing called this and so the compiler never emitted\na warning.\n\nChange-Id: Ied96164ff1869e53b1129d5028d20371d4fb6d70\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5734165\nReviewed-by: Daniel Cheng <dcheng@chromium.org>\nCommit-Queue: Hao Liu <haoliuk@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332583}",
    "author": "Hao Liu",
    "author_email": "haoliuk@chromium.org",
    "date": "2024-07-24T21:48:01+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/060056c92ba768e0dff5d8ee124bd8e12371ea1e"
  },
  {
    "id": 110,
    "message": "Fix OOIPF scroll bubbling when filtering for debounce\n\nThere are a few ways gesture events (like scroll, pinch-zoom) can be\nfiltered:\n\n* Scroll debouncing, which removes spurious (unintended) ScrollEnd/Begin\n  pairs\n* Stylus writing which consumes scroll events\n* Browser-side touch action - which uses the renderer-hit-tested CSS\n  touch action property to prevent sending disallowed gesture events.\n* Fling suppressions which suppresses taps while flinging, as well as\n  prevents sending fling events to the renderer.\n\nFiltering works by marking an event as \"Consumed\" with the \"Browser\"\nbeing the source.\n\nThe linked bug is caused by the scroll debounce filter. When the\ndebounce filter removes an event it still calls GestureEventAck. The\nscroll bubbling logic didn't consider this case (and currently has no\nway to differentiate whether a consumed event is filtered or was\nconsumed in the renderer) so the filtered event was confusing the\nbubbling state. See comment #30 in the linked bug for a detailed\ndescription of the issue.\n\nThis CL plumbs through the ACK source so that client code, like the\nscroll bubbling logic, can ignore filtered events.\n\nChange-Id: I8467bc84fed5f307eb1e6a54c89a44cafbc57ba2\nBug: 346629231\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735734\nReviewed-by: James Maclean <wjmaclean@chromium.org>\nCommit-Queue: David Bokan <bokan@chromium.org>\nReviewed-by: Dave Tapuska <dtapuska@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332582}",
    "author": "David Bokan",
    "author_email": "bokan@chromium.org",
    "date": "2024-07-24T21:45:29+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/bb9e92867f59d0d292b3bac94d83e3bee60b4db0"
  },
  {
    "id": 111,
    "message": "[Android] Allow modularized code to create HelpAndFeedbackLauncherImpl\n\nThrough HelpAndFeedbackLauncherFactory, modularized code can create a\nHelpAndFeedbackLauncherImpl without having to pipe it through parameters.\n\nUse android_library_factory to avoid circular dependencies.\n\nBug: 40841118\nChange-Id: I561ba1f4e3f9d8033b33a8e7de082b92db4d9292\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5732094\nCommit-Queue: Henrique Nakashima <hnakashima@chromium.org>\nReviewed-by: Andrew Grieve <agrieve@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nOwners-Override: Andrew Grieve <agrieve@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332581}",
    "author": "Henrique Nakashima",
    "author_email": "hnakashima@chromium.org",
    "date": "2024-07-24T21:45:12+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e92d032d53a33b3a503394679cbc868a757256a1"
  },
  {
    "id": 112,
    "message": "Fix JNI Zero bugs\n\n- Generated header guards now escape `$` in java class names\n- Generated header file names now replace `$` with `__` because siso\n  does not like it.\n- Fixes issue where removing generics can stick a type name with a param\n  name. Now we add a space in place of generics but strip spaces from\n  types.\n- Remove unneeded dependency on build from sample.\n\nChange-Id: Ic4806a467e462d833f2d76329c3f6ef6d94c1c15\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739033\nCommit-Queue: Sam Maier <smaier@chromium.org>\nReviewed-by: Sam Maier <smaier@chromium.org>\nAuto-Submit: Mohamed Heikal <mheikal@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332580}",
    "author": "Mohamed Heikal",
    "author_email": "mheikal@chromium.org",
    "date": "2024-07-24T21:42:49+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/665f7e3dca4a3199fbb69b4c3053db05863e2e23"
  },
  {
    "id": 113,
    "message": "Stop sending enterprise reports for consumer accounts\n\nAccountInfos are used to determine if a gaia account is a consumer account. If an AccountInfo's hosted_domain has not been populated, then we don't know what password reuse type we are dealing with. So, we should admit that.\n\nThis should prevent confusing consumer accounts for GSuite accounts in cases where hosted_domain isn't populated as the field is updated async.\n\nBug: b/347979903\nChange-Id: Ic158ed3058517abcccab4e37eba2683a1b288353\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5727113\nReviewed-by: Xinghui Lu <xinghuilu@chromium.org>\nCommit-Queue: Nwokedi Idika <nwokedi@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332579}",
    "author": "Nwokedi Idika",
    "author_email": "nwokedi@chromium.org",
    "date": "2024-07-24T21:42:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c8d472e3405abedd5fdc9663039f37b7e2369629"
  },
  {
    "id": 114,
    "message": "Roll Perfetto Trace Processor Win from 97caa8a1afc2 to bca707edfa05\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/97caa8a1afc2..bca707edfa05\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-trace-processor-win-chromium\nPlease CC chrometto-team@google.com,perfetto-bugs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: perfetto-bugs@google.com\nChange-Id: If0dd1f0abe220c9d320ae9a80f173162d3f84abb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738116\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332578}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T21:38:31+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b68d787921d7c38d2a384a90f22aaa91cabb4c2a"
  },
  {
    "id": 115,
    "message": "[Android] Instantiate SettingsLauncherImpl through Factory\n\nThis is a simple replace for easy review, will go through usages that\ncan be simplified in a follow-up.\n\nBug: 354019554\nChange-Id: Iabc2b5298eae0019c849aad58fab6d4eb42e9fa0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735739\nOwners-Override: Andrew Grieve <agrieve@chromium.org>\nCommit-Queue: Henrique Nakashima <hnakashima@chromium.org>\nReviewed-by: Andrew Grieve <agrieve@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332577}",
    "author": "Henrique Nakashima",
    "author_email": "hnakashima@chromium.org",
    "date": "2024-07-24T21:38:16+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/524abae6e1c36e90900f20b1eb4aeb784582df6a"
  },
  {
    "id": 116,
    "message": "Updater: Create 3pp CIPD packages for Windows arm64 updater.\n\nBug: 40063600\nChange-Id: I9b87e2ff63e0a567cc875cc2af321f7728ee5589\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738957\nAuto-Submit: Joshua Pawlicki <waffles@chromium.org>\nCommit-Queue: Sorin Jianu <sorin@chromium.org>\nReviewed-by: Sorin Jianu <sorin@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332576}",
    "author": "Joshua Pawlicki",
    "author_email": "waffles@chromium.org",
    "date": "2024-07-24T21:37:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/31e3122e3041178ab46fa716e1f23e08836917ff"
  },
  {
    "id": 156,
    "message": "[headless] Re-enable more Mac ASAN tests\n\nA number of headless browser tests were disabled when running\non Mac with ASAN. After https://crrev.com/c/5637879 and\nhttps://crrev.com/c/5640568 landed those tests appear to\nbe stable, so this CL re-enables them.\n\nBug: 40694526\nChange-Id: Ic829881a0e8f673e159c8d6cd36c5c7c2881216d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709649\nReviewed-by: Eric Seckler <eseckler@chromium.org>\nCommit-Queue: Peter Kvitek <kvitekp@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328136}",
    "author": "Peter Kvitek",
    "author_email": "kvitekp@chromium.org",
    "date": "2024-07-16T15:04:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e6914f9447b1ea9b73bc4b514590e82ea2882698"
  },
  {
    "id": 117,
    "message": "Roll Catapult from 98802e1e928a to c0b92670a9fc (1 revision)\n\nhttps://chromium.googlesource.com/catapult.git/+log/98802e1e928a..c0b92670a9fc\n\n2024-07-24 bsheedy@chromium.org [typ] Surface associated bugs in test finish\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/catapult-autoroll\nPlease CC chrome-speed-operations@google.com,dberris@chromium.org,jbudorick@chromium.org,johnchen@chromium.org,sullivan@chromium.org,wenbinzhang@google.com,zhanliang@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nBug: chromium:333739949\nTbr: zhanliang@google.com\nChange-Id: If2aedb2f2c93af1d9a29c23e91c8accb48afd077\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735011\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332575}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T21:37:07+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a6a5bc333b3d93ef48dcf1670c2175bb207ee5f0"
  },
  {
    "id": 118,
    "message": "Add a timer to debounce incoming network cost notifications\n\nPerformance traces have shown a large number of network cost change\nnotifications occurring on device wake/network connection.  This change\nadds a timer to debounce these notifications to reduce cross thread\ncommunication.  This is a similar strategy used in the\nNetworkChangeCalculator for other network events.\n\nBug: 347047639\nChange-Id: I172e2466b7db0d9fbac1555ebe4a97f793896802\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5731034\nReviewed-by: Cliff Smolinsky <cliffsmo@microsoft.com>\nReviewed-by: Ryan Hamilton <rch@chromium.org>\nCommit-Queue: Chris Davis <chrdavis@microsoft.com>\nCr-Commit-Position: refs/heads/main@{#1332574}",
    "author": "Chris Davis",
    "author_email": "chrdavis@microsoft.com",
    "date": "2024-07-24T21:35:54+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0e794a5aa3653912fafdee665873613a6942798d"
  },
  {
    "id": 119,
    "message": "Delete unimplemented/unused methods in the RC-related files\n\nBug: b/344593927\nChange-Id: If0039117d160ef515110da588ce2c81ad83df3c6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736355\nReviewed-by: Dominique Fauteux-Chapleau <domfc@chromium.org>\nCommit-Queue: Dominique Fauteux-Chapleau <domfc@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332573}",
    "author": "Nancy Xiao",
    "author_email": "nancylanxiao@google.com",
    "date": "2024-07-24T21:35:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3fec1ff55cfc3cdec9d255a6a2cfd6fbe7e0087c"
  },
  {
    "id": 120,
    "message": "Use CreateFromCanvasSharedImage() in CreateImageFromVideoFrame()\n\nAs part of the ClientSharedImage refactoring, this CL lets\nCreateImageFromVideoFrame() use CreateFromCanvasSharedImage() to create\nthe StaticBitmapImage when the video frame has ClientSharedImage.\n\nBug: 1494911\nChange-Id: I368bfb3cbae66a12c6c18032586a08c565a34feb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738895\nCommit-Queue: Mingjing Zhang <mjzhang@chromium.org>\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332572}",
    "author": "Mingjing Zhang",
    "author_email": "mjzhang@chromium.org",
    "date": "2024-07-24T21:33:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a0811d0b58c8e25fbc0992785992150beb0db068"
  },
  {
    "id": 121,
    "message": "Bump Conversions.DebugReport histograms\n\nBug: 332314309\nChange-Id: Ie1d734a0612d2f81672d767477909ce3f426f501\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738953\nCommit-Queue: John Delaney <johnidel@chromium.org>\nAuto-Submit: Anthony Garant <anthonygarant@chromium.org>\nReviewed-by: John Delaney <johnidel@chromium.org>\nCommit-Queue: Anthony Garant <anthonygarant@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332571}",
    "author": "Anthony Garant",
    "author_email": "anthonygarant@chromium.org",
    "date": "2024-07-24T21:27:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fe2b63a86d3a681e27458a97033a0e8850f46706"
  },
  {
    "id": 122,
    "message": "Roll Chrome Linux PGO Profile\n\nRoll Chrome Linux PGO profile from chrome-linux-main-1721821589-eceab329c4db37d41ba47923e16e9a41fdcfd369-0c2054dfadd09f54ab68446c63976579ade7525a.profdata to chrome-linux-main-1721843954-40a71aff8d29217d10cab3b8cceb8208543c36dd-3f75879e72af7aa4330999a92b65a0caa78fee24.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-linux-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I14c11fc52c5a92316e5be5458f007dc6ea908fac\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738640\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332570}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-24T21:27:14+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2573dc4cd3633a4fcf7d18e6f6cdc99566e3c125"
  },
  {
    "id": 123,
    "message": "Create a method to add enterprise reports to a Delivery\n\nCreate AddEnterpriseReports() to add enterprise reports to a Delivery.\nThis will be used for enterprise reporting.\n\ngo/3pcd-enterprise-reporting (Googlers-only)\n\nBug: 353981785\nChange-Id: I2baaab16acd36de5355f1640721b5f3270729536\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5725753\nCommit-Queue: Aaron Selya <selya@google.com>\nReviewed-by: Maks Orlovich <morlovich@chromium.org>\nAuto-Submit: Erin Stewart <erinstewart@google.com>\nReviewed-by: Aaron Selya <selya@google.com>\nCr-Commit-Position: refs/heads/main@{#1332569}",
    "author": "Erin Stewart",
    "author_email": "erinstewart@google.com",
    "date": "2024-07-24T21:25:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b40c386c77c89456065989ff7a54b3d6abd191f1"
  },
  {
    "id": 157,
    "message": "[windows] Add instructions for building mini_installer\n\nChange-Id: I5047acb2f36d6aac6eec75e11848e22de26b550f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713610\nCommit-Queue: Greg Thompson <grt@chromium.org>\nReviewed-by: Aaron Leventhal <aleventhal@chromium.org>\nAuto-Submit: Greg Thompson <grt@chromium.org>\nCommit-Queue: Aaron Leventhal <aleventhal@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328135}",
    "author": "Greg Thompson",
    "author_email": "grt@chromium.org",
    "date": "2024-07-16T15:03:33+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/530c2a98c0b4914b05cf7134a56df243d8e1f730"
  },
  {
    "id": 124,
    "message": "Don't transitively depend on `//printing:metafile` from `//tp/b/p/mojom:mojom_platform`\n\nIn a follow up change we are trying to introduce dependency from\n`//services/viz/privileged/mojom/compositing:compositing` to\n`//third_party/blink/public/common:common`. This fails to compile on\nCrOS targets[1] due to a cyclic dependency. This change breaks this\ndependency by refactoring out \"//printing:printing\".\n\n[1] https://ci.chromium.org/ui/p/chromium/builders/try/lacros-amd64-generic-rel-gtest/158511/overview\n\nBug: b/353667454\nChange-Id: Ib4a97ccffd61c2cecbf66b62ef02bd4c5dcbd90f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735389\nReviewed-by: Dave Tapuska <dtapuska@chromium.org>\nReviewed-by: Lei Zhang <thestig@chromium.org>\nCommit-Queue: Kartar Singh <kartarsingh@google.com>\nCr-Commit-Position: refs/heads/main@{#1332568}",
    "author": "Kartar Singh",
    "author_email": "kartarsingh@google.com",
    "date": "2024-07-24T21:23:51+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4ea0a1c89de48c521ef92c966aafd25a304f460b"
  },
  {
    "id": 125,
    "message": "cryptohome: Update Cryptohome metrics\n\nRemoval of obsolete TimeToAuthSessionRemoveAuthFactorVK,\nupdate of other metrics set to expire soon.\n\nBUG=b:347710123, b:354713125\nTEST=validate_format.py\n\nOBSOLETE_HISTOGRAM[Cryptohome.TimeToAuthSessionRemoveAuthFactorVK]=Removed metric code after mostly complete transition from VaultKeyets to AuthFactor in cryptohome\n\nChange-Id: I741c95d0fad37d1cfe2b19ceb975d808489c271f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735747\nAuto-Submit: Thomas Cedeno <thomascedeno@google.com>\nReviewed-by: John Admanski <jadmanski@chromium.org>\nReviewed-by: Maksim Ivanov <emaxx@chromium.org>\nCommit-Queue: Thomas Cedeno <thomascedeno@google.com>\nCr-Commit-Position: refs/heads/main@{#1332567}",
    "author": "Thomas Cedeno",
    "author_email": "thomascedeno@google.com",
    "date": "2024-07-24T21:23:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1e50c784a9869cb937fa9cf5b58fda510b5b54a9"
  },
  {
    "id": 126,
    "message": "[gpu] Remove unused BufferPlane util methods\n\nRemove BufferPlane utility methods ie. BufferPlaneToString,\nGetPlaneSize, GetPlaneIndex, GetPlaneBufferFormat and\nIsPlaneValidForGpuMemoryBufferFormat that were used with legacy\nmultiplanar shared images with GpuMemoryBuffers that has been removed.\n\nBug: 332564976\nChange-Id: Ife2ff480f681f019dccb3304152adefbc8a6a362\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735671\nReviewed-by: Scott Violet <sky@chromium.org>\nCommit-Queue: Saifuddin Hitawala <hitawala@chromium.org>\nReviewed-by: vikas soni <vikassoni@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332566}",
    "author": "Saifuddin Hitawala",
    "author_email": "hitawala@chromium.org",
    "date": "2024-07-24T21:20:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/717258e097909bc6b9c470dd5eaff0b8072116bd"
  },
  {
    "id": 127,
    "message": "[declutter] Enable archive & auto-delete via finch params\n\nBug: 353337191\nChange-Id: Ic0e1990da159c595fdbfce8de28dd55ce18980d8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5719252\nCommit-Queue: Brandon Wylie <wylieb@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nAuto-Submit: Brandon Wylie <wylieb@google.com>\nReviewed-by: Calder Kitagawa <ckitagawa@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332565}",
    "author": "Brandon Wylie",
    "author_email": "wylieb@chromium.org",
    "date": "2024-07-24T21:19:18+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d0680d36fd3c5a4981506285997605341bf28e6d"
  },
  {
    "id": 128,
    "message": "Reland \"jni_zero: add a hash to ensure multiplexing lines up\"\n\nThis reverts commit aa2445185c97ecbf9afb646e9f2d8cd44e0c37f6.\n\nReason for revert: Monochrome shares the J/N.java, so we also need\nto check if Java's priority targets match C++'s all targets. PS#1->2\n\nOriginal change's description:\n> Revert \"jni_zero: add a hash to ensure multiplexing lines up\"\n>\n> This reverts commit 10c89dc6adc3dabdb1101b63f1414a6a5121d85a.\n>\n> Reason for revert: crbug.com/354025384\n>\n> Original change's description:\n> > jni_zero: add a hash to ensure multiplexing lines up\n> >\n> > Bug: 348183775\n> > Change-Id: I91c0dbeca74a38159bd13cdc8381aab485a225be\n> > Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5680442\n> > Reviewed-by: Mohannad Farrag <aymanm@google.com>\n> > Commit-Queue: Sam Maier <smaier@chromium.org>\n> > Auto-Submit: Sam Maier <smaier@chromium.org>\n> > Cr-Commit-Position: refs/heads/main@{#1329481}\n>\n> Bug: 348183775\n> Change-Id: Ia0a2ae11c3cacc0c333e9860c76bd92beb7368e7\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5731035\n> Commit-Queue: Krishna Govind <govind@chromium.org>\n> Owners-Override: Krishna Govind <govind@chromium.org>\n> Bot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\n> Reviewed-by: Sam Maier <smaier@chromium.org>\n> Reviewed-by: Krishna Govind <govind@chromium.org>\n> Cr-Commit-Position: refs/heads/main@{#1331338}\n\nBug: 348183775\nChange-Id: Id0c0f5c20701b2e32b8c674f94d1616d4f5b317e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735732\nReviewed-by: Andrew Grieve <agrieve@chromium.org>\nOwners-Override: Sam Maier <smaier@chromium.org>\nCommit-Queue: Sam Maier <smaier@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332564}",
    "author": "Sam Maier",
    "author_email": "smaier@chromium.org",
    "date": "2024-07-24T21:18:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/970be2aec7448c8c6eeaa3ba5eac68a80376647f"
  },
  {
    "id": 129,
    "message": "graphite: Fix max_texture_size in PaintOpBuffer::Playback\n\nPaintOpBuffer::Playback was querying max texture size from the canvas'\nGanesh recording context. Update it to use the canvas' Graphite recorder\nwhen running under Graphite.\n\nBug: 337980338\nChange-Id: I8c47a9f62fd155bbb09626e490debfdae6ee589d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736079\nReviewed-by: Saifuddin Hitawala <hitawala@chromium.org>\nAuto-Submit: Sunny Sachanandani <sunnyps@chromium.org>\nCommit-Queue: Saifuddin Hitawala <hitawala@chromium.org>\nCommit-Queue: Sunny Sachanandani <sunnyps@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332563}",
    "author": "Sunny Sachanandani",
    "author_email": "sunnyps@chromium.org",
    "date": "2024-07-24T21:17:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/738293c08a4ba0653e0ff2ecb66b7406ebc5c23b"
  },
  {
    "id": 130,
    "message": "p13n: format sea_pen_provider_base.cc\n\nManual format the file as git cl format didn't work properly in the\nprevious CL.\n\nBUG=b:349441414\nTEST=none\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nChange-Id: I1ade285f4586665034c3c289f0709a7ad40f66eb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735006\nReviewed-by: Jeffrey Young <cowmoo@google.com>\nCommit-Queue: Jason Thai <jasontt@chromium.org>\nAuto-Submit: Jason Thai <jasontt@chromium.org>\nCommit-Queue: Jeffrey Young <cowmoo@google.com>\nCr-Commit-Position: refs/heads/main@{#1332562}",
    "author": "Jason Thai",
    "author_email": "jasontt@chromium.org",
    "date": "2024-07-24T21:17:48+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/120a9caca649397448e0ef75ca1e9d7d3f0a450e"
  },
  {
    "id": 158,
    "message": "Implement FloatingSsoSyncBridge::GetDataForCommit\n\nBug: b:346354248\nChange-Id: I2a2de5cd87e35e409b1db9b7ea0e4632da7a7af9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5703519\nReviewed-by: Maria Petrisor <mpetrisor@chromium.org>\nReviewed-by: Marc Treib <treib@chromium.org>\nCommit-Queue: Andrey Davydov <andreydav@google.com>\nCr-Commit-Position: refs/heads/main@{#1328134}",
    "author": "Andrey Davydov",
    "author_email": "andreydav@google.com",
    "date": "2024-07-16T15:02:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3f68f78870b9d2a2e35f3f3cd753844fd73974a5"
  },
  {
    "id": 131,
    "message": "Add basic HTTP/2 support in HttpStreamPool\n\nThis CL implements basic HTTP/2 support in HttpStreamPool:\n* HttpStreamPool is owned by HttpNetworkSession. To avoid dangling\n  pointers, it should be destroyed after SpdySessionPool.\n* Add RequestStream() method in HttpStreamPool. It returns an\n  HttpStream synchronously when there is an active SPDY session,\n  otherwise, delegates stream creation to HttpStreamPool::Group.\n* Once a SPDY session is successfully created while attempting\n  connections, HttpStreamPool::Job cancels in-flight attempts and\n  idle connections. All on-going requests receive HttpStreams\n  on top of the created SPDY session.\n\nThis CL doesn't implement IP-based pooling. Subsequent CLs should\nadd IP-based pooling support.\n\nBug: 346835898\nChange-Id: I5c91d79f1685eafd1f70bc85e48cfba3295ccca6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5725636\nCommit-Queue: Kenichi Ishibashi <bashi@chromium.org>\nReviewed-by: Adam Rice <ricea@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332561}",
    "author": "Kenichi Ishibashi",
    "author_email": "bashi@chromium.org",
    "date": "2024-07-24T21:16:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1ede3e46663856fc4742d7cbd3e56e3a3ab33e78"
  },
  {
    "id": 132,
    "message": "[SHoA] Update strings for update checks\n\n- Update title and summary for update module states.\n- Add an UNAVAILABLE state when update check is not fetched.\n- UNAVAILABLE state should not show \"all good\" indicator.\n\nBug: 324562205\nChange-Id: I7d6bcaa72c8b835ddfadcbdb363089380a3ca7a3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5710150\nReviewed-by: Rubin Deliallisi <rubindl@chromium.org>\nCommit-Queue: Zaina Al-Mashni <zalmashni@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328160}",
    "author": "Zaina Al-Mashni",
    "author_email": "zalmashni@google.com",
    "date": "2024-07-16T15:52:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/36620b8761c08a2b79e58ed99da241e94e8423ac"
  },
  {
    "id": 133,
    "message": "[PaymentRequest] Fix histograms.xml for IsCanMakePaymentAllowedByPref\n\nIn https://chromium-review.googlesource.com/c/chromium/src/+/5655392,\na histogram was accidentally given the enum type 'Boolean' instead\nof the correct 'CanMakePaymentPreferenceSetter'. This resulted in\nincorrect rendering on the UMA dashboard. This CL corrects the mistake.\n\nBug: 349411309\nChange-Id: Idd1289de1704cc82858d5e48ac14a3e5e3f21cb8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713488\nReviewed-by: Kavita Soni <kavitasoni@chromium.org>\nCommit-Queue: Stephen McGruer <smcgruer@chromium.org>\nAuto-Submit: Stephen McGruer <smcgruer@chromium.org>\nCommit-Queue: Kavita Soni <kavitasoni@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328159}",
    "author": "Stephen McGruer",
    "author_email": "smcgruer@chromium.org",
    "date": "2024-07-16T15:47:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6f8147cc50a4d07b1150087e4888756d4c45143e"
  },
  {
    "id": 134,
    "message": "[visited-url-ranking] Fix `total_foreground_duration` signal\n\nFixes an issue where sentinel values for `total_foreground_duration`\nvalues are being aggregated and used to compute a signal value.\n\nBug: 353336975\nChange-Id: I4746483fcba058fd1ca1ca2729a1d69b9b116651\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709470\nReviewed-by: Siddhartha S <ssid@chromium.org>\nCommit-Queue: Roman Arora <romanarora@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328158}",
    "author": "Roman Arora",
    "author_email": "romanarora@chromium.org",
    "date": "2024-07-16T15:41:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2a8374a85ce5a2c455298364e2720426cb4c2ad3"
  },
  {
    "id": 135,
    "message": "[sync] Expose ModelTypeController::GetModelTypeLocalDataBatchUploader()\n\nThe new API is not used yet, so no behavior change. The empty\nBOOKMARKS, READING_LIST, PASSWORDS implementations of\nModelTypeLocalDataBatchUploader are hooked to the new method.\n\nBug: 40072432\nChange-Id: I9af014fc166f804ae4350b76c410cce97f3de824\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711850\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Victor Vianna <victorvianna@google.com>\nReviewed-by: Mikel Astiz <mastiz@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328157}",
    "author": "Victor Hugo Vianna Silva",
    "author_email": "victorvianna@google.com",
    "date": "2024-07-16T15:36:17+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4efa9e51b75b8a02a3d892a80cbca62bd6052c3d"
  },
  {
    "id": 136,
    "message": "Create a custom layout for the PinnedToolbarActionsContainer\n\nThis is needed since the container has a divider that's visibility is\ndependant on there being any pinned buttons included in the layout to\nthe left. Currently our layouts cannot support this.\n\nBug: 348453824\nChange-Id: I5d670ab2275ec89f1a37d7f59318e3fed9aaa2e1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5676468\nReviewed-by: Dana Fried <dfried@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Caroline Rising <corising@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328156}",
    "author": "Caroline Rising",
    "author_email": "corising@chromium.org",
    "date": "2024-07-16T15:33:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c3c5f30b20e766d4d127269bdc28fc8a00a71e06"
  },
  {
    "id": 137,
    "message": "Roll Dawn from 16be0748ef57 to 72f2b14be5bb (2 revisions)\n\nhttps://dawn.googlesource.com/dawn.git/+log/16be0748ef57..72f2b14be5bb\n\n2024-07-16 cwallez@chromium.org [Kotlin]: Introduce a JNIContext helper, use it to free JNI allocations\n2024-07-16 cwallez@chromium.org generator_lib.py: Add support for importing templates in templates\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/dawn-chromium-autoroll\nPlease CC cwallez@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Dawn: https://bugs.chromium.org/p/dawn/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:dawn-android-arm-deps-rel;luci.chromium.try:dawn-android-arm64-deps-rel;luci.chromium.try:dawn-linux-x64-deps-rel;luci.chromium.try:dawn-mac-x64-deps-rel;luci.chromium.try:dawn-mac-arm64-deps-rel;luci.chromium.try:dawn-win10-x64-deps-rel;luci.chromium.try:dawn-win10-x86-deps-rel;luci.chromium.try:dawn-win11-arm64-deps-rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:gpu-fyi-cq-android-arm64\nBug: chromium:330293719\nTbr: cwallez@google.com\nInclude-Ci-Only-Tests: true\nChange-Id: Ie7591a1fc94eb5c97d25826e172a236628f57335\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712790\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328155}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T15:29:07+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3b53d7e25145f6dcc26262b181863f276a812f71"
  },
  {
    "id": 138,
    "message": "[Chromoting] Add proto defs for UpdateRemoteHostRequest\n\nChange-Id: I25563a6af0cbe060880f676ec4fcf637ae3cd69e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5710175\nReviewed-by: Joe Downing <joedow@chromium.org>\nCommit-Queue: Joe Downing <joedow@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328154}",
    "author": "Gary Kacmarcik",
    "author_email": "garykac@google.com",
    "date": "2024-07-16T15:28:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f224338e6fa0866fc1d8072bbe908f31fe6857dc"
  },
  {
    "id": 141,
    "message": "Roll libc++ from 81d177a0a844 to 928433a141bd (4 revisions)\n\nhttps://chromium.googlesource.com/external/github.com/llvm/llvm-project/libcxx.git/+log/81d177a0a844..928433a141bd\n\n2024-07-16 nikolasklauser@berlin.de [libc++] Simplify the implementation of is_null_pointer a bit (#98728)\n2024-07-16 nikolasklauser@berlin.de [libc++] Merge is_member{,_object,_function}_pointer.h (#98727)\n2024-07-15 siujoeng.lau@gmail.com [libc++] P2389R2: `dextents` Index Type Parameter (#97393)\n2024-07-15 ldionne.2@gmail.com [libc++] Handle _LIBCPP_HAS_NO_{THREADS,LOCALIZATION} consistently with other carve-outs (#98319)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/libcxx-chromium\nPlease CC hans@chromium.org,thakis@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: hans@chromium.org,thakis@chromium.org\nChange-Id: I14de40da6deff51074da4fa59acc6147c5b4bd9e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712614\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328151}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T15:25:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/196d34193b24d3b58ad34f385490e4e8962de61c"
  },
  {
    "id": 142,
    "message": "Tracking Protection 100%: Add tests for managed icons\n\nBug: 330745124\nChange-Id: Id951429ff9084ffa5dc2e17a3c9c0872160d3a52\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713029\nCommit-Queue: Andrey Zaytsev <andzaytsev@google.com>\nReviewed-by: Matt Dembski <dembski@google.com>\nCr-Commit-Position: refs/heads/main@{#1328150}",
    "author": "Andrey Zaytsev",
    "author_email": "andzaytsev@google.com",
    "date": "2024-07-16T15:23:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b78f006735fde4d887cf09c82e3b1f35a6b30525"
  },
  {
    "id": 143,
    "message": "Add clear-undecryptable-passwords flag to ios\n\nBug: 40286735\nChange-Id: I1fa3d9d4abae70978d751cfae65167a6c457f753\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712689\nReviewed-by: Vasilii Sukhanov <vasilii@chromium.org>\nCommit-Queue: Karol Sygiet <sygiet@google.com>\nCr-Commit-Position: refs/heads/main@{#1328149}",
    "author": "Karol Sygiet",
    "author_email": "sygiet@google.com",
    "date": "2024-07-16T15:22:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6ac82c84ea702e5f088466785ac8013ce5ef7377"
  },
  {
    "id": 144,
    "message": "[iOS] Move cellForItemAtIndexPath to protected\n\nThis CL makes cellForItemAtIndexPath:itemIdentifier: protected on the\nBaseGridViewController so the subclass can override it.\nIt also moves the registration of the InactiveTabsButton to the\nsubclass, allowing to configure it with the different infos in the\nfuture.\n\nBug: 352722446\nChange-Id: Ib817ad7feae688959cdaf46376f83a85d3d5e157\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713251\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Louis Romero <lpromero@google.com>\nAuto-Submit: Gauthier Ambard <gambard@chromium.org>\nCommit-Queue: Gauthier Ambard <gambard@chromium.org>\nCommit-Queue: Louis Romero <lpromero@google.com>\nCr-Commit-Position: refs/heads/main@{#1328148}",
    "author": "Gauthier Ambard",
    "author_email": "gambard@chromium.org",
    "date": "2024-07-16T15:22:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4e9e98dc51ae8838ae1dc6cd9f5f1fddb50f7340"
  },
  {
    "id": 145,
    "message": "[Autofill] Factor some common payments test code into a new utils file\n\nThis CL adds a new utils file for Autofill Payments tests specifically,\nautofill_payments_test_utils.h, which factors some common code out of\npayments_network_interface_unittest.cc and\nupload_card_request_unittest.cc. It will also be used in future CLs\nintroducing more payments_requests/ unittests.\n\nBug: None\nChange-Id: I86653bdd4c72128810e0f267227bfb123fb65c27\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5708648\nReviewed-by: Vinny Persky <vinnypersky@google.com>\nCommit-Queue: Stephen McGruer <smcgruer@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328147}",
    "author": "Stephen McGruer",
    "author_email": "smcgruer@chromium.org",
    "date": "2024-07-16T15:20:29+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b23daf4ae8472bbd18544717c533067d04efe8ff"
  },
  {
    "id": 146,
    "message": "Roll Perfetto Trace Processor Linux from 2535d757a6bf to 63560013c817\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/2535d757a6bf..63560013c817\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-trace-processor-linux-chromium\nPlease CC chrometto-team@google.com,perfetto-bugs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: perfetto-bugs@google.com\nChange-Id: I77e619f10c561e1ce9185dc72735494c236acba0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712849\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328146}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T15:17:09+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1cf066c80fcf0fc2208347718c8b70b9c5b994b0"
  },
  {
    "id": 147,
    "message": "Move known unexpected Fend behaviour into disabled test\n\nThese tests are effectively change detectors, which make automatic\nfend-core rolls difficult due to test breakages. Move these tests into a\ndisabled parameterised test so they can be safely ignored in CQ.\n\nSee https://crrev.com/c/5707389/comments/dfa4a328_f797804a for more\ndetails.\n\nBug: b:345610113\nChange-Id: I6a6a21002cae7a1339c01ee6744618f7dfdd34ad\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711588\nReviewed-by: Dustin Mitchell <djmitche@chromium.org>\nCommit-Queue: Dustin Mitchell <djmitche@chromium.org>\nAuto-Submit: Michael Cui <mlcui@google.com>\nCr-Commit-Position: refs/heads/main@{#1328145}",
    "author": "mlcui",
    "author_email": "mlcui@google.com",
    "date": "2024-07-16T15:16:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5d2ae784596b7748b5a0a22bffcd6115347e274a"
  },
  {
    "id": 159,
    "message": "Updating XTBs based on .GRDs from branch main\n\nChange-Id: I2487e14ec1ca0139fec2d3e19996b6a4c8e1a701\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713368\nAuto-Submit: Ben Mason <benmason@chromium.org>\nBot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCommit-Queue: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328133}",
    "author": "Ben Mason",
    "author_email": "benmason@chromium.org",
    "date": "2024-07-16T14:57:18+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/07a3fffd7a781839a2fcf992f3c04e393f4ffd26"
  },
  {
    "id": 148,
    "message": "Add methods for connecting to and dialing CECA\n\nIntroduce methods to the client library for connecting to and dialing\nthe Mojo service. Additionally, the service stub unittests have been\nupdated to use this client implementation.\n\nThe runtime of the test suite has been improved by immediately returning\nfrom the multi-process test child if an IPC callback was dropped.\nPreviously, tests in which the caller was rejected as untrustworthy\ndepended on no action being observed over a time period.\n\nBug: 342180612\nChange-Id: Ibf338969ed5b4fef0caac1b8000516038f7f4962\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5691032\nReviewed-by: Sorin Jianu <sorin@chromium.org>\nCommit-Queue: Noah Rose Ledesma <noahrose@google.com>\nCr-Commit-Position: refs/heads/main@{#1328144}",
    "author": "Noah Rose Ledesma",
    "author_email": "noahrose@google.com",
    "date": "2024-07-16T15:14:07+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9ad8ff11de051d8f720d79260184ffc1285db367"
  },
  {
    "id": 149,
    "message": "[a11y] Update rebase script's output\n\nNow outputs the changed expectation files as a sorted list, per request.\n\nAlso prints out a notification \".\" of each file changed in realtime.\n\nChange-Id: I00a94c1ba8b5c1ce938736a5bf6a215e6b566123\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5703538\nCommit-Queue: Mark Schillaci <mschillaci@google.com>\nAuto-Submit: Jonny Wang <jonnywang@google.com>\nReviewed-by: Mark Schillaci <mschillaci@google.com>\nReviewed-by: Aaron Leventhal <aleventhal@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328143}",
    "author": "Jonny Wang",
    "author_email": "jonnywang@google.com",
    "date": "2024-07-16T15:10:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/83bcd944cd58ad6665c84b8a0b36f4cb131adb19"
  },
  {
    "id": 150,
    "message": "Remove DIPS database prepopulation\n\nBug: 336563548\nChange-Id: I27214cd060f60a9b4a1d7390bd39b243fdb24db5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709591\nReviewed-by: Anton Maliev <amaliev@chromium.org>\nCommit-Queue: Andrew Liu <liu@chromium.org>\nReviewed-by: Alex Rudenko <alexrudenko@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328142}",
    "author": "Andrew Liu",
    "author_email": "liu@chromium.org",
    "date": "2024-07-16T15:10:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3f168b5e3551c2355cdb5dc0afa9f971883cc32e"
  },
  {
    "id": 151,
    "message": "[ios] Fix test testLeaveSwitcherAfterEnteringAndExittingInactiveTabs\n\nThis CL adds a workaround to fix test in\nTabSwitcherTransitionTestCase when run on iOS18 devices.\n\nBug: 349816094\nChange-Id: I97c98787fc53c2abab43a0dfd27334ff6cf2dbe4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713310\nAuto-Submit: Federica Germinario <fedegermi@google.com>\nCommit-Queue: Federica Germinario <fedegermi@google.com>\nReviewed-by: Louis Romero <lpromero@google.com>\nCommit-Queue: Louis Romero <lpromero@google.com>\nCr-Commit-Position: refs/heads/main@{#1328141}",
    "author": "Federica Germinario",
    "author_email": "fedegermi@google.com",
    "date": "2024-07-16T15:07:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/84370d74ba652da29b093155e71d69b6ec66a1d9"
  },
  {
    "id": 152,
    "message": "Handle abandoned callbacks in IpProtectionAuthClient\n\nBug: b:343937808\nChange-Id: I06037f9dc3e33dc75f8d26744e3ea542969418e8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5665312\nReviewed-by: Abhijith Nair <abhijithnair@chromium.org>\nCommit-Queue: Ashley Newson <ashleynewson@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328140}",
    "author": "Ashley Newson",
    "author_email": "ashleynewson@chromium.org",
    "date": "2024-07-16T15:07:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8e0d4adc53fc69193929970481fb5b24763a5b8a"
  },
  {
    "id": 153,
    "message": "[headless] Re-enable proxy config tests\n\nAfter https://crrev.com/c/5637879 and\nhttps://crrev.com/c/5640568 landed these tests appear to\nbe stable.\n\nBug: 40694526\nChange-Id: Ifaff39abb054dc745d26802b00b277f50d8a247e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5710233\nCommit-Queue: Peter Kvitek <kvitekp@chromium.org>\nReviewed-by: Eric Seckler <eseckler@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328139}",
    "author": "Peter Kvitek",
    "author_email": "kvitekp@chromium.org",
    "date": "2024-07-16T15:06:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/acdbe193486415c285e2ed3b6ec91ef5558bc4fd"
  },
  {
    "id": 154,
    "message": "Roll Skia from 07bd32512db2 to 57ae89c0ffae (1 revision)\n\nhttps://skia.googlesource.com/skia.git/+log/07bd32512db2..57ae89c0ffae\n\n2024-07-16 robertphillips@google.com [graphite] Add BackendApiToStr and use it in the PaintParamsKeyTest\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/skia-autoroll\nPlease CC nicolettep@google.com,skiabot@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Skia: https://bugs.chromium.org/p/skia/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux-blink-rel;luci.chromium.try:linux-chromeos-compile-dbg;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nCq-Do-Not-Cancel-Tryjobs: true\nBug: None\nTbr: nicolettep@google.com\nChange-Id: Iadf2044f74590262739a87e7f6dd0fef67907598\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713090\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328138}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T15:05:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3bd821e5c87f3762c54e2af00a20b131d105f6b9"
  },
  {
    "id": 155,
    "message": "[iOS] Use Inactive Tabs Button item in Grid\n\nThis CL adds the inactive tabs button item in the grid when necessary.\nFor now it is using a dummy cell.\n\nIn order to be able to display the cell, it is also creating a first\nversion of the section in the compositional layout.\n\nIt is also doing a light cleanup:\n- Move the registrations to ivar\n- Reorder the registrations creation\n- Remove dead \"action\" code in the BaseGrid\n- Add \"inactive tab button\" as part of the sections that can be tapped\n\nBug: 352722446\nChange-Id: If95d5a73c7a169df43cfaf5f95808f195ae62a85\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713228\nReviewed-by: Louis Romero <lpromero@google.com>\nCommit-Queue: Louis Romero <lpromero@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nAuto-Submit: Gauthier Ambard <gambard@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328137}",
    "author": "Gauthier Ambard",
    "author_email": "gambard@chromium.org",
    "date": "2024-07-16T15:05:01+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8e4808fcdc9343c9dd86bff28ed90706131f84f0"
  },
  {
    "id": 160,
    "message": "[oobe] Get user preferred locale from PeopleAPI\n\nBug: b:330384267\nChange-Id: Ib4cde8c023a57f841abfffb23da2a0afad4917b8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5645663\nCommit-Queue: Danila Kuzmin <dkuzmin@google.com>\nReviewed-by: David Roger <droger@chromium.org>\nReviewed-by: Bohdan Tyshchenko <bohdanty@google.com>\nReviewed-by: Side YILMAZ <sideyilmaz@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328132}",
    "author": "Danila Kuzmin",
    "author_email": "dkuzmin@google.com",
    "date": "2024-07-16T14:56:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7efc94e40a765846f3cbd8a3ba14d5b4564cfc2e"
  },
  {
    "id": 161,
    "message": "[Gardener] Disable flaky testActivityDoesNotLeak\n\nThe AwContentsGarbageCollectionTest#testActivityDoesNotLeak has average\nflakiness in runs 11.25%.\n\nBug: 353484967\nChange-Id: Ic29dd36bcb48730c68c2b1db758ef3b6de68e075\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713309\nAuto-Submit: Ivana Žužić <izuzic@google.com>\nOwners-Override: Ivana Žužić <izuzic@google.com>\nCommit-Queue: Theo Cristea <theocristea@google.com>\nReviewed-by: Theo Cristea <theocristea@google.com>\nCr-Commit-Position: refs/heads/main@{#1328131}",
    "author": "Ivana Žužić",
    "author_email": "izuzic@google.com",
    "date": "2024-07-16T14:56:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/eb8d0341b947ec9ae587f2568f88e45eb6b8dfb6"
  },
  {
    "id": 162,
    "message": "Extend Extension Telemetry FileData histograms.\n\nBug: 346342130\nChange-Id: I5e24c7d3182818e6299054aea98a388e0aa14abe\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709652\nCommit-Queue: Richard Chen <richche@google.com>\nReviewed-by: thefrog <thefrog@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328130}",
    "author": "Richard Chen",
    "author_email": "richche@chromium.org",
    "date": "2024-07-16T14:55:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7856f51ebf623084341fd982aa91c89befe4642f"
  },
  {
    "id": 163,
    "message": "Resolve CSSGradientValue with CSSToLengthConversionData\n\nReplace calls for Get{Double/Float}Value by Compute{Number,Percentage}\nwith conversion data for CSSGradientValue.\n\nBug: 40946458\nChange-Id: I095e31c0532c1be8279f11ab0cf9dc8c67aae37b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5694260\nReviewed-by: Morten Stenshorne <mstensho@chromium.org>\nCommit-Queue: Daniil Sakhapov <sakhapov@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328129}",
    "author": "Daniil Sakhapov",
    "author_email": "sakhapov@chromium.org",
    "date": "2024-07-16T14:53:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d7fd37e9a6bb572eea7cf976a8a870d1c17e863e"
  },
  {
    "id": 164,
    "message": "[Reauth] Introduce biometric auth for cred man controller\n\nBiometric auth should be triggered when filling passwords from CredMan too. CL triggers re-auth in CredManController before filling a username/password when required.\n\nBug: 343879727\nChange-Id: I9dc081ac9d48b44cc8048049ab2b15d9a152d9a5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711873\nCommit-Queue: Anna Tsvirchkova <atsvirchkova@google.com>\nReviewed-by: Ioana Pandele <ioanap@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328128}",
    "author": "Anna Tsvirchkova",
    "author_email": "atsvirchkova@google.com",
    "date": "2024-07-16T14:52:09+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b182c6ec2513060e31dce8295a9619287b49ea8e"
  },
  {
    "id": 165,
    "message": "Roll clank/internal/apps from 973f3ee91312 to 8bf4ccb93387 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/973f3ee91312..8bf4ccb93387\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,izuzic@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: izuzic@google.com\nNo-Try: true\nChange-Id: Ief72395c2a381106cc748b70b1284f9664c79d6b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713808\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328127}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-16T14:51:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5ad80c4d016a93a33bcc4690ec09f1f62bae1a66"
  },
  {
    "id": 166,
    "message": "Roll ios_internal from 53d0a975b482 to f1079cb86dcc\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/53d0a975b482..f1079cb86dcc\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,danieltwhite@google.com,thegreenfrog@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: Icec342420559a57e6b16703c6af7eb688b027d6b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712616\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328126}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-16T14:50:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/96d0c2bcb98a4f024b07476b58517b613b35cadf"
  },
  {
    "id": 167,
    "message": "[3PCD Eligibility] Implement IsNewProfile\n\nThis CL adds functionality to check if a profile is new or not. This\naffects which onboarding notices are showed to the profile.\n\nBug: b:349788009\n\nChange-Id: I1b320016d718e52f57625bd55bd855191f8462f9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5703137\nReviewed-by: Abe Boujane <boujane@google.com>\nReviewed-by: Fiona Macintosh <fmacintosh@google.com>\nCommit-Queue: Aashna Sheth <aashnas@google.com>\nCr-Commit-Position: refs/heads/main@{#1328125}",
    "author": "Aashna Sheth",
    "author_email": "aashnas@google.com",
    "date": "2024-07-16T14:47:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b2bb9a538c11705a6b3ae210e23adcd7bb2f6e4b"
  },
  {
    "id": 168,
    "message": "[Start] Remove tools/metrics/histograms/metadata/start_surface.\n\nOBSOLETE_HISTOGRAM[StartSurface.ColdStartup.IsLastActiveTabNtp]=Start surface goes away.\nOBSOLETE_HISTOGRAM[StartSurface.Module.Click]=Start surface goes away.\nOBSOLETE_HISTOGRAM[StartSurface.Show.State]=Start surface goes away.\nOBSOLETE_HISTOGRAM[StartSurface.SpareTab.TimeBetweenShowAndCreate]=Start surface goes away.\nOBSOLETE_HISTOGRAM[StartSurface.TimeSpent]=Start surface goes away.\n\nBug: 344651414\nChange-Id: I7d4ab95d148e523164e8cb496265c6993d9ef467\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5704217\nReviewed-by: Caitlin Fischer <caitlinfischer@google.com>\nCommit-Queue: Xi Han <hanxi@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328124}",
    "author": "Xi Han",
    "author_email": "hanxi@google.com",
    "date": "2024-07-16T14:47:14+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9a48da6b97980e489836dbb94dfeb9cc02d6d1eb"
  },
  {
    "id": 169,
    "message": "[SHoA] Add the help center article to the dashboard\n\nWhen clicking on the info icon, open the safety check help center\narticle in CCT.\n\nBug: 324562205\nChange-Id: I67c7c3a33460905ca39ff9965853827d50e82b5b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712768\nCommit-Queue: Zaina Al-Mashni <zalmashni@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nAuto-Submit: Zaina Al-Mashni <zalmashni@google.com>\nCommit-Queue: Rubin Deliallisi <rubindl@chromium.org>\nReviewed-by: Rubin Deliallisi <rubindl@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328123}",
    "author": "Zaina Al-Mashni",
    "author_email": "zalmashni@google.com",
    "date": "2024-07-16T14:45:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f2908a5d13db7ce5f821c67d4e92c36aebc83976"
  },
  {
    "id": 170,
    "message": "Roll Chrome Linux PGO Profile\n\nRoll Chrome Linux PGO profile from chrome-linux-main-1721109360-8c0a193819d641255c121f965be697754a8750af-5395bd516765e60c4cdbdf9f7758522d2f88a382.profdata to chrome-linux-main-1721131103-04be9ecfec2ec668b99a1ecf03df2d5d241a30c8-f1ce5ea1e387d15eb36e31879f11a08cd241e09c.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-linux-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I3bb1aa87e9838c0453bb9058c348f63e2228722c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712554\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328122}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T14:43:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d461b33ee788739e900ef8f3dff2a407d0174bed"
  },
  {
    "id": 171,
    "message": "Fix StorageAccess to work with RemoteFrames.\n\nWhen a page with StorageAccess is loaded into a cross-process frame,\nthe `storage_access_api_status` value is missing. This CL plumbs the\nvalue through the RenderFrameProxyHost, and updates tests to exercise\nboth the same-process and cross-process states.\n\nBug: 40259221\nChange-Id: I63173be82f5b8fd956f2456efb2c4a09c67cc2f8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5669691\nReviewed-by: Brendon Tiszka <tiszka@chromium.org>\nReviewed-by: Charlie Reis <creis@chromium.org>\nReviewed-by: Chris Fredrickson <cfredric@chromium.org>\nCommit-Queue: James Maclean <wjmaclean@chromium.org>\nReviewed-by: Daniel Cheng <dcheng@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328121}",
    "author": "W. James MacLean",
    "author_email": "wjmaclean@chromium.org",
    "date": "2024-07-16T14:42:34+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/443ef3e6db52ee76bbcaaec6a76dbdb0ed38c829"
  },
  {
    "id": 172,
    "message": "[remoting] Remove multi-stream capability from ChromeOS host.\n\nThe capability was added in https://crrev.com/c/5664350 but it is\nstill not working. The crash is fixed, but the video-layout is\nincorrect, and mouse-events are not correctly injected.\n\nThis removes the advertised capability, until these new issues are\nfixed.\n\nBug: 353279144\nChange-Id: Id75056978f54db3fd9795471cd117deb176a70d5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709594\nCommit-Queue: Joe Downing <joedow@chromium.org>\nReviewed-by: Joe Downing <joedow@chromium.org>\nAuto-Submit: Lambros Lambrou <lambroslambrou@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328120}",
    "author": "Lambros Lambrou",
    "author_email": "lambroslambrou@chromium.org",
    "date": "2024-07-16T14:41:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a82b4dd7b2b03b0d58c3e6c3e4fc7dfde1a158eb"
  },
  {
    "id": 173,
    "message": "Remove unnecessary ScopedProfileSelectionsForFactoryTesting usages.\n\nBug: n/a\nChange-Id: Ifb439d02349a87c38bde2e250244251a329f8111\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712308\nReviewed-by: Jan Keitel <jkeitel@google.com>\nCommit-Queue: Timofey Chudakov <tchudakov@google.com>\nCr-Commit-Position: refs/heads/main@{#1328119}",
    "author": "Timofey Chudakov",
    "author_email": "tchudakov@google.com",
    "date": "2024-07-16T14:41:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/69c659f14d31c3c349102cf7d5e55b7a64395e93"
  },
  {
    "id": 174,
    "message": "Roll Chrome Win32 PGO Profile\n\nRoll Chrome Win32 PGO profile from chrome-win32-main-1721109360-70fd5274c70dfc5c3fc17f4d6e54ece9e0435978-5395bd516765e60c4cdbdf9f7758522d2f88a382.profdata to chrome-win32-main-1721120389-8381d2d1fbf270c070309b37c9ee4e4ec2453c51-18cdc55117478bdeba28fa6c5fa3c563613fcd3c.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win32-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:win-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I3b8ec127bb1ddf8ce9b6802de9db2547b629a47a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711799\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328118}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T14:38:07+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/77ad78557522bb01a1f5db02612c7c575bda7abc"
  },
  {
    "id": 175,
    "message": "[Blink] Have CanvasResource::SetMailboxSyncMode just take in bool\n\nThe only thing that `MailboxSyncMode` is used for internally is to set\na bool based on whether the mode is `kVerifiedSyncToken`. This CL\nchanges the method name and signature accordingly. Followup will expand\nthis outward.\n\nBug: 352263194\nChange-Id: I2760693d54a364d1b6eec66bbade18952ff9c774\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5710132\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCommit-Queue: Colin Blundell <blundell@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328117}",
    "author": "Colin Blundell",
    "author_email": "blundell@chromium.org",
    "date": "2024-07-16T14:37:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1bb4a3134d4bfa72bc7ce0c562f31ed7a8d4d4af"
  },
  {
    "id": 176,
    "message": "Remove feature ForceRestartGpuKillSwitch.\n\nIt's been enabled fully for 1.5 years now.\n\nTEST=bots\n\nBug: 1172229\nChange-Id: I3397be0be04d12ea255d59a76539600950dd5e35\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709538\nCommit-Queue: Vasiliy Telezhnikov <vasilyt@chromium.org>\nAuto-Submit: Zhenyao Mo <zmo@chromium.org>\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328116}",
    "author": "Zhenyao Mo",
    "author_email": "zmo@chromium.org",
    "date": "2024-07-16T14:36:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f057f62b5a7c8f3ff340ce6af88fccd8d389af1f"
  },
  {
    "id": 177,
    "message": "[iOS] Clean Up Startup Improvements Feature Flag in Omnibox\n\nClean up startup improvements feature flag related code.\n\nBug: 352833570\nChange-Id: Ifbf47aa36e9e2f8e705c2973097670d925939dca\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5704222\nReviewed-by: Rohit Rao <rohitrao@chromium.org>\nCommit-Queue: Joemer Ramos <joemerramos@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nAuto-Submit: Joemer Ramos <joemerramos@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328115}",
    "author": "Joemer Ramos",
    "author_email": "joemerramos@chromium.org",
    "date": "2024-07-16T14:33:49+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5481e09a4865925a8d1412191695a27ad9fcbd3c"
  },
  {
    "id": 178,
    "message": "Roll devtools-internal from 07a9a068d30b to 363b3fd04d07 (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/07a9a068d30b..363b3fd04d07\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/devtools/devtools-frontend.git/+log/37848592ea7180f71e330a7a54a273e98a9f52b9..f9d2a54729e43674bdec477149e79fabefec3c2e\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: I0663c9e74501476bd52e69053606c319d6787f2a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712549\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328114}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-16T14:32:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/96c98f14742c2bacd46286e12784649bda36bb19"
  },
  {
    "id": 179,
    "message": "Update logic for moving pinned buttons in PinnedToolbarActionsContainer.\n\nIts possible for actions that are 'pinned' in the model to not exist in\nthe toolbar container (ie a user pinned CSC then changed their default\nsearch engine so it is no longer available). This cl updates how we move\nbuttons to be based on what button precedes the button being moved in\nthe ui.\n\nBug: 350752543\nChange-Id: I200bbda5b09470308a50fe81b6a8877f62543b25\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5708191\nCommit-Queue: Caroline Rising <corising@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: David Pennington <dpenning@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328113}",
    "author": "Caroline Rising",
    "author_email": "corising@chromium.org",
    "date": "2024-07-16T14:31:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f524ce704481ce1501c219c0ab123d244e08e8d4"
  },
  {
    "id": 180,
    "message": "Roll out latest changes in FR datasets.\n\nBug: 347859030\nChange-Id: I8d29a632eda918f190aa038f200c6ad3b9b21cd1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713128\nReviewed-by: Florian Leimgruber <fleimgruber@google.com>\nCommit-Queue: Florian Leimgruber <fleimgruber@google.com>\nAuto-Submit: Norge Vizcay <vizcay@google.com>\nCommit-Queue: Norge Vizcay <vizcay@google.com>\nCr-Commit-Position: refs/heads/main@{#1328112}",
    "author": "Norge Vizcay",
    "author_email": "vizcay@google.com",
    "date": "2024-07-16T14:24:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3b9e2678c19b8d88805c4744a2860a0720025ca4"
  },
  {
    "id": 181,
    "message": "[TrackingProtection] Final strings and TP version\n\nTracking protection refactor. In the final code we merge ModeB and FullTP into one segment, but it's important to know that TrackingProtectionOnboardingController is not used for current ModeB. ModeB in this class is used only for backcompat if we'd need it.\n\nBug: 341968245\nChange-Id: I7323b619bfd4562596f4746da0e669614070367f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5703539\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Andrey Zaytsev <andzaytsev@google.com>\nAuto-Submit: Matt Dembski <dembski@google.com>\nCommit-Queue: Matt Dembski <dembski@google.com>\nCr-Commit-Position: refs/heads/main@{#1328111}",
    "author": "Matt Dembski",
    "author_email": "dembski@google.com",
    "date": "2024-07-16T14:24:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0db78424094afe1734e6ad3267c0ab677c3f6937"
  },
  {
    "id": 182,
    "message": "[ios] Add feature flag for proactive password generation bottom sheet\n\nBug: 346533806\nChange-Id: Id79e2cd781aee3306013fad7696539c17c6b48d3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5708450\nReviewed-by: Vincent Boisselle <vincb@google.com>\nReviewed-by: Noémie St-Onge <noemies@google.com>\nCommit-Queue: Charlotte Cloutier <cloutierc@google.com>\nCr-Commit-Position: refs/heads/main@{#1328110}",
    "author": "Charlotte Cloutier",
    "author_email": "cloutierc@google.com",
    "date": "2024-07-16T14:23:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/18f0a5361f1fc210c1415e4699d8dbe505290004"
  },
  {
    "id": 183,
    "message": "Create non-GWS AbandonedPageLoadMetricsObserver for WebView\n\nGWSAbandonedPageLoadMetricsObserver was introduced in\ncrrev.com/c/5631386 but it's currently only registered for\nChrome. We're interested in gathering the metrics on WebView\nas well, so this CL moves the class to components/ and\nregisters it on WebView as well. However, there's still ongoing\ndiscussions about having a GWS observer on WebView, so this CL\nmakes a general AbandonedPageLoadMetricsObserver instead that\nis made a parent class of the GWS-specific\nGWSAbandonedPageLoadMetricsObserver\n\nSee also:\nhttps://docs.google.com/document/d/1LAC_oLTFGqRxN3RlNcKKq52Ob69Zqk4MRJNM-v23Jvg/edit\n\nNO_IFTTT=File relocation\n\nBug: 347706997\nChange-Id: Ib667b26ca86c159797744187a1d6d6dd61c5dcba\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5675788\nReviewed-by: Takashi Toyoshima <toyoshim@chromium.org>\nCommit-Queue: Rakina Zata Amni <rakina@chromium.org>\nReviewed-by: Richard (Torne) Coles <torne@chromium.org>\nReviewed-by: Shunya Shishido <sisidovski@chromium.org>\nReviewed-by: Minoru Chikamune <chikamune@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328109}",
    "author": "Rakina Zata Amni",
    "author_email": "rakina@chromium.org",
    "date": "2024-07-16T14:20:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b6b0b6031ca90b21b458ff560557be9a17cba064"
  },
  {
    "id": 219,
    "message": "Offer plus address filling on affiliated domains in OTR mode.\n\nPrior to this CL, a user will not see filling suggestions for plus\naddresses when they are in incognito mode and there is no exact match.\n\nThis CL includes affiliated matches in determining whether to show\nplus address suggestions.\n\nBug: 353240084\nChange-Id: I790e6c0a026ed1379894a1627aca26c87d01f86f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5707970\nReviewed-by: Norge Vizcay <vizcay@google.com>\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nCr-Commit-Position: refs/heads/main@{#1328073}",
    "author": "Jan Keitel",
    "author_email": "jkeitel@google.com",
    "date": "2024-07-16T12:42:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/65ff8d5909cf8849085ba571c2c293afba0f33bf"
  },
  {
    "id": 184,
    "message": "Add kite badge on secondary Supervised User sign-in\n\nThis change makes the avatar badge of sign-in in intercept bubble\nparametrized, in order to display the enterprize or the supervised\nuser badge conditionally. The decision on the badge to use is made\nin the bubble's handler.\nFor screenshots for the SU badge see: b/352486197#comment2.\n\nThe change also re-enabled the previously disabled\nDiceWebSigninInterceptionBubblePixelTest test suite.\n\nBug: b/352486197\nChange-Id: I982885ebf33e9e3dc6e5402c5d54d8e33fffcba9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5691641\nReviewed-by: David Roger <droger@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Anthi Orfanou <anthie@google.com>\nCr-Commit-Position: refs/heads/main@{#1328108}",
    "author": "anthie@google.com",
    "author_email": "anthie@google.com",
    "date": "2024-07-16T14:19:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3398f35460a2c8bc92a0ef93249777471d9e879a"
  },
  {
    "id": 185,
    "message": "Roll Perfetto from 2535d757a6bf to 63560013c817 (1 revision)\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/2535d757a6bf..63560013c817\n\n2024-07-16 mayzner@google.com Merge \"tp: Add Flatten function.\" into main\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-chromium-autoroll\nPlease CC perfetto-bugs@google.com,primiano@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-perfetto-rel\nBug: None\nTbr: perfetto-bugs@google.com\nChange-Id: I3106fddfbeb51f19d8de8b618cb0c4f4c0f104a3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712550\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328107}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T14:18:14+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8400beb1e4d45d6f0276011aea32de4499c4e6b2"
  },
  {
    "id": 186,
    "message": "[SHoA] Add the managed state for passwords module\n\nCheck if password manager is disabled by policy, if so we should show a summary describing this and hide the secondary buttons. We will still check for the no passwords/compromised passwords state and update the title/icon.\n\nBug: 324562205\nChange-Id: I6d0bc1c9aaf0ca663e61b4c255986155ddf4a4a6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5707949\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nAuto-Submit: Zaina Al-Mashni <zalmashni@google.com>\nReviewed-by: Rubin Deliallisi <rubindl@chromium.org>\nCommit-Queue: Rubin Deliallisi <rubindl@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328106}",
    "author": "Zaina Al-Mashni",
    "author_email": "zalmashni@google.com",
    "date": "2024-07-16T14:15:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/62de987e459823b868d8a3f09d1a9f6ce1ec734b"
  },
  {
    "id": 187,
    "message": "Fix css math function clamp()\n\nSince in C++ -0.0 == 0.0, std::min(0.0, -0.0) returns 0, which\nleads to accumulating error and results in failing of some tests.\nFix it by manually checking such situations.\n\nChange-Id: I542cbb83817c4d5bf75db2763dc230e17381b988\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5707994\nCommit-Queue: Daniil Sakhapov <sakhapov@chromium.org>\nReviewed-by: Morten Stenshorne <mstensho@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328105}",
    "author": "Daniil Sakhapov",
    "author_email": "sakhapov@chromium.org",
    "date": "2024-07-16T14:15:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/11348920d17f028275b83262a719951e838306bc"
  },
  {
    "id": 188,
    "message": "Fix css math functions min() and max()\n\nSince in C++ -0.0 == 0.0, std::min(0.0, -0.0) returns 0, which\nleads to accumulating error and results in failing of some tests.\nFix it by manually checking such situations.\n\nChange-Id: I87f919af78abd83043fc7de75a06d13616c4f996\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5707930\nReviewed-by: Morten Stenshorne <mstensho@chromium.org>\nCommit-Queue: Daniil Sakhapov <sakhapov@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328104}",
    "author": "Daniil Sakhapov",
    "author_email": "sakhapov@chromium.org",
    "date": "2024-07-16T14:14:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4eb25ce3e8c6bacd8e93e75609ec0fdbf784c4e4"
  },
  {
    "id": 189,
    "message": "Fix css math function mod()\n\nWhen the arguments have the same absolute value, but are on opposite\nsides of zero, we should act differently to put result between zero and\nB. We need to return the output of std::fmod function with sign\nflipped.\n\nSpec: https://www.w3.org/TR/css-values-4/#round-func\nFixed: 40947685\nChange-Id: I9fab0adba7569fa6d057b3fc9c138d4f43475023\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5703225\nReviewed-by: Morten Stenshorne <mstensho@chromium.org>\nCommit-Queue: Daniil Sakhapov <sakhapov@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328103}",
    "author": "Daniil Sakhapov",
    "author_email": "sakhapov@chromium.org",
    "date": "2024-07-16T14:14:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8f0d8517881da152e1324e308824b0b1e8773407"
  },
  {
    "id": 190,
    "message": "Updates to the Safety Hub Extensions Review Panel\n\nCurrently, a completion message is shown in the review panel even if the last item was removed using the button in an extension card. This CL restricts the display of the completion message to only those moments when an action is taken from within the Review Panel. The CL also\nremoves the panels \"Remove All\" button if only one extension is\ndisplayed and reworks the spacing on the completion message.\n\nBug: 350000412\nChange-Id: Iaf2c07833a8b427e315aa5793c72ed7068a7bcc8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5673232\nCommit-Queue: Adam Psarouthakis <psarouthakis@google.com>\nReviewed-by: Anunoy Ghosh <anunoy@chromium.org>\nReviewed-by: John Lee <johntlee@chromium.org>\nReviewed-by: Adam Psarouthakis <psarouthakis@google.com>\nReviewed-by: Richard Chen <richche@google.com>\nCr-Commit-Position: refs/heads/main@{#1328102}",
    "author": "Adam Psarouthakis",
    "author_email": "psarouthakis@google.com",
    "date": "2024-07-16T14:10:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7b073860993b5581331c1d897499ce28a9e488b1"
  },
  {
    "id": 191,
    "message": "[Autofill] Add experimental timeout for UploadCardRequest\n\nAs part of Loading + Confirmation, we want to enforce client-side\ntimeouts on the related Payments APIs. Currently there is no\nclient-side timeout (there is a server-side timeout, but it's\nvery high).\n\nThis CL adds a range of timeouts for upload card request, behind a\nnew feature flag (kAutofillUploadCardRequestTimeout). We will then\nexperiment with the timeouts to see which has the right balance\nof user experience vs false negatives (timing out requests that\nwould succeed).\n\nBug: 40287257\nChange-Id: Ib810104e93b7be90863650f94a48242508e451c6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5689335\nReviewed-by: Olivia Saul <jsaul@google.com>\nCommit-Queue: Stephen McGruer <smcgruer@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328101}",
    "author": "Stephen McGruer",
    "author_email": "smcgruer@chromium.org",
    "date": "2024-07-16T14:10:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8fe98b6d2570d9a003191645c01cd0fcf71413d2"
  },
  {
    "id": 192,
    "message": "[iOS] Toggle module visibility in customization menu\n\nPopulates the customization menu's UISwitches with a preference\nrepresenting the visibility of their respective modules:\n- The mediator now populates the data using a map instead of a vector so\nthat it can include the preference value for each type.\n- The mediator is set as the view controller's mutator so that it can\nhandle changes to the switches.\n- Adds an EG test which opens the menu, toggles a few switches, closes\nit and finally reopens it to check that their values were properly\nsaved in the preferences.\n\nBug: 350990359\nChange-Id: Iae81facd7c26422718c1ebcf645cf12949f8c669\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5708453\nAuto-Submit: Adam Arcaro <adamta@google.com>\nCommit-Queue: Adam Arcaro <adamta@google.com>\nReviewed-by: Chris Lu <thegreenfrog@chromium.org>\nReviewed-by: Sergio Collazos <sczs@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328100}",
    "author": "adamta",
    "author_email": "adamta@google.com",
    "date": "2024-07-16T14:09:34+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/18cedf14fb37ee46a158f7f572996dc992d97de9"
  },
  {
    "id": 193,
    "message": "Roll Perfetto Trace Processor Linux from dc7197034f80 to 2535d757a6bf\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/dc7197034f80..2535d757a6bf\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-trace-processor-linux-chromium\nPlease CC chrometto-team@google.com,perfetto-bugs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: perfetto-bugs@google.com\nChange-Id: Ic8b2bfaa1e2d2c1b050f00526d9ff3755ca7f167\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712888\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328099}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T14:09:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e7c7d5ad481e867defe1d5a8db26e8a34006e0e2"
  },
  {
    "id": 194,
    "message": "Document apply_edits.py extension of deletions\n\napply_edits.py handles edit directives where the replacement text is\nempty specially: if it detects list separators like ',' next to the\nrange to-be-deleted, it expands the range to-be-deleted to also delete\nthe list separators.\n\nFor example, if we want to delete \"*\" in \"foo(x, *y)\", the result is\n\"foo(xy)\" instead of \"foo(x, y)\".\n\nThis CL adds some documentation of this behavior and prints on\nstdout when this happens.\n\nBug: 352614160\nChange-Id: I50c580f0f259619950cfe451d72c4f9d011c1155\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5690954\nCommit-Queue: Christoph Schwering <schwering@google.com>\nReviewed-by: Daniel Cheng <dcheng@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328098}",
    "author": "Christoph Schwering",
    "author_email": "schwering@google.com",
    "date": "2024-07-16T14:04:17+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1723f5285fa2d392677d7e5cd5784ea31b6919b4"
  },
  {
    "id": 195,
    "message": "Roll Chrome Win ARM64 PGO Profile\n\nRoll Chrome Win ARM64 PGO profile from chrome-win-arm64-main-1721109360-8d427ec7c237b27b7a126f8fc66a0400c720c583-5395bd516765e60c4cdbdf9f7758522d2f88a382.profdata to chrome-win-arm64-main-1721131103-29c068c3c7a6375f2149210066d8ffe6705c5444-f1ce5ea1e387d15eb36e31879f11a08cd241e09c.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: Id279cb3adb11ad1695290fad0581055708968b8b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712791\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328097}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T13:55:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/afcddd74c286e85643f18cf42d7ca3f067e96763"
  },
  {
    "id": 196,
    "message": "Roll Arm AFDO from 128.0.6597.0_rc-r1-merged to 128.0.6598.0_rc-r1-merged\n\nThis CL may cause a small binary size increase, roughly proportional\nto how long it's been since our last AFDO profile roll. For larger\nincreases (around or exceeding 100KB), please file go/crostc-bug.\n\nPlease note that, despite rolling to chrome/android, this profile is\nused for both Linux and Android.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/afdo-arm-chromium\nPlease CC c-compiler-chrome@google.com,mobiletc-prebuild@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium Main: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: mobiletc-prebuild@google.com\nChange-Id: Ib46e83700bd683fda6561880cbbcb930a2c773dd\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713089\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328096}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T13:53:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d3db3b770e8bf909f5bc6c7711bf0980c61d523a"
  },
  {
    "id": 197,
    "message": "Change prediction_source to experimental in capture site tests\n\nIt doesn't make a difference [1], but is wrong regardless, since the\nnextgen file doesn't exist anymore.\n\n[1] https://source.chromium.org/chromium/chromium/src/+/main:components/autofill/core/browser/heuristic_source.cc;l=21-23;drc=99d1b90e58c0aeaf0693f0652bfe1569c8663f95\n\nBug: 40280853\nChange-Id: I81505b033a6f6db75378d58d09e28363d1ad74f3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712728\nAuto-Submit: Florian Leimgruber <fleimgruber@google.com>\nCommit-Queue: Christoph Schwering <schwering@google.com>\nReviewed-by: Christoph Schwering <schwering@google.com>\nCr-Commit-Position: refs/heads/main@{#1328095}",
    "author": "Florian Leimgruber",
    "author_email": "fleimgruber@google.com",
    "date": "2024-07-16T13:52:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/46144b04c6300d0b9619bd40b4108d5ac3ad89a8"
  },
  {
    "id": 198,
    "message": "Roll ios_internal from e058eccb6d9c to 53d0a975b482\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/e058eccb6d9c..53d0a975b482\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,danieltwhite@google.com,thegreenfrog@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: chromium:40935924\nChange-Id: I17095ef4054dc5a6659ec518b565a076b8344ae3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712610\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328094}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-16T13:48:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/82c74dc97f2bd9406bcd151cff7e45124cfe69a1"
  },
  {
    "id": 199,
    "message": "Roll ANGLE from 9e60edc7ca54 to b8d21cd23034 (1 revision)\n\nhttps://chromium.googlesource.com/angle/angle.git/+log/9e60edc7ca54..b8d21cd23034\n\n2024-07-16 angle-autoroll@skia-public.iam.gserviceaccount.com Roll SwiftShader from f23c77132e0b to 389854967d78 (1 revision)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/angle-chromium-autoroll\nPlease CC angle-team@google.com,yuxinhu@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in ANGLE: https://bugs.chromium.org/p/angleproject/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86\nBug: None\nTbr: yuxinhu@google.com\nChange-Id: I8427833bada41332804542928f7cdef0a6518340\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711802\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328093}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T13:46:16+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/56a0fe3bdf00434d4114ce39ccc9f39d73fe41a9"
  },
  {
    "id": 200,
    "message": "Certificate Management UI V2: Move strings to grdp file\n\n* Move more strings into the grdp file, so that they can be translated\n  (once all the screenshots are taken).\n\n* Add a few missing headers.\n\nBug: 40928765\nChange-Id: I458b2bb106b6a5ccedf6e549460a342ac8f6c9a8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5707533\nCommit-Queue: Hubert Chao <hchao@chromium.org>\nReviewed-by: Matt Mueller <mattm@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328092}",
    "author": "Hubert Chao",
    "author_email": "hchao@chromium.org",
    "date": "2024-07-16T13:43:15+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d3540b90427e54fd4449c7366e4d38495973ff1f"
  },
  {
    "id": 201,
    "message": "[iOS] Password \"Delete\" button still in place when user interacts with multiple rows\n\nThe 'setEditing' method has some sub-calls that employ batch update, a method that calls `setEditing` with `editing = NO`, which disables editing on all table view's cells, causing the abrupt disappearance of the 'delete' buttons. Therefore, the sub-calls should only take place when there is a change in the editing state (i.e. no editing -> editing and editing -> no editing).\n\nFixed: 330404655\nChange-Id: Ic007f392fad8bd1413e8d17d0f614348e6d344dc\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5694708\nReviewed-by: Viktor Semeniuk <vsemeniuk@google.com>\nCommit-Queue: Cristian Ciacu <cristianciacu@google.com>\nReviewed-by: Filipa Senra <fsenra@google.com>\nCr-Commit-Position: refs/heads/main@{#1328091}",
    "author": "Cristian Ciacu",
    "author_email": "cristianciacu@google.com",
    "date": "2024-07-16T13:42:47+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5ef019240764c86e6f63831feda0c62591cd0016"
  },
  {
    "id": 202,
    "message": "Roll Chrome Mac Arm PGO Profile\n\nRoll Chrome Mac Arm PGO profile from chrome-mac-arm-main-1721109360-76a9e576a1ce24e838a3df282f59e517dc2ccf10-5395bd516765e60c4cdbdf9f7758522d2f88a382.profdata to chrome-mac-arm-main-1721131103-ae09384c2abd0a2fa096ecef424bee1b5431660c-f1ce5ea1e387d15eb36e31879f11a08cd241e09c.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-mac-arm-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:mac-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I502b2c55583d3532c9ab97b19d908e859f3b5d3c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713088\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328090}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T13:42:05+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/33dbe25614277cd177f0cbca7b7a2360b8a4296d"
  },
  {
    "id": 203,
    "message": "[Autofill] Remove feature-name exception from presubmit check\n\nThe feature AutofillAddressEnhancementVotes does not exist\nanymore.\n\nBug: 40100455\nChange-Id: Ib3baab197be1dc7b10d85f4bcaa401d43a73dbf3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5703904\nReviewed-by: Matthias Körber <koerber@google.com>\nAuto-Submit: Christoph Schwering <schwering@google.com>\nCommit-Queue: Christoph Schwering <schwering@google.com>\nCr-Commit-Position: refs/heads/main@{#1328089}",
    "author": "Christoph Schwering",
    "author_email": "schwering@google.com",
    "date": "2024-07-16T13:39:52+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/038b21830737b3a942329133bc6614cd40248bce"
  },
  {
    "id": 204,
    "message": "[Profiles][Windows] Change bagded profile avatar icon\n\nBehind the feature flag \"OutlineSilhouetteIcon\", change the profile\navatar icon displayed next to the browser icon on Windows to be the\nnew outline silhouette icon.\n\nThe new icon is set once the profile is opened with the feature flag\nenabled (https://screencast.googleplex.com/cast/NjQ2MTE1Mzk2MTA1MDExMnw3OGEyMjdkMS05YQ).\n\nScreenshots:\n- Icon badge on shortcut: https://screenshot.googleplex.com/7gW2rQLCw7REnhP\n- Icon badge in taskbar & window switcher: https://screenshot.googleplex.com/9ZptW6Vo5bx8Cm5\n\nBug: b:307739359\nRequested changes: https://docs.google.com/presentation/d/1fVZMbZ5zpkpAcaSn31gTaX5F-_mOZRjK_azWMDUBXqE/edit#slide=id.g29410333a18_0_31\n\nChange-Id: Iaa3bb4744d73205f86dff861c844b7ba570d4126\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5678767\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Amelie Schneider <amelies@google.com>\nReviewed-by: Ryan Sultanem <rsult@google.com>\nReviewed-by: Alex Ilin <alexilin@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328088}",
    "author": "Amelie Schneider",
    "author_email": "amelies@google.com",
    "date": "2024-07-16T13:39:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/14ffda24fec97d11fbc38133d61a939b3cecb1c3"
  },
  {
    "id": 205,
    "message": "[Frameworks roll] Roll to 652719529 piper revision\n\nChange-Id: Iadf387a45a6b47b05ac3b1bec3174585e795ca3b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5713068\nCommit-Queue: Internal Frameworks Autoroller <bling-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com>\nBot-Commit: Internal Frameworks Autoroller <bling-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328087}",
    "author": "Internal Frameworks Autoroller",
    "author_email": "bling-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com",
    "date": "2024-07-16T13:36:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8b28f5f40f31c04ff73565cf03323d6adcb149d5"
  },
  {
    "id": 206,
    "message": "Roll V8 from db8922240fa9 to 0ca6756a98d2 (20 revisions)\n\nhttps://chromium.googlesource.com/v8/v8.git/+log/db8922240fa9..0ca6756a98d2\n\n2024-07-16 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Version 12.8.327\n2024-07-16 evih@chromium.org Revert \"Reland^2 \"[builtins] Optimize and generate builtins off-main-thread\"\"\n2024-07-16 dmercadier@chromium.org [turboshaft] Maglev-to-ts: enable Maglev escape analysis\n2024-07-16 mliedtke@chromium.org [wasm][turboshaft] Simplify rtt depth check for ref.cast and ref.test\n2024-07-16 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Update V8 DEPS (trusted)\n2024-07-15 jkummerow@chromium.org [wasm][sandbox] Harden WasmImportData against concurrent modification\n2024-07-15 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Update V8 DEPS (trusted)\n2024-07-15 victorgomes@chromium.org [maglev] Iterate the phi list when reducing predecessor count\n2024-07-15 verwaest@chromium.org [parser] reenable scope info reuse\n2024-07-15 snek@chromium.org [api] add Get/Set with receiver\n2024-07-15 clemensb@chromium.org [compiler] Load 32-bit constants as 32-bit\n2024-07-15 dmercadier@chromium.org [maglev] Extract post-hoc processors from maglev-compiler.cc\n2024-07-15 junyan@redhat.com s390x: [maglev] Support extending property backing stores\n2024-07-15 verwaest@chromium.org [streaming] Use UniqueIdInScript instead of StartPosition\n2024-07-15 junyan@redhat.com [maglev] Fix endianess issue on GapMove\n2024-07-15 dmercadier@chromium.org [wasm][arm64] Only store 32 bits in thread_in_wasm_flag instead of 64\n2024-07-15 verwaest@chromium.org [parser] Fix eval handling in scope info reuse\n2024-07-15 yahan@iscas.ac.cn [riscv] In Abort, Use Call directly to avoid any unneeded overhead\n2024-07-15 syg@chromium.org Reland^2 \"[builtins] Optimize and generate builtins off-main-thread\"\n2024-07-15 marja@chromium.org [compile hints magic] Fix a bug where the compile hints magic comment value contains two-byte data\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/v8-chromium-autoroll\nPlease CC liviurau@google.com,machenbach@google.com,v8-waterfall-gardener@grotations.appspotmail.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-blink-rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:dawn-linux-x64-deps-rel\nChange-Id: I1c18474cd4996667cda5fdd47dea253b62a1c24b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712548\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328086}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T13:31:52+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c669d62895219ff2621d5756e0e03a58b5683b60"
  },
  {
    "id": 207,
    "message": "Roll clank/internal/apps from 74887d9acac3 to 973f3ee91312 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/74887d9acac3..973f3ee91312\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,izuzic@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: izuzic@google.com\nNo-Try: true\nChange-Id: I4d03d59f0edcef26d5f64c00d13b31ff6284d583\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712612\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328085}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-16T13:30:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8c842cf1eb2801133297492d3432f53dff87f13d"
  },
  {
    "id": 208,
    "message": "Update `TestExpectations` with bugs filed for crrev.com/c/5709969\n\nR=rubber-stamper@appspot.gserviceaccount.com\n\nBug: 353228154, 353228155\nChange-Id: Ic91c67126a69e72c22219ad7a90f6e322c737412\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712748\nAuto-Submit: WPT Autoroller <wpt-autoroller@chops-service-accounts.iam.gserviceaccount.com>\nBot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCommit-Queue: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328084}",
    "author": "Chromium WPT Sync",
    "author_email": "wpt-autoroller@chops-service-accounts.iam.gserviceaccount.com",
    "date": "2024-07-16T13:26:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ab32509539db12f9ceb66af7b7c4b1ae61579cff"
  },
  {
    "id": 209,
    "message": "Adopt base::test::RunOnceCallback[Repeatedly] in FCMHandlerTest\n\nTest-only change. More concise than a custom action.\n\nChange-Id: I2bddc9148deb1a33b2e2913714a8ec400d4a061d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712050\nAuto-Submit: Victor Vianna <victorvianna@google.com>\nReviewed-by: Rushan Suleymanov <rushans@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Rushan Suleymanov <rushans@google.com>\nCr-Commit-Position: refs/heads/main@{#1328083}",
    "author": "Victor Hugo Vianna Silva",
    "author_email": "victorvianna@google.com",
    "date": "2024-07-16T13:26:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/51c2b65b2b01129a604b616760d9ce17e2982942"
  },
  {
    "id": 210,
    "message": "Add options to specify supervised user family members in test.\n\nChange-Id: I95084c2c4a22c8363b809cff7c620d4189f4bfa3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5705298\nReviewed-by: Anthi Orfanou <anthie@google.com>\nCommit-Queue: Tomek Jurkiewicz <tju@google.com>\nCr-Commit-Position: refs/heads/main@{#1328082}",
    "author": "Tomasz Jurkiewicz",
    "author_email": "tju@google.com",
    "date": "2024-07-16T13:24:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6e0a3db4d0344d1619740ab7dac6958cd6b1d7b9"
  },
  {
    "id": 211,
    "message": "kiosk: Comment KioskAppsMixin and rename constant to kTestChromeAppId\n\nThis also fixes includes in kiosk_crash_restore_browsertest.cc.\n\nBug: none\nChange-Id: I7f5787f4a8651c6de386e091c415bdc2ffd24f0e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712049\nReviewed-by: Maksim Ivanov <emaxx@chromium.org>\nReviewed-by: Irina Fedorova <irfedorova@google.com>\nCommit-Queue: Edman Anjos <edman@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328081}",
    "author": "Edman Anjos",
    "author_email": "edmanp@google.com",
    "date": "2024-07-16T13:22:47+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/86b1c9d2079c950e09cbf88b52e5143002f92306"
  },
  {
    "id": 212,
    "message": "Roll Perfetto from dc7197034f80 to 2535d757a6bf (1 revision)\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/dc7197034f80..2535d757a6bf\n\n2024-07-16 android-test-infra-autosubmit@system.gserviceaccount.com Merge \"ui: implement \"Show From Frame\" in flamegraph\" into main\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-chromium-autoroll\nPlease CC perfetto-bugs@google.com,primiano@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-perfetto-rel\nBug: None\nTbr: perfetto-bugs@google.com\nChange-Id: Ie3050f3de26706a71edfb1c69ac98e36093f6434\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711803\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328080}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T13:13:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/aba8d5f484b81cd2696aa65b49af593bc8712202"
  },
  {
    "id": 213,
    "message": "[ios] add confirmation when deleting a tab group\n\nThis CL adds a confirmation when deleting a tab group.\n\nThe confirmation appears\n- in the tab grid view\n- in the tab group view (not in this CL)\nAlso the confirmation appears when the user is signed out.\n\nThe confirmation and the delete group action in incognito doesn't show\nup. (implemented in https://chromium-review.googlesource.com/c/chromium/src/+/5701960)\n\nOn iPhone, the confirmation buttons are placed at the bottom of the tab\ngrid.\nOn iPad, the confirmation popover points to the item that was long\npressed\n\ndemo on iPhone: https://drive.google.com/file/d/1Es72IjrjocyitvYn206t3gBZ2wbWmIZU/view?usp=sharing\ndemo on iPad: https://drive.google.com/file/d/1ZzZGbq1vefaM27sqC8nSKEwATpLul3Tv/view?usp=sharing\n\nBug: 329627336\nChange-Id: Iabebc4e26529649080c60829ac2b5d8a204e14b5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5683003\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Gauthier Ambard <gambard@chromium.org>\nCommit-Queue: Asami Doi <asamidoi@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328079}",
    "author": "Asami Doi",
    "author_email": "asamidoi@chromium.org",
    "date": "2024-07-16T13:12:36+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/21734de5c94ab337a5a13d9fe756cee39d8676b7"
  },
  {
    "id": 214,
    "message": "Roll vulkan-deps from 3fefe147dce7 to 99f840c7c7ad (3 revisions)\n\nhttps://chromium.googlesource.com/vulkan-deps.git/+log/3fefe147dce7..99f840c7c7ad\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/external/github.com/KhronosGroup/Vulkan-Headers/+log/fc6c06ac529e4b4b6e34c17cc650a8f62dee2eb0..b379292b2ab6df5771ba9870d53cf8b2c9295daf\n  https://chromium.googlesource.com/external/github.com/KhronosGroup/Vulkan-ValidationLayers/+log/86e776e02529060dc0bd5bbc566239b8a0d852fb..98d78ad635340c289286e5494ce3fc388e840b55\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/vulkan-deps-chromium-autoroll\nPlease CC angle-team@google.com,cwallez@google.com,radial-bots+chrome-roll@google.com,radial-bots@google.com,webgpu-developers@google.com,yuxinhu@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86;luci.chromium.try:dawn-linux-x64-deps-rel;luci.chromium.try:win-asan;luci.chromium.try:linux_chromium_cfi_rel_ng\nBug: None\nTbr: cwallez@google.com,radial-bots+chrome-roll@google.com,yuxinhu@google.com\nChange-Id: Ibae6ec845f2cf3331ca48757ce0b8b85f96ed095\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711774\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328078}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T13:10:36+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/018d40506e8afb66d661e03074add206a91c0e3b"
  },
  {
    "id": 215,
    "message": "Add enums for abandoned navigation metrics\n\nCurrently we only log the timestamps between the last milestone to\nabandonment, which is good for analyzing the latency details but a\nbit hard to get an overall view of what abandonment reasons and\nmilestones are most common. This CL adds an enum to make it easier\nto get that, and refactors the histogram name creation as well,\nto make the code less bloated (and to prepare for a future CL to\ngeneralize the abandon observer).\n\nOBSOLETE_HISTOGRAMS=NewNavigation histograms has been updated\nto use the more detailed names.\n\nBug: 347706997\nChange-Id: I4c7cf0739d767cd26d90b9ef408ab1a285a8cfa0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5693886\nReviewed-by: Minoru Chikamune <chikamune@chromium.org>\nCommit-Queue: Rakina Zata Amni <rakina@chromium.org>\nReviewed-by: Shunya Shishido <sisidovski@chromium.org>\nReviewed-by: Takashi Toyoshima <toyoshim@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328077}",
    "author": "Rakina Zata Amni",
    "author_email": "rakina@chromium.org",
    "date": "2024-07-16T13:06:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5afdf28f670b703609dc8b580cf0c813ba80696a"
  },
  {
    "id": 216,
    "message": "Remove V8 slow histograms from field trial default\n\nChromium perf bots pick up the first experiment, to which\nhttps://crrev.com/c/5706063 added the slow histograms, which then caused\na bunch of regression alerts. This CL adds a third arm now without the\nslow histograms and moves that to first place.\n\nBug: 42204161, 353178386, 353182422, 353182421\nChange-Id: Ieeb77d11935286e24f1a9dc0108d7fe7a8feb476\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712448\nReviewed-by: Clemens Backes <clemensb@chromium.org>\nCommit-Queue: Daniel Lehmann <dlehmann@chromium.org>\nReviewed-by: Camillo Bruni <cbruni@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328076}",
    "author": "Daniel Lehmann",
    "author_email": "dlehmann@chromium.org",
    "date": "2024-07-16T13:02:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f619d9a4e133c52f6e2036ce3dc662765e549525"
  },
  {
    "id": 217,
    "message": "Roll ChromeOS Bigcore AFDO profile from 128-6582.0-1721039692-benchmark-128.0.6597.0-r1 to 128-6582.0-1721039692-benchmark-128.0.6598.0-r1\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/cros-afdo-bigcore-chromium\nPlease CC c-compiler-chrome@google.com,mobiletc-prebuild@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium Main: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: mobiletc-prebuild@google.com\nChange-Id: Ic7c675233ace988bc7a76ba0d85cc0070fec4655\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712609\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328075}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T12:58:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f1ce5ea1e387d15eb36e31879f11a08cd241e09c"
  },
  {
    "id": 218,
    "message": "[sync] Add scaffold of new batch upload internals\n\nNo behavior change in this CL, just a new interface that isn't called\nyet and empty implementations.\n\nThis CL:\n- Introduces a new interface ModelTypeLocalDataBatchUploader, with\n  one method for previewing the local data to migrate, and another\n  one for actually performing the migration.\n- Adds trivial implementations of the interface for BOOKMARKS,\n  READING_LIST and PASSWORDS. The actual implementations will come\n  in next CLs.\n- Adds a mock implementation for tests.\n\nBug: 40072432\nChange-Id: I1f46bb4b1192838693a22cd4730e64612a06cf1d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709609\nCommit-Queue: Mikel Astiz <mastiz@chromium.org>\nReviewed-by: Mikel Astiz <mastiz@chromium.org>\nAuto-Submit: Victor Vianna <victorvianna@google.com>\nCommit-Queue: Victor Vianna <victorvianna@google.com>\nCr-Commit-Position: refs/heads/main@{#1328074}",
    "author": "Victor Hugo Vianna Silva",
    "author_email": "victorvianna@google.com",
    "date": "2024-07-16T12:49:05+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/49a57f5c52950cfa1fa257aae03210045b3915f5"
  },
  {
    "id": 220,
    "message": "Refactor PlusAddressService::SupportsPlusAddresses.\n\nThis CL contains no functionality change - it simply refactors the\ncurrent SupportsPlusAddresses method by\n- renaming it to ShouldShowManualFallback and continuing to use it in\n  the AutofillContextMenuManager,\n- inlining its content inside the only other call-site,\n  PlusAddressService::GetAvailableSuggestions.\n\nThis is a preparation for\na) fixing a bug that does not take affiliations into account when\n   showing suggestions in incognito mode,\nb) including the state of the global account toggle\n\nBug: 353241201, 353240084\nChange-Id: I15f980c66e68330ca88c5214d5f124af6fc27354\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5708010\nReviewed-by: Norge Vizcay <vizcay@google.com>\nAuto-Submit: Jan Keitel <jkeitel@google.com>\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nCr-Commit-Position: refs/heads/main@{#1328072}",
    "author": "Jan Keitel",
    "author_email": "jkeitel@google.com",
    "date": "2024-07-16T12:38:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9f7f1ab79ee7eab570b365077a4a6ab1ff0285dd"
  },
  {
    "id": 221,
    "message": "[Gardener] Disable flaky tab switcher test\n\ntestTabGroupOverflowMenuInTabSwitcher_deleteGroupDecline has the average\nflaky rate in runs of 24.24%.\n\nBug: 353463207\nChange-Id: I6e37f9860ae84f413c31ea36247276f48eaba7eb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712468\nAuto-Submit: Ivana Žužić <izuzic@google.com>\nOwners-Override: Ivana Žužić <izuzic@google.com>\nCommit-Queue: Ivana Žužić <izuzic@google.com>\nReviewed-by: Theo Cristea <theocristea@google.com>\nCommit-Queue: Theo Cristea <theocristea@google.com>\nCr-Commit-Position: refs/heads/main@{#1328071}",
    "author": "Ivana Žužić",
    "author_email": "izuzic@google.com",
    "date": "2024-07-16T12:35:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c093387ff9a360ddf86c8d5e2aa68a9c59ef660e"
  },
  {
    "id": 222,
    "message": "[sync] Add comment in ReadingListSpecifics about time since unix epoch\n\nIt makes sense to highlight this since BookmarksSpecifics has the same\nfields measured from windows epoch.\n\nBug: 353449026\nChange-Id: Ic404764b54a708025d68972793312864cdbf2c5b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5710909\nCommit-Queue: Ankush Singh <ankushkush@google.com>\nReviewed-by: Marc Treib <treib@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328070}",
    "author": "Ankush Singh",
    "author_email": "ankushkush@google.com",
    "date": "2024-07-16T12:27:43+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/47a5f5d62165fdc748f390756b2f2c54f80cdcdb"
  },
  {
    "id": 223,
    "message": "Fix freshness score to be float.\n\nAll inputs need to be floats in segmentation models.\nDisable auto execute when model is ondemand.\n\nBug: 337878177\nChange-Id: Ia61edec13010e96856c2ab45664be3bd1fbdafad\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5710446\nCommit-Queue: Xi Han <hanxi@chromium.org>\nAuto-Submit: Siddhartha S <ssid@chromium.org>\nReviewed-by: Xi Han <hanxi@chromium.org>\nReviewed-by: Ritika Gupta <ritikagup@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328069}",
    "author": "ssid",
    "author_email": "ssid@chromium.org",
    "date": "2024-07-16T12:24:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8a3c6cee7c38d9e5a6951ee25650a041bab96e4c"
  },
  {
    "id": 224,
    "message": "spanification: //third_party/blink/renderer/core/layout/svg/\n\n svg_text_layout_algorithm.cc and\n resolved_text_layout_attributes_iterator.h\n\nBug: 351564777\nChange-Id: I3b25aac79a94cdbce2d594cd429f76a56c72ffed\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5694825\nCommit-Queue: Fredrik Söderquist <fs@opera.com>\nReviewed-by: Kent Tamura <tkent@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328068}",
    "author": "Fredrik Söderqvist",
    "author_email": "fs@opera.com",
    "date": "2024-07-16T12:23:43+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a6459c2a88a0f02972bb0c94e48c69a264953566"
  },
  {
    "id": 225,
    "message": "[ios] Make sure the removed synthetic form ids are sorted\n\nBefore running base::ranges::includes on the synthetic form fields\n(extracted from cached forms) the renderer ids have to be sorted,\nthe algorithm will fail otherwise.\n\nThis fixes crashes we see on Canary regarding that.\n\nChange-Id: I3d638dc47e9171f413c3fb32ae36aa696782bd3c\nBug: 352532310\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5709871\nCommit-Queue: Vincent Boisselle <vincb@google.com>\nReviewed-by: Noémie St-Onge <noemies@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328067}",
    "author": "Vincent Boisselle",
    "author_email": "vincb@google.com",
    "date": "2024-07-16T12:23:09+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2e0584eb342ed6a1aa7396454c2cc18c0d354d7e"
  },
  {
    "id": 226,
    "message": "[Start] Clean up cached flags in SyncPromoController.\n\nIn this CL, we remove cached flags of StartSurfaceConfiguration and\nuse static variables in SyncPromoControler to replcae their values:\n- SIGNIN_PROMO_NTP_COUNT_LIMIT\n- SIGNIN_PROMO_NTP_SINCE_FIRST_TIME_SHOWN_LIMIT_HOURS\n- SIGNIN_PROMO_NTP_RESET_AFTER_HOURS\nAlso clean up another two unused cached flags.\n\nBug: 344651414\nChange-Id: Ic4d27e9f7c01297330d6863a3d2ded0f15148d3d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5710231\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Tanmoy Mollik <triploblastic@google.com>\nCommit-Queue: Xi Han <hanxi@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1328066}",
    "author": "Xi Han",
    "author_email": "hanxi@google.com",
    "date": "2024-07-16T12:19:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/80dc2cbd77bdfbdb97f8bd7488abcc4c8a859e3d"
  },
  {
    "id": 227,
    "message": "[Gardener] Disable flaky StatusBarColorControllerTest\n\nThe test is 14.85% flaky in runs.\n\nBug: 353460498\nChange-Id: I7cba032045b59d7809f8463ff032ec3c739429c7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5712368\nCommit-Queue: Theo Cristea <theocristea@google.com>\nOwners-Override: Ivana Žužić <izuzic@google.com>\nReviewed-by: Theo Cristea <theocristea@google.com>\nAuto-Submit: Ivana Žužić <izuzic@google.com>\nCr-Commit-Position: refs/heads/main@{#1328065}",
    "author": "Ivana Žužić",
    "author_email": "izuzic@google.com",
    "date": "2024-07-16T12:15:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/66e5966d00f480c416189c33fff5f0829e29b42e"
  },
  {
    "id": 228,
    "message": "Roll devtools-internal from 7a3ec8eaef63 to 07a9a068d30b (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/7a3ec8eaef63..07a9a068d30b\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/devtools/devtools-frontend.git/+log/d4c016959f5026bf795ea853d5363861d97749a9..37848592ea7180f71e330a7a54a273e98a9f52b9\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: I06ddd82ba4851bc5bbadbfee54eeb134ed6bc557\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711934\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328064}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-16T12:14:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/aef4c4d0a16b1fcefbe6fe5e98257be384354f13"
  },
  {
    "id": 229,
    "message": "[UNOp2][Android] Use the new sign-in flow on Automotive\n\nThis CL enables the bottom-sheet based new sign-in flow for Android\nAutomotive. The upgrade promo will remain suppressed for Automotive.\n\nSee https://crbug.com/352394844#comment3 for recording.\n\nFixed: 352394844\nChange-Id: I3901bcf9b35c4e05e7211578e726516f6e924ce8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5695645\nReviewed-by: Samar Chehade-Lepleux <samarchehade@google.com>\nCommit-Queue: Menghan Yang <myuu@google.com>\nCr-Commit-Position: refs/heads/main@{#1328063}",
    "author": "Menghan Yang",
    "author_email": "myuu@google.com",
    "date": "2024-07-16T12:14:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d55fec704396f09f61e7792b381f3ee3695446c3"
  },
  {
    "id": 230,
    "message": "Roll Skia from 3a0067d2f105 to 07bd32512db2 (1 revision)\n\nhttps://skia.googlesource.com/skia.git/+log/3a0067d2f105..07bd32512db2\n\n2024-07-16 skia-autoroll@skia-public.iam.gserviceaccount.com Roll vulkan-deps from 6ab83f393122 to 99f840c7c7ad (6 revisions)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/skia-autoroll\nPlease CC nicolettep@google.com,skiabot@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Skia: https://bugs.chromium.org/p/skia/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux-blink-rel;luci.chromium.try:linux-chromeos-compile-dbg;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nCq-Do-Not-Cancel-Tryjobs: true\nBug: None\nTbr: nicolettep@google.com\nChange-Id: If0827e7e403cfed4a011f5b8e968abbcea73589b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711784\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328062}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-16T12:12:31+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3dc87b7dd9a2e02696113d97591db030f775a3be"
  },
  {
    "id": 231,
    "message": "Roll ios_internal from 7e20d5405b84 to e058eccb6d9c\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/7e20d5405b84..e058eccb6d9c\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,danieltwhite@google.com,thegreenfrog@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: Iec662493cd72cb793855d80e8d7a1427a07a9b3e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5711956\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1328061}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-16T12:12:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/abc57903ca6d7d08cc840792b5d7cb6ee615ed95"
  },
  {
    "id": 285,
    "message": "Roll Chrome Linux PGO Profile\n\nRoll Chrome Linux PGO profile from chrome-linux-main-1720655837-797818dbb03d1b09d8b3b2b781a7b0c5a3f0d5e8-bb52a88220825ae3b01e1149182da0bbe22ef78c.profdata to chrome-linux-main-1720677402-be59c8e8cedee7074baa5569ed7e348af7010b3d-904d57e857b393029c0c549e9378b85f4adc0afb.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-linux-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I911dac70c5e9d9ce172f049efff9b7740cd2b8f4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5693157\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325971}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T09:14:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9d92e92235b955740553ae7340cba16c1a5417cd"
  },
  {
    "id": 286,
    "message": "[Reauth] Fix memory leak in the AutofillPaymentMethodsFragment\n\nAutofillPaymentMethodsFragment creates the ReauthenticatorBridge, which\ncreates its C++ counterpart and never destroys it causing a memory leak.\nReauthenticatorBridge.destroy() needs to be called when fragment gets\ndestroyed, otherwise the C++ ReauthenticatorBridge and all its\nreferences never get garbage collected.\n\nBug: 351890412\nChange-Id: I4c269c68dc31ccd028297cf98c8a12b82e1f2438\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5690032\nReviewed-by: Florian Leimgruber <fleimgruber@google.com>\nCommit-Queue: Anna Tsvirchkova <atsvirchkova@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325970}",
    "author": "Anna Tsvirchkova",
    "author_email": "atsvirchkova@google.com",
    "date": "2024-07-11T09:10:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8c6f83c4cfa205674d68fc454621d8dfbc875e93"
  },
  {
    "id": 287,
    "message": "[Reauth] Fix memory leak in the ReauthenticatorBridge\n\nReauthenticatorBridge's C++ counterpart currently never gets destroyed and causes a memory leak in case when ChromeTabbedActivity gets re-created.\n\nBug: 351890412\nChange-Id: Ia22c36af3657cde37eb6088582f53a60fc40dc7c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5689331\nReviewed-by: Vasilii Sukhanov <vasilii@chromium.org>\nReviewed-by: Christian Dullweber <dullweber@chromium.org>\nReviewed-by: Sky Malice <skym@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Anna Tsvirchkova <atsvirchkova@google.com>\nCr-Commit-Position: refs/heads/main@{#1325969}",
    "author": "Anna Tsvirchkova",
    "author_email": "atsvirchkova@google.com",
    "date": "2024-07-11T09:10:34+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2994af5c42e54a52b72a8e2f7ff55c9f306ea9ba"
  },
  {
    "id": 288,
    "message": "Roll V8 from a3976fec3e52 to cc08e720be5d (2 revisions)\n\nhttps://chromium.googlesource.com/v8/v8.git/+log/a3976fec3e52..cc08e720be5d\n\n2024-07-11 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Version 12.8.281\n2024-07-11 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Update V8 DEPS (trusted)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/v8-chromium-autoroll\nPlease CC liviurau@google.com,machenbach@google.com,v8-waterfall-gardener@grotations.appspotmail.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-blink-rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:dawn-linux-x64-deps-rel\nChange-Id: I68fc2e288667eca84bca5d162eab2fbb69e8411a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5693830\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325968}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T09:08:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c1f6bbc07e2a621315416e6e98d0a3b94307f841"
  },
  {
    "id": 289,
    "message": "[iOS][STG] Implements close group locally\n\nWhen a group is closed, it shouldn't be deleted  from the sync service.\n\nI'm planing to restrict this action (incognito / signed out) in a\nfollow-up.\n\nBug: 329627016, 352297050\nChange-Id: Ic37eac605ee644e6b2e829a7f24d3eee6d5ef72f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5687589\nReviewed-by: Gauthier Ambard <gambard@chromium.org>\nCommit-Queue: Ewann Pellé <ewannpv@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325967}",
    "author": "Ewann Pelle",
    "author_email": "ewannpv@chromium.org",
    "date": "2024-07-11T09:05:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/217cfc583c827e3b76263679fec3e718f9a787fe"
  },
  {
    "id": 290,
    "message": "[ios] Migrate tests away from LegacyBookmarkModel\n\nTest-only change, no behavioral changes.\n\nProduction (non-test) code is already migrated away from\nLegacyBookmarkModel and interacts directly with\nbookmarks::BookmarkModel.\n\nIn this patch, the same is done for tests, migrating all remaining\noccurrences. Most of the changes are about Earl Grey tests, but\nthere are a few trivial changes in unit-tests too.\n\nBug: 346918509\nChange-Id: Ib9f99a2b7f03f3aa728d300428a493df8afe614c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5688510\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Jérôme Lebel <jlebel@chromium.org>\nReviewed-by: Gauthier Ambard <gambard@chromium.org>\nCommit-Queue: Mikel Astiz <mastiz@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325966}",
    "author": "Mikel Astiz",
    "author_email": "mastiz@chromium.org",
    "date": "2024-07-11T09:05:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/874edf178cf8d99bd715ed606a86691a6d4bf10d"
  },
  {
    "id": 291,
    "message": "Add counters to 'delete all password manager data' dialog\n\nThis CL also includes updated color of the warning text and icon.\n\nImpl: http://screen/6Rihrmxr4DaxhAc.png\nMock: http://screen/9rEQCtbz4Q3k2fU.png\n\nBug: 342366264\nChange-Id: I516b123045e0353a8b17a545a157623aa41ee513\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5684980\nAuto-Submit: Andrii Natiahlyi <natiahlyi@google.com>\nReviewed-by: Adem Derinel <derinel@google.com>\nReviewed-by: Mohamed Amir Yosef <mamir@chromium.org>\nCommit-Queue: Adem Derinel <derinel@google.com>\nCr-Commit-Position: refs/heads/main@{#1325965}",
    "author": "Andrii Natiahlyi",
    "author_email": "natiahlyi@google.com",
    "date": "2024-07-11T08:57:12+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d09b6897f741267f6b5e3b977c29971372da79cd"
  },
  {
    "id": 292,
    "message": "Roll PDFium from a715706df8a1 to a08da34685b1 (3 revisions)\n\nhttps://pdfium.googlesource.com/pdfium.git/+log/a715706df8a1..a08da34685b1\n\n2024-07-11 rhalavati@google.com Update FPDFImageObj_GetRenderedBitmap matrix processing.\n2024-07-09 idtl@google.com Avoid float-cast-overflow in IccTransform::Translate\n2024-07-09 rop@google.com Updating FP16's README.chromium to include Revision\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pdfium-autoroll\nPlease CC andyphan@google.com,dhoss@chromium.org,thestig@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in PDFium: https://bugs.chromium.org/p/pdfium/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: chromium:350851535,chromium:42271013\nTbr: andyphan@google.com\nChange-Id: Iff0311a07dc8d7b779404dc3eee9c027e34c01f1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5693156\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325964}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T08:56:08+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b6c4ff68da6158696a7468937762483a8177b07e"
  },
  {
    "id": 293,
    "message": "Fix VO annoucement for TR1.5\n\nThere is no sessionLabel for TR1.5.\nThe session is part of the hostnameAndSyncTimeLabel.\n\nFixed: 352057933\nChange-Id: Ie4e3f19fe2ef067e4f1e0e133afa9e7a1578311e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5691357\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Gauthier Ambard <gambard@chromium.org>\nAuto-Submit: Olivier Robin <olivierrobin@chromium.org>\nCommit-Queue: Gauthier Ambard <gambard@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325963}",
    "author": "Olivier Robin",
    "author_email": "olivierrobin@google.com",
    "date": "2024-07-11T08:52:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3e65472df63fa01e13af93640e8251dbf546d640"
  },
  {
    "id": 294,
    "message": "HaTS: Prioritized camera HaTS for specific Chromebook models\n\nThis patch will read list of models from finch and decide whether the\ncurrent user will get a prioritized HaTS or not.\nThe difference between prioritized/not is just the probability of being\nselected, while also utilizing \"prioritized\" type of HaTS.\n\nMore details: go/cros-camera:dd:prioritized-hats\n\nBUG=b:350423691\nTEST=Manual: Modify /etc/chrome_dev.conf to enforce camera HaTS\nTEST=unit_tests --gtest_filter=\"*CameraGeneralSurveyHandler*\"\n\nChange-Id: I8fd33914e49a698fbff239bab54727a4d44c7407\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5671999\nReviewed-by: Sean Li <seannli@google.com>\nReviewed-by: Ren-Pei Zeng <kamesan@chromium.org>\nCommit-Queue: Yudhistira Erlandinata <yerlandinata@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325962}",
    "author": "Yudhistira Erlandinata",
    "author_email": "yerlandinata@google.com",
    "date": "2024-07-11T08:46:40+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0056cb6bd676c2f9d8489d3016acbafb925ceb68"
  },
  {
    "id": 295,
    "message": "[High5] Remove obsolete UI elements in Settings\n\nRemove the cr-radio-group which allows switching of password factor,\nsince we do not allow two-way switch.\n\nBug: b:290916811, b:348326316\nChange-Id: I91afa14fcdcc30ba570460ccadca5abfe24f1d41\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5669863\nCommit-Queue: Yunke Zhou <yunkez@google.com>\nReviewed-by: Wes Okuhara <wesokuhara@google.com>\nCr-Commit-Position: refs/heads/main@{#1325961}",
    "author": "Yunke Zhou",
    "author_email": "yunkez@google.com",
    "date": "2024-07-11T08:44:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2ba86a0a3c3ea9adb05386b06c0dded44adb331f"
  },
  {
    "id": 296,
    "message": "Roll devtools-internal from f930141be71a to 62d668b97df8 (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/f930141be71a..62d668b97df8\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: Ic5e3c6f335d0e649dc591059163bf5dffd6bd535\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5694042\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325960}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-11T08:42:40+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4ebbcd5c871193f23660c47c75393349e4da940c"
  },
  {
    "id": 297,
    "message": "PA: Let LightweightQuarantineBranch.Quarantine() take usable_size\n\nIn case of the Extreme LUD, `usable_size` is a known value on the\ncall site. So, it's more performant to pass the value to\nLightweightQuarantineBranch.Quarantine() as an argument and avoid\ncalling PartitionRoot::GetSlotUsableSize() twice.\n\nPinpoint results:\nhttps://pinpoint-dot-chromeperf.appspot.com/job/13243fbfc10000\nhttps://pinpoint-dot-chromeperf.appspot.com/job/1038d580210000\n\nBug: 40944045\nChange-Id: I63d20f2ad06d43f9bf9c2e70741a9c54e5fec6fb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5681540\nCommit-Queue: Yuki Shiino <yukishiino@chromium.org>\nReviewed-by: Sergei Glazunov <glazunov@google.com>\nReviewed-by: Takashi Sakamoto <tasak@google.com>\nReviewed-by: Mikihito Matsuura <mikt@google.com>\nCr-Commit-Position: refs/heads/main@{#1325959}",
    "author": "Yuki Shiino",
    "author_email": "yukishiino@chromium.org",
    "date": "2024-07-11T08:36:06+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ed86e9a2322cd11ae49fe5dd44a9c115b0a72672"
  },
  {
    "id": 298,
    "message": "Roll Skia from 0120468beacf to 1f179dfc01b1 (2 revisions)\n\nhttps://skia.googlesource.com/skia.git/+log/0120468beacf..1f179dfc01b1\n\n2024-07-11 skia-autoroll@skia-public.iam.gserviceaccount.com Roll Skia Infra from 3c9cca49b866 to 67fb1a1b693f (8 revisions)\n2024-07-11 skia-autoroll@skia-public.iam.gserviceaccount.com Roll Dawn from f986604e2165 to c7002cf03eb0 (14 revisions)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/skia-autoroll\nPlease CC fmalita@google.com,skiabot@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Skia: https://bugs.chromium.org/p/skia/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux-blink-rel;luci.chromium.try:linux-chromeos-compile-dbg;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nCq-Do-Not-Cancel-Tryjobs: true\nBug: None\nTbr: fmalita@google.com\nChange-Id: I4de5c35c7f88df2368ebfac0cd254e5339386383\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5694078\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325958}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T08:33:33+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1bd7b5541488f947343e7694f942330014be921e"
  },
  {
    "id": 299,
    "message": "Plumb storage_key_override further into CreateWorker\n\nThis is a followup from a suggestion in https://crrev.com/c/5683463,\nto move the computation of the StorageAccessApiStatus value closer to\nwhere it is actually used.\n\nChange-Id: Iabec02f3de37fe42d5bfddea2150e19329703dd6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5689896\nReviewed-by: Asami Doi <asamidoi@chromium.org>\nCommit-Queue: Asami Doi <asamidoi@chromium.org>\nAuto-Submit: Chris Fredrickson <cfredric@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325957}",
    "author": "Chris Fredrickson",
    "author_email": "cfredric@chromium.org",
    "date": "2024-07-11T08:33:05+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/71b001399a3a7f024c21260fb1b44280fb608e76"
  },
  {
    "id": 300,
    "message": "Roll Amd64 AFDO from 128.0.6586.0_rc-r1-merged to 128.0.6588.0_rc-r1-merged\n\nThis CL may cause a small binary size increase, roughly proportional\nto how long it's been since our last AFDO profile roll. For larger\nincreases (around or exceeding 100KB), please file go/crostc-bug.\n\nPlease note that, despite rolling to chrome/android, this profile is\nused for both Linux and Android.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/afdo-chromium\nPlease CC c-compiler-chrome@google.com,mobiletc-prebuild@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium Main: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: mobiletc-prebuild@google.com\nChange-Id: I70ddcd30d8e12c1c8a562d64f3e684a9a76bd4c5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5694041\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325956}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T08:24:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6aa1d6cdde464a7979f506ff4023e8b8c2b51433"
  },
  {
    "id": 301,
    "message": "ash: Create //chrome/browser/ash/remote_apps:browser_tests\n\nBug: b:335293662\nChange-Id: I5509ab8c57f4192536d06a451f186f9cd6c05643\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5693600\nReviewed-by: Hidehiko Abe <hidehiko@chromium.org>\nCommit-Queue: Georg Neis <neis@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325955}",
    "author": "Georg Neis",
    "author_email": "neis@chromium.org",
    "date": "2024-07-11T08:21:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/79aa90228371dbc7e042e53a80086ccb15665c9a"
  },
  {
    "id": 302,
    "message": "Roll Chrome Win32 PGO Profile\n\nRoll Chrome Win32 PGO profile from chrome-win32-main-1720655837-95377004efed2b57b5b4fe537a5dc298dc7f488c-bb52a88220825ae3b01e1149182da0bbe22ef78c.profdata to chrome-win32-main-1720666785-5c7cd6e770ed4bf2ff4377e7d20028998b6c6759-7025c763169d1e435dbcc2a500e71176f3650ac6.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win32-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:win-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: Ib429ace37f65c19910e46e1a33eaf692223e4c20\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5693127\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325954}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T08:20:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5a783d76b72a60dba12c886a465d01ad7791fffe"
  },
  {
    "id": 303,
    "message": "ash: Add chrome/browser/ash/system/BUILD.gn\n\nBug: b:335294372\nChange-Id: I04a439168b0ee4687fa27a4821cf4ae21709a186\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5692247\nCommit-Queue: Georg Neis <neis@chromium.org>\nReviewed-by: Hidehiko Abe <hidehiko@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325953}",
    "author": "Georg Neis",
    "author_email": "neis@chromium.org",
    "date": "2024-07-11T08:19:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1ab8a09ab25ed749592cf27367beb8b9f3e86570"
  },
  {
    "id": 304,
    "message": "[views-ax] Migrating kPopupForId to new system\n\nThis CL migrates the kPopupForId attribute in Views to be updated\nwhenever its value should change, rather than querying the value\nand computing it only when needed.\n\nThis CL is part of the ViewsAX project:\nhttps://docs.google.com/document/d/1Ku7HOyDsiZem1yaV6ccZ-tz3lO2XR2NEcm8HjR6d-VY/edit#heading=h.ke1u3utej413\n\nBug: 325137417\nChange-Id: Ib08c80471db46ec625fc4a7cafc671231e7d544d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5648297\nCommit-Queue: Sejal Anand <sejalanand@microsoft.com>\nReviewed-by: Patrick Noland <pnoland@chromium.org>\nReviewed-by: Javier Contreras <javiercon@microsoft.com>\nCr-Commit-Position: refs/heads/main@{#1325952}",
    "author": "Sejal Anand",
    "author_email": "sejalanand@microsoft.com",
    "date": "2024-07-11T08:19:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2a27e9e3a603bdbe1023dd47b71a7600b7eec8dc"
  },
  {
    "id": 305,
    "message": "Simplify ConvertToPdfOrigin() inside pdfium_searchify.cc and add tests\n\nExpose the function as ConvertToPdfOriginForTesting() to add some unit\ntests. Also remove the unused width parameter.\n\nChange-Id: Iaa7242f138fbe1a132c58643652daac80cc03aac\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5691718\nCommit-Queue: Lei Zhang <thestig@chromium.org>\nReviewed-by: Chu-Hsuan Yang <chuhsuan@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325951}",
    "author": "Lei Zhang",
    "author_email": "thestig@chromium.org",
    "date": "2024-07-11T08:10:57+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ece448f5597f8a0588447575e10944327047a1a5"
  },
  {
    "id": 306,
    "message": "Roll Chrome Android ARM32 PGO Profile\n\nRoll Chrome Android ARM32 PGO profile from chrome-android32-main-1720655837-4cf2179daeae501b51ac3cbe90cd7496a7a7fb43-bb52a88220825ae3b01e1149182da0bbe22ef78c.profdata to chrome-android32-main-1720677402-738fb1e0100b3714d97e56c02d4b9c567e124bc9-904d57e857b393029c0c549e9378b85f4adc0afb.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-android-arm32-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: Ie265dab0a10f12976f6973cc0b1ae15bb290554c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5694081\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325950}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T08:03:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c3e6dac20189ebf42c9bac60c0478a608ee75103"
  },
  {
    "id": 307,
    "message": "[Blink] Eliminate CanvasResource::Abandon()\n\nThis method is called from CanvasResource::OnDestroy(), which itself is\ncalled from subclass' destructors, in the case where the instance is\nbeing destroyed cross-thread. However, it is both confusing and not\nactually needed:\n\n* Its default implementation calls TearDown()\n* However, all subclasses override that default implementation\n* The GPU-based subclasses override the default implementation to be\n  empty (since effectively nothing can be done in this case)\n* CanvasResourceSharedBitmap overrides the default implementation to\n  clear the reference to the shared bitmap, but it can equally well\n  just do that from its destructor post-invoking OnDestroy() as\n  OnDestroy() invokes TearDown() when on the owning thread, and\n  TearDown() also clears the reference to the shared bitmap\n\nThis CL makes the last change and eliminates CanvasResource::Abandon()\naltogether.\n\nBug: 352263194\nChange-Id: If3bb4b58de8cfe714f8a034f76d5d554f7e36978\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5691295\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCommit-Queue: Colin Blundell <blundell@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325949}",
    "author": "Colin Blundell",
    "author_email": "blundell@chromium.org",
    "date": "2024-07-11T07:59:43+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5d05cc1e0f8bbbe1f3b99cc5b2152215c1e5a64c"
  },
  {
    "id": 308,
    "message": "Roll Chrome Win ARM64 PGO Profile\n\nRoll Chrome Win ARM64 PGO profile from chrome-win-arm64-main-1720655837-65c7b3f0d520262e0dfdf44de7b55533e9d2a250-bb52a88220825ae3b01e1149182da0bbe22ef78c.profdata to chrome-win-arm64-main-1720677402-0103a00fc7f1c2d927083d89f014a48b64a93b89-904d57e857b393029c0c549e9378b85f4adc0afb.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: Ibab6f00e55501568177f6beb880ee53196241d48\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5692926\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325948}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T07:57:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/73dbe77018eeb6d31d0aed7f554d8abbad934ce9"
  },
  {
    "id": 309,
    "message": "[lacros skew tests] Refresh skew tests for M128\n\nThis CL updates the ash version ['128.0.6589.0'] for Lacros version skew testing.\nThis cl only affect linux-lacros config builders like\nlinux-lacros-tester-rel, linux-lacros-rel.\nThis cl will certainly NOT affect Lacros on-device builders\n(lacros-amd64-generic-rel, lacros-amd64-generic-chrome-skylab,\netc) or any other platforms.\n\nIf this CL caused regressions, please revert and stop the autoroller\nat https://luci-scheduler.appspot.com/jobs/chrome/lacros-version-skew-roller\nAlso please file a bug to OS>LaCrOS>Partner, and CC kuanhuang@chromium.org.\n\nR=rubber-stamper@appspot.gserviceaccount.com\n\nBug: None\nChange-Id: Iaa688d5b5b74e2ea440c06deb8d38c543de2c74b\nRequires-Testing: True\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5689218\nAuto-Submit: chrome-weblayer-builder <chrome-weblayer-builder@chops-service-accounts.iam.gserviceaccount.com>\nBot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCommit-Queue: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325947}",
    "author": "chrome-weblayer-builder",
    "author_email": "chrome-weblayer-builder@chops-service-accounts.iam.gserviceaccount.com",
    "date": "2024-07-11T07:49:51+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/52b1914cbbd07ea303b714bc031f6d63ca26416b"
  },
  {
    "id": 310,
    "message": "Updating libjingle XMPP and xmllite libraries's README.chromium to include Revision\n\nRevision is already provided in Version field. No version tags were found to be associated with this revision, so leaving Version as is.\n\nR=jamiewalch@chromium.org, rkjnsn@chromium.org\n\nBug: 350851535\nChange-Id: I69521efac29fec42d8aa151dcf433284be6a65a7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5679611\nReviewed-by: Erik Jensen <rkjnsn@chromium.org>\nAuto-Submit: Jordan Brown <rop@google.com>\nReviewed-by: Jamie Walch <jamiewalch@chromium.org>\nCommit-Queue: Erik Jensen <rkjnsn@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325946}",
    "author": "Jordan",
    "author_email": "rop@google.com",
    "date": "2024-07-11T07:47:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/61bebc14f8e2ac2e5552a86b88abb5ef2eb8de99"
  },
  {
    "id": 311,
    "message": "[lensoverlay] Log new object and tap-to-region selection types\n\nImage context menu search is logged as REGION_SEARCH\n\nBug: 349897105\nChange-Id: I1d129d64ce96855e688f6f741e3dc462f4433b79\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5689952\nReviewed-by: Daniel Cheng <dcheng@chromium.org>\nCommit-Queue: Jason Hu <hujasonx@google.com>\nReviewed-by: Duncan Mercer <mercerd@google.com>\nCr-Commit-Position: refs/heads/main@{#1325945}",
    "author": "Jason Hu",
    "author_email": "hujasonx@google.com",
    "date": "2024-07-11T07:46:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d29fb2099ca2a79024e21037803e3e5416e1a640"
  },
  {
    "id": 312,
    "message": "Roll GoogleTest from 3ef16ef8b30f to b4aaf97d8f7e (1 revision)\n\nhttps://chromium.googlesource.com/external/github.com/google/googletest.git/+log/3ef16ef8b30f..b4aaf97d8f7e\n\n2024-07-10 dmauro@google.com Workaround GCC 12 -Wrestrict false-positive\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/googletest-chromium-autoroll\nPlease CC asully@google.com,jonathanjlee@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:linux_chromium_cfi_rel_ng;luci.chrome.try:win-chrome\nTbr: asully@google.com,jonathanjlee@google.com\nChange-Id: I5f5f989e2f2ca32156c068c6e9b9b124bfb5ba03\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5693561\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1325944}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-11T07:43:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f514f3f94a040cf1a5edf8af390881f01629025e"
  },
  {
    "id": 313,
    "message": "[Blink] Remove usages of CanvasResource::TextureTarget()\n\nThis CL changes calls of CanvasResource::TextureTarget() that are on\nconcrete subclasses to instead inline their implementations as part of\ntrimming down usage of and eventually eliminating the need for this\nmethod.\n\nBug: 351275962\nChange-Id: Ibd8aa25437e07fc52483a7a43ed307165e90877b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5691640\nCommit-Queue: Colin Blundell <blundell@chromium.org>\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325943}",
    "author": "Colin Blundell",
    "author_email": "blundell@chromium.org",
    "date": "2024-07-11T07:41:07+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0238804c3fdc0261a955bf86a11536e13c46d9e1"
  },
  {
    "id": 314,
    "message": "Add ink drop to more emojis button.\n\nFixed: b:347108048\nChange-Id: I4f27ad098006cebe8d0e1b593b1be7b0c432f192\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5685094\nReviewed-by: Michael Cui <mlcui@google.com>\nCommit-Queue: Darren Shen <shend@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1325942}",
    "author": "Darren Shen",
    "author_email": "shend@chromium.org",
    "date": "2024-07-11T07:40:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c23fd9cac5f5c0f9d8bc021fb2cfccb78a394601"
  },
  {
    "id": 315,
    "message": "Roll Xcode beta targets to Xcode16 B4\n\nAlso remove appending\n__PLATFORMS__/iPhoneSimulator.platform/Developer/usr/lib/libXCTestBundleInject.dylib\nto DYLD_INSERT_LIBRARIES in xctestrun for EG tests. As of Xcode16 Beta 4 the presence of that value for that environment variable began causing\nthe following error when attempting to launch EG tests:\n'NSInternalInconsistencyException', reason: 'Cannot initiate shared\nsession more than once.'\n\n1. it was initially added in:\nhttps://codereview.chromium.org/2805133005\nto enable running EG tests on device using the xcodebuild CLI.\n\n2. xctestrun creation logic was moved from test_runner.py to\ntest_apps.py in:\nhttps://chromium-review.googlesource.com/c/chromium/src/+/1961166\nand during that move DYLD_INSERT_LIBRARIES went from\n__PLATFORMS__/iPhoneOS.platform/Developer/usr/lib/libXCTestBundleInject.dylib\nto\n__PLATFORMS__/iPhoneSimulator.platform/Developer/usr/lib/libXCTestBundleInject.dylib\n\nSince 2. even on-device EG tests are using that dylib from iPhoneSimulator.platform, which doesn't seem to be needed (or right).\n\nBug: 355030048\nChange-Id: Ie461d5846696d5412e89db125953a4d93074096d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738253\nReviewed-by: Justin Cohen <justincohen@chromium.org>\nCommit-Queue: Will Yeager <wyeager@google.com>\nReviewed-by: Yue She <yueshe@google.com>\nCr-Commit-Position: refs/heads/main@{#1332961}",
    "author": "Will Yeager",
    "author_email": "wyeager@google.com",
    "date": "2024-07-25T16:13:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2ecf7b4bcdac2d084b035b5cdea183a303fedce6"
  },
  {
    "id": 316,
    "message": "[Editing] Append extra line break if the paragraph ends\nwith or without any layout content.\n\nWhen the paragraph contains some content followed by a comment or\nany element which doesn't have a layout object and the copy\noperation is executed, the content copied does not contain extra\nnew line break character after the paragraph that contained element\nwhich didn't have layout object.\n\nThis happens because the element does not have any layout object\nand when the iterator iterates through the last text node in a\nparagraph, it appends extra newline only when the text node has\na layout object. Thus, the condition where the |ExitNode| is called\nonly when the last child has a layout object is removed and it\nshould be called for every last child. This helps to add extra\nnewline after the paragraph and fixes the bug.\n\nBug: 41350470\nChange-Id: I8242a354977ac5a76a99972156c54d4a646f9eb0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5707431\nCommit-Queue: Pranav Modi <pranavmodi@microsoft.com>\nReviewed-by: Kent Tamura <tkent@chromium.org>\nReviewed-by: Sanket Joshi <sajos@microsoft.com>\nCr-Commit-Position: refs/heads/main@{#1332960}",
    "author": "Pranav Modi",
    "author_email": "pranavmodi@microsoft.com",
    "date": "2024-07-25T16:12:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2a25e3550de22a28539dd73ae93ba33df9181794"
  },
  {
    "id": 325,
    "message": "Add NativeUI dialog for product specifications FRE\n\nThis CL adds NativeUI dialog that hosts the WebUI for product\nspecifications first-run experience. The dialog will show below the top\ncontrol of the window (see Demo). This CL also makes changes to\nProductSpecificationsUI to allow WebUI to communicate with this dialog,\nwhich will be integrated in later CLs.\n\nDemo: https://drive.google.com/file/d/1tzW9gJL_jeVpJCThlfOI5mH2jSIm8TSa/view?usp=sharing&resourcekey=0-R_iEt8EcsMAxb5zaepBApg\n\nBug: 343110207\nChange-Id: I9a6aef9881bf2ab38c31c5502c425389c3753cf9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736463\nReviewed-by: Mei Liang <meiliang@chromium.org>\nCommit-Queue: Yue Zhang <yuezhanggg@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332951}",
    "author": "Yue Zhang",
    "author_email": "yuezhanggg@chromium.org",
    "date": "2024-07-25T16:04:52+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/372683467aa66f1ae1c7663fb58a0daf55686868"
  },
  {
    "id": 317,
    "message": "Roll Dawn from e00312602742 to 60032f110416 (5 revisions)\n\nhttps://dawn.googlesource.com/dawn.git/+log/e00312602742..60032f110416\n\n2024-07-25 cwallez@chromium.org dawn.node: Access AsyncRunner through a weak_ptr.\n2024-07-25 cwallez@chromium.org dawn.node: Fix leak of WGPUAdapterInfo\n2024-07-25 jiawei.shao@intel.com Don't allow using multiple color targets with dual source blending\n2024-07-25 lokokung@google.com Roll third_party/webgpu-cts/ 5167b7163..6c58a1958 (4 commits)\n2024-07-25 dawn-autoroll@skia-public.iam.gserviceaccount.com Roll ANGLE from 84e54d885f87 to 4f498eaa1426 (1 revision)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/dawn-chromium-autoroll\nPlease CC cwallez@google.com,lokokung@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Dawn: https://bugs.chromium.org/p/dawn/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:dawn-android-arm-deps-rel;luci.chromium.try:dawn-android-arm64-deps-rel;luci.chromium.try:dawn-linux-x64-deps-rel;luci.chromium.try:dawn-mac-x64-deps-rel;luci.chromium.try:dawn-mac-arm64-deps-rel;luci.chromium.try:dawn-win10-x64-deps-rel;luci.chromium.try:dawn-win10-x86-deps-rel;luci.chromium.try:dawn-win11-arm64-deps-rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:gpu-fyi-cq-android-arm64\nBug: chromium:341973423,chromium:355006329\nTbr: lokokung@google.com\nTest: Test: dawn_end2end_tests\nInclude-Ci-Only-Tests: true\nChange-Id: I7d2521e9fe186038e95014643b1bdf3b11ad4e49\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741172\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332959}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-25T16:11:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/70ff8ff34647b9e0965d0bcdb8892852010977f0"
  },
  {
    "id": 318,
    "message": "Move common metric code to shared location\n\nThere is some code in UKM that can be leveraged by other metrics. To prevent code duplication, we want to consolidate this code in a common location so all metrics can leverage it.\n\nFixed: 40870309\nChange-Id: Iefd25896380679777030191809a4bb0483b1f669\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5727639\nReviewed-by: Zainab Rizvi <rizvis@google.com>\nReviewed-by: Alexei Svitkine <asvitkine@chromium.org>\nCommit-Queue: Jay Zhou <zhouzj@google.com>\nCr-Commit-Position: refs/heads/main@{#1332958}",
    "author": "Jay Zhou",
    "author_email": "zhouzj@google.com",
    "date": "2024-07-25T16:11:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/afa91424094928eeb227716bcf026a8990273475"
  },
  {
    "id": 319,
    "message": "Disable failing TokenStreamMatcherWithLoadingLazy\n\nBug: 355413130\nChange-Id: Ib9908405e743530fdda9a1fe613977812c6231ab\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741756\nReviewed-by: Igor Kraskevich <kraskevich@google.com>\nAuto-Submit: Peter Pakkenberg <pbirk@chromium.org>\nCommit-Queue: Peter Pakkenberg <pbirk@chromium.org>\nOwners-Override: Peter Pakkenberg <pbirk@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332957}",
    "author": "Peter Pakkenberg",
    "author_email": "pbirk@chromium.org",
    "date": "2024-07-25T16:10:29+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/35790f07c3ca75d4c8165b7d35a0b7ff2bdb37de"
  },
  {
    "id": 320,
    "message": "Add niharm to searchbox OWNERS file.\n\nChange-Id: If2b1cde84ac33fcec12228f7baaab6916cec8825\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740136\nReviewed-by: Moe Ahmadi <mahmadi@chromium.org>\nCommit-Queue: Nihar Majmudar <niharm@google.com>\nCr-Commit-Position: refs/heads/main@{#1332956}",
    "author": "Nihar Majmudar",
    "author_email": "niharm@google.com",
    "date": "2024-07-25T16:10:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/207f0e075bda678a681d8a764587316f396d018e"
  },
  {
    "id": 321,
    "message": "sideways: Fix drop-down select rendering\n\nBug: 40501131\nChange-Id: I4944681f15505cfb244822b1502db881d70b903c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740173\nReviewed-by: Di Zhang <dizhangg@chromium.org>\nCommit-Queue: Di Zhang <dizhangg@chromium.org>\nAuto-Submit: Kent Tamura <tkent@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332955}",
    "author": "Kent Tamura",
    "author_email": "tkent@chromium.org",
    "date": "2024-07-25T16:09:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9e8a51be70c45c8ed7a5839d089260eea5dd102a"
  },
  {
    "id": 322,
    "message": "[ios] Improve IOSChromeScopedTestingLocalState usage\n\nSince IOSChromeScopedTestingLocalState is a scoped object that\ninstall a local state in ApplicationContext, rename the member\nvariables of that type (to avoid clash with member variables\nthat would store a raw_ptr<PrefService> to the local state).\n\nAlso remove IOSChromeScopedTestingLocalState::Get() method and\nchange code to use ApplicationContext to get the local state\nin test too (since this is how it is accessed by production\ncode).\n\nBug: none\nChange-Id: I444c157fa9e0341a5860b561f37fa3ef5f926c0b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735281\nCommit-Queue: Sylvain Defresne <sdefresne@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Aliona Dangla <alionadangla@chromium.org>\nAuto-Submit: Sylvain Defresne <sdefresne@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332954}",
    "author": "Sylvain Defresne",
    "author_email": "sdefresne@chromium.org",
    "date": "2024-07-25T16:09:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/52ed665da783b5ddc809deae82d249d2a80f068b"
  },
  {
    "id": 323,
    "message": "Roll clank/internal/apps from 6d162acff6b7 to 094a4771d2ac (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/6d162acff6b7..094a4771d2ac\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,pbirk@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: chromium:349641010\nTbr: pbirk@google.com\nNo-Try: true\nChange-Id: I1e49705da779719c48d053794f02ab88ec07f239\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740318\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332953}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-25T16:07:31+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9e6a188883e570fe33895c346525a14e33e26dfa"
  },
  {
    "id": 324,
    "message": "Accelerators: Update Dictation accelerator category\n\nIf a user goes to the shortcuts app and browses accessibility shortcuts,\nDictation does not currently show up in that list. This CL updates the\ncategory for the Dictation accelerator to accessibility, with sub\ncategory of text editing, to accurately reflect the type of shortcut\nit is.\n\nBug: N/A\nChange-Id: I5d3bd1cf96ceb115178001bcd1a10fc0b9863428\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735746\nCommit-Queue: Akihiro Ota <akihiroota@chromium.org>\nReviewed-by: Jimmy Gong <jimmyxgong@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332952}",
    "author": "Akihiro Ota",
    "author_email": "akihiroota@chromium.org",
    "date": "2024-07-25T16:05:16+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b011e99ca175e279e24b3ff92d6e8a778e9fc4d2"
  },
  {
    "id": 326,
    "message": "[iOS][Omnibox] Remove remoteSuggestionService observation on disconnect\n\nCleanup remoteSuggestionService observation on disconnect\n\nBug: None\nChange-Id: I834ed87bcd38d14bfd250cf05dc1d932038abf65\nLow-Coverage-Reason: EXPERIMENTAL_CODE remoteSuggestionService is only used for debug.\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736141\nCommit-Queue: Christian Xu <christianxu@chromium.org>\nAuto-Submit: Christian Xu <christianxu@chromium.org>\nReviewed-by: Stepan Khapugin <stkhapugin@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332950}",
    "author": "Christian Xu",
    "author_email": "christianxu@google.com",
    "date": "2024-07-25T16:01:52+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/81f065c894bbfd13606c133a97b2d367b3a9ac4f"
  },
  {
    "id": 327,
    "message": "Enable SecondaryAccountAllowedInArcPolicy by default\n\nThis feature is scheduled to release in M129.\n\nBug: 301157620\nChange-Id: Ib6c502c7b0b9a4c161424e5392278b4074390e6a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738413\nAuto-Submit: Muhammad Hasan Khan <mhasank@chromium.org>\nReviewed-by: Kush Sinha <sinhak@chromium.org>\nCommit-Queue: Muhammad Hasan Khan <mhasank@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332949}",
    "author": "Muhammad Hasan Khan",
    "author_email": "mhasank@chromium.org",
    "date": "2024-07-25T16:01:33+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/10b478f62711151d3f27daa9e37e5c50261d11a9"
  },
  {
    "id": 328,
    "message": "Updating trunk VERSION from 6617.0 to 6618.0\n\n* This is an automated release commit.\n* Do not revert without consulting chrome-pmo@google.com.\nNOAUTOREVERT=true\n\nChange-Id: I3228836cc1cdad47279c350741dc8cb95b46304f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741775\nBot-Commit: Chrome Release Bot (LUCI) <chrome-official-brancher@chops-service-accounts.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332948}",
    "author": "Chrome Release Bot (LUCI)",
    "author_email": "chrome-official-brancher@chops-service-accounts.iam.gserviceaccount.com",
    "date": "2024-07-25T16:00:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ec4525ac64f4cbafa3e3a6f39d0437894948b5f1"
  },
  {
    "id": 329,
    "message": "moveBefore: keep popover open\n\nBug: 353441315\nChange-Id: I85054374a71234db6329d90c3e866dad3cfaf5a2\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5717751\nCommit-Queue: Noam Rosenthal <nrosenthal@chromium.org>\nReviewed-by: Dominic Farolino <dom@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332947}",
    "author": "Noam Rosenthal",
    "author_email": "nrosenthal@chromium.org",
    "date": "2024-07-25T15:59:49+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ea09a8c9fa3c02115e71053033747b1f8d6bdb2f"
  },
  {
    "id": 330,
    "message": "Automatic update from google3\n\nAutomatic update for 2024-07-25 UTC\n\nChange-Id: Ibefbb14b3cf96495c6e2a9b6db9a9d6e5a6ecbf9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739241\nCommit-Queue: PKI Metadata Updates Bot <mdb.chrome-pki-metadata-release-jobs@google.com>\nBot-Commit: PKI Metadata Updates Bot <mdb.chrome-pki-metadata-release-jobs@google.com>\nCr-Commit-Position: refs/heads/main@{#1332946}",
    "author": "PKI Metadata Updates Bot",
    "author_email": "mdb.chrome-pki-metadata-release-jobs@google.com",
    "date": "2024-07-25T15:58:43+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/04a44400156ca0d16497fb447a2605aa74a8dee5"
  },
  {
    "id": 331,
    "message": "Enable limited entropy mode in stable.\n\nWith https://chromium-review.googlesource.com/c/chromium/src/+/5709808 submitted this CL will in effect bring stable to 1%.\n\nFor stable clients, this means a change in behavior at the following two places:\n1. The limited entropy synthetic trial will now be registered.\n2. The limited source will be given to EntropyState when it's constructed, and thus a limited layer will not be filtered out, enabling the limited entropy mode.\n\nThis would equate the behavior of such stable clients to their previous behavior when they were in pre-stable channels.\n\nAll clients at this point either have limited entropy mode enabled currently (M128+) or have it enabled when they were in pre-stable channels (M127). This means they already generated the required pref values, which are:\n- \"user_experience_metrics.limited_entropy_randomization_source\"\n- \"variations_limited_entropy_synthetic_trial_seed_v2\"\n\nBug: 41484361\nChange-Id: Ie4d8e1c293d061a5306c96b500eaa6a928e5b9cc\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738334\nReviewed-by: Caitlin Fischer <caitlinfischer@google.com>\nCommit-Queue: Yulun Zeng <yulunz@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332945}",
    "author": "Yulun Zeng",
    "author_email": "yulunz@chromium.org",
    "date": "2024-07-25T15:57:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ee8c5ba5a2c8752ece290e741cf975024ba60db7"
  },
  {
    "id": 332,
    "message": "webnn: refactor WebNN API conformance tests infrastructure\n\nThis CL is to refactor WebNN API conformance tests infrastructure by\noptimizing utils.js helper and moving tests from JSON files into each\ntest file.\nIt also removes tests of dropped `constant(fillSequence)` op of WebNN API\nchanges [1].\n\n[1] [Remove sequential filling overload of\nconstant()](https://github.com/webmachinelearning/webnn/pull/656)\n\nBug: 331692961\nChange-Id: Ie57095d76ed1a87bcbd93dbade8962a1d4461627\nCq-Include-Trybots: luci.chromium.try:win11-blink-rel,mac14-blink-rel,mac14.arm64-blink-rel\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5668527\nAuto-Submit: Feng Dai <feng.dai@intel.com>\nCommit-Queue: Feng Dai <feng.dai@intel.com>\nReviewed-by: ningxin hu <ningxin.hu@intel.com>\nReviewed-by: David Baron <dbaron@chromium.org>\nReviewed-by: Austin Sullivan <asully@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332944}",
    "author": "BruceDai",
    "author_email": "feng.dai@intel.com",
    "date": "2024-07-25T15:57:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8afcbc6cec783eb93acae7181ab83e688f05b6d2"
  },
  {
    "id": 333,
    "message": "[Sync] Add metrics for the cross user sharing key pair\n\nIn some cases the private key can be missing leading to an incorrect key\npair state. This CL adds some metrics to investigate when the clients\ncan enter such state.\n\nBug: 349736958\nChange-Id: Ic3494189039c586c1d8f3890387997e9745cc1aa\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741494\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Maksim Moskvitin <mmoskvitin@google.com>\nCommit-Queue: Rushan Suleymanov <rushans@google.com>\nCr-Commit-Position: refs/heads/main@{#1332943}",
    "author": "Rushan Suleymanov",
    "author_email": "rushans@google.com",
    "date": "2024-07-25T15:56:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6f264b19dd9af2d5f955fb32986ab29117fcbc25"
  },
  {
    "id": 334,
    "message": "kiosk: browsertest for OnAddGuestView\n\nAdd browser tests for web and chrome app kiosks to cover\nKioskSystemSession::OnAddGuestView().\n\nBug: b/352490540\nTest: browser_tests --gtest_filter=\"*KioskGuestViewTest*\"\nChange-Id: I8f9cd802f48469774b6a9419ebf38414414ff141\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5724352\nReviewed-by: Edman Anjos <edman@chromium.org>\nAuto-Submit: Polina Bondarenko <pbond@chromium.org>\nCommit-Queue: Polina Bondarenko <pbond@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332942}",
    "author": "Polina Bondarenko",
    "author_email": "pbond@google.com",
    "date": "2024-07-25T15:56:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c4cc8287d15b5fbd6578bf4ce5d81ffba66fe3cf"
  },
  {
    "id": 335,
    "message": "Roll Skia from 75994403a8a4 to 529569ce068a (1 revision)\n\nhttps://skia.googlesource.com/skia.git/+log/75994403a8a4..529569ce068a\n\n2024-07-25 sunnyps@chromium.org graphite: Skip adding genID listeners for immutable bitmaps\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/skia-autoroll\nPlease CC jamesgk@google.com,skiabot@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Skia: https://bugs.chromium.org/p/skia/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux-blink-rel;luci.chromium.try:linux-chromeos-compile-dbg;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nCq-Do-Not-Cancel-Tryjobs: true\nBug: chromium:337980338\nTbr: jamesgk@google.com\nChange-Id: Ifcd8bcda58e7757e21dc6479bd92352b3236ce0f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5739861\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332941}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-25T15:53:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3e4cad957609750d5bd5718b2809915f75cd8f10"
  },
  {
    "id": 336,
    "message": "Refactor CreateFromCanvasMailbox() in ExternalCanvasResource\n\nAs part of the ClientSharedImage refactoring, this CL replaces a\nCreateFromCanvasMailbox() call site in ExternalCanvasResource with\nCreateFromCanvasSharedImage().\n\nBug: 1494911\nChange-Id: I606facd8825d950c22c21027025b6925cce8998c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738956\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCommit-Queue: Mingjing Zhang <mjzhang@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332940}",
    "author": "Mingjing Zhang",
    "author_email": "mjzhang@chromium.org",
    "date": "2024-07-25T15:52:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ae317cae887d5a1a46fe0baba52c004ce45aac3c"
  },
  {
    "id": 337,
    "message": "Roll Chromium Variations from 2bf89b61d978 to fed121be014f (11 revisions)\n\nhttps://chromium.googlesource.com/chromium-variations.git/+log/2bf89b61d978..fed121be014f\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/chromium-variations-chromium\nPlease CC chrome-metrics-team@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium Variations: https://bugs.chromium.org/p/chromium/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: chrome-metrics-team@google.com\nChange-Id: I61b1a508ffa8b03198cfff99007cab3a7795519a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741171\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332939}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-25T15:50:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/48068dd59aafd34db197ed19ac0a7d7887a4a03d"
  },
  {
    "id": 338,
    "message": "ash: Add SessionObserver::OnFirstSessionReady\n\nIt is wired under SessionManager::OnUserSessionStartUpTaskCompleted\nso that ash code could use the signal without depending on\nSessionManager.\n\nBug: b:341172705\nChange-Id: I4598108714b5f792bb6958b0d9cf7cab1c160f82\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5730653\nReviewed-by: Hidehiko Abe <hidehiko@chromium.org>\nCommit-Queue: Xiyuan Xia <xiyuan@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332938}",
    "author": "Xiyuan Xia",
    "author_email": "xiyuan@chromium.org",
    "date": "2024-07-25T15:49:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2e3d1c9d793d94eb9af7367c6c173e35ea2bfdcb"
  },
  {
    "id": 339,
    "message": "[iOS][Omnibox] Add GuideName provider in OmniboxPopupPresenterDelegate\n\nThe OmniboxPopupPresenter uses an omnibox LayoutGuide to layout the\npopup at the right location. To supports multiple omnibox instances at\ndifferent positions, this CL adds `omniboxGuideNameForPresenter`.\n\nBug: 353259997\nChange-Id: If62322cbf313daec7c07c086eef91df7d5c88331\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735988\nReviewed-by: Stepan Khapugin <stkhapugin@chromium.org>\nCommit-Queue: Federica Germinario <fedegermi@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Federica Germinario <fedegermi@google.com>\nAuto-Submit: Christian Xu <christianxu@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332937}",
    "author": "Christian Xu",
    "author_email": "christianxu@google.com",
    "date": "2024-07-25T15:49:02+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a57f7b521d55b089c224d2e100de0b88015b702e"
  },
  {
    "id": 340,
    "message": "Roll ios_internal from aeedc36fb241 to e9ae0f17521a\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/aeedc36fb241..e9ae0f17521a\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,joemerramos@google.com,rkgibson@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: I9d23aa684d11b5850c0e3eae61e91ea6936a60e8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741170\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332936}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-25T15:48:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/19c999550331eee0bc04e9d2afce1bbdd7a2115b"
  },
  {
    "id": 358,
    "message": "Change DoIdleWork to not return any value.\n\nThe function has only returned false for a long time and keeping\nthe return value complicates understanding of message pump\ninternal workings.\n\nBug: 40226913\nChange-Id: Ie330faa85b2c039e85a90c5a5e53059a21ac7be0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5570405\nCommit-Queue: Olivier Li <olivierli@chromium.org>\nReviewed-by: Francois Pierre Doray <fdoray@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306458}",
    "author": "Olivier Li",
    "author_email": "olivierli@chromium.org",
    "date": "2024-05-27T17:19:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c01b2134e9b16e55e4fa6e0860a795c525b60294"
  },
  {
    "id": 341,
    "message": "ash: Move `OnAshRestart` from PreProfileInit to PostProfileInit\n\n`login_unlock_throughput_recorder()->OnAshRestart()` lets go of\nthe deferred task runner. Moving it from `PreProfileInit` stage\nto `PostProfileInit` stage where ChromeSessionManager sets up\nthe primary user session. The intentions are:\n- Profile loading is heavy (could take seconds on slow devices) so\n  keeping the deferred tasks until it loads\n- New ash `SessionObserver::OnFirstSessionReady` happens as one\n  deferred task. It should happen after `OnFirstSessionStarted`.\n  `OnFirstSessionStarted` is currently after user profile loading.\n\nBug: b:341172705\nChange-Id: I7395d84f11985a23c3587c0fc2247ad1f520347a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5737468\nReviewed-by: Hidehiko Abe <hidehiko@chromium.org>\nCommit-Queue: Xiyuan Xia <xiyuan@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332935}",
    "author": "Xiyuan Xia",
    "author_email": "xiyuan@chromium.org",
    "date": "2024-07-25T15:47:53+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b51b618f1fd91be6e98aa545796561da9f64a4a8"
  },
  {
    "id": 342,
    "message": "Remove CCT TrackingProtection Onboarding logic\n\nThe intent is to delete the TrackingProtection Onboarding call from the\nBaseCustomRootUiCoordinator as this no longer needs to launch\n\nBug: 355002762\nChange-Id: I2440b7e3c40aa757ddf41f3625ef1b794919314b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738936\nAuto-Submit: Abe Boujane <boujane@google.com>\nReviewed-by: Jinsuk Kim <jinsukkim@chromium.org>\nCommit-Queue: Jinsuk Kim <jinsukkim@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332934}",
    "author": "Abe Boujane",
    "author_email": "boujane@google.com",
    "date": "2024-07-25T15:46:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5d0a4c41abb8f20121e6ff8bafbc1d390fa4d218"
  },
  {
    "id": 343,
    "message": "Roll Chrome Win ARM64 PGO Profile\n\nRoll Chrome Win ARM64 PGO profile from chrome-win-arm64-main-1721886080-35a721c2413ea151bb811c810866010b5445ed96-fee276fa6d547c0aeceda4345e4a2164d7e40af0.profdata to chrome-win-arm64-main-1721908653-6958a39c83779287b627cdbdf891a21633cc426f-1b9101b139af3908dc74bd315eb2963777bfc3ca.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I45665cc1f04311f214058f4dbb231b8f9201f931\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740900\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1332933}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-25T15:45:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a4eddd72d94b1dbba5430f2e21284002e5e08270"
  },
  {
    "id": 344,
    "message": "Allow client arrays except for command decoders.\n\nClient arrays in GL are disabled in the command decoders because it is\nnot feasiable to calculate which data needs to be streamed but they can\nbe used with contexts that are used entirely within the GPU process.\n\nSkia already uses client arrays when running directly on the driver.\n\nBug: angleproject:355034868\nChange-Id: Ie07ec3fed545fc01b2869a0c02c23a0772974f65\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5736354\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCommit-Queue: Geoff Lang <geofflang@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1332932}",
    "author": "Geoff Lang",
    "author_email": "geofflang@chromium.org",
    "date": "2024-07-25T15:41:08+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/52469f3880be97d52b603589fc6b6f35ac3c6300"
  },
  {
    "id": 345,
    "message": "Automatic update from google3\n\nAutomatic update for 2024-05-27 UTC\n\nChange-Id: Ib5b9f57aa69bf448ec74477ad5713c33d6172aff\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569462\nBot-Commit: PKI Metadata Updates Bot <mdb.chrome-pki-metadata-release-jobs@google.com>\nCommit-Queue: PKI Metadata Updates Bot <mdb.chrome-pki-metadata-release-jobs@google.com>\nCr-Commit-Position: refs/heads/main@{#1306471}",
    "author": "PKI Metadata Updates Bot",
    "author_email": "mdb.chrome-pki-metadata-release-jobs@google.com",
    "date": "2024-05-27T18:05:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/90bf773cf426dc382a8841b1523f76502096931f"
  },
  {
    "id": 346,
    "message": "[iOS] Increase keyboard constraint priority on TabGroup creation\n\nThe priority of the constraint ensuring that the buttons of the TabGroup\ncreation screen doesn't overlap with the keyboard was too low,\npreventing it from working sometimes.\n\nIncrease it to make sure it is respected.\n\nFixed: 342282474\nChange-Id: I650a2fba5c59b732fa5be99502a913cee810a2e7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5568826\nReviewed-by: Tommy Martino <tmartino@chromium.org>\nCommit-Queue: Tommy Martino <tmartino@chromium.org>\nCommit-Queue: Gauthier Ambard <gambard@chromium.org>\nAuto-Submit: Gauthier Ambard <gambard@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306470}",
    "author": "Gauthier Ambard",
    "author_email": "gambard@chromium.org",
    "date": "2024-05-27T18:04:49+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1dffc156b7e79e757f88b5c68927e27b7705ce96"
  },
  {
    "id": 347,
    "message": "[IOS][Omnibox] Introduce new content configuration for actions row\n\nThis CL aims to make it easier to review the following CLs in chain.\nThe content view is exactly the same as the one we have for popup row.\nThe actions content configuration subclass the popup row content config\n\nCL.\n\nLow-Coverage-Reason: The feature will be covered by tests in different\nBug: 331344638\nChange-Id: I14631d8c2e6730422a0023fdcab22292a7128dac\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5554491\nReviewed-by: Stepan Khapugin <stkhapugin@chromium.org>\nCommit-Queue: Ameur Hosni <ameurhosni@google.com>\nCr-Commit-Position: refs/heads/main@{#1306469}",
    "author": "Ameur Hosni",
    "author_email": "ameurhosni@google.com",
    "date": "2024-05-27T18:03:18+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/05bc5f1dd2ff3506025c5a6de2a29e8341c35241"
  },
  {
    "id": 348,
    "message": "Undisable failing tests after improving infrastructure.\n\nTests should benefit from explicit wait until sync system is ready.\n\nBug: b/328036610\nChange-Id: I8e33d3fb5ffa1add9c2e5e50a31e9c6959efca63\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573390\nReviewed-by: Anthi Orfanou <anthie@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Tomek Jurkiewicz <tju@google.com>\nCr-Commit-Position: refs/heads/main@{#1306468}",
    "author": "Tomasz Jurkiewicz",
    "author_email": "tju@google.com",
    "date": "2024-05-27T17:58:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9605fed0393cc3b4072606b4677e3fa601307287"
  },
  {
    "id": 349,
    "message": "Fix History and tabs toggle string in settings panel\n\nFixed: 338567028\nChange-Id: I09c32273a3cf8e5e20bb07036d3a0e0cd9f327d1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569772\nReviewed-by: Boris Sazonov <bsazonov@chromium.org>\nCommit-Queue: Mahmoud Rashad <mmrashad@google.com>\nCr-Commit-Position: refs/heads/main@{#1306467}",
    "author": "Mahmoud Rashad",
    "author_email": "mmrashad@google.com",
    "date": "2024-05-27T17:54:09+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c3b9211d357d241eb09b39dfee103d8dbf2c3c4c"
  },
  {
    "id": 350,
    "message": "Rename CompareSpecifics to ProductComparisonSpecifics\n\nSync team requested a more descriptive name. This change is safe\nbecause proto serialization/deserialization is independent of\nthe message name.\n\nBug: b:342471685\nChange-Id: I05dd613aff50bf5bbb828b2f1c7e62a36da95917\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5563408\nReviewed-by: Ayman Almadhoun <ayman@chromium.org>\nReviewed-by: Marc Treib <treib@chromium.org>\nReviewed-by: Ankush Singh <ankushkush@google.com>\nReviewed-by: Rushan Suleymanov <rushans@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: David Maunder <davidjm@chromium.org>\nReviewed-by: Rohit Rao <rohitrao@chromium.org>\nReviewed-by: Ayu Ishii <ayui@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306466}",
    "author": "David Maunder",
    "author_email": "davidjm@google.com",
    "date": "2024-05-27T17:44:49+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f351b3bbcb1de243c78089edf52f0f012a0828b6"
  },
  {
    "id": 351,
    "message": "Use consistent name for InlineCapacity non-type template param\n\nThis follows the naming convention abseil (and presumably Google C++\nstyle) generally use.\n\nChange-Id: I37ee34bf839a163b35b9194f5aef19c1d05dc0be\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5572568\nReviewed-by: Michael Lippautz <mlippautz@chromium.org>\nCommit-Queue: Daniel Cheng <dcheng@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306465}",
    "author": "Daniel Cheng",
    "author_email": "dcheng@chromium.org",
    "date": "2024-05-27T17:44:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/179a7038a8ee6b04b408ed0f39017b2ac8cf11fc"
  },
  {
    "id": 352,
    "message": "Roll Chrome Mac Arm PGO Profile\n\nRoll Chrome Mac Arm PGO profile from chrome-mac-arm-main-1716818245-0470379b3e1a198869e28b20387a8eebbbacfd3a-198e2d983c5e9742450e8e742bc946b0ca524c7b.profdata to chrome-mac-arm-main-1716825535-8006a77934f1092a721eac30122b0a09e06cd987-064bf98dfae995738418058266512389b03ecc28.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-mac-arm-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:mac-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: Ic16c1cfaa0b96f7fb9bba39b3c7e3f27df15243a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5572583\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306464}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-27T17:41:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/35a0ce2a2c6f5b74ac43431ff4f7b4dfba73038b"
  },
  {
    "id": 353,
    "message": "Update SiteInstanceTest expectations\n\nThis CL updates expectations in site_instance_impl_unittest to be\ncompatible with OriginKeyedProcessesByDefault being either enabled or\ndisabled.\n\nBug: 40259221\nChange-Id: I7f88812bea8cf1abe1eec91765cf08baf23608d0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5570981\nReviewed-by: Nasko Oskov <nasko@chromium.org>\nCommit-Queue: W. James Maclean <wjmaclean@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306463}",
    "author": "W. James MacLean",
    "author_email": "wjmaclean@chromium.org",
    "date": "2024-05-27T17:38:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5d010e5c5c2a624aa9e85e9e6983a74ef0228f24"
  },
  {
    "id": 354,
    "message": "[Sync] Rename GetData() -> GetDataForCommit()\n\nThis is a part of ongoing effort to rename GetData() to\nGetDataForCommit() (to reflect the only legit usage of this function in\nthe name).\n\nCurrent ModelTypeSyncBridge definition allows bridges to implement\neither of GetData() or GetDataForCommit(), just for the purpose of this\nrefactoring, that is better to do in multiple CLs (in case you are\nsurprised to not see rename in the base class in the same CL).\n\nThis CL was uploaded by git cl split.\n\nR=ellyjones@chromium.org\n\nBug: 331763450\nChange-Id: Ia392e16995e30f301ce227503b2b3c69f782f239\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5568232\nAuto-Submit: Maksim Moskvitin <mmoskvitin@google.com>\nCommit-Queue: Elly FJ <ellyjones@chromium.org>\nReviewed-by: Elly FJ <ellyjones@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306462}",
    "author": "Maksim Moskvitin",
    "author_email": "mmoskvitin@google.com",
    "date": "2024-05-27T17:32:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4152f3f3e8778739d9dde768c40e354680862f8b"
  },
  {
    "id": 355,
    "message": "Cleanup killswitch AutofillUseTypedCreditCardNumber\n\nBug: 1498664\nChange-Id: I0b41f281ff2185c60f7599c449ad67545d011135\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569735\nReviewed-by: Peter Kotwicz <pkotwicz@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Vidhan Jain <vidhanj@google.com>\nCr-Commit-Position: refs/heads/main@{#1306461}",
    "author": "Vidhan Jain",
    "author_email": "vidhanj@google.com",
    "date": "2024-05-27T17:29:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4e41a6e08fc2a18cb38f324d6bd3f7c92b0f5e2f"
  },
  {
    "id": 356,
    "message": "Roll clank/internal/apps from 86864fe6b0e8 to 5e75c8b0e1da (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/86864fe6b0e8..5e75c8b0e1da\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,fhorschig@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: fhorschig@google.com\nNo-Try: true\nChange-Id: I0c1365a70be05a5b97971c439482811ab22860ba\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573268\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306460}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-05-27T17:22:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c4645406b3ddf6ea2b68db458d55c9a0828db9b4"
  },
  {
    "id": 357,
    "message": "Disable new authentication API on windows\n\nBUg: 342316548\nChange-Id: I032beff616eb37bfb51995e1917c2c2512c3cdaa\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5560985\nReviewed-by: Vasilii Sukhanov <vasilii@chromium.org>\nCommit-Queue: Karol Sygiet <sygiet@google.com>\nCr-Commit-Position: refs/heads/main@{#1306459}",
    "author": "Karol Sygiet",
    "author_email": "sygiet@google.com",
    "date": "2024-05-27T17:20:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0dca1c8bf87e5758e1c58d97503055553f1ad1e1"
  },
  {
    "id": 359,
    "message": "[Uno-D] Add browser test for feature rollback\n\nFixed: 322510754\nChange-Id: I39835bcb31c4388ef8f53ca6a83c576cd1956567\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569746\nCommit-Queue: David Roger <droger@chromium.org>\nReviewed-by: Monica Basta <msalama@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306457}",
    "author": "David Roger",
    "author_email": "droger@chromium.org",
    "date": "2024-05-27T17:18:43+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5c3429f1218f2835587eec0d11889497c3cea421"
  },
  {
    "id": 360,
    "message": "[iOS] Do not change the color of the button if disabled\n\nBased on the UX feedback, the save button's color and the label color\nshould not change when it is disabled.\n\nScreenshot: https://screenshot.googleplex.com/4bxeS638s98in7c.png\nBug: 342562229, 1482269\nChange-Id: Id656b8a607638173b099290230d2b67ab0674bdd\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569146\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Gauthier Ambard <gambard@chromium.org>\nCommit-Queue: Vidhan Jain <vidhanj@google.com>\nCr-Commit-Position: refs/heads/main@{#1306456}",
    "author": "Vidhan Jain",
    "author_email": "vidhanj@google.com",
    "date": "2024-05-27T17:18:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7d26bb675d83b835525156ffae12b04b0c48cd00"
  },
  {
    "id": 361,
    "message": "Implement IOSTabModelURLVisitDataFetcher\n\nImplement the iOS version of the URL data fetcher for tab resumption.\n\nFixed: 342389640\nChange-Id: Ibe9793c35cbd45dd4945a2aeabd06f1d89195a46\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5476371\nAuto-Submit: Olivier Robin <olivierrobin@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Olivier Robin <olivierrobin@chromium.org>\nReviewed-by: Gauthier Ambard <gambard@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306455}",
    "author": "Olivier Robin",
    "author_email": "olivierrobin@google.com",
    "date": "2024-05-27T17:15:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/361344bf7c8552764c3b8589c3f4c4ce3f46d7b0"
  },
  {
    "id": 362,
    "message": "Update UMA code definition to match XML definition\n\nBug: b/321240396, 342378638\nChange-Id: Ie8f0f09d3de142c82ae3664857dd153372162cfa\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5572686\nReviewed-by: Anthi Orfanou <anthie@google.com>\nCommit-Queue: James Lee <ljjlee@google.com>\nCr-Commit-Position: refs/heads/main@{#1306454}",
    "author": "James Lee",
    "author_email": "ljjlee@google.com",
    "date": "2024-05-27T17:14:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9c54c150664f04854914f32e0e9ba37f1458162b"
  },
  {
    "id": 363,
    "message": "Roll Help App from 7zm5XJ4Hgoz8AM4O3... to yqZO7NQ8YxQCedBaR...\n\nRelease_Notes: http://go/help_app-x20/relnotes/Main/help_app_nightly_202405270800_RC00.html\n\nhttps://chrome-infra-packages.appspot.com/p/chromeos_internal/apps/help_app/app/+/yqZO7NQ8YxQCedBaRmGW3KkG0O5v6G5MSQosMgdHB9cC\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/help-app-chromium-autoroll\nPlease CC cros-essential-apps-dev@chromium.org,help-app@grotations.appspotmail.com,jomag@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: b/332931850\nTbr: help-app@grotations.appspotmail.com\nChange-Id: I159e4f0cc034782661e04c116bf9099b4402c7b8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573747\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306453}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-05-27T17:13:01+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/165f9eac8f7e909be3a12df31da135e598073fb6"
  },
  {
    "id": 364,
    "message": "Update CrossOriginOpenerPolicyBrowserTest for OriginKeyedProcessesByDefault\n\nThis CL updates a number of CrossOriginOpenerPolicyBrowserTest tests\nto be compatible with features::kOriginKeyedProcessesByDefault being\nenabled by either (i) disabling the feature if it's incompatible with\nthe intent of the test, or (ii) updating test expectations to account\nfor the fact that the feature forces same-site, cross-origin frames\ninto different SiteInstances and processes.\n\nBug: 40259221\nChange-Id: I1f4ad9eee0382ff6037c05cd46dc33ca37177e5a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5570602\nCommit-Queue: W. James Maclean <wjmaclean@chromium.org>\nReviewed-by: Nasko Oskov <nasko@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306452}",
    "author": "W. James MacLean",
    "author_email": "wjmaclean@chromium.org",
    "date": "2024-05-27T17:12:48+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1959acd32d292e220badb8fdbb436341791bfe42"
  },
  {
    "id": 365,
    "message": "[//gpu] Remove client_id from HostGMBManager::DestroyGpuMemoryBuffer()\n\nThis method is callend only privately and only with `client_id_` as the\nvalue.\n\nBug: 342905184\nChange-Id: I08f88f0b8fcb40ac42328db8745249284dab0e6d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569195\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCommit-Queue: Colin Blundell <blundell@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306451}",
    "author": "Colin Blundell",
    "author_email": "blundell@chromium.org",
    "date": "2024-05-27T17:09:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/797d7502d06f5f10c393a37cfd99dafb79f056cf"
  },
  {
    "id": 366,
    "message": "ozone/drm: Restrict maximum cursor buffer size\n\nSome platforms, e.g. MediaTek, support the cursor plane as an additional\nhardware plane. As a result, the supported cursor size from DRM driver\ncan be as large as 2048 or 4096 (or potentially larger).\n\nIn such cases, allocating two 4096*4096*4 buffers for a cursor plane\nis wasteful - that's 128MB of RAM solely for rendering the cursor.\n\nOn the other hand, Intel GPUs support a maximum cursor size of 256,\nwhich should be sufficiently large for most non-Intel platforms as well.\n\nTherefore, limit the cursor size to a maximum of 256, so that only two\n256*256*4 buffers will be allocated i.e. 0.5MB.\nThat directly frees up hundred of RAM back to the system on platforms\nlike MediaTek.\n\nBug: b/342940839\n\nChange-Id: I140a692613cc563b8e8cbe3f054f088a8b0f83ab\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5568541\nReviewed-by: Peter McNeeley <petermcneeley@chromium.org>\nReviewed-by: Chen-Yu Tsai <wenst@google.com>\nCommit-Queue: Peter McNeeley <petermcneeley@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306450}",
    "author": "Fei Shao",
    "author_email": "fshao@chromium.org",
    "date": "2024-05-27T17:06:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5c65041f496a122d09e5e2a8fcfe28a9e0884734"
  },
  {
    "id": 367,
    "message": "Record disk-space metrics in PM's metrics provider\n\nChange-Id: Ib1b32da3bcb397b744bf6336f2711110dd7ddeac\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5392180\nReviewed-by: Francois Pierre Doray <fdoray@chromium.org>\nCommit-Queue: Anthony Vallée-Dubois <anthonyvd@chromium.org>\nReviewed-by: Olivier Li <olivierli@chromium.org>\nReviewed-by: Olivier Li Shing Tat-Dupuis <olivierli@google.com>\nCr-Commit-Position: refs/heads/main@{#1306449}",
    "author": "Anthony Vallée-Dubois",
    "author_email": "anthonyvd@google.com",
    "date": "2024-05-27T16:58:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/064bf98dfae995738418058266512389b03ecc28"
  },
  {
    "id": 368,
    "message": "Roll Amd64 AFDO from 127.0.6501.0_rc-r1-merged to 127.0.6501.0_rc-r2-merged\n\nThis CL may cause a small binary size increase, roughly proportional\nto how long it's been since our last AFDO profile roll. For larger\nincreases (around or exceeding 100KB), please file go/crostc-bug.\n\nPlease note that, despite rolling to chrome/android, this profile is\nused for both Linux and Android.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/afdo-chromium\nPlease CC c-compiler-chrome@google.com,mobiletc-prebuild@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium Main: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: mobiletc-prebuild@google.com\nChange-Id: I4250fc780cd14e79bfca59b7ad9a3949d5d8b790\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573150\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306448}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-27T16:58:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c8cf0b8e893f0c4827b7f986073dbcaa93a7e161"
  },
  {
    "id": 369,
    "message": "Roll ANGLE from 81fe86ae2053 to 8db2d61a815c (1 revision)\n\nhttps://chromium.googlesource.com/angle/angle.git/+log/81fe86ae2053..8db2d61a815c\n\n2024-05-27 lehoangquyen@chromium.org Metal: properly handle base/max level for immutable textures.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/angle-chromium-autoroll\nPlease CC angle-team@google.com,solti@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in ANGLE: https://bugs.chromium.org/p/angleproject/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86\nBug: None\nTbr: solti@google.com\nChange-Id: I0ff546a860e6c30396e4d8f2905e913ddd322978\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5572483\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306447}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-27T16:56:08+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/14ccaa3f9d9e73fcc867b7a03e6bd73fb805c7c2"
  },
  {
    "id": 370,
    "message": "Fix testReopenGroupFromAnotherWindow\n\nThe test was failing because when simulating key presses when entering\nthe group name, it relied on the default lowercase/uppercase mode of the\nkeyboard, which can differ iPhone vs iPad, simulator vs device.\n\nThis CL uses a group name starting with a number, eschewing the problem.\n\nFixed: 339415297\nChange-Id: I57f5c035b14a7fc3c00c5fd017ad17484fbedb68\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569770\nAuto-Submit: Louis Romero <lpromero@google.com>\nCommit-Queue: Ameur Hosni <ameurhosni@google.com>\nReviewed-by: Ameur Hosni <ameurhosni@google.com>\nCr-Commit-Position: refs/heads/main@{#1306446}",
    "author": "Louis Romero",
    "author_email": "lpromero@google.com",
    "date": "2024-05-27T16:55:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9f1ddba873233efeee79a104c2470ed9ba95c656"
  },
  {
    "id": 371,
    "message": "Add pixeltest for the pin keyboard view\n\nBug: b:286527724\nChange-Id: I487fa862e35f8ed8a505eea761fed0e1a00e470b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5482077\nCommit-Queue: Istvan Nagy <iscsi@google.com>\nReviewed-by: Denis Kuznetsov <antrim@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306445}",
    "author": "Istvan Nagy",
    "author_email": "iscsi@google.com",
    "date": "2024-05-27T16:52:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/aea32f1f82c959344d05dada54da46e4c1992471"
  },
  {
    "id": 372,
    "message": "Add LazyBlinkTimezoneInit to testing config.\n\nBug: 40287434\nChange-Id: I2aa0db0d04e1c1c24889f5d99b1e4d5f7fadc945\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5559484\nCommit-Queue: Francois Pierre Doray <fdoray@chromium.org>\nReviewed-by: Colin Blundell <blundell@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306444}",
    "author": "François Doray",
    "author_email": "fdoray@chromium.org",
    "date": "2024-05-27T16:47:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/dd7bd692cb4cfa8d35d375f8c20b964734622351"
  },
  {
    "id": 373,
    "message": "Roll winapi-util: 0.1.6 => 0.1.8 in //third_party/rust.\n\nThis change moves the dependency for windows bindings from winapi to\nthe windows-rs crates published by Microsoft. Those crates come with\n.lib files that must be made available for linking, so we add them\nto gnrt_config.toml under the `native_libs_roots` variable.\n\nNative libs are copied into the toolchain's `./rustlib` directory\nand that directory is added to the library search path in order to\nfind them there. This ensures the correct architecture library is\nused for host toolchains.\n\nThis CL has been created semi-automatically.  The expected review\nprocess and other details can be found at\n//tools/crates/create_update_cl.md\n\nUpdated crates:\n\n* winapi-util: 0.1.6 => 0.1.8\n\nNew crates:\n\n* windows-sys@0.52.0\n* windows-targets@0.52.5\n* windows_aarch64_gnullvm@0.52.5 (as placeholder)\n* windows_aarch64_msvc@0.52.5\n* windows_i686_gnu@0.52.5 (as placeholder)\n* windows_i686_gnullvm@0.52.5 (as placeholder)\n* windows_i686_msvc@0.52.5\n* windows_x86_64_gnu@0.52.5 (as placeholder)\n* windows_x86_64_gnullvm@0.52.5 (as placeholder)\n* windows_x86_64_msvc@0.52.5\n\nRemoved crates:\n\n* winapi-i686-pc-windows-gnu@0.4.0\n* winapi-x86_64-pc-windows-gnu@0.4.0\n* winapi@0.3.9\n\nChromium `supply-chain/config.toml` policy requires that the following\naudit criteria are met (note that these are the *minimum* required\ncriteria and `supply-chain/audits.toml` can and should record a stricter\ncertification if possible;  see also //docs/rust-unsafe.md):\n\n* winapi-util@0.1.8, windows-sys@0.52.0, windows-targets@0.52.5,\n  windows_aarch64_gnullvm@0.52.5, windows_aarch64_msvc@0.52.5,\n  windows_i686_gnu@0.52.5, windows_i686_gnullvm@0.52.5,\n  windows_i686_msvc@0.52.5, windows_x86_64_gnu@0.52.5,\n  windows_x86_64_gnullvm@0.52.5, windows_x86_64_msvc@0.52.5:\n  does-not-implement-crypto, safe-to-run\n\nFixed: 342194487\nChange-Id: I156d557e63a295abcd7fae5bc025b9531aec5adc\nCq-Include-Trybots: chromium/try:android-rust-arm32-rel\nCq-Include-Trybots: chromium/try:android-rust-arm64-dbg\nCq-Include-Trybots: chromium/try:android-rust-arm64-rel\nCq-Include-Trybots: chromium/try:linux-rust-x64-dbg\nCq-Include-Trybots: chromium/try:linux-rust-x64-rel\nCq-Include-Trybots: chromium/try:win-rust-x64-dbg\nCq-Include-Trybots: chromium/try:win-rust-x64-rel\nDisable-Rts: True\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5533267\nReviewed-by: Łukasz Anforowicz <lukasza@chromium.org>\nAuto-Submit: danakj <danakj@chromium.org>\nCommit-Queue: danakj <danakj@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306443}",
    "author": "danakj",
    "author_email": "danakj@chromium.org",
    "date": "2024-05-27T16:45:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2a140ac6fc6943d7689fc41f8c4573e3d27cba7d"
  },
  {
    "id": 374,
    "message": "[DBSC] Verify rotation request signature before sending\n\nThe metrics data suggests that the message signing on TPMs might be\nflaky. Investigate this by adding an extra signature verification step\ninto the assertion token generation process.\n\nIf the verification fails, try one more time before giving up. Also\nrecord a histogram about how often the second attempt succeeds to\nunderstand whether retries make sense.\n\nBug: b/340578557\nChange-Id: Id5bea36b7e6c256a162dabde6a9a190616354cf0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569235\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Monica Basta <msalama@chromium.org>\nCommit-Queue: Alex Ilin <alexilin@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306442}",
    "author": "Alex Ilin",
    "author_email": "alexilin@chromium.org",
    "date": "2024-05-27T16:43:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8b3cf809cd724fbf16cc060d7174876920ea2ba8"
  },
  {
    "id": 375,
    "message": "[skyvault] Add migration policy\n\nAdds a policy to control if local files are uploaded when SkyVault is\nturned on.\n\nR=poromov@chromium.org,hidehiko@chromium.org\n\nFixed: b/343400162\nChange-Id: Ia97f5d5986f1bc8a6c57d8d7f08d24b49a675b58\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5576818\nReviewed-by: Sergey Poromov <poromov@chromium.org>\nCommit-Queue: Aida Zolic <aidazolic@chromium.org>\nReviewed-by: Hidehiko Abe <hidehiko@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309831}",
    "author": "Aida Zolic",
    "author_email": "aidazolic@chromium.org",
    "date": "2024-06-04T10:37:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/28b4a4ee7ed4374aa399a38182c0e0fa7ddde78b"
  },
  {
    "id": 376,
    "message": "privacy-hub: Hide turn on camera access button in CRD session\n\nThis CL makes sure we do not display a button to turn on camera access\nin a CRD session.\n\nBug: b:330137918\nChange-Id: I6441d9832e685872da452f59af7b4d440ed09fa7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575474\nCommit-Queue: Md Shahadat Hossain Shahin <shahinmd@google.com>\nReviewed-by: Xiaohui Chen <xiaohuic@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309830}",
    "author": "Md Shahadat Hossain Shahin",
    "author_email": "shahinmd@google.com",
    "date": "2024-06-04T10:27:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/381f35db91ed1c2cd3d25c89b4ab513cc6ef5273"
  },
  {
    "id": 377,
    "message": "Roll clank/internal/apps from 22e9f448420d to 3e923fa04ebc (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/22e9f448420d..3e923fa04ebc\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC agarwaltushar@google.com,chrome-brapp-engprod@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: agarwaltushar@google.com\nNo-Try: true\nChange-Id: I72cd4f4a8fd8dbca0d389c126d005e173e6ec8b0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5596508\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309829}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-04T10:19:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/73be3f305683dfdc5f81a4346b01fbd90a802ec0"
  },
  {
    "id": 378,
    "message": "Roll ios_internal from a4b594ba836d to d9c065773edd\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/a4b594ba836d..d9c065773edd\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC alionadangla@google.com,chrome-brapp-engprod@google.com,vidhanj@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: Iecb5777126475b740ec340329e935ae0325931b1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5596241\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309828}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-04T10:18:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/61ecc0192a42471a5b7c35701047f93517c5b8b4"
  },
  {
    "id": 379,
    "message": "[views-ax] Migrate GetAccessibleName to ViewAccessibility getter in\nchrome/* folder\n\nThis CL migrates callers in path chrome/* from using View::GetAccessibleName to\nnew ViewAccessibility::GetCachedName. Also, this CL does not introduce a change \nin behavior. You can find the design doc for the ViewsAX project here:\n\nhttps://docs.google.com/document/d/1Ku7HOyDsiZem1yaV6ccZ-tz3lO2XR2NEcm8HjR6d-VY/edit\n\nBug: 325137417\nChange-Id: I37508e8ba8ad292d2adbf880f77964b8002273a2\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5577512\nReviewed-by: Javier Contreras <javiercon@microsoft.com>\nReviewed-by: Benjamin Beaudry <benjamin.beaudry@microsoft.com>\nReviewed-by: David Trainor <dtrainor@chromium.org>\nCommit-Queue: Divyansh Mangal <dmangal@microsoft.com>\nCr-Commit-Position: refs/heads/main@{#1309827}",
    "author": "Divyansh Mangal",
    "author_email": "dmangal@microsoft.com",
    "date": "2024-06-04T10:11:33+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f75226a669d512f9513270c62dc43b0f7468ee29"
  },
  {
    "id": 380,
    "message": "Roll Skia from 761dba1a7c41 to 43cb2d4f238a (3 revisions)\n\nhttps://skia.googlesource.com/skia.git/+log/761dba1a7c41..43cb2d4f238a\n\n2024-06-04 skia-autoroll@skia-public.iam.gserviceaccount.com Roll SwiftShader from 4bb94d6c235c to 31bee5b9c506 (2 revisions)\n2024-06-04 skia-autoroll@skia-public.iam.gserviceaccount.com Roll Skia Infra from 5e66633855fd to 2d509f23bc91 (11 revisions)\n2024-06-04 skia-autoroll@skia-public.iam.gserviceaccount.com Roll ANGLE from 42a61f6e70c9 to af72bf7f4270 (15 revisions)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/skia-autoroll\nPlease CC jmbetancourt@google.com,skiabot@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Skia: https://bugs.chromium.org/p/skia/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux-blink-rel;luci.chromium.try:linux-chromeos-compile-dbg;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nCq-Do-Not-Cancel-Tryjobs: true\nBug: chromium:303266564,chromium:343991300\nTbr: jmbetancourt@google.com\nTest: Test: Test: presubmit, and the problematic files mentioned in\nChange-Id: Ia30ce95fa546f671f6f44b6c17d5ab8c3c42f3e1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5596238\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309826}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-04T10:08:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7f4f23dc8896f5e2f423d9b74dd8b56c67a56e72"
  },
  {
    "id": 381,
    "message": "[Android] Use monospaced font for the suggested plus address.\n\nScreenshot: https://screenshot.googleplex.com/9bbxMxxTWq5TcSe.\n\nBug: 342126606\nChange-Id: I45736a2fde597a9e5372ceba75baa74eb2dbaeec\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5593871\nReviewed-by: Jan Keitel <jkeitel@google.com>\nCommit-Queue: Timofey Chudakov <tchudakov@google.com>\nCr-Commit-Position: refs/heads/main@{#1309825}",
    "author": "Timofey Chudakov",
    "author_email": "tchudakov@google.com",
    "date": "2024-06-04T09:55:09+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/549168cdf073de7a7915a3f2e18e6101474337e6"
  },
  {
    "id": 382,
    "message": "Roll eche_app from Fxh0n44f722q569k_... to V7l2yr0da33J7Ty2l...\n\nRelease_Notes: http://go/eche-x20/relnotes/Main/eche_20240603_RC00.html\n\nhttps://chrome-infra-packages.appspot.com/p/chromeos_internal/apps/eche_app/app/+/V7l2yr0da33J7Ty2l57gj9ddh_o1B1pLJmoEazmx-qwC\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/eche-app-chromium\nPlease CC christyfranks@google.com,dnishi@google.com,eche-app@grotations.appspotmail.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: b/337912574\nTbr: eche-app@grotations.appspotmail.com\nChange-Id: I014ab558d60bdb016abfad7df9a7947fc8c65e7d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5595216\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309824}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-04T09:52:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/107031718da6746be2fb40bd44123b1a5a134b12"
  },
  {
    "id": 383,
    "message": "[blink/heap] Report pooled memory in memory dumps\n\nBug: 326303884\nChange-Id: Id852d2446b6aec95a4e299c9b957554162cd40d7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5594108\nCommit-Queue: Benoit Lize <lizeb@chromium.org>\nReviewed-by: Michael Lippautz <mlippautz@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309823}",
    "author": "Benoît Lizé",
    "author_email": "lizeb@chromium.org",
    "date": "2024-06-04T09:51:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e4812c0714c851b5ff4a2abe395f4398398d980a"
  },
  {
    "id": 384,
    "message": "[canvas] Remove the feature to lose context in background\n\nThis was added in\nhttps://chromium-review.googlesource.com/c/chromium/src/+/3297289, but\nhasn't been an active experiment for years, and is likley not shippable\nas clients do not expect the context (and content) to be lost.\n\nRemove the feature to simplify the code.\n\nBug: 343013357\nChange-Id: I965841ba8870854d5f2b8ba17600b224d493dd9a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5589810\nReviewed-by: Colin Blundell <blundell@chromium.org>\nReviewed-by: Kentaro Hara <haraken@chromium.org>\nCommit-Queue: Benoit Lize <lizeb@chromium.org>\nReviewed-by: Yi Xu <yiyix@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309822}",
    "author": "Benoît Lizé",
    "author_email": "lizeb@chromium.org",
    "date": "2024-06-04T09:51:12+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9167501f173b2ee1c718e324b0fa7e7b179bb0e4"
  },
  {
    "id": 385,
    "message": "Roll Help App from 0g464vhcx_TFPbKWG... to G3-wD8C0PqfO7_JUA...\n\nRelease_Notes: http://go/help_app-x20/relnotes/Main/help_app_20240603_RC00.html\n\nhttps://chrome-infra-packages.appspot.com/p/chromeos_internal/apps/help_app/app/+/G3-wD8C0PqfO7_JUAVE3lFJt2TMX0bOn0B4yr62KYGYC\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/help-app-chromium-autoroll\nPlease CC cros-essential-apps-dev@chromium.org,help-app@grotations.appspotmail.com,jomag@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: b/332931850\nTbr: help-app@grotations.appspotmail.com\nChange-Id: I973262f53086e62b0fffc23fef94330c81ad3986\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5595184\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309821}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-04T09:49:16+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1d4b5cb060069c8faed0125a71a6b3024be4ad52"
  },
  {
    "id": 386,
    "message": "Reland \"[android] Enable shipping both snapshots on android\"\n\nThis is a reland of commit dbe657ecce09220eee9c16d8cfd2e585349ca368\n\nRelanding without changes, the android-x86 build issues are fixed in\nhttps://chromium-review.googlesource.com/c/chromium/src/+/5526637\n\nOriginal change's description:\n> [android] Enable shipping both snapshots on android\n>\n> Reland of crrev.com/c/3219567. We want to experiment with this again,\n> potentially limiting it to high-end Android only (depending on impact).\n>\n> Binary-Size: Increase is temporary for experiment.\n> Bug: 40200623, 40539769\n> Change-Id: Ide7e11ab9b9c1fcf9c42cb97df9e831dbd0cb6e7\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5518514\n> Auto-Submit: Leszek Swirski <leszeks@chromium.org>\n> Commit-Queue: Andrew Grieve <agrieve@chromium.org>\n> Reviewed-by: Andrew Grieve <agrieve@chromium.org>\n> Cr-Commit-Position: refs/heads/main@{#1297053}\n\nBinary-Size: Increase is temporary for experiment.\nBug: 40200623, 40539769\nChange-Id: I879b06477a36aa9f69495d877878b0836deb92cd\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5592691\nCommit-Queue: Leszek Swirski <leszeks@chromium.org>\nAuto-Submit: Leszek Swirski <leszeks@chromium.org>\nReviewed-by: Andrew Grieve <agrieve@chromium.org>\nReviewed-by: Kentaro Hara <haraken@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309820}",
    "author": "Leszek Swirski",
    "author_email": "leszeks@google.com",
    "date": "2024-06-04T09:45:15+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9a6b6c3edba4305b3a31660c0eba7454d1084209"
  },
  {
    "id": 387,
    "message": "[DevTools] Add browser test for `getHostConfig`\n\nPrevious CL: https://crrev.com/c/5344413\nCorresponding DevTools CL: https://crrev.com/c/5538381\n\nChange-Id: I1d1489f76fee7f3924cdb52467a9b66f4a2a1e33\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5578016\nReviewed-by: Alex Rudenko <alexrudenko@chromium.org>\nReviewed-by: Simon Zünd <szuend@chromium.org>\nCommit-Queue: Wolfgang Beyer <wolfi@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309819}",
    "author": "Wolfgang Beyer",
    "author_email": "wolfi@chromium.org",
    "date": "2024-06-04T09:42:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/bbdd7f713523d5dd30bf0bfef174e31aebf31b63"
  },
  {
    "id": 388,
    "message": "Roll src-internal from 656a685e929f to 598508bff793 (2 revisions)\n\nhttps://chrome-internal.googlesource.com/chrome/src-internal.git/+log/656a685e929f..598508bff793\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/src-internal-chromium-autoroll\nPlease CC chrome-browser-infra-team,rushans@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: rushans@google.com\nChange-Id: I50d97a36b35ba5c7a49c9781d02c625232feba90\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5595186\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309818}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-04T09:41:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/479d856f94bc6e3bf7974838b315c2b110475285"
  },
  {
    "id": 389,
    "message": "Roll Chrome Mac Arm PGO Profile\n\nRoll Chrome Mac Arm PGO profile from chrome-mac-arm-main-1717466336-6303ea66413777b82b6cfc01d41ecaf01efc45d9-6f8e2dd141d44755500c702ba4549bb8c89a11c0.profdata to chrome-mac-arm-main-1717487976-f877eddc575309d76d2acaf29c14a2e511078bbc-954c375f0b63ef4b333c85bcb3de47351b0a3b69.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-mac-arm-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:mac-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: If771ac7342296338147164b110442ac7505dc931\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5595900\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309817}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-04T09:41:15+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c034ab717f61d5861894c6668a7ce5abf72e81f0"
  },
  {
    "id": 390,
    "message": "Roll androidx from rMlRImWiede5FBK_M... to Afv3GhUvvnWPAgJf4...\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/androidx-chromium\nPlease CC clank-build@google.com,clank-library-failures@google.com,wnwen@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:android-internal-binary-size;luci.chrome.try:android-internal-rel\nTbr: clank-library-failures@google.com\nChange-Id: I6abc26fe27de7c4a00f596f45dd42fd40b24dc92\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5595219\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309816}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-04T09:38:16+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fca31267bc6ef824e314e3680d9cc21ad85584ca"
  },
  {
    "id": 391,
    "message": "WebNN: Use DML_OPERATOR_DIAGONAL_MATRIX1 when DirectML feature level >= 5.1\n\nThis CL updates the triangular implementation of DirectML backend to\nsupport DML_DIAGONAL_MATRIX1_OPERATOR_DESC when DirectML feature level\nis greater than or equal to DML_FEATURE_LEVEL_5_1.\n\nBug: 332574921\nChange-Id: I5e2e23de7457a88f7c34bb47aa1b50a101e595bc\nCq-Include-Trybots: luci.chromium.try:win11-blink-rel\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5540226\nCommit-Queue: Lisha Guo <lisha.guo@intel.com>\nReviewed-by: ningxin hu <ningxin.hu@intel.com>\nReviewed-by: Rafael Cintron <rafael.cintron@microsoft.com>\nCr-Commit-Position: refs/heads/main@{#1309815}",
    "author": "lisa0314",
    "author_email": "lisha.guo@intel.com",
    "date": "2024-06-04T09:35:12+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6a481cdb1c54d7c2baec74f8ac0443ea0e6a3ebc"
  },
  {
    "id": 392,
    "message": "GWP-ASan: Rename experiment\n\nThis change is trivial. Rather than describing the experiment arm\n(\"larger browser reservation\"), let's provide the name of the larger\nexperiment (it's just GWP-ASan, being tinkered with in the year 2024.)\n\nBug: b:326007059\nChange-Id: Ie669e4adc27e66963c6e7b25b5e1717412321833\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5595712\nReviewed-by: Matthew Denton <mpdenton@chromium.org>\nCommit-Queue: Kalvin Lee <kdlee@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309814}",
    "author": "Kalvin Lee",
    "author_email": "kdlee@chromium.org",
    "date": "2024-06-04T09:33:36+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a2d4b7948ac381cdb43574be45c0bd579243b947"
  },
  {
    "id": 393,
    "message": "privacy-hub: Make app permissions v2 strings translatable 2/3\n\nBug: b:343181924\nChange-Id: I41b8a062342ee35698e428a4d884e43d7d95ad9f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5588408\nReviewed-by: Dominick Ng <dominickn@chromium.org>\nCommit-Queue: Md Shahadat Hossain Shahin <shahinmd@google.com>\nCr-Commit-Position: refs/heads/main@{#1309813}",
    "author": "Md Shahadat Hossain Shahin",
    "author_email": "shahinmd@google.com",
    "date": "2024-06-04T09:32:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0688bb38888f0420bb3a85fc27e1ca90d378332f"
  },
  {
    "id": 394,
    "message": "privacy-hub: Make app permissions v2 strings translatable 1/3\n\nBug: b:343181924\nChange-Id: I9d1be8b33f36b081864ad00e122bce1c54129274\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5591675\nReviewed-by: Nikki Moteva <moteva@google.com>\nReviewed-by: Dominick Ng <dominickn@chromium.org>\nReviewed-by: Wes Okuhara <wesokuhara@google.com>\nCommit-Queue: Md Shahadat Hossain Shahin <shahinmd@google.com>\nCr-Commit-Position: refs/heads/main@{#1309812}",
    "author": "Md Shahadat Hossain Shahin",
    "author_email": "shahinmd@google.com",
    "date": "2024-06-04T09:32:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ca18464aff77a7bd453f4f89b447d030d4d20a4c"
  },
  {
    "id": 395,
    "message": "privacy-hub: Make app permissions v2 strings translatable 3/3\n\nBug: b:343181924\nChange-Id: I068c248ea2ac88870b3b9dd954a2f573a904632a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5591679\nReviewed-by: Dominick Ng <dominickn@chromium.org>\nCommit-Queue: Md Shahadat Hossain Shahin <shahinmd@google.com>\nCr-Commit-Position: refs/heads/main@{#1309811}",
    "author": "Md Shahadat Hossain Shahin",
    "author_email": "shahinmd@google.com",
    "date": "2024-06-04T09:31:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8ec4b819acc79d15609d0e44b854dbb04cb8bc19"
  },
  {
    "id": 396,
    "message": "Partition Alloc: Cutting the Cord!\n\nPartition Alloc is all grown up and ready to fly the coop! This CL\nremoves the dependency on Chromium's //build files, allowing Partition\nAlloc to build independently.\n\nBug: 41481467\nChange-Id: I125b94bf950b3c806d8c52927a76319b21d3375a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5587524\nCommit-Queue: Arthur Sonzogni <arthursonzogni@chromium.org>\nReviewed-by: Bartek Nowierski <bartekn@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309810}",
    "author": "Arthur Sonzogni",
    "author_email": "arthursonzogni@chromium.org",
    "date": "2024-06-04T09:31:15+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/dca75e17e9a13fcddc21cde8cc2a1da98785d000"
  },
  {
    "id": 397,
    "message": "partition_alloc: Check no external dependencies in GN.\n\nBug: 41481467\nChange-Id: I93763ea3b9f2e2f12d39b97e1f7b4444706c2885\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5587870\nCommit-Queue: Arthur Sonzogni <arthursonzogni@chromium.org>\nReviewed-by: Takashi Sakamoto <tasak@google.com>\nCr-Commit-Position: refs/heads/main@{#1309809}",
    "author": "Arthur Sonzogni",
    "author_email": "arthursonzogni@chromium.org",
    "date": "2024-06-04T09:28:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/837e9d5fe30858072763b3b73f86962afdeba90b"
  },
  {
    "id": 398,
    "message": "ide_query: handle use_reclient\n\nuse_remoteexec doesn't control reclient, but it just controls\nremote execution in the build.\nuse_reclient control whether it uses reclient or not.\n\ncheck use_reclient instead of use_remoteexec\n\nBug: b/342038479\nChange-Id: I0e4865ab974801291ae78197900c68c96a909113\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5595376\nReviewed-by: Junji Watanabe <jwata@google.com>\nCommit-Queue: Fumitoshi Ukai <ukai@google.com>\nAuto-Submit: Fumitoshi Ukai <ukai@google.com>\nReviewed-by: Richard Wang <richardwa@google.com>\nCr-Commit-Position: refs/heads/main@{#1309808}",
    "author": "Fumitoshi Ukai",
    "author_email": "ukai@google.com",
    "date": "2024-06-04T09:25:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/128f48d59f8389248ebab912505062f962a04e23"
  },
  {
    "id": 399,
    "message": "[Autofill] Log user action Autofill_ProfileDisabled\n\nLogs user action Autofill_ProfileDisabled on disabling pref\nautofill.profile_enabled, if the change was caused by a user or\nextension.\n\nChange-Id: I842282edf55bb833bcc69ef3278bb02ab93864c2\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5576820\nReviewed-by: Christoph Schwering <schwering@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Florian Leimgruber <fleimgruber@google.com>\nCommit-Queue: Wolfgang Billenstein <bwolfgang@google.com>\nCr-Commit-Position: refs/heads/main@{#1309807}",
    "author": "Wolfgang Billenstein",
    "author_email": "bwolfgang@google.com",
    "date": "2024-06-04T09:18:02+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9aa45d1dac55aadb9f4833bf9e1f822775a0b2a0"
  },
  {
    "id": 400,
    "message": "Delay expiry of histograms causing alerts.\n\nUpdates the expires_after attribute for 63 histograms that have\nbeen used to generate alerts in the past 180 days and have an expiry date\n(or milestone equivalent) between 2024-07-19 and 2024-10-02.\n\nTBR=chromium-metrics-reviews@google.com\n\nChange-Id: I240e770c39579ab95420ed9f18e0e4d994b172ba\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5596169\nBot-Commit: Chrome Metrics Logs <chrome-metrics-team+robot@google.com>\nCommit-Queue: Chrome Metrics Logs <chrome-metrics-team+robot@google.com>\nCr-Commit-Position: refs/heads/main@{#1309806}",
    "author": "Chrome Metrics Logs",
    "author_email": "chrome-metrics-team+robot@google.com",
    "date": "2024-06-04T09:09:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a518052223a42789875a4d4c390c73bcada3402c"
  },
  {
    "id": 401,
    "message": "Fix testBrowsingContextTreeIsNotEmpty\n\nIn the fix ChromeDriver awaits the promise returned by runMapperInstance\nfunction of BiDiMapper to be sure that the mapper launch process has\nfinished. Without this there can be situations that the mapper leaves\nsome commands without response.\n\nBug: 343891977\nChange-Id: I638848217ce78fd6c2ce5ffb98149fd82722db69\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5594070\nReviewed-by: Maksim Sadym <sadym@chromium.org>\nCommit-Queue: Vladimir Nechaev <nechaev@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309805}",
    "author": "Vladimir Nechaev",
    "author_email": "nechaev@chromium.org",
    "date": "2024-06-04T09:08:14+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c3a29559395f3fd550249e0eef5cf54fe101d667"
  },
  {
    "id": 402,
    "message": "[Autofill] Fix styling of disabled icons in the autofill popup\n\nUnfortunately when the disabled style is applied on the whole ImageView,\nMacOS grays out the whole bounding rectangle (screen/4ibAzybF5CKDnQy).\n\nThis CL moves the graying out to the ImageModel before creating the\nImageView.\n\nThe CL also IWYUfies the files it touches.\n\nImpl: screen/98j8n9oKQG65bxw\n\nFixed: 344470541\nChange-Id: I2b3822165a77cb07097738894dddedaaf496a271\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5588315\nReviewed-by: Jan Keitel <jkeitel@google.com>\nCommit-Queue: Adem Derinel <derinel@google.com>\nCr-Commit-Position: refs/heads/main@{#1309804}",
    "author": "Adem Derinel",
    "author_email": "derinel@google.com",
    "date": "2024-06-04T09:06:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/632ba09ef59a706d3d614e72ac78575f9e5fd5f0"
  },
  {
    "id": 403,
    "message": "Roll clank/internal/apps from 4c3578a439db to 22e9f448420d (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/4c3578a439db..22e9f448420d\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: \nNo-Try: true\nChange-Id: Icf9dd3de8583d17f2ad0393fac066a50f1687a64\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5595897\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1309803}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-04T09:01:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ec93c8cb52ded9343fb6ee659a28d656c79c00d0"
  },
  {
    "id": 404,
    "message": "[Telemetry] Refactor ProbeServiceAsh\n\nRefactor ProbeServiceAsh to remove outdated usage of factory.\n\nBUG=b/335129312\nTEST=CQ\n     testing/xvfb.py tools/autotest.py -C out_linux_lacros/Release --run_all chrome/browser/lacros/cros_apps/api/*\n     testing/xvfb.py out/Default/browser_tests --gtest_filter=\"*TelemetryExtensionTelemetryApiBrowserTest*\"\n     testing/xvfb.py out/Default/browser_tests --gtest_filter=\"*TelemetryExtensionApiGuard*\"\n     testing/xvfb.py out/Default/unit_tests --gtest_filter=\"*ApiGuardDelegate*\"\n     testing/xvfb.py out/Default/unit_tests --gtest_filter=\"*ProbeServiceAsh*\"\n\nChange-Id: Ib376ac9ab7989b1a50d164278e3a1c1c58ba24be\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5568397\nCommit-Queue: Ethan Cheng <yycheng@google.com>\nReviewed-by: Chung-sheng Wu <chungsheng@google.com>\nReviewed-by: Kerker Yang <kerker@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1309802}",
    "author": "Ethan Cheng",
    "author_email": "yycheng@google.com",
    "date": "2024-06-04T09:01:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5d7a42f46bb3e118499698c1b5355614a6a6185d"
  },
  {
    "id": 405,
    "message": "Implement a hash of the cookies specified by Cookie-Indices.\n\nThis can be used to compare the cookies for an intercepted request with\nthose sent on a previous request. Since equality is the only operation\ncurrently required, this is equivalent to storing the request cookies,\nbut doesn't require us to actually do so.\n\nBug: 328628231\nChange-Id: I6e64d871422e1e5e6c7c9827ace2bc86c0f8a1c9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5544845\nCommit-Queue: Jeremy Roman <jbroman@chromium.org>\nReviewed-by: Adam Rice <ricea@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306741}",
    "author": "Jeremy Roman",
    "author_email": "jbroman@chromium.org",
    "date": "2024-05-28T15:20:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cb37d66b16a15b2ab45c940aca356ef2f11aa9ab"
  },
  {
    "id": 406,
    "message": "Roll Perfetto from e02750a8e14f to a886234f224a (1 revision)\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/e02750a8e14f..a886234f224a\n\n2024-05-28 carlscab@google.com Remove unneeded dependency\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-chromium-autoroll\nPlease CC perfetto-bugs@google.com,primiano@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-perfetto-rel\nBug: None\nTbr: perfetto-bugs@google.com\nChange-Id: I91fc96af26eee0ee4cc6901f47d18aa80e91bbf4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575522\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306740}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-28T15:13:08+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e00e0c3f117350da18bf4dbb348cbc05ecc23f48"
  },
  {
    "id": 407,
    "message": "[IOS][Omnibox] Add answer proto migration flag in ios\n\nThis CL Includes the answer proto migration flag to iOS about_flags.\n\nBug: 327497146\nChange-Id: I9ce6cc1ac78edf35b18c2335532cdbbaf2317008\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5563724\nAuto-Submit: Ameur Hosni <ameurhosni@google.com>\nCommit-Queue: Ameur Hosni <ameurhosni@google.com>\nReviewed-by: Stepan Khapugin <stkhapugin@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306739}",
    "author": "Ameur Hosni",
    "author_email": "ameurhosni@google.com",
    "date": "2024-05-28T15:11:12+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3d764c91ef334cbf8b7d5c68d39dbd010c4c3a6c"
  },
  {
    "id": 408,
    "message": "Fix AvoidScheduleWorkDuringNativeEventProcessing behavior in nested loops.\n\nThis also requires a small fix to MessagePumpEpoll to avoid marking\nthe handling the wakeup as native work.\n\nBug: 1523528\nChange-Id: I9dea042916338f4765dc04a1c750ddcdc71e7630\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5574025\nCommit-Queue: Olivier Li <olivierli@chromium.org>\nReviewed-by: Francois Pierre Doray <fdoray@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306738}",
    "author": "Olivier Li Shing Tat-Dupuis",
    "author_email": "olivierli@chromium.org",
    "date": "2024-05-28T15:09:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a72d31de82d5a3e41d0189c30594e3ebc9931870"
  },
  {
    "id": 409,
    "message": "Update DEPS to point to latest libjpeg_turbo.\n\nThe latest update resolves an O(n^2) slowdown issue discovered by\nthe fuzzer.\n\nBug: 339704200\nChange-Id: Ifa8f9c6965b2489b923ae449fb70953c637daf86\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5570282\nReviewed-by: Leon Scroggins <scroggo@google.com>\nAuto-Submit: John Stiles <johnstiles@google.com>\nCommit-Queue: Leon Scroggins <scroggo@google.com>\nCr-Commit-Position: refs/heads/main@{#1306737}",
    "author": "John Stiles",
    "author_email": "johnstiles@google.com",
    "date": "2024-05-28T15:06:56+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c7e12c4980152d1508e303ceba1ba5b0e7ed5568"
  },
  {
    "id": 410,
    "message": "Update GameModeEnabler window when it reused across WindowTrackers\n\nWhen GameModeController detects a change in window focus, it may pass\nan existing GameModeEnabler to the new WindowTracker, which avoids\nmomentarily disabling game mode. Prior to this CL, the reused\nGameModeEnabler would contain an not-owned reference to the initially\ntracked window, which could possibly become stale. This change forcibly\nupdates the reference to use the newly tracked window whenever a\nGameModeEnabler is reused.\n\nBug: b:341220423\nChange-Id: Ibff3ff7dd1013858cd71dcf95c7da179b445a76f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5560127\nReviewed-by: Mitsuru Oshima <oshima@chromium.org>\nCommit-Queue: Andrew Wolfers <aswolfers@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306736}",
    "author": "Andrew Wolfers",
    "author_email": "aswolfers@chromium.org",
    "date": "2024-05-28T15:05:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/30e212de7077084aab0da37eb433701c055ea6df"
  },
  {
    "id": 411,
    "message": "[Android] Add icon before the suggested plus address.\n\nThe icon and text views are placed inside a linear layout\noriented horizontally. The background is implemented using\na custom drawable, because the standard rounded_rectangle_surface_1\nhas a different radius for the rounded corners.\n\nNote: the icon for this view hasn't been designed yet, so this\nCL used a placeholder icon, which will be changed later.\n\nBefore: https://screenshot.googleplex.com/AcEU8KJfv84RqXz\nAfter: https://screenshot.googleplex.com/69bPvfz7wbxsnof\nMocks: https://screenshot.googleplex.com/5cY3FEknUVc2SQV\n\nBug: 342126606\nChange-Id: Id9b0b9ed44033356acb2c71e26d91a0c3c263c69\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573527\nReviewed-by: Jan Keitel <jkeitel@google.com>\nReviewed-by: Norge Vizcay <vizcay@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Friedrich Horschig <fhorschig@chromium.org>\nCommit-Queue: Timofey Chudakov <tchudakov@google.com>\nCr-Commit-Position: refs/heads/main@{#1306735}",
    "author": "Timofey Chudakov",
    "author_email": "tchudakov@google.com",
    "date": "2024-05-28T15:05:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4374d3eafb4e3213aaf27e43ab01844ecc8e45c1"
  },
  {
    "id": 433,
    "message": "[iOS][QD] Trigger menu when any part of the popup cell is tapped\n\nThis CL makes the popup menu be triggered when any part of the\nTableViewPopupCell is tapped. To accomplish this, a UIControl is\ncustomized to have a static title on the left, a subtitle with the\ncurrent selected menu item on the right along with a up and down\nchefron.\n\nScreenrecording: http://dr/file/d/116J3nDr6jxEnDC-up4ylHKRHbuUsCFUD\n\nFixed: 341077648, 341306396\nBug: 335387869\nChange-Id: I11873f14b28a28919bbeeb4ec28cc84784586314\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5563907\nCommit-Queue: Filipa Senra <fsenra@google.com>\nReviewed-by: Jérôme Lebel <jlebel@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306713}",
    "author": "Filipa Senra",
    "author_email": "fsenra@google.com",
    "date": "2024-05-28T14:08:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e4a6034c07ddd085731e18eaa0bc266898f0135c"
  },
  {
    "id": 412,
    "message": "Roll Chrome Linux PGO Profile\n\nRoll Chrome Linux PGO profile from chrome-linux-main-1716875762-1e408da82017edbceca14c813613164aa0d7bca0-fb0fdb1001863cf64dd23b425114ca9b2300a415.profdata to chrome-linux-main-1716897398-da360a04ee393068e8ccb96e635e516a741a1912-85f98417ff9cdf4b22fcd1d5c4df34c82542d846.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-linux-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I72a2b839337da1aab68286cc049696d4a57acf3c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575236\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306734}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-28T15:03:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4be4ffe8e723f32e629e1f28342c7680962d5326"
  },
  {
    "id": 413,
    "message": "Reland \"Verify SkColor4f when reading DrawLooper\"\n\nThis reverts commit c466444c28fb2e43606ada66436b4804f8490ea7.\n\nReason for revert: Put everything behind `valid_` check.\n\nOriginal change's description:\n> Revert \"Verify SkColor4f when reading DrawLooper\"\n>\n> This reverts commit 466991fb11a90cc033e7ee977c8d9df05bf7d848.\n>\n> Reason for revert: causes MSan failures:\n> https://ci.chromium.org/ui/p/chromium/builders/ci/Linux%20MSan%20Tests/47926/overview\n>\n> Original change's description:\n> > Verify SkColor4f when reading DrawLooper\n> >\n> > Call `Read(&color)` instead of `ReadSimple(&color)` to properly verify\n> > the color read. Also verify alpha values in [0, 1].\n> >\n> > Bug: 334082978\n> > Change-Id: I66f3e676deab1924b28730b0197a769b41b0eebe\n> > Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5570149\n> > Reviewed-by: Xianzhu Wang <wangxianzhu@chromium.org>\n> > Commit-Queue: Ben Wagner <bungeman@chromium.org>\n> > Cr-Commit-Position: refs/heads/main@{#1305932}\n>\n> Bug: 334082978\n> Change-Id: I2f9f43178ceb4ae4648e4aed373806061638d796\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5572590\n> Auto-Submit: Daniel Cheng <dcheng@chromium.org>\n> Commit-Queue: Daniel Cheng <dcheng@chromium.org>\n> Bot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\n> Owners-Override: Daniel Cheng <dcheng@chromium.org>\n> Cr-Commit-Position: refs/heads/main@{#1306208}\n\nBug: 334082978\nChange-Id: Ib7c907c0ff2f84325fba8b687eeb4bb7c4c0d7d4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5570786\nReviewed-by: Xianzhu Wang <wangxianzhu@chromium.org>\nReviewed-by: Daniel Cheng <dcheng@chromium.org>\nCommit-Queue: Ben Wagner <bungeman@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306733}",
    "author": "Ben Wagner",
    "author_email": "bungeman@chromium.org",
    "date": "2024-05-28T14:57:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a17a0307c844b361305364f17444787ffe83189c"
  },
  {
    "id": 414,
    "message": "Adding extension telemetry event router\n\nChange-Id: I3eac8e8dd3d48c107ab748b69409542924555819\nBug: b/331794658\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5485326\nReviewed-by: Shanthanu Bhardwaj <xanth@google.com>\nReviewed-by: Dominique Fauteux-Chapleau <domfc@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nAuto-Submit: Alexander-Joseph Silva <ajsilva@google.com>\nReviewed-by: Sébastien Lalancette <seblalancette@chromium.org>\nCommit-Queue: Alexander-Joseph Silva <ajsilva@google.com>\nCr-Commit-Position: refs/heads/main@{#1306732}",
    "author": "Alexander-Joseph Silva",
    "author_email": "ajsilva@google.com",
    "date": "2024-05-28T14:57:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/76ff27d02bb6ce818c7c4046c63d5cd53fd08f59"
  },
  {
    "id": 415,
    "message": "Updating XTBs based on .GRDs from branch main\n\nChange-Id: I30cb998b7bf544cce694e082ff4720a441c0202a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575934\nCommit-Queue: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nAuto-Submit: Ben Mason <benmason@chromium.org>\nBot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306731}",
    "author": "Ben Mason",
    "author_email": "benmason@chromium.org",
    "date": "2024-05-28T14:52:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d7d37d4fd2d720993ff03a6e9e75194a0b2cb2ff"
  },
  {
    "id": 416,
    "message": "Roll Chrome Lacros Arm64 Generic PGO Profile\n\nRoll Chrome Lacros Arm64 PGO profile from chrome-chromeos-arm64-generic-main-1716811142-122047e7a411782f96db57dca36e602895238999-50580ff929617d096636b2a3759511713d52ceba.profdata to chrome-chromeos-arm64-generic-main-1716897398-05e5861a0e463d6611c6a2704eb6ee3fe5b4d3e7-85f98417ff9cdf4b22fcd1d5c4df34c82542d846.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-lacros-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I70ebbf8f316e15553a01b8838f70571ce5f7669c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575211\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306730}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-28T14:51:48+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5eb4510d233fecbe9e3766684146b787ab25d09f"
  },
  {
    "id": 417,
    "message": "Use \"In your Google Account\" as a section title and omit page title\n\nBug: 338567294\nChange-Id: I334cc9ed7dd7840aff48424b83b462d64eafd418\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573533\nCommit-Queue: Mahmoud Rashad <mmrashad@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Boris Sazonov <bsazonov@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306729}",
    "author": "Mahmoud Rashad",
    "author_email": "mmrashad@google.com",
    "date": "2024-05-28T14:51:15+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f0b0895a55bc5df5e5c8182f48d6123af8bc32c6"
  },
  {
    "id": 434,
    "message": "Decouple AuthPanelEventDispatcher from the LoginTextfield\n\nBug: b:286527724\nChange-Id: Id8b3efa92e37f46bd5d212e51c55c60a1cbc694a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5574068\nCommit-Queue: Istvan Nagy <iscsi@google.com>\nReviewed-by: Denis Kuznetsov <antrim@chromium.org>\nAuto-Submit: Istvan Nagy <iscsi@google.com>\nReviewed-by: Elie Maamari <emaamari@google.com>\nCr-Commit-Position: refs/heads/main@{#1306712}",
    "author": "Istvan Nagy",
    "author_email": "iscsi@google.com",
    "date": "2024-05-28T14:07:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/249bd0ae73988f66a445caec7b6211279871687e"
  },
  {
    "id": 418,
    "message": "Roll ios_internal from 4c059d7bc861 to c47f7adb825e\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/4c059d7bc861..c47f7adb825e\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC christianxu@google.com,chrome-brapp-engprod@google.com,stkhapugin@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: I47fd77b4356979cfc1fffba44f2f78b3e1409648\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575235\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306728}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-05-28T14:47:29+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/351326de183a84f22b5938bec7af5250ad8d9abb"
  },
  {
    "id": 419,
    "message": "Remove searchbox focus when submitting query.\n\nBug: 339003039\nChange-Id: I7b4f12523b71cdcbf1ee5794c1680b0da2346fc5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5570369\nReviewed-by: Moe Ahmadi <mahmadi@chromium.org>\nCommit-Queue: Nihar Majmudar <niharm@google.com>\nCr-Commit-Position: refs/heads/main@{#1306727}",
    "author": "Nihar Majmudar",
    "author_email": "niharm@google.com",
    "date": "2024-05-28T14:46:05+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ccbc7bf50b978baee170ad3569905a0697ca90ab"
  },
  {
    "id": 420,
    "message": "Roll WebRTC from d297c4c23f44 to 06815534d2da (5 revisions)\n\nhttps://webrtc.googlesource.com/src.git/+log/d297c4c23f44..06815534d2da\n\n2024-05-28 perkj@webrtc.org Add SchedulableNetworkBehavior and tests.\n2024-05-28 webrtc-version-updater@webrtc-ci.iam.gserviceaccount.com Update WebRTC code version (2024-05-28T04:05:18).\n2024-05-27 phancke@meta.com Split SSL adapters from main ssl build target 2/2\n2024-05-27 manashi@google.com Revert \"Propagate arrival time inside NetEq\"\n2024-05-27 perkj@webrtc.org Remove dependency from rtp_rtcp module to remote_bitrate_estimator\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/webrtc-chromium-autoroll\nPlease CC webrtc-chromium-sheriffs-robots@google.com,webrtc-infra@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in WebRTC: https://bugs.chromium.org/p/webrtc/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: webrtc-chromium-sheriffs-robots@google.com\nChange-Id: Ifb63f4f39ed664a5302ef0767f47ea5ef191be92\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575521\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306726}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-28T14:45:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5587990cb9134334a524301ee4936001e870cfa0"
  },
  {
    "id": 421,
    "message": "policy: Refactor URLBlocklistManagerTest\n\n* Move helper functions to anonymous namespace.\n* Rename helper functions to more specific names.\n* Create wrapper functions for `GetMatch` to avoid boolean params.\n* Delete unnecessary `policy::`.\n* Use local alias in URLBlocklistManagerTest.UseBlocklistState.\n\nTest: ./components_unittests --gtest_filter=*URLBlocklistManagerTest*\nBug: b:339377401\nChange-Id: I36c5376651202f8ffe7cbae2ecdc249b8fb5ae01\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573968\nCommit-Queue: Emmanuel Arias Soto <eariassoto@google.com>\nReviewed-by: Roland Bock <rbock@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306725}",
    "author": "Emmanuel Arias Soto",
    "author_email": "eariassoto@google.com",
    "date": "2024-05-28T14:41:20+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c87dfa68f43c22f901b33ba2ee36b776c513e7c7"
  },
  {
    "id": 422,
    "message": "Roll Chrome Lacros Amd64 Generic PGO Profile\n\nRoll Chrome Lacros Amd64 PGO profile from chrome-chromeos-amd64-generic-main-1716854702-755ae0e8aaf8e5346d9ab35691181fa4c565cbda-48aae9b1e004951f346d50c49752048af759084e.profdata to chrome-chromeos-amd64-generic-main-1716897398-1524e851c68f28d249f5d45233b5722f1618a3fa-85f98417ff9cdf4b22fcd1d5c4df34c82542d846.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-lacros-amd64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I69eb1c4bcb427e18d1faa31242acd622fe289fcd\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575465\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306724}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-28T14:37:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b4a42e4227e904a1080cfabcde8e9c15b8bc4a12"
  },
  {
    "id": 423,
    "message": "[Viz] Remove client_id->buffers mapping from HostGMBManager\n\nHostGpuMemoryBufferManager now manages requests from the browser only,\nfor which `client_id_` is always used. This CL collapses the\nclient_id -> gmb_id -> gmb_handle mapping used by HostGMBManager with\ngmb_id -> gmb_handle.\n\nBug: 342905184\nChange-Id: Ib9e8bccf508f9d11ce78be54263e11dd99c98990\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569461\nReviewed-by: Vasiliy Telezhnikov <vasilyt@chromium.org>\nCommit-Queue: Colin Blundell <blundell@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306723}",
    "author": "Colin Blundell",
    "author_email": "blundell@chromium.org",
    "date": "2024-05-28T14:35:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1ddbb5f57c8d828b4c7c432c94f817eedf210a80"
  },
  {
    "id": 424,
    "message": "dnt: Remove early swap for back forward visual transitions.\n\nThe feature is no longer necessary since we don't use full screen\noverlays for the UX.\n\nR=rakina@chromium.org\n\nChange-Id: I012f80b3715540805fb9d8b9541da5a1ff9854bf\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5572654\nReviewed-by: Rakina Zata Amni <rakina@chromium.org>\nAuto-Submit: Khushal Sagar <khushalsagar@chromium.org>\nCommit-Queue: Khushal Sagar <khushalsagar@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306722}",
    "author": "Khushal Sagar",
    "author_email": "khushalsagar@chromium.org",
    "date": "2024-05-28T14:34:48+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3e7ea37caf164e43ba960e196588544f31eee424"
  },
  {
    "id": 425,
    "message": "Roll devtools-internal from 626927be85cd to a2ef4c8fadff (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/626927be85cd..a2ef4c8fadff\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/devtools/devtools-frontend.git/+log/39381be348f8171b3387b9affc8881fee7d1c50d..437b27d3926277278af0fdf78ff627af52bc12d7\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: Id355ec50572d349903647d967032ef9dd162f684\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575460\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306721}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-05-28T14:29:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a5fd23a66647c96b17a72cb82c9a178f8be46973"
  },
  {
    "id": 426,
    "message": "Serialization added for blocking bit\n\nBug: b/326058805\nChange-Id: I0442f41965f5c1b6f429454905602ab68e352b18\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569183\nAuto-Submit: Alexander-Joseph Silva <ajsilva@google.com>\nReviewed-by: Dominique Fauteux-Chapleau <domfc@chromium.org>\nReviewed-by: thefrog <thefrog@chromium.org>\nCommit-Queue: thefrog <thefrog@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306720}",
    "author": "Alexander-Joseph Silva",
    "author_email": "ajsilva@google.com",
    "date": "2024-05-28T14:26:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b72921340243dfb1e654fe3b9aa2c51d06dc4643"
  },
  {
    "id": 427,
    "message": "policy: Refactor URLBlocklist::GetURLBlocklistState\n\n* Move the logic that finds the highest priority filter for an URL\n  to a helper function.\n* Add helper function IsWildcardBlocklist.\n\nTest: ./components_unittests --gtest_filter=*URLBlocklistManagerTest*\nBug: b:339377401\n\nChange-Id: I5f74d9a92c1fb952b2b4605d1209e8bb6ffd9a2d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573668\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Emmanuel Arias Soto <eariassoto@google.com>\nReviewed-by: Jeroen Dhollander <jeroendh@google.com>\nReviewed-by: Roland Bock <rbock@google.com>\nCr-Commit-Position: refs/heads/main@{#1306719}",
    "author": "Emmanuel Arias Soto",
    "author_email": "eariassoto@google.com",
    "date": "2024-05-28T14:26:07+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/73985c20359fdf63604054a5275da9aa65a10f89"
  },
  {
    "id": 428,
    "message": "policy: Clean up URLBlocklistManager\n\n* Delete unused using statements.\n* Delete unused function URLBlocklist::Size().\n* Add missing <string> include.\n* Delete unused includes.\n* Guard ios-only includes with a build flag.\n\nTest: ./components_unittests --gtest_filter=*URLBlocklistManagerTest*\nBug: b:339377401\n\nChange-Id: Ifff431de8f0668899e317fb92c13e49baedaf2da\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5573574\nCommit-Queue: Emmanuel Arias Soto <eariassoto@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Roland Bock <rbock@google.com>\nCr-Commit-Position: refs/heads/main@{#1306718}",
    "author": "Emmanuel Arias Soto",
    "author_email": "eariassoto@google.com",
    "date": "2024-05-28T14:19:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5ee4b4ac657169264e2139c955fe2beeb71062b5"
  },
  {
    "id": 429,
    "message": "Make rust_bindgen_root a configurable argument.\n\nUpstream doesn't provide rust-toolchain for all platforms;\nFor them, using the system's bindgen is an alternative to support.\n\nBug: 342705679\nChange-Id: I53b0b216403ce4f6002ed55713bd820defc39034\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569698\nReviewed-by: Adrian Taylor <adetaylor@chromium.org>\nReviewed-by: danakj <danakj@chromium.org>\nCommit-Queue: Nico Weber <thakis@chromium.org>\nReviewed-by: Nico Weber <thakis@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306717}",
    "author": "Nathan Pratta Teodosio",
    "author_email": "nathan.teodosio@canonical.com",
    "date": "2024-05-28T14:18:49+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/84cbfd44bbc212fa80ecedad9439f20c40c9cb85"
  },
  {
    "id": 430,
    "message": "Roll Perfetto Trace Processor Linux from 480fd109dc7f to e02750a8e14f\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/480fd109dc7f..e02750a8e14f\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-trace-processor-linux-chromium\nPlease CC chrometto-team@google.com,perfetto-bugs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: perfetto-bugs@google.com\nChange-Id: I0f183c7a1592acb5c267e509a384c3e7fa9f2a15\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5575209\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1306716}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-05-28T14:16:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/888d897d972ff83858c048e78510deb6593057f4"
  },
  {
    "id": 431,
    "message": "tools/cr: remove references to goma\n\ngoma is maintained anymore.\n\nBug: 41489832\nChange-Id: Ic61c44b1929bba27f6799d6f7f052f54022fdb42\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5568440\nReviewed-by: Nico Weber <thakis@chromium.org>\nCommit-Queue: Nico Weber <thakis@chromium.org>\nAuto-Submit: Takuto Ikuta <tikuta@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1306715}",
    "author": "Takuto Ikuta",
    "author_email": "tikuta@chromium.org",
    "date": "2024-05-28T14:15:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/669ec93b001af9afb19e39a7d491ee6482b105ff"
  },
  {
    "id": 432,
    "message": "policy: Refactor URLBlocklist::FilterTakesPrecedence\n\n* This function does not use internal state and can be moved to\n  internal linkage.\n* No changes to the original implementation.\n* Added documentation.\n\nTest: ./components_unittests --gtest_filter=*URLBlocklistManagerTest*\nBug: b:339377401\nChange-Id: Ie4d9152060fec4bcd34c0380ce7de7d1f1b8bddc\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5569199\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Roland Bock <rbock@google.com>\nCommit-Queue: Emmanuel Arias Soto <eariassoto@google.com>\nCr-Commit-Position: refs/heads/main@{#1306714}",
    "author": "Emmanuel Arias Soto",
    "author_email": "eariassoto@google.com",
    "date": "2024-05-28T14:09:51+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/124e2f65ca0946fa18ee92e818f503c83be28508"
  },
  {
    "id": 435,
    "message": "[Perf] Expand crossbench runs\n\nAdd crossbench-based Speedometer 3 to additional desktop platforms\n(one platform per OS type). Add JetStream 2 and MotionMark 1.3 to\nlinux-perf-fyi.\n\nBug: b/338629299, b/338630226\nChange-Id: Ie5fcc6729cfffc7ccd23c1c821117ef566299d83\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5639636\nCommit-Queue: John Chen <johnchen@chromium.org>\nReviewed-by: Jeff Yoon <jeffyoon@google.com>\nCr-Commit-Position: refs/heads/main@{#1317541}",
    "author": "John Chen",
    "author_email": "johnchen@chromium.org",
    "date": "2024-06-20T20:05:16+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a54ab353d8138ffaf4c48b122e15e7132b732eac"
  },
  {
    "id": 436,
    "message": "[visited-url-ranking] Visited URL Action UMA metric\n\nNO_IFTTT=New enumeration\n\nBug: 345098154\nChange-Id: I5a75467b792e6082d56686815efead927115d11b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5632985\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Tibor Goldschwendt <tiborg@chromium.org>\nReviewed-by: Robert Kaplow <rkaplow@chromium.org>\nReviewed-by: Siddhartha S <ssid@chromium.org>\nCommit-Queue: Roman Arora <romanarora@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317540}",
    "author": "Roman Arora",
    "author_email": "romanarora@chromium.org",
    "date": "2024-06-20T20:04:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e1ddc0446af0fcc40ba8bbf5720fc282eabadd9f"
  },
  {
    "id": 437,
    "message": "Revert \"Do not trigger win11-arm64-blink-rel in Wpt Importer\"\n\nThis reverts commit 6387cf915fc0e35058d5697b4d63bf0c77280062.\n\nReason for revert: The issue on the bot is resolved,\nrebaseline for win11-arm64 so that the builder can\nbe added back to CQ.\n\nOriginal change's description:\n> Do not trigger win11-arm64-blink-rel in Wpt Importer\n>\n> This builder seems to be extremely unhappy. It takes more than\n> 3.5 hours to finish. And Wpt importer runs all timed out due to this.\n> Skip it for now.\n>\n> Bug: 342424178\n> Change-Id: I6faf045edb02b55ba5d34a38b39a74049bee2792\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5579514\n> Commit-Queue: Jonathan Lee <jonathanjlee@google.com>\n> Auto-Submit: Weizhong Xia <weizhong@google.com>\n> Reviewed-by: Jonathan Lee <jonathanjlee@google.com>\n> Cr-Commit-Position: refs/heads/main@{#1307196}\n\nCq-Include-Trybots: luci.chromium.try:win11-arm64-blink-rel\n\nBug: 342424178\nChange-Id: I299d0b4f76fbf72cebe40057bd2dae7197a32d90\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5637176\nAuto-Submit: Weizhong Xia <weizhong@google.com>\nReviewed-by: Kuan Huang <kuanhuang@chromium.org>\nCommit-Queue: Kuan Huang <kuanhuang@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317539}",
    "author": "Weizhong Xia",
    "author_email": "weizhong@google.com",
    "date": "2024-06-20T20:04:02+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/79156178efb9a63744375c84520913d3de414258"
  },
  {
    "id": 438,
    "message": "media/chromeos/VDPipeline: switch to FrameResourceConverter\n\nInstead of using MailboxVideoFrameConverter, this changes\nVideoDecoderPipeline to use the FrameResourceConverter base class. This\nunblocks using other implementations of FrameResourceConverter.\n\nChanging the MailboxVideoFrameConverter unit tests to store an\nstd::unique_ptr<FrameResourceConverter> instead of a\nstd::unique_ptr<MailboxVideoFrameeConverter> unblocks deleting the\nMailboxVideoFrameConverter `default_delete` struct.\n\nNote: the change to mojo_stable_video_decoder.cc was needed because\nthere was a type definition that had previously been transitively\nincluded, but is not anymore.\n\nBug: b:332382416\nChange-Id: I8b19117eafa285c316b7e28c5376d5d0210e215c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5631412\nReviewed-by: Andres Calderon Jaramillo <andrescj@chromium.org>\nCommit-Queue: Nathan Hebert <nhebert@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317538}",
    "author": "Nathan Hebert",
    "author_email": "nhebert@chromium.org",
    "date": "2024-06-20T20:03:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c81f89850ac1a762e9624a94d8d74b87b99f906d"
  },
  {
    "id": 439,
    "message": "[Fuchsia] Disable OopPixelTest.DrawGainmapImageFiltering\n\nCq-Include-Trybots: luci.chromium.try:fuchsia-fyi-arm64-dbg\nChange-Id: Id3e25ccba4cf429c6324c31e6e576758cd143c62\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5645824\nReviewed-by: Zijie He <zijiehe@google.com>\nCommit-Queue: Guocheng Wei <guochengwei@google.com>\nReviewed-by: David Song <wintermelons@google.com>\nCr-Commit-Position: refs/heads/main@{#1317537}",
    "author": "Guocheng Wei",
    "author_email": "guochengwei@google.com",
    "date": "2024-06-20T20:01:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/53e631d089c24026f2d1418f1965e46bd7ab12b7"
  },
  {
    "id": 440,
    "message": "[infra] Update the active milestones.\n\nM118 and M124 projects have been turned down because they are no longer\nnecessary.\n\nChange-Id: I7fed1d3934f451e1362c02b2a29fcdad612471d3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5646345\nAuto-Submit: Garrett Beaty <gbeaty@google.com>\nReviewed-by: Struan Shrimpton <sshrimp@google.com>\nCommit-Queue: Struan Shrimpton <sshrimp@google.com>\nCr-Commit-Position: refs/heads/main@{#1317536}",
    "author": "Garrett Beaty",
    "author_email": "gbeaty@chromium.org",
    "date": "2024-06-20T19:58:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/83dd2ba350a38bdb626f1cef390f6c989108274b"
  },
  {
    "id": 441,
    "message": "Roll Chrome Mac Arm PGO Profile\n\nRoll Chrome Mac Arm PGO profile from chrome-mac-arm-main-1718899006-c3ed1fa0dbecdbedf391d05091d58f85aa404ba0-1d398d9bdb36e616edf2101e34525f62689d41fc.profdata to chrome-mac-arm-main-1718906365-879e0861c9771052950e66b9c7a33c6f35e15c4c-e824e9e6b6bdfaf0696c9e5e7c4a8f05c8275989.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-mac-arm-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:mac-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: Id33eaa710d52e1d1289a12ad69c23fd8334104c5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5646071\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1317535}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-20T19:56:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/715a8cb8bdeccdd87677012dc8eb5e155ec8bb54"
  },
  {
    "id": 454,
    "message": "[Read Anything] Add Atkinson Hyperlegible font to reading mode.\n\nBug: 1266555\nChange-Id: Icbfeb3dfa62bc045a434db1849f9f10c52d2b015\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5634164\nCommit-Queue: Abigail Klein <abigailbklein@google.com>\nReviewed-by: Kristi Saney <kristislee@google.com>\nCr-Commit-Position: refs/heads/main@{#1317522}",
    "author": "Abigail Klein",
    "author_email": "abigailbklein@google.com",
    "date": "2024-06-20T19:36:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3c7ee695ad7f33c64e3e4cc7066918836f06c74e"
  },
  {
    "id": 442,
    "message": "unbundle: add missing dav1d targets\n\nfixes \"ERROR Unresolved dependencies.\n//third_party/crabbyavif:crabbyavif_dav1d_bindings(//build/toolchain/linux/unbundle:default)\n  needs //third_party/dav1d:dav1d_config(//build/toolchain/linux/unbundle:default)\n//third_party/crabbyavif:crabbyavif_dav1d_bindings(//build/toolchain/linux/unbundle:default)\n  needs //third_party/dav1d:dav1d_headers(//build/toolchain/linux/unbundle:default)\"\n\nChange-Id: I85442e5fb67a804985354570fba453cc619c83d7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5642761\nReviewed-by: Lei Zhang <thestig@chromium.org>\nReviewed-by: Thomas Anderson <thomasanderson@chromium.org>\nCommit-Queue: Thomas Anderson <thomasanderson@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317534}",
    "author": "lauren n. liberda",
    "author_email": "lauren@selfisekai.rocks",
    "date": "2024-06-20T19:55:57+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/89dcd2d419755421290f85e32617acabdd81cac1"
  },
  {
    "id": 443,
    "message": "utr: add target_cpu to android builders\n\nTarget_os and target_cpu are currently assumed based on the bot the\nbuilder is assigned to for many builders. We are working on enforcing\nthese to be explicitly set. This should make those implicit args\nexplicit for this set of builders\n\nBug: 41492686\nChange-Id: I6a9a5aa51ce06c26082b8a2129f060620e96c007\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5639539\nReviewed-by: Haiyang Pan <hypan@google.com>\nCommit-Queue: Struan Shrimpton <sshrimp@google.com>\nCr-Commit-Position: refs/heads/main@{#1317533}",
    "author": "Struan Shrimpton",
    "author_email": "sshrimp@google.com",
    "date": "2024-06-20T19:54:36+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0d8dfb210d471425b6ddad9eec086f8ea0688ea4"
  },
  {
    "id": 444,
    "message": "unbundle: add missing icu target\n\nsolves \"ERROR Unresolved dependencies.\n//chrome/browser/ui/lens:unit_tests(//build/toolchain/linux/unbundle:default)\n  needs //third_party/icu:icuuc_public(//build/toolchain/linux/unbundle:default)\"\n\nChange-Id: I216415b1437f950a34b250e47086983212c1b43a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5643371\nReviewed-by: Lei Zhang <thestig@chromium.org>\nReviewed-by: Thomas Anderson <thomasanderson@chromium.org>\nCommit-Queue: Thomas Anderson <thomasanderson@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317532}",
    "author": "lauren n. liberda",
    "author_email": "lauren@selfisekai.rocks",
    "date": "2024-06-20T19:49:48+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cfab7038d62d51457e50274c9dd543cfd7523e80"
  },
  {
    "id": 445,
    "message": "Roll Chrome Android ARM64 PGO Profile\n\nRoll Chrome Android ARM64 PGO profile from chrome-android64-main-1718891786-f0d0d700e96f42ed787663252663579bdfde37ef-b7dc321164bfac1119cc8bbff789bddd996e2961.profdata to chrome-android64-main-1718899006-8004bc27f89226bf40e5f2ac2672cf807b9a539b-1d398d9bdb36e616edf2101e34525f62689d41fc.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-android-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: Ib8fdcd18872a9704279374d4e6ac960484eb2042\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5645336\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1317531}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-20T19:48:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2f858b2613b9de0da67940924511193f97cdea62"
  },
  {
    "id": 446,
    "message": "Roll ANGLE from 0d6869c2d9b3 to e30871281e34 (2 revisions)\n\nhttps://chromium.googlesource.com/angle/angle.git/+log/0d6869c2d9b3..e30871281e34\n\n2024-06-20 syoussefi@chromium.org Tighten FixedVector access asserts\n2024-06-20 donghwan.yu@lge.com GCC: Define MemoryAllocInfoMapKey's hash before using in unordered_map\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/angle-chromium-autoroll\nPlease CC angle-team@google.com,ynovikov@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in ANGLE: https://bugs.chromium.org/p/angleproject/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86\nBug: chromium:40565911\nTbr: ynovikov@google.com\nChange-Id: I28abe35126464c1560cbb20bc98661b6f05fd46b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5645335\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1317530}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-20T19:47:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8b44045f714071c38fe9b1a8f2fba182120d4a10"
  },
  {
    "id": 447,
    "message": "[lacros skew tests] Refresh skew tests for M128\n\nThis CL updates the ash version ['128.0.6548.0'] for Lacros version skew testing.\nThis cl only affect linux-lacros config builders like\nlinux-lacros-tester-rel, linux-lacros-rel.\nThis cl will certainly NOT affect Lacros on-device builders\n(lacros-amd64-generic-rel, lacros-amd64-generic-chrome-skylab,\netc) or any other platforms.\n\nIf this CL caused regressions, please revert and stop the autoroller\nat https://luci-scheduler.appspot.com/jobs/chrome/lacros-version-skew-roller\nAlso please file a bug to OS>LaCrOS>Partner, and CC kuanhuang@chromium.org.\n\nR=rubber-stamper@appspot.gserviceaccount.com\n\nBug: None\nChange-Id: I4717001b0564ff2bc868912d06734aa6e19649fb\nRequires-Testing: True\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5645983\nBot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nAuto-Submit: chrome-weblayer-builder <chrome-weblayer-builder@chops-service-accounts.iam.gserviceaccount.com>\nCommit-Queue: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1317529}",
    "author": "chrome-weblayer-builder",
    "author_email": "chrome-weblayer-builder@chops-service-accounts.iam.gserviceaccount.com",
    "date": "2024-06-20T19:46:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/19d1b42f744ed84dd4c08b1f53ce19dd0a9b347f"
  },
  {
    "id": 484,
    "message": "Pass in a fake manifest when on-device model is overridden\n\nBug: 346672139\nChange-Id: Iccae5649be462f5f2c5e8e18bae6ae18eea31fed\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5624531\nCommit-Queue: Sophie Chang <sophiechang@chromium.org>\nReviewed-by: Marco Georgaklis <mgeorgaklis@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313832}",
    "author": "Sophie Chang",
    "author_email": "sophiechang@chromium.org",
    "date": "2024-06-12T06:08:01+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0887e7b331effc3216084b94cccb9020af537d8d"
  },
  {
    "id": 448,
    "message": "`span`-ify `mojo/.../data_pipe.h`: Migrate `services/tracing`.\n\n`span`-ification: `mojo::DataPipe[Consumer|Producer]Handle`: Migrating\ncallers.  This CL `span`-ifies the callers of the APIs in\n`mojo/public/cpp/system/data_pipe.h`:\n\n* `BeginReadData` and `BeginWriteData` return (via an out parameter)\n  a `span` instead of returning a ptr+size pair.\n* `ReadData` and `WriteData` consume a `span` instead of consuming\n  a ptr+size pair.\n\nFor more details see the design note here:\nhttps://docs.google.com/document/d/1c4NKpXwpQ9MKK1SbJ4C6MvhXI8-KJZ4jq7N4VHTHJoI/edit?usp=sharing\n\nThis CL was uploaded by git cl split.\n\nR=skyostil@chromium.org\n\nBug: 40284755\nChange-Id: I457c0d678de7faf7f706043f101fc3af8652443b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5618992\nReviewed-by: Alexander Timin <altimin@chromium.org>\nAuto-Submit: Łukasz Anforowicz <lukasza@chromium.org>\nCommit-Queue: Łukasz Anforowicz <lukasza@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317528}",
    "author": "Lukasz Anforowicz",
    "author_email": "lukasza@chromium.org",
    "date": "2024-06-20T19:46:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f6344efb8ea9051f42b8e1d788ce54317e324d72"
  },
  {
    "id": 449,
    "message": "[Compose] Disable nudging on unspecified hints.\n\nThis behavior can cause proactive nudges to be shown on sites that\nviolate policy.  Therefore we turn the sever-side config into an\nallowlist.\n\nBug: b/347341350\nChange-Id: I2188f0a5a4a407852410d781c4f141ed06c38cfb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5631419\nReviewed-by: Trevor Perrier <perrier@chromium.org>\nCommit-Queue: Justin DeWitt <dewittj@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317527}",
    "author": "Justin DeWitt",
    "author_email": "dewittj@chromium.org",
    "date": "2024-06-20T19:45:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2a4a400a5249a4843d000463d9c2fae837e4a74f"
  },
  {
    "id": 450,
    "message": "[Parcel Tracking] Handle unset estimated delivery dates\n\nThe estimated delivery date from the server may be unset, represented\nas a zero value in the proto. This was previously parsed into a valid,\nalbeit null, 'zero time'. When this happened, the user would see an\nestimated delivery date of the Windows epoch, either December 31st or\nJanuary 1st - although thankfully we don't show the year in the UX!\n\nThis CL fixes the bug by:\n\n1. Making estimated_delivery_date a std::optional<base::Time>\n2. Handling a missing estimated delivery date inside the\n   parcel_tracking_view.mm for both in-flight and delivered parcels.\n    i. In-flight parcels are displayed as if in the 'new' state.\n    ii. Finished parcels display just \"Delivered\" with no date.\n\nA view unittest is also added for the functionality. Preferably we\nwould test only the mediator and move the business logic there, but\nthat would be a substantial refactor and so is left for the future.\n\nBug: 330351789\nChange-Id: I6bd971b289a0a1979e90668fd12ee421889ae234\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5630146\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Matthew Jones <mdjones@chromium.org>\nReviewed-by: Chris Lu <thegreenfrog@chromium.org>\nReviewed-by: Slobodan Pejic <slobodan@chromium.org>\nCommit-Queue: Stephen McGruer <smcgruer@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317526}",
    "author": "Stephen McGruer",
    "author_email": "smcgruer@chromium.org",
    "date": "2024-06-20T19:45:01+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6c1e95faceec795cbc5dbf7b1e56d1cfe0d79a36"
  },
  {
    "id": 451,
    "message": "Roll Catapult from 904d293e09a3 to c8eadcacfbba (1 revision)\n\nhttps://chromium.googlesource.com/catapult.git/+log/904d293e09a3..c8eadcacfbba\n\n2024-06-20 eduardoyap@google.com update chromeperf deployment for pinpoint\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/catapult-autoroll\nPlease CC chrome-speed-operations@google.com,dberris@chromium.org,eduardoyap@google.com,jbudorick@chromium.org,johnchen@chromium.org,sullivan@chromium.org,wenbinzhang@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nBug: None\nTbr: eduardoyap@google.com\nChange-Id: I33b73b97b07147390da0ee431dea706b14942560\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5646046\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1317525}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-20T19:44:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/29a8d99cc3ca4d02db68ec2c49971e47696610f2"
  },
  {
    "id": 452,
    "message": "Tidy some rules in PRESUBMIT.py\n\n- Add some separators to a few lists to improve readability.\n- Do more sorting.\n\nChange-Id: I40716aa6d42db12ab2fd8981b5b44104720183af\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5644761\nReviewed-by: Dirk Pranke <dpranke@google.com>\nCommit-Queue: Lei Zhang <thestig@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317524}",
    "author": "Lei Zhang",
    "author_email": "thestig@chromium.org",
    "date": "2024-06-20T19:42:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1f9d9ec46b7aa6c2d543a09e53ba507d27bddb2f"
  },
  {
    "id": 453,
    "message": "Roll V8 from eae24a5d4958 to 780fd425b6f5 (10 revisions)\n\nhttps://chromium.googlesource.com/v8/v8.git/+log/eae24a5d4958..780fd425b6f5\n\n2024-06-20 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Version 12.8.119\n2024-06-20 verwaest@chromium.org Reland \"[compiler] Ensure nested functions always have higher ID than the parent\"\n2024-06-20 dmercadier@chromium.org [turboshaft] Maglev-to-ts: remove assertOptimized in inner-function.js\n2024-06-20 mfarazma@redhat.com [turboshaft] Fix LoadStackArgument on big endian\n2024-06-20 dmercadier@chromium.org [turboshaft] Maglev-to-ts: support TaggedNotEqual\n2024-06-20 clemensb@chromium.org [liftoff] Move two frequenly called methods to the header\n2024-06-20 dmercadier@chromium.org [turboshaft] Collect more feedback in dataview.js test\n2024-06-20 dmercadier@chromium.org [turboshaft] Maglev-to-ts: support TargetType::Any in CallForwardVarargs\n2024-06-20 dmercadier@chromium.org [turboshaft] Maglev-to-ts: properly retag exception phis\n2024-06-20 clemensb@chromium.org [liftoff] Allow constant memory and table indexes\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/v8-chromium-autoroll\nPlease CC liviurau@google.com,machenbach@google.com,v8-waterfall-gardener@grotations.appspotmail.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-blink-rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:dawn-linux-x64-deps-rel\nChange-Id: I6c088ab8ce852b6c8073e941b6d68acc56e24b00\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5645815\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1317523}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-20T19:36:56+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6e4bc3849bab4f527253905df7a2ac7bd09ea620"
  },
  {
    "id": 491,
    "message": "Roll Dawn from 5ddf4e5b6d97 to 6f5358203435 (25 revisions)\n\nhttps://dawn.googlesource.com/dawn.git/+log/5ddf4e5b6d97..6f5358203435\n\n2024-06-11 jaswant.panchumarti@kitware.com Add struct keyword before structures in new callback function pointers\n2024-06-11 senorblanco@chromium.org OpenGLES: enable or suppress Buffer and Texture end2end tests.\n2024-06-11 chrome-branch-day@chops-service-accounts.iam.gserviceaccount.com Activate dawn M127\n2024-06-11 dawn-automated-expectations@chops-service-accounts.iam.gserviceaccount.com Roll third_party/webgpu-cts/ 5fb5afd13..412878330 (11 commits)\n2024-06-11 jrprice@google.com [msl] Emit read_write access for storage textures\n2024-06-11 jrprice@google.com [msl] Add polyfill for textureDimensions()\n2024-06-11 jrprice@google.com [msl] Add member function call validation tests\n2024-06-11 jrprice@google.com [msl] Handle textureSample builtins\n2024-06-11 jrprice@google.com [msl] Add MSL subclass of MemberBuiltinCall\n2024-06-11 jrprice@google.com [ir] Add base MemberBuiltinCall instruction class\n2024-06-11 jrprice@google.com [intrinsics] Consider @member_function in Lookup*\n2024-06-11 jrprice@google.com [intrinsics] Add @member_function attribute\n2024-06-11 dawn-autoroll@skia-public.iam.gserviceaccount.com Roll DirectX Shader Compiler from 4b7993c78be1 to 0b9acdb75e17 (1 revision)\n2024-06-11 rharrison@chromium.org [tint][ir][fuzz] Explictly depend on ir.proto in fuzz.proto\n2024-06-11 senorblanco@chromium.org OpenGL: clean up / enable e2e texture init tests.\n2024-06-11 dsinclair@chromium.org [hlsl] Emit empty functions in HLSL IR.\n2024-06-11 lehoangquyen@chromium.org dawn: Allow internal extensions when creating shader module.\n2024-06-11 dsinclair@chromium.org Enable building without the proto library.\n2024-06-11 dawn-autoroll@skia-public.iam.gserviceaccount.com Roll Depot Tools from 927e02b156cb to e30d8fac3437 (1 revision)\n2024-06-11 senorblanco@chromium.org OpenGLES: enable CopyTextureForBrowser tests.\n2024-06-11 dsinclair@chromium.org [hlsl] Stub out the HLSL IR backend.\n2024-06-11 dawn-autoroll@skia-public.iam.gserviceaccount.com Roll SwiftShader from 2ff3212615da to 085997ddb08b (1 revision)\n2024-06-11 enga@chromium.org Fix TSAN errors in WaitAny\n2024-06-11 dawn-autoroll@skia-public.iam.gserviceaccount.com Roll ANGLE from bb908741db33 to 81452425d73f (11 revisions)\n2024-06-11 dsinclair@chromium.org Remove ir_to_program include from MSL printer.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/dawn-chromium-autoroll\nPlease CC cwallez@google.com,senorblanco@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Dawn: https://bugs.chromium.org/p/dawn/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:dawn-android-arm-deps-rel;luci.chromium.try:dawn-android-arm64-deps-rel;luci.chromium.try:dawn-linux-x64-deps-rel;luci.chromium.try:dawn-mac-x64-deps-rel;luci.chromium.try:dawn-mac-arm64-deps-rel;luci.chromium.try:dawn-win10-x64-deps-rel;luci.chromium.try:dawn-win10-x86-deps-rel;luci.chromium.try:dawn-win11-arm64-deps-rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:gpu-fyi-cq-android-arm64\nBug: chromium:345541614,chromium:346356622,chromium:346377856,chromium:42240662,chromium:42251016,chromium:42251303,chromium:42251304\nTbr: senorblanco@google.com\nInclude-Ci-Only-Tests: true\nChange-Id: Ia2d71726313f01bad1a7d95c3b355630cac53d34\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5621019\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313825}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-12T05:32:34+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/06f5c6df850f19e2a58c78d22b2d05316a8f3f8d"
  },
  {
    "id": 455,
    "message": "Remove unnecessary tests involving WebContents replacement\n\nTabStripModel::DiscardWebContentsAt is only used for WebContents\ndiscarding.\n\nThis CL removes tests for session restore and fullscreen controller\nthat use this API to test behavior when a Tab is replaced, which\nused to occur in the original prerendering impl.\n\nThe prerender implementation has since moved to NoStatePrefetch\n(which does not replace the contents) and these tests are no longer\nrelevant.\n\nBug: 347770670\nChange-Id: I462ddec0919da31114838874a04f96effcc29426\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5639419\nReviewed-by: Lingqi Chi <lingqi@chromium.org>\nReviewed-by: Lei Zhang <thestig@chromium.org>\nCommit-Queue: Thomas Lukaszewicz <tluk@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317521}",
    "author": "Thomas Lukaszewicz",
    "author_email": "tluk@chromium.org",
    "date": "2024-06-20T19:35:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8ece4221dd210c257b3db20627f2fc49c24243d8"
  },
  {
    "id": 456,
    "message": "[Hub] Remove isHubEnabled from tests\n\nThe ANDROID_HUB flag is removed and isHubEnabled now defaults to true.\n\nThis CL removes isHubEnabled from all tests and test infrastructure. As\na side effect a number of classes/lines/methods become unreachable and\nare removed.\n\nBug: 346852431\nChange-Id: I8cc346e1c61564a1fee680128a95f82ffef90e99\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5644332\nReviewed-by: Henrique Nakashima <hnakashima@chromium.org>\nCommit-Queue: Calder Kitagawa <ckitagawa@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317520}",
    "author": "Calder Kitagawa",
    "author_email": "ckitagawa@chromium.org",
    "date": "2024-06-20T19:35:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/21c748cb67ff0a0928f875f746e9a1149121aecc"
  },
  {
    "id": 457,
    "message": "Ref-counting of `SubresourceSignedExchangeURLLoaderFactory::orb_state_`.\n\nBug: 345261068\nChange-Id: I04eae2126ba041d41f44eb4268b4f302f352defd\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5608334\nAuto-Submit: Łukasz Anforowicz <lukasza@chromium.org>\nReviewed-by: mmenke <mmenke@chromium.org>\nCommit-Queue: Łukasz Anforowicz <lukasza@chromium.org>\nCommit-Queue: mmenke <mmenke@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317519}",
    "author": "Lukasz Anforowicz",
    "author_email": "lukasza@chromium.org",
    "date": "2024-06-20T19:32:58+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d9105883c329d6c66e5e2afa516c0899079b28e8"
  },
  {
    "id": 458,
    "message": "[WebAuthn] Add gpm locked pin sheet\n\nThe sheet has placeholder strings and illustrations for now until\nthe mocks are finalised. Adds displaying it in correct moments.\n\nBug: 40274370\nChange-Id: I29a297a5210ee2bc76e3811ec6b91a5d1cbc8e40\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5639745\nReviewed-by: Ken Buchanan <kenrb@chromium.org>\nCommit-Queue: Rafał Godlewski <rgod@google.com>\nCr-Commit-Position: refs/heads/main@{#1317518}",
    "author": "rgod",
    "author_email": "rgod@google.com",
    "date": "2024-06-20T19:29:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d4e8878a00f0ccb173dbf7d80b4823ef2eef6598"
  },
  {
    "id": 459,
    "message": "Do not allow linux arm for xnnpack tflite\n\nBug: 348117454\nChange-Id: Ia28240a8636f60066ea93d890796cb7b98eb5d06\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5643875\nReviewed-by: Robert Ogden <robertogden@chromium.org>\nCommit-Queue: Sophie Chang <sophiechang@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317517}",
    "author": "Sophie Chang",
    "author_email": "sophiechang@chromium.org",
    "date": "2024-06-20T19:28:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c7d4577e86beaebfe2d5515f66149520fa91f8bb"
  },
  {
    "id": 460,
    "message": "snap-groups: Update caption button phantom bounds (4/N)\n\nThis CL updates the snap phantom bounds for the caption button.\n\nDemo: https://buganizer.corp.google.com/issues/346624805#comment17\n\nTest: *FrameSizeButtonTest*\nBug: b/346624805\nChange-Id: Iebbd0e39943ac52f0071765c059b886f14dcca7f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5640633\nReviewed-by: Ahmed Fakhry <afakhry@chromium.org>\nCommit-Queue: Sophie Wen <sophiewen@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1317516}",
    "author": "Sophie Wen",
    "author_email": "sophiewen@google.com",
    "date": "2024-06-20T19:27:36+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/68b8f2fa4fd07372acbe66fcf008e3abcaeb3c7b"
  },
  {
    "id": 461,
    "message": "Roll src/third_party/libaom/source/libaom/ 49c02efb6..afedaf9da (16 commits)\n\nhttps://aomedia.googlesource.com/aom.git/+log/49c02efb61e1..afedaf9da5a1\n\n$ git log 49c02efb6..afedaf9da --date=short --no-merges --format='%ad %ae %s'\n2024-06-17 jonathan.wright Add Arm Neon I8MM implementation of aom_scaled_2d\n2024-06-15 jonathan.wright Add Arm Neon DotProd implementation of aom_scaled_2d\n2024-06-11 jonathan.wright Refactor and optimize aom_scaled_2d_neon\n2024-06-18 jonathan.wright Move aom_scaled_2d_neon to aom_dsp/arm\n2024-06-11 jonathan.wright Add unit tests for aom_scaled_2d_(c|ssse3|neon)\n2024-06-13 jonathan.wright Halve filter values in Armv8.0 Neon convolve8 filter functions\n2024-06-18 jzern add av1/encoder/blockiness.h\n2024-06-14 jzern comp_avg_pred_test.cc: fix class names\n2024-06-14 jzern move comp_avg_pred_test.h contents to cc\n2024-06-18 jzern Revert \"av1_quantize: use optimized aom_quantize_b_helper()\"\n2024-06-14 jzern rename simd_cmp_impl.h to simd_cmp_impl.inc\n2024-06-12 marpan rtc: Add logic for setting sb_size for MT\n2024-06-14 jzern av1_fwd_txfm2d.c: exclude 1:4/4:1 fns w/CONFIG_REALTIME_ONLY\n2024-06-14 jzern sad{,_av1}.c: exclude 1:4/4:1 fns w/CONFIG_REALTIME_ONLY\n2024-06-14 jzern av1_quantize: use optimized aom_quantize_b_helper()\n2024-06-04 marpan rtc: Speedup for dynamic screen content\n\nCreated with:\n  roll-dep src/third_party/libaom/source/libaom\nR=jzern@google.com\n\nBug: b:307414544\nChange-Id: I8d3289e0c1fd0e052b69e7039aa4d6add2e997a6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5646322\nReviewed-by: Wan-Teh Chang <wtc@google.com>\nCommit-Queue: James Zern <jzern@google.com>\nReviewed-by: James Zern <jzern@google.com>\nCommit-Queue: Marco Paniconi <marpan@google.com>\nCr-Commit-Position: refs/heads/main@{#1317515}",
    "author": "Marco Paniconi",
    "author_email": "marpan@google.com",
    "date": "2024-06-20T19:25:05+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fc632b2c58f409ceb774c2a4df3eaa259fa7cdc6"
  },
  {
    "id": 462,
    "message": "Change USM routing id to use new the new models.\n\nConfig was added in cl/644378516.\n\nBug: b:331457898\nChange-Id: I462ab396ef08d69a22f3c6005d2aa9fbd70fbfb6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5639533\nReviewed-by: Benjamin Zielinski <bzielinski@google.com>\nCommit-Queue: Ahmed Nasr <anasr@google.com>\nCr-Commit-Position: refs/heads/main@{#1317514}",
    "author": "Ahmed Nasr",
    "author_email": "anasr@google.com",
    "date": "2024-06-20T19:24:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2b868aeda6de0696fa8c49742ef9c3847f139b53"
  },
  {
    "id": 475,
    "message": "[cjk-fonts] Update the signature of `GetFallbackFamily`\n\n* Change the mandatory out argument to a ref, and move to the\n  last as per the Google C++ Style Guide[1].\n* Change the return type to `AtomicString`, in order to allow\n  caching the `AtomicString` to avoid hash table lookups in\n  future.\n\nThis patch has no behavior changes.\n\n[1] https://google.github.io/styleguide/cppguide.html#Inputs_and_Outputs\n\nBug: 343302583\nChange-Id: I75046b33540e5a59c9ef9e43f537a99aa63f275a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5622210\nCommit-Queue: Koji Ishii <kojii@chromium.org>\nReviewed-by: Kent Tamura <tkent@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313841}",
    "author": "Koji Ishii",
    "author_email": "kojii@chromium.org",
    "date": "2024-06-12T06:38:02+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1fa0d68d20dd46e3a4b576eb9e0090d0df80925f"
  },
  {
    "id": 463,
    "message": "Roll src/third_party/libvpx/source/libvpx/ a2508b571..253d6365e (11 commits)\n\nhttps://chromium.googlesource.com/webm/libvpx.git/+log/a2508b5711b6..253d6365e3f1\n\n$ git log a2508b571..253d6365e --date=short --no-merges --format='%ad %ae %s'\n2024-06-17 marpan rtc-vp9: Allow scene detection for all speeds\n2024-06-14 jzern set_analyzer_env.sh: remove -fno-strict-aliasing\n2024-06-13 jzern rtcd.pl: add license header to generated files\n2024-06-12 angiebird Add missing header in vp9_firstpass.c\n2024-06-12 angiebird Fix typo of received again\n2024-06-12 angiebird Remove redundant setting of max_layer_depth.\n2024-06-12 angiebird Typo recieved -> received\n2024-06-11 jzern configure: add -c to ASFLAGS for android + AS=clang\n2024-06-11 jzern configure: remove unused NM & RANLIB variables\n2024-06-06 angiebird Move ext_rc_define_gf_group_structure\n2024-06-07 jzern tiny_ssim: mv read error checks closer to assignment\n\nCreated with:\n  roll-dep src/third_party/libvpx/source/libvpx\nR=jzern@google.com\n\nBug: b:308446709, b:346846607\nChange-Id: I87b5d0970704cd42ed16c600e50c7d10063d3a6e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5644752\nReviewed-by: Wan-Teh Chang <wtc@google.com>\nCommit-Queue: James Zern <jzern@google.com>\nReviewed-by: James Zern <jzern@google.com>\nCr-Commit-Position: refs/heads/main@{#1317513}",
    "author": "Marco Paniconi",
    "author_email": "marpan@google.com",
    "date": "2024-06-20T19:23:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/edfc3402d1c657bf13537ac61e4dd5729bc3e254"
  },
  {
    "id": 464,
    "message": "Roll Perfetto Trace Processor Linux from 34e4cd49026b to 4d27a177ff6f\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/34e4cd49026b..4d27a177ff6f\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-trace-processor-linux-chromium\nPlease CC chrometto-team@google.com,perfetto-bugs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: perfetto-bugs@google.com\nChange-Id: I692ab723f62a5dd308a98b6c2b0fdefb5c466407\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5645971\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1317512}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-20T19:23:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0c7e9be613ccb14342efdfdfa9946793282051a8"
  },
  {
    "id": 465,
    "message": "Optimize internal state of FormFetcher.\n\nCombine non_federated_ and non_federated_same_scheme_ to reduce memory\nfootprint.\n\nSide effect of this change is that GetNonFederatedMatches() will\nreturn same scheme matches in the same order as in\nGetAllRelevantMatches() and GetBestMatches().\n\nBug: 343879843\nChange-Id: Iac410fde91af49567584336a8df136153563e1db\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5598786\nCommit-Queue: Kirill Davletkildeev <kirilld@google.com>\nReviewed-by: Vasilii Sukhanov <vasilii@chromium.org>\nReviewed-by: Anna Tsvirchkova <atsvirchkova@google.com>\nCr-Commit-Position: refs/heads/main@{#1313851}",
    "author": "Kirill Davletkildeev",
    "author_email": "kirilld@google.com",
    "date": "2024-06-12T08:17:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/982268fed7ce1e05c6ae6c97865d229d15902af4"
  },
  {
    "id": 466,
    "message": "Move UnmaskAuthMethod into PaymentsAutofillClient\n\nBug: 40937065\nChange-Id: I6a8e4985b9e07803712acd35e02fec49e176bcae\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5622977\nReviewed-by: Olivia Saul <jsaul@google.com>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Jan Keitel <jkeitel@google.com>\nCommit-Queue: Vinny Persky <vinnypersky@google.com>\nCr-Commit-Position: refs/heads/main@{#1313850}",
    "author": "Vinny Persky",
    "author_email": "vinnypersky@google.com",
    "date": "2024-06-12T08:10:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e77a4e0b7bbf5f4c640bdf13ffaa95e0b4a5fe3e"
  },
  {
    "id": 467,
    "message": "[Mahi] Implement `OnLinkClick` in the Magicboost disclaimer view\n\nBug: b/339044721\nChange-Id: I22ed603898e3715620877af57a2f6e5e1ffa39ee\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5622705\nCommit-Queue: Jiaming Cheng <jiamingc@chromium.org>\nReviewed-by: Andre Le <leandre@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313849}",
    "author": "Jiaming Cheng",
    "author_email": "jiamingc@chromium.org",
    "date": "2024-06-12T08:06:53+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8c246c32e7c6c51b2c695a0c22c986aeb5fddb92"
  },
  {
    "id": 468,
    "message": "Add system PackageId PackageType\n\nSystem apps are used in Launcher Ordering.\n\nBug: b/327058986\nChange-Id: I6ba422ff578f804f055e643740f35c64dae5c443\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5619687\nReviewed-by: Tim Sergeant <tsergeant@chromium.org>\nReviewed-by: Sam McNally <sammc@chromium.org>\nCommit-Queue: Sam McNally <sammc@chromium.org>\nAuto-Submit: Joel Hockey <joelhockey@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313848}",
    "author": "Joel Hockey",
    "author_email": "joelhockey@chromium.org",
    "date": "2024-06-12T08:05:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/560b96206c5bac7d120b1218e4235f0eb207266f"
  },
  {
    "id": 469,
    "message": "infra: Use 32 core bots for ci/linux-win-cross-rel\n\nThe builder often fails for OOM.\nhttps://ci.chromium.org/ui/p/chromium/builders/ci/linux-win-cross-rel/930/infra\n\nBug: 343922410\nChange-Id: I8b227853433d67f64686c381d28636ab30a375f3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5622955\nAuto-Submit: Junji Watanabe <jwata@google.com>\nReviewed-by: Takuto Ikuta <tikuta@chromium.org>\nCommit-Queue: Junji Watanabe <jwata@google.com>\nCr-Commit-Position: refs/heads/main@{#1313847}",
    "author": "Junji Watanabe",
    "author_email": "jwata@google.com",
    "date": "2024-06-12T07:55:48+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/62114f568675eb0c3d18307a3f42d9d3b6100ab1"
  },
  {
    "id": 483,
    "message": "Roll Depot Tools from 441cd5c465c1 to 6f180c0a2304 (1 revision)\n\nhttps://chromium.googlesource.com/chromium/tools/depot_tools.git/+log/441cd5c465c1..6f180c0a2304\n\n2024-06-12 richardwa@google.com Use $HOME/.config/depot_tools on linux for .cfg files\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/depot-tools-chromium-autoroll\nPlease CC chops-source-team@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: \nChange-Id: I7e5e14f903fd65e135cbe0c8c494221fb88b9e0d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5624308\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313833}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-12T06:12:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ada75f4c584f5dd943b064c76646b387906d3c38"
  },
  {
    "id": 470,
    "message": "Roll Chrome Win32 PGO Profile\n\nRoll Chrome Win32 PGO profile from chrome-win32-main-1718031540-f4208ed6c14ed676748ba330febe2576831b3c37-131982a2bae5824f2af783d0efa0b9f861fa4ebc.profdata to chrome-win32-main-1718161127-58655325d450af6559b6deb62f2630a178629ffe-9d65339121490aa783690d8806ff255ff1289b25.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win32-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:win-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I221fd9cf999cae18b9c114e4d46b997e3f639c1a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5624809\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313846}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-12T07:55:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/89aef31f61ca9da7cffb88b1bd77ffa21f3bcce0"
  },
  {
    "id": 471,
    "message": "Roll clank/internal/apps from 9677dccda0e2 to 2fb1a3798ec5 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/9677dccda0e2..2fb1a3798ec5\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: \nNo-Try: true\nChange-Id: Ia442b60d9a3fe2f8282afd40eb5ae67a9766d363\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5621445\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313845}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-12T07:05:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1e16f300f0bd05aed7eef84f944871ebd087b3b5"
  },
  {
    "id": 472,
    "message": "Reland \"[E2E] Dynamic safe-area-inset* as browser control scrolls\"\n\nThis reverts commit dbd46654c908ae4709927eaea4c1dd5d0320c1e3.\n\nReason for revert: The original CL failed for browser tests due to DisplayCutoutClientImpl received IPCs are all in different frames. The reland attempts to make the client associate with the Document instead. \n\nOriginal change's description:\n> Revert \"[E2E] Dynamic safe-area-inset* as browser control scrolls\"\n>\n> This reverts commit 8da040ad0865e003d29a97c5a306cc133cc32c2d.\n>\n> Reason for revert:\n> LUCI Bisection has identified this change as the cause of a test failure. See the analysis: https://ci.chromium.org/ui/p/chromium/bisection/test-analysis/b/5482680839831552\n>\n> Sample build with failed test: https://ci.chromium.org/b/8746129781505968577\n> Affected test(s):\n> [ninja://chrome/android:chrome_public_test_apk/org.chromium.chrome.browser.display_cutout.DisplayCutoutTest#testViewportFitDefault](https://ci.chromium.org/ui/test/chromium/ninja:%2F%2Fchrome%2Fandroid:chrome_public_test_apk%2Forg.chromium.chrome.browser.display_cutout.DisplayCutoutTest%23testViewportFitDefault?q=VHash%3A359dcabdb2bd6ffc)\n>\n> If this is a false positive, please report it at http://b.corp.google.com/createIssue?component=1199205&description=Analysis%3A+https%3A%2F%2Fci.chromium.org%2Fui%2Fp%2Fchromium%2Fbisection%2Ftest-analysis%2Fb%2F5482680839831552&format=PLAIN&priority=P3&title=Wrongly+blamed+https%3A%2F%2Fchromium-review.googlesource.com%2Fc%2Fchromium%2Fsrc%2F%2B%2F5543639&type=BUG\n>\n> Original change's description:\n> > [E2E] Dynamic safe-area-inset* as browser control scrolls\n> >\n> > Dynamically change safe-area-inset* css attribute as browser control scrolls.\n> >\n> > This CL implements that the safe-area-inset-bottom will increase as the    browser control scrolls below the display cutout area, keeping the bottom attached item remains in the touchable area. When browser controls scroll back to the screen, the safe-area-inset* will decrease as the viewport is taken over by the browser control.\n> >\n> > Recording and more design details please see the attached bug.\n> >\n> > Bug: 324436581\n> > Change-Id: I039da021d150939513f3da1465b8f88fe95f4842\n> > Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5543639\n> > Reviewed-by: Charles Hager <clhager@google.com>\n> > Reviewed-by: Robert Flack <flackr@chromium.org>\n> > Commit-Queue: Wenyu Fu <wenyufu@chromium.org>\n> > Cr-Commit-Position: refs/heads/main@{#1309385}\n> >\n>\n> Bug: 324436581\n> Change-Id: I9ae376aa6d3441c223944c9e51046bb1a8ce5bd2\n> No-Presubmit: true\n> No-Tree-Checks: true\n> No-Try: true\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5594472\n> Commit-Queue: Wenyu Fu <wenyufu@chromium.org>\n> Reviewed-by: Charles Hager <clhager@google.com>\n> Reviewed-by: Robert Flack <flackr@chromium.org>\n> Cr-Commit-Position: refs/heads/main@{#1310006}\n\nBug: 324436581\nChange-Id: Ie22ba6cb2c8a840f63ebd1988967a77b306e7733\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5597281\nCommit-Queue: Wenyu Fu <wenyufu@chromium.org>\nReviewed-by: Robert Flack <flackr@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313844}",
    "author": "Wenyu Fu",
    "author_email": "wenyufu@chromium.org",
    "date": "2024-06-12T07:05:43+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/685b8d917dfa9a75fcbdb902e84e2518a9006016"
  },
  {
    "id": 473,
    "message": "Roll devtools-internal from 9d7ba6e5cbd5 to f2c6d8312653 (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/9d7ba6e5cbd5..f2c6d8312653\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/devtools/devtools-frontend.git/+log/e77d086ad0337f8d5f6326418f669712a3a0ce2c..8c7d4283a5fdc252e10f2a4ded4c80b87a731d13\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: Ib79defa62a3eab379103f95ee2a9a730288d0e25\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5624810\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313843}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-12T06:44:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/7a3e9aa840ff28ec9b3b14c482a1d0b036ea1d2e"
  },
  {
    "id": 474,
    "message": "[PDF Ink Signatures] Convert input positions to canonical format\n\nEvent positions collected during an ink stroke are provided as screen\ncoordinates, as an offset from the top left corner of an available\nmarking area.  To understand the real position within the page being\nstroked, the position needs to take page orientation, zoom level, and\noffsets into account.\n\nStore the stroke input points in a canonical format, representing the\nposition as CSS pixels within a page in its original orientation.  This\nwould directly match for rendering to a page at 100% zoom which fully\nuses an embedder's viewport.\n\nBug: 335517471\nChange-Id: Ib009c47c9afc3a114f9ce142a7c234c18f7e98be\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5571881\nReviewed-by: Lei Zhang <thestig@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Alan Screen <awscreen@chromium.org>\nReviewed-by: Andy Phan <andyphan@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313842}",
    "author": "Alan Screen",
    "author_email": "awscreen@chromium.org",
    "date": "2024-06-12T06:42:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8c19ddc7cb1e8063680ca0a680c7c4c9c98295a1"
  },
  {
    "id": 476,
    "message": "Roll Perfetto from 9cb4dc8b3bed to 2f21ab4e2a44 (2 revisions)\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/9cb4dc8b3bed..2f21ab4e2a44\n\n2024-06-12 chinglinyu@google.com Merge \"traced_relay: bind to the listening socket from Android init\" into main\n2024-06-12 vaage@google.com Merge changes I63ade06a,I0d042b48 into main\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-chromium-autoroll\nPlease CC perfetto-bugs@google.com,primiano@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-perfetto-rel\nBug: None\nTbr: perfetto-bugs@google.com\nChange-Id: Ib069e69b614088c92792181dfa88c892b44d2f18\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5624535\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313840}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-12T06:35:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3435dd4f2a4d6ec83dd8178db7390a33ed69e6ec"
  },
  {
    "id": 477,
    "message": "Roll src/net/third_party/quiche/src/ 76863f97e..f50a7ee73 (4 commits)\n\nhttps://quiche.googlesource.com/quiche.git/+log/76863f97e957..f50a7ee73f0d\n\n$ git log 76863f97e..f50a7ee73 --date=short --no-merges --format='%ad %ae %s'\n2024-06-11 bnc Set `header_decoding_delay_` before calling methods.\n2024-06-10 ericorth No public description\n2024-06-10 martinduke Update MoQT to draft-04\n2024-06-07 wub Split `QuicConnectionTracer` into two interfaces: - A `QuicPlatformConnectionContext` interface that get notified when a `QuicConnection` becomes active or inactive. - A new smaller `QuicConnectionTracer` that only contains PrintXXX methods.\n\nCreated with:\n  roll-dep src/net/third_party/quiche/src src/third_party/quic_trace/src\n\nChange-Id: I79481530f9e22cc7de239fc7e1c0414cf5a6c926\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5622350\nCommit-Queue: Adam Rice <ricea@chromium.org>\nReviewed-by: Adam Rice <ricea@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313839}",
    "author": "Bence Béky",
    "author_email": "bnc@chromium.org",
    "date": "2024-06-12T06:33:40+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/ceaa1064baf254a1a4797e2104621553e7023873"
  },
  {
    "id": 478,
    "message": "Roll Chrome Win64 PGO Profile\n\nRoll Chrome Win64 PGO profile from chrome-win64-main-1718117710-9b7abb48ee731161c1ef96bb15c150099925b4ee-ef26015e3807236c81d0dbd6c3aaff81dba01f7c.profdata to chrome-win64-main-1718150158-6403bb09a9ba149354a9caff8f062e8c943796c9-db76c843f48e4a74f3d16379246be3cb318b34ba.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:win64-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: Iff10c97effc8ba389ac2a640136a4ab03d7c7398\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5623822\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313838}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-12T06:28:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5ac40bb889e4132e3892f9c36ee28fdddaf06191"
  },
  {
    "id": 479,
    "message": "[PDF Ink Signatures] Add InkStroke interface to get inputs\n\nExpose the ability to get the inputs that were provided into an ink\nstroke.  Requires extension of the stub implementation to retain and\npropagate the input values.  This prepares for allowing InkModule to\ncheck the validity of captured ink input positions in\nhttps://crrev.com/c/5571881.\n\nBug: 335517471\nChange-Id: Ia3f74c4a803efa8566c2baba2bd35008883d2b9c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5622978\nCommit-Queue: Alan Screen <awscreen@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Lei Zhang <thestig@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313837}",
    "author": "Alan Screen",
    "author_email": "awscreen@chromium.org",
    "date": "2024-06-12T06:27:54+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/770a74530b55eef2c54db3005fe253282238fb66"
  },
  {
    "id": 480,
    "message": "Remove HttpStreamFactory::JobController::OnResolveProxyError()\n\nIt's not defined (probably removed at some point).\n\nBug: N/A\nChange-Id: I3a270a340dc3188d02e647011e167185187d16b8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5619640\nReviewed-by: Adam Rice <ricea@chromium.org>\nCommit-Queue: Kenichi Ishibashi <bashi@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313836}",
    "author": "Kenichi Ishibashi",
    "author_email": "bashi@chromium.org",
    "date": "2024-06-12T06:26:02+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3bef8c36bcab97c570723088e8dad92b9c453df9"
  },
  {
    "id": 481,
    "message": "Extend WebShare flag expiry\n\nBug: 346617229\nChange-Id: I42f98ba1a497c633086cff2cff58b898d6d89c6e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5623576\nReviewed-by: Daniel Murphy <dmurph@chromium.org>\nCommit-Queue: Hassan Talat <hatalat@microsoft.com>\nCr-Commit-Position: refs/heads/main@{#1313835}",
    "author": "Hassan Talat",
    "author_email": "hatalat@microsoft.com",
    "date": "2024-06-12T06:20:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3129c05ce30a0fb9f84521e95e7dce35ab8dac14"
  },
  {
    "id": 482,
    "message": "Roll optimization-guide from 84e631720516 to e10a3ab857a9\n\nhttps://chrome-internal.googlesource.com/chrome/components/optimization_guide.git/+log/84e631720516..e10a3ab857a9\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/optimization-guide-chromium\nPlease CC chrome-intelligence-core@google.com,sophiechang@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: I253cec00628e45850fff00d976af5c69cf6e4217\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5621438\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313834}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-12T06:18:47+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/51de850d7261a5847919577e2ccb34d9484acd33"
  },
  {
    "id": 485,
    "message": "Roll clank/internal/apps from 5358bb5f45d6 to 9677dccda0e2 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/5358bb5f45d6..9677dccda0e2\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: \nNo-Try: true\nChange-Id: Ie30b7a05dfd704a7a4f2154dc241ab7a815a9a02\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5624794\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313831}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-12T06:05:47+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/445b8d161e8506966fa25d389a263b271156130d"
  },
  {
    "id": 486,
    "message": "Roll Perfetto Trace Processor Linux from 69fa769cde3a to 9cb4dc8b3bed\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/69fa769cde3a..9cb4dc8b3bed\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-trace-processor-linux-chromium\nPlease CC chrometto-team@google.com,perfetto-bugs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: perfetto-bugs@google.com\nChange-Id: Ibeb86c71c6ae642ad758b559ebfd4388d802e5ae\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5624272\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1313830}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-12T06:02:31+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3116fe3c7aa3712a15e9a0bd35a0e279c6419b2b"
  },
  {
    "id": 487,
    "message": "[FedCM] Show account chooser for non-returning user on approved client\n\nCurrently on button mode, for this particular case of a logged out\nuser signing in as a non-returning user on an RP which is in the list\nof approved clients, we would show disclosure UI without the\ndisclosure text. This patch makes it such that we show the account\nchooser instead because the disclosure UI without the disclosure text\nis unnecessary new UI surface which doesn't convey any additional\ninformation compared to the account chooser.\n\nNote that we show the multi account picker instead of the single account\npicker. This is because when we were showing disclosure UI, the user was\nable to go back and retrieve the original list of accounts by clicking\nthe \"Back\" button. However, the \"back\" button is not available on the\naccount pickers. We prefer a multi account picker compared to having an\nadditional UI surface. The most recently signed in accounts are shown\nat the top of the multi account picker.\n\nThis patch also refactors the state machine. We used to set the state_\nto State::REQUEST_PERMISSION for single account pickers on the bubble\nbut that is confusing. This patch changes it to have state_ be\nState::SINGLE_ACCOUNT_PICKER.\n\nBug: 342557704\nChange-Id: I9435ee5a34cf47567c6a4a16bbffd99ba8eee798\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5578254\nReviewed-by: Yi Gu <yigu@chromium.org>\nCommit-Queue: Zachary Tan <tanzachary@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313829}",
    "author": "Zachary Tan",
    "author_email": "tanzachary@chromium.org",
    "date": "2024-06-12T05:45:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e51a9e0550d6b95679fc64f40f4a094ac6fe2d83"
  },
  {
    "id": 488,
    "message": "arc: Add metric for apk cache hit/miss\n\nArc.AppInstall.CacheHit metric will help us track usefulness of apk\ncache feature and detect any regressions.\n\nBUG=b:323459171\nTEST=./out_Default/Release/ash_components_unittests --gtest_filter='*ReportApkCache*'\n\nChange-Id: I12e8d452a0c574f6acb79763197f9286b2ed5120\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5617673\nReviewed-by: Yuichiro Hanada <yhanada@chromium.org>\nReviewed-by: Luzanne Batoon <batoon@google.com>\nReviewed-by: Jorge Lucangeli Obes <jorgelo@chromium.org>\nCommit-Queue: Muhammad Hasan Khan <mhasank@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313828}",
    "author": "Muhammad Hasan Khan",
    "author_email": "mhasank@google.com",
    "date": "2024-06-12T05:44:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/450ad987d7c2c17208e615b218905a8e315dd3db"
  },
  {
    "id": 489,
    "message": "Refactor zoom propagation\n\nThis recent patch made zoom level a property on WebFrameWidget rather\nthan WebView:\n\nhttps://chromium-review.googlesource.com/c/chromium/src/+/5571113\n\nHowever, it left the propagation routes for zoom information in a bit of\na crazy state. This CL is a pure refactor with no behavioral change,\nintended only to make the propagation mechanism more consistent and\nsane.\n\nAll of the inputs to the zoom factor calculation arrive via IPC to a\nframe widget. Zoom level is handled directly by WebFrameWidget; all\nother factors are delegated to WebView, which propagates them down to\nall WebFrameWidgets contained by the WebView. The WebFrameWidget in turn\nis responsible for propagating zoom factor to the root LocalFrame of the\nwidget, and the root frame is responsible for propagating it to\ndescendant non-root frames.\n\nThis also activates the path for propagating zoom level from embedder to\nprocess-isolated iframe via VisualProperties::zoom_level. Currently,\nthat value is ignored by RenderWidgetHostImpl::GetVisualProperties in\nfavor of the WebContents-level value. After this change, the value set\nby the embedder will be used instead. Currently, the two values are\nalways identical, so this is a pure plumbing change with no behavior\ndifference. A follow-up CL will allow the value set by the embedder to\ndeviate from the WebContents-level value based on CSS zoom.\n\nBug: chromium:329482480\nChange-Id: I18991c9db9f1d1be8ffad7677d6a551ab70495f9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5611483\nReviewed-by: Ken Buchanan <kenrb@chromium.org>\nCommit-Queue: Stefan Zager <szager@chromium.org>\nReviewed-by: Andrey Kosyakov <caseq@chromium.org>\nReviewed-by: Philip Rogers <pdr@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313827}",
    "author": "Stefan Zager",
    "author_email": "szager@chromium.org",
    "date": "2024-06-12T05:42:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/da8a353ba2c90996af18b0308eb95038232b1f15"
  },
  {
    "id": 490,
    "message": "[Extensions c2s] Extend flags expiry milestone to M132\n\nExtend the Click to Script project expiry milestone to M132\n\nClosed: 346617114,346616939, 346617066\nChange-Id: I1edc833c9a8cbae256447fe865925fa35022ec8c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5621868\nCommit-Queue: Emilia Paz <emiliapaz@chromium.org>\nAuto-Submit: Emilia Paz <emiliapaz@chromium.org>\nReviewed-by: Kelvin Jiang <kelvinjiang@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313826}",
    "author": "EmiliaPaz",
    "author_email": "emiliapaz@chromium.org",
    "date": "2024-06-12T05:33:10+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/072b1a84850e9bfaec94a023cc3e7ce651176bff"
  },
  {
    "id": 493,
    "message": "[lensoverlay] Shader should fade out after region focus.\n\nThe cursor transition should only happen when the previous shimmer\ncontroller was NONE / steady state. The circles should fade out in place\notherwise.\n\nDemo: http://shortn/_ohkjjn040P\n\nChange-Id: I505e8e6c0f840ae3e3c07ac7c33d7f2a05ea2f59\nBug: b:345532379\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5622772\nCommit-Queue: Juan Mojica <juanmojica@google.com>\nReviewed-by: Duncan Mercer <mercerd@google.com>\nCr-Commit-Position: refs/heads/main@{#1313823}",
    "author": "Juan Mojica",
    "author_email": "juanmojica@google.com",
    "date": "2024-06-12T05:30:13+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f59c9f1b64a076b75214682240484226963d32fc"
  },
  {
    "id": 494,
    "message": "[manta] include channel info in the Manta request proto\n\n\nBug: b:346427098\nTest: autoninja -C out_brya/Release chrome\nChange-Id: I4c7d7c9866ec75b297a3719a7c8b9f48c9d9da07\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5619951\nReviewed-by: Andrew Moylan <amoylan@google.com>\nCommit-Queue: Xinglong Luan <alanlxl@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1313822}",
    "author": "Xinglong Luan",
    "author_email": "alanlxl@google.com",
    "date": "2024-06-12T05:29:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/321e86795d57216829179e70672f820be9284571"
  },
  {
    "id": 495,
    "message": "CSS highlight pseudos: fix ‘currentColor’ resolution and painting\n\nWhen ‘color’ is set to ‘currentColor’ in highlight pseudos, the color\n(and any other properties that use ‘currentColor’) needs to be taken\nfrom the ‘currentColor’ of the next active highlight overlay below.\n\nOur impl was reworked significantly in CL:5332250 and CL:5368297, and\nthis patch improves on that by:\n\n• updating HighlightPaintingStyle and ResolveColorsFromPreviousLayer\n  to let ‘background-color’ participate in ‘currentColor’ resolution\n\n• fixing ResolveColorsFromPreviousLayer to always read ‘currentColor’,\n  not the property being resolved, from the previous layer\n\n• making PaintHighlightOverlays pass the correct ‘currentColor’ for\n  the purposes of painting ‘text-shadow’\n\n• extending PaintHighlightOverlays to paint ‘background-color’ and\n  ‘text-shadow’ by as little as one part at a time, if affected by\n  ‘currentColor’, while keeping parts merged where possible\n\n• removing a stray LineRelativeRect::Unite call that made ::selection\n  parts fail to clip their decorations to the part\n\nThis patch adds eleven new tests with overlapping highlights and a\nvariety of properties set to ‘currentColor’:\n\n• 001.html tests ‘color’ only (full ::selection)\n• 001a.html tests ‘color’ only (partial ::selection)\n• 002.html tests ‘color’ and ‘background-color’ (full ::selection)\n• 002a.html tests ‘color’ and ‘background-color’ (partial ::selection)\n• 002b.html tests ‘background-color’ only (full ::selection)\n• 003.html tests ‘color’ and ‘text-decoration’ (full ::selection)\n• 003a.html tests ‘color’ and ‘text-decoration’ (partial ::selection)\n• 003b.html tests ‘text-decoration’ only (full ::selection)\n• 004.html tests ‘color’ and ‘text-shadow’ (full ::selection)\n• 004a.html tests ‘color’ and ‘text-shadow’ (partial ::selection)\n• 004b.html tests ‘text-shadow’ only (full ::selection)\n\nBug: 339298411\nChange-Id: I2e42a0d655683e76daf507aee3e34085f5eb080b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5644614\nReviewed-by: Stephen Chenney <schenney@chromium.org>\nCommit-Queue: Delan Azabani <dazabani@igalia.com>\nCr-Commit-Position: refs/heads/main@{#1322551}",
    "author": "Delan Azabani",
    "author_email": "dazabani@igalia.com",
    "date": "2024-07-03T04:16:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6632618e8216b87a594486ec30d920ae6326c750"
  },
  {
    "id": 496,
    "message": "Add UKMs for data URL length and parse time\n\nThis CL adds a UKM metric measuring the data URL length and the duration\nit takes the parse the URL. This will help us understand the\ndistribution of data URL sizes and potential loading optimization\nopportunities.\n\nNew UKM Collection Review: https://docs.google.com/document/d/1nVYsibOCoWLkQKbcBF0fYUpCZuaB3kVCYHJGOd6nPaM/edit?usp=sharing&resourcekey=0-2M3q7wzZEM8x4LzUEbvlWA\n\nBug: 348442535\nChange-Id: Ic465ff70df671fbaba7b43a46be2ed2a1ab59ec0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5656891\nReviewed-by: Takashi Toyoshima <toyoshim@chromium.org>\nCommit-Queue: Nidhi Jaju <nidhijaju@chromium.org>\nReviewed-by: Sun Yueru <yrsun@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322550}",
    "author": "Nidhi Jaju",
    "author_email": "nidhijaju@chromium.org",
    "date": "2024-07-03T04:14:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/59089e36bd07aab03641e8d476e526cbfcfdfc3e"
  },
  {
    "id": 497,
    "message": "Roll Amd64 AFDO from 128.0.6570.0_rc-r1-merged to 128.0.6571.0_rc-r1-merged\n\nThis CL may cause a small binary size increase, roughly proportional\nto how long it's been since our last AFDO profile roll. For larger\nincreases (around or exceeding 100KB), please file go/crostc-bug.\n\nPlease note that, despite rolling to chrome/android, this profile is\nused for both Linux and Android.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/afdo-chromium\nPlease CC c-compiler-chrome@google.com,mobiletc-prebuild@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium Main: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: mobiletc-prebuild@google.com\nChange-Id: I9b9b61caaed350c52a104d06ac158b28d0ebf916\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5673604\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322549}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T04:10:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c91e63db8b1cf6bd3e1ab2a09830096b9f07cee3"
  },
  {
    "id": 498,
    "message": "Roll Depot Tools from d7d82215816c to ca091f0d1678 (1 revision)\n\nhttps://chromium.googlesource.com/chromium/tools/depot_tools.git/+log/d7d82215816c..ca091f0d1678\n\n2024-07-03 ayatane@chromium.org [gerrit_util] Update docstring for new param\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/depot-tools-chromium-autoroll\nPlease CC chops-source-team@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: \nChange-Id: I6ee2d8319310572015902ddcf00ccb0ebae4490d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5672791\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322548}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T04:06:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3fc9245cb63c089c783d729c4fdd955dc4d3fd9a"
  },
  {
    "id": 499,
    "message": "Updating trunk VERSION from 6573.0 to 6574.0\n\n* This is an automated release commit.\n* Do not revert without consulting chrome-pmo@google.com.\nNOAUTOREVERT=true\n\nChange-Id: I2ac8cb20af793dad08ac98575a829f698223ec21\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5671015\nBot-Commit: Chrome Release Bot (LUCI) <chrome-official-brancher@chops-service-accounts.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322547}",
    "author": "Chrome Release Bot (LUCI)",
    "author_email": "chrome-official-brancher@chops-service-accounts.iam.gserviceaccount.com",
    "date": "2024-07-03T04:00:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0199daa6b4cd6896fede1418cc0d4e331791f971"
  },
  {
    "id": 500,
    "message": "WebNN: Destroy command recorder for graph initialization\n\nThis CL releases command recorder and its recorded command resources for\ngraph initialization, such as constants uploading buffers, after\nsubmission to command queue. This change ensures no graph initialization\nresources are kept unnecessarily until the first `dispatch()`\ninvocation. The worst case is that these graph initialization resources\nare kept during the whole `MLGraph` live time, because the apps may only\ncall `compute()`.\n\nBecause the initialization command recorder is not re-used anymore, a\nnew command recorder is created for graph `dispatch()`.\n\nBug: 350292237\nChange-Id: Id8922bc94e2831f043639f266968238ce91b8ac8\nCq-Include-Trybots: luci.chromium.try:win11-blink-rel\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5669072\nReviewed-by: Rafael Cintron <rafael.cintron@microsoft.com>\nCommit-Queue: ningxin hu <ningxin.hu@intel.com>\nCr-Commit-Position: refs/heads/main@{#1322546}",
    "author": "Ningxin Hu",
    "author_email": "ningxin.hu@intel.com",
    "date": "2024-07-03T03:54:12+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/77710f5a66fd6397d9075ad93fe5b1efc26581ae"
  },
  {
    "id": 553,
    "message": "[Nearby Share] Respect is_p2p_supported for WiDi\n\nUses the `is_p2p_supported` property when determining if the WiDi\ninterface is valid.\n\nTests: unittested\nBug: b/345545524\nChange-Id: I8e7204a56cb0cd767757a2c49d7a42e28bca9c88\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606210\nReviewed-by: Brando Socarras <brandosocarras@google.com>\nCommit-Queue: Jack Shira <jackshira@google.com>\nCr-Commit-Position: refs/heads/main@{#1311633}",
    "author": "Jack Shira",
    "author_email": "jackshira@google.com",
    "date": "2024-06-07T00:51:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3c505639fe2cbabf2e4674707418ec3177884b55"
  },
  {
    "id": 501,
    "message": "Roll clank/internal/apps from cf9e3d118ccf to 99ff2d18af3a (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/cf9e3d118ccf..99ff2d18af3a\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: \nNo-Try: true\nChange-Id: Ic14cc73e4be899c03fd61be343f9d532e1456e50\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5674083\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322545}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-03T03:50:02+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2d0fd38ac930433275adfaec6b6d3ffcb8023e07"
  },
  {
    "id": 502,
    "message": "Roll Chrome Mac Arm PGO Profile\n\nRoll Chrome Mac Arm PGO profile from chrome-mac-arm-main-1719964707-dc24bfda35e48423f4af91420d5c99ed7b303bc2-6f18562a487c9923bf9c6fc8250d2eb2626bed58.profdata to chrome-mac-arm-main-1719971996-c1186dec327e76e354428d821525b88a96b6c2c6-080be1a73624d3bb9b4dd31d57cd5939e19693fa.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-mac-arm-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:mac-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I82621ca44934a7b839fc5fa661979234da465922\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5674080\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322544}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T03:41:16+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fd75c37a1ba043e170212c4cba6f93e008074edb"
  },
  {
    "id": 503,
    "message": "[Extensions] Move some DNR-related helper functions. (4)\n\nThis CL moves another set of functions and also cleans up a bit in\nthe helper class implementation file.\n\nBug: 40593486\nChange-Id: I0143dbaa63f7eb64a472535a092fb1271210df12\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5670887\nCommit-Queue: David Bertoni <dbertoni@chromium.org>\nReviewed-by: Tim <tjudkins@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322543}",
    "author": "David Bertoni",
    "author_email": "dbertoni@chromium.org",
    "date": "2024-07-03T03:40:52+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/77dce22e3c232dad4a759dec42969af1662038aa"
  },
  {
    "id": 504,
    "message": "Make all WebAudio histograms expire at the same time.\n\nTechnically we are supposed to carefully consider each histogram as it\nexpires, and decide if we really need to keep it.  But for the past\nyear we have just automatically extended the histograms since we don't\nhave bandwidth to do this consideration right now.\n\nBy making all the histograms expire at the same time we actually\nincrease the chance that we will set aside some time to evaluate what\nis actually necessary and what is not.  And, if we do decide to\nautomatically extend again we will only have to do it twice per year\ninstead of the current several times per year.\n\nNote that some histograms will expire earlier after this change than\nbefore this change.\n\nBug: 346342640\nChange-Id: I9cb671c288b19d446a43b7cc8bbce5f9787d053d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5671544\nCommit-Queue: Hongchan Choi <hongchan@chromium.org>\nAuto-Submit: Michael Wilson <mjwilson@chromium.org>\nReviewed-by: Hongchan Choi <hongchan@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322542}",
    "author": "Michael Wilson",
    "author_email": "mjwilson@chromium.org",
    "date": "2024-07-03T03:39:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/976cb4846288f62ab55b2496c4673512950e1208"
  },
  {
    "id": 505,
    "message": "[shared storage] Record usage for when addModule() is called with cross-origin script\n\nWhy: We proposed a potentially breaking change to addModule() in\nhttps://github.com/WICG/shared-storage/pull/158. That is, currently,\na website can call addModule() everywhere and expect it to fail when\nthey're not in their own context; however, with cross-origin script\nsupport, these calls could unexpectedly succeed.\n\nThus, we track potentially non-forward-compatible addModule() usages to\nassess the potential impact and guide further decisions. We anticipate\nthis usage to be low.\n\nBug: 350764023\nChange-Id: I5da4583ed583baa398bdb2cadf8955d341474050\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5665746\nReviewed-by: Josh Karlin <jkarlin@chromium.org>\nCommit-Queue: Yao Xiao <yaoxia@chromium.org>\nReviewed-by: Sun Yueru <yrsun@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322541}",
    "author": "Yao Xiao",
    "author_email": "yaoxia@chromium.org",
    "date": "2024-07-03T03:37:27+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/259f6e5ba908edb5742c5b56ee93ccd2b39e9ddf"
  },
  {
    "id": 506,
    "message": "Revert \"Log histograms for MediaPlayer HLS playback\"\n\nThis reverts commit 1ccf5f198a4e706a98eeff2119e844e396b64dcd.\n\nReason for revert: Link failure on `UBSan vptr Release`\n\nOriginal change's description:\n> Log histograms for MediaPlayer HLS playback\n>\n> Records Media.HLS.AdvancedFeatureTagType, Media.Hls.HLSEncryptionMode,\n> Media.HLS.PlaylistSegmentExtension, and Media.HLS.UnparsableManifest\n> by parsing a media playlist with an implementation of TagRecorder.\n>\n> Bug: 338277331\n> Change-Id: I18972f09875979c8ecd41ffe3d089c0a04b99247\n> Reviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5637853\n> Commit-Queue: Ted (Chromium) Meyer <tmathmeyer@chromium.org>\n> Reviewed-by: Eugene Zemtsov <eugene@chromium.org>\n> Reviewed-by: Evan Liu <evliu@google.com>\n> Cr-Commit-Position: refs/heads/main@{#1322523}\n\nBug: 338277331\nChange-Id: I697d9d4b31aa4c2344a0a3468041b48b354aa7c3\nNo-Presubmit: true\nNo-Tree-Checks: true\nNo-Try: true\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5674099\nCommit-Queue: Kalvin Lee <kdlee@chromium.org>\nBot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nOwners-Override: Kalvin Lee <kdlee@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322540}",
    "author": "Kalvin Lee",
    "author_email": "kdlee@chromium.org",
    "date": "2024-07-03T03:34:08+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/cdafd251e91fb7e12b84173defe26d11e7cb7b10"
  },
  {
    "id": 554,
    "message": "Add fieldtrial test configs for Growth Framework predefined params.\n\nBug: 345539381\nChange-Id: I710346f741d3c18c96d818d9d6df27f1c1405047\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5604672\nCommit-Queue: Li Lin <llin@chromium.org>\nReviewed-by: Tao Wu <wutao@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311632}",
    "author": "Li Lin",
    "author_email": "llin@chromium.org",
    "date": "2024-06-07T00:50:59+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fc9145f43003803b5d0906de7fc8e3a2bfd16464"
  },
  {
    "id": 507,
    "message": "Roll Chrome Lacros Amd64 Generic PGO Profile\n\nRoll Chrome Lacros Amd64 PGO profile from chrome-chromeos-amd64-generic-main-1719921560-e1cf704dd00c5e478120bb0a86a45e1dcfbebac7-e8f9773d53fd3015c991126e794ab7a56a7d2f71.profdata to chrome-chromeos-amd64-generic-main-1719965278-80b42b6255bff46bddd59ec3f9f1cabb7e8440a5-a5e6ff3fdc202f42c23c3850881567fd705e6c29.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-lacros-amd64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I7cbe9bcca57cf66368c91ec8103295f71c726e4c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5673609\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322539}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T03:32:04+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/426bd1bcf82242d9ae1fddafc9f3da2340f219f1"
  },
  {
    "id": 508,
    "message": "Align enumerations for AI metrics in xml and code\n\nThe CL aligns the enumeration definition in enums.xml and ai_metrics.h.\nLINT checks are also added to prevent future disalignment.\n\nNO_IFTTT=Adding missing LINTS, no code change required\n\nBug: None\nChange-Id: I7198089d935af32f2691684aa27afdce38344a18\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5671195\nCommit-Queue: Jiacheng Guo <gjc@google.com>\nReviewed-by: Takashi Toyoshima <toyoshim@chromium.org>\nReviewed-by: Mingyu Lei <leimy@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322538}",
    "author": "Jiacheng Guo",
    "author_email": "gjc@google.com",
    "date": "2024-07-03T03:30:49+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1dcbe99bd5c0532d8a5348bb7ff16379a523acb6"
  },
  {
    "id": 509,
    "message": "Roll Chrome Android ARM64 PGO Profile\n\nRoll Chrome Android ARM64 PGO profile from chrome-android64-main-1719899919-8c74f5513622cbc58cb787155b0a340522f4531a-765fe7058aa26f3903ce870284b4803ace47b8e9.profdata to chrome-android64-main-1719964707-2705fe43495b72c643ef8d72e7f24bd2f8abe5da-6f18562a487c9923bf9c6fc8250d2eb2626bed58.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-android-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I4bb4a7b5fb9240ed66a520a8dde5145cf139cd43\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5673596\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322537}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T03:30:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/bd56ddbce8e8d02660e4d32c755011bf3eec1aa4"
  },
  {
    "id": 510,
    "message": "Add relevant members as owners of IPP histograms.\n\nWe rely on these histograms and would like to be notified if we\nnear expiry.\n\nBug: b:347731426\nChange-Id: Ic54d363256d6763e4da75a24d682ac3820fb404e\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5666112\nReviewed-by: Ashley Newson <ashleynewson@chromium.org>\nCommit-Queue: Takashi Toyoshima <toyoshim@chromium.org>\nAuto-Submit: Abhijith Nair <abhijithnair@chromium.org>\nReviewed-by: Takashi Toyoshima <toyoshim@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322536}",
    "author": "Abhijith Nair",
    "author_email": "abhijithnair@chromium.org",
    "date": "2024-07-03T03:26:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d343d82162a5100c134e1260e1ef6f9b4f9f8a36"
  },
  {
    "id": 511,
    "message": "Extend expiration date of histograms related to IPP.\n\nChange-Id: I4b1ca363d88ff89d9562b70bd94e758ec6c41977\nBug: b:344591903\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5665704\nCommit-Queue: Takashi Toyoshima <toyoshim@chromium.org>\nReviewed-by: Ashley Newson <ashleynewson@chromium.org>\nReviewed-by: Takashi Toyoshima <toyoshim@chromium.org>\nAuto-Submit: Abhijith Nair <abhijithnair@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322535}",
    "author": "Abhijith Nair",
    "author_email": "abhijithnair@chromium.org",
    "date": "2024-07-03T03:26:14+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5101fc07370a52fabfbcdf9a29466dfe2825819d"
  },
  {
    "id": 512,
    "message": "[Ash] BUILD.gn file for //chrome/browser/ash/printing/enterprise\n\nThis CL is in preparation for the bigger refactoring of\nb/335294351, i.e., create BUILD.gn file for\n//chrome/browser/ash/printing.\n\nFixed: 349929005\nChange-Id: Ica977c77b90e544a67ba05235ff9ae135e67a21d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5671705\nReviewed-by: Kyle Horimoto <khorimoto@chromium.org>\nCommit-Queue: Di Wu <diwux@google.com>\nCr-Commit-Position: refs/heads/main@{#1322534}",
    "author": "Di Wu",
    "author_email": "diwux@google.com",
    "date": "2024-07-03T03:22:51+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f3b7c00ed532c792b044b5b66874360b6579fe6d"
  },
  {
    "id": 513,
    "message": "SearchPrefetch: Remove kSearchPrefetchSkipsCancel\n\nThis feature was enabled by default by https://crrev.com/c/4469310.\n\nNO_IFTTT=Changes will be done in the separate repository later.\n\nChange-Id: I04e933b7dd49e7c842bfd106b2536f5d516396c3\nBug: b/262915418\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5670607\nReviewed-by: Takashi Toyoshima <toyoshim@chromium.org>\nCommit-Queue: Hiroki Nakagawa <nhiroki@chromium.org>\nReviewed-by: Lingqi Chi <lingqi@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322533}",
    "author": "Hiroki Nakagawa",
    "author_email": "nhiroki@chromium.org",
    "date": "2024-07-03T03:22:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/41a083672130d62fc2bdc063992fd29f92ae1652"
  },
  {
    "id": 536,
    "message": "[NTP][Enterprise] Add attachments to NTP Google Calendar card\n\n* Pulled out some styles in cr-chip into variables so they could be\n  tweaked for this UI.\n* Set cr-chip white-space to nowrap because of wrapping that was\n  happening when the line of attachments was overflowing.\n* Optimized images.\n* Updated handler unittests to support multiple test server response\n  json files.\n\nscreenshot: http://screenshot/atkRdBXyD3m6gy2\n\nBug: 345258413\nChange-Id: Iba569ec40286d4233e478647d9d7e9e0635fcfd1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5601712\nReviewed-by: John Lee <johntlee@chromium.org>\nCommit-Queue: Riley Tatum <rtatum@google.com>\nReviewed-by: Tibor Goldschwendt <tiborg@chromium.org>\nReviewed-by: Mustafa Emre Acer <meacer@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311650}",
    "author": "Riley Tatum",
    "author_email": "rtatum@google.com",
    "date": "2024-06-07T01:25:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2a3f596567f102ac864379a80b0dad4a6852f591"
  },
  {
    "id": 514,
    "message": "Roll Chrome Linux PGO Profile\n\nRoll Chrome Linux PGO profile from chrome-linux-main-1719943160-babba220d4051f723f1bb4c843bf8922e8a9504e-c0491a6d993bc5e9298b223fcaf8f2ed407d80ea.profdata to chrome-linux-main-1719964707-03c5ff9666a2eaf127fd2b6ed594fbf11b290d28-6f18562a487c9923bf9c6fc8250d2eb2626bed58.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-linux-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I8dbb327a55b42623601afe2b6bb50d8b20ea9554\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5673805\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322532}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T03:15:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b0aeea74e31c35291c34219030ab43ef7f1e6d3c"
  },
  {
    "id": 515,
    "message": "Split cbcm_enroll.py script into 2\n\nThe original one will stay as simple, the new one with reporting events logic will be moved into a new webdriver script, and that one serves report_cbcm_event test only.\n\nTESTED= Run 2 related tests locally, both passed. The tests in tryjob also passed.\n\nBug: b:327797500\nChange-Id: I37e5495f27b82514295fc5a1e324ca6c03ba9f3f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5671656\nReviewed-by: Josh Hilke <jrhilke@google.com>\nCommit-Queue: Xiang Ji <jxiang@google.com>\nCr-Commit-Position: refs/heads/main@{#1322531}",
    "author": "Xiang Ji",
    "author_email": "jxiang@google.com",
    "date": "2024-07-03T03:13:53+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/00ed464976d468ef0ac404249d743afb7384b45f"
  },
  {
    "id": 516,
    "message": "[Extensions] Allow service worker requests to continue without a cert\n\nCertain requests may request, but not require, a client cert. Today,\nthis will result in one of three things happening if there isn't a\nclient cert available (either from a previous request or from\nenterprise policy):\n1) The request will fail\n2) The user will be shown a cert picker dialog\n3) The request will continue without a certificate\n\nOn desktop platforms for WebContents-based requesting contexts:\n* If there are no certs to choose from, the request will continue\n  without a cert.\n* If there are client certs and the WebContents supports showing\n  dialogs, the cert picker will be shown.\n* If there are certs and the WebContents does not support showing\n  dialogs, the request will fail.\n\nOn Android for WebContents-based requesting contexts:\n* We will always call out to the native cert-picker, which may or\n  may not show a dialog (depending on other device configurations\n  which Chrome doesn't know about).\n\nOn desktop and Android platforms for service worker requesting\ncontexts:\n* The request will always fail.\n\nThis CL changes this behavior to allow requests from extension\nbackground service workers to proceed without a client cert if there\nare no certs to choose from; this then matches the behavior for\nextension background and offscreen WebContents (which do not support\nshowing dialogs).\n\nBug: 333954429\nChange-Id: Ia6111f3bd499998d6244945daa13ac67774804bf\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5455480\nReviewed-by: Emily Stark <estark@chromium.org>\nReviewed-by: David Benjamin <davidben@chromium.org>\nReviewed-by: Richard (Torne) Coles <torne@chromium.org>\nReviewed-by: Alex Moshchuk <alexmos@chromium.org>\nReviewed-by: Luke Halliwell <halliwell@chromium.org>\nReviewed-by: Andrey Kosyakov <caseq@chromium.org>\nCommit-Queue: Devlin Cronin <rdevlin.cronin@chromium.org>\nReviewed-by: Wez <wez@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322530}",
    "author": "Devlin Cronin",
    "author_email": "rdevlin.cronin@chromium.org",
    "date": "2024-07-03T03:13:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/df260f5e98ba465534694727f573ddb1222c199e"
  },
  {
    "id": 517,
    "message": "lobster: add feature flag\n\nBUG=b:348280305\n\nChange-Id: I7ce3135f0b10710783298604b8ebe71ffd7e0149\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5647454\nReviewed-by: Darren Shen <shend@chromium.org>\nCommit-Queue: Curtis McMullan <curtismcmullan@chromium.org>\nReviewed-by: Scott Violet <sky@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322529}",
    "author": "Curtis McMullan",
    "author_email": "curtismcmullan@chromium.org",
    "date": "2024-07-03T03:04:19+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9e3e20cd4875442fc1aaab7b65604a53cab1f1a3"
  },
  {
    "id": 518,
    "message": "[Discovery] Delay sink discovery feature enabled by default on Win\n\nCurrently, background discovery is delayed on Windows to avoid showing\nthe firewall settings prompt. crrev.com/c/5597929 made background discovery controlled by a feature flag and the feature is disabled by default, which caused background discovery to start immediately on Windows.\n\nThis CL enables the feature by default on Windows.\n\nBug: 345056325\nChange-Id: I98adb7f849d5f1ba500c70e12a04a634dcde2f85\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5640711\nCommit-Queue: Muyao Xu <muyaoxu@google.com>\nReviewed-by: Takumi Fujimoto <takumif@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322528}",
    "author": "Muyao Xu",
    "author_email": "muyaoxu@google.com",
    "date": "2024-07-03T03:01:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2443820f1afae0c243dae9c729ee94a4be1fd811"
  },
  {
    "id": 519,
    "message": "[LCPP] Add UpdateFrequencyStatAndTryGetEntry function\n\nThis function updates protobuf::Map<string, T> entry following the entry\nappearance in LcppStringFrequencyStatData.\n\nThis CL is a preparation for double keyed LCPP.\nDesign doc(google internal): https://docs.google.com/document/d/1AVi2qvPvGioUcuLMwgMfFyvhLy9scc_-8fVI8GxeV3w/edit?usp=sharing\n\nBug: 343093433\nChange-Id: I9215deaf13024c7db1ebdc3dccb2d01b00595f19\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5670080\nCommit-Queue: Yoichi Osato <yoichio@chromium.org>\nReviewed-by: Minoru Chikamune <chikamune@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322527}",
    "author": "Yoichi Osato",
    "author_email": "yoichio@chromium.org",
    "date": "2024-07-03T03:00:56+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/128a843564458626150709af5a90f9706b92824f"
  },
  {
    "id": 520,
    "message": "Roll Arm AFDO from 128.0.6570.0_rc-r1-merged to 128.0.6571.0_rc-r1-merged\n\nThis CL may cause a small binary size increase, roughly proportional\nto how long it's been since our last AFDO profile roll. For larger\nincreases (around or exceeding 100KB), please file go/crostc-bug.\n\nPlease note that, despite rolling to chrome/android, this profile is\nused for both Linux and Android.\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/afdo-arm-chromium\nPlease CC c-compiler-chrome@google.com,mobiletc-prebuild@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium Main: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: mobiletc-prebuild@google.com\nChange-Id: Ie212328cfbd53560a2d5161c85da385477c0def0\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5673808\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322526}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T02:59:56+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/080be1a73624d3bb9b4dd31d57cd5939e19693fa"
  },
  {
    "id": 521,
    "message": "Roll Skia from d02998fba957 to 7c69f39fa85b (1 revision)\n\nhttps://skia.googlesource.com/skia.git/+log/d02998fba957..7c69f39fa85b\n\n2024-07-02 kjlubick@google.com Remove staging gni file groups for sksl\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/skia-autoroll\nPlease CC robertphillips@google.com,skiabot@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Skia: https://bugs.chromium.org/p/skia/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux-blink-rel;luci.chromium.try:linux-chromeos-compile-dbg;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel\nCq-Do-Not-Cancel-Tryjobs: true\nBug: None\nTbr: robertphillips@google.com\nChange-Id: Ic116e8b6e0f9e91a7e32e2233ed294d6e0fed1ff\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5672968\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322525}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T02:59:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/193362e4d7025ac1f636cf7e2fed5b223606e73a"
  },
  {
    "id": 522,
    "message": "Roll vulkan-deps from d5f8cf18fe16 to 492062c10738 (2 revisions)\n\nhttps://chromium.googlesource.com/vulkan-deps.git/+log/d5f8cf18fe16..492062c10738\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/external/github.com/KhronosGroup/Vulkan-Tools/+log/35eb0bf879cf10dd7bd5835ba8399682f741a0c0..7e13360e42364fdd1f07fe00f19d0432b12db055\n  https://chromium.googlesource.com/external/github.com/KhronosGroup/Vulkan-ValidationLayers/+log/803ef169856697d93f69b5d01489763a114e81cb..1326072ad05e2cee4339f59df305988dd421faa8\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/vulkan-deps-chromium-autoroll\nPlease CC abdolrashidi@google.com,angle-team@google.com,dneto@google.com,radial-bots+chrome-roll@google.com,radial-bots@google.com,webgpu-developers@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86;luci.chromium.try:dawn-linux-x64-deps-rel;luci.chromium.try:win-asan;luci.chromium.try:linux_chromium_cfi_rel_ng\nBug: None\nTbr: abdolrashidi@google.com,dneto@google.com,radial-bots+chrome-roll@google.com\nChange-Id: I9ab6d34bb78b970fbd5d2e34be2ebd5f84593e40\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5671634\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1322524}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-03T02:58:47+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/88c6408b6631b9c1e126edaac2cd6e9f8e550a8b"
  },
  {
    "id": 523,
    "message": "Log histograms for MediaPlayer HLS playback\n\nRecords Media.HLS.AdvancedFeatureTagType, Media.Hls.HLSEncryptionMode,\nMedia.HLS.PlaylistSegmentExtension, and Media.HLS.UnparsableManifest\nby parsing a media playlist with an implementation of TagRecorder.\n\nBug: 338277331\nChange-Id: I18972f09875979c8ecd41ffe3d089c0a04b99247\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5637853\nCommit-Queue: Ted (Chromium) Meyer <tmathmeyer@chromium.org>\nReviewed-by: Eugene Zemtsov <eugene@chromium.org>\nReviewed-by: Evan Liu <evliu@google.com>\nCr-Commit-Position: refs/heads/main@{#1322523}",
    "author": "Ted Meyer",
    "author_email": "tmathmeyer@chromium.org",
    "date": "2024-07-03T02:58:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1ccf5f198a4e706a98eeff2119e844e396b64dcd"
  },
  {
    "id": 524,
    "message": "Add category specific illustrations and strings for no results screen.\n\nBug: b:348067874\nChange-Id: I5ca73fe420a960dd6e32528cda13fa91f24a8dd3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5647860\nReviewed-by: Michelle Chen <michellegc@google.com>\nReviewed-by: Michael Cui <mlcui@google.com>\nCommit-Queue: Darren Shen <shend@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1322522}",
    "author": "Darren Shen",
    "author_email": "shend@chromium.org",
    "date": "2024-07-03T02:57:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/29d49633745dc9a879ce4987969b7fe4992bceff"
  },
  {
    "id": 525,
    "message": "Roll Perfetto Trace Processor Win from a29e84549d26 to 9395152fa95a\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/a29e84549d26..9395152fa95a\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-trace-processor-win-chromium\nPlease CC chrometto-team@google.com,perfetto-bugs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: perfetto-bugs@google.com\nChange-Id: Idce25d465b84fa22c0324e2398b1448f86f4263f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606662\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311661}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-07T01:39:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c7008e15a88d3a2ea9d24ea40ad9da8923027fc1"
  },
  {
    "id": 526,
    "message": "CrOS Settings: Add usage documentation for settings-slider-row\n\nBug: b/333454343\nChange-Id: I0abe49e12b94f6844d9f2af6c0b311bfd044df11\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606003\nReviewed-by: Wes Okuhara <wesokuhara@google.com>\nCommit-Queue: Connie Xu <conniekxu@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311660}",
    "author": "conniekxu",
    "author_email": "conniekxu@chromium.org",
    "date": "2024-06-07T01:38:52+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5f383ecd2f62c0047634ad92afd6e204f9e79175"
  },
  {
    "id": 527,
    "message": "[shared storage] Allow createWorklet() in opaque origin contexts\n\nPR: https://github.com/WICG/shared-storage/pull/156\n\nBug: 345274915\nChange-Id: I748df668635d77a2d158bd86222aa2fdecfed3eb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5564969\nReviewed-by: Cammie Smith Barnes <cammie@chromium.org>\nCommit-Queue: Yao Xiao <yaoxia@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311659}",
    "author": "Yao Xiao",
    "author_email": "yaoxia@chromium.org",
    "date": "2024-06-07T01:37:29+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6b39114afad0ac8e7ec517dc83a2dbf4dd9549ea"
  },
  {
    "id": 552,
    "message": "[GM2 Cleanup] find_bar_view.cc\n\n#chrome-refresh-2023-cleanup\n\nChange-Id: I3102fddbb78a6bbe7fc5d7059abc531960c26dd1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5605654\nReviewed-by: Joseph Park <josephjoopark@chromium.org>\nCommit-Queue: Joseph Park <josephjoopark@chromium.org>\nAuto-Submit: Dana Fried <dfried@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311634}",
    "author": "Dana Fried",
    "author_email": "dfried@chromium.org",
    "date": "2024-06-07T00:54:05+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/0a96f14a8b50fba0bd86967c32a79595a39893df"
  },
  {
    "id": 528,
    "message": "Roll clank/internal/apps from 3f7cd1d137e2 to 29dd0d981187 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/3f7cd1d137e2..29dd0d981187\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,gangwu@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: gangwu@google.com\nNo-Try: true\nChange-Id: I97b91b9109e1fbe01a4ef78bb38fe4b60edf1bda\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606376\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311658}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-07T01:34:14+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/582321cd59e0d91a74302c78a0fb3c5410f823f9"
  },
  {
    "id": 529,
    "message": "Expose TabFeatures and BrowserWindowFeatures.\n\nThis CL has no intended user-visible effect.\n\nThis CL updates TabInterface to expose TabFeatures.\nThis CL updates BrowserWindowInterface to expose the foreground tab, as\nwell as BrowserWindowFeatures.\n\nThis CL moves the headers for TabFeatures and BrowserWindowFeatures to\nreflect this change.\n\nThe main effect of this CL is that consumers of TabInterface and\nBrowserWindowInterface can selectively acquire access to controllers in\nTabFeatures and BrowserWindowFeatures. The build target dependency graph\nonly reflects controllers that are used.\n\nChange-Id: I37fbb776b4c83be9241443f3798f12856863d4df\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5602163\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Erik Chen <erikchen@chromium.org>\nReviewed-by: Scott Violet <sky@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311657}",
    "author": "Erik Chen",
    "author_email": "erikchen@chromium.org",
    "date": "2024-06-07T01:33:14+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/707c621adf698abdc68ad942dfcde5fd9472ae5b"
  },
  {
    "id": 530,
    "message": "Clean up IsChromeRefresh2023 in c/b/ui/view/frame/\n\nThis includes removing quite a bit of dead code calculating the distance\nbetween the top of the window and the top of the tabstrip. Prior to\nRefresh, this was up to the frame to decide; now it's a constant owned\nby the tab strip.\n\nBug: 335903961\nChange-Id: Ie66afac1744e9e47ea6b3c70558798616307c42d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5590464\nReviewed-by: Elly FJ <ellyjones@chromium.org>\nCommit-Queue: Taylor Bergquist <tbergquist@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: David Pennington <dpenning@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311656}",
    "author": "Taylor Bergquist",
    "author_email": "tbergquist@chromium.org",
    "date": "2024-06-07T01:31:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/81dee0f4f06b4015b8d9176f1adcd13f6b3bd96f"
  },
  {
    "id": 531,
    "message": "Mark http/tests/media/mixed-range-response.html as flaky timeout.\n\nBug: 40226500\nChange-Id: I3c83e529a14f96c14c8fa10e77e1e45d659b36fa\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5604679\nAuto-Submit: Dale Curtis <dalecurtis@chromium.org>\nCommit-Queue: Dale Curtis <dalecurtis@chromium.org>\nBot-Commit: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCommit-Queue: Rubber Stamper <rubber-stamper@appspot.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311655}",
    "author": "Dale Curtis",
    "author_email": "dalecurtis@chromium.org",
    "date": "2024-06-07T01:27:56+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b6af31fdae4115994a003e3bc851d8bf3f64e074"
  },
  {
    "id": 532,
    "message": "Introduce printing/windows_types.h\n\nCreate a printing-specific Windows types include file, similar to what\nis generally done in base/win/windows_types.h to help source avoid\nincluding windows.h.  This new file contains Windows system definitions\nthat are specific to printing.\n\nChange-Id: I8a68aefcd9677b807f1df97a6d4d9d89874d1af6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606124\nReviewed-by: Lei Zhang <thestig@chromium.org>\nCommit-Queue: Alan Screen <awscreen@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311654}",
    "author": "Alan Screen",
    "author_email": "awscreen@chromium.org",
    "date": "2024-06-07T01:27:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f84d71455559d697ecbcd7940d317c295a48de0e"
  },
  {
    "id": 533,
    "message": "Update desktop iOS's User-Agent OS version\n\nThis makes it consistent with desktop Chrome, Firefox, and Safari,\nwhich have all been frozen to 10_15_7 for web compatibility reasons.\n\nSee https://crbug.com/40167872 for more history on that.\n\nBug: 40231533\nChange-Id: Ic90b4a559e139d621a9b5d545eb391616070164a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5599372\nReviewed-by: Ali Juma <ajuma@chromium.org>\nCommit-Queue: Mike Taylor <miketaylr@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311653}",
    "author": "Mike Taylor",
    "author_email": "miketaylr@chromium.org",
    "date": "2024-06-07T01:26:47+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5953dfab5bd342ef0c0ca0aa8f59b20531496a29"
  },
  {
    "id": 534,
    "message": "[NTP][Enterprise] Add time until next event when expanded\n\nNote: I made the timeStatus_ string empty when the event is expanded\nand used that to determine whether to show timeStatus_ because when more\nlogic is added, expanded events will show a timeStatus_ in certain\nedge cases. Therefore, determining visibility based off of the string\nbeing empty or not is more accurate than just expanded or not expanded.\n\nScreenshots:\nhttp://screenshot/AEa6cc68dixKXLK\nhttp://screenshot/4r3Ci593vwTNDCT\nhttp://screenshot/k8xwn4USV6EXMV9\n\nBug: 345327808\nChange-Id: I445d9056986a13119584cc627fd8609b16a53fc4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5598391\nReviewed-by: Thomas Lukaszewicz <tluk@chromium.org>\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Paul Adedeji <pauladedeji@google.com>\nCommit-Queue: Riley Tatum <rtatum@google.com>\nCr-Commit-Position: refs/heads/main@{#1311652}",
    "author": "Riley Tatum",
    "author_email": "rtatum@google.com",
    "date": "2024-06-07T01:26:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b6b0d5a48f941b369388b01c155aa2729199a61d"
  },
  {
    "id": 535,
    "message": "[NTP][Enterprise] Add button to join video call to calendar card\n\nScreenshot: http://screenshot/8ENJLf8Zexe5WYt\n\nBug: 345295618\nChange-Id: Ib3326a36d18c50cd41c4665921adcb88f94b244a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5598232\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Tibor Goldschwendt <tiborg@chromium.org>\nCommit-Queue: Riley Tatum <rtatum@google.com>\nReviewed-by: Alex Gough <ajgo@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311651}",
    "author": "Riley Tatum",
    "author_email": "rtatum@google.com",
    "date": "2024-06-07T01:25:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/5c8b1f6ce478a8f7880f2170e25d662584563c4d"
  },
  {
    "id": 537,
    "message": "Roll Chrome Android ARM64 PGO Profile\n\nRoll Chrome Android ARM64 PGO profile from chrome-android64-main-1717696741-53aed4c33cdd2284cba7e4a901d44ad12405276d-0563a68f3dd37c8c090fb3b01590c2b5e356840d.profdata to chrome-android64-main-1717710885-7598b6da629bbe9e26a349c460233dd260d575d8-642d09383f2e68d1d3ea5069f40c6b247e9dc229.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-android-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: If07bdba47f61df852fc860c8710e206e6370bcd5\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606357\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311649}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-07T01:24:14+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d6a308c312a0897ee0ac284a6a2801ac0f4e4c1c"
  },
  {
    "id": 538,
    "message": "[GM2 Cleanup] title_origin_label.cc\n\n#chrome-refresh-2023-cleanup\n\nChange-Id: I6c5c4e31bce07aa41506690b89f0f8933b71df38\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606504\nReviewed-by: Joseph Park <josephjoopark@chromium.org>\nCommit-Queue: Dana Fried <dfried@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311648}",
    "author": "Dana Fried",
    "author_email": "dfried@chromium.org",
    "date": "2024-06-07T01:23:09+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d8dc0ea621c3c8f7248aa527ad046b2158f95e3a"
  },
  {
    "id": 539,
    "message": "[CC Number Validation Target] Network identifiers include (7d/9)\n\nThis patch removes the temporary \"credit_card_network_identifiers.h\"\ninclude from the \"credit_card.h\" header (which does not need it) and\nupdates the files that were using credit card network identifiers (e.g.,\nkVisaCard, kMasterCard) to include \"credit_card_network_identifiers.h\"\ndirectly. In the few cases where the \"credit_card.h\" was included only\nfor the network identifiers, that include is removed.\n\nDesign doc: https://bit.ly/cc-number-validation-target-dd\n\nBug: b:281812289\nChange-Id: I3ff883335cc2be5aad31fc54836d4709886c886b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5594452\nAuto-Submit: Rouslan Solomakhin <rouslan@chromium.org>\nCommit-Queue: Christoph Schwering <schwering@google.com>\nReviewed-by: Christoph Schwering <schwering@google.com>\nReviewed-by: Mikel Astiz <mastiz@chromium.org>\nReviewed-by: Vidhan Jain <vidhanj@google.com>\nCr-Commit-Position: refs/heads/main@{#1311647}",
    "author": "Rouslan Solomakhin",
    "author_email": "rouslan@chromium.org",
    "date": "2024-06-07T01:21:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/db13dae3afb514f487535452a42f9e28365d7f04"
  },
  {
    "id": 540,
    "message": "Roll V8 from 35f96db75367 to 919383d7c777 (2 revisions)\n\nhttps://chromium.googlesource.com/v8/v8.git/+log/35f96db75367..919383d7c777\n\n2024-06-06 v8-ci-autoroll-builder@chops-service-accounts.iam.gserviceaccount.com Version 12.7.210\n2024-06-06 dmercadier@chromium.org [turboshaft] Fix VariableReducer bug with loop phis\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/v8-chromium-autoroll\nPlease CC liviurau@google.com,machenbach@google.com,v8-waterfall-gardener@grotations.appspotmail.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-blink-rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:dawn-linux-x64-deps-rel\nChange-Id: I27953aa67129d8b136e351021bff8490aa6eb4c4\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606358\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311646}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-07T01:20:30+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9c75677cc0e7243c83e104740c353fc936f9f645"
  },
  {
    "id": 541,
    "message": "Roll Website from ff6eddb8ebf3 to 27928a8bdd0f (1 revision)\n\nhttps://chromium.googlesource.com/website.git/+log/ff6eddb8ebf3..27928a8bdd0f\n\n2024-06-06 jocelyntran@google.com Update chromium-os guide with vncviewer command\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/website-chromium\nPlease CC dpranke@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Website: https://bugs.chromium.org/p/chromium/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: dpranke@google.com\nChange-Id: I31b7d9034dfc4a8da0adeb3745dcad7bc6d89e4d\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5603935\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311645}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-06-07T01:19:38+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/126ae7dadb98fdcfa0b123bfb344b13346b3b971"
  },
  {
    "id": 542,
    "message": "Roll clank/internal/apps from edc04c3ccf66 to 3f7cd1d137e2 (1 revision)\n\nhttps://chrome-internal.googlesource.com/clank/internal/apps.git/+log/edc04c3ccf66..3f7cd1d137e2\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/clank-apps-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,gangwu@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nTbr: gangwu@google.com\nNo-Try: true\nChange-Id: I79cd41bd00c5eec59aaf7ca5b20494b05ca693ea\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606634\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311644}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-07T01:19:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1a772fd75238e06f2b9a579674b40867fd41f2aa"
  },
  {
    "id": 543,
    "message": "Disable tests in DumpAccessibilityTreeTest\n\nDisable tests\nAll/DumpAccessibilityTreeTest.AccessibilityOffscreenSelect/android\nAll/DumpAccessibilityTreeTest.AccessibilitySvgTextAlternativeComputation/android\nAll/YieldingParserDumpAccessibilityTreeTest.AccessibilitySvgTextAlternativeComputation/android\nbecause they failed in\nhttps://ci.chromium.org/ui/p/chrome/builders/ci/android-automotive-12l-x64-rel-tests/933\n\nBug: b/345567344\nChange-Id: I2760d7dbfd50149696b6fc7e50da529e6e63131b\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606508\nReviewed-by: Brandon Wylie <wylieb@google.com>\nCommit-Queue: Gang Wu <gangwu@chromium.org>\nCommit-Queue: Brandon Wylie <wylieb@google.com>\nAuto-Submit: Gang Wu <gangwu@chromium.org>\nOwners-Override: Gang Wu <gangwu@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311643}",
    "author": "Gang Wu",
    "author_email": "gangwu@chromium.org",
    "date": "2024-06-07T01:18:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/86748466137982bda75895a9a20520f127faa705"
  },
  {
    "id": 544,
    "message": "[lensoverlay] Change overflow menu style.\n\nBefore: https://screenshot.googleplex.com/BXUm4N4nXkpF7dd\nAfter: https://screenshot.googleplex.com/C7eVhAiiqeQYvZj\n\nBug: 344679866, 343949820\nChange-Id: I82982a2976f10c90c849c835e94a3ff7aea71558\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5599003\nReviewed-by: Duncan Mercer <mercerd@google.com>\nCommit-Queue: Nihar Majmudar <niharm@google.com>\nCr-Commit-Position: refs/heads/main@{#1311642}",
    "author": "Nihar Majmudar",
    "author_email": "niharm@google.com",
    "date": "2024-06-07T01:13:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/6a282b186cd984ff82842cc89db612c2a5b66e5f"
  },
  {
    "id": 545,
    "message": "Disable the DWARF line number reader in sanitizer builds\n\nGate enable_stack_trace_line_numbers on print_unsymbolized_stack_traces.\nThis is typically set to true in sanitizer builds, where the binary\nrelies on an external symbolizer script; in that case, there's no point\nin using the (currently much) slower in-process line number reader. This\nalso fixes the check in WillSymbolizeToStreamForTesting() to check the\ncorrect buildflag rather than assuming it is only implicitly set in\nsanitizer builds.\n\nBug: 340987671\nChange-Id: Id3a23100d6b69bb923daf7f6173dbed4d76f7e0a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5605953\nCommit-Queue: Daniel Cheng <dcheng@chromium.org>\nReviewed-by: Thomas Anderson <thomasanderson@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311641}",
    "author": "Daniel Cheng",
    "author_email": "dcheng@chromium.org",
    "date": "2024-06-07T01:12:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/8d764e57e872ba9efb1b8f3ce0e5d88094cc651e"
  },
  {
    "id": 546,
    "message": "[Read Aloud] Convert the language for the pack manager for the toast\n\nPreviously we were converting to an available language, but the language\nwe're showing a toast for was just downloaded and may not yet be\navailable, so we would previously convert to the wrong locale. Instead,\nconvert to a pack manager supported locale.\n\nFixed: 343821828\nChange-Id: I19546568ab3340ef8b8d2f0c8e960c175618e7e8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606213\nCommit-Queue: Lauren Winston <lwinston@google.com>\nReviewed-by: Lauren Winston <lwinston@google.com>\nAuto-Submit: Kristi Saney <kristislee@google.com>\nCr-Commit-Position: refs/heads/main@{#1311640}",
    "author": "Kristi Saney",
    "author_email": "kristislee@google.com",
    "date": "2024-06-07T01:11:45+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c2892350827b283685a666f169bd3e3c89dab8f7"
  },
  {
    "id": 547,
    "message": "Roll optimization-guide from 4190c31cd456 to 79fea6341855\n\nhttps://chrome-internal.googlesource.com/chrome/components/optimization_guide.git/+log/4190c31cd456..79fea6341855\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/optimization-guide-chromium\nPlease CC chrome-intelligence-core@google.com,sophiechang@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: I8acbf099d3b1be44f3fe18d6f8e50d8d7cde7acf\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606274\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1311639}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-06-07T01:02:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fba35207822bebd7260ae00a08b41f692182e39c"
  },
  {
    "id": 548,
    "message": "[floss] Remove Battery Manager observer\n\nWe need to remove all adapter observers on shut down.\n\nBUG=b:344972231\nTEST=build and device_unittests\n\nChange-Id: I93d878425e49e0d03c92bdf0d39adeca91dbd427\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5606006\nCommit-Queue: Michael Sun <michaelfsun@google.com>\nAuto-Submit: Katherine Lai <laikatherine@chromium.org>\nReviewed-by: Michael Sun <michaelfsun@google.com>\nCr-Commit-Position: refs/heads/main@{#1311638}",
    "author": "Katherine Lai",
    "author_email": "laikatherine@google.com",
    "date": "2024-06-07T01:01:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/918b5f850fe9d19e06c2b7820cf143bc3d3988e7"
  },
  {
    "id": 549,
    "message": "[TabGroupCollapse] Expand tabs on bulk ungroup\n\nExpand tabs on bulk ungroup.\n-Check if StripLayoutTab is collapsed instead of group when determining\n whether to expand. The group collapse state is updated before the rest\n of the tabs are removed.\n\nBug: 338332424\nChange-Id: I94b4f1310100bb64432ae98c02589be93d3db134\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5593600\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nReviewed-by: Sirisha Kavuluru <skavuluru@google.com>\nCommit-Queue: Neil Coronado <nemco@google.com>\nCr-Commit-Position: refs/heads/main@{#1311637}",
    "author": "Neil Coronado",
    "author_email": "nemco@google.com",
    "date": "2024-06-07T01:01:47+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/902ee7b0fdc860d6e32ba37b05e1dbaa239d294d"
  },
  {
    "id": 550,
    "message": "[chromecast] Enable chrome://inspect on debug builds in web runtime.\n\nBug: b/344966095\nTest: use chrome://inspect on a Cast device running Disney Plus\nChange-Id: I76d02df841c61b4d5bbcaa35732f1fe853b01c00\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5598705\nAuto-Submit: Simeon Anfinrud <sanfin@chromium.org>\nCommit-Queue: Luke Halliwell <halliwell@chromium.org>\nReviewed-by: Luke Halliwell <halliwell@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311636}",
    "author": "Simeon Anfinrud",
    "author_email": "sanfin@chromium.org",
    "date": "2024-06-07T00:59:51+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/24e870240c7063b808d3731fa23ee779fc46fd11"
  },
  {
    "id": 551,
    "message": "[BRP] Un-rewrite DrawQuad::shared_quad_state\n\nKudos to nuskos@ for finding the offending pointer!\n\nBug: 345298647, 331840472\nChange-Id: Ibffc24c1f0afdc15f13d4a0d7bf76c03b24b7080\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5604742\nCommit-Queue: Sunny Sachanandani <sunnyps@chromium.org>\nCommit-Queue: Bartek Nowierski <bartekn@chromium.org>\nReviewed-by: Sunny Sachanandani <sunnyps@chromium.org>\nAuto-Submit: Bartek Nowierski <bartekn@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1311635}",
    "author": "Bartek Nowierski",
    "author_email": "bartekn@chromium.org",
    "date": "2024-06-07T00:54:41+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/be9ffb8b041cacb32584e0d9790cdd9dc572ff82"
  },
  {
    "id": 555,
    "message": "Re-enable flaky test AttributionUponAttestationsLoading\n\nThis cl fixed the attestations file version due to recent pre-installed\nattestations change.\n\nWe also registered the source using attributionsrc instead of navigation\nto hopefully reduce the test time.\n\nBug: 355215758\nChange-Id: I9a511c0f0dc2b77b45a0d66594ea5c76701dfd48\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740360\nReviewed-by: Andrew Paseltiner <apaseltiner@chromium.org>\nCommit-Queue: Nan Lin <linnan@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333529}",
    "author": "Nan Lin",
    "author_email": "linnan@chromium.org",
    "date": "2024-07-26T15:08:21+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/f20f839616dc2eadb38af2b4b25ab2f9319c806f"
  },
  {
    "id": 556,
    "message": "Mark long-press-drag-drop-touch-editing test as slow\n\nBug: 354579691\nChange-Id: I9e5432d4a3ebb87dfb2cc7571bb78386481bcaf7\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738755\nReviewed-by: Robert Flack <flackr@chromium.org>\nCommit-Queue: Kevin Ellis <kevers@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333528}",
    "author": "Kevin Ellis",
    "author_email": "kevers@google.com",
    "date": "2024-07-26T15:05:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/138e9c403ed7c57dcdec9a1f5eb1c7cc108ed93a"
  },
  {
    "id": 557,
    "message": "Reset PrefChangeRegistrar in UpgradeDetector::Shutdown.\n\nUpgradeDetector::pref_change_registrar_ keeps a raw pointer to the\nlocal state's PrefService. To avoid triggering detection of a leaked\ndangling pointer, reset it during UpgradeDetector::Shutdown.\n\nBug: 355610030\nChange-Id: I40bbdc4005205f35e345508de8a10a1e3801b20f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5742343\nReviewed-by: Greg Thompson <grt@chromium.org>\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nCr-Commit-Position: refs/heads/main@{#1333527}",
    "author": "Jan Keitel",
    "author_email": "jkeitel@google.com",
    "date": "2024-07-26T15:03:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/2bfc82bd28e6b99c1c646d7577f4630cdbb02f35"
  },
  {
    "id": 558,
    "message": "spanification: Add `#pragma allow_unsafe_buffers` to ui/*\n\nSee `//docs/unsafe_buffers.md`\n\nThis is a preparation to fix each files.\nThis CL has no behavior changes.\n\nThis patch was fully automated using script:\nhttps://paste.googleplex.com/5614491201175552\n\nNote that in patchset2, change to:\n/build/config/unsafe_buffers_paths.txt\nwas reverted. Indeed, running too many (~3) CQ run touching this file is\nmaking the builder cache much slower. I will bundle every change to this\nfile in a subsequent CL. I will limit myself to 1-2 CQ run per day.\n\nSee internal doc about it:\nhttps://docs.google.com/document/d/1erdcokeh6rfBqs_h0drHqSLtbDbB61j7j3O2Pz8NH78/edit?resourcekey=0-hNe6w1hYAYyVXGEpWI7HVA&tab=t.0\n\nAX-Relnotes: n/a.\nBug: 40285824\nChange-Id: I008d48abd0b756a5c93a6c937b00679e5d2194f1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5717833\nReviewed-by: Nico Weber <thakis@chromium.org>\nAuto-Submit: Arthur Sonzogni <arthursonzogni@chromium.org>\nCommit-Queue: Arthur Sonzogni <arthursonzogni@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333526}",
    "author": "Arthur Sonzogni",
    "author_email": "arthursonzogni@chromium.org",
    "date": "2024-07-26T15:00:54+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/24d53e3e9a861c76d0e1626522c93e371f5d9223"
  },
  {
    "id": 559,
    "message": "Adjust destruction order of ProfileImpl members.\n\nPrefChangeRegistrar keeps a raw pointer to the PrefService it is\nobserving. As a result, it should be destroyed prior to the PrefService.\n\nThis CL does not yet modify the dangling annotation of the\nPrefRegistrar member because there seem to be (several) additional\ninstances in which the member becomes dangling.\n\nBug: 355610030\nChange-Id: I9f516e983b4c757b78d2dc302e0236ca1a12ecce\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743593\nReviewed-by: David Roger <droger@chromium.org>\nReviewed-by: Dominic Battré <battre@chromium.org>\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nCr-Commit-Position: refs/heads/main@{#1333525}",
    "author": "Jan Keitel",
    "author_email": "jkeitel@google.com",
    "date": "2024-07-26T15:00:03+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/bcb1754d6dd82ee448b3ce6ede8755c314959de6"
  },
  {
    "id": 560,
    "message": "Roll ios_internal from 501e67df31c7 to 5e4d598fc9be\n\nhttps://chrome-internal.googlesource.com/chrome/ios_internal.git/+log/501e67df31c7..5e4d598fc9be\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/ios-internal-chromium-autoroll\nPlease CC chrome-brapp-engprod@google.com,joemerramos@google.com,rkgibson@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nBug: None\nChange-Id: I74be47390446ff9a9d7efcbdfe63b0c5e223fb48\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5742522\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333524}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-26T14:57:11+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/719ca2d9993093caa31282269af19be49b3f9993"
  },
  {
    "id": 561,
    "message": "[lensoverlay] Show searchbox in lens search bubble.\n\nSelecting matches in the searchbox is non functional, however the\nsearchbox input and dropdown are functional. Feature flag is in\nlens_features.cc.\n\nScreencast: screencast/cast/NTQ1NTE0OTcxMDQ0MjQ5Nnw3ZjIxZTg3Mi1hYg\n\nBug: 328632272\nChange-Id: I6b37ff47b66f1ca14c2b8f9933262f685ab8d9cb\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5463167\nReviewed-by: Moe Ahmadi <mahmadi@chromium.org>\nCommit-Queue: Nihar Majmudar <niharm@google.com>\nCr-Commit-Position: refs/heads/main@{#1333523}",
    "author": "Nihar Majmudar",
    "author_email": "niharm@google.com",
    "date": "2024-07-26T14:56:44+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9c8e9823e7586361bec781dec3c40dad050e6399"
  },
  {
    "id": 562,
    "message": "Roll Chrome Linux PGO Profile\n\nRoll Chrome Linux PGO profile from chrome-linux-main-1721973548-72ec7edd41800a61cf7fbce7b7d39033ee7c32d2-321aaaf7ae9ca72b98adf17ed74aafafaf276da2.profdata to chrome-linux-main-1721995131-f4680ba138ea4278f4a631b7af1f07ed8e5366c5-1a765d8f98899564966915d2451bc73189822e5d.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-linux-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chrome\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I8df00fc1c7e8cfac386a77c4a0ca5f9a40ce32f9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743422\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333522}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-26T14:56:01+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/72cb923a154ef11d9417e18b0f8754c59e934adf"
  },
  {
    "id": 563,
    "message": "Append OAuth client ID to IssueToken's request headers\n\nX-OAuth-Client-ID: <OAuth2 client ID>\n\nFixed: b:341371239\nChange-Id: I04e4b6cefbba5cf9563215ead7c46a82e94d0a70\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741345\nAuto-Submit: Alex Ilin <alexilin@chromium.org>\nReviewed-by: David Roger <droger@chromium.org>\nCommit-Queue: David Roger <droger@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333521}",
    "author": "Alex Ilin",
    "author_email": "alexilin@chromium.org",
    "date": "2024-07-26T14:51:32+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/b4ced426ad28fe4336a1f3fc809daeb206654edd"
  },
  {
    "id": 564,
    "message": "Cleanup TrackingProtectionOnboardingDelegate\n\nThis CL removes the TrackingProtectionOnboardingDelegate, as we don't\nneed browser sides methods after the recent announcement.\n\nBug: b:355002866\n\nChange-Id: I2868d007c6b3a3343622eddafe2ef821a5e3e7a9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738417\nReviewed-by: Andrey Zaytsev <andzaytsev@google.com>\nCommit-Queue: Aashna Sheth <aashnas@google.com>\nReviewed-by: Abe Boujane <boujane@google.com>\nCr-Commit-Position: refs/heads/main@{#1333520}",
    "author": "Aashna Sheth",
    "author_email": "aashnas@google.com",
    "date": "2024-07-26T14:50:42+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/a7443d0942350783db65e3d4c13027f1c36ea5e2"
  },
  {
    "id": 565,
    "message": "Roll Perfetto from e9b88b1f4968 to 439b00c39be4 (2 revisions)\n\nhttps://android.googlesource.com/platform/external/perfetto.git/+log/e9b88b1f4968..439b00c39be4\n\n2024-07-26 ivankc@google.com Merge \"bigtrace: Add headless service minikube worker and modify cluster script\" into main\n2024-07-26 ivankc@google.com Merge \"bigtrace: Buffer results from workers instead of writing directly\" into main\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/perfetto-chromium-autoroll\nPlease CC perfetto-bugs@google.com,primiano@chromium.org on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:linux-perfetto-rel\nBug: None\nTbr: perfetto-bugs@google.com\nChange-Id: I33faf74a756c3ec6fa9ea2b599408dab5d2c5345\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743419\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333519}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-26T14:50:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/fc0061d8dd6dcfc042a7000a133d32b0b4001b17"
  },
  {
    "id": 566,
    "message": "🧇 Fix crash when \"More\" is clicked multiple times\n\nIf the user manages to click the \"More\" button on the choice screen\nmore than once before we scroll to the end and disable it, we end\nup attempting to run a base::OnceCallback more than once, which\ncauses a crash. With this fix, we will ignore the subsequent calls\nto the callback, which only results in not logging the fact that\n\"More\" was tapped more than once per choice screen.\n\nBug: 353573024\nChange-Id: Id6bb06b640c38455a2743ac82eef98d2da01de53\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741468\nCommit-Queue: Nicolas Dossou-Gbété <dgn@chromium.org>\nAuto-Submit: Nicolas Dossou-Gbété <dgn@chromium.org>\nReviewed-by: David Roger <droger@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333518}",
    "author": "Nicolas Dossou-Gbete",
    "author_email": "dgn@chromium.org",
    "date": "2024-07-26T14:48:23+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9386833aaf2e29dfa5407fc1a381ce29d9849593"
  },
  {
    "id": 567,
    "message": "Fix test flake in animation-event-destroy-renderer.\n\nThe ?animationiteration variant of the test could fail due to a stall\non the testing machine causing to animation to be finished before\nanimationiteration is fired. An animationiteration event does not fire\non the last iteration. Switched to a long duration animation with\nsetting the current time to trigger each event.\n\nTested locally with hundreds of repeats in virtual/threaded.\n\nBug: 355327241\nChange-Id: I823ea9c3cb22c299e2f8c78ed0e051305827f404\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740864\nReviewed-by: Robert Flack <flackr@chromium.org>\nCommit-Queue: Kevin Ellis <kevers@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333517}",
    "author": "Kevin Ellis",
    "author_email": "kevers@google.com",
    "date": "2024-07-26T14:46:33+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/546b4c5975d5a31c3638e049f3ee6e862185b3ab"
  },
  {
    "id": 568,
    "message": "Roll ANGLE from 48d19755a37e to a0a832de0a0c (7 revisions)\n\nhttps://chromium.googlesource.com/angle/angle.git/+log/48d19755a37e..a0a832de0a0c\n\n2024-07-26 geofflang@chromium.org Revert \"GL: Forward client-side arrays to the driver when possible\"\n2024-07-26 angle-autoroll@skia-public.iam.gserviceaccount.com Roll vulkan-deps from 49b5e420b19a to 068847956e41 (2 revisions)\n2024-07-26 angle-autoroll@skia-public.iam.gserviceaccount.com Roll Chromium from 15e30105c06e to 2e8f18571b3e (242 revisions)\n2024-07-26 mark@lunarg.com Tests: Add Shovel Knight Pocket Dungeon trace\n2024-07-25 angle-autoroll@skia-public.iam.gserviceaccount.com Manual roll Chromium from 99513aa4b41c to 15e30105c06e (813 revisions)\n2024-07-25 geofflang@chromium.org GL: Forward client-side arrays to the driver when possible\n2024-07-25 angle-autoroll@skia-public.iam.gserviceaccount.com Manual roll Chromium from 834617e372fd to 99513aa4b41c (90 revisions)\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/angle-chromium-autoroll\nPlease CC angle-team@google.com,solti@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in ANGLE: https://bugs.chromium.org/p/angleproject/issues/entry\nTo file a bug in Chromium: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chromium.try:android_optional_gpu_tests_rel;luci.chromium.try:linux_optional_gpu_tests_rel;luci.chromium.try:mac_optional_gpu_tests_rel;luci.chromium.try:win_optional_gpu_tests_rel;luci.chromium.try:linux-swangle-try-x64;luci.chromium.try:win-swangle-try-x86\nBug: None\nTbr: solti@google.com\nChange-Id: Ie9dcc8a476184d81d64e0591f999fc24a12e1590\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743417\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333516}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-26T14:44:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/c477e74aec2b9257a9f1fa6893976cf5e6fd340e"
  },
  {
    "id": 569,
    "message": "Delete Full3pcd Clank Onboarding Logic\n\nBug: 355002762\nChange-Id: I299f3350c6de77c60385337d7684fffc4a9d44b3\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740817\nAuto-Submit: Abe Boujane <boujane@google.com>\nReviewed-by: Calder Kitagawa <ckitagawa@chromium.org>\nReviewed-by: Matt Dembski <dembski@google.com>\nCommit-Queue: Calder Kitagawa <ckitagawa@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333515}",
    "author": "Abe Boujane",
    "author_email": "boujane@google.com",
    "date": "2024-07-26T14:40:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/089fe853bf036a4d7d64290c71923e5d5fe88307"
  },
  {
    "id": 578,
    "message": "[AutoReload] Add url paramter for autoreload\n\nThis CL appends a url paramter to the gaia url when an automatic reload of the authentication flow is executed.\nAs of now no specific functionality utilizes this url parameter but it will be used for future debugging purposes and for future proofing.\n\nBug: b/352025280\nChange-Id: I36cdf534ea9c54e6288374fa8909f4cd54223c3a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5729977\nCommit-Queue: Aya Elgendy‎ <ayag@chromium.org>\nReviewed-by: Maksim Ivanov <emaxx@chromium.org>\nReviewed-by: Andrey Davydov <andreydav@google.com>\nCr-Commit-Position: refs/heads/main@{#1333506}",
    "author": "Aya El Gendy",
    "author_email": "ayag@chromium.org",
    "date": "2024-07-26T14:07:08+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/589fb91ec454bb08f2309f40bd7a0141784f1bfa"
  },
  {
    "id": 570,
    "message": "Roll Chrome Win ARM64 PGO Profile\n\nRoll Chrome Win ARM64 PGO profile from chrome-win-arm64-main-1721973548-f166d4d2c53038380101e46c634993efba9bca99-321aaaf7ae9ca72b98adf17ed74aafafaf276da2.profdata to chrome-win-arm64-main-1721995131-f779e870da14a04a56a23f0dee44cb6c7e21bff7-1a765d8f98899564966915d2451bc73189822e5d.profdata\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://autoroll.skia.org/r/pgo-win-arm64-chromium\nPlease CC chrome-brapp-engprod@google.com,pgo-profile-sheriffs@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo file a bug in Chromium main branch: https://bugs.chromium.org/p/chromium/issues/entry\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nTbr: pgo-profile-sheriffs@google.com\nChange-Id: I8185782182109799c875a002baa01dac40a9697f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743285\nCommit-Queue: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nBot-Commit: chromium-autoroll <chromium-autoroll@skia-public.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333514}",
    "author": "chromium-autoroll",
    "author_email": "chromium-autoroll@skia-public.iam.gserviceaccount.com",
    "date": "2024-07-26T14:38:22+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/192a406ccc3f73be2a785c3da12cdfb784ad414c"
  },
  {
    "id": 571,
    "message": "Extend safe-browsing-hash-prefix flag\n\nWebView has not rolled out yet.\n\nFixed: 355308112\nChange-Id: I9b9c4e5ac519bc24342e9916ac9611fedbf04b5f\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740507\nAuto-Submit: Xinghui Lu <xinghuilu@chromium.org>\nReviewed-by: thefrog <thefrog@chromium.org>\nCommit-Queue: thefrog <thefrog@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333513}",
    "author": "Xinghui Lu",
    "author_email": "xinghuilu@chromium.org",
    "date": "2024-07-26T14:35:00+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4ef6847efb49a366322c2ce3bfc95ae7e876ac67"
  },
  {
    "id": 572,
    "message": "Switch riscv64 build to use stable r27 NDK\n\nBug: 353699308\nTest: CQ\nChange-Id: Ic87b2877e69ef6936cbb9043ef28053524a4581a\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5740908\nReviewed-by: Andrew Grieve <agrieve@chromium.org>\nReviewed-by: Nico Weber <thakis@chromium.org>\nAuto-Submit: Prashanth Swaminathan <prashanthsw@google.com>\nCommit-Queue: Nico Weber <thakis@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333512}",
    "author": "Prashanth Swaminathan",
    "author_email": "prashanthsw@google.com",
    "date": "2024-07-26T14:32:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/dc1044af4484b79db65f3cfb53f4b488f55e70d7"
  },
  {
    "id": 573,
    "message": "Field trial config for AutofillI18nAddressModelNewCountries.\n\nBug: b/41489892\nChange-Id: I45a7b1f892adf26670ea2a38213b6699d73d9fa8\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741284\nReviewed-by: Jan Keitel <jkeitel@google.com>\nCommit-Queue: Norge Vizcay <vizcay@google.com>\nCommit-Queue: Jan Keitel <jkeitel@google.com>\nAuto-Submit: Norge Vizcay <vizcay@google.com>\nCr-Commit-Position: refs/heads/main@{#1333511}",
    "author": "Norge Vizcay",
    "author_email": "vizcay@google.com",
    "date": "2024-07-26T14:31:35+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/419c75b3b4c79e1095f516e9ee12e21e2a39c992"
  },
  {
    "id": 574,
    "message": "[Mullhet] Split kGrouped MatchedFormType.\n\nBefore this CL, PasswordManager.PotentialBestMatchFormType didn't\ndifferentiate between grouped Android and other credentials.\nThis CL splits the MatchedFormType::kGrouped bucket into 2\nseparate buckets. The existing bucket gets deprecated.\n\nBug: 332673987\nChange-Id: I461b9e752c7f10cd3bcabcb62276a27719234889\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735375\nReviewed-by: Vasilii Sukhanov <vasilii@chromium.org>\nCommit-Queue: Timofey Chudakov <tchudakov@google.com>\nCr-Commit-Position: refs/heads/main@{#1333510}",
    "author": "Timofey Chudakov",
    "author_email": "tchudakov@google.com",
    "date": "2024-07-26T14:29:26+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/554c994dd740d5345ae6f279a4ef258dfe5a6b54"
  },
  {
    "id": 575,
    "message": "DocumentRules: Handle base URL change in ScriptLoader\n\nMoves some of the logic handling base URL changes to ScriptLoader, to\nallow it to update |speculation_rule_set_| to the new rule set\ncreated with the updated base URL.\n\nBug: 353099744\nChange-Id: I3e0ec98f9a57ccac98555a21d9fcf03e913506b9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5738414\nReviewed-by: Jeremy Roman <jbroman@chromium.org>\nCommit-Queue: Adithya Srinivasan <adithyas@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333509}",
    "author": "Adithya Srinivasan",
    "author_email": "adithyas@chromium.org",
    "date": "2024-07-26T14:26:37+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/65a74c3b16d7638e62a4e16fb5676b191265c195"
  },
  {
    "id": 576,
    "message": "Roll devtools-internal from 894b13578bf0 to 5425428d1f4b (1 revision)\n\nhttps://chrome-internal.googlesource.com/devtools/devtools-internal.git/+log/894b13578bf0..5425428d1f4b\n\nAlso rolling transitive DEPS:\n  https://chromium.googlesource.com/devtools/devtools-frontend.git/+log/dee5906e292296195923cfc3f69cc6cdb8c58829..2c1d92ad7b1a4b5247b09ae1709a2d99da76e6af\n\nIf this roll has caused a breakage, revert this CL and stop the roller\nusing the controls here:\nhttps://skia-autoroll.corp.goog/r/devtools-internal-chromium\nPlease CC devtools-waterfall-sheriff-onduty@rotations.google.com,liviurau@google.com on the revert to ensure that a human\nis aware of the problem.\n\nTo report a problem with the AutoRoller itself, please file a bug:\nhttps://issues.skia.org/issues/new?component=1389291&template=1850622\n\nDocumentation for the AutoRoller is here:\nhttps://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md\n\nCq-Include-Trybots: luci.chrome.try:linux-chromeos-chrome\nBug: None\nTbr: devtools-waterfall-sheriff-onduty@rotations.google.com\nChange-Id: Icaf4cb9356e17ba0e9e7138238b9ea4bb83debca\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5742968\nCommit-Queue: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nBot-Commit: chromium-internal-autoroll <chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com>\nCr-Commit-Position: refs/heads/main@{#1333508}",
    "author": "chromium-internal-autoroll",
    "author_email": "chromium-internal-autoroll@skia-corp.google.com.iam.gserviceaccount.com",
    "date": "2024-07-26T14:23:43+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/4b29d609a8be45f439bb2ea7d370ff1ebcb61ac6"
  },
  {
    "id": 577,
    "message": "[Passwords] Extend expiring histograms.\n\nFixed: 350504411, 354713400, 354713638, 354713483\nChange-Id: Ifc1364dd4788912fe26d793ffbc00b45d1c0465c\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743615\nCommit-Queue: Maria Kazinova <kazinova@google.com>\nAuto-Submit: Maria Kazinova <kazinova@google.com>\nReviewed-by: Ioana Pandele <ioanap@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333507}",
    "author": "Maria Kazinova",
    "author_email": "kazinova@google.com",
    "date": "2024-07-26T14:21:28+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/9b07a944b4784ac593f8d6deb63b270ef9f01d9d"
  },
  {
    "id": 579,
    "message": "[Android] Rename address sheet keyboard accessory actions.\n\nBug: 327838324\nChange-Id: I96d8476e8f0872dfb8d7e5ca1c9d6cbda59c3ca9\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741546\nCommit-Queue: Timofey Chudakov <tchudakov@google.com>\nReviewed-by: Friedrich Horschig <fhorschig@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333505}",
    "author": "Timofey Chudakov",
    "author_email": "tchudakov@google.com",
    "date": "2024-07-26T14:01:55+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3a0e070ba2732e0664f64de691d0af941a5b1ed0"
  },
  {
    "id": 580,
    "message": "[Gardener] Disable flaky UkmBrowserTest\n\nUkmBrowserTest.RegularPlusIncognitoCheck has caused failure on\nandroid-oreo-x86-rel bot.\nUkmBrowserTest.InstallDateProviderPopulatesSystemProfile has caused\nfailure on android-bfcache-rel bot.\n\nBug: 355609356\nChange-Id: I6bfec433d155f12fe20fe4df2134a23510509cd1\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5743594\nAuto-Submit: Anna Tsvirchkova <atsvirchkova@google.com>\nCommit-Queue: Anna Tsvirchkova <atsvirchkova@google.com>\nReviewed-by: Adem Derinel <derinel@google.com>\nOwners-Override: Anna Tsvirchkova <atsvirchkova@google.com>\nCommit-Queue: Adem Derinel <derinel@google.com>\nCr-Commit-Position: refs/heads/main@{#1333504}",
    "author": "Anna Tsvirchkova",
    "author_email": "atsvirchkova@google.com",
    "date": "2024-07-26T14:01:46+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/e4622e3f77d6e92c40edd01ff2e4c981985ef354"
  },
  {
    "id": 581,
    "message": "[ios] Use a single gn variable to control use of copy_xctrunner_app\n\nBuilding XCTests requires copying XCTRunner.app which is part of the iOS\nSDK (and shipped inside Xcode.app) into the application. When using the\nsystem installation of Xcode, those files are outside of the checkout.\nUsing absolute path works with gn, however the distributed build system\nrequires that all paths are relative to the checkout. This is faked by\nusing symbolic links to the SDK inside of Xcode. Additionally, each build\ndirectory may use a distinct version of Xcode (e.g. to build with beta),\nso the symlink needs to be present in the $root_build_dir. However, when\ndoing that, we need to list inputs pointing to file in $root_build_dir,\nand gn requires all failes in $root_build_dir to be listed as outputs of\nanother target.\n\nTo fullfill all of those requirements, we 1. create symlinks pointing to\nthe SDK files in Xcode, 2. declare a target listing the files as outputs\n(the target is a script that does nothing, it only pretends to created\nthe files but they already exists).\n\nThis works, but results in some files in $root_build_dir being links to\nfiles outside of the build directory. Running `ninja -t clean` will try\nto delete those files breaking Xcode installation. The recommendation is\nto use `gn clean` or `ninja -t cleandead` instead.\n\nAdd a single variable to control whether the workaround is needed.\n\nBug: none\nChange-Id: I0686232ebf59e9f39dec7cd73e464969479f5122\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741540\nCommit-Queue: Sylvain Defresne <sdefresne@chromium.org>\nReviewed-by: Rohit Rao <rohitrao@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333503}",
    "author": "Sylvain Defresne",
    "author_email": "sdefresne@chromium.org",
    "date": "2024-07-26T13:56:24+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/948b92504b84cf9b3f091052e708ace2cfae5651"
  },
  {
    "id": 582,
    "message": "[sync] Remove CreateDataTypeManager\n\nNo behavioral changes outside tests, as it always instantiates\nDataTypeManagerImpl.\n\nIn tests, before this class, a subclass was instantiated for the purpose\nof accessing some internal state. Instead, SyncEngine can be used for\nsimilar purposes, and everything else isn't externally visible and\narguably shouldn't be verified in tests.\n\nOne benefit is that SyncApiComponentFactory has a better-defined\nscope, which is dealing with SyncEngine instances. A TODO is added\nto find a better name for this class and make it less abstract.\n\nChange-Id: Ia54821245f07f09c49bb0c3d5dc595d1ac61bf0a\nBug: 335688372\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741644\nCode-Coverage: findit-for-me@appspot.gserviceaccount.com <findit-for-me@appspot.gserviceaccount.com>\nCommit-Queue: Mikel Astiz <mastiz@chromium.org>\nReviewed-by: Marc Treib <treib@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333502}",
    "author": "Mikel Astiz",
    "author_email": "mastiz@chromium.org",
    "date": "2024-07-26T13:55:39+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/d66d47c65b5180387e321d05bffcf37be1d9112a"
  },
  {
    "id": 583,
    "message": "[Profiles] Profile picker no longer navigates in browser being destroyed\n\nBug: 40064092, 40242414\nChange-Id: Id8283b435a99254788225748800d7fec409fb9c6\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5741701\nReviewed-by: Greg Thompson <grt@chromium.org>\nCommit-Queue: David Roger <droger@chromium.org>\nCr-Commit-Position: refs/heads/main@{#1333501}",
    "author": "David Roger",
    "author_email": "droger@chromium.org",
    "date": "2024-07-26T13:48:25+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/1cd71739a1661436a24c9b8ea057dc9061e73ef0"
  },
  {
    "id": 584,
    "message": "[High5] ActiveSessionAuthController unittests\n\nWe add several unittests that test the behavior of\n`ActiveSessionAuthController`. We assert that it behaves correctly in\nthe case of correct and wrong password/pin inputs, and in the case of\ncanceling the dialog.\n\nBug: b:352238958, b:348326316\nChange-Id: I141d45f932ad9884253480e578c413ec61d948ab\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/5735972\nReviewed-by: Xiyuan Xia <xiyuan@chromium.org>\nReviewed-by: Maksim Ivanov <emaxx@chromium.org>\nReviewed-by: Hardik Goyal <hardikgoyal@chromium.org>\nCommit-Queue: Elie Maamari <emaamari@google.com>\nCr-Commit-Position: refs/heads/main@{#1333500}",
    "author": "Elie Maamari",
    "author_email": "emaamari@google.com",
    "date": "2024-07-26T13:41:50+01:00",
    "url": "https://api.github.com/repos/chromium/chromium/git/commits/3d5950913dbbd130539cca48ada2812498e5cf48"
  }
]
```
# Swagger Documentation
The swagger documentation can be found here: 
http://localhost:8082/swagger-ui/index.html

## Acknowledgements
 - Github: https://github.com

