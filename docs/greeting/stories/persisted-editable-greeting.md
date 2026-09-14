# Story — Persisted editable greeting

Module: `greeting`
Plan item: Persisted editable greeting

## User story

As a Visitor, I want to view and save one shared greeting, so that the page reflects persisted data after reload.

## In scope

- Load one global greeting from PostgreSQL through the Go API and show it on the single Next.js page.
- Seed or recreate initial greeting text as `Hello, World!` when no greeting has been changed since install or stored row is absent.
- Show current greeting as large centered `h1` and pre-fill the text input with same value.
- Let any Visitor save a changed greeting through form submit.
- Trim saved greeting, accept 1 to 200 characters, reject empty, whitespace-only, or over-200-character values without changing stored greeting.
- Persist latest completed save so reload shows latest greeting.
- Match approved one-section minimal screen and design-system tokens.

## Out of scope

- Sign-in, roles, per-user greetings, or permissions; SRS says no sign-in and one shared greeting.
- Multiple greetings, greeting history, versioning, or conflict UI; last completed save wins.
- Navigation, extra sections, loading states, or error-message UI; approved design has one centered editor only.
- External services; stakeholder specified none.
- Rich text, multiline input, localization, or greeting templates; story stores one plain text value exactly after trimming.

## UI scope

- Screen: Greeting editor only, matching approved `main aria-label="Greeting editor"` design.
- Default state: white page, centered main region, one large centered `h1`, one visually hidden `Greeting` label, one single-line text input, one `Save` submit button.
- Interaction states: button hover uses `#1D4ED8`; input and button focus-visible outlines use blue token; invalid empty or whitespace submit focuses input and keeps existing heading.
- Responsive state: at viewport widths up to 520px, input and button stack vertically, fill available width, and avoid horizontal scroll.

## Acceptance criteria

- SC-1 [GREETING-001 AC-1]: Given no greeting has been changed since install, when Visitor opens page, heading text is `Hello, World!`.
- SC-2 [GREETING-001 AC-2]: Given stored greeting is `Pipeline accepted`, when Visitor opens page, heading text is `Pipeline accepted`.
- SC-3 [GREETING-001 AC-3]: Given stored greeting is `Pipeline accepted`, when Visitor opens page, text input value is `Pipeline accepted`.
- SC-4 [GREETING-002 AC-1]: Given page shows `Hello, World!`, when Visitor enters `Saved value` and submits form, heading text becomes `Saved value`.
- SC-5 [GREETING-002 AC-2]: Given Visitor has saved `Saved value`, when Visitor reloads page, heading text is `Saved value`.
- SC-6 [GREETING-002 AC-3]: Given Visitor enters `  Saved value  `, when Visitor submits form, stored greeting is `Saved value`.
- SC-7 [GREETING-002 AC-4]: Given page shows `Current value`, when Visitor submits empty or whitespace-only greeting, heading remains `Current value` and input receives focus.
- SC-8 [GREETING-003 AC-1]: Given Visitor opens page, when page renders, there is exactly one main region labelled `Greeting editor`.
- SC-9 [GREETING-003 AC-2]: Given Visitor opens page, when page renders, there is exactly one `h1` showing current greeting.
- SC-10 [GREETING-003 AC-3]: Given Visitor opens page, when page renders, there is one text input with accessible label `Greeting`.
- SC-11 [GREETING-003 AC-4]: Given Visitor opens page, when page renders, there is one submit button labelled `Save`.
- SC-12 [GREETING-003 AC-5]: Given Visitor opens page, when page renders, button background color is `#2563EB` in default state.
- SC-13 [GREETING-003 AC-6]: Given Visitor hovers `Save` button with pointer device, when hover state renders, button background color is `#1D4ED8`.
- SC-14 [GREETING-003 AC-7]: Given Visitor opens page at 320px viewport width, when page renders, page has no horizontal scroll and form controls fit within viewport.
- SC-15 [GREETING-003 AC-8]: Given Visitor tabs through controls, when input or button receives keyboard focus, focus-visible outline is blue and visible.

## Dependencies

- PostgreSQL database exists and is reachable by backend through `DATABASE_URL`.
- Go API implements `GET /v1/greeting` and `PUT /v1/greeting` with singleton greeting row and shared JSON error envelope.
- Backend migrations seed or recreate singleton greeting with initial value `Hello, World!`.
- Next.js frontend reads API base from `NEXT_PUBLIC_API_URL`.
- No external accounts, credentials, or stakeholder decisions required.
