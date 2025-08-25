# TiCDC Remote Storage Sink Support
# TiCDC Remote Storage Sink Support

## Background

TiCDC is a change data capture component for TiDB that provides the ability to replicate data changes from TiDB to downstream systems. As of TiDB v5.4.0, TiCDC supports writing directly to remote storage services such as Amazon S3, Google Cloud Storage, Azure Blob Storage, and NFS. This feature is useful for creating data lakes, archiving data, or feeding data to other processing systems.

However, TiDB Operator currently lacks the ability to configure TiCDC to write to these remote storage services. This proposal aims to add support for configuring TiCDC sink to remote storage in TiDB Operator.

## Proposal

### TidbCluster CRD Changes

We propose to add a new `Sink` field to the `TiCDCSpec` section in the `TidbCluster` CRD. This field will allow users to configure remote storage sinks for TiCDC.
## Background

TiCDC is a change data capture component for TiDB that provides the ability to replicate data changes from TiDB to downstream systems. As of TiDB v5.4.0, TiCDC supports writing directly to remote storage services such as Amazon S3, Google Cloud Storage, Azure Blob Storage, and NFS. This feature is useful for creating data lakes, archiving data, or feeding data to other processing systems.

However, TiDB Operator currently lacks the ability to configure TiCDC to write to these remote storage services. This proposal aims to add support for configuring TiCDC sink to remote storage in TiDB Operator.

## Proposal

### TidbCluster CRD Changes

We propose to add a new `Sink` field to the `TiCDCSpec` section in the `TidbCluster` CRD. This field will allow users to configure remote storage sinks for TiCDC.
