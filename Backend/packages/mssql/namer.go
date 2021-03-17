package mssql

import (
	"fmt"

	"gorm.io/gorm/schema"
)

type MniNamer struct {
	TablePrefix   string
	SingularTable bool
}

func (MniNamer) TableName(table string) string {
	return table
}

func (MniNamer) ColumnName(table, column string) string {
	return fmt.Sprintf("%v", column)
}

func (MniNamer) JoinTableName(joinTable string) string {
	return joinTable
}

func (MniNamer) RelationshipFKName(relationship schema.Relationship) string {
	return relationship.Name
}

func (MniNamer) CheckerName(table, column string) string {
	return fmt.Sprintf("%v.%v", table, column)
}

func (MniNamer) IndexName(table, column string) string {
	return fmt.Sprintf("%v.%v", table, column)
}

// // TableName convert string to table name
// func (ns MniNamer) TableName(str string) string {
// 	if ns.SingularTable {
// 		return ns.TablePrefix + ns.toDBName(str)
// 	}
// 	return ns.TablePrefix + inflection.Plural(ns.toDBName(str))
// }

// // ColumnName convert string to column name
// func (ns MniNamer) ColumnName(table, column string) string {
// 	return ns.toDBName(column)
// }

// // JoinTableName convert string to join table name
// func (ns MniNamer) JoinTableName(str string) string {
// 	if !ns.NoLowerCase && strings.ToLower(str) == str {
// 		return ns.TablePrefix + str
// 	}

// 	if ns.SingularTable {
// 		return ns.TablePrefix + ns.toDBName(str)
// 	}
// 	return ns.TablePrefix + inflection.Plural(ns.toDBName(str))
// }

// // RelationshipFKName generate fk name for relation
// func (ns MniNamer) RelationshipFKName(rel schema.Relationship) string {
// 	return ns.formatName("fk", rel.Schema.Table, ns.toDBName(rel.Name))
// }

// // CheckerName generate checker name
// func (ns MniNamer) CheckerName(table, column string) string {
// 	return ns.formatName("chk", table, column)
// }

// // IndexName generate index name
// func (ns MniNamer) IndexName(table, column string) string {
// 	return ns.formatName("idx", table, ns.toDBName(column))
// }

// func (ns MniNamer) formatName(prefix, table, name string) string {
// 	formatedName := strings.Replace(fmt.Sprintf("%v_%v_%v", prefix, table, name), ".", "_", -1)

// 	if utf8.RuneCountInString(formatedName) > 64 {
// 		h := sha1.New()
// 		h.Write([]byte(formatedName))
// 		bs := h.Sum(nil)

// 		formatedName = fmt.Sprintf("%v%v%v", prefix, table, name)[0:56] + string(bs)[:8]
// 	}
// 	return formatedName
// }

// var (
// 	// https://github.com/golang/lint/blob/master/lint.go#L770
// 	commonInitialisms         = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
// 	commonInitialismsReplacer *strings.Replacer
// )

// func init() {
// 	commonInitialismsForReplacer := make([]string, 0, len(commonInitialisms))
// 	for _, initialism := range commonInitialisms {
// 		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
// 	}
// 	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
// }

// func (ns MniNamer) toDBName(name string) string {
// 	if name == "" {
// 		return ""
// 	}

// 	if ns.NoLowerCase {
// 		return name
// 	}

// 	var (
// 		value                          = commonInitialismsReplacer.Replace(name)
// 		buf                            strings.Builder
// 		lastCase, nextCase, nextNumber bool // upper case == true
// 		curCase                        = value[0] <= 'Z' && value[0] >= 'A'
// 	)

// 	for i, v := range value[:len(value)-1] {
// 		nextCase = value[i+1] <= 'Z' && value[i+1] >= 'A'
// 		nextNumber = value[i+1] >= '0' && value[i+1] <= '9'

// 		if curCase {
// 			if lastCase && (nextCase || nextNumber) {
// 				buf.WriteRune(v + 32)
// 			} else {
// 				if i > 0 && value[i-1] != '_' && value[i+1] != '_' {
// 					buf.WriteByte('_')
// 				}
// 				buf.WriteRune(v + 32)
// 			}
// 		} else {
// 			buf.WriteRune(v)
// 		}

// 		lastCase = curCase
// 		curCase = nextCase
// 	}

// 	if curCase {
// 		if !lastCase && len(value) > 1 {
// 			buf.WriteByte('_')
// 		}
// 		buf.WriteByte(value[len(value)-1] + 32)
// 	} else {
// 		buf.WriteByte(value[len(value)-1])
// 	}
// 	ret := buf.String()
// 	return ret
// }
