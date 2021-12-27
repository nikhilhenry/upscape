export default function useDragSort() {
  // list components
  function dragStart(event: any, index: Number) {
    event.dataTransfer.dropEffect = "move";
    event.dataTransfer.effectAllowed = "move";

    console.log(index);
  }

  // li funcitons
  // 'dragover'
  function dragOver() {}
  // 'drop'
  function dragDrop() {}
  // 'dragenter'
  function dragEnter() {}
  // 'dragleave'
  function dragLeave() {}

  return {
    dragStart,
    dragOver,
    dragDrop,
    dragEnter,
    dragLeave,
  };
}
