
# DHIS2 Go SDK

This is a Go-based SDK for interacting with the [DHIS2 Web API](https://docs.dhis2.org/), supporting:

- Tracker and Aggregate payloads
- Query parameter construction
- Typed request and response models
- Structured error handling
- Auto-generated schema models from DHIS2's OpenAPI spec

---

## 📦 Package Structure

```text
dhis2/
├── client.go             # HTTP client using Resty
├── query/                # Chainable query param builder
├── response/             # Response summaries (e.g. import status)
├── error/                # Custom DHIS2-aware error types
├── schema/               # Auto-generated models (Event, TEI, etc.)
├── generate-schema.sh    # Rebuild schema/ using OpenAPI spec
```

---

## ✅ Getting Started

### Installation

```bash
go get github.com/HISP-Uganda/dhis2-sdk/dhis2
```

### Example Usage

```go
package main

import (
    "log"
    "github.com/HISP-Uganda/dhis2-sdk/dhis2"
    "github.com/HISP-Uganda/dhis2-sdk/dhis2/schema"
)

func main() {
    client := dhis2.NewClient("https://your-dhis2.org", "admin", "district")

    event := &schema.Event{
        Program: "abc123",
        OrgUnit: "xyz789",
        Status:  "ACTIVE",
    }

    resp, err := client.Post("/api/events", event)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    log.Println("✅ Submitted event successfully:", resp.Status())
}
```

---

## 🧱 Schema Models (`schema/`)

The `schema/` package contains Go types auto-generated from DHIS2's OpenAPI spec (v40+). Only the models listed in [`models.txt`](./dhis2/schema/models.txt) are retained and cleaned.

Each model file is:
- Formatted with `goimports` or `gofmt`
- Version-controlled and reproducible

---

## ⚙️ Regenerating Models

To regenerate the schema types after a DHIS2 upgrade:

```bash
./generate-schema.sh
```

This will:
1. Download the latest OpenAPI spec
2. Generate all models
3. Retain only models listed in `models.txt`
4. Rename and format `.go` files

---

## 🧰 Dependencies

- Go 1.19+
- [Resty](https://github.com/go-resty/resty)
- [`goimports`](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) (optional but recommended)

Install dependencies:

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

---

## 📚 Resources

- [DHIS2 Web API Docs](https://docs.dhis2.org/master/en/developer/html/webapi.html)
- [OpenAPI Generator](https://openapi-generator.tech/)

---

## 📬 Contributions

Pull requests, issues, and feature suggestions are welcome. If you’d like to support more DHIS2 endpoints, helper methods, or automatic retries — open an issue!

---

## 📄 License
Copyright (c) HISP-Uganda 2025,

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
