echo "Generating GoLang sdk for given spec: $1"
if  [ ! -f openapi-generator-cli.jar ]; then
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.4.0/openapi-generator-cli-5.4.0.jar -O openapi-generator-cli.jar
fi
mkdir codegen
java -jar openapi-generator-cli.jar generate -i $1 -g go -o codegen/
cd codegen