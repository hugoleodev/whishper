// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TranscriptionsColumns holds the columns for the "transcriptions" table.
	TranscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString, Default: "pending"},
		{Name: "diarize", Type: field.TypeBool, Default: false},
		{Name: "language", Type: field.TypeString, Default: "auto"},
		{Name: "task", Type: field.TypeString, Default: "transcribe"},
		{Name: "device", Type: field.TypeString, Default: "cpu"},
		{Name: "model_size", Type: field.TypeString, Default: "small"},
		{Name: "source_url", Type: field.TypeString, Nullable: true},
		{Name: "file_name", Type: field.TypeString, Nullable: true},
		{Name: "result", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// TranscriptionsTable holds the schema information for the "transcriptions" table.
	TranscriptionsTable = &schema.Table{
		Name:       "transcriptions",
		Columns:    TranscriptionsColumns,
		PrimaryKey: []*schema.Column{TranscriptionsColumns[0]},
	}
	// TranslationsColumns holds the columns for the "translations" table.
	TranslationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "source_language", Type: field.TypeString},
		{Name: "target_language", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt},
		{Name: "result", Type: field.TypeJSON},
		{Name: "transcription_translations", Type: field.TypeInt, Nullable: true},
	}
	// TranslationsTable holds the schema information for the "translations" table.
	TranslationsTable = &schema.Table{
		Name:       "translations",
		Columns:    TranslationsColumns,
		PrimaryKey: []*schema.Column{TranslationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "translations_transcriptions_translations",
				Columns:    []*schema.Column{TranslationsColumns[5]},
				RefColumns: []*schema.Column{TranscriptionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TranscriptionsTable,
		TranslationsTable,
	}
)

func init() {
	TranslationsTable.ForeignKeys[0].RefTable = TranscriptionsTable
}