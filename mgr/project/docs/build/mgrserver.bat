
echo "1. " 

cd ..\..\mgrserver\web


npm run build

echo "2. " 

cd ..

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build --tags="prod"


echo "3. "

scp ./mgrserver root@121.196.168.242:/tmp


ssh   root@121.196.168.242 "cd /root/work/mgrserver/bin;./mgrserver stop;mv ./mgrserver ./mgrserver_${dt} ;cp /tmp/mgrserver ./;sleep 3;./mgrserver start;rm -rf /tmp/mgrserver"


rm ./mgrserver

echo "3. "
