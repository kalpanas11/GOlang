echo "------stopping container -----"
containerid = `docker ps|awk -F“ “‘{Print $1}’`
docker -rm -f $containerid
