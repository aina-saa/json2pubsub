# json2pubsub

Publish JSON object (stream) into GCP Pub/Sub topic based on a field value.

```
Usage: json2pubsub --project=STRING <mapping> ...

Reads JSON object (stream) from file and routes it/them to GCP Pub/Sub topics.

Arguments:
  <mapping> ...    Format: VALUE=json.field:my-topic-for-value VALUE2=json.field:my-topic-for-value2

Flags:
  -h, --help              Show context-sensitive help.
      --project=STRING    Google Cloud Platform project id where the Pub/Sub topics in mappings are located in.
  -f, --file="-"          Input file or '-' for stdin
      --version           Print version information and quit
      --[no-]quiet        Be quiet.
```

* project is a project-id for your google cloud project that hosts the Pub/Sub topics provided in mapping
* mapping maps JSON field and its value to certain pub-sub topic (see example below)
  * Message can also be a "*" which will map all if the field is present in the JSON object.

Known limitations:
* This utility expects JSON objects to be presented in input stream (file or stdin) as one object per line. No multiline objects are supported (so disable any pretty printing in your output for example by using ```jq -c```)

## Usage example (from file):

Lets assume that we have a following JSON objects in a file:

```json
{"Message": "Hello World!", "Kind": "EHLO"}
{"Message": "Hello cruel world!", "Kind": "HELLO"}
```

We wish to send these two messages to seperate Pub/Sub topics. We would achieve this by doing:

```sh
json2pubsub --project=my-project-id --file=my-file.json EHLO=Kind:ehlo-topic HELLO=Kind:hello-topic
```

Login credentials for GCP are pulled from GCE metadata server or you can provide Service Account key file in JSON form by declaring it in environment. Methods are described in <https://cloud.google.com/docs/authentication/production>

## Usage example (from stdin)

```sh
cat my-file.json | json2pubsub --project=my-project-id EHLO=Kind:ehlo-topic HELLO=Kind:hello-topic
```

Again, credentials are detected based on service account login flow.
