Simple Translator
===================
Simple Translate Server with Google Translate API

1. Download Directory
<pre><code>go get github.com/Junhyeok99/SimpleTranslate</code></pre>
Make sure you set *$GOPATH*

2. Download Dependancy
<pre><code>cd $GOPATH/src/github.com/Junhyeok99/SimpleTranslate && make</code></pre>
This command will download Related Packages

3. Set your env
Google API needs **json key** path in env variable

Fix ```sample_env.sh``` with your Google API KEY (Go to the link below)
<https://console.developers.google.com/apis>

After change the value run below:

<pre><code>source sample_env.sh</code></pre>

4. Run Program
<pre><code>go run main.go -flags</code></pre>
Description of flags are below

>1. -f
>>Bool type:
>>>true: get text source from file/ false: get text source from argument

>2. -l
>>String type:
>>>Location of file (default "example.txt")

>3. -lang
>>String type:
>>>Destination language to Translate text (default "ko")

>4. -t
>>String type:
>>>Text which you want to Translate

Still in Develop!
