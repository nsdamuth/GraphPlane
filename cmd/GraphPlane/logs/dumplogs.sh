#/bin/bash

echo "Dumping GraphPlane Logs"

for entry in *
do
    if [ "${entry: -4}" == ".log" ]; then
        > $entry
    fi
done