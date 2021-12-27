import { computed } from "vue";

export default function useColorClass(duration: Number) {
  // colors class
  const colorClass = computed(() => {
    if (duration < 2) return "light";
    if (duration < 4.5) return "medium";
    if (duration < 6) return "medium-heavy";
    if (duration >= 6) return "heavy";
    return "none";
  });

  return colorClass;
}
