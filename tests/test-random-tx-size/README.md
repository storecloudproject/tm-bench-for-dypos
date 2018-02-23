# Consensus Efficiency of Storecoin DyPoS - Test 1 (test-random-tx-size)

**Scope of Test 1**

The purpose of this test is to evaluate the “consensus efficiency” of Storecoin DyPoS (https://storeco.in/) engine with random,transaction sizes. Consensus efficiency is the rate at which the participating validator nodes agree on the new block. It is measured as transactions per second processed and added to the blocks. This test doesn’t include validating the transactions, executing the transactions after they are added to the block, etc., but attempts to measure raw performance of the consensus engine.
 
Storecoin DyPoS is built on top of Tendermint (https://tendermint.com/). In this test, its benchmarking tool, TM-Bench, is customized to provide the necessary setup described below.

**Test Setup**
 
 
Tm-bench is customized where the transaction sizes are randomly computed subjected to the maximum size. The test setup will be same as Phase 1, except that the transaction sizes are randomized.  
Transaction Size (100 byte to 10k byte)
 
 The number of clients (“c” = 4) is set to be same as the number of validator nodes (“N” = 4). Each client generates transactions of random size and sends them to the configured validator node. All the 4 validator nodes receive transactions concurrently and the elected proposer node proposes the new blocks with transactions received by all the validators.  The tests are run with random transaction sizes of 100, 500, 1K, 5K and 10K bytes for a duration (“T” = 5) of 5 seconds. The transaction rate ("r") will be 500, 1000, 2000, 5000, and 10000.

**Environment and Tools**
 
A cluster containing 4 validator nodes is set up on an Amazon Web Services (AWS) Elastic Compute Cloud (EC2). Each validator node runs on r3.xlarge instance with 4 CPUs, 30.5GB memory, and 80GB SSD drive. The nodes are located in the following regions.
 
Nodes :  r3.xlarge instances
California 
Virginia
Ohio  
Oregon 
The transactions are generated from clients running on an m4.2xlarge instance located in Canada.
 
The 4 validator nodes are geographically spread across the United States and the clients are configured in a 5th region in order to simulate real world setup.

**Running the Test**

C1 sends r transactions of random size s to N1 for the duration T. 
C2 sends r transactions of random size s to N2 for the duration T. 
C3 sends r transactions of random size s to N3 for the duration T. 
C4 sends r transactions of random size s to N4 for the duration T.

**Test Summary**

5 different transaction sizes are used and for each transaction size 5 tests are performed, resulting in a total of 25 tests. Some tests, at higher rate, failed due to timeouts resulting from excessive load on the system.

Time (T)
Transactions (r)
Connections (c)

**Test Results**

In the following test results, “Blocks/sec” indicates the number of blocks added to the blockchain per second and “Txs/sec” indicates the consensus efficiency. “Txs/sec” is the rate at which the transactions are processed and included in the blocks. 

* See the transaction logs for failures.

Test Tx-Random

Test for Transaction Size : Random
Connections : 4
Rate : 500. 1000,2000,5000,10000
Time : 5 Sec



Transaction Size 5000 Byte
Transactions
Stats
Avg
Stdev
Max

**Results Summary**

Consensus efficiency is directly proportional to the transaction volume. The transaction volume is (transaction size * number of transactions). 
At smaller transaction sizes (see the results for 100 and 500 byte transactions) tests passed for higher input transaction rates, but as the transaction sizes increases, the test failed for higher rates. The failures here result in rejected transactions.
The validator nodes remained stable even at failure levels above. See the transaction logs for details.

**What’s Next?**

Test 3 (test-random-distributed-tx-size) is a special case of Test 2 where the transaction sizes are distributed based on a % distribution. For example, 60% of the transactions are 500 bytes or less, 20% of the transactions are 500 to 2,000 bytes in length, 10%, 2,000 to 5,000 bytes in length, and another 10%, 5,000 to 10,000 bytes in length. Within each threshold, the transaction sizes are randomized. This test serves as a basis for rest of the tests discussed below.
Test 4 (test-8-node-random-distributed-tx-size) uses 8 validator nodes with 4 and 8 clients sending transactions as described in Test 3. The purpose of this test is to measure the consensus efficiency with a larger set of validator nodes.
Test 5 (test-8-node-burst-mode-random-distributed-tx-size) uses the same setup as Test 4, but the clients send transactions in bursts to the validator nodes. Each test is run for 10 minutes with clients sending predetermined number of transactions in short bursts. The purpose of this test is to observe consensus efficiency over a longer duration and to evaluate if burst mode results in better throughput.
Test 6 (test-21-node-burst-mode-random-distributed-tx-size) extends test 5 to 21 validator nodes.
Test 7 (test-21-node-burst-mode-real-tx) uses real transactions which need to be validated before they are included in the block. The purpose of this test is to measure the throughput with transaction overheads included in the throughput numbers.

The dates for tests 3 to 7 will be published as we make progress.

