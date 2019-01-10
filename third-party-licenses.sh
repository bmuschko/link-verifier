#!/usr/bin/env bash

MERGE_FILE="THIRD-PARTY-LICENSES"
rm -f ${MERGE_FILE}

echo "Third party libraries" >> ${MERGE_FILE}
echo "=====================" >> ${MERGE_FILE}

licenseFiles=(`find vendor -name 'LICENSE' -o -name 'LICENSE.txt'`)

for i in "${licenseFiles[@]}"
do
   :
   echo "" >> ${MERGE_FILE}
   path="${i#vendor/}"
   len=${#path}
   echo ${path} >> ${MERGE_FILE}
   underscore=$(printf "%0.s=" $(seq 1 $len))
   echo ${underscore} >> ${MERGE_FILE}
   echo "" >> ${MERGE_FILE}
   cat ${i} >> ${MERGE_FILE}
done