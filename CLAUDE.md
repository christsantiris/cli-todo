# cli-todo — Claude Coding Rules

## 1. Always use braces on control flow

Prefer:
```go
if err != nil {
    return err
}
```

Never:
```go
if err != nil { return err }
```

This applies to all control flow: `if`, `else`, `for`, `switch`.

---

## 2. Single space around operators — no alignment padding

Prefer:
```go
a := 5
bc := 10
def := 15
```

Never:
```go
a   := 5
bc  := 10
def := 15
```

This applies to variable declarations, struct field assignments, and all `=` / `:=` usage.

---

## 3. Function signatures on one line

Prefer:
```go
func (t *Todos) CompleteItem(index int) error {
```

Never:
```go
func (t *Todos) CompleteItem(
    index int) error {
```

If a signature is too long, shorten parameter names rather than wrapping.

---

## 4. Minimal code per change set

- Provide the smallest possible change per step
- Wait for an explicit "next step" prompt before providing the next change
- Never combine multiple features into one code block
- Never add unrequested changes alongside a requested change

---

## 5. Explain each change

For every code change, briefly explain:
- What the change does
- Why it is needed
- Any side effects or dependencies

---

## 6. Be maximally specific about code location

When instructing where to add or modify code, always specify:
- The exact filename and path
- The exact function name
- The exact surrounding lines to find the insertion point

Prefer:

> In `cmd/todo/main.go` in `main()`, after `todos.StoreAddedItem(todoFile)` in the `*add` case, add:
> ```go
> todos.PrintToDos()
> ```

Never:

> Add this after the store call.

---

## 7. All command additions or changes must be documented in README.md

When adding a new flag or modifying an existing command's behaviour, update the Commands table in `README.md` in the same change set. The table must stay the authoritative reference for how to use the tool.

---

## 8. All model changes must be reflected in storage

Any new field added to `Todo` in `todo.go` must also be:
- Serialized correctly (verify JSON tags are present)
- Handled in `PrintToDos()` in `todos.go` if it should appear in output

Omitting this causes data to silently drop or display incorrectly.
