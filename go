#! /bin/bash

VOW_PATH="node_modules/vows/bin/vows"
PRODUCTION_HOST="ec2-user@54.243.60.165"
STAGING_HOST="ec2-user@54.197.20.199"
UTILITIES_HOST="ec2-user@54.242.37.176"
ASSETS_HOST="ec2-user@54.243.226.216"
PRIVATE_KEY="$HOME/.ssh/stagingventurepactcom.pem"

echo
echo "***Running about_me bash script*****"
echo

case "$1" in
  tests)
    echo "beginning tests"
    echo
    "$VOW_PATH" test --spec
    ;;
  production)
    echo "beginning ssh into production server"
    echo
    ssh -i "$PRIVATE_KEY" "$PRODUCTION_HOST"
    ;;
  staging)
    echo "beginning ssh into staging server"  
    echo
    ssh -i "$PRIVATE_KEY" "$STAGING_HOST"
    ;;
  utilities)
    echo "beginning ssh into utilities server"
    echo
    ssh -i "$PRIVATE_KEY" "$UTILITIES_HOST"
    ;;
  assets)
    echo "beginning ssh into assets server"
    echo
    ssh -i "$PRIVATE_KEY" "$ASSETS_HOST"
    ;;    
  *)
    echo "usage is: firefly {production|staging|tests}"
    echo
    ;;
esac    
    