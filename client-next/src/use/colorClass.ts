import { ComputedRef } from "vue";

export default function useColorClass(totalDuration: ComputedRef<string>) {
  const duration = parseFloat(totalDuration.value);
  // colors class

  if (duration < 2) return "light";
  if (duration < 4.5) return "medium";
  if (duration < 6) return "medium-heavy";
  if (duration >= 6) return "text-cyan-300";
  return "none";
}
