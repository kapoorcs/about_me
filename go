#! /bin/bash

VOW_PATH="node_modules/vows/bin/vows"
STAGING_HOST="ec2-user@54.197.20.199"
PRIVATE_KEY="$HOME/.ssh/stagingventurepactcom.pem"

echo
echo "***Running about_me bash script*****"
echo

case "$1" in
  tests)
    echo "beginning tests"
    echo
    #"$VOW_PATH" test --spec
    ;;
 
  staging)
    echo "beginning ssh into staging server"  
    echo
    ssh -i "$PRIVATE_KEY" "$STAGING_HOST"
    ;;
   
  *)
    echo "usage is: firefly {production|staging|tests}"
    echo
    ;;
esac    
    