#!/bin/bash
count=1
sleep=6
while sleep $sleep
  echo 'Running TM-bench...'


do ( ./tm-bench -T 5 -v -r 10000 -c 4 54.67.122.49:46657,34.227.97.184:46657,18.217.98.51:46657,34.209.158.158:46657 >> ./output/Phase2_T5_C4_N4_R10000.out &);


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

