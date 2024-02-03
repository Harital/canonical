package shredder

import (
	"errors"
	"go/internal/osInterface"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShredder(t *testing.T) {
	tests := []struct {
		name string
		args string
		mock func(m *osInterface.MockOsInterface)
		want error
	}{
		{
			name: "shred file open error",
			args: "/tmp/randomfile",
			mock: func(m *osInterface.MockOsInterface) {
				m.On("OpenFile", "/tmp/randomfile", os.O_WRONLY|os.O_CREATE, fs.ModePerm).Return(nil, errors.New("random error"))
			},
		},
		/* other tests cases should be implemented here */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOs := &osInterface.MockOsInterface{}
			tt.mock(mockOs)

			shredder := NewShredder(mockOs)
			err := shredder.Shred(tt.args)
			if tt.want != nil {
				assert.Equal(t, tt.want, err)
			}
			mockOs.AssertExpectations(t)

		})
	}

}
