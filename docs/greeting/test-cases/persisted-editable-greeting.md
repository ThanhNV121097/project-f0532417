# Test cases — Persisted editable greeting

Module: `greeting`
Function: Persisted editable greeting
Risk level: Medium. Story writes shared persisted state and must prove full frontend, backend, and PostgreSQL path; scope has one actor and one singleton record, so no role matrix beyond public Visitor.

## Page behaviour and design

**Scenario**: Initial install shows default greeting heading
**Given**: No greeting has been changed since install and the singleton row is absent or contains the installed value
**When**: Visitor opens the greeting page
**Then**: Exactly one `h1` heading shows `Hello, World!`
Traces: SC-1 (GREETING-001 AC-1)
Check: render_url

**Scenario**: Stored greeting shows in heading on load
**Given**: Stored greeting is `Pipeline accepted`
**When**: Visitor opens the greeting page
**Then**: Exactly one `h1` heading shows `Pipeline accepted`
Traces: SC-2 (GREETING-001 AC-2), SC-9 (GREETING-003 AC-2)
Check: render_url

**Scenario**: Stored greeting pre-fills input on load
**Given**: Stored greeting is `Pipeline accepted`
**When**: Visitor opens the greeting page
**Then**: Text input labelled `Greeting` has value `Pipeline accepted`
Traces: SC-3 (GREETING-001 AC-3)
Check: render_url

**Scenario**: Save button updates heading
**Given**: Page shows `Hello, World!`
**When**: Visitor enters `Saved value` in the `Greeting` input and submits the form with `Save`
**Then**: Heading text becomes `Saved value` and input value is `Saved value`
Traces: SC-4 (GREETING-002 AC-1)
Check: interact_page

**Scenario**: Saved greeting persists after reload
**Given**: Visitor has saved `Saved value`
**When**: Visitor reloads the page
**Then**: Heading text is `Saved value` and input value is `Saved value`
Traces: SC-5 (GREETING-002 AC-2)
Check: interact_page

**Scenario**: Surrounding whitespace is trimmed before save
**Given**: Page is open and current stored greeting is any valid value
**When**: Visitor enters `  Saved value  ` in the `Greeting` input and submits the form
**Then**: Heading text becomes `Saved value`; after reload, heading text is `Saved value`
Traces: SC-6 (GREETING-002 AC-3)
Check: interact_page

**Scenario**: Whitespace-only greeting is rejected without losing current heading
**Given**: Page shows `Current value`
**When**: Visitor enters `   ` in the `Greeting` input and submits the form
**Then**: Heading remains `Current value`; focus is on the `Greeting` input; no error text is shown
Traces: SC-7 (GREETING-002 AC-4)
Check: interact_page

**Scenario**: Empty greeting is rejected without losing current heading
**Given**: Page shows `Current value`
**When**: Visitor clears the `Greeting` input and submits the form
**Then**: Heading remains `Current value`; focus is on the `Greeting` input; no error text is shown
Traces: SC-7 (GREETING-002 AC-4)
Check: interact_page

**Scenario**: Page exposes one labelled main region
**Given**: Visitor opens the page
**When**: Page renders
**Then**: There is exactly one `main` region with accessible label `Greeting editor`
Traces: SC-8 (GREETING-003 AC-1)
Check: render_url

**Scenario**: Greeting input has required accessible contract
**Given**: Visitor opens the page
**When**: Page renders
**Then**: There is one single-line text input with accessible label `Greeting`, name `greeting`, `required` set, and `autocomplete` off
Traces: SC-10 (GREETING-003 AC-3)
Check: render_url

**Scenario**: Save submit button is present
**Given**: Visitor opens the page
**When**: Page renders
**Then**: There is one submit button labelled `Save`
Traces: SC-11 (GREETING-003 AC-4)
Check: render_url

**Scenario**: Button default color matches design
**Given**: Visitor opens the page
**When**: Page renders without pointer hover
**Then**: `Save` button computed background color is `#2563EB`
Traces: SC-12 (GREETING-003 AC-5)
Check: measure_styles

**Scenario**: Button hover color matches design
**Given**: Visitor opens the page on a pointer-capable viewport
**When**: Visitor hovers the `Save` button
**Then**: `Save` button computed background color is `#1D4ED8`
Traces: SC-13 (GREETING-003 AC-6)
Check: measure_styles

**Scenario**: Compact viewport fits controls without horizontal scroll
**Given**: Visitor opens the page at 320px viewport width
**When**: Page renders
**Then**: Document has no horizontal scroll; the text input and `Save` button each fit within the viewport width
Traces: SC-14 (GREETING-003 AC-7)
Check: measure_styles

**Scenario**: Keyboard focus outline is visible and blue
**Given**: Visitor opens the page
**When**: Visitor tabs to the `Greeting` input and then to the `Save` button
**Then**: Each focused control has a visible blue focus outline while focused
Traces: SC-15 (GREETING-003 AC-8)
Check: interact_page

**Scenario**: Page has one section only and no navigation
**Given**: Visitor opens the page
**When**: Page renders
**Then**: Page contains no navigation landmarks and no content section besides the labelled `Greeting editor` main region
Traces: SC-8 (GREETING-003 AC-1)
Check: render_url

**Scenario**: Page uses minimal static styling with no animation
**Given**: Visitor opens the page
**When**: Page renders
**Then**: Page background is `#FFFFFF`, text color is `#000000`, and no rendered element has running CSS animation or transition duration greater than `0s`
Traces: SC-12 (GREETING-003 AC-5)
Check: measure_styles

## Boundary, conflict, and recovery behaviour

**Scenario**: One-character greeting is accepted and persisted
**Given**: Page is open and current stored greeting is `Current value`
**When**: Visitor enters `A` in the `Greeting` input, submits the form, and reloads the page
**Then**: Heading text is `A`
Traces: GREETING-002 boundary
Check: interact_page

**Scenario**: Two-hundred-character greeting is accepted and persisted
**Given**: Page is open and current stored greeting is `Current value`
**When**: Visitor enters a 200-character greeting, submits the form, and reloads the page
**Then**: Heading text equals the same 200-character greeting
Traces: GREETING-002 boundary
Check: interact_page

**Scenario**: Over-limit greeting is rejected without changing stored value
**Given**: Page shows `Current value`
**When**: Visitor enters a 201-character greeting and submits the form
**Then**: Heading remains `Current value`; after reload, heading text is `Current value`; no error text is shown
Traces: GREETING-002 boundary
Check: interact_page

**Scenario**: Missing stored row is recreated without not-found screen
**Given**: The singleton greeting row is absent
**When**: Visitor opens the greeting page
**Then**: Heading text is `Hello, World!` and no not-found screen is shown
Traces: GREETING-001 not found
Check: render_url

**Scenario**: Unsigned visitor can view and save greeting
**Given**: Visitor has no sign-in session and page shows `Hello, World!`
**When**: Visitor enters `Public save` and submits the form
**Then**: Heading text becomes `Public save`; no sign-in prompt is shown
Traces: GREETING-002 permission
Check: interact_page

**Scenario**: Last completed save wins across two visitors
**Given**: Visitor A and Visitor B both have the greeting page open with heading `Start value`
**When**: Visitor A saves `First save`, then Visitor B saves `Second save`, and Visitor A reloads
**Then**: Visitor A sees heading `Second save`
Traces: GREETING-002 conflict
Check: interact_page

**Scenario**: Save API failure preserves displayed current value
**Given**: Page shows `Current value` and the API or database becomes unavailable before save
**When**: Visitor enters `Lost save` and submits the form
**Then**: Heading remains `Current value`; no error text is shown
Traces: GREETING-002 upstream failure
Check: interact_page

## Service contract cases

**Scenario**: GET greeting returns current greeting JSON
**Given**: Stored greeting is `Hello, World!`
**When**: Client sends `GET /api/v1/greeting` with no request body
**Then**: Response status is `200` and body is exactly `{"greeting":"Hello, World!"}`
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting recreates absent singleton row
**Given**: The singleton greeting row is absent
**When**: Client sends `GET /api/v1/greeting`
**Then**: Response status is `200` and body is exactly `{"greeting":"Hello, World!"}`; a following `GET /api/v1/greeting` returns the same body
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting database failure returns internal error envelope
**Given**: Database query for greeting fails after connection succeeds
**When**: Client sends `GET /api/v1/greeting`
**Then**: Response status is `500`; body is `{"error":{"code":"INTERNAL_ERROR","message":"Internal error."}}` or another safe plain-text message with code `INTERNAL_ERROR`; no database detail appears
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting saves trimmed valid text
**Given**: Stored greeting is `Hello, World!`
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"greeting":"  Saved value  "}`
**Then**: Response status is `200` and body is exactly `{"greeting":"Saved value"}`; a following `GET /api/v1/greeting` returns `{"greeting":"Saved value"}`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting accepts one-character boundary
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"greeting":"A"}`
**Then**: Response status is `200` and body is exactly `{"greeting":"A"}`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting accepts two-hundred-character boundary
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with JSON body containing `greeting` as exactly 200 characters
**Then**: Response status is `200` and response `greeting` equals the 200-character value
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects malformed JSON
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with malformed JSON body `{`
**Then**: Response status is `400`; error code is `INVALID_REQUEST`; a following `GET /api/v1/greeting` returns `Current value`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects missing greeting field
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{}`
**Then**: Response status is `400`; error code is `INVALID_REQUEST`; a following `GET /api/v1/greeting` returns `Current value`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects non-string greeting
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"greeting":123}`
**Then**: Response status is `400`; error code is `INVALID_REQUEST`; a following `GET /api/v1/greeting` returns `Current value`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects undefined request field
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"greeting":"New value","unexpected":true}`
**Then**: Response status is `400`; error code is `INVALID_REQUEST`; a following `GET /api/v1/greeting` returns `Current value`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects empty text
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"greeting":""}`
**Then**: Response status is `422`; error code is `INVALID_GREETING`; a following `GET /api/v1/greeting` returns `Current value`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects whitespace-only text
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"greeting":"   "}`
**Then**: Response status is `422`; error code is `INVALID_GREETING`; a following `GET /api/v1/greeting` returns `Current value`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects over-limit text
**Given**: Stored greeting is `Current value`
**When**: Client sends `PUT /api/v1/greeting` with JSON body containing `greeting` as 201 characters after trimming
**Then**: Response status is `422`; error code is `INVALID_GREETING`; a following `GET /api/v1/greeting` returns `Current value`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting database failure returns internal error envelope
**Given**: Stored greeting is `Current value` and database write fails after connection succeeds
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"greeting":"New value"}`
**Then**: Response status is `500`; error code is `INTERNAL_ERROR`; message is safe plain text with no database detail; stored greeting remains `Current value`
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: Health check succeeds when database is ready
**Given**: Migration succeeded and database `SELECT 1` succeeds
**When**: Client sends `GET /api/healthz`
**Then**: Response status is `200` and body is exactly `{"status":"ok"}`
Traces: contract (GET /healthz)
Check: fetch_url

**Scenario**: Health check reports unavailable when database cannot be reached
**Given**: Database refuses connection or migration has not succeeded
**When**: Client sends `GET /api/healthz`
**Then**: Response status is `503`; body uses error envelope with code `SERVICE_UNAVAILABLE` and safe plain-text message
Traces: contract (GET /healthz)
Check: fetch_url
