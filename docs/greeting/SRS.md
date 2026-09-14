# SRS — Greeting

Module: `greeting`
Design: [View the approved design](http://localhost:8080/design/f0532417-c6f6-4c4d-8dcf-3a7321363ae4)
Design system: `design/design-system.md`

> One file per module, at `docs/{module}/SRS.md`. It covers only the functions
> that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

Greeting module lets any visitor view and change one persisted greeting on the single page for "Hello World Acceptance 4". It proves the full pipeline end to end: PostgreSQL persistence, Go API access, and Next.js display. Without it, the product becomes a static page and cannot prove saved data survives reload.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Visitor | Anyone opening the public page, with no sign-in | View the current greeting, edit the greeting text, and save it |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Persisted editable greeting

**Out of scope** — name what a reader would reasonably expect here and say
where it lives instead. This section prevents the same argument twice.

- Sign-in and roles — deliberately not built; stakeholder specified no sign-in.
- Navigation and other page sections — deliberately not built; approved design has one centered section only.
- External services — deliberately not built; stakeholder specified no external services.
- Multiple greetings or greeting history — deliberately not built; scope is one stored greeting.

## 4. Functional requirements

### 4.1 Persisted editable greeting

**Requirement GREETING-001 — Load stored greeting**

*As a* Visitor, *I want to* see the current stored greeting when I open the page, *so that* the page reflects persisted data.

Behaviour:

1. Visitor opens the greeting page.
2. Page obtains the current greeting through the product API.
3. Page shows the greeting as the large centered `h1` heading.
4. Page pre-fills the text input with the same greeting.
5. First installed value is `Hello, World!` until a Visitor saves a different non-empty value.

**Acceptance criteria** — each is proved by at least one test case in
`docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it
(`SC-1 [GREETING-001 AC-1]`). Given/When/Then, no compound
conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | No greeting has been changed since install | Visitor opens the page | Heading text is `Hello, World!` |
| AC-2 | Stored greeting is `Pipeline accepted` | Visitor opens the page | Heading text is `Pipeline accepted` |
| AC-3 | Stored greeting is `Pipeline accepted` | Visitor opens the page | Text input value is `Pipeline accepted` |

**Requirement GREETING-002 — Save changed greeting**

*As a* Visitor, *I want to* save a changed greeting, *so that* the new greeting remains after reload.

Behaviour:

1. Visitor changes the text input value.
2. Visitor submits the form with the `Save` button or native form submit.
3. If trimmed input is non-empty, product API stores the new greeting.
4. Page shows the saved greeting in the heading.
5. Page keeps the text input value equal to the saved greeting.
6. Reloading the page shows the saved greeting again.

**Acceptance criteria** — each is proved by at least one test case in
`docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it
(`SC-1 [GREETING-002 AC-1]`). Given/When/Then, no compound
conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Page shows `Hello, World!` | Visitor enters `Saved value` and submits the form | Heading text becomes `Saved value` |
| AC-2 | Visitor has saved `Saved value` | Visitor reloads the page | Heading text is `Saved value` |
| AC-3 | Visitor enters `  Saved value  ` | Visitor submits the form | Stored greeting is `Saved value` |
| AC-4 | Page shows `Current value` | Visitor submits an empty or whitespace-only greeting | Heading remains `Current value` and input receives focus |

**Requirement GREETING-003 — Match approved screen**

*As a* Visitor, *I want to* use the approved greeting editor screen, *so that* the product matches the accepted design.

Behaviour:

1. Page uses one centered `main` region labelled `Greeting editor`.
2. Page shows one large centered `h1` greeting.
3. Page shows one form below the heading.
4. Form includes one visually hidden label with text `Greeting` associated to the text input.
5. Form includes one single-line text input named `greeting`, required, with autocomplete off.
6. Form includes one submit button labelled `Save`.
7. Page has white background, black text, blue `#2563EB` primary button, darker blue `#1D4ED8` primary button hover state, no navigation, and no animation.
8. At widths up to 520px, input and button stack vertically and fill available width.

**Acceptance criteria** — each is proved by at least one test case in
`docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it
(`SC-1 [GREETING-003 AC-1]`). Given/When/Then, no compound
conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Visitor opens the page | Page renders | There is exactly one main region labelled `Greeting editor` |
| AC-2 | Visitor opens the page | Page renders | There is exactly one `h1` showing the current greeting |
| AC-3 | Visitor opens the page | Page renders | There is one text input with accessible label `Greeting` |
| AC-4 | Visitor opens the page | Page renders | There is one submit button labelled `Save` |
| AC-5 | Visitor opens the page | Page renders | Button background color is `#2563EB` in default state |
| AC-6 | Visitor hovers the `Save` button with a pointer device | Hover state renders | Button background color is `#1D4ED8` |
| AC-7 | Visitor opens the page at 320px viewport width | Page renders | Page has no horizontal scroll and form controls fit within viewport |
| AC-8 | Visitor tabs through controls | Input or button receives keyboard focus | Focus-visible outline is blue and visible |

**Failure, boundary and permission behaviour** — the part most often skipped
and most often the source of bugs. Every case this function actually has needs a
defined outcome; "should not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Greeting is empty or whitespace-only | Nothing is saved; existing heading stays unchanged; input receives focus. No error text is shown because approved design has no error state. |
| Boundary | Greeting length is 1 to 200 characters after trimming | Accepted and persisted. |
| Boundary | Greeting length is more than 200 characters after trimming | Rejected by API contract; nothing is saved; existing stored greeting and heading stay unchanged. No error text is shown because approved design has no error state. |
| Not found | Stored greeting row is absent | Product recreates or returns initial greeting `Hello, World!`; Visitor never sees a not-found screen. |
| Not permitted | Visitor is not signed in | Not applicable: module has no sign-in or permission differences. |
| Conflict | Two Visitors save different greetings | Last completed save wins; reload shows the latest stored greeting. |
| Upstream failure | API or database is unavailable during load or save | No error or loading state is part of the approved design; API contract error envelope is specified in service contract. |

**Data touched** — the fields this function reads and writes, in product terms.
The physical schema is TL's job in `docs/architecture/erd.md`; this is the list
that document has to satisfy.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Trimmed value must be 1 to 200 characters; initial value is `Hello, World!`; latest saved value persists across reload |

## 5. Screens

The design is the source of truth for appearance; this section maps functions
onto it so nothing in the design is unaccounted for and nothing specified here
is missing from the design.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting editor | `main aria-label="Greeting editor"` in approved design | GREETING-001, GREETING-002, GREETING-003 | default, button hover |

## 6. Non-functional requirements

| Area | Requirement |
|---|---|
| Accessibility | Keyboard reachable input and button; visible focus; input has accessible label `Greeting`; text/button contrast ratio is at least 4.5:1. |
| Responsive | Page works at 320px viewport width and up with no horizontal page scroll. |
| Localisation | Static UI copy is English: `Greeting`, `Save`; saved greeting text is displayed exactly as stored after trimming. |
| Privacy | No personal data is required; only one greeting text value is stored. |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL, for persisting the greeting across reloads.
- **Depends on:** Go API, for reading and saving the greeting.
- **Depends on:** Next.js frontend, for rendering the approved one-page screen.
- **Assumption:** One global greeting is shared by all Visitors. If per-user greetings are needed later, sign-in and ownership become new scope.

| Open question | Proposed default | Who decides |
|---|---|---|
| — | None; stakeholder already confirmed scope, design, and no external services. | — |

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Persisted editable greeting | GREETING-001, GREETING-002, GREETING-003 | `test-cases/persisted-editable-greeting.md` |
