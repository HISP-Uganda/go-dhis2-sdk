#!/bin/bash

set -e

# CONFIGURATION
GENERATOR_VERSION="7.6.0"
GENERATOR_JAR="$HOME/openapi-generator-cli.jar"
SCHEMA_DIR="dhis2/schema"
OPENAPI_FILE="dhis2.openapi.json"
OPENAPI_URL="https://play.im.dhis2.org/stable-2-42-0/api/openapi.json"
#MODEL_LIST_FILE="$SCHEMA_DIR/models.txt"
CONFIG_JSON="generator-config.json"

# Ensure Java is installed
if ! command -v java >/dev/null; then
  echo "❌ Java not found. Please install Java."
  exit 1
fi

# Check for openapi-generator-cli or fallback JAR
if command -v openapi-generator-cli >/dev/null 2>&1; then
  GENERATOR_CMD="openapi-generator-cli"
  echo "✅ Using openapi-generator-cli binary."
elif [ -f "$GENERATOR_JAR" ]; then
  GENERATOR_CMD="java -jar $GENERATOR_JAR"
  echo "✅ Using $GENERATOR_JAR"
else
  echo "❌ Neither 'openapi-generator-cli' nor '$GENERATOR_JAR' found."
  echo "   Please install OpenAPI Generator CLI or provide the JAR."
  exit 1
fi

# Download JAR if needed (only if not using binary and file is missing)
if [ "$GENERATOR_CMD" = "java -jar $GENERATOR_JAR" ] && [ ! -f "$GENERATOR_JAR" ]; then
  echo "⬇️  Downloading OpenAPI Generator CLI v$GENERATOR_VERSION..."
  curl -L "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${GENERATOR_VERSION}/openapi-generator-cli-${GENERATOR_VERSION}.jar" -o "$GENERATOR_JAR"
fi

# Fetch DHIS2 OpenAPI spec
echo "🌐 Fetching DHIS2 OpenAPI spec..."
curl -s -u admin:district "$OPENAPI_URL" -o "$OPENAPI_FILE"

## Clean previous schema
echo "🧹 Cleaning old schema directory..."
rm -rf "$SCHEMA_DIR"

##Extract model names
echo "🔍 Extracting referenced models..."
mkdir -p "$SCHEMA_DIR"
##grep '"$ref":"#/components/schemas/' "$OPENAPI_FILE" \
##  | cut -d'/' -f4 \
##  | sed 's/[",]//g' \
##  | sort -u \
##  > "$MODEL_LIST_FILE"
##
#
# Generate selected models
echo "⚙️  Generating Go models..."
java -jar "$GENERATOR_JAR" generate \
  -i "$OPENAPI_FILE" \
  -g go \
  -o "$SCHEMA_DIR" \
  --package-name=schema \
  --additional-properties=enumClassPrefix=true,modelPropertyNaming=original

## Clean non-Go files
echo "🚮 Cleaning non-Go files..."
find "$SCHEMA_DIR" -type f ! -name '*.go' ! -name 'models.txt' -delete
find "$SCHEMA_DIR" -type d -not -path "$SCHEMA_DIR" -exec rm -rf {} +

#
for file in "$SCHEMA_DIR"/model_*.go; do
  [ -f "$file" ] || continue
  base=$(basename "$file")
  name="${base#model_}"         # remove model_ prefix
  name="${name%.go}"            # strip .go
  #pascal=$(to_pascal_case "$name")
  mv "$file" "$SCHEMA_DIR/$name.go"
  echo "✅ $base → $name.go"
done
#
echo "🎉 Done: Generated models into directory $SCHEMA_DIR"

echo "🛠 Patching schema.configuration with Debug, HTTPClient, UserAgent, and auth types..."

CONFIG_FILE=$(find "$SCHEMA_DIR" -name 'configuration.go')
if [ -f "$CONFIG_FILE" ]; then
  # Ensure net/http and context are imported
  perl -i -pe '
    BEGIN { $added = 0 }
    if (/^import \($/ && !$added) {
      $added = 1;
      $_ .= qq{    "net/http"\n    "context"\n};
    }
  ' "$CONFIG_FILE"

  # Patch fields into Configuration struct
  perl -0777 -i -pe '
    s/(type Configuration struct \{.*?)(\n\})/$1
    Host string `json:"-"`
    Scheme string `json:"-"`
    DefaultHeader map[string]string `json:"-"`
    Debug bool `json:"-"`
    UserAgent string `json:"-"`
    HTTPClient *http.Client `json:"-"`$2/s
  ' "$CONFIG_FILE"

  # Append auth-related types and method if not already present
  if ! grep -q 'type BasicAuth struct' "$CONFIG_FILE"; then
    cat <<EOF >> "$CONFIG_FILE"

type BasicAuth struct {
    UserName string
    Password string
}

type ContextKey string

const (
    ContextBasicAuth ContextKey = "basic"
)

func (c *Configuration) ServerURLWithContext(ctx context.Context, operation string) (string, error) {
    return "", nil // Customize if needed
}
EOF
  fi

  echo "✅ Configuration struct patched."
else
  echo "❌ Configuration.go not found in $SCHEMA_DIR"
fi

echo "🧽 Formatting generated Go files..."
perl -pi -e 's/\bAny\b/any/g' "$SCHEMA_DIR"/*.go

if command -v goimports >/dev/null 2>&1; then
  goimports -w "$SCHEMA_DIR"/*.go
else
  echo "⚠️  goimports not found. Falling back to gofmt."
  gofmt -s -w "$SCHEMA_DIR"/*.go
fi

echo "✅ Go files formatted."
#

