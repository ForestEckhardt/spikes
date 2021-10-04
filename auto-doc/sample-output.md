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

## Build Plan
```
[[provides]]
name = "lorem"

[[requires]]
name = "ipsum"

[requires.metadata]
build = true
```

## Caching Reuse Logic
| `lorem` | `ipsum` | `dolor` | Command |
| ------- | ------- | ------- | ------- |
| X | X | X | `sum` |
| X | X | ✓ | `sum` |
| X | ✓ | X | `es` |
| X | ✓ | ✓ | `es` |
| ✓ | X | X | `est` |
| ✓ | X | ✓ | `est` |
| ✓ | ✓ | X | `es` |
| ✓ | ✓ | ✓ | `est` |

