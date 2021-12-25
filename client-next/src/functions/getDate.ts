function getFormattedDate(date: Date) {
  const formattedDate =
    date.toLocaleString("default", { month: "short" }).toUpperCase() +
    " " +
    new Date(date).getDate() +
    " " +
    new Date(date).getFullYear();
  return formattedDate;
}

function getQueryDate(date: Date) {
  let month = new Date(date).getMonth() + 1;
  let day = new Date(date).getDate();
  const year = new Date(date).getFullYear().toString().substr(-2);

  let monthNum, dayNum;

  if (day < 10) {
    dayNum = "0" + day;
  }

  if (month < 10) {
    monthNum = "0" + month;
  }

  const queryDate = monthNum + "/" + dayNum + "/" + year;
  return queryDate;
}

const getDate = function (dateIndex: number) {
  if (typeof dateIndex != "number") return null;
  if (dateIndex > 1) dateIndex = 1;

  const formattedDate = [];

  // today with index 0
  if (dateIndex == 0) {
    formattedDate[0] = "today";
    formattedDate[1] = getFormattedDate(new Date());
    formattedDate[2] = "today";

    return formattedDate;
  }

  const today = new Date();
  const date = new Date(today);

  // tomorrow with index 1
  if (dateIndex > 0) {
    formattedDate[0] = "tomorrow";
    date.setDate(date.getDate() + 1);
    formattedDate[1] = getFormattedDate(date);
    formattedDate[2] = "tomorrow";

    return formattedDate;
  }

  // get date from index
  date.setDate(date.getDate() + dateIndex);

  formattedDate[0] = date.toLocaleString("default", { weekday: "long" });
  formattedDate[1] = getFormattedDate(date);
  formattedDate[2] = getQueryDate(date);

  return formattedDate;
};

export default getDate;
