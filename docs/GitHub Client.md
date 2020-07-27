# GitHub Client

## Features

- [ ] Know how likely a PR is going to be approved.
- [ ] The approximate time for a PR to be resolved.
- [ ] An repository activity indicator with trends.
- [ ] Comparisson of stats across repositories.
- [ ] An agility score for good practices.

## RoadMap

* Hello World Request. (Done)
* Define Strategy to Retrieve Results. (Done)
* Connect to MongoDB Client.
* Store Results to MongoDB.
* Design CRON Pipeline to Fetch Results.
* Draft Article from Insights.
* Clean Data for Front-end.
* Authenticate request to GitHub API.

### User Stories

1. User can see the average time it takes to close a PR.
2. User sees % of Accepted PRs & time it takes on them.
3. User see trend on time and % of Acceptance with moving average.
4. User can see number of PRs merged and closed on defined time intervals, one week & one month, serverless?
5. User sees activity trend by rolling average.
6. User can compare up to five repositories in the above metrics.
7. The repository gets a score, based on the metrics.

### Tasks

* Define DB Model. (Done)
* Fetch Closed PRs and Filter relevant Stats. ()
    * Create Struct. ()
    * Retrieve Stats. ()
    * Compute Additionals. ()
        * Handle time stamps. ()
    * Instantiate Stats Object. ()
* Store PR Stats on MongoDB. ()
* Design GET Queries on Stitch Console. ()
* Create PR and Document Front End Tasks. ()

### DB Model

#### Collection : *Stats*

```
"_id": <BSON Object>
"owner": <string>
"repository": <string>

"number": <int>,
"state": <enum> || "closed",
"merged": <bool>,
"title": <string>,
"created_at": <dateTime>,
"closed_at": <dateTime>,
"author_association": <enum>,

"assignees_count": <int>,
"requested_reviewers_count": <int>
"comments": <int>, 
"review_comments": <int>,
"maintainer_can_modify": <bool>,
"commits": <int>,
"additions": <int>,
"deletions": <int>,
"changed_files": <int>,

"time_diff": <int>,
"lines_diff": <int>
```

## Useful Links

**API Libraries**:

* [GitHub V3 Client](https://github.com/google/go-github)
* [GraphQL Client](https://github.com/shurcooL/githubv4)
* [Sample PR Response](https://api.github.com/repos/cypress-io/cypress/pulls/8072)
