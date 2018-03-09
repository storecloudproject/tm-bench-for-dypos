#!/bin/bash
count=1
sleep=6
while sleep $sleep
  echo 'Running TM-bench...'


do ( ./tm-bench -T 5 -v -r10000 -c 8 54.67.122.49:46657,52.53.67.196:46657,34.227.97.184:46657,34.225.8.67:46657,18.217.98.51:46657,18.218.220.224:46657,34.209.158.158:46657,34.215.61.224:46657 >> ./../logs/C8N8_fixed-tx-size/Phase4_P1_T5_C8_N8_R10000_S10000.out &);


  count=`expr $count  + 1`
  if [ $count -gt 1 ]
  then
      break
  fi

  round=`expr $count % 10`

  if [ $round -eq 0 ]
  then
      echo 'Sleeping for catch up...'
      sleep 30
  fi

done

