export type GreetingResponse = {
  greeting: string;
};

let storedGreeting = "Hello, World!";

export function getGreeting(): GreetingResponse {
  return { greeting: storedGreeting };
}

export function saveGreeting(greeting: string): GreetingResponse | null {
  const nextGreeting = greeting.trim();

  if (!nextGreeting || nextGreeting.length > 200) {
    return null;
  }

  storedGreeting = nextGreeting;
  return { greeting: storedGreeting };
}
