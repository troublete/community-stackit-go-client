package validate_test

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/pkg/errors"
)

func TestUUID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{consts.SCHWARZ_ORGANIZATION_ID}, false},
		{"not ok", args{"bad-uuid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.UUID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrganizationID(t *testing.T) {
	type args struct {
		orgID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{consts.SCHWARZ_ORGANIZATION_ID}, false},
		{"not ok", args{"bad-uuid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.OrganizationID(tt.args.orgID); (err != nil) != tt.wantErr {
				t.Errorf("OrganizationID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectID(t *testing.T) {
	type args struct {
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"5dae0612-f5b1-4615-b7ca-b18796aa7e78"}, false},
		{"not ok", args{"bad-uuid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.ProjectID(tt.args.projectID); (err != nil) != tt.wantErr {
				t.Errorf("ProjectID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok", args{"project name!"}, true},
		{"ok [1]", args{"project name"}, false},
		{"ok [2]", args{"project-name"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.ProjectName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ProjectName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBillingRef(t *testing.T) {
	type args struct {
		billingRef string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok", args{"invalid!"}, true},
		{"ok", args{"T-123456B"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.BillingRef(tt.args.billingRef); (err != nil) != tt.wantErr {
				t.Errorf("BillingRef() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultResponseErrorHandler(t *testing.T) {
	r := io.NopCloser(strings.NewReader("ABC"))
	resp := &http.Response{StatusCode: 400, Body: r, ContentLength: 3, Request: &http.Request{URL: &url.URL{}}}
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{&http.Response{StatusCode: 202}}, false},
		{"not ok", args{resp}, true},
		{"not ok 2", args{resp}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.DefaultResponseErrorHandler(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("DefaultResponseErrorHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRole(t *testing.T) {
	type args struct {
		role string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"validate admin", args{consts.ROLE_PROJECT_ADMIN}, false},
		{"validate owner", args{consts.ROLE_PROJECT_OWNER}, false},
		{"validate auditor", args{consts.ROLE_PROJECT_AUDITOR}, false},
		{"validate member", args{consts.ROLE_PROJECT_MEMBER}, false},
		{"error", args{"something"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.Role(tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("Role() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResourceType(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"resource type project", args{consts.RESOURCE_TYPE_PROJECT}, false},
		{"resource type organization", args{consts.RESOURCE_TYPE_ORG}, false},
		{"resource type invalid", args{"something"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.ResourceType(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("ResourceType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserOrigin(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{consts.SCHWARZ_AUTH_ORIGIN}, false},
		{"bad origin", args{"something"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.UserOrigin(tt.args.origin); (err != nil) != tt.wantErr {
				t.Errorf("UserOrigin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSemVer(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"1.2.3"}, false},
		{"not ok 1", args{"ab1.2.3"}, true},
		{"not ok 2", args{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.SemVer(tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("SemVer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRFC3339(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"2020-09-04T00:00:00Z"}, false},
		{"not ok 1", args{"2020-09-04 00:00:00"}, true},
		{"not ok 2", args{"2020/09/04 00:00"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.RFC3339(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("RFC3339() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestISO8601(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"2020-09-04T00:00:00.605Z"}, false},
		{"not ok 1", args{"2020-09-04 00:00:00"}, true},
		{"not ok 2", args{"2020/09/04 00:00"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.ISO8601(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("ISO8601() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"1s"}, false},
		{"ok 2", args{"1m"}, false},
		{"ok 3", args{"60s"}, false},
		{"not ok 1", args{"abcd"}, true},
		{"not ok 2", args{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := validate.Duration(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Duration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetClientError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{"all ok", args{errors.New("test")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.SetClientError(tt.args.err); err.Error() != errors.Wrap(tt.args.err, "client validation error (Bad Request)").Error() {
				t.Errorf("SetClientError() error = %v, wantErr %v", err, errors.Wrap(tt.args.err, "client validation error (Bad Request)"))
			}
		})
	}
}
