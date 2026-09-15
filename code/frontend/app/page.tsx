import PersistedEditableGreeting from "../components/PersistedEditableGreeting";
import { getGreeting } from "../lib/mock/persisted-editable-greeting";

export default function Page() {
  const { greeting } = getGreeting();

  return <PersistedEditableGreeting initialGreeting={greeting} />;
}
