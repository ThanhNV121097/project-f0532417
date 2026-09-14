# Design System — Hello World Acceptance 4

> Source of truth: approved `index.html`.
> Every value below is extracted from it. Changing a value here without changing approved design is defect.

Last updated: 2026-09-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background, input background, button text |
| `--color-text` | `#000000` | Heading text, input text, input border |
| `--color-primary-action` | `#2563EB` | Primary Save button background and border |
| `--color-primary-action-hover` | `#1D4ED8` | Primary Save button hover background and border |
| `--color-focus-ring` | `#2563EB` | Input and button focus-visible outline |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21.0:1` | AA |
| `--color-bg` | `--color-primary-action` | `5.2:1` | AA |
| `--color-bg` | `--color-primary-action-hover` | `6.7:1` | AA |
| `--color-text` border | `--color-bg` | `21.0:1` | UI AA |
| `--color-focus-ring` | `--color-bg` | `5.2:1` | UI AA |

### 1.2 Spacing

Base unit: `4px`. Margins, gaps, and layout padding use this scale.

| Token | Value |
|---|---|
| `--space-3` | `12px` |
| `--space-6` | `24px` |
| `--space-8` | `32px` |

Known one-off values from approved design: `14px` input horizontal padding, `22px` button horizontal padding, `1px` border, `3px` focus outline and offset, and visually-hidden label utility values (`1px`, `-1px`). Listed under Known deviations.

### 1.3 Typography

Font families:

- Body: `Arial, Helvetica, sans-serif`, loaded as system fonts.
- Headings: `Arial, Helvetica, sans-serif`, inherited from body.

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-control` | `16px` | normal browser line-height | `400` input, `700` button | Input and button text |
| `--text-greeting` | `clamp(48px, 10vw, 88px)` | `1` | `700` | h1 greeting |

Heading levels are used in order: one `h1` only.

| Token | Value | Used for |
|---|---|---|
| `--font-weight-body` | `400` | Input text and inherited body text |
| `--font-weight-strong` | `700` | h1 greeting and Save button |
| `--tracking-tight-greeting` | `-0.04em` | h1 greeting |

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-control` | `6px` | Input and Save button |
| `--border-width-control` | `1px` | Input and Save button border |
| `--focus-outline-width` | `3px` | Input and button focus-visible outline |
| `--focus-outline-offset` | `3px` | Input and button focus-visible offset |

No shadows appear in approved design.

No transition or animation appears in approved design. State changes are immediate.

### 1.5 Layout and breakpoints

| Name | Max width rule | Container | Columns | Gutter |
|---|---|---|---|---|
| `base` | default | `min(100%, 560px)` centered main | 1 | form gap `12px` |
| `compact` | `max-width: 520px` | `min(100%, 560px)` centered main | 1 stacked form controls | form gap `12px` |

Page layout uses viewport-centered grid on `body`, `min-height: 100vh`, and `padding: 24px`.

No z-index tokens: approved design uses no positioned stacking layers.

## 2. Components

### 2.1 Greeting editor

**Purpose** — Single page section for viewing current greeting and saving a changed greeting. Do not use for multi-section pages or navigation.

**Anatomy** — `[main region] [h1 greeting] [form] [visually-hidden label] [text input] [Save button]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Centered editor | `--color-bg`, `--color-text`, `--space-6`, `--space-8`, `--text-greeting` | One-page greeting edit flow |

**Sizes**

| Size | Width | Padding | Text token |
|---|---|---|---|
| Default | Main `min(100%, 560px)` | Body `24px` | `--text-greeting` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Large centered greeting above centered form | `--color-bg`, `--color-text`, `--space-6`, `--space-8`, `--text-greeting` |

**Accessibility** — `main` uses `aria-label="Greeting editor"`. Heading is real `h1`. Form submit updates greeting. No navigation appears.

### 2.2 Text input

**Purpose** — Enter replacement greeting. Do not use for long-form text.

**Anatomy** — `[visually-hidden label] [text input]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default text input | `--color-bg`, `--color-text`, `--radius-control`, `--border-width-control` | Single-line greeting entry |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `min-height: 48px` | `0 14px` | `--text-control` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White field, black text, black 1px border | `--color-bg`, `--color-text`, `--radius-control`, `--border-width-control` |
| Focus-visible | 3px blue outline with 3px offset | `--color-focus-ring`, `--focus-outline-width`, `--focus-outline-offset` |

**Accessibility** — Input has associated label text `Greeting`, visually hidden. Use `required` and `autocomplete="off"`. Minimum hit target is `48px` high. Keyboard focus uses visible `:focus-visible` outline.

### 2.3 Primary button

**Purpose** — Submit greeting change. Use for sole primary action in this flow.

**Anatomy** — `[label]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Primary action | `--color-primary-action`, `--color-bg`, `--radius-control`, `--border-width-control` | Main Save action |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `min-height: 48px` | `0 22px` | `--text-control` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Blue fill and border, white bold text | `--color-primary-action`, `--color-bg`, `--font-weight-strong` |
| Hover | Darker blue fill and border | `--color-primary-action-hover` |
| Focus-visible | 3px blue outline with 3px offset | `--color-focus-ring`, `--focus-outline-width`, `--focus-outline-offset` |

**Accessibility** — Native `button type="submit"`. Minimum hit target is `48px` high. Keyboard activation uses native Enter/Space behavior. Keyboard focus uses visible `:focus-visible` outline.

## 3. Content and formatting

- Voice and tone: plain, direct, product-neutral.
- Date, time, number, and currency formats: none used.
- Capitalization: heading preserves saved greeting exactly; button uses title case label `Save`; label uses sentence/title case `Greeting`.
- Empty-state and error-message wording pattern: none drawn. Empty input submit refocuses field and leaves existing greeting unchanged.

## 4. Known deviations

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Spacing | `14px` input horizontal padding and `22px` button horizontal padding are outside 4px scale | Approved design uses these exact control paddings | Keep unless stakeholder asks for tighter token scale |
| Focus ring | `3px` outline and `3px` offset are outside 4px spacing scale | Approved design uses visible blue focus ring | Keep for accessibility unless design is revised |
| Button states | No active or disabled state drawn | Approved design only draws default, hover, and focus-visible | Add only if later requirements introduce disabled or pressed button |
| Input states | No hover, invalid, disabled, loading, or error state drawn | Approved design only draws default and focus-visible | Add only if later requirements introduce validation UI or async status |
| AI defaults check | Minimal flat design avoids purple/indigo palette, gradients, heavy shadows, emoji, filler copy, and generic multi-section layout | Matches approved minimal project style | No action |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2026-09-14 | Initial design system extracted from approved `index.html` | This PR |
