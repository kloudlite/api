package app

import (
	"fmt"
	"github.com/kloudlite/api/apps/infra/internal/domain"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/nats"
)

type ResourceEventPublisherImpl struct {
	cli    *nats.Client
	logger logging.Logger
}

func (r *ResourceEventPublisherImpl) PublishClusterEvent(cluster *entities.Cluster, msg domain.PublishMsg) {
	subject := clusterResUpdateSubject(cluster)
	if err := r.cli.Conn.Publish(subject, []byte(msg)); err != nil {
		r.logger.Errorf(err, "failed to publish message to subject %q", subject)
	}
}

func (r *ResourceEventPublisherImpl) PublishNodePoolEvent(np *entities.NodePool, msg domain.PublishMsg) {
	subject := nodePoolResUpdateSubject(np)
	if err := r.cli.Conn.Publish(subject, []byte(msg)); err != nil {
		r.logger.Errorf(err, "failed to publish message to subject %q", subject)
	}
}

func (r *ResourceEventPublisherImpl) PublishVpnDeviceEvent(dev *entities.VPNDevice, msg domain.PublishMsg) {
	subject := vpnDeviceResUpdateSubject(dev)
	if err := r.cli.Conn.Publish(subject, []byte(msg)); err != nil {
		r.logger.Errorf(err, "failed to publish message to subject %q", subject)
	}
}

func (r *ResourceEventPublisherImpl) PublishDomainResEvent(domain *entities.DomainEntry, msg domain.PublishMsg) {
	subject := domainResUpdateSubject(domain)
	if err := r.cli.Conn.Publish(subject, []byte(msg)); err != nil {
		r.logger.Errorf(err, "failed to publish message to subject %q", subject)
	}
}

func (r *ResourceEventPublisherImpl) PublishPvcResEvent(pvc *entities.PersistentVolumeClaim, msg domain.PublishMsg) {
	subject := pvcResUpdateSubject(pvc)
	if err := r.cli.Conn.Publish(subject, []byte(msg)); err != nil {
		r.logger.Errorf(err, "failed to publish message to subject %q", subject)
	}
}

func (r *ResourceEventPublisherImpl) PublishCMSEvent(cms *entities.ClusterManagedService, msg domain.PublishMsg) {
	subject := clusterManagedServiceUpdateSubject(cms)
	if err := r.cli.Conn.Publish(subject, []byte(msg)); err != nil {
		r.logger.Errorf(err, "failed to publish message to subject %q", subject)
	}
}

func NewResourceEventPublisher(cli *nats.Client, logger logging.Logger) domain.ResourceEventPublisher {
	return &ResourceEventPublisherImpl{
		cli,
		logger,
	}
}

func clusterResUpdateSubject(cluster *entities.Cluster) string {
	return fmt.Sprintf("res-updates.account.%s.cluster.%s", cluster.AccountName, cluster.Cluster.Name)
}

func nodePoolResUpdateSubject(nodePool *entities.NodePool) string {
	return fmt.Sprintf("res-updates.account.%s.cluster.%s.node-pool.%s", nodePool.AccountName, nodePool.ClusterName, nodePool.Name)
}

func domainResUpdateSubject(domainEntry *entities.DomainEntry) string {
	return fmt.Sprintf("res-updates.account.%s.cluster.%s.domain.%s", domainEntry.AccountName, domainEntry.ClusterName, domainEntry.DomainName)
}

func vpnDeviceResUpdateSubject(device *entities.VPNDevice) string {
	return fmt.Sprintf("res-updates.account.%s.cluster.%s.vpn-device.%s", device.AccountName, device.ClusterName, device.Name)
}

func pvcResUpdateSubject(pvc *entities.PersistentVolumeClaim) string {
	return fmt.Sprintf("res-updates.account.%s.cluster.%s.vpn-device.%s", pvc.AccountName, pvc.ClusterName, pvc.Name)
}

func clusterManagedServiceUpdateSubject(cmsvc *entities.ClusterManagedService) string {
	return fmt.Sprintf("res-updates.account.%s.cluster.%s.cluster-managed-service.%s", cmsvc.AccountName, cmsvc.ClusterName, cmsvc.Name)
}