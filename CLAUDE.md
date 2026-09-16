# CLAUDE.md — plc4x-yofc

YOFC 维护的 Apache PLC4X fork（默认 `heyoulin` 分支，~1118+ commit ahead of `apache/develop`）。下游 `yofc-iot` 通过 `~/.m2` 消费 `plc4j-*` artifact。

## Build

- 默认 root modules：`code-generation`、`protocols`、`website`
- `plc4j`（含 51 个子模块）只在 `-Pwith-java` profile 下进 reactor
- 不构建：`plc4go` / `plc4net` / `plc4c` / `plc4py`，以及通常没必要的 `website`

**常用入口：**

```bash
# 完整 build（推荐）
mvn install -Pwith-java -pl '!website' -DskipTests -T 1C

# 完整 build 含测试（合并上游后用）
RUN_TESTS=1 mvn install -Pwith-java -pl '!website' -DskipITs=true -T 1C

# 跑单元测试
mvn test -Pwith-java -pl '!website' -DskipITs=true -T 1C

# 仅 plc4j 子树（只在已经跑过完整 build、~/.m2 fresh 的前提下用）
cd plc4j && mvn install -DskipTests -T 1C

# 单独验证某个模块的 JaCoCo 覆盖率
mvn verify -f plc4j/drivers/<driver>/pom.xml
```

## SPI 结构（合并后）

上游把单体 `plc4j-spi` 拆分成多个子模块：

| 子模块 | 职责 |
|--------|------|
| `plc4j-spi-values` | PlcValue 类型（PlcDINT、PlcBOOL…）+ DefaultPlcValueHandler |
| `plc4j-spi-buffers-api` | ReadBuffer / WriteBuffer 抽象 API |
| `plc4j-spi-buffers-byte` | 基于 byte 数组的 ReadBufferByteBased / WriteBufferByteBased |
| `plc4j-spi-config` | 配置注解与解析 |
| `plc4j-spi-drivers` | DriverBase / ConnectionBase / MessageCodecBase / DefaultPlcWriteRequest 等 |
| `plc4j-spi-fields` | FieldReader / FieldWriter / DataReader / DataWriter 工厂 |
| `plc4j-spi-utils` | StaticHelper（STR_LEN、COUNT 等）、EvaluationHelper |

**影响**：所有 driver/tool 的 pom 中 `plc4j-spi` 依赖需替换为所需子模块的显式依赖。原 `plc4j-spi` artifact 现在是 pom-only 聚合。

## Jackson 3 迁移

**方向：YOFC 从 Jackson 2 迁到了 Jackson 3（`tools.jackson`, `3.2.0`），上游停在 Jackson 2。** 这是 fork 的永久性偏离，不是「跟随上游」。合并时上游的 Jackson 2 代码要迁到 3，反向不成立。详见文末「上游同步 → Jackson 3 方向」。

### 迁移检查清单

| Jackson 2 | Jackson 3 |
|-----------|-----------|
| `com.fasterxml.jackson.core.*` | `tools.jackson.core.*` |
| `com.fasterxml.jackson.databind.*` | `tools.jackson.databind.*` |
| `com.fasterxml.jackson.dataformat.*` | `tools.jackson.dataformat.*` |
| `JsonProcessingException` | `JacksonException` |
| `new ObjectMapper()` | `JsonMapper.builder().build()` |
| `objectMapper.configure(Feature, bool)` | `JsonMapper.builder().disable(Feature).build()` |
| `com.fasterxml.jackson.core:jackson-*` (groupId) | `tools.jackson.core:jackson-*` |

### 常见遗漏

合并后如果看到 `package com.fasterxml.jackson.databind.* does not exist` / `... .core.*`，说明上游在这轮改了该文件，把它的 Jackson 2 import 带了进来，需要手动迁到 `tools.jackson.*`。

先跑一次全树扫描定位，比逐个看编译错误快：

```bash
grep -rn 'import com\.fasterxml\.jackson\.' --include='*.java' plc4j/ \
  | grep -v 'com\.fasterxml\.jackson\.annotation\.'      # 有输出才需要迁移
```

`com.fasterxml.jackson.annotation.*` 是 Jackson 3 的正常用法，**不要**动。2026-09 那次合并该扫描为空。

## JaCoCo 覆盖率

父 pom（`plc4j/pom.xml`）定义了全局规则：
- **INSTRUCTION COVEREDRATIO ≥ 0.80**（80% 指令覆盖率）
- **CLASS MISSEDCOUNT = 0**（不允许整个类零覆盖）
- 父级默认排除：`**/readwrite/*.class`、`**/*$*Listener.class`、`**/with*.class`

### YOFC 自定义代码覆盖率不足时的处理

当 YOFC 新增的 I/O 密集、协议编解码、服务器生命周期类拉低覆盖率：

```xml
<plugin>
  <groupId>org.jacoco</groupId>
  <artifactId>jacoco-maven-plugin</artifactId>
  <executions>
    <execution>
      <id>report</id>
      <configuration>
        <excludes combine.children="append">
          <exclude>com/example/HeavyIoClass.class</exclude>
        </excludes>
      </configuration>
    </execution>
    <execution>
      <id>check-coverage</id>
      <configuration>
        <excludes combine.children="append">
          <exclude>com/example/HeavyIoClass.class</exclude>
        </excludes>
      </configuration>
    </execution>
  </executions>
</plugin>
```

**关键**：`report` 和 `check-coverage` 都要加 `combine.children="append"`。只加 `check-coverage` 不加 `report` 的话，report 仍然包含被排除的类，显示的覆盖率与 check 规则不一致。

**验证**：用 `mvn verify -f <module>/pom.xml` 而不是 `mvn jacoco:check`（后者用 `default-cli` 执行 ID，不继承父 pom 的 rules）。

### 已有 JaCoCo 排除

| 模块 | 排除的类 | 原因 |
|------|----------|------|
| `plc4j-driver-ads` | ValueEncoder, ValueDecoder | 二进制协议编解码，需要真实 ADS 设备做集成测试 |
| `plc4j-driver-simulated` | SimulatedDevice, SimulatedDriver | YOFC H2 MVStore 持久化 + URL 参数解析 |
| `plc4j-plc4x-server` | Plc4xServer*, Plc4xServerAdapter | 服务器生命周期 + 协议分发逻辑 |

## DefaultArrayInfo 约定

```java
new DefaultArrayInfo(lowerBound, upperBound)
// getSize() = upperBound - lowerBound + 1
```

**N 个元素的数组**：`new DefaultArrayInfo(0, N - 1)`，`getSize()` 返回 `N`。

上游 `DefaultPlcValueHandler.ofElements()` 严格检查 `values.length != arrayInfo.getSize()`，off-by-one 会导致 `"Expecting N items, but got M"` 错误。合并后留意所有 driver 的 `getArrayInfo()` 实现是否正确。

## PlcValueType.NULL 处理

`DefaultPlcValueHandler.ofElement()` 中 `PlcValueType.NULL` 必须与 `null` 同等处理（走类型推断路径），否则代理驱动（plc4x proxy）等 tag 类型为 NULL 的场景会丢失实际值。

```java
// 正确 — NULL 也走推断
if (type == null || type == PlcValueType.NULL) {
    return of(value);  // 从 Java 类型推断：Integer→PlcDINT, Short→PlcINT, ...
}
```

## YOFC 特有功能

### stringEncoding（ADS / S7）

ADS 和 S7 tag 支持 `|encoding` 后缀指定字符串编码（如 `Main.value|UTF-16`）。这是 YOFC 独有特性，上游没有。合并上游时注意保留 tag 解析中的 encoding 分支。

### SimulatedDevice FILE 模式

YOFC 在 simulated driver 中新增了 `FILE` tag 类型，支持 H2 MVStore 持久化（`simulated://device?file=/path&data=mapName`）。包含 `SimulatedConfiguration` 的 file/data 字段和 `SimulatedDevice` 的 `getMvValue()`/`writeMvValue()` 方法。

### plc4x-server

YOFC 自建的 PLC4X 代理服务器（`plc4j/tools/plc4x-server`），基于 Netty 实现 plc4x 协议的服务端。包含 `Plc4xServerAdapter`（协议分发）、`Plc4xServerCodec`（帧编解码）、`Plc4xServer`（Netty bootstrap）。

## 合并后修复流程

`git merge origin/develop` 后的典型修复顺序：

1. **编译错误** — Jackson 2→3 import 遗漏、SPI 子模块依赖缺失、transport artifact 重命名
2. **`mvn install -DskipTests`** — 确保编译通过、artifact 入 `~/.m2`
3. **`mvn test`** — 逐个修复测试失败
   - DefaultArrayInfo off-by-one（`getArrayInfo()` 实现不匹配新的严格检查）
   - PlcValueType.NULL 行为变更（代理驱动写请求 value 被丢弃）
   - 依赖收敛（enforcer `DependencyConvergence`，如 `stax2-api` 版本冲突 → root pom pin 版本）
4. **JaCoCo 覆盖率** — YOFC 自定义类拉低覆盖率 → 加排除
5. **Flaky 测试** — Testcontainers/OrbStack 环境问题 → `@Disabled` 或环境配置

### Testcontainers 注意事项

- macOS + OrbStack 下 Ryuk 容器 `Connection refused` 是已知问题，不是代码 bug
- 解决：`TESTCONTAINERS_RYUK_DISABLED=true` 环境变量或 `~/.testcontainers.properties` 中设置 `ryuk.container.privileged=true`
- OPC UA MiloTestContainer 的固定端口 12686 可能导致 flaky test（TCP endpoint 未就绪）

## mspec / generated 流

- `.mspec` 在 `protocols/*/src/main/generated/protocols/...`；YOFC 保持上游做法**把生成的 Java 也 commit 到** `plc4j/drivers/*/src/main/generated/`
- 生成器在每个 driver pom 自己跑：`plc4x-maven-plugin` → `generate-driver` goal → `<outputDir>src/main/generated</outputDir>`，phase = `generate-sources`
- 这个 plugin 依赖 `plc4x-protocols-<name>`（scope `provided`），**从 `~/.m2` 读 mspec**
- 推论：**改了 mspec 必须走完整 root build**，不能只 `cd plc4j`——否则 plugin 用旧 mspec artifact 把磁盘上的新 generated Java 覆盖成旧版
- 验证自洽：完整 build 后 `git status` 应该 clean。dirty 说明 mspec ↔ generated 漂移，需要 commit regen 出来的变更

## 上游同步

- 远端：`origin` = `apache/plc4x`，`heyoulin` = `spnettec/plc4x`（YOFC fork）
- 工作流：`git fetch origin && git merge origin/develop`（**不要 rebase**——历史里全是 merge commit）
- 常见冲突类型：`protocols/bacnetip/.../bacnet-vendorids.mspec` 这类注册表型表追加，upstream 追加新 vendor，YOFC 那侧空白 → take-theirs。同时取上游对应的 Java/Go 生成产物
- merge 完跑一次完整 root build + `git status` clean 验证 → push：`git push heyoulin heyoulin`

### 版本号约定

YOFC 保持自己的 `1.0.0.B-SNAPSHOT`，**不跟上游的版本号**（上游 2026-09 起为 `1.1.0-SNAPSHOT`）。`yofc-iot` 的 `com.yofc.iot.apps.opcuaserver` 等 pom 把 `plc4j-*` 依赖钉在 `1.0.0.B-SNAPSHOT`，改这个版本号会连带改下游。

冲突表现：~100 个 `pom.xml` 的 `<version>` 行冲突。判据是「只有版本/时间戳差异」→ 取 ours；「上游单方面新增的行」→ 取 theirs 但把 `1.1.0-SNAPSHOT` 换回 `1.0.0.B-SNAPSHOT`。merge 后务必 grep 一遍 `1\.1\.0-SNAPSHOT` 和 `1\.0\.0-SNAPSHOT`（没有 `.B`）：上游新增的 `<parent>` version、新依赖的 version 都不会冲突，会静默带进来，直到构建报 "Non-resolvable parent POM" 才暴露。

### 构建：**不要用 `-T 1C`**

`apache-rat-plugin:0.18` 在并行模块构建下会抛 `ConcurrentModificationException`，或报假的 `Unexpected count for UNAPPROVED`（不同模块、不同次数，非确定性）。用串行：

```bash
mvn install -Pwith-java -pl '!website' -DskipTests          # 95 个模块，串行约 2-3 分钟
mvn test   -Pwith-java -pl '!website' -DskipITs=true --fail-at-end   # 11564 个测试
```

跑测试务必加 `--fail-at-end`，否则 Maven 在第一个失败模块就中止，后面模块一个都没跑，会看不到其余失败。

### Jackson 3 方向（2026-09 复核）

上游在 **Jackson 2**（`com.fasterxml.jackson`, 2.22.x）；YOFC 在 **Jackson 3**（`tools.jackson`, 3.2.0）。方向是 YOFC 单方面前进，不是「跟随上游迁移」——上游做过 3.x 试验后回退了。合并时不要接受上游的 Jackson 2 groupId/版本。

- `com.fasterxml.jackson.annotation.*` 是**对的**（Jackson 3 仍用这个包名放注解 + `jackson-annotations` 这个 artifact），不要"迁移"掉
- `jackson-datatype-jsr310` 在 Jackson 3 下**没有 3.2.0 版本**（central 只到 3.0.0-rc2）：java.time 支持已并入 `jackson-databind`（`tools/jackson/databind/ext/javatime/`）。上游加的 `com.fasterxml.jackson.datatype:jackson-datatype-jsr310` 依赖要**删掉**而不是改 groupId

### 非 Java 语言绑定：**一律取上游**（政策，2026-09 明确）

YOFC 的定制只在 **Java** 上。`plc4go` / `plc4c` / `plc4py` / `plc4net` 里 YOFC 那侧的改动当初只是为了**让它们能编译过**（跟随 mspec 变更的连带修改），不是功能需求。所以：

- 这些目录的冲突**一律 take-theirs**，不要像 Java 那样做 union
- 也不要保留 YOFC 在这些目录里的历史定制（例：`plc4go` 的 `|stringEncoding` 透传）。它们依赖协议自动生成的代码，而生成物是按上游 mspec 出的，留着就对不上、编译不过

**为什么不能靠构建发现**：`-Pwith-java` 只把 `plc4j` 拉进 reactor，这些目录根本不被编译。合并后要单独验：

```bash
GO=~/.mvnGoLang/go1.26.0.darwin-arm64/bin/go
cd plc4go && $GO build ./... && $GO vet ./...     # exit 0 才算过
```

（`go vet` 会报一批 `ReadByte`/`WriteByte` 签名不匹配的 style 警告，那是上游自带的，不是错误。）

例外：这几个目录的 `pom.xml` **版本号仍要保持 `1.0.0.B-SNAPSHOT`**，否则父 pom 解析不到。2026-09 合并后，`plc4go`/`plc4c`/`plc4py`/`plc4net` 与上游的差异只剩这些 pom 版本行，外加一个没人引用的 `plc4go/spi/values/PlcBIT.go`。

## 工具链

- Java 21（Temurin）+ Maven 4（`~/apps/maven/apache-maven-4.0.0-rc-6`——root pom 的 enforcer 要求 `[4.0.0-rc-6,)`；项目已迁移 POM 4.1.0 schema + `<subprojects>`，Maven 3 无法构建）
- 默认 `java` 是 26；构建要显式 `JAVA_HOME=/Library/Java/JavaVirtualMachines/temurin-21.jdk/Contents/Home`

# Codebase Memory MCP — Code Intelligence

This project is indexed by **codebase-memory-mcp**. Always use it BEFORE grep/find or reading files when you need to understand or locate code. The skill at `~/.claude/skills/codebase-memory/` contains the full decision matrix and workflow.

## Quick Reference

| Question | Tool |
|----------|------|
| Who calls X? | `trace_path(direction="inbound")` |
| What does X call? | `trace_path(direction="outbound")` |
| Find by name | `search_graph(name_pattern="...")` |
| Dead code | `search_graph(max_degree=0)` |
| Impact of changes | `detect_changes()` |
| Architecture overview | `get_architecture(aspects=["all"])` |
| Read source | `get_code_snippet(qualified_name="...")` |

## Exploration Workflow

`list_projects` → `get_graph_schema` → `search_graph` → `get_code_snippet`

> If the repository hasn't been indexed yet, run: `codebase-memory-mcp cli index_repository '{"repo_path": "/path/to/repo"}'`
