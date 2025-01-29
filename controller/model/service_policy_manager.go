/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package model

import (
	"github.com/openziti/storage/boltz"
	"ztna-core/ztna/common/pb/edge_cmd_pb"
	"ztna-core/ztna/controller/change"
	"ztna-core/ztna/controller/command"
	"ztna-core/ztna/controller/db"
	"ztna-core/ztna/controller/fields"
	"ztna-core/ztna/controller/models"
	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

func NewServicePolicyManager(env Env) *ServicePolicyManager {
	manager := &ServicePolicyManager{
		baseEntityManager: newBaseEntityManager[*ServicePolicy, *db.ServicePolicy](env, env.GetStores().ServicePolicy),
	}
	manager.impl = manager

	RegisterManagerDecoder[*ServicePolicy](env, manager)

	return manager
}

type ServicePolicyManager struct {
	baseEntityManager[*ServicePolicy, *db.ServicePolicy]
}

func (self *ServicePolicyManager) newModelEntity() *ServicePolicy {
	return &ServicePolicy{}
}

func (self *ServicePolicyManager) Create(entity *ServicePolicy, ctx *change.Context) error {
	return DispatchCreate[*ServicePolicy](self, entity, ctx)
}

func (self *ServicePolicyManager) ApplyCreate(cmd *command.CreateEntityCommand[*ServicePolicy], ctx boltz.MutateContext) error {
	_, err := self.createEntity(cmd.Entity, ctx)
	return err
}

func (self *ServicePolicyManager) Update(entity *ServicePolicy, checker fields.UpdatedFields, ctx *change.Context) error {
	return DispatchUpdate[*ServicePolicy](self, entity, checker, ctx)
}

func (self *ServicePolicyManager) ApplyUpdate(cmd *command.UpdateEntityCommand[*ServicePolicy], ctx boltz.MutateContext) error {
	return self.updateEntity(cmd.Entity, cmd.UpdatedFields, ctx)
}

func (self *ServicePolicyManager) Marshall(entity *ServicePolicy) ([]byte, error) {
	tags, err := edge_cmd_pb.EncodeTags(entity.Tags)
	if err != nil {
		return nil, err
	}

	msg := &edge_cmd_pb.ServicePolicy{
		Id:                entity.Id,
		Name:              entity.Name,
		Tags:              tags,
		Semantic:          entity.Semantic,
		IdentityRoles:     entity.IdentityRoles,
		ServiceRoles:      entity.ServiceRoles,
		PostureCheckRoles: entity.PostureCheckRoles,
		PolicyType:        entity.PolicyType,
	}

	return proto.Marshal(msg)
}

func (self *ServicePolicyManager) Unmarshall(bytes []byte) (*ServicePolicy, error) {
	msg := &edge_cmd_pb.ServicePolicy{}
	if err := proto.Unmarshal(bytes, msg); err != nil {
		return nil, err
	}

	return &ServicePolicy{
		BaseEntity: models.BaseEntity{
			Id:   msg.Id,
			Tags: edge_cmd_pb.DecodeTags(msg.Tags),
		},
		Name:              msg.Name,
		Semantic:          msg.Semantic,
		IdentityRoles:     msg.IdentityRoles,
		ServiceRoles:      msg.ServiceRoles,
		PostureCheckRoles: msg.PostureCheckRoles,
		PolicyType:        msg.PolicyType,
	}, nil
}

type AssociatedIdsResult struct {
	ServiceIds      []string
	IdentityIds     []string
	PostureCheckIds []string
}

func (self *ServicePolicyManager) ListAssociatedIds(tx *bbolt.Tx, id string) *AssociatedIdsResult {
	return &AssociatedIdsResult{
		IdentityIds:     self.env.GetStores().ServicePolicy.GetRelatedEntitiesIdList(tx, id, db.EntityTypeIdentities),
		ServiceIds:      self.env.GetStores().ServicePolicy.GetRelatedEntitiesIdList(tx, id, db.EntityTypeServices),
		PostureCheckIds: self.env.GetStores().ServicePolicy.GetRelatedEntitiesIdList(tx, id, db.EntityTypePostureChecks),
	}
}
