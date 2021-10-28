# Paketo BellSoft Liberica Buildpack

## Environment Variable Configuration
### BPL_JVM_HEAD_ROOM
the headroom in memory calculation
Default Value: `0`
This environment variable is used during launch

### BPL_JVM_LOADED_CLASS_COUNT
the number of loaded classes in memory calculation
Default Value: `35% of classes`
This environment variable is used during launch

### BPL_JVM_THREAD_COUNT
the number of threads in memory calculation
Default Value: `250`
This environment variable is used during launch

### BPL_HEAP_DUMP_PATH
write heap dumps on error to this path
This environment variable is used during launch

### BPL_JAVA_NMT_ENABLED
enables Java Native Memory Tracking (NMT)
Default Value: `true`
This environment variable is used during launch

### BPL_JAVA_NMT_LEVEL
configure level of NMT, summary or detail
Default Value: `summary`
This environment variable is used during launch

### BPL_JMX_ENABLED
enables Java Management Extensions (JMX)
Default Value: `false`
This environment variable is used during launch

### BPL_JMX_PORT
configure the JMX port
Default Value: `5000`
This environment variable is used during launch

### BPL_DEBUG_ENABLED
enables Java remote debugging support
Default Value: `false`
This environment variable is used during launch

### BPL_DEBUG_PORT
configure the remote debugging port
Default Value: `8000`
This environment variable is used during launch

### BPL_DEBUG_SUSPEND
configure whether to suspend execution until a debugger has attached
Default Value: `false`
This environment variable is used during launch

### BP_JVM_VERSION
the Java version
Default Value: `11`
This environment variable is used during build

### BP_JVM_TYPE
the JVM type - JDK or JRE
Default Value: `JRE`
This environment variable is used during build

### JAVA_TOOL_OPTIONS
the JVM launch flags
This environment variable is used during launch

## Behavior

This buildpack will participate if any of the following conditions are met

* Another buildpack requires `jdk`
* Another buildpack requires `jre`

The buildpack will do the following if a JDK is requested:

* Contributes a JDK to a layer marked `build` and `cache` with all commands on `$PATH`
* Contributes `$JAVA_HOME` configured to the build layer
* Contributes `$JDK_HOME` configure to the build layer

The buildpack will do the following if a JRE is requested:

* Contributes a JRE to a layer with all commands on `$PATH`
* Contributes `$JAVA_HOME` configured to the layer
* Contributes `-XX:ActiveProcessorCount` to the layer
* Contributes `-XX:+ExitOnOutOfMemoryError` to the layer
* Contributes `-XX:+UnlockDiagnosticVMOptions`,`-XX:NativeMemoryTracking=summary` & `-XX:+PrintNMTStatistics` to the layer (Java NMT)
* If `BPL_JMX_ENABLED = true`
  * Contributes `-Djava.rmi.server.hostname=127.0.0.1`, `-Dcom.sun.management.jmxremote.authenticate=false`, `-Dcom.sun.management.jmxremote.ssl=false` & `-Dcom.sun.management.jmxremote.rmi.port=5000`
* If `BPL_DEBUG_ENABLED = true`
  * Contributes `-agentlib:jdwp=transport=dt_socket,server=y,address=*:8000,suspend=n`. If Java version is 8, address parameter is `address=:8000`
* Contributes `$MALLOC_ARENA_MAX` to the layer
* Disables JVM DNS caching if link-local DNS is available
* If `metadata.build = true`
  * Marks layer as `build` and `cache`
* If `metadata.launch = true`
  * Marks layer as `launch`
* Contributes Memory Calculator to a layer marked `launch`
* Contributes Heap Dump helper to a layer marked `launch`

## Bindings

The buildpack optionally accepts the following bindings:

### Type: `dependency-mapping`

| Key                   | Value   | Description                                                                                       |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------- |
| `<dependency-digest>` | `<uri>` | If needed, the buildpack will fetch the dependency with digest `<dependency-digest>` from `<uri>` |

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0


