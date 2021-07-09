import axios from "axios";

export default async function(page = 1, sort = "desc", field = "created_at") {
  const response = await axios
    .get(`/api/objective?page=${page}&limit=6&sort=${sort}&field=${field}`)
    .then((res) => {
      return res.data;
    })
    .catch((err) => {
      return err.response.data;
    });

  return response;
}
