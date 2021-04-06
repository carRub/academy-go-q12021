package service

import (
	"reflect"
	"sync"
	"testing"

	"github.com/carRub/academy-go-q12021/model"
)

func TestNewCharacterService(t *testing.T) {
	type args struct {
		url  string
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCharacterService(tt.args.url, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCharacterService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCharacterService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetCharacters(t *testing.T) {
	type fields struct {
		url  string
		file string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Character
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				url:  tt.fields.url,
				file: tt.fields.file,
			}
			got, err := s.GetCharacters()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetCharacters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetCharacterByID(t *testing.T) {
	type fields struct {
		url  string
		file string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Character
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				url:  tt.fields.url,
				file: tt.fields.file,
			}
			got, err := s.GetCharacterByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetCharacterByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCharacterByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_InsertExternalCharacter(t *testing.T) {
	type fields struct {
		url  string
		file string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Character
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				url:  tt.fields.url,
				file: tt.fields.file,
			}
			got, err := s.InsertExternalCharacter(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.InsertExternalCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.InsertExternalCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetCharactersConcurrently(t *testing.T) {
	type fields struct {
		url  string
		file string
	}
	type args struct {
		t               string
		items           int
		itemsPerWorkers int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Character
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				url:  tt.fields.url,
				file: tt.fields.file,
			}
			got, err := s.GetCharactersConcurrently(tt.args.t, tt.args.items, tt.args.itemsPerWorkers)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetCharactersConcurrently() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCharactersConcurrently() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOfType(t *testing.T) {
	type args struct {
		t  string
		id int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOfType(tt.args.t, tt.args.id); got != tt.want {
				t.Errorf("isOfType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readRecordFromCsv(t *testing.T) {
	type args struct {
		s  *Service
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    model.Character
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readRecordFromCsv(tt.args.s, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("readRecordFromCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readRecordFromCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readAllRecordsFromCsv(t *testing.T) {
	type args struct {
		s *Service
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Character
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readAllRecordsFromCsv(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("readAllRecordsFromCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readAllRecordsFromCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worker(t *testing.T) {
	type args struct {
		jobs       <-chan model.Character
		shutdown   <-chan struct{}
		results    chan model.Character
		wg         *sync.WaitGroup
		t          string
		fileLength int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			worker(tt.args.jobs, tt.args.shutdown, tt.args.results, tt.args.wg, tt.args.t, tt.args.fileLength)
		})
	}
}
