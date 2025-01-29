# Release 1.3.0

## What's New

* Router Data Model enabled by default
* Bug fixes
* Controller Health Check HA Update (from @nenkoru)

## Router Data Model

As part of the controller HA work, a stripped down version of the data model can now be distributed to the routers, 
allowing routers to make some authorization/authentication decisions. This code has existed for some time, but
after testing and validation, is now enabled by default. 

It can still be disabled at the controller level using new configuration. Note that the router data model is required
for HA functionality, so if the controller is running in HA mode, it cannot be disabled. 

```yaml
routerDataModel:
  # Controls whether routers are told to enable functionality dependent on the router data model
  # Defaults to true
  enabled: true 

  # How many model changes to buffer so that routers can be updated iteratively. If a router requests
  # data that's no longer available, it will receive the full data model
  logSize: 10000
```

## HA Changes

Routers no longer require the `ha: enabled` flag be set in the configuration. Routers should work correctly
whether connecting to HA or non-HA controllers. 

NOTE: If the controller a router is connected changes modes, specifically if the controller goes from
      supporting the router data model to not, or vice-versa, the router will shutdown so that it can
      restart with the correct mode.

## Controller Health Check HA Update

This feature was contributed by @nenkoru.

The controller health check can now optionally return information about raft and leadership when the `/controller/raft` path is provided.

```
$ curl -k https://localhost:1280/health-checks/controller/raft
{
    "data": {
        "checks": [
            {
                "healthy": true,
                "id": "bolt.read",
                "lastCheckDuration": "0s",
                "lastCheckTime": "2025-01-14T19:42:13Z"
            }
        ],
        "healthy": true
    },
    "meta": {},
    "raft": {
        "isLeader": true,
        "isRaftEnabled": true
    }
}
```

Note the `raft` section, which indicates if raft is enabled and if the queried controller is currently the leader. If the 
`controller/raft` path isn't present in the request, the result should be unchanged from previous releases. 

When querying the controller/raft health, if raft is enabled but the controller is not the leader, the check will
return an HTTP status of 429.

## Component Updates and Bug Fixes

* github.com/openziti/agent: [v1.0.20 -> v1.0.23](https://github.com/openziti/agent/compare/v1.0.20...v1.0.23)
* github.com/openziti/channel/v3: [v3.0.16 -> v3.0.26](https://github.com/openziti/channel/compare/v3.0.16...v3.0.26)
* ztna-core/edge-api: [v0.26.35 -> v0.26.38](https://ztna-core/edge-api/compare/v0.26.35...v0.26.38)
    * [Issue #138](https://ztna-core/edge-api/issues/138) - management api deletes were generally not mapping 404 properly

* github.com/openziti/foundation/v2: [v2.0.52 -> v2.0.56](https://github.com/openziti/foundation/compare/v2.0.52...v2.0.56)
* github.com/openziti/identity: [v1.0.90 -> v1.0.94](https://github.com/openziti/identity/compare/v1.0.90...v1.0.94)
* github.com/openziti/metrics: [v1.2.61 -> v1.2.65](https://github.com/openziti/metrics/compare/v1.2.61...v1.2.65)
* github.com/openziti/runzmd: [v1.0.55 -> v1.0.59](https://github.com/openziti/runzmd/compare/v1.0.55...v1.0.59)
* github.com/openziti/secretstream: [v0.1.26 -> v0.1.28](https://github.com/openziti/secretstream/compare/v0.1.26...v0.1.28)
* github.com/openziti/storage: [v0.3.8 -> v0.3.15](https://github.com/openziti/storage/compare/v0.3.8...v0.3.15)
    * [Issue #91](https://github.com/openziti/storage/issues/91) - Support dashes in identifier segments after the first dot

* github.com/openziti/transport/v2: [v2.0.153 -> v2.0.159](https://github.com/openziti/transport/compare/v2.0.153...v2.0.159)
* ztna-core/ztna: [v1.2.2 -> v1.3.0](https://ztna-core/ztna/compare/v1.2.2...v1.3.0)
    * [Issue #2674](https://ztna-core/ztna/issues/2674) - 404 not found on well-known OIDC configuration with default ports/localhost
    * [Issue #2669](https://ztna-core/ztna/issues/2669) - Router api session tracker leaks memory.
    * [Issue #2659](https://ztna-core/ztna/issues/2659) - OIDC Login Panics On Unsupported Media Type
    * [Issue #2582](https://ztna-core/ztna/issues/2582) - An endpoint to determine whether a node is a raft leader
    * [Issue #2619](https://ztna-core/ztna/issues/2619) - Add source id to all events
    * [Issue #2644](https://ztna-core/ztna/issues/2644) - enhance mismapped external identity logging
    * [Issue #2636](https://ztna-core/ztna/issues/2636) - Enable HA smoketest
    * [Issue #2586](https://ztna-core/ztna/issues/2586) - Ziti Controller in HA mode doesn't update binding address in a bolt database after config changed
    * [Issue #2639](https://ztna-core/ztna/issues/2639) - Change cluster events namespace from fabric.cluster to cluster
    * [Issue #2184](https://ztna-core/ztna/issues/2184) - Add Event(s) For Controller Leader Connection State
    * [Issue #2548](https://ztna-core/ztna/issues/2548) - Generate a log message if the cluster is without a leader for some configurable period of time
    * [Issue #2624](https://ztna-core/ztna/issues/2624) - Remove uri/params from connect events
    * [Issue #2596](https://ztna-core/ztna/issues/2596) - Add DisableRouterDataModel config flag to controller
    * [Issue #2599](https://ztna-core/ztna/issues/2599) - Routers should only stream model data from one controller
    * [Issue #2232](https://ztna-core/ztna/issues/2232) - Standardized REST API Error For Mutation on Non-Consensus Controller
    * [Issue #2566](https://ztna-core/ztna/issues/2566) - Remove HA config flag from router
    * [Issue #2550](https://ztna-core/ztna/issues/2550) - Router Data Model Chaos Test
    * [Issue #2625](https://ztna-core/ztna/issues/2625) - edge sessions for an ERT may not be cleaned up when the ER/T is deleted 
    * [Issue #2591](https://ztna-core/ztna/issues/2591) - Split Edge APIs can cause `ziti edge login` to fail

# Release 1.2.2

## What's New

* Bug fixes and continuing progress on controller HA

## Component Updates and Bug Fixes

* github.com/openziti/secretstream: [v0.1.25 -> v0.1.26](https://github.com/openziti/secretstream/compare/v0.1.25...v0.1.26)
* github.com/openziti/storage: [v0.3.6 -> v0.3.8](https://github.com/openziti/storage/compare/v0.3.6...v0.3.8)
    * [Issue #87](https://github.com/openziti/storage/issues/87) - negative URL filter returns incorrect results

* ztna-core/ztna: [v1.2.1 -> v1.2.2](https://ztna-core/ztna/compare/v1.2.1...v1.2.2)
    * [Issue #2559](https://ztna-core/ztna/issues/2559) - expired JWTs are allowed to enroll
    * [Issue #2543](https://ztna-core/ztna/issues/2543) - Support adding adding an uninitialized node to a cluster (rather than the reverse)


# Release 1.2.1

## What's New

* Bug fixes and continuing progress on controller HA

## Component Updates and Bug Fixes

* github.com/openziti/agent: [v1.0.19 -> v1.0.20](https://github.com/openziti/agent/compare/v1.0.19...v1.0.20)
* github.com/openziti/channel/v3: [v3.0.10 -> v3.0.16](https://github.com/openziti/channel/compare/v3.0.10...v3.0.16)
* github.com/openziti/foundation/v2: [v2.0.50 -> v2.0.52](https://github.com/openziti/foundation/compare/v2.0.50...v2.0.52)
* github.com/openziti/identity: [v1.0.88 -> v1.0.90](https://github.com/openziti/identity/compare/v1.0.88...v1.0.90)
* github.com/openziti/metrics: [v1.2.59 -> v1.2.61](https://github.com/openziti/metrics/compare/v1.2.59...v1.2.61)
* github.com/openziti/runzmd: [v1.0.53 -> v1.0.55](https://github.com/openziti/runzmd/compare/v1.0.53...v1.0.55)
* github.com/openziti/storage: [v0.3.2 -> v0.3.6](https://github.com/openziti/storage/compare/v0.3.2...v0.3.6)
* github.com/openziti/transport/v2: [v2.0.150 -> v2.0.153](https://github.com/openziti/transport/compare/v2.0.150...v2.0.153)
* ztna-core/ztna: [v1.2.0 -> v1.2.1](https://ztna-core/ztna/compare/v1.2.0...v1.2.1)
    * [Issue #2543](https://ztna-core/ztna/issues/2543) - Support adding adding an uninitialized node to a cluster (rather than the reverse)
    * [Issue #2541](https://ztna-core/ztna/issues/2541) - Add cluster id, to prevent merging disparate clusters
    * [Issue #2532](https://ztna-core/ztna/issues/2532) - When adding an existing HA cluster member, remove/add if suffrage has changed
    * [Issue #2217](https://ztna-core/ztna/issues/2217) - Controller list is empty until peers connect
    * [Issue #2533](https://ztna-core/ztna/issues/2533) - Handle concurrent raft connections
    * [Issue #2534](https://ztna-core/ztna/issues/2534) - Ziti ID with leading hyphen causes command-line parameter ambiguity
    * [Issue #2528](https://ztna-core/ztna/issues/2528) - Updated router costs are not use when evaluating current path cost in the context of smart rerouting

# Release 1.2.0

## What's New

* New Router Metrics
* Changes to identity connect status
* HA Bootstrap Changes
* Connect Events
* SDK Events
* Bug fixes and other HA work

## New Router Metrics

The following new metrics are available for edge routers:

1. edge.connect.failures - meter tracking failed connect attempts from sdks
   This tracks failures to not having a valid token. Other failures which 
   happen earlier in the connection process may not be tracked here.
2. edge.connect.successes - meter tracking successful connect attempts from sdks
3. edge.disconnects - meter tracking disconnects of previously successfully connected sdks
4. edge.connections - gauge tracking count of currently connected sdks

## Identity Connect Status

Ziti tracks whether an identity is currently connected to an edge router. 
This is the `hasEdgeRouterConnection` field on Identity. 

Identity connection status used to be driven off of heartbeats from the edge router.
This feature doesn't work correctly when running with controller HA. 

To address this, while also providing more operation insight, connect events were added
(see below for more details on the events themselves).

The controller can be configured to use status from heartbeats, connect events or both.
If both are used as source, then if either reports the identity as connected, then it 
will show as connected. This is intended for when you have a mix of routers and they
don't all yet supported connect events.

The controller now also aims to be more precise about identity state. There is a new
field on Identity: `edgeRouterConnectionStatus`. This field can have one of three
values:

* offline
* online
* unknown

If the identity is reported as connected to any ER, it will be marked as `online`. 
If the identity has been reported as connected, but the reporting ER is now
offline, the identity may still be connected to the ER. While in this state
it will be marked as 'unknown'. After a configurable interval, it will be marked
as offline.

New controller config options:

```
identityStatusConfig:
  # valid values ['heartbeats', 'connect-events', 'hybrid']
  # defaults to 'hybrid' for now
  source: connect-events 

  # determines how often we scan for disconnected routers
  # defaults to 1 minute
  scanInterval: 1m

  # determines how long an identity will stay in unknown status before it's marked as offline
  # defaults to 5m
  unknownTimeout: 5m
```
  
## HA Bootstrapping Changes

Previously bootstrapping the RAFT cluster and initializing the controller with a 
default administrator were separate operations.
Now, the raft cluster will be bootstrapped whenever the controller is initialized. 

The controller can be initialized as follows:

1. Using `ziti agent controller init`
2. Using `ziti agent controller init-from-db`
3. Specifying a `db:` entry in the config file. This is equivalent to using `ziti agent controller init-from-db`.

Additionally:

1. `minClusterSize` has been removed. The cluster will always be initialized with a size of 1.
2. `bootstrapMembers` has been renamed to `initialMembers`. If `initialMembers` are specified,
   the bootstrapping controller will attempt to add them after bootstrap has been complete. If
   they are invalid they will be ignored. If they can't be reached (because they're not running
   yet), the controller will continue to retry until they are reached, or it is restarted.


## Connect Events

These are events generated when a successful connection is made to a controller, from any of:

1. Identity, using the REST API
2. Router
3. Controller (peer in an HA cluster)

They are also generated when an SDK connects to a router. 

**Controller Configuration**

```yml
events:
  jsonLogger:
    subscriptions:
      - type: connect
    handler:
      type: file
      format: json
      path: /tmp/ziti-events.log
```

**Router Configuration**
```yml
connectEvents:
  # defaults to true. 
  # If set to false, minimal information about which identities are connected will still be 
  # sent to the controller, so the `edgeRouterConnectionStatus` field can be populated, 
  # but connect events will not be generated.
  enabled: true

  # The interval at which connect information will be batched up and sent to the controller. 
  # Shorter intervals will improve data resolution on the controller. Longer intervals could
  # more efficient.
  batchInterval: 3s

  # The router will also periodically sent the full state to the controller, to ensure that 
  # it's in sync. It will do this automatically if the router gets disconnected from the 
  # controller, or if the router is unable to send a connect events messages to the controller.
  # This controls how often the full state will be sent under ordinairy conditions
  fullSyncInterval: 5m

  # If enabled is set to true, the router will collect connect events and send them out
  # at the configured batch interval. If there are a huge number of connecting identities
  # or if the router is disconnected from the controller for a time, it may be unable to
  # send events. In order to prevent queued events from exhausting memory, a maximum 
  # queue size is configured. 
  # Default value 100,000
  maxQueuedEvents: 100000
  
```

**Example Events**

```json
{
  "namespace": "connect",
  "src_type": "identity",
  "src_id": "ji2Rt8KJ4",
  "src_addr": "127.0.0.1:59336",
  "dst_id": "ctrl_client",
  "dst_addr": "localhost:1280/edge/management/v1/edge-routers/2L7NeVuGBU",
  "timestamp": "2024-10-02T12:17:39.501821249-04:00"
}
{
  "namespace": "connect",
  "src_type": "router",
  "src_id": "2L7NeVuGBU",
  "src_addr": "127.0.0.1:42702",
  "dst_id": "ctrl_client",
  "dst_addr": "127.0.0.1:6262",
  "timestamp": "2024-10-02T12:17:40.529865849-04:00"
}
{
  "namespace": "connect",
  "src_type": "peer",
  "src_id": "ctrl2",
  "src_addr": "127.0.0.1:40056",
  "dst_id": "ctrl1",
  "dst_addr": "127.0.0.1:6262",
  "timestamp": "2024-10-02T12:37:04.490859197-04:00"
}
```

## SDK Events

Building off of the connect events, there are events generated when an identity/sdk comes online or goes offline.

```yml
events:
  jsonLogger:
    subscriptions:
      - type: sdk
    handler:
      type: file
      format: json
      path: /tmp/ziti-events.log
```

```json
{
  "namespace": "sdk",
  "event_type" : "sdk-online",
  "identity_id": "ji2Rt8KJ4",
  "timestamp": "2024-10-02T12:17:39.501821249-04:00"
}

{
  "namespace": "sdk",
  "event_type" : "sdk-status-unknown",
  "identity_id": "ji2Rt8KJ4",
  "timestamp": "2024-10-02T12:17:40.501821249-04:00"
}

{
  "namespace": "sdk",
  "event_type" : "sdk-offline",
  "identity_id": "ji2Rt8KJ4",
  "timestamp": "2024-10-02T12:17:41.501821249-04:00"
}
```

## Component Updates and Bug Fixes

* github.com/openziti/agent: [v1.0.18 -> v1.0.19](https://github.com/openziti/agent/compare/v1.0.18...v1.0.19)
* github.com/openziti/channel/v3: [v3.0.5 -> v3.0.10](https://github.com/openziti/channel/compare/v3.0.5...v3.0.10)
* ztna-core/edge-api: [v0.26.32 -> v0.26.35](https://ztna-core/edge-api/compare/v0.26.32...v0.26.35)
* github.com/openziti/foundation/v2: [v2.0.49 -> v2.0.50](https://github.com/openziti/foundation/compare/v2.0.49...v2.0.50)
* github.com/openziti/identity: [v1.0.85 -> v1.0.88](https://github.com/openziti/identity/compare/v1.0.85...v1.0.88)

* github.com/openziti/metrics: [v1.2.58 -> v1.2.59](https://github.com/openziti/metrics/compare/v1.2.58...v1.2.59)
* github.com/openziti/runzmd: [v1.0.51 -> v1.0.53](https://github.com/openziti/runzmd/compare/v1.0.51...v1.0.53)
* github.com/openziti/sdk-golang: [v0.23.43 -> v0.23.44](https://github.com/openziti/sdk-golang/compare/v0.23.43...v0.23.44)
* github.com/openziti/transport/v2: [v2.0.146 -> v2.0.150](https://github.com/openziti/transport/compare/v2.0.146...v2.0.150)
* ztna-core/ztna: [v1.1.15 -> v1.2.0](https://ztna-core/ztna/compare/v1.1.15...v1.2.0)
    * [Issue #2212](https://ztna-core/ztna/issues/2212) - Rework distributed control bootstrap mechanism
    * [Issue #1835](https://ztna-core/ztna/issues/1835) - Add access log for rest and router connections
    * [Issue #2234](https://ztna-core/ztna/issues/2234) - Emit an event when hasEdgeRouterConnection state changes for an Identity
    * [Issue #2506](https://ztna-core/ztna/issues/2506) - Identity service config overrides referential integrity issues
    * [Issue #2491](https://ztna-core/ztna/issues/2491) - fix router CSR loading
    * [Issue #2478](https://ztna-core/ztna/issues/2478) - JWT signer secondary auth: not enough information to continue
    * [Issue #2482](https://ztna-core/ztna/issues/2482) - router run command - improperly binds 127.0.0.1:53/udp when tunnel mode is not tproxy
    * [Issue #2474](https://ztna-core/ztna/issues/2474) - Enable Ext JWT Enrollment/Generic Trust Bootstrapping
    * [Issue #2471](https://ztna-core/ztna/issues/2471) - Service Access for Legacy SDKs in  HA does not work
    * [Issue #2468](https://ztna-core/ztna/issues/2468) - enrollment signing cert is not properly identified

# Release 1.1.15

## What's New

* Bug fixes, enhancements and continuing progress on controller HA

## Component Updates and Bug Fixes

* github.com/openziti/channel/v3: [v3.0.3 -> v3.0.5](https://github.com/openziti/channel/compare/v3.0.3...v3.0.5)
    * [Issue #146](https://github.com/openziti/channel/issues/146) - Transport options aren't being set in dialer
    * [Issue #144](https://github.com/openziti/channel/issues/144) - Add ReadAdapter utility

* ztna-core/edge-api: [v0.26.31 -> v0.26.32](https://ztna-core/edge-api/compare/v0.26.31...v0.26.32)
* github.com/openziti/sdk-golang: [v0.23.42 -> v0.23.43](https://github.com/openziti/sdk-golang/compare/v0.23.42...v0.23.43)
    * [Issue #629](https://github.com/openziti/sdk-golang/issues/629) - JWT session refresh interprets expiration date incorrectly

* github.com/openziti/secretstream: [v0.1.24 -> v0.1.25](https://github.com/openziti/secretstream/compare/v0.1.24...v0.1.25)
* ztna-core/ztna: [v1.1.14 -> v1.1.15](https://ztna-core/ztna/compare/v1.1.14...v1.1.15)
    * [Issue #2460](https://ztna-core/ztna/issues/2460) - Panic on JWT token refresh

# Release 1.1.14

## What's New

* Bug fixes, enhancements and continuing progress on controller HA

## Component Updates and Bug Fixes

* ztna-core/edge-api: [v0.26.30 -> v0.26.31](https://ztna-core/edge-api/compare/v0.26.30...v0.26.31)
* github.com/openziti/jwks: [v1.0.5 -> v1.0.6](https://github.com/openziti/jwks/compare/v1.0.5...v1.0.6)
* ztna-core/ztna: [v1.1.13 -> v1.1.14](https://ztna-core/ztna/compare/v1.1.13...v1.1.14)
    * [Issue #2119](https://ztna-core/ztna/issues/2119) - Add authentication events
    * [Issue #2424](https://ztna-core/ztna/issues/2424) - Enabling any health check causes WARNING to be logged
    * [Issue #2454](https://ztna-core/ztna/issues/2454) - Fix release archive
    * [Issue #1479](https://ztna-core/ztna/issues/1479) - ziti edge list ... show paginated output but no suggestions on how to go to next page
    * [Issue #1420](https://ztna-core/ztna/issues/1420) - ziti-cli comma+space causes unhelpful error
    * [Issue #2207](https://ztna-core/ztna/issues/2207) - ziti edge login --token -- gets "username and password fields are required"
    * [Issue #2444](https://ztna-core/ztna/issues/2444) - Change default semantic for policies created from the CLI from AllOf to AnyOf

* github.com/openziti/xweb/v2: [v2.1.2 -> v2.1.3](https://github.com/openziti/xweb/compare/v2.1.2...v2.1.3)
  * [Issue #2454](https://ztna-core/ztna/issues/2454) - Fix release archive
  * [Issue #2429](https://ztna-core/ztna/issues/2429) - Controller configurations without default Edge API binding panics 
* ztna-core/ztna: [v1.1.12 -> v1.1.13](https://ztna-core/ztna/compare/v1.1.12...v1.1.13)
  * [Issue #2427](https://ztna-core/ztna/issues/2427) - Add low overhead xgress protocol for DTLS links
  * [Issue #2422](https://ztna-core/ztna/issues/2422) - Busy first hop links should backpressure to xgress senders
  * support using "\*" in host.v1/host.v2 allowedAddresses


# Release 1.1.13

This release will not be promoted, as a test binary was unintentionally released in the release archives.


# Release 1.1.12

## What's New

* Bug fixes, enhancements and continuing progress on controller HA
* Data corruption Fix

## Data Corruption Fix

Previous to version 1.1.12, the controller would not handle changes to the policy type of a service policy.
Specifically if the type was changed from Bind -> Dial, or Dial -> Bind, a set of denormalized data would
be left behind, leaving the permissions with the old policy type. 

Example:

1. Identity A has Bind access to service B via Bind service policy C. 
2. The policy type of service policy C is changed from Bind to Dial.
3. The service list would now likely show that Identity A has Dial and Bind access to service B, instead of
  just Dial access.

### Mitigation/Fixing Bad Data

If you encounter this problem, the easiest and safest way to solve the problem is to to delete and recreate
the affected service policy.

If changing policy types is something you do on a regular basis, and can't upgrade to a version with the fix,
you can work around the issue by deleting and recreating policies, instead of updating them. 

If you're not sure if you have ever changed a policy type, there is a database integrity check tool which can
 be run which looks for data integrity errors. It is run against a running system. 

Start the check using:

```
ziti fabric db start-check-integrity
```

This kicks off the operation in the background. The status of the check can be seen using:

```
ziti fabric db check-integrity-status 
```

By default this is a read-only operation. If the read-only run reports errors, it can be run 
with the `-f` flag, which will have it try to fix errors. The data integrity errors caused
by this bug should all be fixable by the integrity checker.

```
ziti fabric db start-check-integrity -f
```

**WARNINGS**: 
* Always make a database snapshot before running the integrity checker: `ziti db fabric snapshot <optional path`
* The integrity checker can be very resource intensive, depending on the size of your data model. 
  It is recommended that you run the integrity checker when the system is otherwise not busy.

## Component Updates and Bug Fixes

* github.com/openziti/agent: [v1.0.17 -> v1.0.18](https://github.com/openziti/agent/compare/v1.0.17...v1.0.18)
* github.com/openziti/channel/v3: [v2.0.143 -> v3.0.3](https://github.com/openziti/channel/compare/v2.0.143...v3.0.3)
    * [Issue #138](https://github.com/openziti/channel/issues/138) - Allow custom message serialization. Add support for a 'raw' message type.
    * [Issue #82](https://github.com/openziti/channel/issues/82) - Remove transport.Configuration from UnderlayFactory.Create

* ztna-core/edge-api: [v0.26.29 -> v0.26.30](https://ztna-core/edge-api/compare/v0.26.29...v0.26.30)
* github.com/openziti/foundation/v2: [v2.0.48 -> v2.0.49](https://github.com/openziti/foundation/compare/v2.0.48...v2.0.49)
* github.com/openziti/identity: [v1.0.84 -> v1.0.85](https://github.com/openziti/identity/compare/v1.0.84...v1.0.85)
* github.com/openziti/jwks: [v1.0.4 -> v1.0.5](https://github.com/openziti/jwks/compare/v1.0.4...v1.0.5)
    * [Issue #9](https://github.com/openziti/jwks/issues/9) - Using NewKey w/ RSA key results in nil pointer exception

* github.com/openziti/metrics: [v1.2.57 -> v1.2.58](https://github.com/openziti/metrics/compare/v1.2.57...v1.2.58)
* github.com/openziti/runzmd: [v1.0.50 -> v1.0.51](https://github.com/openziti/runzmd/compare/v1.0.50...v1.0.51)
* github.com/openziti/sdk-golang: [v0.23.40 -> v0.23.42](https://github.com/openziti/sdk-golang/compare/v0.23.40...v0.23.42)
    * [Issue #625](https://github.com/openziti/sdk-golang/issues/625) - traffic optimization: implement support for receiving multi-part edge payloads

* github.com/openziti/secretstream: [v0.1.21 -> v0.1.24](https://github.com/openziti/secretstream/compare/v0.1.21...v0.1.24)
* github.com/openziti/storage: [v0.3.0 -> v0.3.2](https://github.com/openziti/storage/compare/v0.3.0...v0.3.2)
* github.com/openziti/transport/v2: [v2.0.143 -> v2.0.146](https://github.com/openziti/transport/compare/v2.0.143...v2.0.146)
    * [Issue #92](https://github.com/openziti/transport/issues/92) - Implement simple traffic shaper

* github.com/openziti/xweb/v2: [v2.1.1 -> v2.1.2](https://github.com/openziti/xweb/compare/v2.1.1...v2.1.2)
* github.com/openziti-incubator/cf: v0.0.3 (new)
* github.com/openziti/dilithium: [v0.3.3 -> v0.3.5](https://github.com/openziti/dilithium/compare/v0.3.3...v0.3.5)
* ztna-core/ztna: [v1.1.11 -> v1.1.12](https://ztna-core/ztna/compare/v1.1.11...v1.1.12)
    * [Issue #2413](https://ztna-core/ztna/issues/2413) - Add db anonymization utility
    * [Issue #2415](https://ztna-core/ztna/issues/2415) - Fix policy denormalization when service policy type is changed
    * [Issue #2406](https://ztna-core/ztna/issues/2406) - ziti agent controller snapshot-db exit code is always successful
    * [Issue #2405](https://ztna-core/ztna/issues/2405) - Investigate Older SDKs Not Enrolling Not Connecting in HA
    * [Issue #2403](https://ztna-core/ztna/issues/2403) - Fix terminator costing concurrency issue
    * [Issue #2397](https://ztna-core/ztna/issues/2397) - JWKS endpoints w/ new keys do not get refreshed
    * [Issue #2390](https://ztna-core/ztna/issues/2390) - Update to github.com/openziti/channel/v3
    * [Issue #2388](https://ztna-core/ztna/issues/2388) - Remove use of ziti fabric add-identity commands in 004-controller-pki.md

# Release 1.1.11

# What's New

* This release updates to Go v1.23
* Updates to the latest version of golangci-lint, to allow it to work with the new version of Go
* Linter fixes to address issues caught by updated linter

# Release 1.1.10

## What's New

* Bug fixes, enhancements and continuing progress on controller HA

## Component Updates and Bug Fixes

* github.com/openziti/cobra-to-md: v1.0.1 (new)
* ztna-core/edge-api: [v0.26.25 -> v0.26.29](https://ztna-core/edge-api/compare/v0.26.25...v0.26.29)
* github.com/openziti/jwks: [v1.0.3 -> v1.0.4](https://github.com/openziti/jwks/compare/v1.0.3...v1.0.4)
* ztna-core/ztna: [v1.1.9 -> v1.1.10](https://ztna-core/ztna/compare/v1.1.9...v1.1.10)
    * [Issue #2374](https://ztna-core/ztna/issues/2374) - Edge Routers Do Not Accept JWTs for `openziti`/`native` audiences
    * [Issue #2353](https://ztna-core/ztna/issues/2353) - Update go-jose to avoid CVEs
    * [Issue #2333](https://ztna-core/ztna/issues/2333) - Give zit agent controller snapshot-db same capabilities as ziti fabric db snapshot
    * [Issue #2343](https://ztna-core/ztna/issues/2343) - Transferring leadership to another controller fails sometimes

# Release 1.1.9

## What's New

* Bug fixes, enhancements and continuing progress on controller HA
* Includes a performance update ([Issue #2340](https://ztna-core/ztna/issues/2340)) which should improve
  connection ramp times. Previously circuits would start with a relatively low window size and ramp slowly. Circuits
  will now start with a large initial window size and scale back if they can't keep up
* Added `ziti ops verify-network`. A command to aid when configuring the overlay network, this command will check config
  files for obvious mistakes
* Added `ziti ops verify-traffic`. Another command to aid when configuring the overlay network, this command will ensure
  the overlay network is able to pass traffic

## Component Updates and Bug Fixes

* github.com/openziti/agent: [v1.0.16 -> v1.0.17](https://github.com/openziti/agent/compare/v1.0.16...v1.0.17)
* github.com/openziti/channel/v2: [v2.0.136 -> v2.0.143](https://github.com/openziti/channel/compare/v2.0.136...v2.0.143)
    * [Issue #136](https://github.com/openziti/channel/issues/136) - Fix timeout on classic dialer 
    * [Issue #134](https://github.com/openziti/channel/issues/134) - Support the dtls transport

* ztna-core/edge-api: [v0.26.23 -> v0.26.25](https://ztna-core/edge-api/compare/v0.26.23...v0.26.25)
* github.com/openziti/foundation/v2: [v2.0.47 -> v2.0.48](https://github.com/openziti/foundation/compare/v2.0.47...v2.0.48)
* github.com/openziti/identity: [v1.0.81 -> v1.0.84](https://github.com/openziti/identity/compare/v1.0.81...v1.0.84)
* github.com/openziti/metrics: [v1.2.56 -> v1.2.57](https://github.com/openziti/metrics/compare/v1.2.56...v1.2.57)
* github.com/openziti/runzmd: [v1.0.49 -> v1.0.50](https://github.com/openziti/runzmd/compare/v1.0.49...v1.0.50)
* github.com/openziti/sdk-golang: [v0.23.39 -> v0.23.40](https://github.com/openziti/sdk-golang/compare/v0.23.39...v0.23.40)
    * [Issue #601](https://github.com/openziti/sdk-golang/issues/601) - Only send config types on service list if controller version supports it

* github.com/openziti/transport/v2: [v2.0.138 -> v2.0.143](https://github.com/openziti/transport/compare/v2.0.138...v2.0.143)
    * [Issue #85](https://github.com/openziti/transport/issues/85) - Update to latest dtls library

* ztna-core/ztna: [v1.1.8 -> v1.1.9](https://ztna-core/ztna/compare/v1.1.8...v1.1.9)
    * [Issue #2343](https://ztna-core/ztna/issues/2343) - Transferring leadership to another controller fails sometimes
    * [Issue #2340](https://ztna-core/ztna/issues/2340) - Update xgress defaults
    * [Issue #2336](https://ztna-core/ztna/issues/2336) - Re-enable optional xgress payload mtu, with message framing
    * [Issue #2091](https://ztna-core/ztna/issues/2091) - Add `scope` and `cliend_id` configuration to ext jwt signers
    * [Issue #2318](https://ztna-core/ztna/issues/2318) - Unable to update appData on existing edge routers using PATCH
    * [Issue #2281](https://ztna-core/ztna/issues/2281) - Session Certificates Should Return a Chain
    * [Issue #2285](https://ztna-core/ztna/issues/2285) - routers sometimes report link metrics for closed links 
    * [Issue #2282](https://ztna-core/ztna/issues/2282) - Investigate OIDC TOTP Redirect w/ application/json
    * [Issue #2279](https://ztna-core/ztna/issues/2279) - Ensure xweb initialized before RAFT
    * [Issue #2277](https://ztna-core/ztna/issues/2277) - docker controller and router deployments - generate a config by default
    * [Issue #2154](https://ztna-core/ztna/issues/2154) - HA MFA Enrollment returns 500
    * [Issue #2159](https://ztna-core/ztna/issues/2159) - API Session in HA return 400

# Release 1.1.8

## What's New

* Bug fixes, enhancements and continuing progress on controller HA

## Component Updates and Bug Fixes

* ztna-core/edge-api: [v0.26.20 -> v0.26.23](https://ztna-core/edge-api/compare/v0.26.20...v0.26.23)
    * [Issue #120](https://ztna-core/edge-api/issues/120) - Add API for retrieving services referencing a config
    * [Issue #121](https://ztna-core/edge-api/issues/121) - Add API for retrieving the set of attribute roles used by posture checks

* github.com/openziti/sdk-golang: [v0.23.38 -> v0.23.39](https://github.com/openziti/sdk-golang/compare/v0.23.38...v0.23.39)
    * [Issue #596](https://github.com/openziti/sdk-golang/issues/596) - SDK should submit selected config types to auth and service list APIs
    * [Issue #593](https://github.com/openziti/sdk-golang/issues/593) - SDK Golang OIDC Get API Session Returns Wrong Value

* github.com/openziti/storage: [v0.2.47 -> v0.3.0](https://github.com/openziti/storage/compare/v0.2.47...v0.3.0)
    * [Issue #80](https://github.com/openziti/storage/issues/80) - Set indexes aren't cleaned up when referenced entities are deleted, only when they change
    * [Issue #78](https://github.com/openziti/storage/issues/78) - Allow searching for things without case sensitivity

* ztna-core/ztna: [v1.1.7 -> v1.1.8](https://ztna-core/ztna/compare/v1.1.7...v1.1.8)
    * [Issue #2121](https://ztna-core/ztna/issues/2121) - Use router data model for edge router tunnel
    * [Issue #2245](https://ztna-core/ztna/issues/2245) - Add ability to retrieve a list of services that reference a config
    * [Issue #2089](https://ztna-core/ztna/issues/2089) - Enhance Management API to list Posture Check Roles
    * [Issue #2209](https://ztna-core/ztna/issues/2209) - `/edge/v1/external-jwt-signers` needs to be open
    * [Issue #2010](https://ztna-core/ztna/issues/2010) - Add config information to router data model
    * [Issue #1990](https://ztna-core/ztna/issues/1990) - Implement subscriber model for identity/service events in router
    * [Issue #2240](https://ztna-core/ztna/issues/2240) - Secondary ext-jwt Auth Policy check incorrectly requires primary ext-jwt auth to be enabled


# Release 1.1.7

## What's New

* Release actions fixes
* Fix for a flaky acceptance test

# Release 1.1.6

## What's New

* Trust Domain Configuration
* Controller HA Beta 2

## Trust Domain Configuration

OpenZiti controllers from this release forward will now require a `trust domain` to be configured. 
High Availability (HA) controllers already have this requirement. HA Controllers configure their trust domain via SPIFFE 
ids that are embedded in x509 certificates.

For feature parity, non-HA controllers will now have this same requirement. However, as re-issuing certificates is not
always easily done. To help with the transition, non-HA controllers will have the ability to have their trust domain
sourced from the controller configuration file through the root configuration value `trustDomain`. The configuration
field which takes a string that must be URI hostname compatible (see: https://github.com/spiffe/spiffe/blob/main/standards/SPIFFE-ID.md).
If this value is not defined, a trust domain will be generated from the root CA certificate of the controller. 

For networks that will be deployed after this change, it is highly suggested that a SPIFFE id is added to certificates.
The `ziti pki create ...` tooling supports the `--spiffe-id` option to help handle this scenario.

### Generated Trust Domain Log Messages

The following log messages are examples of warnings produced when a controller is using a generated trust domain:

```
WARNING this environment is using a default generated trust domain [spiffe://d561decf63d229d66b07de627dbbde9e93228925], 
  it is recommended that a trust domain is specified in configuration via URI SANs or the 'trustDomain' field

WARNING this environment is using a default generated trust domain [spiffe://d561decf63d229d66b07de627dbbde9e93228925], 
  it is recommended that if network components have enrolled that the generated trust domain be added to the 
  configuration field 'additionalTrustDomains'
```

### Trust domain resolution:

- Non-HA controllers
  - Prefers SPIFFE ids in x509 certificate URI SANs, looking at the leaf up the signing chain
  - Regresses to `trustDomain` in the controller configuration file if not found
  - Regress to generating a trust domain from the server certificates root CA, if the above do not resolve

- HA Controllers
  - Requires x509 SPIFFE ids in x509 certificate URI SANs

### Additional Trust Domains

When moving between trust domains (i.e. from the default generated to a new named one), the controller supports having
other trust domains. The trust domains do not replace certificate chain validation, which is still checked and enforced.

Additional trust domains are configured in the controller configuration file under the root field 
`additionalTrustDomains`. This field is an array of hostname safe strings.

The most common use case for this is field is if a network has issued certificates using the generated trust domain and 
now wants to transition to a explicitly defined one.

## Controller HA Beta 2

This release can be run in HA mode. The code is still beta, as we're still finding and fixing bugs. Several bugs
have been fixed since Beta 1 and c-based SDKs and tunnelers now work in HA mode. The smoketest can now be run
with HA controllers and clients.

* Latest ZET release supporting HA control: https://github.com/openziti/ziti-tunnel-sdk-c/releases/tag/v2.0.0-alpha9
* Windows, Mac and Mobile clients are in the process of being updated

For more information:

* HA overview/getting started/migration: [HA Documentation](https://ztna-core/ztna/tree/release-next/doc/ha)
* Open Issues: [HA Project Board](https://github.com/orgs/openziti/projects/9/views/1)

## Component Updates and Bug Fixes

* github.com/openziti/storage: [v0.2.45 -> v0.2.46](https://github.com/openziti/storage/compare/v0.2.45...v0.2.46)
    * [Issue #76](https://github.com/openziti/storage/issues/76) - Add support for non-boltz symbols to the the boltz stores

* ztna-core/ztna: [v1.1.5 -> v1.1.6](https://ztna-core/ztna/compare/v1.1.5...v1.1.6)
    * [Issue #2171](https://ztna-core/ztna/issues/2171) - Routers should consider control channels unresponsive if they are not connected
    * [Issue #2219](https://ztna-core/ztna/issues/2219) - Add inspection for router connections
    * [Issue #2195](https://ztna-core/ztna/issues/2195) - cached data model file set to
    * [Issue #2222](https://ztna-core/ztna/issues/2222) - Add way to get read-only status from cluster nodes
    * [Issue #2191](https://ztna-core/ztna/issues/2191) - Change raft list cluster members element name from values to data to match rest of REST api
    * [Issue #785](https://ztna-core/ztna/issues/785) - ziti edge update service-policy to empty/no posture checks fails
    * [Issue #2205](https://ztna-core/ztna/issues/2205) - Merge fabric and edge model code
    * [Issue #2165](https://ztna-core/ztna/issues/2165) - Add network id

# Release 1.1.5

## What's New

* Bug fixes

## Component Updates and Bug Fixes

* github.com/openziti/channel/v2: [v2.0.133 -> v2.0.136](https://github.com/openziti/channel/compare/v2.0.133...v2.0.136)
    * [Issue #132](https://github.com/openziti/channel/issues/132) - reconnecting dialer doesn't take local binding into account when reconnecting

* github.com/openziti/identity: [v1.0.80 -> v1.0.81](https://github.com/openziti/identity/compare/v1.0.80...v1.0.81)
* github.com/openziti/transport/v2: [v2.0.136 -> v2.0.138](https://github.com/openziti/transport/compare/v2.0.136...v2.0.138)
    * [Issue #83](https://github.com/openziti/transport/issues/83) - tls.Dial should use proxy configuration if provided

* github.com/openziti/xweb/v2: [v2.1.0 -> v2.1.1](https://github.com/openziti/xweb/compare/v2.1.0...v2.1.1)
* ztna-core/ztna: [v1.1.4 -> v1.1.5](https://ztna-core/ztna/compare/v1.1.4...v1.1.5)
    * [Issue #2173](https://ztna-core/ztna/issues/2173) - panic on HA peer connect
    * [Issue #2171](https://ztna-core/ztna/issues/2171) - Routers should consider control channels unresponsive if they are not connected
    * [Issue #2086](https://ztna-core/ztna/issues/2086) - Enable File Watching for Router/Controller Identities
    * [Issue #2087](https://ztna-core/ztna/issues/2087) - Ext JWT not setting provider value in auth query

# Release 1.1.4

## What's New

* Controller HA Beta 1
* Bug fixes

## Controller HA Beta 1

This release can be run in HA mode. The code is still beta, as we're still finding and fixing bugs. Several bugs 
have been fixed since Alpha 3 and c-based SDKs and tunnelers now work in HA mode. The smoketest can now be run
with HA controllers and clients.

* Initial ZET release support HA control: https://github.com/openziti/ziti-tunnel-sdk-c/releases/tag/v2.0.0-alpha1
* Windows, Mac and Mobile clients are in the process of being updated

For more information:

* HA overview/getting started/migration: [HA Documentation](https://ztna-core/ztna/tree/release-next/doc/ha)
* Open Issues: [HA Project Board](https://github.com/orgs/openziti/projects/9/views/1)

## Component Updates and Bug Fixes 

* github.com/openziti/channel/v2: [v2.0.130 -> v2.0.133](https://github.com/openziti/channel/compare/v2.0.130...v2.0.133)
* ztna-core/edge-api: [v0.26.19 -> v0.26.20](https://ztna-core/edge-api/compare/v0.26.19...v0.26.20)
    * [Issue #113](https://ztna-core/edge-api/issues/113) - RecoveryCodesEnvelope is wrong

* github.com/openziti/foundation/v2: [v2.0.45 -> v2.0.47](https://github.com/openziti/foundation/compare/v2.0.45...v2.0.47)
    * [Issue #407](https://github.com/openziti/foundation/issues/407) - Remove Branch from build info

* github.com/openziti/identity: [v1.0.77 -> v1.0.80](https://github.com/openziti/identity/compare/v1.0.77...v1.0.80)
* github.com/openziti/metrics: [v1.2.54 -> v1.2.56](https://github.com/openziti/metrics/compare/v1.2.54...v1.2.56)
* github.com/openziti/runzmd: [v1.0.47 -> v1.0.49](https://github.com/openziti/runzmd/compare/v1.0.47...v1.0.49)
* github.com/openziti/sdk-golang: [v0.23.37 -> v0.23.38](https://github.com/openziti/sdk-golang/compare/v0.23.37...v0.23.38)
    * [Issue #573](https://github.com/openziti/sdk-golang/issues/573) - api session refresh spins in a tight loop if there is no current api session
    * [Issue #562](https://github.com/openziti/sdk-golang/issues/562) - Support sticky dials

* github.com/openziti/secretstream: [v0.1.20 -> v0.1.21](https://github.com/openziti/secretstream/compare/v0.1.20...v0.1.21)
* github.com/openziti/storage: [v0.2.41 -> v0.2.45](https://github.com/openziti/storage/compare/v0.2.41...v0.2.45)
    * [Issue #73](https://github.com/openziti/storage/issues/73) - db integrity checker doesn't take nullable flag into account when checking unique indices
    * [Issue #71](https://github.com/openziti/storage/issues/71) - Add AddFkIndexCascadeDelete

* github.com/openziti/transport/v2: [v2.0.133 -> v2.0.136](https://github.com/openziti/transport/compare/v2.0.133...v2.0.136)
* ztna-core/ztna: [v1.1.3 -> v1.1.4](https://ztna-core/ztna/compare/v1.1.3...v1.1.4)
    * [Issue #2084](https://ztna-core/ztna/issues/2084) - Bug: Router enrollment is missing its server chain
    * [Issue #2124](https://ztna-core/ztna/issues/2124) - api session certs should be deleted when related api sessions are deleted

# Release 1.1.3

## What's New

* Sticky Terminator Selection
* Linux and Docker deployments log formats no longer default to the simplified format option and now use logging library
  defaults: `json` for non-interactive, `text` for interactive.

NOTE: This release is the first since 1.0.0 to be marked promoted from pre-release. Be sure to check the release notes
      for the rest of the post-1.0.0 releases to get the full set of changes.

## Stick Terminator Strategy

This release introduces a new terminator selection strategy `sticky`. On every dial it will return a token to the 
dialer, which represents the terminator used in the dial. This token maybe passed in on subsequent dials. If no token
is passed in, the strategy will work the same as the `smartrouting` strategy. If a token is passed in, and the 
terminator is still valid, the same terminator will be used for the dial. A terminator will be consideder valid if
it still exists and there are no terminators with a higher precedence. 

This is currently only supported in the Go SDK.

### Go SDK Example

```
ziti edge create service test --terminator-strategy sticky
```

```
	conn := clientContext.Dial("test")
	token := conn.Conn.GetStickinessToken()
	_ = conn.Close()

	dialOptions := &ziti.DialOptions{
		ConnectTimeout:  time.Second,
		StickinessToken: token,
	}
	conn = clientContext.DialWithOptions("test", dialOptions))
	nextToken := conn.Conn.GetStickinessToken()
	_ = conn.Close()
```

## Component Updates and Bug Fixes

* github.com/openziti/channel/v2: [v2.0.128 -> v2.0.130](https://github.com/openziti/channel/compare/v2.0.128...v2.0.130)
* ztna-core/edge-api: [v0.26.18 -> v0.26.19](https://ztna-core/edge-api/compare/v0.26.18...v0.26.19)
* github.com/openziti/foundation/v2: [v2.0.42 -> v2.0.45](https://github.com/openziti/foundation/compare/v2.0.42...v2.0.45)
* github.com/openziti/identity: [v1.0.75 -> v1.0.77](https://github.com/openziti/identity/compare/v1.0.75...v1.0.77)
* github.com/openziti/metrics: [v1.2.51 -> v1.2.54](https://github.com/openziti/metrics/compare/v1.2.51...v1.2.54)
* github.com/openziti/runzmd: [v1.0.43 -> v1.0.47](https://github.com/openziti/runzmd/compare/v1.0.43...v1.0.47)
* github.com/openziti/sdk-golang: [v0.23.35 -> v0.23.37](https://github.com/openziti/sdk-golang/compare/v0.23.35...v0.23.37)
    * [Issue #562](https://github.com/openziti/sdk-golang/issues/562) - Support sticky dials

* github.com/openziti/secretstream: [v0.1.19 -> v0.1.20](https://github.com/openziti/secretstream/compare/v0.1.19...v0.1.20)
* github.com/openziti/storage: [v0.2.37 -> v0.2.41](https://github.com/openziti/storage/compare/v0.2.37...v0.2.41)
* github.com/openziti/transport/v2: [v2.0.131 -> v2.0.133](https://github.com/openziti/transport/compare/v2.0.131...v2.0.133)
* ztna-core/ztna: [v1.1.2 -> v1.1.3](https://ztna-core/ztna/compare/v1.1.2...v1.1.3)
    * [Issue #2064](https://ztna-core/ztna/issues/2064) - Fix panic on link close
    * [Issue #2062](https://ztna-core/ztna/issues/2062) - Link connection retry delays should contain some randomization 
    * [Issue #2055](https://ztna-core/ztna/issues/2055) - Controller panics on 'ziti agent cluster list'
    * [Issue #2019](https://ztna-core/ztna/issues/2019) - Support mechanism for sticky dials

# Release 1.1.2

## What's New

* Bug fixes and minor enhancements

## Component Updates and Bug Fixes
* github.com/openziti/sdk-golang: [v0.23.32 -> v0.23.35](https://github.com/openziti/sdk-golang/compare/v0.23.32...v0.23.35)
* ztna-core/ztna: [v1.1.1 -> v1.1.2](https://ztna-core/ztna/compare/v1.1.1...v1.1.2)
  * [Issue #2032](https://ztna-core/ztna/issues/2032) - Auto CA Enrollment Fails w/ 400 Bad Request
  * [Issue #2026](https://ztna-core/ztna/issues/2026) - Root Version Endpoint Handling 404s
  * [Issue #2002](https://ztna-core/ztna/issues/2002) - JWKS endpoints may not refresh on new KID
  * [Issue #2007](https://ztna-core/ztna/issues/2007) - Identities for edge routers with tunneling enabled sometimes show hasEdgeRouterConnection=false even though everything is OK
  * [Issue #1983](https://ztna-core/ztna/issues/1983) - delete of non-existent entity causes panic when run on follower controller


# Release 1.1.1

## What's New

* HA Alpha-3
* Bug fixes and minor enhancements
* [The all-in-one quickstart compose project](./quickstart/docker/all-in-one/README.md) now uses the same environment variable to configure the controller's address as the ziti command line tool

## HA Alpha 3

This release can be run in HA mode. The code is still alpha, as we're still finding and fixing bugs. 

For more information:

* HA overview/getting started/migration: [HA Documementation](https://ztna-core/ztna/tree/release-next/doc/ha)
* Open Issues: [HA Project Board](https://github.com/orgs/openziti/projects/9/views/1) 

## New Contributors

Thanks to new contributors

* @Vrashabh-Sontakke

## Component Updates and Bug Fixes
* ztna-core/edge-api: [v0.26.17 -> v0.26.18](https://ztna-core/edge-api/compare/v0.26.17...v0.26.18)
* github.com/openziti/sdk-golang: [v0.23.27 -> v0.23.32](https://github.com/openziti/sdk-golang/compare/v0.23.27...v0.23.32)
    * [Issue #554](https://github.com/openziti/sdk-golang/issues/554) - Passing in config types on service list breaks on older controller

* github.com/openziti/storage: [v0.2.36 -> v0.2.37](https://github.com/openziti/storage/compare/v0.2.36...v0.2.37)
    * [Issue #64](https://github.com/openziti/storage/issues/64) - Add support for transaction complete listeners

* ztna-core/ztna: [v1.1.0 -> v1.1.1](https://ztna-core/ztna/compare/v1.1.0...v1.1.1)
    * [Issue #1973](https://ztna-core/ztna/issues/1973) - Raft should not initialize if db is misconfigured
    * [Issue #1971](https://ztna-core/ztna/issues/1971) - BUG: OIDC authentication does not convert config type names to ids
    * [Issue #1966](https://ztna-core/ztna/issues/1966) - Handle multi-entity updates in router data model
    * [Issue #1772](https://ztna-core/ztna/issues/1772) - provide a better error when the user is not logged in
    * [Issue #1964](https://ztna-core/ztna/issues/1964) - Add API Session Token Update Messaging
    * [Issue #1960](https://ztna-core/ztna/issues/1960) - JWT Session exchange isn't working
    * [Issue #1962](https://ztna-core/ztna/issues/1962) - permissions enum doesn't contain "Invalid"

# Release 1.1.0

## What's New

* HA Alpha2
* Deployments Alpha
    * Linux packages provide systemd services for controller and router. Both depend on existing package `openziti` which provides the `ziti` command line tool.
        * `openziti-controller` provides `ziti-controller.service`
        * `openziti-router` provides `ziti-router.service`
    * Container images for controller and router now share the bootstrapping logic with the packages, so they
      support the same configuration options.

## HA Alpha2

This release can be run in HA mode. The code is still alpha, so there are still some bugs and missing features,
however basic functionality work with the exceptions noted. See the [HA Documementation](https://ztna-core/ztna/tree/release-next/doc/ha)
for instructions on setting up an HA cluster.

### Known Issues

* JWT Session exchange isn't working with Go SDK clients
    * This means Go clients will need to be restarted once their sessions expire
* Service/service policy changes might not be reflected in routers
    * Changes to policy may not yet properly sync to the routers, causing unexpected behavior with ER/Ts running in HA mode

More information can be found on the [HA Project Board](https://github.com/orgs/openziti/projects/9/views/1)

## Component Updates and Bug Fixes

* ztna-core/edge-api: [v0.26.16 -> v0.26.17](https://ztna-core/edge-api/compare/v0.26.16...v0.26.17)
    * [Issue #107](https://ztna-core/edge-api/issues/107) - Add configTypes param to service list

* github.com/openziti/sdk-golang: [v0.23.19 -> v0.23.27](https://github.com/openziti/sdk-golang/compare/v0.23.19...v0.23.27)
    * [Issue #545](https://github.com/openziti/sdk-golang/issues/545) - Set config types on query when listing services
    * [Issue #541](https://github.com/openziti/sdk-golang/issues/541) - Token exchange in Go SDK not working
    * [Issue #540](https://github.com/openziti/sdk-golang/issues/540) - Switch to EdgeRouter.SupportedProtocols from deprecated URLs map

* ztna-core/ztna: [v1.0.0 -> v1.1.0](https://ztna-core/ztna/compare/v1.0.0...v1.1.0)
    * [Issue #1952](https://ztna-core/ztna/issues/1952) - Remove support for fabric only identities in CLI
    * [Issue #1950](https://ztna-core/ztna/issues/1950) - Add policy type to service policy router events
    * [Issue #1951](https://ztna-core/ztna/issues/1951) - Add more attributes to route data model Identity
    * [Issue #1942](https://ztna-core/ztna/issues/1942) - Rework ER/T intercept code to be sessionless or use JWT sessions
    * [Issue #1936](https://ztna-core/ztna/issues/1936) - SDK Hosted HA sessions are getting removed when they shouldn't be
    * [Issue #1934](https://ztna-core/ztna/issues/1934) - Don't publish binary builds to artifactory
    * [Issue #1931](https://ztna-core/ztna/issues/1931) - "invalid kid: <kid>" randomly occurs in HA mode

# Release 1.0.0

## About 1.0

What does marking OpenZiti as 1.0 mean?

### Backwards Compatibility
We've guaranteed API stability for SDK clients for years and worked hard to ensure that routers 
and controllers would be backwards and forward compatible. However, we have had a variety of 
management API changes and CLI changes. For post 1.0 releases we expect to make additions to the 
APIs and CLI, but won't remove anything until it's been first marked as deprecated and then only
with a major version bump. 

### Stability and Scale
Recent releases have seen additional testing using chaos testing techniques. These tests involve
setting up relatively large scale environments, knocking out various components and then verifying
that the network is able to return to a stable state. These test are run for hours to try and 
eliminate race conditions and distributed state machine problems. 

OpenZiti is also being used as underlying infrastrcture for the zrok public service. Use of this 
network has grown quickly and proven that it's possible to build ziti native apps that can scale
up.

## Backward Incompatible Changes to pre-1.0 releases

Administrators no longer have access to dial/bind all services by default. See below for details.

## What's New

* Administrators no longer have access to dial/bind all services by default.
* TLS Handshakes can now be rate limited in the controller
* TLS Handshake timeouts can now be set on the controller when using ALPN
* Bugfixes

## DEFAULT Bind/Dial SERVICE PERMISSIONS FOR Admin IDENTITIES HAVE CHANGED

Admin identities were able to Dial and Bind all services regardless of the effective service policies
prior to this release. This could lead to a confusing situation where a tunneler that was assuming an Admin
identity would put itself into an infinite connect-loop when a service's host.v1 address overlapped with
any addresses in its intercept configuration.

Please create service policies to grant Bind or Dial permissions to Admin identities as needed.

## TLS Handshake

A TLS handhshake rate limiter can be enabled. This is useful in cases where there's a flood of TLS requests and the
controller can't handle them all. It can get into a state where it can't respond to TLS handshakes quickly enough,
so the clients time out. They then retry, adding to the the load. The controller ends up wasting time doing work 
that isn't use. 

This uses the same rate limiting as the auth rate limiter. 

Additionally the server side handshake timeout can now be configured.

Configuration:

```
tls: 
  handshakeTimeout: 15s

  rateLimiter:
    # if disabled, no tls handshake rate limiting with be enforced
    enabled: true
    # the smallest window size for tls handshakes
    minSize: 5
    # the largest allowed window size for tls handshakes
    maxSize: 5000
    # after how long to consider a handshake abandoned if neither success nor failure was reported
    timeout: 30s
```

New metrics:

* `tls_handshake_limiter.in_process` - number of TLS handshakes in progress
* `tls_handshake_limiter.window_size` - number of TLS handhshakes allowed concurrently
* `tls_handshake_limiter.work_timer` - timer tracking how long TLS handshakes are taking


## Component Updates and Bug Fixes

* github.com/openziti/channel/v2: [v2.0.122 -> v2.0.128](https://github.com/openziti/channel/compare/v2.0.122...v2.0.128)
* ztna-core/edge-api: [v0.26.14 -> v0.26.16](https://ztna-core/edge-api/compare/v0.26.14...v0.26.16)
* github.com/openziti/foundation/v2: [v2.0.40 -> v2.0.42](https://github.com/openziti/foundation/compare/v2.0.40...v2.0.42)
* github.com/openziti/identity: [v1.0.73 -> v1.0.75](https://github.com/openziti/identity/compare/v1.0.73...v1.0.75)
* github.com/openziti/metrics: [v1.2.48 -> v1.2.51](https://github.com/openziti/metrics/compare/v1.2.48...v1.2.51)
* github.com/openziti/runzmd: [v1.0.41 -> v1.0.43](https://github.com/openziti/runzmd/compare/v1.0.41...v1.0.43)
* github.com/openziti/sdk-golang: [v0.23.15 -> v0.23.19](https://github.com/openziti/sdk-golang/compare/v0.23.15...v0.23.19)
* github.com/openziti/secretstream: [v0.1.18 -> v0.1.19](https://github.com/openziti/secretstream/compare/v0.1.18...v0.1.19)
* github.com/openziti/storage: [v0.2.33 -> v0.2.36](https://github.com/openziti/storage/compare/v0.2.33...v0.2.36)
* github.com/openziti/transport/v2: [v2.0.125 -> v2.0.131](https://github.com/openziti/transport/compare/v2.0.125...v2.0.131)
    * [Issue #79](https://github.com/openziti/transport/issues/79) - Add adaptive rate limiting to shared tls listener

* ztna-core/ztna: [v0.34.2 -> v1.0.0](https://ztna-core/ztna/compare/v0.34.2...v1.0.0)
    * [Issue #1923](https://ztna-core/ztna/issues/1923) - Add release validation test suite
    * [Issue #1904](https://ztna-core/ztna/issues/1904) - Add TLS handshake rate limiter
    * [Issue #1921](https://ztna-core/ztna/issues/1921) - Tidy CLI
    * [Issue #1916](https://ztna-core/ztna/issues/1916) - SDK dials fails with 'token is malformed' error
    * [Issue #1911](https://ztna-core/ztna/issues/1911) - Fix panic on first HA controller startup
    * [Issue #1914](https://ztna-core/ztna/issues/1914) - Fix panic in PeerConnected
    * [Issue #1781](https://ztna-core/ztna/issues/1781) - Admin identities have bind and dial permissions to services
