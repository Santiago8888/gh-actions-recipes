# GitHub Client

## Features

- [ ] Get time when Branch was Opened, PR created and Merged.
- [ ] Get number of commits during each step.
- [ ] Store results in DB.
- [ ] Get Activity counts as in total numner of PRs & Merged branches.
- [ ] Compare across GitHub Repos.

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

* Define DB Model. ()
* Fetch Closed PRs and Filter relevant Stats. ()
* Store PR Stats on MongoDB. ()
* Design GET Queries on Stitch Console. ()
* Create PR and Document Front End Tasks. ()

## Useful Links

**API Libraries**:

* [GitHub V3 Client](https://github.com/google/go-github)
* [GraphQL Client](https://github.com/shurcooL/githubv4)
* [Sample PR Response](https://api.github.com/repos/cypress-io/cypress/pulls/8072)
