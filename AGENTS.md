# Project Guidelines & Directives (`vk_google`)

This document summarizes the core directives, architectural patterns, and project rules established during development of the Vulkan Go bindings repository (`vk_google`).

---

## 1. Code Generation & Package Architecture (`cmd/vk-gen`)

- **Multi-API Output Packages**: Codegen must emit 3 separate API packages:
  - `vulkan`: Vulkan Core API (`api="vulkan"`)
  - `vulkanbase`: Vulkan Base API (`api="vulkanbase"`)
  - `vulkansc`: Vulkan Safety Critical API (`api="vulkansc"`)
- **File Naming & Suffixes**:
  - All generated code files MUST end with `_gen.go` (e.g. `types_gen.go`, `constants_gen.go`, `enums_gen.go`, `bitmasks_gen.go`, `structs_gen.go`, `commands_gen.go`, `metadata_gen.go`, `doc_gen.go`).
  - Handwritten helper files (`errors.go`, `loader.go`, `utils.go`) are emitted to all 3 API packages and must NOT contain "This file has been generated" header warnings.
- **Function Pointer Invocation Policy**:
  - **NEVER** check function pointers against `0` and return silently in generated command wrappers.
  - If a function pointer is uninitialized (nil/0), let execution panic/crash on null pointer invocation rather than swallowing errors silently.

---

## 2. API Branch Filtering & Indexing Rules

- **XML `api` Attribute Scoping**:
  - The `api` XML attribute applies across `<type>`, `<command>`, `<enum>`, `<member>`, and `<extension>` elements in `vk.xml`.
  - Maintain isolated index maps (`Handles`, `Bitmasks`, `Structs`, `EnumGroups`, `Constants`, `TypeAliases`, `FuncPointerDefs`) for each `ApiBranchInfo` to prevent SC or extension definitions from overwriting or corrupting Vulkan core/base definitions.
- **Strict Type Validation (`isValidBranchType`)**:
  - When emitting structs, command wrappers, callback types, and struct extender interfaces for a specific branch, validate that all referenced types (parameter types, struct member types, return types) actually exist within that API branch package.
  - Omit struct members or commands referencing types not supported in that API branch to ensure zero undefined symbols.

---

## 3. Extension Documentation & Parsing

- **AsciiDoc Overview Append**:
  - For extension packages (`extensions/<ext_name>`), parse the `.adoc` file from `Vulkan-Docs/appendices/<ext_name>.adoc`.
  - Strip all AsciiDoc syntax (including `include::`, `slink:`, `pname:`, `ename:`, `[eq]#`, block delimiters, and macro tags).
  - Append the cleaned, plain-text overview as package documentation in `doc_gen.go`.

---

## 4. Loader & Command Initialization API

- `Init()` in `loader.go` loads the Vulkan shared library (`libvulkan.so.1` / `vulkan-1.dll`).
- `InitCommands(instance Instance, device Device)` in `commands_gen.go` resolves instance and device procedure addresses without symbol collisions.

---

## 5. Sample & Application Conventions

- Keep samples (`samples/multi_draw_indirect`, `samples/base`) clean, minimal, and performant.
- Camera implementations handle view, projection, and combined projection-view matrix math.
- Utilize shader reflection where applicable to automatically configure descriptor set layouts, pipeline layouts, and shader modules with a thin API.
