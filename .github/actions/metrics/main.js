const MongoClient = require('mongodb').MongoClient;
const github = require("@actions/github");
const core = require("@actions/core");


const dbName = 'actions';
const cluster = 'cluster0-8pgr7.mongodb.net';
const config = 'retryWrites=true&w=majority';
const pwd = core.getInput("mongo-password")

const uri = `mongodb+srv://Admin:${pwd}@${cluster}/${dbName}?${config}`;
const client = new MongoClient(uri, { useNewUrlParser: true });
const COLLECTION = 'metrics';

const CREATE = 'CREATE';
const PUSH = 'PUSH';
const PULL = 'PULL';
const MERGE = 'MERGE';
const hooks = [ CREATE, PUSH, PULL, MERGE ]


async function run() {
    try {
        console.log('Context: ', github.context);

        const hook = github.context.event;
        client.connect(_ => {
            const collection = client.db(dbName).collection(COLLECTION);
            collection.insertOne({
                repository: github.context.repository,
                branch: github.context.ref,
                hook: github.context.event,
                time: new Date()
            })

            client.close();
        });

        if (hook === CREATE ){
            const issueTitle = "Provisional Title.";
            const jokeBody = "Provisional Body.";
            const token = core.getInput("repo-token");
            const octokit = github.getOctokit(token);
    
            await octokit.issues.create({
                repo: github.context.repo.repo,
                owner: github.context.repo.owner,
                title: issueTitle,
                body: jokeBody
            });    
        }
    } catch (err) {
        core.setFailed(err.message);
    }
}


run();
