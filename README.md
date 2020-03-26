# temporal-go-demo
Temporal demo for GO meetups

# GO Temporal Sample for GO Meetup
This package contains a sample demoed at [Go's 10th Anniversary Seattle Meetup](https://www.meetup.com/golang/events/265858683/)

More Temporal info at:

* [Temporal Service](https://github.com/temporalio/temporal)
* [Temporal Go SDK](https://github.com/temporalio/temporal-go-sdk)
* [Temporal Java SDK](https://github.com/temporalio/temporal-java-sdk)

## Overview of the Samples

### Money Transfer Sample

Demonstrates a simple transfer from one account to another. 

## Get the Samples

Run the following commands:

      git clone git@github.com:temporalio/temporal-go-demo.git
      cd temporal-go-demo

## Build the Samples

      make

## Run Temporal Server

Run Temporal Server using Docker Compose:

    curl -O https://raw.githubusercontent.com/uber/cadence/master/docker/docker-compose.yml
    docker-compose up

If this does not work, see the instructions for running Temporal Server at https://github.com/uber/cadence/blob/master/README.md.

## Install Temporal CLI

[Command Line Interface Documentation](https://docs.temporal.io/docs/08_cli)

## See Temporal UI

The Temporal Server running in a docker container includes a Web UI.

Connect to [http://localhost:8088](http://localhost:8088).

Enter the *samples* domain. You'll see a "No Results" page. After running any sample, change the 
filter in the
top right corner from "Open" to "Closed" to see the list of the completed workflows.

Click on a *RUN ID* of a workflow to see more details about it. Try different view formats to get a different level
of details about the execution history.

## Run the samples

Each sample has specific requirements for running it. The following sections contain information about
how to run each of the samples after you've built them using the preceding instructions.


### Money Transfer Sample

Workflow Worker:
```
./bins/workflowWorker
```
Activities Worker:
```
./bins/activityWorker
```
Initiate Transfer:
```
tctl wf start --tl transfer --wt transfer --et 1200  --if ./transferRequest.json
```
Initiate Transfer Specifying Id of "T1":
```
tctl wf start --w T1 --tl transfer --wt transfer --et 1200  --if ./transferRequest.json
```
Initiate Transfer Specifying Id of "T1" and "reject duplicate workflow ID" ID reuse policy:
```
tctl wf start --w T1 --wrp 1 --tl transfer --wt transfer --et 1200  --if ./transferRequest.json
```

### Batch Transfer Sample

Workflow Worker:
```
./bins/workflowWorker
```
Activities Worker:
```
./bins/activityWorker
```

Initiate Batch Transfer:
```
tctl wf start -w batch-transfer-1 --tl transfer --wt batch-transfer --et 1200 --if ./batchTransferRequest.json
```

Send Signal For First Withdrawal:
```
tctl wf signal -w batch-transfer-1 --name withdraw --if ./withdrawSignalPayload1.json
```

Send Signal For Second Withdrawal:
```
tctl wf signal -w batch-transfer-1 --name withdraw --if ./withdrawSignalPayload2.json
```

Send Signal For Third Withdrawal:
```
tctl wf signal -w batch-transfer-1 --name withdraw --if ./withdrawSignalPayload3.json
```

Query Withdrawal Count:
```
tctl wf query -w batch-transfer-1 --qt get-count
```

Query Balance:
```
tctl wf query -w batch-transfer-1 --qt get-balance
```