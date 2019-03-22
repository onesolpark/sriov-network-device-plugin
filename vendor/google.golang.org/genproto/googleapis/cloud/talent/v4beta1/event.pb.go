// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/event.proto

package talent // import "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An enumeration of an event attributed to the behavior of the end user,
// such as a job seeker.
type JobEvent_JobEventType int32

const (
	// The event is unspecified by other provided values.
	JobEvent_JOB_EVENT_TYPE_UNSPECIFIED JobEvent_JobEventType = 0
	// The job seeker or other entity interacting with the service has
	// had a job rendered in their view, such as in a list of search results in
	// a compressed or clipped format. This event is typically associated with
	// the viewing of a jobs list on a single page by a job seeker.
	JobEvent_IMPRESSION JobEvent_JobEventType = 1
	// The job seeker, or other entity interacting with the service, has
	// viewed the details of a job, including the full description. This
	// event doesn't apply to the viewing a snippet of a job appearing as a
	// part of the job search results. Viewing a snippet is associated with an
	// [impression][google.cloud.talent.v4beta1.JobEvent.JobEventType.IMPRESSION]).
	JobEvent_VIEW JobEvent_JobEventType = 2
	// The job seeker or other entity interacting with the service
	// performed an action to view a job and was redirected to a different
	// website for job.
	JobEvent_VIEW_REDIRECT JobEvent_JobEventType = 3
	// The job seeker or other entity interacting with the service
	// began the process or demonstrated the intention of applying for a job.
	JobEvent_APPLICATION_START JobEvent_JobEventType = 4
	// The job seeker or other entity interacting with the service
	// submitted an application for a job.
	JobEvent_APPLICATION_FINISH JobEvent_JobEventType = 5
	// The job seeker or other entity interacting with the service
	// submitted an application for a job with a single click without
	// entering information. If a job seeker performs this action, send only
	// this event to the service. Do not also send
	// [JobEventType.APPLICATION_START][google.cloud.talent.v4beta1.JobEvent.JobEventType.APPLICATION_START]
	// or
	// [JobEventType.APPLICATION_FINISH][google.cloud.talent.v4beta1.JobEvent.JobEventType.APPLICATION_FINISH]
	// events.
	JobEvent_APPLICATION_QUICK_SUBMISSION JobEvent_JobEventType = 6
	// The job seeker or other entity interacting with the service
	// performed an action to apply to a job and was redirected to a different
	// website to complete the application.
	JobEvent_APPLICATION_REDIRECT JobEvent_JobEventType = 7
	// The job seeker or other entity interacting with the service began the
	// process or demonstrated the intention of applying for a job from the
	// search results page without viewing the details of the job posting.
	// If sending this event, JobEventType.VIEW event shouldn't be sent.
	JobEvent_APPLICATION_START_FROM_SEARCH JobEvent_JobEventType = 8
	// The job seeker, or other entity interacting with the service, performs an
	// action with a single click from the search results page to apply to a job
	// (without viewing the details of the job posting), and is redirected
	// to a different website to complete the application. If a candidate
	// performs this action, send only this event to the service. Do not also
	// send
	// [JobEventType.APPLICATION_START][google.cloud.talent.v4beta1.JobEvent.JobEventType.APPLICATION_START],
	// [JobEventType.APPLICATION_FINISH][google.cloud.talent.v4beta1.JobEvent.JobEventType.APPLICATION_FINISH]
	// or
	// [JobEventType.VIEW][google.cloud.talent.v4beta1.JobEvent.JobEventType.VIEW]
	// events.
	JobEvent_APPLICATION_REDIRECT_FROM_SEARCH JobEvent_JobEventType = 9
	// This event should be used when a company submits an application
	// on behalf of a job seeker. This event is intended for use by staffing
	// agencies attempting to place candidates.
	JobEvent_APPLICATION_COMPANY_SUBMIT JobEvent_JobEventType = 10
	// The job seeker or other entity interacting with the service demonstrated
	// an interest in a job by bookmarking or saving it.
	JobEvent_BOOKMARK JobEvent_JobEventType = 11
	// The job seeker or other entity interacting with the service was
	// sent a notification, such as an email alert or device notification,
	// containing one or more jobs listings generated by the service.
	JobEvent_NOTIFICATION JobEvent_JobEventType = 12
	// The job seeker or other entity interacting with the service was
	// employed by the hiring entity (employer). Send this event
	// only if the job seeker was hired through an application that was
	// initiated by a search conducted through the Cloud Talent Solution
	// service.
	JobEvent_HIRED JobEvent_JobEventType = 13
	// A recruiter or staffing agency submitted an application on behalf of the
	// candidate after interacting with the service to identify a suitable job
	// posting.
	JobEvent_SENT_CV JobEvent_JobEventType = 14
	// The entity interacting with the service (for example, the job seeker),
	// was granted an initial interview by the hiring entity (employer). This
	// event should only be sent if the job seeker was granted an interview as
	// part of an application that was initiated by a search conducted through /
	// recommendation provided by the Cloud Talent Solution service.
	JobEvent_INTERVIEW_GRANTED JobEvent_JobEventType = 15
	// The job seeker or other entity interacting with the service showed
	// no interest in the job.
	JobEvent_NOT_INTERESTED JobEvent_JobEventType = 16
)

var JobEvent_JobEventType_name = map[int32]string{
	0:  "JOB_EVENT_TYPE_UNSPECIFIED",
	1:  "IMPRESSION",
	2:  "VIEW",
	3:  "VIEW_REDIRECT",
	4:  "APPLICATION_START",
	5:  "APPLICATION_FINISH",
	6:  "APPLICATION_QUICK_SUBMISSION",
	7:  "APPLICATION_REDIRECT",
	8:  "APPLICATION_START_FROM_SEARCH",
	9:  "APPLICATION_REDIRECT_FROM_SEARCH",
	10: "APPLICATION_COMPANY_SUBMIT",
	11: "BOOKMARK",
	12: "NOTIFICATION",
	13: "HIRED",
	14: "SENT_CV",
	15: "INTERVIEW_GRANTED",
	16: "NOT_INTERESTED",
}
var JobEvent_JobEventType_value = map[string]int32{
	"JOB_EVENT_TYPE_UNSPECIFIED":       0,
	"IMPRESSION":                       1,
	"VIEW":                             2,
	"VIEW_REDIRECT":                    3,
	"APPLICATION_START":                4,
	"APPLICATION_FINISH":               5,
	"APPLICATION_QUICK_SUBMISSION":     6,
	"APPLICATION_REDIRECT":             7,
	"APPLICATION_START_FROM_SEARCH":    8,
	"APPLICATION_REDIRECT_FROM_SEARCH": 9,
	"APPLICATION_COMPANY_SUBMIT":       10,
	"BOOKMARK":                         11,
	"NOTIFICATION":                     12,
	"HIRED":                            13,
	"SENT_CV":                          14,
	"INTERVIEW_GRANTED":                15,
	"NOT_INTERESTED":                   16,
}

func (x JobEvent_JobEventType) String() string {
	return proto.EnumName(JobEvent_JobEventType_name, int32(x))
}
func (JobEvent_JobEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_event_9cc0bb8b89753637, []int{1, 0}
}

// The enum represents types of client events for a candidate profile.
type ProfileEvent_ProfileEventType int32

const (
	// Default value.
	ProfileEvent_PROFILE_EVENT_TYPE_UNSPECIFIED ProfileEvent_ProfileEventType = 0
	// The profile is displayed.
	ProfileEvent_IMPRESSION ProfileEvent_ProfileEventType = 1
	// The profile is viewed.
	ProfileEvent_VIEW ProfileEvent_ProfileEventType = 2
	// The profile is bookmarked.
	ProfileEvent_BOOKMARK ProfileEvent_ProfileEventType = 3
	// Comment added to the profile.
	ProfileEvent_COMMENT_ADDED ProfileEvent_ProfileEventType = 4
	// Email sent regarding the profile.
	ProfileEvent_EMAIL_SENT ProfileEvent_ProfileEventType = 5
	// The resume of the profile is viewed.
	ProfileEvent_RESUME_VIEWED ProfileEvent_ProfileEventType = 6
	// The profile is added to a job.
	ProfileEvent_ADD_TO_JOB ProfileEvent_ProfileEventType = 7
	// The recruiter explicitly highlights that a given profile is interesting
	// enough for later review or is a good result for the search.
	ProfileEvent_POSITIVE_REACTION_TO_PROFILE ProfileEvent_ProfileEventType = 8
	// The recruiter explicitly highlights that a given profile is not
	// interesting enough for later review or is not a good result for the
	// search.
	ProfileEvent_NEGATIVE_REACTION_TO_PROFILE ProfileEvent_ProfileEventType = 9
	// The candidate is interesting enough to set up an initial screen with the
	// recruiter. This step may be skipped based on the interview process.
	ProfileEvent_SCREEN ProfileEvent_ProfileEventType = 10
)

var ProfileEvent_ProfileEventType_name = map[int32]string{
	0:  "PROFILE_EVENT_TYPE_UNSPECIFIED",
	1:  "IMPRESSION",
	2:  "VIEW",
	3:  "BOOKMARK",
	4:  "COMMENT_ADDED",
	5:  "EMAIL_SENT",
	6:  "RESUME_VIEWED",
	7:  "ADD_TO_JOB",
	8:  "POSITIVE_REACTION_TO_PROFILE",
	9:  "NEGATIVE_REACTION_TO_PROFILE",
	10: "SCREEN",
}
var ProfileEvent_ProfileEventType_value = map[string]int32{
	"PROFILE_EVENT_TYPE_UNSPECIFIED": 0,
	"IMPRESSION":                     1,
	"VIEW":                           2,
	"BOOKMARK":                       3,
	"COMMENT_ADDED":                  4,
	"EMAIL_SENT":                     5,
	"RESUME_VIEWED":                  6,
	"ADD_TO_JOB":                     7,
	"POSITIVE_REACTION_TO_PROFILE":   8,
	"NEGATIVE_REACTION_TO_PROFILE":   9,
	"SCREEN":                         10,
}

func (x ProfileEvent_ProfileEventType) String() string {
	return proto.EnumName(ProfileEvent_ProfileEventType_name, int32(x))
}
func (ProfileEvent_ProfileEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_event_9cc0bb8b89753637, []int{2, 0}
}

// An event issued when an end user interacts with the application that
// implements Cloud Talent Solution. Providing this information improves the
// quality of search and recommendation for the API clients, enabling the
// service to perform optimally. The number of events sent must be consistent
// with other calls, such as job searches, issued to the service by the client.
type ClientEvent struct {
	// Required.
	//
	// A unique ID generated in the API responses. It can be found in
	// [ResponseMetadata.request_id][google.cloud.talent.v4beta1.ResponseMetadata.request_id].
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Required.
	//
	// A unique identifier, generated by the client application.
	EventId string `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Required.
	//
	// The timestamp of the event.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Required.
	//
	// The detail information of a specific event type.
	//
	// Types that are valid to be assigned to Event:
	//	*ClientEvent_JobEvent
	//	*ClientEvent_ProfileEvent
	Event isClientEvent_Event `protobuf_oneof:"event"`
	// Optional.
	//
	// Extra information about this event. Used for storing information with no
	// matching field in [event][google.cloud.talent.v4beta1.event] payload, for
	// example, user application specific context or details.
	//
	// At most 20 keys are supported. The maximum total size of all keys and
	// values is 2 KB.
	ExtraInfo            map[string]string `protobuf:"bytes,7,rep,name=extra_info,json=extraInfo,proto3" json:"extra_info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClientEvent) Reset()         { *m = ClientEvent{} }
func (m *ClientEvent) String() string { return proto.CompactTextString(m) }
func (*ClientEvent) ProtoMessage()    {}
func (*ClientEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_9cc0bb8b89753637, []int{0}
}
func (m *ClientEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEvent.Unmarshal(m, b)
}
func (m *ClientEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEvent.Marshal(b, m, deterministic)
}
func (dst *ClientEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEvent.Merge(dst, src)
}
func (m *ClientEvent) XXX_Size() int {
	return xxx_messageInfo_ClientEvent.Size(m)
}
func (m *ClientEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEvent proto.InternalMessageInfo

func (m *ClientEvent) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *ClientEvent) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *ClientEvent) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type isClientEvent_Event interface {
	isClientEvent_Event()
}

type ClientEvent_JobEvent struct {
	JobEvent *JobEvent `protobuf:"bytes,5,opt,name=job_event,json=jobEvent,proto3,oneof"`
}

type ClientEvent_ProfileEvent struct {
	ProfileEvent *ProfileEvent `protobuf:"bytes,6,opt,name=profile_event,json=profileEvent,proto3,oneof"`
}

func (*ClientEvent_JobEvent) isClientEvent_Event() {}

func (*ClientEvent_ProfileEvent) isClientEvent_Event() {}

func (m *ClientEvent) GetEvent() isClientEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ClientEvent) GetJobEvent() *JobEvent {
	if x, ok := m.GetEvent().(*ClientEvent_JobEvent); ok {
		return x.JobEvent
	}
	return nil
}

func (m *ClientEvent) GetProfileEvent() *ProfileEvent {
	if x, ok := m.GetEvent().(*ClientEvent_ProfileEvent); ok {
		return x.ProfileEvent
	}
	return nil
}

func (m *ClientEvent) GetExtraInfo() map[string]string {
	if m != nil {
		return m.ExtraInfo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClientEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ClientEvent_OneofMarshaler, _ClientEvent_OneofUnmarshaler, _ClientEvent_OneofSizer, []interface{}{
		(*ClientEvent_JobEvent)(nil),
		(*ClientEvent_ProfileEvent)(nil),
	}
}

func _ClientEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ClientEvent)
	// event
	switch x := m.Event.(type) {
	case *ClientEvent_JobEvent:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JobEvent); err != nil {
			return err
		}
	case *ClientEvent_ProfileEvent:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProfileEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClientEvent.Event has unexpected type %T", x)
	}
	return nil
}

func _ClientEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ClientEvent)
	switch tag {
	case 5: // event.job_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JobEvent)
		err := b.DecodeMessage(msg)
		m.Event = &ClientEvent_JobEvent{msg}
		return true, err
	case 6: // event.profile_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProfileEvent)
		err := b.DecodeMessage(msg)
		m.Event = &ClientEvent_ProfileEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClientEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ClientEvent)
	// event
	switch x := m.Event.(type) {
	case *ClientEvent_JobEvent:
		s := proto.Size(x.JobEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientEvent_ProfileEvent:
		s := proto.Size(x.ProfileEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An event issued when a job seeker interacts with the application that
// implements Cloud Talent Solution.
type JobEvent struct {
	// Required.
	//
	// The type of the event (see
	// [JobEventType][google.cloud.talent.v4beta1.JobEvent.JobEventType]).
	Type JobEvent_JobEventType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.talent.v4beta1.JobEvent_JobEventType" json:"type,omitempty"`
	// Required.
	//
	// The [job name(s)][google.cloud.talent.v4beta1.Job.name] associated with
	// this event. For example, if this is an
	// [impression][google.cloud.talent.v4beta1.JobEvent.JobEventType.IMPRESSION]
	// event, this field contains the identifiers of all jobs shown to the job
	// seeker. If this was a
	// [view][google.cloud.talent.v4beta1.JobEvent.JobEventType.VIEW] event, this
	// field contains the identifier of the viewed job.
	Jobs                 []string `protobuf:"bytes,2,rep,name=jobs,proto3" json:"jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobEvent) Reset()         { *m = JobEvent{} }
func (m *JobEvent) String() string { return proto.CompactTextString(m) }
func (*JobEvent) ProtoMessage()    {}
func (*JobEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_9cc0bb8b89753637, []int{1}
}
func (m *JobEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobEvent.Unmarshal(m, b)
}
func (m *JobEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobEvent.Marshal(b, m, deterministic)
}
func (dst *JobEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobEvent.Merge(dst, src)
}
func (m *JobEvent) XXX_Size() int {
	return xxx_messageInfo_JobEvent.Size(m)
}
func (m *JobEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_JobEvent.DiscardUnknown(m)
}

var xxx_messageInfo_JobEvent proto.InternalMessageInfo

func (m *JobEvent) GetType() JobEvent_JobEventType {
	if m != nil {
		return m.Type
	}
	return JobEvent_JOB_EVENT_TYPE_UNSPECIFIED
}

func (m *JobEvent) GetJobs() []string {
	if m != nil {
		return m.Jobs
	}
	return nil
}

// An event issued when a profile searcher interacts with the application
// that implements Cloud Talent Solution.
type ProfileEvent struct {
	// Required.
	//
	// Type of event.
	Type ProfileEvent_ProfileEventType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.talent.v4beta1.ProfileEvent_ProfileEventType" json:"type,omitempty"`
	// Required.
	//
	// The [profile name(s)][google.cloud.talent.v4beta1.Profile.name] associated
	// with this client event.
	Profiles []string `protobuf:"bytes,2,rep,name=profiles,proto3" json:"profiles,omitempty"`
	// Optional.
	//
	// The job ID associated with this client event if there is one. Leave it
	// empty if the event isn't associated with a job.
	//
	// The job ID should be consistent with the
	// [JobApplication.job.requisition_id][] in the profile.
	JobId                string   `protobuf:"bytes,3,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileEvent) Reset()         { *m = ProfileEvent{} }
func (m *ProfileEvent) String() string { return proto.CompactTextString(m) }
func (*ProfileEvent) ProtoMessage()    {}
func (*ProfileEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_9cc0bb8b89753637, []int{2}
}
func (m *ProfileEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileEvent.Unmarshal(m, b)
}
func (m *ProfileEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileEvent.Marshal(b, m, deterministic)
}
func (dst *ProfileEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileEvent.Merge(dst, src)
}
func (m *ProfileEvent) XXX_Size() int {
	return xxx_messageInfo_ProfileEvent.Size(m)
}
func (m *ProfileEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileEvent proto.InternalMessageInfo

func (m *ProfileEvent) GetType() ProfileEvent_ProfileEventType {
	if m != nil {
		return m.Type
	}
	return ProfileEvent_PROFILE_EVENT_TYPE_UNSPECIFIED
}

func (m *ProfileEvent) GetProfiles() []string {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *ProfileEvent) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientEvent)(nil), "google.cloud.talent.v4beta1.ClientEvent")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.talent.v4beta1.ClientEvent.ExtraInfoEntry")
	proto.RegisterType((*JobEvent)(nil), "google.cloud.talent.v4beta1.JobEvent")
	proto.RegisterType((*ProfileEvent)(nil), "google.cloud.talent.v4beta1.ProfileEvent")
	proto.RegisterEnum("google.cloud.talent.v4beta1.JobEvent_JobEventType", JobEvent_JobEventType_name, JobEvent_JobEventType_value)
	proto.RegisterEnum("google.cloud.talent.v4beta1.ProfileEvent_ProfileEventType", ProfileEvent_ProfileEventType_name, ProfileEvent_ProfileEventType_value)
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/event.proto", fileDescriptor_event_9cc0bb8b89753637)
}

var fileDescriptor_event_9cc0bb8b89753637 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xde, 0xc4, 0xf9, 0x3d, 0x49, 0xc3, 0xec, 0x68, 0x17, 0x85, 0xb0, 0x3f, 0x21, 0x02, 0x51,
	0x6e, 0x1c, 0x51, 0x90, 0x40, 0xbb, 0xdc, 0x38, 0xf6, 0x64, 0x3b, 0x6d, 0xfd, 0xc3, 0xd8, 0x0d,
	0x5a, 0x6e, 0x46, 0x4e, 0x33, 0x8d, 0x5c, 0x52, 0x4f, 0x48, 0xdc, 0x6a, 0xfb, 0x06, 0x5c, 0xf3,
	0x08, 0xbc, 0x06, 0xaf, 0xc5, 0x1d, 0x37, 0x68, 0xc6, 0x4e, 0xe5, 0xc2, 0x52, 0x55, 0x7b, 0x77,
	0x7e, 0xbe, 0xf3, 0x9d, 0xe3, 0x2f, 0x9f, 0x1d, 0xf8, 0x72, 0x29, 0xe5, 0x72, 0x25, 0xc6, 0x67,
	0x2b, 0x79, 0xb5, 0x18, 0x67, 0xf1, 0x4a, 0xa4, 0xd9, 0xf8, 0xfa, 0xdb, 0xb9, 0xc8, 0xe2, 0xaf,
	0xc7, 0xe2, 0x5a, 0xa4, 0x99, 0xb9, 0xde, 0xc8, 0x4c, 0xe2, 0x4f, 0x73, 0xa0, 0xa9, 0x81, 0x66,
	0x0e, 0x34, 0x0b, 0xe0, 0xe0, 0x59, 0xc1, 0x12, 0xaf, 0x93, 0x71, 0x9c, 0xa6, 0x32, 0x8b, 0xb3,
	0x44, 0xa6, 0xdb, 0x7c, 0x74, 0xf0, 0xb2, 0xe8, 0xea, 0x6c, 0x7e, 0x75, 0x3e, 0xce, 0x92, 0x4b,
	0xb1, 0xcd, 0xe2, 0xcb, 0x75, 0x0e, 0x18, 0xfd, 0x69, 0x40, 0xc7, 0x5e, 0x25, 0x22, 0xcd, 0x88,
	0xda, 0x88, 0x9f, 0x03, 0x6c, 0xc4, 0xaf, 0x57, 0x62, 0x9b, 0xf1, 0x64, 0xd1, 0xaf, 0x0c, 0x2b,
	0xfb, 0x6d, 0xd6, 0x2e, 0x2a, 0x74, 0x81, 0x3f, 0x81, 0x96, 0xbe, 0x4c, 0x35, 0xab, 0xba, 0xd9,
	0xd4, 0x39, 0x5d, 0xe0, 0xd7, 0xd0, 0x39, 0xdb, 0x88, 0x38, 0x13, 0x5c, 0xed, 0xe8, 0xd7, 0x86,
	0x95, 0xfd, 0xce, 0xc1, 0xc0, 0x2c, 0x6e, 0xdf, 0x1d, 0x60, 0x46, 0xbb, 0x03, 0x18, 0xe4, 0x70,
	0x55, 0xc0, 0x0e, 0xb4, 0x2f, 0xe4, 0x9c, 0x6b, 0xae, 0x7e, 0x5d, 0x8f, 0x7e, 0x61, 0xde, 0xf3,
	0xd8, 0xe6, 0x91, 0x9c, 0xeb, 0x83, 0x0f, 0x1f, 0xb1, 0xd6, 0x45, 0x11, 0xe3, 0x00, 0xf6, 0xd6,
	0x1b, 0x79, 0x9e, 0xac, 0x44, 0xc1, 0xd4, 0xd0, 0x4c, 0x5f, 0xdd, 0xcb, 0x14, 0xe4, 0x13, 0x3b,
	0xb6, 0xee, 0xba, 0x94, 0xe3, 0x19, 0x80, 0x78, 0x97, 0x6d, 0x62, 0x9e, 0xa4, 0xe7, 0xb2, 0xdf,
	0x1c, 0x1a, 0xfb, 0x9d, 0x83, 0xef, 0xee, 0xa5, 0x2b, 0x89, 0x69, 0x12, 0x35, 0x4a, 0xd3, 0x73,
	0x49, 0xd2, 0x6c, 0x73, 0xc3, 0xda, 0x62, 0x97, 0x0f, 0x7e, 0x80, 0xde, 0xdd, 0x26, 0x46, 0x60,
	0xfc, 0x22, 0x6e, 0x0a, 0xc5, 0x55, 0x88, 0x9f, 0x40, 0xfd, 0x3a, 0x5e, 0x5d, 0x89, 0x42, 0xe8,
	0x3c, 0x79, 0x55, 0xfd, 0xbe, 0x32, 0x69, 0x42, 0x5d, 0x3f, 0xdf, 0xe8, 0x6f, 0x03, 0x5a, 0x3b,
	0x25, 0xf0, 0x14, 0x6a, 0xd9, 0xcd, 0x5a, 0x68, 0x8a, 0xde, 0xc1, 0xc1, 0x83, 0xe4, 0xbb, 0x0d,
	0xa2, 0x9b, 0xb5, 0x60, 0x7a, 0x1e, 0x63, 0xa8, 0x5d, 0xc8, 0xf9, 0xb6, 0x5f, 0x1d, 0x1a, 0xfb,
	0x6d, 0xa6, 0xe3, 0xd1, 0xef, 0x06, 0x74, 0xcb, 0x50, 0xfc, 0x02, 0x06, 0x47, 0xfe, 0x84, 0x93,
	0x19, 0xf1, 0x22, 0x1e, 0xbd, 0x0d, 0x08, 0x3f, 0xf5, 0xc2, 0x80, 0xd8, 0x74, 0x4a, 0x89, 0x83,
	0x1e, 0xe1, 0x1e, 0x00, 0x75, 0x03, 0x46, 0xc2, 0x90, 0xfa, 0x1e, 0xaa, 0xe0, 0x16, 0xd4, 0x66,
	0x94, 0xfc, 0x84, 0xaa, 0xf8, 0x31, 0xec, 0xa9, 0x88, 0x33, 0xe2, 0x50, 0x46, 0xec, 0x08, 0x19,
	0xf8, 0x29, 0x3c, 0xb6, 0x82, 0xe0, 0x84, 0xda, 0x56, 0x44, 0x7d, 0x8f, 0x87, 0x91, 0xc5, 0x22,
	0x54, 0xc3, 0x1f, 0x03, 0x2e, 0x97, 0xa7, 0xd4, 0xa3, 0xe1, 0x21, 0xaa, 0xe3, 0x21, 0x3c, 0x2b,
	0xd7, 0x7f, 0x3c, 0xa5, 0xf6, 0x31, 0x0f, 0x4f, 0x27, 0x2e, 0xcd, 0xb7, 0x35, 0x70, 0x1f, 0x9e,
	0x94, 0x11, 0xb7, 0xab, 0x9a, 0xf8, 0x33, 0x78, 0xfe, 0x9f, 0x55, 0x7c, 0xca, 0x7c, 0x97, 0x87,
	0xc4, 0x62, 0xf6, 0x21, 0x6a, 0xe1, 0xcf, 0x61, 0xf8, 0xbe, 0xe1, 0x3b, 0xa8, 0xb6, 0x12, 0xa0,
	0x8c, 0xb2, 0x7d, 0x37, 0xb0, 0xbc, 0xb7, 0xf9, 0x19, 0x11, 0x02, 0xdc, 0x85, 0xd6, 0xc4, 0xf7,
	0x8f, 0x5d, 0x8b, 0x1d, 0xa3, 0x0e, 0x46, 0xd0, 0xf5, 0xfc, 0x88, 0x4e, 0x0b, 0x38, 0xea, 0xe2,
	0x36, 0xd4, 0x0f, 0x29, 0x23, 0x0e, 0xda, 0xc3, 0x1d, 0x68, 0x86, 0x4a, 0x46, 0x7b, 0x86, 0x7a,
	0x4a, 0x0b, 0xea, 0x45, 0x84, 0x69, 0x8d, 0xde, 0x30, 0xcb, 0x8b, 0x88, 0x83, 0x3e, 0xc2, 0x18,
	0x7a, 0x9e, 0x1f, 0x71, 0xdd, 0x22, 0xa1, 0xaa, 0xa1, 0xd1, 0x6f, 0x06, 0x74, 0xcb, 0xee, 0xc5,
	0xde, 0x1d, 0x07, 0xbc, 0x7a, 0xb0, 0xed, 0xef, 0x24, 0x25, 0x27, 0x0c, 0xa0, 0x55, 0xbc, 0x0d,
	0x3b, 0x37, 0xdc, 0xe6, 0xf8, 0x29, 0x34, 0xd4, 0x1b, 0x9b, 0x2c, 0xfa, 0x46, 0x6e, 0xcf, 0x0b,
	0x39, 0xa7, 0x8b, 0xd1, 0x5f, 0x15, 0x40, 0xff, 0x66, 0xc3, 0x23, 0x78, 0x11, 0x30, 0x7f, 0x4a,
	0x4f, 0xc8, 0x87, 0x18, 0xa6, 0xac, 0xa4, 0xa1, 0xec, 0x63, 0xfb, 0xae, 0xab, 0x58, 0x2c, 0xc7,
	0x21, 0x0e, 0xaa, 0xa9, 0x51, 0xe2, 0x5a, 0xf4, 0x84, 0x2b, 0x15, 0x51, 0x5d, 0x41, 0x18, 0x09,
	0x4f, 0x5d, 0xc2, 0x15, 0x03, 0x71, 0x50, 0x43, 0x41, 0x2c, 0xc7, 0xe1, 0x91, 0xcf, 0x8f, 0xfc,
	0x09, 0x6a, 0x2a, 0x0b, 0x05, 0x7e, 0x48, 0x23, 0x3a, 0x23, 0x9c, 0x11, 0xcb, 0xd6, 0xbf, 0x61,
	0xe4, 0xf3, 0xe2, 0x4c, 0xd4, 0x52, 0x08, 0x8f, 0xbc, 0xb1, 0xfe, 0x17, 0xd1, 0xc6, 0x00, 0x8d,
	0xd0, 0x66, 0x84, 0x78, 0x08, 0x26, 0xef, 0xe0, 0xe5, 0x99, 0xbc, 0xbc, 0x4f, 0xf0, 0x09, 0x68,
	0x3d, 0x02, 0xf5, 0x1d, 0x0c, 0x2a, 0x3f, 0x5b, 0x05, 0x74, 0x29, 0x57, 0x71, 0xba, 0x34, 0xe5,
	0x66, 0x39, 0x5e, 0x8a, 0x54, 0x7f, 0x25, 0xc7, 0x79, 0x2b, 0x5e, 0x27, 0xdb, 0xf7, 0xfe, 0x37,
	0xbc, 0xce, 0xd3, 0x3f, 0xaa, 0x86, 0x1d, 0x85, 0xf3, 0x86, 0x9e, 0xf9, 0xe6, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0a, 0x67, 0xd0, 0x69, 0x4e, 0x06, 0x00, 0x00,
}