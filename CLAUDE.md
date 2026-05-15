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

# 跑单元测试
mvn test -Pwith-java -pl '!website' -DskipITs=true -T 1C

# 仅 plc4j 子树（只在已经跑过完整 build、~/.m2 fresh 的前提下用）
cd plc4j && mvn install -DskipTests -T 1C
```

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

## 工具链

- Java 21（Temurin）+ Maven 3.9+

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **plc4x** (24105 symbols, 65721 relationships, 300 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

> If any GitNexus tool warns the index is stale, run `npx gitnexus analyze` in terminal first.

## Always Do

- **MUST run impact analysis before editing any symbol.** Before modifying a function, class, or method, run `gitnexus_impact({target: "symbolName", direction: "upstream"})` and report the blast radius (direct callers, affected processes, risk level) to the user.
- **MUST run `gitnexus_detect_changes()` before committing** to verify your changes only affect expected symbols and execution flows.
- **MUST warn the user** if impact analysis returns HIGH or CRITICAL risk before proceeding with edits.
- When exploring unfamiliar code, use `gitnexus_query({query: "concept"})` to find execution flows instead of grepping. It returns process-grouped results ranked by relevance.
- When you need full context on a specific symbol — callers, callees, which execution flows it participates in — use `gitnexus_context({name: "symbolName"})`.

## Never Do

- NEVER edit a function, class, or method without first running `gitnexus_impact` on it.
- NEVER ignore HIGH or CRITICAL risk warnings from impact analysis.
- NEVER rename symbols with find-and-replace — use `gitnexus_rename` which understands the call graph.
- NEVER commit changes without running `gitnexus_detect_changes()` to check affected scope.

## Resources

| Resource | Use for |
|----------|---------|
| `gitnexus://repo/plc4x/context` | Codebase overview, check index freshness |
| `gitnexus://repo/plc4x/clusters` | All functional areas |
| `gitnexus://repo/plc4x/processes` | All execution flows |
| `gitnexus://repo/plc4x/process/{name}` | Step-by-step execution trace |

## CLI

| Task | Read this skill file |
|------|---------------------|
| Understand architecture / "How does X work?" | `.claude/skills/gitnexus/gitnexus-exploring/SKILL.md` |
| Blast radius / "What breaks if I change X?" | `.claude/skills/gitnexus/gitnexus-impact-analysis/SKILL.md` |
| Trace bugs / "Why is X failing?" | `.claude/skills/gitnexus/gitnexus-debugging/SKILL.md` |
| Rename / extract / split / refactor | `.claude/skills/gitnexus/gitnexus-refactoring/SKILL.md` |
| Tools, resources, schema reference | `.claude/skills/gitnexus/gitnexus-guide/SKILL.md` |
| Index, status, clean, wiki CLI commands | `.claude/skills/gitnexus/gitnexus-cli/SKILL.md` |

<!-- gitnexus:end -->
