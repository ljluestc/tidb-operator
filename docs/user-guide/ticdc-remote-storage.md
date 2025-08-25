# TiCDC Remote Storage Sink

TiCDC is a change data capture component for TiDB that can replicate data changes to various downstream systems. Starting from TiDB v5.4.0, TiCDC supports writing directly to remote storage services like Amazon S3, Google Cloud Storage, Azure Blob Storage, and NFS.

This document describes how to configure TiCDC to write to remote storage services using TiDB Operator.

## Prerequisites

- TiDB Operator version >= 1.3.0
- TiDB version >= v5.4.0

## Configuration

To enable TiCDC to write to remote storage, add a `sink` section to the `spec.ticdc` field in your `TidbCluster` custom resource.

### Amazon S3
