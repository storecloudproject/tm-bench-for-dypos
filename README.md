# Storecoin Dynamic Proof-of-Stake (DyPoS) Consensus Engine

## Introduction

Storecoin Dynamic Proof-of-Stake (DyPoS) consensus engine is built partially on top of [Tendermint](https://tendermint.com/). Tendermint is a Proof-of-Stake (PoS) based Byzantine Fault Tolerant (BFT) state machine replication engine. Storecoin DyPoS uses Tendermint as its foundation. It consists of 6 interdependent engines. Each engine is based on similar dynamic supply and demand principles as Uber Surge Pricing. Prices for all six engines are driven by real-time threat levels in the **Storecoin Security System.**

Governance in DyPoS is inspired by the checks and balances of the U.S. Constitution. Govnodes – a new type of Worker – ensure that Governance is decentralized.

## 6 Engines of DyPoS

The six engines of DyPoS are as follows.

### Dynamic (Free) Transactions

To secure transactions, a small security bond is required with each transaction that is returned to the sender, with interest, once the transaction is determined by dGuards to not be spam or an attack. The transactions are permanently free. Merchants and app developers who use Storecoin as their payments platform never pay fees to have their transactions processed on Storecoin.

### Dynamic Validation 

An unlimited number of Validators can participate in DyPoS consensus. This leader-less consensus is comparable to community consensus versus competitive or leader-based consensus. This maximizes decentralization and rewards all Work proportional to Effort. To improve the throughput even further, especially with large number of Validators, each Validator has an eAgent that functions within a Master Node network and participates in block validations on behalf of the Validator. Dynamic validation uses Tendermint core in the first version of DyPoS.

### Dynamic Block Rewards

Capped, yearly token inflation of 5% is used to pay Block Rewards to four types of blockchain Workers: Validators, eAgents, dGuards, and Govnodes. Their rewards are dynamic and based upon Security Threat Levels and real-time supply and demand. Storecoin takes 1% of inflation to endow blockchain security, engineering, governance, and operations.

### Dynamic dSecurity

The Storecoin Threat Level System determines real-time prices for DyPoS. dGuards study transactions and Validator activity and report potential spam or attacks to Storecoin, Inc. dGuards earn +50% of the stake or the bond of the found bad actors. The rest is burnt.

### Scaling with eAgents

To scale blockchain transactions even further by making validation instant – without expensive and unnecessary network hops – Storecoin is introducing the concept of an encrypted Agent – or an eAgent. It's similar to Power of Attorney where eAgents represent corresponding participating validator nodes inside a single MasterNode, so they finalize the blocks instantly.

### Governance with Checks and Balances

Storecoin Governance is inspired by the United States Constitution where checks and balances ensure that: a) Security Matters Most and b) No Centralization of Power. For the first 4 years, during the Storecoin blockchain rollout and stabilization phases, Storecoin, Inc. will have executive power. All communication and voting data that enables Governance are hosted by decentralized Govnodes.

Visit [Storecoin](https://storeco.in/) for more details on the six engines that make DyPoS, different types of decentralized workers, how free transactions are financed, how block rewards are allocated, and how Storecoin is secured.

The DyPoS changes on top of Tendermint core will be open-sourced in the second half of 2018.

## Early Testing

This repository contains the the first set of tests performed on DyPoS. The tests are run to characterize the behavior of Validator nodes under different load conditions. Since the goal is to learn the behavior of Validator nodes, we focused on the consensus engine and how the blocks are finalized under different loads. We focused on the **consensus efficiency**, which is defined as _the rate at which the incoming transactions are added to the blocks._ It is measured in transactions processed per second. We measured the consensus efficiency under a 4-node (minimum setup required to qualify for BFT), 8-node, and 21-node setups. In each setup, we configured N clients to generate and send bursts of transactions to the Validator nodes, where N is chosen based on the number of Validator nodes. In the 4-node setup, 4 clients are configured. In 8-node setup, 4 and 8 clients are configured in each of the sub-tests. In the 21-node setup, 2, 8, and 12 clients are configured. 

We conducted tests where clients generate both fixed and random sized transactions. Fixed size transactions are used to model the consensus efficiency when transaction sizes can be estimated. Since the consensus efficiency is a function of transaction volume (transaction size x number of transactions) fixed size transactions help us understand how Validators react to varying transaction volumes. We used transactions of size 100, 500, 1,000,  5,000, and 10,000 bytes.

For random transaction sizes, we allocated the sizes as follows.

<table>
  <tr>
    <td>Transaction Group</td>
    <td>Transaction Size</td>
    <td>Percent Allocation</td>
  </tr>
  <tr>
    <td>Group 1</td>
    <td>50-250 bytes</td>
    <td>60%</td>
  </tr>
  <tr>
    <td>Group 2</td>
    <td>251-500 bytes</td>
    <td>20%</td>
  </tr>
  <tr>
    <td>Group 3</td>
    <td>501-1,000 bytes</td>
    <td>15%</td>
  </tr>
  <tr>
    <td>Group 4</td>
    <td>1,001+ bytes</td>
    <td>5%</td>
  </tr>
</table>


Each client is configured to send _bursts_ of transactions to each of the Validator nodes. Storecoin is a payments platform that is expected to be used by merchants and app developers to process transactions generated by the crypto-powered apps (cApps). The transactions are expected to arrive in bursts of different sizes from cApps, so we modeled the tests with these use cases in mind. Generally, most distributed systems can handle continuous load of low volume transactions, but bursts create unpredictable behavior. In our tests the bursts are of size 500, 1,000, 2,000, 5,000, and 10,000 transactions. Each burst lasts for 5 seconds in which the configured number of transactions are sent to each of the Validator nodes. For example, in a 21-node setup with 12 clients and 2,000 transactions per burst, the total transactions sent to one Validator node would be `12 x 2,000 = 24,000` transactions. The total transactions across all the Validator nodes would be `24,000 x 21 = 504,000` transactions and the incoming transaction rate is `504,000 / 5 = 100,800` transactions per second.

The above combination of transaction sizes and rates result in multiple tests in each of the test categories. For each test, consensus efficiency and other statistics are captured as follows.

This observation is recorded for the configuration used in the example above.

<table>
  <tr>
    <td>How To Read Detailed Test Results</td>
    <td></td>
    <td></td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>Avg.</td>
    <td>Std. Dev.</td>
    <td>Max.</td>
  </tr>
  <tr>
    <td>Block Latency</td>
    <td>39.08ms</td>
    <td>25.52ms</td>
    <td>77ms</td>
  </tr>
  <tr>
    <td>Blocks/sec</td>
    <td>215.5</td>
    <td>372.102</td>
    <td>860</td>
  </tr>
  <tr>
    <td>Txs/sec</td>
    <td>62,941</td>
    <td>109,017</td>
    <td>251,763</td>
  </tr>
</table>

**Block Latency:** This is the time taken by the validator to create a new block with transactions included in them. The block latency is determined by the transaction rate and size. The lower the block latency, the higher the consensus efficiency.

**Block/sec:** This is the number of blocks created per second. The higher the number of blocks produced per second, the higher the consensus efficiency.

**Txs/sec:** This is the consensus efficiency measured as the number of transactions processed per second. We use the terms, consensus efficiency, efficiency, throughput, performance, etc. interchanbeably to mean the same thing.

We use Tendermint’s [tm-bench](https://github.com/tendermint/tools/tree/master/tm-bench) to generate transactions. Each test modifies the tm-bench code and configuration to suit the requirements of the tests.

The [tests](https://github.com/StorecoinProject/tm-bench/tree/master/tests) folder contains the list of tests conducted so far. Specifically, the following tests are run.

* [Fixed Size Transactions (4-node)](https://github.com/StorecoinProject/tm-bench/tree/master/tests/test-fixed-tx-size)

* [Random Size (100 - 10K bytes) Transactions (4-node)](https://github.com/StorecoinProject/tm-bench/tree/master/tests/test-random-tx-size)

* [Random Size (Percentage Method) Transactions (4-node)](https://github.com/StorecoinProject/tm-bench/tree/master/tests/test-random-trx-size-percentage)

* [Fixed and Random Size Transactions (8-node)](https://github.com/StorecoinProject/tm-bench/tree/master/tests/test-8-node-fixed-random-tx-size)

We’ll add the test results for the 21-node setup soon.

Each test folder contains the description of the test setup and the results. The logs folder for each of the tests (for [example](https://github.com/StorecoinProject/tm-bench/tree/master/tests/test-8-node-fixed-random-tx-size/random-tx-size/logs)) contains the raw transaction logs for each of tests. The logs can be used to peek into the exact consensus process, the transactions being sent, blocks created, and any failures during transaction processing.

