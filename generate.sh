#!/bin/bash

csv_list=`find csv -type f`
format_list="skk contacts"
for format in ${format_list} ; do for csv in ${csv_list} ; do ./syosyo --format ${format} `basename ${csv} .csv` ; done ; done