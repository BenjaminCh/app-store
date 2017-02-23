#!/bin/bash

if [ -z "$1" ] ; then
    echo 'Env param needed (dev, prod or all)'
    exit 1
fi

echo "YARN TASK"
sudo yarn run build

if [ $1 = "prod" -o $1 = "all" ]; then

  echo "DIST TO PROD S3 MASTER"
  aws s3 sync dist/ s3://search.tests.sh --acl public-read

  echo "Change default index & error documents"
  aws s3 website s3://search.tests.sh --index-document `find . -maxdepth 1 -name "*index.html" | sed -e "s/.\///g"` --error-document `find . -maxdepth 1 -name "*index.html" | sed -e "s/.\///g"`

fi