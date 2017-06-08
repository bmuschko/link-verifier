#!/usr/bin/env bash

BUILD_TMP_DIR="build/tmp"
MERGE_FILE="${BUILD_TMP_DIR}/THIRD-PARTY-LICENSES"

mkdir -p ${BUILD_TMP_DIR}
rm -f ${MERGE_FILE}

echo "Third party libraries" >> ${MERGE_FILE}
echo "=====================" >> ${MERGE_FILE}

licenseFiles=(`find vendor -name 'LICENSE'`)

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