
# SavannahTech



## Description
This is a project that retrieves commit data about a specified GitHub repository, the application provides three main features for retrieving repository information.
## Key Features

&nbsp;
__Get Top N Commit Authors__: This feature returns the top N Committers for the given GitHub Repository.\
&nbsp;
__Get All Commits for Repository__: This feature retrieves all commits for the given GitHub Repository. \
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
`AUTH_TOKEN` (gitHub access token)
`FETCH_DATE_SINCE`(Format: MM-DD-YYYY)



## Getting Started
## Usage
To Run the project, navigate to the projects root directory and do the following:
1. Clone the project from the repository [here](https://github.com/djfemz/savannah_Tech)

2. Before running any commands to build/run the application, make sure that all environment variables in the Environment Variables section [here](#environment-variables) are filled in, especially the `AUTH_TOKEN` (gitHub access token), variable.

3. Ensure that docker is up and running on your computer.
4. If building or running the application the first time, execute the following command:
```shell
docker network create fullstack
```
5. execute the following command to build, test and start the application:
```shell
docker-compose up --build
```

## Testing
- The test suites for all the components of the application are located in the controllers, repositories and services directories in the api directory. The tests are executed as a part of the build process, so running the command below, will execute the tests and the build will fail if any test fails:
```shell
docker-compose up --build
``` 
## Technologies Used
Golang
## API Documentation

### Commits

#### GET http://localhost:8082/commits?size=3

- Get Top N commits
  **Request Sample**
```shell
curl --location 'localhost:8082/commits/authors/top?size=3' \
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


#### GET
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
- Get Commits for specified repository

```shell
curl -X 'GET' \
  'http://localhost:8082/api/v1/commits/chromium' \
  -H 'accept: application/json'
```

**Response Examples:**
```json
[
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
  }
]
```
# Swagger Documentation
The swagger documentation can be found here:
http://localhost:8082/swagger-ui/index.html

## Acknowledgements
- Github: https://github.com

