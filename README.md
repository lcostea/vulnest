## Build

go build -o vne .

## Run

### AWS S3

./vne s3 scan --bucket {BUCKET_NAME}

If you found any issues you can run with extended output to get more details on those misconfigurations

./vne s3 scan --bucket {BUCKET_NAME} --extended-output

### JWT 

Elevate to Admin: user=admin, role=admin

./vne jwt elevate-to-admin --jwt <<jwt-goes-here>>
