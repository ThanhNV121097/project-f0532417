# Story — Persisted editable greeting

Module: `greeting`
Plan item: Persisted editable greeting

## User story

As a Visitor, I want to view and save one shared greeting, so that the greeting shown on the page comes from persisted PostgreSQL data and remains after reload.

## In scope

- Load the one global stored greeting through the Go API and show it in the Next.js page.
- Seed or recreate the initial greeting value as `Hello, World!` when no stored value exists.
- Show the current greeting as the only large centered `h1` heading.
- Pre-fill one single-line text input with the current greeting.
- Save a changed greeting with the `Save` button or native form submit.
- Trim saved input before persistence.
- Accept trimmed greeting text from 1 to 200 characters.
- Reject empty, whitespace-only, and over-200-character greetings without changing stored value.
- Persist accepted changes in PostgreSQL so reload shows the latest saved greeting.
- Match approved one-section minimal greeting editor screen and design tokens.

## Out of scope

- Sign-in, roles, and per-visitor greetings — stakeholder specified no sign-in; greeting is global.
- Navigation and extra page sections — approved design has one centered section only.
- Multiple greetings, greeting history, audit trail, and undo — SRS scope is one latest stored greeting.
- External services — stakeholder specified none.
- Error banners, toast messages, loading indicators, disabled states, and active button states — approved design does not include these states.
- Conflict resolution beyond last completed save wins — architecture and SRS define latest save as winner.

## UI scope

- Screen: single `Greeting editor` page from approved design.
- Region: one centered `main` with `aria-label="Greeting editor"`.
- Default state: white page, black large centered greeting heading, one visually hidden `Greeting` label, one text input named `greeting`, and one blue `Save` submit button.
- Hover state: `Save` button uses darker blue `#1D4ED8` on pointer hover.
- Focus state: input and button show visible blue focus-visible outline.
- Compact layout: at widths up to 520px, input and button stack vertically and fit without horizontal scroll.
- Invalid submit state: no error text; existing heading stays unchanged and input receives focus.

## Acceptance criteria

- SC-1 [GREETING-001 AC-1]: Given no greeting has been changed since install, when Visitor opens page, then heading text is `Hello, World!`.
- SC-2 [GREETING-001 AC-2]: Given stored greeting is `Pipeline accepted`, when Visitor opens page, then heading text is `Pipeline accepted`.
- SC-3 [GREETING-001 AC-3]: Given stored greeting is `Pipeline accepted`, when Visitor opens page, then text input value is `Pipeline accepted`.
- SC-4 [GREETING-002 AC-1]: Given page shows `Hello, World!`, when Visitor enters `Saved value` and submits form, then heading text becomes `Saved value`.
- SC-5 [GREETING-002 AC-2]: Given Visitor has saved `Saved value`, when Visitor reloads page, then heading text is `Saved value`.
- SC-6 [GREETING-002 AC-3]: Given Visitor enters `  Saved value  `, when Visitor submits form, then stored greeting is `Saved value`.
- SC-7 [GREETING-002 AC-4]: Given page shows `Current value`, when Visitor submits empty or whitespace-only greeting, then heading remains `Current value` and input receives focus.
- SC-8 [GREETING-003 AC-1]: Given Visitor opens page, when page renders, then there is exactly one main region labelled `Greeting editor`.
- SC-9 [GREETING-003 AC-2]: Given Visitor opens page, when page renders, then there is exactly one `h1` showing current greeting.
- SC-10 [GREETING-003 AC-3]: Given Visitor opens page, when page renders, then there is one text input with accessible label `Greeting`.
- SC-11 [GREETING-003 AC-4]: Given Visitor opens page, when page renders, then there is one submit button labelled `Save`.
- SC-12 [GREETING-003 AC-5]: Given Visitor opens page, when page renders, then button background color is `#2563EB` in default state.
- SC-13 [GREETING-003 AC-6]: Given Visitor hovers `Save` button with pointer device, when hover state renders, then button background color is `#1D4ED8`.
- SC-14 [GREETING-003 AC-7]: Given Visitor opens page at 320px viewport width, when page renders, then page has no horizontal scroll and form controls fit within viewport.
- SC-15 [GREETING-003 AC-8]: Given Visitor tabs through controls, when input or button receives keyboard focus, then focus-visible outline is blue and visible.

## Dependencies

- PostgreSQL exists and is reachable through `DATABASE_URL` for greeting persistence.
- Backend exposes singleton greeting contract: `GET /v1/greeting` and `PUT /v1/greeting`.
- Backend stores one global greeting row with last completed save wins.
- Frontend reads API base URL from `NEXT_PUBLIC_API_URL`.
- Next.js page scaffold exists and mounts story component from `code/frontend/components/`.
- No external accounts, credentials, or stakeholder decisions required.
