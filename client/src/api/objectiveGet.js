import axios from "axios";

export default async function(page = 1) {
  const response = await axios
    .get(`/api/objective?page=${page}&limit=4`)
    .then((res) => {
      return res.data;
    })
    .catch((err) => {
      return err.response.data;
    });

  return response;
}
