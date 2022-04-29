#!/bin/bash
# $1 $2 dst. merupakan posisi argumen yang digunakan untuk menjalankan script ini
maxRate=$1 # rate maksimal
rateInc=$2 # kenaikan rate
incDuration=$3 # durasi tiap kenaikan
startAt=$4 # rate awal
currentRate=$startAt

while [ $currentRate -le $maxRate ] # selama request rate saat ini kurang dari sama dengan maxrate
do
  jq -ncM 'while(true; .+1) | {method: "POST", url: "http://localhost:8090/movie", body: {episode: ., name: "bocek"} | @base64}' \
  | vegeta attack -lazy --format=json -rate=$currentRate -duration=$incDuration \
  > reel-$maxRate-$currentRate-test.bin 
  # ^ melakukan vegeta attack dan menyimpan hasil dalam reel-xx-xx-test.bin
  currentRate=$((currentRate+rateInc)) # Meningkatkan request
done
