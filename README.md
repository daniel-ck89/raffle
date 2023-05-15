# Raffle

**Raffle** is a rollup chain for Celestia. Anyone can create a Raffle and proceed fairly on the blockchain. Anyone can query the results easily.

## Setup
### 1. Install Ignite
```curl https://get.ignite.com/cli | bash```

- https://docs.ignite.com/

### 2. clone raffle & configure init sciprt
```
git clone https://github.com/daniel-ck89/raffle.git
cd raffle
git checkout v0.1.0
vi init.sh

> On line 49, update base_url to your celestia light node's gateway url.
```

### 3. Run init script
`./init.sh`

- Simply run this script to start the chain.

## Cmd list
- `raffle tx raffle create-simple-raffle`
- `raffle tx raffle start-simple-raffle`
- `raffle q raffle raffles`
- `raffle q raffle raffle`
- `raffle q raffle raffle-result`

## Usage
### 1. Create a raffle
```raffled tx raffle create-simple-raffle "Event1" "This is a custom keyboard limited edition raffle." "https://docs.google.com/spreadsheets/d/1jEnEBlYIf_GfQs32WfaCJJNtrylpfQdL_2NVtegtPGM/edit?usp=sharing" 10 25 --from raffle-key-1 --broadcast-mode block --keyring-backend test --node http://localhost:27657```
- If raffle creation is successful, Id (raffle_id) will be returned.
#### Usage :
```raffled tx raffle create-simple-raffle [title] [description] [participant-list-url] [number-of-winners] [number-of-participants] [flags]```
- title : Raffle's title. ( max 256 characters )
- description : Raffle's description. ( max 1024 characters )
- participant-list-url : The url of the site contains the list of participants. Give each participant a unique id in sequential order, starting from 0. `ex) https://docs.google.com/spreadsheets/d/1jEnEBlYIf_GfQs32WfaCJJNtrylpfQdL_2NVtegtPGM/edit#gid=0`
- number-of-winners : This is the number of winners. ( max winners 10,000 )
- number-of-participants : The number of participants participating in the raffle. ( max participants 30,000 )


&nbsp;&nbsp;&nbsp;
### 2. Query your raffle when after your tx is processed successfully
```raffled q raffle raffle 0 --output json --node http://localhost:27657```

- Response
  - ```{"raffle":{"creator":"raffle187p7dxsk3s4fujhx2q7a8ftd35fj2jmx97kfu0","id":"0","status":0,"title":"Event1","description":"This is a custom keyboard limited edition raffle.","participantListUrl":"https://docs.google.com/spreadsheets/d/1jEnEBlYIf_GfQs32WfaCJJNtrylpfQdL_2NVtegtPGM/edit?usp=sharing","numberOfWinners":10,"numberOfParticipants":25}}```
- where status is the current status of raffle. 0 means raffle has not yet exited.

#### Usage :
```raffled q raffle raffle [raffle_id] [flags]```

If the raffle_id is not returned due to broadcast timeout in the previous step, you can find your raffle by querying raffles.

```raffled q raffle raffles --reverse --output=json --node http://localhost:27657```


&nbsp;&nbsp;&nbsp;
### 3. Start raffle
```raffled tx raffle start-simple-raffle 0 --from raffle-key-1 --keyring-backend test --node http://localhost:27657```
- Proceed to raffle. 
#### Usage :
```raffled tx raffle start-simple-raffle [raffle_id] [flags]```


&nbsp;&nbsp;&nbsp;
### 4. Query the raffle result
```raffled q raffle raffle-result 0 --output json --node http://localhost:27657```

- Response
  - ```{"raffleResult":"{"result":[1,5,7,8,9,10,15,16,17,21]}"}```
- This returns the participants_id(in the docs submitted through participant-list-url) of the winners.

#### Usage :
```raffled q raffle raffle-result [raffle_id] [flags]```



