# Consensus Efficiency of Storecoin DyPoS -Test 4 a (8 nodes  Transaction Size distribution by fixed size) 

**Scope of Test 4**

The purpose of this test is to evaluate the “consensus efficiency” of Storecoin DyPoS (https://storeco.in/) engine with fixed, transaction sizes on 8 nodes,  Consensus efficiency is the rate at which the participating validator nodes agree on the new block. It is measured as transactions per second processed and added to the blocks. This test doesn’t include validating the transactions, executing the transactions after they are added to the block, etc., but attempts to measure raw performance of the consensus engine.
 
Storecoin DyPoS is built on top of Tendermint (https://tendermint.com/). In this test, its benchmarking tool, TM-Bench, is customized to provide the necessary setup described below.


**Test Setup**

 Test has been executed in 2 phases : 

- 4 Connections 

- 8 Connections

Tm-bench is customized where the transaction sizes are fixed, computed subjected to the maximum size. The test setup will be same as Test 1, except that the nodes are increased to 8 and are two set of connections 4 and 8 respectively are used. 
 
TM-Bench is customized to generate transactions of fixed size (in bytes). The number of clients (“c” = 4/”c”=8 ) is set to be same as the number of validator nodes (“N” = 8). Each client generates transactions of specified size and sends them to the configured validator node. All the 8 validator nodes receive transactions concurrently and the elected proposer node proposes the new blocks with transactions received by all the validators. The tests are run with fixed transaction sizes of 100 bytes, 500 bytes, 1 KB, 5 KB and 10 KB  for a duration (“T” = 5) of 5 seconds. The transaction rate ("r") will be 500, 1000, 2000, 5000, and 10000.


**Environment and Tools**

A cluster containing 8 validator nodes is set up on an Amazon Web Services (AWS) Elastic Compute Cloud (EC2). Each validator node runs on r3.xlarge instance with 4 CPUs, 30.5GB memory, and 80GB SSD drive. The nodes are located in the following regions.
 
Nodes :  r3.xlarge instances

- California - ( 2 nodes in the region) 

- Virginia  - ( 2 nodes in the region) 

- Ohio   - ( 2 nodes in the region) 

- Oregon  - ( 2 nodes in the region) 

The 8 validator nodes are geographically spread across the United States and the clients are configured in a 5th region in order to simulate real world setup.

**Running the Test**

4 Connections: 

- C1 sends r transactions of fixed size s to N1 for the duration T. 

- C2 sends r transactions of fixed size s to N2 for the duration T. 

- C3 sends r transactions of fixed size s to N3 for the duration T. 

- C4 sends r transactions of fixed size s to N4 for the duration T.

- C1 sends r transactions of fixed size s to N5 for the duration T. 

- C2 sends r transactions of fixed size s to N6 for the duration T. 

- C3 sends r transactions of fixed size s to N7 for the duration T. 

- C4 sends r transactions of fixed size s to N8 for the duration T.

8 Connections: 

- C1 sends r transactions of fixed size s to N1 for the duration T. 

- C2 sends r transactions of fixed size s to N2 for the duration T. 

- C3 sends r transactions of fixed size s to N3 for the duration T. 

- C4 sends r transactions of fixed size s to N4 for the duration T.

- C5 sends r transactions of fixed size s to N5 for the duration T. 

- C6 sends r transactions of fixed size s to N6 for the duration T. 

- C7 sends r transactions of fixed size s to N7 for the duration T. 

- C8 sends r transactions of fixed size s to N8 for the duration T.


**Test Summary**

Test was performed for fixed transaction  size and 5 transaction rates were used resulting in a total of 25 tests for 4 connections and 25 test for 8 connections  Some tests, at higher rate, failed due to timeouts resulting from excessive load on the system.

4 Connections : 

See here : https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/4a-4-node.png

8 Connections: 

See here : https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/4a-8-node.png

**Test Results**


In the following test results, “Blocks/sec” indicates the number of blocks added to the blockchain per second and “Txs/sec” indicates the consensus efficiency. “Txs/sec” is the rate at which the transactions are processed and included in the blocks.

See the transaction logs for failures.

4 Connections : 

See here :

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C4_N8_R500.png

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C4_N8_R1000.png

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C4_N8_R2000.png

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C4_N8_R5000.png

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C4_N8_R10000.png

8 Connections : 

See here :

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C8_N8_R500.png

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C8_N8_R1000.png

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C8_N8_R2000.png

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C8_N8_R5000.png

- https://github.com/StorecoinProject/tm-bench/blob/test-8-node-fixed-random-tx-size/tests/test-8-node-fixed-random-tx-size/fixed-tx-size/images/Phase4_P1_T5_C8_N8_R1000.png



**Results Summary**

Consensus efficiency is directly proportional to the transaction volume. The transaction volume is (transaction size * number of transactions). At smaller transaction sizes (see the results for 500 byte transactions) tests passed for higher input transaction rates, but as the transaction sizes increases, the test failed for higher rates. The failures here result in rejected transactions. The validator nodes remained stable even at failure levels above. 

See the transaction logs for details.

**What’s Next?**


- Test 5 (test-8-node-burst-mode-random-distributed-tx-size) uses the same setup as Test 4, but the clients send transactions in bursts to the validator nodes. Each test is run for 10 minutes with clients sending predetermined number of transactions in short bursts. The purpose of this test is to observe consensus efficiency over a longer duration and to evaluate if burst mode results in better throughput.

- Test 6 (test-21-node-burst-mode-random-distributed-tx-size) extends test 5 to 21 validator nodes.

- Test 7 (test-21-node-burst-mode-real-tx) uses real transactions which need to be validated before they are included in the block. The purpose of this test is to measure the throughput with transaction overheads included in the throughput numbers.

The dates for tests 5 to 7 will be published as we make progress.
