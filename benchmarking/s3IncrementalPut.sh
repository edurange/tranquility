#!/usr/bin/env bash

for i in {1..1000}
do
	echo "foo 1566429105 echo bar" > testOutput${i}.log

	date=`date +%Y%m%d`
	dateFormatted=`date -R`
	s3Bucket="storagespeedbench"
	fileName="testOutput${i}.log"
	relativePath="/${s3Bucket}/${fileName}"
	contentType="text/plain"
	stringToSign="PUT\n\n${contentType}\n${dateFormatted}\n${relativePath}"
	s3AccessKey="AKIAVODDJPWATJPXUHVP"
	s3SecretKey="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"
	signature=`echo -en ${stringToSign} | openssl sha1 -hmac ${s3SecretKey} -binary | base64`
	curl -X PUT -T "${fileName}" \
	-H "Host: ${s3Bucket}.s3.amazonaws.com" \
	-H "Date: ${dateFormatted}" \
	-H "Content-Type: ${contentType}" \
	-H "Authorization: AWS ${s3AccessKey}:${signature}" \
	http://${s3Bucket}.s3-us-west-2.amazonaws.com/${fileName}

	rm testOutput${i}.log
done

#curl here


