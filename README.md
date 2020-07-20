# GitHub Actions Recipes

**Github Actions** automate all software workflows, with world-class CI/CD. Build, test, and deploy your code right from GitHub.

Using real-world scenarios, a variety of workflows are explored. For the first project a bot to measure and improve Software Productivity was developed. The bot has an accompanying Angular application to visualize the reported metrics.

Included in the application we see examples of the following *Github Actions* features:

- [ ] Multplie workflow triggers (push & pull request) filtered for specific branches and event types.
- [ ] Configure and pass secrets to actions ( individual tasks that you can combine to create jobs and customize your workflow).
- [ ] Develop and call you own custom actions using Node.
- [ ] Make a comment to a GitHub issue using the OctoKit client (official GitHub API client).
- [ ] Connect to MongoDB to preserve and retrieve the state of your workflows.

The Angular dashboard is a work in progress and being developed using [ngx-charts](https://swimlane.gitbook.io/ngx-charts)

## RoadMap

- [X] Gather Code-lifcycle metrics.
- [ ] Build and deploy a Software Productity Dashboard.
- [ ] Gather metrics for code coverage.
- [ ] Reuse dashboard & metrics to compare other GitHub repos.
- [ ] Export Action to the marketplace.

## Media

[Software Development Metrics automation using GitHub Actions](https://medium.com/@santiagoq/software-development-metrics-automation-using-github-actions-30a51fd56df0)

Benefits discovered from tracking code-lifecyle metrics:

1. Shorter more frequent new features commits.
2. Better product planning and documentation.
3. Increased commitment to contribute.

