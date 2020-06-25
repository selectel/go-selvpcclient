#!/usr/bin/env bash

echo "==> Running go test and creating a coverage profile..."
i=0
failed=0
for testingpkg in $(go list ./selvpcclient/.../testing); do
  coverpkg=${testingpkg::-8}
  go test -v -covermode count -coverprofile "./${i}.coverprofile" -coverpkg $coverpkg $testingpkg
  if [ $? -eq 1 ]; then
     failed+=1
  fi
  ((i++))
done
gocovmerge $(ls ./*.coverprofile) > coverage.out
rm *.coverprofile

if ((failed>0)); then
  exit 1
fi
exit 0