On branch create. Write to DB. [MVP]

On branch push. Write to DB. [MVP]
On branch push, after build step. Write to DB.
On branch push, after unit tests. Write to DB.

On pull-request to master. Write to DB. [MVP]
On pull-request to master after integration tests. Write to DB.
On pull-request to master after E2E tests. Write to DB.

On push to master. Write to DB. [MVP]


Acceptance Tests:
1. Write to MongoDB from actions directory. (Done)
2. Write to MongoDB from pipeline run. (Done)
3. Conclude write step after push. (Done)
4. Develop write request on pull request. (Done)
5. Write request on create branch. (Done)
6. Write request on merge. (Done)
7. Create issue on branch create. (Done)
8. Request metrics on merge. (Done)
9. Write metrics on merge. (Done)
10. Close issue on merge. (Done)


Data Model (MVP): {
    id: <uuid>,
    branch: <string>,
    hook: <enum> (create|push|pull_request|merge),
    time: <date>
}


Metrics MVP:

    TIME/ COMMIT Count:
    1. From open to merge.
    2. From PR to merge.
    3. From open to PR.



Useful Links:
On getting GIT info for merge:
https://octokit.github.io/rest.js/v18#actions # Get a reference


On master exclude when pushing:
https://stackoverflow.com/a/57903434


TODO:
1. Split words on issue title & capitalize.
2. Format issue body: add link to author & new line.
3. Create metrics label & attach to issue.

4. Implement strategy when switching branches.
