package core_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gotesting.example/videotracker/core"
)

type stubRepo struct {
	lastSavedRecord *core.VideoRecord
}

func (r *stubRepo) LoadRecord(userId string, videoId string) *core.VideoRecord {
	return r.lastSavedRecord
}

func (r *stubRepo) SaveRecord(userId string, videoId string, record *core.VideoRecord) {
	r.lastSavedRecord = record
}

var _ = Describe("Core", func() {
	Context("Given a new VideoTracker", func() {
		var tracker *core.VideoTracker

		BeforeEach(func() {
			tracker = core.NewVideoTracker(nil)
		})
		It("When GetActivity with empty userId, Then should return error", func() {
			result, err := tracker.GetActivity("", "")
			Expect(result).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
		It("When GetActivity with empty videoId, Then return error", func() {
			result, err := tracker.GetActivity("U", "")
			Expect(result).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
		It("When GetActivity with non empty ids, Then return nil", func() {
			result, err := tracker.GetActivity("U", "V")
			Expect(result).To(BeNil())
			Expect(err).To(BeNil())
		})
		It("When UpdateCheckpoint with empty userId, Then should return error", func() {
			activity, err := tracker.UpdateCheckpoint("", "", core.VideoCheckpoint{})
			Expect(activity).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
		It("When UpdateCheckpoint with empty videoId, Then should return error", func() {
			activity, err := tracker.UpdateCheckpoint("U", "", core.VideoCheckpoint{})
			Expect(activity).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
		It("When UpdateCheckpoint with non empty ids, Then should return activity with 1 checkpoint", func() {
			activity, err := tracker.UpdateCheckpoint("U", "V", core.VideoCheckpoint{
				At:      10,
				Elapsed: 7,
			})
			Expect(activity).NotTo(BeNil())
			Expect(activity.Checkpoints).To(Equal([]core.VideoCheckpoint{{At: 10, Elapsed: 7}}))
			Expect(err).To(BeNil())
		})
		It("When UpdateCheckpoint with Elapsed > At, Then should return error", func() {
			activity, err := tracker.UpdateCheckpoint("U", "V", core.VideoCheckpoint{
				At:      2,
				Elapsed: 3,
			})
			Expect(activity).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
		It("When UpdateCheckpoint with non empty ids, Then GetActivity should return the same checkpoint", func() {
			activity, _ := tracker.UpdateCheckpoint("U", "V", core.VideoCheckpoint{
				At:      10,
				Elapsed: 7,
			})
			_, err := tracker.GetActivity("U", "V")
			Expect(activity).NotTo(BeNil())
			Expect(activity.Checkpoints).To(Equal([]core.VideoCheckpoint{{At: 10, Elapsed: 7}}))
			Expect(err).To(BeNil())
		})
	})

	Context("Given a VideoTracker with 1 checkpoint", func() {
		var tracker *core.VideoTracker

		BeforeEach(func() {
			tracker = core.NewVideoTracker(nil)
			_, _ = tracker.UpdateCheckpoint("U", "V", core.VideoCheckpoint{At: 10, Elapsed: 7})
		})
		It("When UpdateCheckpoint with equal `At` but greater `Elapsed` value, Then replace", func() {
			activity, err := tracker.UpdateCheckpoint("U", "V", core.VideoCheckpoint{
				At:      10,
				Elapsed: 9,
			})
			Expect(activity.Checkpoints).To(Equal([]core.VideoCheckpoint{{At: 10, Elapsed: 9}}))
			Expect(err).To(BeNil())
		})
		It("When UpdateCheckpoint with different `At`, Then keep both checkpoints", func() {
			activity, err := tracker.UpdateCheckpoint("U", "V", core.VideoCheckpoint{
				At:      13,
				Elapsed: 3,
			})
			Expect(len(activity.Checkpoints)).To(Equal(2))
			Expect(err).To(BeNil())
		})
	})

	Context("Given a user U1 with video V1 and 1 checkpoint", func() {
		var tracker *core.VideoTracker

		BeforeEach(func() {
			tracker = core.NewVideoTracker(nil)
			_, _ = tracker.UpdateCheckpoint("U1", "V1", core.VideoCheckpoint{At: 10, Elapsed: 7})
		})
		It("When U1 UpdateCheckpoint with video V2, Then each video have separate checkpoints", func() {
			_, _ = tracker.UpdateCheckpoint("U1", "V2", core.VideoCheckpoint{At: 13, Elapsed: 3})
			v1Activity, _ := tracker.GetActivity("U1", "V1")
			v2Activity, _ := tracker.GetActivity("U1", "V2")
			Expect(v1Activity.Checkpoints).To(Equal([]core.VideoCheckpoint{{At: 10, Elapsed: 7}}))
			Expect(v2Activity.Checkpoints).To(Equal([]core.VideoCheckpoint{{At: 13, Elapsed: 3}}))
		})
		It("When U2 UpdateCheckpoint with video V1, Then each video have separate checkpoints", func() {
			_, _ = tracker.UpdateCheckpoint("U2", "V1", core.VideoCheckpoint{At: 13, Elapsed: 3})
			u1Activity, _ := tracker.GetActivity("U1", "V1")
			u2Activity, _ := tracker.GetActivity("U2", "V1")
			Expect(u1Activity.Checkpoints).To(Equal([]core.VideoCheckpoint{{At: 10, Elapsed: 7}}))
			Expect(u2Activity.Checkpoints).To(Equal([]core.VideoCheckpoint{{At: 13, Elapsed: 3}}))
		})
	})

	Context("Given a new VideoTracker with external VideoRepository", func() {
		var tracker *core.VideoTracker
		var stub *stubRepo

		BeforeEach(func() {
			stub = &stubRepo{}
			tracker = core.NewVideoTracker(stub)
		})
		It("When UpdateCheckpoint, Then save to VideoRepository", func() {
			_, _ = tracker.UpdateCheckpoint("U", "V", core.VideoCheckpoint{At: 10, Elapsed: 7})
			Expect(stub.lastSavedRecord).ToNot(Equal(nil))
		})
	})
})
