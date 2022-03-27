package core

import (
	"errors"
)

type VideoActivity struct{ Checkpoints []VideoCheckpoint }

type VideoRecord struct {
	userId      string
	videoId     string
	checkpoints map[uint]VideoCheckpoint
}

type VideoCheckpoint struct {
	At      uint
	Elapsed uint
}

type VideoTracker struct {
	videos VideoRepository
}

type VideoRepository interface {
	LoadRecord(userId string, videoId string) *VideoRecord
	SaveRecord(userId string, videoId string, record *VideoRecord)
}

type inMemVideoRepository struct {
	videos map[string]VideoRecord
}

func newInMemVideoRepository() VideoRepository {
	return &inMemVideoRepository{
		videos: make(map[string]VideoRecord),
	}
}

func (r *inMemVideoRepository) LoadRecord(userId string, videoId string) *VideoRecord {
	if record, exists := r.videos[userId+"-"+videoId]; exists {
		return &record
	}

	return nil
}

func (r *inMemVideoRepository) SaveRecord(userId string, videoId string, record *VideoRecord) {
	r.videos[userId+"-"+videoId] = *record
}

func (r *VideoTracker) UpdateCheckpoint(userId string, videoId string, checkpoint VideoCheckpoint) (
	*VideoActivity,
	error,
) {
	if userId == "" {
		return nil, errors.New("userId can't be nil")
	}
	if videoId == "" {
		return nil, errors.New("videoId can't be nil")
	}
	if checkpoint.At < checkpoint.Elapsed {
		return nil, errors.New("checkpoint.At should be equal or greater than checkpointElapsed")
	}
	record := r.videos.LoadRecord(userId, videoId)
	if record == nil {
		checkpoints := make(map[uint]VideoCheckpoint)
		record = &VideoRecord{userId: userId, videoId: videoId, checkpoints: checkpoints}
	}
	if record.checkpoints == nil {
		record.checkpoints = make(map[uint]VideoCheckpoint)
	}
	record.checkpoints[checkpoint.At] = checkpoint
	r.videos.SaveRecord(userId, videoId, record)

	var checkpoints []VideoCheckpoint
	for _, c := range record.checkpoints {
		checkpoints = append(checkpoints, c)
	}
	return &VideoActivity{Checkpoints: checkpoints}, nil
}

func (r *VideoTracker) GetActivity(userId string, videoId string) (*VideoActivity, error) {
	if userId == "" {
		return nil, errors.New("userId can't be nil")
	}
	if videoId == "" {
		return nil, errors.New("videoId can't be nil")
	}
	record := r.videos.LoadRecord(userId, videoId)
	if record == nil {
		return nil, nil
	}
	if record.checkpoints == nil {
		return nil, nil
	}

	var checkpoints []VideoCheckpoint
	for _, c := range record.checkpoints {
		checkpoints = append(checkpoints, c)
	}
	return &VideoActivity{Checkpoints: checkpoints}, nil
}

func NewVideoTracker(repository VideoRepository) *VideoTracker {
	if repository == nil {
		repository = newInMemVideoRepository()
	}
	return &VideoTracker{videos: newInMemVideoRepository()}
}
