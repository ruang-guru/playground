#!/bin/bash
target=$1 # target yang akan ditest
maxRate=$2 # rate maksimal
rateInc=$3 # kenaikan rate
incDuration=$4 # durasi tiap kenaikan
startAt=$5 # rate awal
currentRate=$startAt

while [ $currentRate -le $maxRate ] # Selama request rate saat ini kurang dari sama dengan maxrate
do
  echo $target | vegeta attack -rate=$currentRate -duration=$incDuration > reel-$maxRate-$currentRate-test.bin
  # ^ melakukan vegeta attack dan menyimpan hasil dalam reel-xx-xx-test.bin
  currentRate=$((currentRate+rateInc)) # Meningkatkan request
done
