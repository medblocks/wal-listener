package listener

import "testing"

func TestEvent_GetSubjectName(t *testing.T) {
	type fields struct {
		Schema string
		Table  string
		Action string
		Data   map[string]interface{}
	}
	type args struct {
		topic string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "empty topic return default format with scheam and table",
			fields: fields{
				Schema: "public",
				Table:  "users",
				Action: "insert",
				Data:   nil,
			},
			args: args{},
			want: "public_users",
		},
		{
			name: "topic without formatting parameters is returned as is",
			fields: fields{
				Schema: "public",
				Table:  "users",
				Action: "insert",
				Data:   nil,
			},
			args: args{
				topic: "constant_topic",
			},
			want: "constant_topic",
		},
		{
			name: "topic with single formatting parameters includes the table name",
			fields: fields{
				Schema: "public",
				Table:  "users",
				Action: "insert",
				Data:   nil,
			},
			args: args{
				topic: "app_%s",
			},
			want: "app_users",
		},
		{
			name: "topic with two formatting parameters includes the schema and table name",
			fields: fields{
				Schema: "public",
				Table:  "users",
				Action: "insert",
				Data:   nil,
			},
			args: args{
				topic: "app_%s_%s",
			},
			want: "app_public_users",
		},
		{
			name: "topic with more than two formatting parameters uses the default",
			fields: fields{
				Schema: "public",
				Table:  "users",
				Action: "insert",
				Data:   nil,
			},
			args: args{
				topic: "app_%s_%s_%s",
			},
			want: "public_users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				Schema: tt.fields.Schema,
				Table:  tt.fields.Table,
				Action: tt.fields.Action,
				Data:   tt.fields.Data,
			}
			if got := e.GetSubjectName(tt.args.topic); got != tt.want {
				t.Errorf("GetSubjectName() = %v, want %v", got, tt.want)
			}
		})
	}
}
