const MongoClient = require('mongodb').MongoClient;
const github = require("@actions/github");
const core = require("@actions/core");
const moment = require('moment');

const OPENED = 'opened';
const CLOSED = 'closed';
const COLLECTION = 'metrics';

const dbName = 'actions';
const cluster = 'cluster0-8pgr7.mongodb.net';
const config = 'retryWrites=true&w=majority';
const pwd = core.getInput("mongo-password");
const uri = `mongodb+srv://Admin:${pwd}@${cluster}/${dbName}?${config}`;
const TOKEN = "repo-token";

const refPrefix = 'refs/heads/';
async function run() {
    try {
        console.log('Context: ', github.context);
        // console.log('Pull Request: ', github.context.payload.pull_request);
        // console.log('Keys: ', Object.keys(github.context))
        // console.log('Repo: ', github.context.repo);
        // console.log('Payload: ', github.context.payload)
        // console.log('Commits: ', github.context.payload.commits);

        const repository = github.context.repo.repo;
        const author = github.context.actor;
        const owner = github.context.repo.owner;

        const action = github.context.payload.action;
        const pull_request = github.context.payload.pull_request;
        const branch = pull_request 
            ? pull_request.head.ref
            : github.context.ref.replace(refPrefix, '');

        // const commit = github.context.payload.head_commit || {};
        // const message = commit.message || '';
        const issue = github.context.payload.number || null;
        const isNewBranch = github.context.payload.created || false;
        const isOpened = action === OPENED || false;
        const isClosed = action === CLOSED;
        const isMerged = isClosed && pull_request.merged;

        // Switch connection outside of try block & use finally.
        const client = await MongoClient.connect(uri, { useNewUrlParser: true });
        const collection = client.db(dbName).collection(COLLECTION);
        const record = {
            repository: repository,
            author: author,
            branch: branch,
            is_created: isNewBranch,
            is_opened: isOpened,
            is_merged: isMerged,
            time: new Date()
        };

        console.log('Record: ', record);
        collection.insertOne(record);

//        if (isMerged){
        if (true){
            const events = await collection.find({branch}).toArray();
            const createEvent = events.find(({ is_created }) => is_created) || {};
            const openEvent = events.find(({ is_opened }) => is_opened) || {};
            console.log('Events: ', events);
    
            const now = moment(new Date());
            const createTime = moment(createEvent.time);
            const openTime = moment(openEvent.time);
    
            const timeToOpen = moment.duration(openTime.diff(createTime)).humanize();
            const timeOpen = moment.duration(now.diff(openTime)).humanize();
            const timeToMerge = moment.duration(now.diff(createTime)).humanize();
    
            const commitsToOpen = events.filter(({ time }) => time < openTime).length;
            const commitsWhileOpen = events.filter(({ time }) => time > openTime).length;
            const totalEvents = events.length;


            const title = `**Pull Request Metrics:** \n\n`;

            const timeToOpenMetric = `Time to Open Pull Request: ${timeToOpen}.\n`;
            const timeOpenMetric = `Time to Merge Pull Request: ${timeOpen}.\n`;
            const timeToMergeMetric = `**Total Time to Merge Branch**: ${timeToMerge}.\n\n`;
            const timeMetrics = `${timeToOpenMetric}${timeOpenMetric}${timeToMergeMetric}`;

            const commitsToOpenMetric = `Number of Commits to Open Pull Request: ${commitsToOpen}.\n`;
            const commitsWhileOpenMetric = `No. of Commits while Pull Request was Open: ${commitsWhileOpen}.\n`;
            const totalEventsMetric = `**Total Events to Merge Branch**: ${totalEvents}.\n\n`;
            const counterMetrics = `${commitsToOpenMetric}${commitsWhileOpenMetric}${totalEventsMetric}`;

            const body = `${title} ${timeMetrics} ${counterMetrics}`;
            console.log('Body: ', body);


            const token = core.getInput(TOKEN);
            const octokit = github.getOctokit(token);
            await octokit.issues.createComment({
                repo: repository,
                owner: owner,
                issue_number: 6,
                // issue_number: issue,
                body: body
            });
        }

        client.close();
    } catch (err) {
        console.error(err);
        core.setFailed(err.message);
    }
}


run();
