#!/bin/bash
target=$1 # target yang akan ditest
maxRate=$2 # rate maksimal
peakDuration=$3 # durasi akses peak
normalRate=$4 # rate awal
normalDuration=$5 # durasi akses normal
currentRate=$normalRate
count=0 # counter awal
repeat=$6 # perulangan spike

while [ $count -lt $repeat ] # Selama nilai counter saat ini kurang dari repeat
do
    echo $target | vegeta attack -rate=$currentRate -duration=$normalDuration > down-$currentRate-test-$count.bin
    # ^ melakukan vegeta attack dan menyimpan hasil dalam file bin
    currentRate=$maxRate # Menaikkan rate secara melonjak

    echo $target | vegeta attack -rate=$currentRate -duration=$peakDuration > up-$currentRate-test-$count.bin
    # ^ melakukan vegeta attack dan menyimpan hasil dalam file bin
    currentRate=$normalRate # Menurunkan rate akses

count=$(( $count + 1 ))
done
