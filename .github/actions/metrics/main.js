const MongoClient = require('mongodb').MongoClient;
const github = require("@actions/github");
const core = require("@actions/core");


const COLLECTION = 'metrics';

const dbName = 'actions';
const cluster = 'cluster0-8pgr7.mongodb.net';
const config = 'retryWrites=true&w=majority';
const pwd = core.getInput("mongo-password")
const uri = `mongodb+srv://Admin:${pwd}@${cluster}/${dbName}?${config}`;

async function run() {
    try {
        console.log('Context: ', github.context);
        console.log('Context: ', github.context.payload.pull_request);

        // console.log('Keys: ', Object.keys(github.context))
        // console.log('Repo: ', github.context.repo);
        // console.log('Payload: ', github.context.payload)
        // console.log('Commits: ', github.context.payload.commits);

        const repository = github.context.repo.repo;
        const branch = github.context.ref.replace('refs/heads/', '');
        const author = github.context.actor;
        const owner = github.context.repo.owner;
        const commit = github.context.payload.head_commit || ''
        const message = commit.message;

        const issue = github.context.payload.number || null;
        console.log('Issue: ', issue);

        const isNewBranch = github.context.payload.created || false;
        const isOpened = github.context.action === 'opened' || false;
        const isPullRequest = github.context.eventName === 'pull_request' || false;

        const client = new MongoClient(uri, { useNewUrlParser: true });
        client.connect(_ => {
            const collection = client.db(dbName).collection(COLLECTION);
            const record = {
                repository: repository,
                author: author,
                branch: branch,
                is_created: isNewBranch,
                is_opened: isOpened,
                is_pull_request: isPullRequest,
                time: new Date()
            }

            console.log('Record: ', record);
            collection.insertOne(record);

            client.close();
        });

        if (isNewBranch && isPullRequest){
            const title = branch;
            const body = `Opened by ${author}, with message : ${message}`;

            const token = core.getInput("repo-token");
            const octokit = github.getOctokit(token);
    
            await octokit.issues.create({
                repo: repository,
                owner: owner,
                title: title,
                body: body
            });
        }
    } catch (err) {
        core.setFailed(err.message);
    }
}


run();
