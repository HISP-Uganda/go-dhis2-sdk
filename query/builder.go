package query

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Params struct {
	values url.Values
}

func New() *Params {
	return &Params{values: url.Values{}}
}

func (p *Params) Set(key, value string) *Params {
	p.values.Set(key, value)
	return p
}

func (p *Params) Add(key, value string) *Params {
	p.values.Add(key, value)
	return p
}

func (p *Params) AddList(key string, values []string) *Params {
	for _, v := range values {
		p.values.Add(key, v)
	}
	return p
}

func (p *Params) ToMap() map[string]string {
	result := make(map[string]string)
	for k, v := range p.values {
		if len(v) > 0 {
			result[k] = v[0]
		}
	}
	return result
}

// Paging helpers
func (p *Params) Paging(enabled bool) *Params {
	p.values.Set("paging", strconv.FormatBool(enabled))
	return p
}

func (p *Params) Page(page int) *Params {
	p.values.Set("page", strconv.Itoa(page))
	return p
}

func (p *Params) PageSize(size int) *Params {
	p.values.Set("pageSize", strconv.Itoa(size))
	return p
}

// Filter helpers
func (p *Params) StartDate(date string) *Params {
	p.values.Set("startDate", date)
	return p
}

func (p *Params) EndDate(date string) *Params {
	p.values.Set("endDate", date)
	return p
}

func (p *Params) LastUpdated(date string) *Params {
	p.values.Set("lastUpdated", date)
	return p
}

func (p *Params) OrgUnitMode(mode string) *Params {
	p.values.Set("ouMode", mode)
	return p
}

// Logical filter helpers
func (p *Params) Filter(field, operator, value string) *Params {
	filterStr := fmt.Sprintf("%s:%s:%s", field, operator, value)
	p.values.Add("filter", filterStr)
	return p
}

// Common logical filter wrappers
func (p *Params) Eq(field, value string) *Params {
	return p.Filter(field, "eq", value)
}

func (p *Params) Ne(field, value string) *Params {
	return p.Filter(field, "ne", value)
}

func (p *Params) Gt(field, value string) *Params {
	return p.Filter(field, "gt", value)
}

func (p *Params) Ge(field, value string) *Params {
	return p.Filter(field, "ge", value)
}

func (p *Params) Lt(field, value string) *Params {
	return p.Filter(field, "lt", value)
}

func (p *Params) Le(field, value string) *Params {
	return p.Filter(field, "le", value)
}

func (p *Params) Like(field, value string) *Params {
	return p.Filter(field, "like", value)
}

func (p *Params) Ilike(field, value string) *Params {
	return p.Filter(field, "ilike", value)
}

func (p *Params) In(field string, values []string) *Params {
	inValue := fmt.Sprintf("%s:in:[%s]", field, join(values, ","))
	p.values.Add("filter", inValue)
	return p
}

func (p *Params) Fields(fields ...string) *Params {
	p.values.Set("fields", strings.Join(fields, ","))
	return p
}

// NestedFields
// params := query.New().NestedFields(map[string][]string{
//     "event":       {},
//     "dataValues":  {"dataElement", "value"},
//     "orgUnit":     {"id", "name"},
// })

func (p *Params) NestedFields(projection map[string][]string) *Params {
	var parts []string
	for parent, children := range projection {
		if len(children) == 0 {
			parts = append(parts, parent)
		} else {
			nested := fmt.Sprintf("%s[%s]", parent, strings.Join(children, ","))
			parts = append(parts, nested)
		}
	}
	return p.Fields(parts...)
}

func (p *Params) ExcludeFields(fields ...string) *Params {
	excludeList := make([]string, len(fields))
	for i, f := range fields {
		excludeList[i] = fmt.Sprintf("!%s", f)
	}
	return p.Fields(excludeList...)
}

func (p *Params) FieldSet(preset string) *Params {
	switch preset {
	case "minimal":
		return p.Fields("id", "name")
	case "detailed":
		return p.Fields("id", "name", "code", "shortName", "description", "lastUpdated")
	case "full":
		return p.Fields("*")
	default:
		return p
	}
}

func join(values []string, sep string) string {
	result := ""
	for i, v := range values {
		if i > 0 {
			result += sep
		}
		result += v
	}
	return result
}

// params := query.New().
//     Set("program", "abc123").
//     Add("paging", "false").
//     AddList("fields", []string{"event", "status", "eventDate"})
//
// resp, err := client.Get("/api/events", &result, params.ToMap())
