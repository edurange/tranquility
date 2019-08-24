# Tranquility

Tranquility is a small golang microservice designed to capture log files in real time, remotly, through an API endpoint. I designed it to make real-time collection of bash histories easy and secure. It uses Redis for storage.

## Development

The .main binary is prebuilt for MacOS, and the .tranquility binary is prebuilt for 64-bit Linux. To develop, you'll need the google/uuid, go-redis/redis, and gin-gonic/gin libraries installed on your computer as well. 

To build for 64-bit Linux, just run:

```shell
GOOS=linux GOARCH=amd64 go build -o tranquility .
```

## Usage

To begin serving, make sure Redis is running on the host. Right now, Redis credentials are hardcoded, although I plan to change that in the future. 

```shell
./tranquility [uuid]
```

Where uuid is a security phrase that must be included in the headers of GET or PUT commands.

Payloads are structured like:

```json
{
	user: [REDIS LOG KEY, STRING],
	time: [UNIX TIMESTAMP, INTEGER],
	command: [TEXT TO LOG, STRING]
}
```

An example command to log a target machine's bash history could look like this: (place the following line inside /etc/bashrc)

```shell
PROMPT_COMMAND='history -a; histJSON=$(echo {\"user\": \"$(whoami)\", \"time\": $(date +%s), \"command\": \"$(history 1 | cut -c 8-)\"}); curl -s -d "$histJSON" -H "Content-Type: application/json" -H "uuid: foo" -X POST <%= hostAddress:8080/logger >> /dev/null'
```

To fetch, then, just run a GET on any machine that can access the host:

```shell
curl -s -H "Content-Type: application/json" -H "uuid: foo" -X GET localhost:8080/results/user > testy.log
```

## Testing

There's no tests right now, but the /benchmarking directory has some benchmarking tests I wrote to compare log read/write speed between this API, pure redis-cli commands, and logging over s3. Use at your own risk. 

## etc

./client $(history 1 | cut -c 8-)