echo "------stopping container -----"
containerid = `docker ps|awk -F“ “‘{Print $1}’`
echo $containerid
docker -rm -f $containerid
