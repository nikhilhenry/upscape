export default function useDragSort() {
  let dragStartIndex: Number, dragEndIndex: Number;

  // list components
  function dragStart(event: any, index: Number) {
    event.dataTransfer.dropEffect = "move";
    event.dataTransfer.effectAllowed = "move";

    dragStartIndex = index;
  }

  // li funcitons
  // 'dragover'
  function dragOver() {}
  // 'drop'
  function dragDrop(event: any, index: Number) {
    dragEndIndex = index;
    swapItems(dragStartIndex, dragEndIndex);
  }
  // 'dragenter'
  function dragEnter() {}
  // 'dragleave'
  function dragLeave() {}

  function swapItems(fromIndex: Number, toIndex: Number) {
    console.log(fromIndex, toIndex);
  }

  return {
    dragStart,
    dragOver,
    dragDrop,
    dragEnter,
    dragLeave,
  };
}
