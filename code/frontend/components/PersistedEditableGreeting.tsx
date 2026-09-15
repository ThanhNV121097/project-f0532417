"use client";

import { FormEvent, useRef, useState } from "react";
import { saveGreeting } from "../lib/mock/persisted-editable-greeting";
import styles from "./PersistedEditableGreeting.module.css";

type PersistedEditableGreetingProps = {
  initialGreeting: string;
};

export default function PersistedEditableGreeting({
  initialGreeting,
}: PersistedEditableGreetingProps) {
  const [greeting, setGreeting] = useState(initialGreeting);
  const [inputValue, setInputValue] = useState(initialGreeting);
  const inputRef = useRef<HTMLInputElement>(null);

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const response = saveGreeting(inputValue);

    if (!response) {
      inputRef.current?.focus();
      return;
    }

    setGreeting(response.greeting);
    setInputValue(response.greeting);
  }

  return (
    <main aria-label="Greeting editor" className={styles.editor}>
      <h1 className={styles.heading}>{greeting}</h1>
      <form className={styles.form} noValidate onSubmit={handleSubmit}>
        <label className={styles.label} htmlFor="greeting-input">
          Greeting
        </label>
        <input
          ref={inputRef}
          id="greeting-input"
          name="greeting"
          type="text"
          value={inputValue}
          required
          autoComplete="off"
          onChange={(event) => setInputValue(event.target.value)}
          className={styles.input}
        />
        <button className={styles.button} type="submit">
          Save
        </button>
      </form>
    </main>
  );
}
