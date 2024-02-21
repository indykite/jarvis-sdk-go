// Copyright (c) 2023 IndyKite
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ingest_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/mock/gomock"

	ingestpb "github.com/indykite/indykite-sdk-go/gen/indykite/ingest/v1beta3"
	knowledgeobjects "github.com/indykite/indykite-sdk-go/gen/indykite/knowledge/objects/v1beta1"
	objects "github.com/indykite/indykite-sdk-go/gen/indykite/objects/v1beta2"
	"github.com/indykite/indykite-sdk-go/ingest"
	ingestm "github.com/indykite/indykite-sdk-go/test/ingest/v1beta3"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var (
	record1 = &ingestpb.Record{
		Id: "1",
		Operation: &ingestpb.Record_Upsert{
			Upsert: &ingestpb.UpsertData{
				Data: &ingestpb.UpsertData_Relationship{
					Relationship: &ingestpb.Relationship{
						Source: &ingestpb.NodeMatch{
							ExternalId: "0000",
							Type:       "Employee",
						},
						Target: &ingestpb.NodeMatch{
							ExternalId: "0001",
							Type:       "Truck",
						},
						Type: "SERVICES",
						Properties: []*knowledgeobjects.Property{
							{
								Type: "since",
								Value: &objects.Value{
									Type: &objects.Value_StringValue{
										StringValue: "production",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	record2 = &ingestpb.Record{
		Id: "2",
		Operation: &ingestpb.Record_Upsert{
			Upsert: &ingestpb.UpsertData{
				Data: &ingestpb.UpsertData_Node{
					Node: &knowledgeobjects.Node{
						ExternalId: "0001",
						Type:       "Truck",
						IsIdentity: false,
						Properties: []*knowledgeobjects.Property{
							{
								Type: "vin",
								Value: &objects.Value{
									Type: &objects.Value_IntegerValue{
										IntegerValue: 1234,
									},
								},
							},
						},
					},
				},
			},
		},
	}
)

var _ = Describe("Ingest", func() {
	It("IngestRecord", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)

		client, err := ingest.NewTestClient(mockClient)
		Ω(err).To(Succeed())

		mockClient.EXPECT().IngestRecord(gomock.Any(), gomock.Eq(&ingestpb.IngestRecordRequest{
			Record: record1,
		}), gomock.Any()).Return(&ingestpb.IngestRecordResponse{
			RecordId: record1.Id,
			Info: &ingestpb.Info{Changes: []*ingestpb.Change{
				{
					Id:       "gid:...",
					DataType: ingestpb.DataType_DATA_TYPE_RELATIONSHIP,
				},
			}},
		}, nil)

		resp, err := client.IngestRecord(context.Background(), record1)
		Expect(err).To(Succeed())
		Expect(resp).To(PointTo(MatchFields(IgnoreExtras, Fields{
			"RecordId": Equal("1"),
			"Info": PointTo(MatchFields(IgnoreExtras, Fields{
				"Changes": ContainElement(PointTo(MatchFields(IgnoreExtras, Fields{
					"Id":       Not(BeEmpty()),
					"DataType": Equal(ingestpb.DataType_DATA_TYPE_RELATIONSHIP),
				}))),
			})),
		})))
	})

	It("StreamRecords", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)
		mockStreamClient := ingestm.NewMockIngestAPI_StreamRecordsClient(mockCtrl)

		client, err := ingest.NewTestClient(mockClient)
		Expect(err).To(Succeed())

		records := []*ingestpb.Record{
			record1, record2,
		}

		mockClient.EXPECT().StreamRecords(gomock.Any()).Return(mockStreamClient, nil)
		mockStreamClient.EXPECT().Send(&ingestpb.StreamRecordsRequest{Record: record1}).Return(nil)
		mockStreamClient.EXPECT().Recv().Return(&ingestpb.StreamRecordsResponse{
			RecordId:    "1",
			RecordIndex: 0,
			Info: &ingestpb.Info{Changes: []*ingestpb.Change{
				{
					Id:       "gid:...",
					DataType: ingestpb.DataType_DATA_TYPE_RELATIONSHIP,
				},
			}}}, nil)
		mockStreamClient.EXPECT().Send(&ingestpb.StreamRecordsRequest{Record: record2}).Return(nil)
		mockStreamClient.EXPECT().Recv().Return(&ingestpb.StreamRecordsResponse{
			RecordId:    "2",
			RecordIndex: 1,
			Info: &ingestpb.Info{Changes: []*ingestpb.Change{
				{
					Id:       "gid:...",
					DataType: ingestpb.DataType_DATA_TYPE_NODE,
				},
			}}}, nil)
		mockStreamClient.EXPECT().CloseSend()

		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()

		err = client.OpenStreamClient(ctx)
		Expect(err).To(Succeed())
		dataTypes := []ingestpb.DataType{
			ingestpb.DataType_DATA_TYPE_RELATIONSHIP,
			ingestpb.DataType_DATA_TYPE_NODE,
		}
		for i, record := range records {
			var resp *ingestpb.StreamRecordsResponse
			err = client.SendRecord(record)
			Expect(err).To(Succeed())
			resp, err = client.ReceiveResponse()
			Expect(err).To(Succeed())
			Expect(resp).To(PointTo(MatchFields(IgnoreExtras, Fields{
				"RecordId":    Equal(fmt.Sprintf("%d", i+1)),
				"RecordIndex": Equal(uint32(i)),
				"Info": PointTo(MatchFields(IgnoreExtras, Fields{
					"Changes": ContainElement(PointTo(MatchFields(IgnoreExtras, Fields{
						"Id":       Not(BeEmpty()),
						"DataType": Equal(dataTypes[i]),
					}))),
				})),
			})))
		}
		err = client.CloseStream()
		Expect(err).To(Succeed())
	})

	It("StreamRecords using helper", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)
		mockStreamClient := ingestm.NewMockIngestAPI_StreamRecordsClient(mockCtrl)

		client, err := ingest.NewTestClient(mockClient)
		Expect(err).To(Succeed())

		records := []*ingestpb.Record{
			record1, record2,
		}

		mockClient.EXPECT().StreamRecords(gomock.Any()).Return(mockStreamClient, nil)
		mockStreamClient.EXPECT().Send(&ingestpb.StreamRecordsRequest{Record: record1}).Return(nil)
		mockStreamClient.EXPECT().Recv().Return(&ingestpb.StreamRecordsResponse{
			RecordId:    "1",
			RecordIndex: 0,
			Info: &ingestpb.Info{Changes: []*ingestpb.Change{
				{
					Id:       "gid:...",
					DataType: ingestpb.DataType_DATA_TYPE_RELATIONSHIP,
				},
			}}}, nil)
		mockStreamClient.EXPECT().Send(&ingestpb.StreamRecordsRequest{Record: record2}).Return(nil)
		mockStreamClient.EXPECT().Recv().Return(&ingestpb.StreamRecordsResponse{
			RecordId:    "2",
			RecordIndex: 1,
			Info: &ingestpb.Info{Changes: []*ingestpb.Change{
				{
					Id:       "gid:...",
					DataType: ingestpb.DataType_DATA_TYPE_NODE,
				},
			}}}, nil)
		mockStreamClient.EXPECT().CloseSend()

		responses, err := client.StreamRecords(records)
		Expect(err).To(Succeed())
		dataTypes := []ingestpb.DataType{
			ingestpb.DataType_DATA_TYPE_RELATIONSHIP,
			ingestpb.DataType_DATA_TYPE_NODE,
		}
		for i, response := range responses {
			Expect(response).To(PointTo(MatchFields(IgnoreExtras, Fields{
				"RecordId":    Equal(fmt.Sprintf("%d", i+1)),
				"RecordIndex": Equal(uint32(i)),
				"Info": PointTo(MatchFields(IgnoreExtras, Fields{
					"Changes": ContainElement(PointTo(MatchFields(IgnoreExtras, Fields{
						"Id":       Not(BeEmpty()),
						"DataType": Equal(dataTypes[i]),
					}))),
				})),
			})))
		}
	})

	It("StreamRecords send before open", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)

		client, err := ingest.NewTestClient(mockClient)
		Expect(err).To(Succeed())

		err = client.SendRecord(record1)
		Expect(err).To(MatchError(ContainSubstring("a stream must be opened first")))
	})

	It("StreamRecords receive before open", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)

		client, err := ingest.NewTestClient(mockClient)
		Expect(err).To(Succeed())

		_, err = client.ReceiveResponse()
		Expect(err).To(MatchError(ContainSubstring("a stream must be opened first")))
	})

	It("StreamRecords close before open", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)

		client, err := ingest.NewTestClient(mockClient)
		Expect(err).To(Succeed())

		err = client.CloseStream()
		Expect(err).To(MatchError(ContainSubstring("the stream has already been closed")))
	})
})

var _ = Describe("Retry client", func() {
	It("StreamRecords", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)
		mockStreamClient := ingestm.NewMockIngestAPI_StreamRecordsClient(mockCtrl)

		client, err := ingest.NewTestRetryClient(mockClient, &ingest.RetryPolicy{
			MaxAttempts:       4,
			InitialBackoff:    1 * time.Second,
			BackoffMultiplier: 2,
		})
		Expect(err).To(Succeed())

		records := []*ingestpb.Record{
			record1, record2,
		}

		mockClient.EXPECT().StreamRecords(gomock.Any()).Return(mockStreamClient, nil)
		mockStreamClient.EXPECT().Send(&ingestpb.StreamRecordsRequest{Record: record1}).Return(nil)
		mockStreamClient.EXPECT().Recv().Return(&ingestpb.StreamRecordsResponse{
			RecordId:    "1",
			RecordIndex: 0,
			Info: &ingestpb.Info{Changes: []*ingestpb.Change{
				{
					Id:       "gid:...",
					DataType: ingestpb.DataType_DATA_TYPE_RELATIONSHIP,
				},
			}}}, nil)
		mockStreamClient.EXPECT().Send(&ingestpb.StreamRecordsRequest{Record: record2}).Return(nil)
		mockStreamClient.EXPECT().Recv().Return(&ingestpb.StreamRecordsResponse{
			RecordId:    "2",
			RecordIndex: 1,
			Info: &ingestpb.Info{Changes: []*ingestpb.Change{
				{
					Id:       "gid:...",
					DataType: ingestpb.DataType_DATA_TYPE_NODE,
				},
			}}}, nil)
		mockStreamClient.EXPECT().CloseSend()

		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()

		err = client.OpenStreamClient(ctx)
		Expect(err).To(Succeed())
		dataTypes := []ingestpb.DataType{
			ingestpb.DataType_DATA_TYPE_RELATIONSHIP,
			ingestpb.DataType_DATA_TYPE_NODE,
		}
		for i, record := range records {
			var resp *ingestpb.StreamRecordsResponse
			err = client.SendRecord(record)
			Expect(err).To(Succeed())
			resp, err = client.ReceiveResponse()
			Expect(err).To(Succeed())
			Expect(resp).To(PointTo(MatchFields(IgnoreExtras, Fields{
				"RecordId":    Equal(fmt.Sprintf("%d", i+1)),
				"RecordIndex": Equal(uint32(i)),
				"Info": PointTo(MatchFields(IgnoreExtras, Fields{
					"Changes": ContainElement(PointTo(MatchFields(IgnoreExtras, Fields{
						"Id":       Not(BeEmpty()),
						"DataType": Equal(dataTypes[i]),
					}))),
				})),
			})))
		}
		err = client.CloseStream()
		Expect(err).To(Succeed())
	})
	It("StreamRecords send before open", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)

		client, err := ingest.NewTestRetryClient(mockClient, &ingest.RetryPolicy{
			MaxAttempts:       4,
			InitialBackoff:    1 * time.Second,
			BackoffMultiplier: 2,
		})
		Expect(err).To(Succeed())

		err = client.SendRecord(record1)
		Expect(err).To(MatchError(ContainSubstring("a stream must be opened first")))
	})
	It("StreamRecords receive before open", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)

		client, err := ingest.NewTestRetryClient(mockClient, &ingest.RetryPolicy{
			MaxAttempts:       4,
			InitialBackoff:    1 * time.Second,
			BackoffMultiplier: 2,
		})
		Expect(err).To(Succeed())

		_, err = client.ReceiveResponse()
		Expect(err).To(MatchError(ContainSubstring("a stream must be opened first")))
	})
	It("StreamRecords close before open", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)

		client, err := ingest.NewTestRetryClient(mockClient, &ingest.RetryPolicy{
			MaxAttempts:       4,
			InitialBackoff:    1 * time.Second,
			BackoffMultiplier: 2,
		})
		Expect(err).To(Succeed())

		err = client.CloseStream()
		Expect(err).To(MatchError(ContainSubstring("the stream has already been closed")))
	})
	It("StreamRecords retry connection", func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockClient := ingestm.NewMockIngestAPIClient(mockCtrl)
		mockStreamClient := ingestm.NewMockIngestAPI_StreamRecordsClient(mockCtrl)

		client, err := ingest.NewTestRetryClient(mockClient, &ingest.RetryPolicy{
			MaxAttempts:       2,
			InitialBackoff:    1 * time.Second,
			BackoffMultiplier: 1,
		})
		Expect(err).To(Succeed())

		mockClient.EXPECT().StreamRecords(gomock.Any()).Times(3).Return(mockStreamClient, nil)
		mockStreamClient.EXPECT().
			Send(&ingestpb.StreamRecordsRequest{Record: record1}).
			Times(3).
			Return(errors.New("something went wrong"))
		mockStreamClient.EXPECT().Recv().AnyTimes().Return(nil, errors.New("something went wrong"))
		mockStreamClient.EXPECT().CloseSend()

		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()

		err = client.OpenStreamClient(ctx)
		Expect(err).To(Succeed())

		go func() {
			err = client.SendRecord(record1)
			Expect(err).To(MatchError(ContainSubstring("something went wrong")))
		}()

		_, err = client.ReceiveResponse()
		Expect(err).To(MatchError(ContainSubstring("unable to reconnect to server")))
	})
})
